package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ghc-golang-hoangth7/finalprj/models"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
)

type FlightService struct {
	pb.UnimplementedFlightServiceServer
	db *sql.DB
}

func (s *FlightService) CreateFlight(ctx context.Context, req *pb.Flight) (*pb.FlightId, error) {
	// TODO: get plane's info
	// convert proto message to sqlboiler model
	flight := &models.Flight{
		PlaneNumber:          req.PlaneNumber,
		DeparturePoint:       req.DeparturePoint,
		DestinationPoint:     req.DestinationPoint,
		DepartureTime:        req.DepartureTime.AsTime(),
		EstimatedArrivalTime: req.EstimatedArrivalTime.AsTime(),
		Status:               "scheduled",
		AvailableSeats:       int(req.AvailableSeats),
	}

	// insert to database
	err := flight.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert flight: %v", err)
	}

	// return the generated flight id
	return &pb.FlightId{Id: flight.FlightID}, nil
}

// GetFlightsList returns a list of flights based on the input query
func (s *FlightService) GetFlightsList(ctx context.Context, req *pb.Flight) (*pb.FlightList, error) {
	flights, err := models.Flights(
		qm.Where("departure_point = ?", req.DeparturePoint),
		qm.Where("destination_point = ?", req.DestinationPoint),
		qm.Where("departure_time >= ?", req.DepartureTime),
		qm.Where("departure_time <= ?", req.EstimatedArrivalTime),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
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

func (s *FlightService) BookFlight(ctx context.Context, req *pb.BookFlightRequest) error {
	// Retrieve the flight by ID
	flight, err := models.FindFlight(ctx, s.db, req.FlightId)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to retrieve flight: %v", err)
	}

	// Check if the flight is scheduled
	if flight.Status != "scheduled" {
		return status.Errorf(codes.FailedPrecondition, "flight is not scheduled")
	}

	// Check if there are available seats
	if flight.AvailableSeats == 0 {
		return status.Errorf(codes.FailedPrecondition, "flight is fully booked")
	}

	// Check if the departure time is at least 45 minutes from now
	if time.Until(flight.DepartureTime) <= 45*time.Minute {
		return status.Errorf(codes.FailedPrecondition, "it is too late to book this flight")
	}

	// Decrease available seats by 1 and save the updated flight
	flight.AvailableSeats -= int(req.SeatNumber)
	if _, err := flight.Update(ctx, s.db, boil.Infer()); err != nil {
		return status.Errorf(codes.Internal, "failed to update flight: %v", err)
	}

	return nil
}

// ChangeFlightStatus updates the status of a flight by its ID
func (s *FlightService) ChangeFlightStatus(ctx context.Context, req *pb.FlightStatusRequest) (*pb.Flight, error) {
	// Get the flight by ID
	flight, err := models.Flights(models.FlightWhere.FlightID.EQ(req.FlightId)).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "flight with ID %s not found", req.FlightId)
		}
		return nil, status.Errorf(codes.Internal, "failed to get flight with ID %s: %v", req.FlightId, err)
	}

	// Update the status of the flight
	flight.Status = req.Status

	// Save the updated flight to the database
	_, err = flight.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update flight with ID %s: %v", req.FlightId, err)
	}

	protoFlight := &pb.Flight{}
	protoFlight.FromModels(flight)

	// Return a success response
	return protoFlight, nil
}
