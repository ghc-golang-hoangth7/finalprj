package flights

import (
	"github.com/ghc-golang-hoangth7/finalprj/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (pb *Flight) FromModels(flight *models.Flight) {
	pb.Id = flight.FlightID
	pb.PlaneNumber = flight.PlaneNumber
	pb.DeparturePoint = flight.DeparturePoint
	pb.DestinationPoint = flight.DestinationPoint
	pb.DepartureTime = timestamppb.New(flight.DepartureTime)
	pb.EstimatedArrivalTime = timestamppb.New(flight.EstimatedArrivalTime)
	pb.RealDepartureTime = timestamppb.New(flight.RealDepartureTime.Time)
	pb.RealArrivalTime = timestamppb.New(flight.RealArrivalTime.Time)
	pb.AvailableSeats = int32(flight.AvailableSeats)
	pb.Status = flight.Status
}
func (pb *Flight) ToModels() *models.Flight {
	return &models.Flight{}
}
