package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ghc-golang-hoangth7/finalprj/models"
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

type PlanesService struct {
	pb.UnimplementedPlanesServiceServer
	flightsSrv pbFlights.FlightServiceClient
	db         *sql.DB
}

func NewPlanesService(db *sql.DB, flightsSrv pbFlights.FlightServiceClient) *PlanesService {
	boil.SetDB(db)
	boil.DebugMode = true
	boil.DebugWriter = os.Stdout
	return &PlanesService{db: db, flightsSrv: flightsSrv}
}

func (s *PlanesService) UpsertPlane(ctx context.Context, req *pb.Plane) (*pb.PlaneId, error) {
	if len(req.PlaneId) == 0 {
		req.PlaneId = uuid.New().String()

		if req.Status == "" {
			req.Status = "ready"
		}

		_, err := models.Planes(qm.Where("plane_number = ?", req.PlaneNumber)).One(ctx, boil.GetContextDB())
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.Internal, "Failed to get plane from database: %v", err)
			}
		} else {
			return nil, status.Errorf(codes.AlreadyExists, "An plane with PlaneNumber [%v] existed", req.PlaneNumber)
		}

		// convert proto message to sqlboiler model
		plane := &models.Plane{
			PlaneID:     req.PlaneId,
			PlaneNumber: req.PlaneNumber,
			TotalSeats:  int(req.TotalSeats),
			Status:      req.Status,
		}

		// insert to database
		err = plane.Insert(ctx, boil.GetContextDB(), boil.Infer())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to insert plane: %v", err)
		}
	} else {
		planeModel, err := models.FindPlane(ctx, boil.GetContextDB(), req.PlaneId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "Plane with ID %s not found", req.PlaneId)
			}
			return nil, status.Errorf(codes.Internal, "Failed to get plane from database: %v", err)
		}

		if planeModel.Status != req.Status {
			return nil, status.Errorf(codes.InvalidArgument, "Please use ChangePlaneStatus function instead")
		}
		if planeModel.PlaneNumber != req.PlaneNumber {
			return nil, status.Errorf(codes.InvalidArgument, "PlaneNumber cannot be update")
		}

		planeModel.TotalSeats = int(req.TotalSeats)
		// update to database
		_, err = planeModel.Update(ctx, boil.GetContextDB(), boil.Infer())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to update plane: %v", err)
		}
	}

	return &pb.PlaneId{
		Id: req.PlaneId,
	}, nil
}

func (s *PlanesService) GetPlanesList(ctx context.Context, req *pb.PlaneQuery) (*pb.PlaneList, error) {
	queries := []qm.QueryMod{}
	if len(req.PlaneId) > 0 {
		queries = append(queries, qm.Where("plane_id = ?", req.PlaneId))
	}
	if len(req.PlaneNumber) > 0 {
		queries = append(queries, qm.Where("plane_number = ?", req.PlaneNumber))
	}

	if len(req.Status) > 0 {
		args := []interface{}{}
		for _, status := range req.Status {
			args = append(args, status)
		}
		queries = append(queries, qm.AndIn("status IN ?", args...))
	}

	if req.TotalSeatsFrom > 0 {
		queries = append(queries, qm.Where("total_seats >= ?", req.TotalSeatsFrom))
	}

	if req.TotalSeatsTo > 0 {
		queries = append(queries, qm.Where("total_seats <= ?", req.TotalSeatsTo))
	}

	planes, err := models.Planes(queries...).All(ctx, boil.GetContextDB())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to get planes list, %v", err)
	}

	planeList := &pb.PlaneList{}
	for _, p := range planes {
		plane := &pb.Plane{}
		plane.FromModels(p)
		planeList.Planes = append(planeList.Planes, plane)
	}

	fmt.Println("Found", len(planes), "planes")
	return planeList, nil
}

func (s *PlanesService) GetPlaneById(ctx context.Context, req *pb.PlaneId) (*pb.Plane, error) {
	// Get plane from database using sqlboiler
	planeModel, err := models.FindPlane(ctx, boil.GetContextDB(), req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Plane with ID %s not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get plane from database: %v", err)
	}

	// Convert sqlboiler model to protobuf message
	plane := &pb.Plane{}
	plane.FromModels(planeModel)

	return plane, nil
}

func (s *PlanesService) GetPlaneByNumber(ctx context.Context, req *pb.PlaneNumber) (*pb.Plane, error) {
	// Get plane from database using sqlboiler
	planeModel, err := models.Planes(qm.Where("plane_number = ?", req.PlaneNumber)).One(ctx, boil.GetContextDB())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "Plane with Number %s not found", req.PlaneNumber)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get plane from database: %v", err)
	}

	// Convert sqlboiler model to protobuf message
	plane := &pb.Plane{}
	plane.FromModels(planeModel)

	return plane, nil
}

func (s *PlanesService) ChangePlaneStatus(ctx context.Context, req *pb.PlaneStatusRequest) (*emptypb.Empty, error) {
	// Get the plane by ID
	plane, err := models.Planes(models.PlaneWhere.PlaneID.EQ(req.PlaneId)).One(ctx, boil.GetContextDB())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "plane with ID %s not found", req.PlaneId)
		}
		return nil, status.Errorf(codes.Internal, "failed to get plane with ID %s: %v", req.PlaneId, err)
	}

	if req.Status == "deleted" {
		list, err := s.flightsSrv.GetFlightsList(ctx, &pbFlights.FlightQuery{
			PlaneNumber:                plane.PlaneNumber,
			ScheduledDepartureTimeFrom: timestamppb.Now(),
			Status:                     []string{"scheduled"},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get list scheduled flight, %v", err)
		}
		if len(list.Flights) > 0 {
			return nil, status.Errorf(codes.InvalidArgument, "This plane has scheduled %v flight(s)", len(list.Flights))
		}
	}

	// Update the status of the plane
	plane.Status = req.Status

	// Save the updated plane to the database
	_, err = plane.Update(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update plane with ID %s: %v", req.PlaneId, err)
	}

	protoPlane := &pb.Plane{}
	protoPlane.FromModels(plane)

	// Return a success response
	return &emptypb.Empty{}, nil
}
