DROP table If EXISTS teams;


CREATE TABLE teams (
   team_id SERIAL PRIMARY KEY,
   team_name VARCHAR,
   Abbreviation VARCHAR,
   hex_code VARCHAR,
   team_logo VARCHAR,
   Stadium VARCHAR  
)
