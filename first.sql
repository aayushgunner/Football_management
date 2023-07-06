DROP TYPE  IF EXISTS player_position;
DROP TABLE  IF EXISTS players ;
DROP TABLE  IF EXISTS coach ;
DROP TABLE  IF EXISTS team ;


CREATE TYPE player_position AS ENUM (
    'GoalKeeper',
    'Left Back',
    'Right Back',
    'Centre Back'
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
  city VARCHAR(50),
  coach_id INT,
  FOREIGN KEY (coach_id) REFERENCES coach(coach_id)
);

CREATE TABLE players (
   player_id SERIAL PRIMARY KEY,
   first_name VARCHAR,
   last_name VARCHAR,
   position player_position,
   team_id INT,
   FOREIGN KEY (team_id) REFERENCES team(team_id)

);






INSERT into coach (first_name, last_name , age ) VALUES ('Mikel', 'Arteta',41);
INSERT into coach (first_name, last_name , age ) VALUES ('Pep', 'Guardiola',51);
INSERT into coach (first_name, last_name , age ) VALUES ('Jurgen', 'Klopp',54);



INSERT INTO team (team_name , city , coach_id) VALUES('Arsenal', 'London', 1 );
INSERT INTO team (team_name , city , coach_id) VALUES('Manchester City', 'Manchester', 2 );
INSERT INTO team (team_name , city , coach_id) VALUES('Liverpool', 'Liverpool', 3 );


INSERT into players(first_name,last_name,position,team_id) VALUES ('Bukayo', 'Saka'   ,'Right Wing', 1 );
INSERT into players(first_name,last_name,position,team_id) VALUES ('Jack', 'Grealish', 'Left Wing', 2);
INSERT into players(first_name,last_name,position,team_id) VALUES ('Darwin', 'Nunez' , 'Striker',3 );
INSERT into players(first_name,last_name,position,team_id) VALUES ('Gabriel', 'Martinelli','Left Wing',1 );
INSERT into players(first_name,last_name,position,team_id) VALUES ('Gabriel', 'Jesus','Striker', 1 );
INSERT into players(first_name,last_name,position,team_id) VALUES ('Granit', 'Xhaka', 'Centre Midfield',1 );
INSERT into players(first_name,last_name,position,team_id) VALUES ('Ben', 'White' , 'Right Back',1 );

SELECT first_name , last_name FROM players;


SELECT first_name , last_name FROM players where team_id IN (SELECT team_id FROM team WHERE coach_id IN (SELECT coach_id FROM coach WHERE first_name ='Mikel')) ORDER BY first_name ASC;
SELECT * FROM coach;
