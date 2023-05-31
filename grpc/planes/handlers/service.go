package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	// "github.com/ghc-golang-hoangth7/finalprj/models"
	"github.com/ghc-golang-hoangth7/finalprj/models"
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

type PlanesService struct {
	pb.UnimplementedPlanesServiceServer
	FlightsService pbFlights.FlightServiceClient
	db             *sql.DB
}

func NewPlanesService(db *sql.DB) *PlanesService {
	return &PlanesService{db: db}
}

func (s *PlanesService) UpsertPlane(ctx context.Context, req *pb.Plane) (*pb.PlaneId, error) {
	if req.PlaneNumber == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing plane number")
	}
	if req.TotalSeats <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid total seats")
	}
	if req.Status == "" {
		req.Status = "ready"
	}
	if req.PlaneId == "" {
		req.PlaneId = uuid.New().String()
	}

	// convert proto message to sqlboiler model
	plane := &models.Plane{
		PlaneNumber: req.PlaneNumber,
		PlaneID:     req.PlaneId,
		TotalSeats:  int(req.TotalSeats),
		Status:      req.Status,
	}

	//TODO: update plane
	// insert to database
	err := plane.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert plane: %v", err)
	}

	// return the generated plane id
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
		queries = append(queries, qm.Where("status = ?", req.Status))
	}
	if req.TotalSeatsFrom > 0 {
		queries = append(queries, qm.Where("total_seats >= ?", req.TotalSeatsFrom))
	}
	if req.TotalSeatsTo > 0 {
		queries = append(queries, qm.Where("total_seats <= ?", req.TotalSeatsTo))
	}

	planes, err := models.Planes(queries...).All(ctx, s.db)
	if err != nil {
		return nil, err
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
	planeModel, err := models.FindPlane(ctx, s.db, req.Id)
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
	// TODO: check scheduler
	// Get the plane by ID
	plane, err := models.Planes(models.PlaneWhere.PlaneID.EQ(req.PlaneId)).One(ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "plane with ID %s not found", req.PlaneId)
		}
		return nil, status.Errorf(codes.Internal, "failed to get plane with ID %s: %v", req.PlaneId, err)
	}

	// Update the status of the plane
	plane.Status = req.Status

	// Save the updated plane to the database
	_, err = plane.Update(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update plane with ID %s: %v", req.PlaneId, err)
	}

	protoPlane := &pb.Plane{}
	protoPlane.FromModels(plane)

	// Return a success response
	return &emptypb.Empty{}, nil
}
