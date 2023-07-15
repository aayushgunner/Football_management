DROP table If EXISTS team;


CREATE TABLE team (
   team_id SERIAL PRIMARY KEY,
   team_name VARCHAR,
   Abbreviation VARCHAR,
   hex_code VARCHAR,
   team_logo VARCHAR,
   Stadium VARCHAR  
)

