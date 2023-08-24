import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link, useHistory } from "react-router-dom/cjs/react-router-dom.min";

// fixture_id	1
// fixture_date	"2022-08-05T19:00:00+00:00"
// team_home_id	7
// team_home_name	"Crystal Palace"
// team_away_id	1
// team_away_name	"Arsenal"
// home_goals	0
// away_goals	2

const Fixtures = () => {
  const [fixtures, setFixtures] = useState([]);
  const history = useHistory()
  const getFixtures = () => {
    axios
      .get(`http://localhost:8080/fixtures`)
      .then((response) => {
        setFixtures(response.data);
      })
      .catch((error) => {
        console.error(error);
      });
  };
  useEffect(() => {
    getFixtures();
  }, []);
  const FixtureCard = ({ fixture, index }) => {
    return (
      <Link
       to={`/matchStats/${fixture.fixture_id}`}
        className="Fixtures_Box"
        key={index}
       
      >
        <div className="Fixtures_Box_Left">
          {" "}
          <p>{fixture.team_home_name}</p>
        </div>

        <div className="Fixtures_Box_Score">
          <p>
            {fixture.home_goals} - {fixture.away_goals}
          </p>
        </div>
        <div className="Fixtures_Box_Right">
          <p>{fixture.team_away_name}</p>
        </div>
      </Link>
    );
  };
  return (
    <div className="Fixtures">
      <h1 style={{ marginLeft: 50 }}>Results</h1>
      {fixtures.map((fixture, index) => {
        return <FixtureCard fixture={fixture} index={index}  onClick={() => {
          history.push(`/fixtures/${fixture.fixture_id}`);
        }}/>;
      })}
    </div>
  );
};

export default Fixtures;
