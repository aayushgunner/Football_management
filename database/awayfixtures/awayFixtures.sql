-- DROP table if EXISTS awayFixtures;
-- CREATE table awayFixtures (
--     team_away_id int,
--     team_away_name VARCHAR,
--     Shots_on_goal int,
--     Shots_off_goal int,
--     total_shots int,
--     blocked_shots int,
--     shots_insideBox int,
--     shots_outsideBox int,
--     fouls int,
--     corner_kicks int,
--     offsides int,
--     possession VARCHAR,
--     yellow_cards int,
--     red_cards int,
--     saves int,
--     total_passes int,
--     accurate_passes int,
--     pass_accuracy VARCHAR,
--     fixture_id int
--     -- FOREIGN KEY (team_home_id) REFERENCES teams(team_id) ,
--     -- FOREIGN KEY (fixture_id) REFERENCES fixtures(fixture_id) 
--  )



-- UPDATE awayFixtures SET team_home_name = 'Newcastle United' WHERE team_home_name = 'Newcastle';
-- UPDATE awayFixtures SET team_away_name = 'Newcastle United' WHERE team_away_name = 'Newcastle';
-- UPDATE awayFixtures SET team_away_name = 'Wolverhampton Wanderers' WHERE team_away_name = 'Wolves';
-- UPDATE awayFixtures SET team_home_name = 'Wolverhampton Wanderers' WHERE team_home_name = 'Wolves';

UPDATE awayFixtures SET team_away_id=1 WHERE  team_away_name='Arsenal';
UPDATE awayFixtures SET team_away_id=2 WHERE  team_away_name='Aston Villa';
UPDATE awayFixtures SET team_away_id=3 WHERE  team_away_name='Bournemouth';
UPDATE awayFixtures SET team_away_id=4 WHERE  team_away_name='Brentford';
UPDATE awayFixtures SET team_away_id=5 WHERE  team_away_name='Brighton';
UPDATE awayFixtures SET team_away_id=6 WHERE  team_away_name='Chelsea';
UPDATE awayFixtures SET team_away_id=7 WHERE  team_away_name='Crystal Palace';
UPDATE awayFixtures SET team_away_id=8 WHERE  team_away_name='Everton';
UPDATE awayFixtures SET team_away_id=9 WHERE  team_away_name='Fulham';
UPDATE awayFixtures SET team_away_id=10 WHERE team_away_name='Leeds';
UPDATE awayFixtures SET team_away_id=11 WHERE team_away_name='Leicester';
UPDATE awayFixtures SET team_away_id=12 WHERE team_away_name='Liverpool';
UPDATE awayFixtures SET team_away_id=13 WHERE team_away_name='Manchester City';
UPDATE awayFixtures SET team_away_id=14 WHERE team_away_name='Manchester United';
UPDATE awayFixtures SET team_away_id=15 WHERE team_away_name='Newcastle United';
UPDATE awayFixtures SET team_away_id=16 WHERE team_away_name='Nottingham Forest';
UPDATE awayFixtures SET team_away_id=17 WHERE team_away_name='Southampton';
UPDATE awayFixtures SET team_away_id=18 WHERE team_away_name='Tottenham';
UPDATE awayFixtures SET team_away_id=19 WHERE team_away_name='West Ham';
UPDATE awayFixtures SET team_away_id=20 WHERE team_away_name='Wolverhampton Wanderers';

