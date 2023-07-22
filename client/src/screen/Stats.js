import React, { useState, useEffect } from 'react';
import axios from 'axios';
// import '../assets/styles/stats.css';
import Modal from 'react-modal';
import footballImage from '../assets/images/saka.png';
import sakaImage from '../assets/images/pitch.png';
import { getPlayer, getStat, getTeam } from '../apiRoutes';
const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
    width:'70%',
    background:'grey',
    height:'80%',
    zIndex: '1000'
  },
};
const Stats = () => {
  const [data, setData] = useState([]);
  const [showTable, setShowTable] = useState(false);
  const [routes, setRoutes] = useState('');
  const [teams, setTeams] = useState([]);
  const [players, setPlayers] = useState([]);
  const [stats, setStats] = useState([]);
  // const [selectedPlayer, setSelectedPlayer] = useState(null);
  // const []
  const [selectedPlayer, setSelectedPlayer] = useState({
    player_name: 'Bukayo Saka',
    games: '12',
    mins_played: '1234',
    goal:'12',
    assists:'12',
    shots:'12',
    key_passes: '1234',
    yellow_cards: '123',
    red_cards: '123',
    position:'Forward'
  })

  useEffect(() => {
    clearData();
  }, [routes]);

  const clearData = () => {
    setData([]);
    setShowTable(false);
  };
  const fetchTeams = () => {
    axios
    .get(`http://localhost:8080/teams`)
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

  const fetchPlayers = () => {
  axios
    .get(`http://localhost:8080/player`)
    .then((response) => {
      console.log(response.data);
      setPlayers(response.data);
      setShowTable(true);
    })
    .catch((error) => {
      console.error(error);
    });
};

const fetchStats = (player_name, player_id) => {
  axios
    .get(`http://localhost:8080/stats/${player_id}`)
    .then((response) => {
      console.log('fetchStats()');
      console.log(response.data);

      setStats(response.data);
      setShowTable(true);
      response.data[0].player_name = player_name;
      setSelectedPlayer(response.data[0]);
    })
    .catch((error) => {
      console.error(error);
    });
};


  const myFunction = () => {
    const input = document.getElementById('myInput');
    const filter = input.value.toUpperCase();
    const table = document.getElementById('mytable');
    const tr = table.getElementsByTagName('tr');

    for (let i = 0; i < tr.length; i++) {
      const td = tr[i].getElementsByTagName('td')[0];
      if (td) {
        const txtValue = td.textContent || td.innerText;
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
          tr[i].style.display = '';
        } else {
          tr[i].style.display = 'none';
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
    <div>
      <div className="Nav-disp">
        <a href="#Teams" onClick={() => { setRoutes('Teams'); fetchTeams(); }}>
          Teams
        </a>
        <a href="#Players" onClick={() => { setRoutes('Players'); fetchPlayers(); }}>
          Players
        </a>
        <a href="#Lineup" onClick={()=> {setRoutes('Lineup')}}>
          Lineup manager
        </a>
      </div>

      {routes === 'Teams' && (
        <div>
          <input className="myInput" type="text" id="myInput" onKeyUp={myFunction} placeholder="Search for names.." />
          <table className="myTable" id="mytable">
            <tr className="header">
              <th>Team Name</th>
              <th>Stadium</th>
            </tr>
            {teams.map((item, index) => (
              <tr key={index}>
                <td>{item.team_name}</td>
                <td>{item.stadium}</td>
              </tr>
            ))}
          </table>
        </div>
      )}

      {routes === 'Players' && (
        <div>
          <input className="myInput" type="text" id="myInput" onKeyUp={myFunction} placeholder="Search for names.." />
          <table className="myTable" id="mytable">
            <tr className="header">
              <th>Player Name</th>
              <th>Club</th>
            </tr>
            {players.map((item, index) => (
              <tr key={index} onClick={() => {openModal(item);fetchStats(item.player_name,item.player_id)}}>
                <td>{item.player_name}</td>
                <td>{item.team_name}</td>
              
              </tr>
            ))}
          </table>
        </div>
      )}  <div>
      <button onClick={openModal}>Open Modal</button>
      <Modal
        isOpen={true}
        // onAfterOpen={afterOpenModal}
        onRequestClose={closeModal}
        style={customStyles}
        contentLabel="Example Modal"
        // className="Model"
      >
       <div className="Model_Content">
          <div className="Model_Content_ImageBox">
            <img className="Model_Content_ImageBox_Image" src={footballImage} alt="Football" className="modal-image" />
          </div>
          <div className="Model_Content_CardContainer">
            <div className="Model_Content_CardContainer_Cards">
            <div className="card">
              <h3>Player Name</h3>
              <p>{selectedPlayer?.player_name}</p>
            </div>
            <div className="card">
              <h3>Games</h3>
              <p>{selectedPlayer?.games}</p>
            </div>
            <div className="card">
              <h3>Minutes Played</h3>
              <p>{selectedPlayer?.mins_played}</p>
            </div>
            <div className="card">
              <h3>Goals</h3>
              <p>{selectedPlayer?.goal}</p>
            </div>
            <div className="card">
              <h3>Assists</h3>
              <p>{selectedPlayer?.assists}</p>
            </div>
            </div>
            <div className="Model_Content_CardContainer_Cards">
           
            <div className="card">
              <h3>Shots</h3>
              <p>{selectedPlayer?.shots}</p>
            </div>
            <div className="card">
              <h3>Key Passes</h3>
              <p>{selectedPlayer?.key_passes}</p>
            </div>
            <div className="card">
              <h3>Yellow Cards</h3>
              <p>{selectedPlayer?.yellow_cards}</p>
            </div>
            <div className="card">
              <h3>Red Cards</h3>
              <p>{selectedPlayer?.red_cards}</p>
            </div>
            <div className="card">
              <h3>Position</h3>
              <p>{selectedPlayer?.position}</p>
            </div>
            </div>
          </div>
        </div>
      </Modal>
    </div>


      {/* {selectedPlayer && ( */}
        <Modal
        isOpen={true}
        onRequestClose={closeModal}
        contentLabel="Player Image"
        // style={customStyles}
        className="modal-container"
        overlayClassName="modal-overlay"
      >
        
      </Modal>
      
      {/* )} */}
    </div>
  );
};

export default Stats;
