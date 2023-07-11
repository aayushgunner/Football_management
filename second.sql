DROP TABLE IF EXISTS stats;

CREATE TABLE stats (
  stats_id SERIAL PRIMARY KEY,
  Player VARCHAR,
  Nation VARCHAR,
  Pos VARCHAR,
  Squad VARCHAR,
  Comp VARCHAR
  
);