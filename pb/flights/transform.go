package flights

import (
	"github.com/ghc-golang-hoangth7/finalprj/models"

	"github.com/volatiletech/null/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (pb *Flight) FromModels(flight *models.Flight) {
	pb.Id = flight.FlightID
	pb.PlaneNumber = flight.PlaneNumber
	pb.DeparturePoint = flight.DeparturePoint
	pb.DestinationPoint = flight.DestinationPoint
	pb.ScheduledDepartureTime = timestamppb.New(flight.ScheduledDepartureTime)
	pb.EstimatedArrivalTime = timestamppb.New(flight.EstimatedArrivalTime)
	pb.RealDepartureTime = timestamppb.New(flight.RealDepartureTime.Time)
	pb.RealArrivalTime = timestamppb.New(flight.RealArrivalTime.Time)
	pb.AvailableSeats = int32(flight.AvailableSeats)
	pb.Status = flight.Status
}

func (pb *Flight) ToModels() *models.Flight {
	return &models.Flight{
		FlightID:               pb.Id,
		PlaneNumber:            pb.PlaneNumber,
		DeparturePoint:         pb.DeparturePoint,
		DestinationPoint:       pb.DestinationPoint,
		ScheduledDepartureTime: pb.ScheduledDepartureTime.AsTime(),
		EstimatedArrivalTime:   pb.EstimatedArrivalTime.AsTime(),
		RealDepartureTime:      null.Time{Time: pb.RealDepartureTime.AsTime(), Valid: true},
		RealArrivalTime:        null.Time{Time: pb.RealArrivalTime.AsTime(), Valid: true},
		AvailableSeats:         int(pb.AvailableSeats),
		Status:                 pb.Status,
	}
}
