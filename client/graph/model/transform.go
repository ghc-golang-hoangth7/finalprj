package model

import (
	pbFlights "github.com/ghc-golang-hoangth7/finalprj/pb/flights"
	pbPlanes "github.com/ghc-golang-hoangth7/finalprj/pb/planes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func (flight *Flight) FromProto(pb *pbFlights.Flight) {
	flight.ID = pb.GetId()
	flight.PlaneNumber = pb.GetPlaneNumber()
	flight.DeparturePoint = pb.GetDeparturePoint()
	flight.DestinationPoint = pb.GetDestinationPoint()
	if pb.GetScheduledDepartureTime() != nil {
		flight.ScheduledDepartureTime = pb.GetScheduledDepartureTime().AsTime()
	}
	if pb.GetEstimatedArrivalTime() != nil {
		estimatedArrivalTime := pb.GetEstimatedArrivalTime().AsTime()
		flight.EstimatedArrivalTime = &estimatedArrivalTime
	}
	if pb.GetRealDepartureTime() != nil {
		realDepartureTime := pb.GetRealDepartureTime().AsTime()
		flight.RealDepartureTime = &realDepartureTime
	}
	if pb.GetRealArrivalTime() != nil {
		realDepartureTime := pb.GetRealArrivalTime().AsTime()
		flight.RealDepartureTime = &realDepartureTime
	}

	flight.Status = pb.GetStatus()
	flight.AvailableSeats = int(pb.GetAvailableSeats())

}

func (flight *FlightQuery) ToProto() *pbFlights.FlightQuery {
	pb := &pbFlights.FlightQuery{}
	if flight == nil {
		return pb
	}
	if flight.ID != nil {
		pb.Id = *flight.ID
	}
	if flight.PlaneNumber != nil {
		pb.PlaneNumber = *flight.PlaneNumber
	}
	if flight.DeparturePoint != nil {
		pb.DeparturePoint = *flight.DeparturePoint
	}
	if flight.DestinationPoint != nil {
		pb.DestinationPoint = *flight.DestinationPoint
	}
	if flight.ScheduledDepartureTimeFrom != nil {
		pb.ScheduledDepartureTimeFrom = &timestamp.Timestamp{Seconds: flight.ScheduledDepartureTimeFrom.Unix()}
	}
	if flight.ScheduledDepartureTimeTo != nil {
		pb.ScheduledDepartureTimeTo = &timestamp.Timestamp{Seconds: flight.ScheduledDepartureTimeTo.Unix()}
	}
	if flight.Status != nil {
		pb.Status = []string{}
		for _, v := range flight.Status {
			pb.Status = append(pb.Status, *v)
		}
	}
	if flight.AvailableSeatsFrom != nil {
		pb.AvailableSeatsFrom = int32(*flight.AvailableSeatsFrom)
	}
	if flight.AvailableSeatsTo != nil {
		pb.AvailableSeatsTo = int32(*flight.AvailableSeatsTo)
	}
	return pb
}

func (plane *PlaneQuery) ToProto() *pbPlanes.PlaneQuery {
	pb := &pbPlanes.PlaneQuery{}
	if plane == nil {
		return pb
	}
	if plane.PlaneID != nil {
		pb.PlaneId = *plane.PlaneID
	}
	if plane.PlaneNumber != nil {
		pb.PlaneNumber = *plane.PlaneNumber
	}
	if plane.TotalSeatsFrom != nil {
		pb.TotalSeatsFrom = int32(*plane.TotalSeatsFrom)
	}
	if plane.TotalSeatsTo != nil {
		pb.TotalSeatsTo = int32(*plane.TotalSeatsTo)
	}
	if plane.Status != nil {
		pb.Status = []string{}
		for _, v := range plane.Status {
			pb.Status = append(pb.Status, *v)
		}
	}
	return pb
}
