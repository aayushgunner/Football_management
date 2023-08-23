import React, { useEffect, useState } from "react";
import { getTeam } from "../apiRoutes";
import { useHistory } from "react-router-dom/cjs/react-router-dom.min";
const Team = () => {
  const [teams, setTeams] = useState([]);
  //   const team = {
  //     team_id: "1",
  //     team_name: "Arsenal",
  //     team_logo: "https://a.espncdn.com/i/teamlogos/soccer/500/359.png",
  //     hex_code: "#EF0107",
  //     abbreviation: "ARS",
  //     Stadium: "Emirates",
  //   };
  useEffect(() => {
    console.log('teams')
    getTeam().then((res) => {
      console.log('teams')
        console.log(res);
      setTeams(res);
    });
  }, []);
  const TeamCard = ({ team, index }) => {
    const [isHovered, setIsHovered] = useState(false);
const history = useHistory()
    const handleMouseEnter = () => {
      setIsHovered(true);
    };

    const handleMouseLeave = () => {
      setIsHovered(false);
    };
    const stopColor1 = '#000000'; // Color for the first stop
    const stopColor2 = 'red'; // Color for the second stop
    return (
      <div
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
        className="Team_Card"
        key={index}
        style={{ borderBottom: `5px solid ${team.hex_code}` }}
        onClick={()=>history.push(`/team/${team.team_id}`)}
      >
        <img
          src={team.logo}
          alt={team.team_name}
          className="Team_Card_Photo"
        />
        <h3>{team.team_name}</h3>
        <img
          src={require("../assets/images/hello.svg")}
          className="Team_Card_Svg club-card__svg club-card__gradient-svg"
          style={{
            background: isHovered && team.hex_code, 
            '--stop-color-1': stopColor1,
            '--stop-color-2': stopColor2,
          }}
        />
      </div>
    );
  };
  return (
    <div className="Team">
      {teams.map((team, index) => {
        return <TeamCard team={team} index={index} />;
      })}
    </div>
  );
};

export default Team;
