-- Create enum types for preferences and units
CREATE TABLE IF NOT EXISTS public.Users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email character varying(255) NOT NULL,
  password character varying(255) NOT NULL,
  preference character varying(100) DEFAULT NULL,
  weight_unit character varying(100) DEFAULT NULL,
  height_unit character varying(100) DEFAULT NULL,
  weight NUMERIC(6, 2),
  height NUMERIC(6, 2),
  name character varying(60),
  image_uri TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the activity types table
CREATE TABLE IF NOT EXISTS public.ActivityTypes (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  calories_per_minute NUMERIC(10, 2) NOT NULL
);

-- Create the activities table without the generated column
CREATE TABLE IF NOT EXISTS public.Activities (
  activityId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  activityType INT NOT NULL REFERENCES ActivityTypes(id),
  doneAt TIMESTAMP NOT NULL,
  durationInMinutes INT NOT NULL CHECK (durationInMinutes >= 1),
  caloriesBurned NUMERIC(10, 2),
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create function to calculate calories
CREATE OR REPLACE FUNCTION calculate_calories_burned()
RETURNS TRIGGER AS $$
BEGIN
    NEW.caloriesBurned := NEW.durationInMinutes * (
        SELECT calories_per_minute 
        FROM ActivityTypes 
        WHERE id = NEW.activityType
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically calculate calories before insert or update
CREATE TRIGGER set_calories_burned
    BEFORE INSERT OR UPDATE OF activityType, durationInMinutes
    ON Activities
    FOR EACH ROW
    EXECUTE FUNCTION calculate_calories_burned();

-- Insert activity types
INSERT INTO
  ActivityTypes (name, calories_per_minute)
VALUES
  ('Walking', 4),
  ('Yoga', 4),
  ('Stretching', 4),
  ('Cycling', 8),
  ('Swimming', 8),
  ('Dancing', 8),
  ('Hiking', 10),
  ('Running', 10),
  ('HIIT', 10),
  ('JumpRope', 10);