package planes

import (
	"github.com/ghc-golang-hoangth7/finalprj/models"
)

func (pb *Plane) FromModels(plane *models.Plane) {
	pb.PlaneId = plane.PlaneID
	pb.PlaneNumber = plane.PlaneNumber
	pb.TotalSeats = int32(plane.TotalSeats)
	pb.Status = plane.Status
}
func (pb *Plane) ToModels() *models.Plane {
	return &models.Plane{
		PlaneID:     pb.PlaneId,
		PlaneNumber: pb.PlaneNumber,
		TotalSeats:  int(pb.TotalSeats),
		Status:      pb.Status,
	}
}
