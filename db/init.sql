CREATE DATABASE planes_db;

CREATE TABLE flights (
    flight_id uuid DEFAULT gen_random_uuid() NOT NULL,
    plane_number character varying(255) NOT NULL,
    departure_point character varying(255) NOT NULL,
    destination_point character varying(255) NOT NULL,
    scheduled_departure_time timestamp without time zone NOT NULL,
    estimated_arrival_time timestamp without time zone NOT NULL,
    available_seats integer NOT NULL,
    real_departure_time timestamp without time zone,
    real_arrival_time timestamp without time zone,
    status character varying(255) NOT NULL,
    CONSTRAINT flights_status_check CHECK (((status)::text = ANY ((ARRAY['scheduled'::character varying, 'canceled'::character varying, 'flying'::character varying, 'finished'::character varying])::text[])))
);

CREATE TABLE planes (
    plane_id uuid DEFAULT gen_random_uuid() NOT NULL,
    plane_number text NOT NULL,
    total_seats integer NOT NULL,
    status text NOT NULL,
    CONSTRAINT planes_status_check CHECK ((status = ANY (ARRAY['flying'::text, 'cleaning'::text, 'repairing'::text, 'ready'::text, 'deleted'::text])))
);

INSERT INTO flights VALUES ('17c48777-56c2-4a87-9c1a-b347c0b7879a', 'A123456', 'New York', 'London', '2023-07-01 09:00:00', '2023-07-01 12:00:00', 140, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('f6dcbe35-e352-4b49-a17f-227c563de15c', 'A123456', 'London', 'New York', '2023-07-02 12:00:00', '2023-07-02 15:00:00', 140, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('c55d0fb9-2a5b-4b77-8dd4-0ac89c1b3103', 'B234567', 'Paris', 'Rome', '2023-07-03 09:00:00', '2023-07-03 12:00:00', 170, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('2c9aea25-3d50-4dfc-bca6-540bc4b21f08', 'B234567', 'Rome', 'Paris', '2023-07-04 12:00:00', '2023-07-04 15:00:00', 170, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('677f6314-6683-41f3-b144-e471c39d1be1', 'C345678', 'Berlin', 'Amsterdam', '2023-07-05 09:00:00', '2023-07-05 12:00:00', 220, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('f62ebd7a-a9ce-4833-b08d-77b3150a75f6', 'C345678', 'Amsterdam', 'Berlin', '2023-07-06 12:00:00', '2023-07-06 15:00:00', 220, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('45af00a0-a27b-4622-8c9e-cb1dfd3be79c', 'D456789', 'Tokyo', 'Shanghai', '2023-07-07 09:00:00', '2023-07-07 12:00:00', 150, NULL, NULL, 'scheduled');
INSERT INTO flights VALUES ('6afb6ac7-9591-4ac4-bcd1-515cb0c6e4d5', 'D456789', 'Shanghai', 'Tokyo', '2023-07-08 12:00:00', '2023-07-08 15:00:00', 150, NULL, NULL, 'scheduled');

INSERT INTO planes VALUES ('8f72f081-d5cc-4bab-b821-8a8a06c6c5b2', 'A123456', 150, 'ready');
INSERT INTO planes VALUES ('de7a65f4-53f5-42d3-bc6c-e68e348a0397', 'B234567', 180, 'flying');
INSERT INTO planes VALUES ('ce5ca06b-031f-4ff2-928a-4952441e58e0', 'C345678', 230, 'cleaning');
INSERT INTO planes VALUES ('a3eb4a67-f5f6-4b9c-9fb8-ad1d67495d33', 'D456789', 160, 'repairing');
INSERT INTO planes VALUES ('fbb89d24-cea6-4763-b558-a58587b33fbc', 'E567890', 170, 'deleted');
INSERT INTO planes VALUES ('8f62872a-6033-4e10-8ca0-bd97363ba673', 'F678901', 180, 'ready');
INSERT INTO planes VALUES ('b52cd7c9-6a1b-49d3-a6d2-3b9663dcbc86', 'G789012', 250, 'flying');
INSERT INTO planes VALUES ('639e282d-7cfa-4ec5-8e8e-7a42070140f7', 'H890123', 270, 'cleaning');
INSERT INTO planes VALUES ('7ff6b191-37c3-4916-86b5-f95aa06042b6', 'I901234', 300, 'repairing');
INSERT INTO planes VALUES ('7cc788c4-c3c0-4b0b-9890-4936fa9b0187', 'J012345', 360, 'ready');


ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (flight_id);

ALTER TABLE ONLY planes
    ADD CONSTRAINT planes_pkey PRIMARY KEY (plane_id);

ALTER TABLE ONLY planes
    ADD CONSTRAINT planes_plane_number_key UNIQUE (plane_number);

ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_plane_number_fkey FOREIGN KEY (plane_number) REFERENCES planes(plane_number);
