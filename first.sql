DROP TYPE  IF EXISTS player_position;
DROP TABLE  IF EXISTS players ;
DROP TABLE  IF EXISTS coach ;
DROP TABLE  IF EXISTS team ;


CREATE TYPE player_position AS ENUM (
    'GoalKeeper',
    'Left Back',
    'Right Back',
    'Centre Back',
    'Left Midfield',
    'Right Midfield',
    'Centre Midfield',
    'Centre Defensive Midfield',
    'Centre Attacking Midfield',
    'Left Wing',
    'Right Wing',
    'Striker'
);



CREATE TABLE coach (
  coach_id SERIAL PRIMARY KEY,
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  age INT
);


CREATE TABLE team (
  team_id SERIAL PRIMARY KEY,
  team_name VARCHAR(50),
  team_logo VARCHAR(300),
  hex_code VARCHAR(10),
  abbreviation VARCHAR(9),
  Stadium VARCHAR(50)
);

CREATE TABLE players (
   player_id SERIAL PRIMARY KEY,
   first_name VARCHAR,
   last_name VARCHAR,
   position player_position,
   team_id INT,
   FOREIGN KEY (team_id) REFERENCES team(team_id)
);
