package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ghc-golang-hoangth7/finalprj/models"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

type FlightService struct {
	pb.UnimplementedFlightServiceServer
	planesSrv pbPlanes.PlanesServiceClient
	db        *sql.DB
}

func NewFlightService(db *sql.DB, planeSrv pbPlanes.PlanesServiceClient) *FlightService {
	boil.SetDB(db)
	boil.DebugMode = true
	boil.DebugWriter = os.Stdout
	return &FlightService{db: db, planesSrv: planeSrv}
}

func (s *FlightService) UpsertFlight(ctx context.Context, req *pb.Flight) (*pb.FlightId, error) {
	pbPlane, err := s.planesSrv.GetPlaneByNumber(ctx, &pbPlanes.PlaneNumber{
		PlaneNumber: req.PlaneNumber,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get plane info: %v", err)
	}

	if len(req.Id) == 0 {
		req.Id = uuid.New().String()
		if req.ScheduledDepartureTime.AsTime().Before(time.Now().AddDate(0, 0, 7)) {
			return nil, status.Errorf(codes.InvalidArgument, "Flight must be scheduled at least a week before departure time")
		}

		// In a 24-hour period, only 1 flight is scheduled for a plane
		if fls, err := s.GetFlightsList(ctx, &pb.FlightQuery{
			PlaneNumber:                req.PlaneNumber,
			Status:                     []string{"scheduled"},
			ScheduledDepartureTimeFrom: timestamppb.New(req.ScheduledDepartureTime.AsTime().Add(-12 * time.Hour)),
			ScheduledDepartureTimeTo:   timestamppb.New(req.ScheduledDepartureTime.AsTime().Add(12 * time.Hour)),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		} else if len(fls.Flights) > 0 {
			return nil, status.Errorf(codes.InvalidArgument, "A flight is scheduled for '%v' in '%v'", req.PlaneNumber, fls.Flights[0].ScheduledDepartureTime)
		}

		req.Status = "scheduled"
		req.AvailableSeats = pbPlane.TotalSeats - 10
		req.RealArrivalTime = nil
		req.RealDepartureTime = nil

		err := req.ToModels().Insert(ctx, boil.GetContextDB(), boil.Infer())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to insert flight: %v", err)
		}
	} else {
		// current, err := models.FindFlight(ctx, s.db, req.Id)
		// if err != nil {
		// 	if errors.Is(err, sql.ErrNoRows) {
		// 		return nil, status.Errorf(codes.NotFound, "Flight with ID %s not found", req.Id)
		// 	}
		// 	return nil, status.Errorf(codes.Internal, "Failed to get flight from database: %v", err)
		// }
		// _, err := req.ToModels().Update(ctx, boil.GetContextDB(), boil.Infer())
		return nil, status.Errorf(codes.Unimplemented, "flight are not updatable at this time")
	}

	// return the generated flight id
	return &pb.FlightId{Id: req.Id}, nil
}

// GetFlightsList returns a list of flights based on the input query
func (s *FlightService) GetFlightsList(ctx context.Context, req *pb.FlightQuery) (*pb.FlightList, error) {
	queries := []qm.QueryMod{}
	if len(req.Id) > 0 {
		queries = append(queries, qm.Where("flight_id = ?", req.Id))
	}
	if len(req.PlaneNumber) > 0 {
		queries = append(queries, qm.Where("plane_number = ?", req.PlaneNumber))
	}
	if len(req.DeparturePoint) > 0 {
		queries = append(queries, qm.Where("departure_point = ?", req.DeparturePoint))
	}
	if len(req.DestinationPoint) > 0 {
		queries = append(queries, qm.Where("destination_point = ?", req.DestinationPoint))
	}
	if req.ScheduledDepartureTimeFrom != nil && !req.ScheduledDepartureTimeFrom.AsTime().IsZero() {
		queries = append(queries, qm.Where("scheduled_departure_time >= ?", req.ScheduledDepartureTimeFrom.AsTime()))
	}
	if req.ScheduledDepartureTimeTo != nil && !req.ScheduledDepartureTimeTo.AsTime().IsZero() {
		queries = append(queries, qm.Where("scheduled_departure_time <= ?", req.ScheduledDepartureTimeTo.AsTime()))
	}
	if len(req.Status) > 0 {
		args := []interface{}{}
		for _, status := range req.Status {
			args = append(args, status)
		}
		queries = append(queries, qm.AndIn("status IN ?", args...))
	}
	if req.AvailableSeatsFrom > 0 {
		queries = append(queries, qm.Where("available_seats >= ?", req.AvailableSeatsFrom))
	}
	if req.AvailableSeatsTo > 0 {
		queries = append(queries, qm.Where("available_seats <= ?", req.AvailableSeatsTo))
	}

	flights, err := models.Flights(queries...).All(ctx, s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get flights list, %v", err)
	}

	flightList := &pb.FlightList{}
	for _, f := range flights {
		flight := &pb.Flight{}
		flight.FromModels(f)
		flightList.Flights = append(flightList.Flights, flight)
	}

	fmt.Println("Found", len(flights), "flights")
	return flightList, nil
}

func (s *FlightService) GetFlightById(ctx context.Context, req *pb.FlightId) (*pb.Flight, error) {
	// Get flight from database using sqlboiler
	flightModel, err := models.FindFlight(ctx, s.db, req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Flight with ID %s not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get flight from database: %v", err)
	}

	// Convert sqlboiler model to protobuf message
	flight := &pb.Flight{}
	flight.FromModels(flightModel)

	return flight, nil
}

func (s *FlightService) BookFlight(ctx context.Context, req *pb.BookFlightRequest) (*emptypb.Empty, error) {
	// Retrieve the flight by ID
	flight, err := models.FindFlight(ctx, s.db, req.FlightId)
	if err == sql.ErrNoRows {
		return &emptypb.Empty{}, status.Errorf(codes.NotFound, "flight not found")
	} else if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to retrieve flight: %v", err)
	}

	// Check if the flight is scheduled
	if flight.Status != "scheduled" {
		return &emptypb.Empty{}, status.Errorf(codes.FailedPrecondition, "flight is not scheduled")
	}

	// Check if there are available seats
	if flight.AvailableSeats == 0 {
		return &emptypb.Empty{}, status.Errorf(codes.FailedPrecondition, "flight is fully booked")
	}

	// Check if the departure time is at least 45 minutes from now
	if time.Until(flight.ScheduledDepartureTime) <= 45*time.Minute {
		return &emptypb.Empty{}, status.Errorf(codes.FailedPrecondition, "it is too late to book this flight")
	}

	// Decrease available seats by 1 and save the updated flight
	flight.AvailableSeats -= int(req.NumberOfSeats)
	if _, err := flight.Update(ctx, s.db, boil.Infer()); err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to update flight: %v", err)
	}

	return &emptypb.Empty{}, nil
}

// ChangeFlightStatus updates the status of a flight by its ID
func (s *FlightService) ChangeFlightStatus(ctx context.Context, req *pb.FlightStatusRequest) (*emptypb.Empty, error) {
	// Get the flight by ID
	flight, err := models.Flights(models.FlightWhere.FlightID.EQ(req.FlightId)).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return &emptypb.Empty{}, status.Errorf(codes.NotFound, "flight with ID %s not found", req.FlightId)
		}
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to get flight with ID %s: %v", req.FlightId, err)
	}

	// Update the status of the flight
	flight.Status = req.Status

	// Save the updated flight to the database
	_, err = flight.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "failed to update flight with ID %s: %v", req.FlightId, err)
	}

	// Return a success response
	return &emptypb.Empty{}, nil
}
