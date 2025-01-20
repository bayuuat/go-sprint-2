-- Create enum types for preferences and units
CREATE TABLE IF NOT EXISTS public.users (
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
CREATE TABLE IF NOT EXISTS public.activity_types (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  calories_per_minute NUMERIC(10, 2) NOT NULL
);

-- Create the activities table without the generated column
CREATE TABLE IF NOT EXISTS public.Activities (
  activityId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  activity_type INT NOT NULL REFERENCES activity_types(id),
  done_at TIMESTAMP NOT NULL,
  duration_in_minutes INT NOT NULL CHECK (duration_in_minutes >= 1),
  calories_burned NUMERIC(10, 2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id UUID NOT NULL
);

-- Create function to calculate calories
CREATE OR REPLACE FUNCTION calculate_calories_burned()
RETURNS TRIGGER AS $$
BEGIN
    NEW.calories_burned := NEW.duration_in_minutes * (
        SELECT calories_per_minute 
        FROM activity_types
        WHERE id = NEW.activity_type
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically calculate calories before insert or update
CREATE TRIGGER set_calories_burned
    BEFORE INSERT OR UPDATE OF activity_type, duration_in_minutes
    ON activities
    FOR EACH ROW
    EXECUTE FUNCTION calculate_calories_burned();

-- Create function to update updated_at
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';


-- Create trigger in users to update updated_at
CREATE TRIGGER update_modified_time
    BEFORE UPDATE
    ON users
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

-- Create trigger in activities to update updated_at
CREATE TRIGGER update_modified_time
    BEFORE UPDATE
    ON activities
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

-- Insert activity types
INSERT INTO
  activity_types (name, calories_per_minute)
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
