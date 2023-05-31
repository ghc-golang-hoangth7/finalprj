CREATE TABLE flights (
  flight_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  plane_number VARCHAR(255) NOT NULL REFERENCES planes(plane_number),
  departure_point VARCHAR(255) NOT NULL,
  destination_point VARCHAR(255) NOT NULL,
  scheduled_departure_time TIMESTAMP NOT NULL,
  estimated_arrival_time TIMESTAMP NOT NULL,
  available_seats INTEGER NOT NULL,
  real_departure_time TIMESTAMP,
  real_arrival_time TIMESTAMP,
  status VARCHAR(255) NOT NULL CHECK (status IN ('scheduled', 'canceled', 'flying', 'finished'))
);