-- DROP table if EXISTS player;
-- CREATE table player (
-- player_id SERIAL PRIMARY KEY,
-- player_link_id INT,
-- player_name VARCHAR,
-- team_name VARCHAR,
-- team_id int,
-- FOREIGN KEY (team_id) REFERENCES teams (team_id)  
-- )


UPDATE player SET team_id = 1 WHERE team_name='Arsenal';
UPDATE player SET team_id = 2 WHERE team_name='Aston Villa';
UPDATE player SET team_id = 3 WHERE team_name='Bournemouth';
UPDATE player SET team_id = 4 WHERE team_name='Brentford';
UPDATE player SET team_id = 5 WHERE team_name='Brighton';
UPDATE player SET team_id = 6 WHERE team_name='Chelsea';
UPDATE player SET team_id = 7 WHERE team_name='Crystal Palace';
UPDATE player SET team_id = 8 WHERE team_name='Everton';
UPDATE player SET team_id = 9 WHERE team_name='Fulham';
UPDATE player SET team_id = 10 WHERE team_name='Leeds';
UPDATE player SET team_id = 11 WHERE team_name='Leicester';
UPDATE player SET team_id = 12 WHERE team_name='Liverpool';
UPDATE player SET team_id = 13 WHERE team_name='Manchester City';
UPDATE player SET team_id = 14 WHERE team_name='Manchester United';
UPDATE player SET team_id = 15 WHERE team_name='Newcastle United';
UPDATE player SET team_id = 16 WHERE team_name='Nottingham Forest';
UPDATE player SET team_id = 17 WHERE team_name='Southampton';
UPDATE player SET team_id = 18 WHERE team_name='Tottenham';
UPDATE player SET team_id = 19 WHERE team_name='West Ham';
UPDATE player SET team_id = 20 WHERE team_name='Wolverhampton Wanderers';


-- UPDATE player
-- SET team_name = TRIM(SPLIT_PART(team_name, ',', 1))
-- WHERE team_name LIKE '%,%';
