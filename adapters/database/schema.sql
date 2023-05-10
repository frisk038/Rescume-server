CREATE TABLE user_locations (
  user_id STRING,
  location GEOGRAPHY(POINT),
  timestamp TIMESTAMPTZ DEFAULT NOW(),
  PRIMARY KEY (user_id, timestamp)
);

