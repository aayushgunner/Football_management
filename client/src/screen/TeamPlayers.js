import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom/cjs/react-router-dom.min";
import Saka from "../assets/images/saka.png";
const TeamPlayerCard = ({ teamPlayer, index }) => {
  const {id} = useParams();
    const lightenColor = (color, percent) => {
        const num = parseInt(color.slice(1), 16);
        const amt = Math.round(2.55 * percent);
        const R = (num >> 16) + amt;
        const G = ((num >> 8) & 0x00ff) + amt;
        const B = (num & 0x0000ff) + amt;
        return `#${(0x1000000 + (R << 16) + (G << 8) + B).toString(16).slice(1)}`;
      };
      const style = {
        backgroundColor: teamPlayer.hex_code,
        filter: 'brightness(200%)', 
        opacity: 0.1// Increase brightness to lighten the color
      };
      let imageSrc = null;

  try {
    // Try to require the image
    imageSrc = require(`../assets/newimages/${teamPlayer.player_id}.png`);
  } catch (error) {
    // Image not found, imageSrc remains null
  }

  return (
    <div className="TeamPlayerCard">
      <div className="TeamPlayerCard_Top">
        <img
          src={require("../assets/images/hello.svg")}
          className="TeamPlayerCard_Top_SVG"
          style={style}
        />
        <div className="TeamPlayerCard_Top_Left">
            <div className="TeamPlayerCard_Top_Left_Box">
                <p>Appearances</p>
                <h3>{teamPlayer.games}</h3>
            </div>
            <div className="TeamPlayerCard_Top_Left_Box">
                <p>Goals</p>
                <h3>{teamPlayer.goal}</h3>
            </div>
            <div className="TeamPlayerCard_Top_Left_Box">
                <p>Assists</p>
                <h3>{teamPlayer.assists}</h3>
            </div>
            <div className="TeamPlayerCard_Top_Left_Box">
                <p>Shots</p>
                <h3>{teamPlayer.shots}</h3>
            </div>
        </div>
        <div
          className="TeamPlayerCard_Top_Right"
          style={{ background: teamPlayer.hex_code }}
        >
          {/* <img src={require(`../assets/newimages/${teamPlayer.player_id}.png`)} alt="Saka" /> */}
          {imageSrc && (
        <img src={imageSrc} alt={teamPlayer.player_name} />
      )}
        </div>
      </div>
      <div className="TeamPlayerCard_Bottom">
        <p>{teamPlayer.player_name.split(" ")[0]}</p>
        <h3>{teamPlayer.player_name.split(" ")[1]}</h3>        
      </div>
    </div>
  );
};
const TeamPlayers = () => {
  const { id } = useParams();
  const [teamPlayers, setTeamPlayers] = useState([]);
  const fetchTeamPlayers = (player_name) => {
    axios
      .get(`http://localhost:8080/team/${id}`)
      .then((response) => {
        setTeamPlayers(response.data);
      })
      .catch((error) => {
        console.error(error);
      });
  };
  useEffect(() => {
    fetchTeamPlayers();
  }, []);

  return (
    <div className="TeamPlayer">
      {teamPlayers.map((teamPlayer, index) => {
        return <TeamPlayerCard teamPlayer={teamPlayer} index={index} />;
      })}
    </div>
  );
};

export default TeamPlayers;
