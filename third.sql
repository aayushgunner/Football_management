DROP TABLE if EXISTS stats;
CREATE TABLE stats (
stats_id SERIAL PRIMARY KEY,
games int,
Mins_played int,
goal int,
assists int,
shots int,
key_passes int,
yellow_cards int,
red_cards int,
position VARCHAR,
player_id int,
FOREIGN KEY (player_id) REFERENCES player (player_id)   

)

