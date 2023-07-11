
import React, { useState , useEffect} from 'react';
import axios from 'axios';
import '../assets/styles/stats.css'

const Stats =() => {
  const [data, setData] = useState([]);
  const [showTable, setShowTable] = useState(false);
  const [coach, setCoach] = useState('');
  const [routes, setRoutes] = useState('');
  const [team , setTeam] = useState('');
  const [teams, setTeams] = useState([])
  const [players, setPlayers] = useState([])
  const [stats, setStats]= useState([])
  const [selectedPlayer, setSelectedPlayer] = useState(null);

  useEffect(() => {
    clearData();
  }, [routes]);

  const clearData = () => {
    setData([]);
    setShowTable(false);
    setCoach('');
  };

  const fetchData = (coachName) => {
    if (showTable && coach === coachName) {
      setShowTable(false);
      setCoach('');
    } else {
      axios
        .get(`http://localhost:8000/records/${coachName}`)
        .then((response) => {
          setData(response.data);
          setShowTable(true);
          setCoach(coachName);
        })
        .catch((error) => {
          console.error(error);
        });
    }
  };
  

  const sendData = (teamName)=> {
 
        axios
        .get(`http://localhost:8000/records/${teamName}`)
        .then ((response)=> {
          setData(response.data);
          setShowTable(true);
          setTeam(teamName);
        })
        .catch((error)=> {
        console.error(error);
        });
      // }
    };

    const handleFormSubmit = (e) => {
      e.preventDefault();
      sendData(team);
    };
  const fetchById = (id) => {
    axios
    .get(`http://localhost:8000/records/${id}`)
    .then ((response)=> {
      console.log(response.data)
      // console.log('teamssss')
      setPlayers(response.data);
      setShowTable(true);
      // setTeams(teamName);
    })
    .catch((error)=> {
    console.error(error);
    });
  }
  const fetchTeams = () => {
    axios
    .get(`http://localhost:8000/teams/`)
    .then ((response)=> {
      console.log(response.data)
      // console.log('teamssss')
      setTeams(response.data);
      setShowTable(true);
      // setTeams(teamName);
    })
    .catch((error)=> {
    console.error(error);
    });
  }
  const fetchPlayerStats = () => {

    axios
    .get(`http://localhost:8000/playerStats`)
    .then ((response)=> {
      console.log(response.data)
      // console.log('teamssss')
      setStats(response.data);
      setShowTable(true);
      // setTeams(teamName);
    })
    .catch((error)=> {
    console.error(error);
    });
  }

  const myFunction = () => {
    // Declare variables
    const input = document.getElementById("myInput");
    const filter = input.value.toUpperCase();
    const table = document.getElementById("mytable");
    const tr = table.getElementsByTagName("tr");
  
    // Loop through all table rows, and hide those that don't match the search query
    for (let i = 0; i < tr.length; i++) {
      const td = tr[i].getElementsByTagName("td")[0];
      if (td) {
        const txtValue = td.textContent || td.innerText;
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
          tr[i].style.display = "";
        } else {
          tr[i].style.display = "none";
        }
      }
    }
  };

  const openModal = (player) => {
    setSelectedPlayer(player);
  };

  const closeModal = () => {
    setSelectedPlayer(null);
  };

  
  

  
  return (
    <div >
      <div className='Nav-disp'>
        {/* <a href  ="#home" onClick={()=>{setRoutes('home')}}>Home</a> */}
        <a href ="#Teams" onClick={()=>{setRoutes('Teams');fetchTeams()}}>Teams</a>
        <a href ="#Players" onClick={()=>{setRoutes('Players');fetchPlayerStats()}}>Players</a>
        {/* <a href ="#Stats" onClick={()=>{setRoutes('stats')}}>Stats</a> */}
      </div>
   {(routes === 'Teams') && (
  <div>
 
    {/* {showTable && ( */}
    <input className="myInput" type = "text" id="myInput" onKeyUp={myFunction} placeholder="Search for names.."></input>
      <table className="myTable" id ="mytable">
        
          <tr classname="header">
            <th>Team Name</th>
            <th>Stadium</th>
          </tr>
          {teams.map((item, index) => (
            <tr key={index}>
              <td>{item.team_name}</td>
              <td>{item.Stadium}  </td>
              {/* <td>{item.last_name}</td> */}
            </tr>
          ))}
      </table>
      {}
    {/* )} */}
  </div>
)}

{(routes === 'Players') && (
  <div>
 
    {/* {showTable && ( */}
    <input className="myInput" type = "text" id="myInput" onKeyUp={myFunction} placeholder="Search for names.."></input>
      <table className="myTable" id ="mytable">
        
          <tr classname="header">
            <th>Player Name</th>
            <th>Nation</th>
            <th>Position</th>
            <th>Team</th>
            <th>Competition</th>
          </tr>
          {stats.map((item, index) => (
            <tr key={index} >
              <td>{item.Player}</td>
              <td>{item.Nation}  </td>
              <td>{item.Pos}</td>
              <td>{item.Squad}</td>
              <td>{item.Comp}</td>
              {/* <td>{item.last_name}</td> */}
            </tr>
          ))}
      </table>
      {}
    {/* )} */}
  </div>
)}


    
        
    </div>
  );
}

export default Stats;