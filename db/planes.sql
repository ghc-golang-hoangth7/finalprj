CREATE TABLE planes (
  plane_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  plane_number TEXT UNIQUE NOT NULL,
  total_seats INTEGER NOT NULL,
  status TEXT NOT NULL CHECK (status IN ('flying', 'cleaning', 'repairing', 'ready', 'deleted'))
);
