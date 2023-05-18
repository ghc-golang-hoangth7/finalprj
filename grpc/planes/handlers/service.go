package handlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	// "github.com/ghc-golang-hoangth7/finalprj/models"
	"github.com/ghc-golang-hoangth7/finalprj/models"
	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

type PlanesService struct {
	pb.UnimplementedPlanesServiceServer
	db *sql.DB
}

func NewPlanesService(db *sql.DB) *PlanesService {
	return &PlanesService{db: db}
}

func (s *PlanesService) ListPlanes(ctx context.Context, req *pb.Plane) (*pb.PlaneList, error) {
	planes, err := models.Planes(
		qm.Where("plane_id = ?", req.PlaneId),
		qm.Where("plane_number = ?", req.PlaneNumber),
		qm.Where("status = ?", req.Status),
		qm.Where("total_seats <= ?", req.TotalSeats),
	).All(ctx, s.db)
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

func (s *PlanesService) UpdatePlaneStatus(ctx context.Context, req *pb.Plane) (*emptypb.Empty, error) {
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

func (s *PlanesService) AddOrUpdatePlane(ctx context.Context, req *pb.Plane) (*pb.Plane, error) {
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

	// insert to database
	err := plane.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert plane: %v", err)
	}

	// return the generated plane id
	return req, nil
}
