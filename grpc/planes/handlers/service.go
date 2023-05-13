package handlers

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
)

type PlanesService struct {
	pb.UnimplementedPlanesServiceServer
	db *sql.DB
}

func NewPlanesService(db *sql.DB) *PlanesService {
	return &PlanesService{db: db}
}

func (s *PlanesService) ListPlanes(ctx context.Context, req *emptypb.Empty) (*pb.PlaneList, error) {
	rows, err := s.db.Query("SELECT plane_id, plane_number, total_seats, status FROM planes;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planes []*pb.Plane
	for rows.Next() {
		var plane pb.Plane
		if err := rows.Scan(&plane.PlaneId, &plane.PlaneNumber, &plane.TotalSeats, &plane.Status); err != nil {
			return nil, err
		}
		planes = append(planes, &plane)
	}

	return &pb.PlaneList{Planes: planes}, nil
}

func (s *PlanesService) UpdatePlaneStatus(ctx context.Context, req *pb.Plane) (*emptypb.Empty, error) {
	// TODO: check scheduler
	_, err := s.db.Exec("UPDATE planes SET status = $1 WHERE plane_id = $2;", req.Status, req.PlaneId)
	if err != nil {
		return nil, err
	}

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
		return nil, status.Errorf(codes.InvalidArgument, "Missing status")
	}

	// If the request doesn't have a plane_id, generate one
	if req.PlaneId == "" {
		req.PlaneId = uuid.New().String()
	}

	res, err := s.db.Exec("UPDATE planes SET plane_number = $1, total_seats = $2, status = $3 WHERE plane_id = $4 OR plane_number = $5", req.PlaneNumber, req.TotalSeats, req.Status, req.PlaneId, req.PlaneNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to add or update plane: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get rows affected: %v", err)
	}

	// If no rows were affected, the plane wasn't found, so we insert a new one
	if rowsAffected == 0 {
		_, err := s.db.Exec("INSERT INTO planes (plane_id, plane_number, total_seats, status) VALUES ($1, $2, $3, $4)", req.PlaneId, req.PlaneNumber, req.TotalSeats, req.Status)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to add or update plane: %v", err)
		}
	}

	return req, nil
}
