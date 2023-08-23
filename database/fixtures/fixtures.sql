-- DROP table if exists fixtures ;

-- CREATE TABLE fixtures (
-- fixture_id SERIAL PRIMARY KEY,
-- fixture_date VARCHAR,
-- team_home_id int,
-- team_home_name VARCHAR,
-- team_home_winner VARCHAR,
-- team_away_id int,
-- team_away_name VARCHAR,
-- team_away_winner VARCHAR,
-- home_goals int,
-- away_goals int
-- )
UPDATE fixtures SET team_home_name = 'Newcastle United' WHERE team_home_name = 'Newcastle';
UPDATE fixtures SET team_away_name = 'Newcastle United' WHERE team_away_name = 'Newcastle';
UPDATE fixtures SET team_away_name = 'Wolverhampton Wanderers' WHERE team_away_name = 'Wolves';
UPDATE fixtures SET team_home_name = 'Wolverhampton Wanderers' WHERE team_home_name = 'Wolves';

UPDATE fixtures SET team_home_id = 1 WHERE  team_home_name='Arsenal';
UPDATE fixtures SET team_home_id = 2 WHERE  team_home_name='Aston Villa';
UPDATE fixtures SET team_home_id = 3 WHERE  team_home_name='Bournemouth';
UPDATE fixtures SET team_home_id = 4 WHERE  team_home_name='Brentford';
UPDATE fixtures SET team_home_id = 5 WHERE  team_home_name='Brighton';
UPDATE fixtures SET team_home_id = 6 WHERE  team_home_name='Chelsea';
UPDATE fixtures SET team_home_id = 7 WHERE  team_home_name='Crystal Palace';
UPDATE fixtures SET team_home_id = 8 WHERE  team_home_name='Everton';
UPDATE fixtures SET team_home_id = 9 WHERE  team_home_name='Fulham';
UPDATE fixtures SET team_home_id = 10 WHERE team_home_name='Leeds';
UPDATE fixtures SET team_home_id = 11 WHERE team_home_name='Leicester';
UPDATE fixtures SET team_home_id = 12 WHERE team_home_name='Liverpool';
UPDATE fixtures SET team_home_id = 13 WHERE team_home_name='Manchester City';
UPDATE fixtures SET team_home_id = 14 WHERE team_home_name='Manchester United';
UPDATE fixtures SET team_home_id = 15 WHERE team_home_name='Newcastle United';
UPDATE fixtures SET team_home_id = 16 WHERE team_home_name='Nottingham Forest';
UPDATE fixtures SET team_home_id = 17 WHERE team_home_name='Southampton';
UPDATE fixtures SET team_home_id = 18 WHERE team_home_name='Tottenham';
UPDATE fixtures SET team_home_id = 19 WHERE team_home_name='West Ham';
UPDATE fixtures SET team_home_id = 20 WHERE team_home_name='Wolverhampton Wanderers';




UPDATE fixtures SET team_away_id=1 WHERE  team_away_name='Arsenal';
UPDATE fixtures SET team_away_id=2 WHERE  team_away_name='Aston Villa';
UPDATE fixtures SET team_away_id=3 WHERE  team_away_name='Bournemouth';
UPDATE fixtures SET team_away_id=4 WHERE  team_away_name='Brentford';
UPDATE fixtures SET team_away_id=5 WHERE  team_away_name='Brighton';
UPDATE fixtures SET team_away_id=6 WHERE  team_away_name='Chelsea';
UPDATE fixtures SET team_away_id=7 WHERE  team_away_name='Crystal Palace';
UPDATE fixtures SET team_away_id=8 WHERE  team_away_name='Everton';
UPDATE fixtures SET team_away_id=9 WHERE  team_away_name='Fulham';
UPDATE fixtures SET team_away_id=10 WHERE team_away_name='Leeds';
UPDATE fixtures SET team_away_id=11 WHERE team_away_name='Leicester';
UPDATE fixtures SET team_away_id=12 WHERE team_away_name='Liverpool';
UPDATE fixtures SET team_away_id=13 WHERE team_away_name='Manchester City';
UPDATE fixtures SET team_away_id=14 WHERE team_away_name='Manchester United';
UPDATE fixtures SET team_away_id=15 WHERE team_away_name='Newcastle United';
UPDATE fixtures SET team_away_id=16 WHERE team_away_name='Nottingham Forest';
UPDATE fixtures SET team_away_id=17 WHERE team_away_name='Southampton';
UPDATE fixtures SET team_away_id=18 WHERE team_away_name='Tottenham';
UPDATE fixtures SET team_away_id=19 WHERE team_away_name='West Ham';
UPDATE fixtures SET team_away_id=20 WHERE team_away_name='Wolverhampton Wanderers';


