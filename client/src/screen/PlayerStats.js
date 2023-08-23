import React, { useEffect, useState } from 'react'
import {useParams} from 'react-router-dom'
import footballImage from '../assets/images/saka.png';
import axios from 'axios';

const PlayerStats = () => {
    const [selectedPlayer, setSelectedPlayer] = useState(null);
    const {id} = useParams();
    const fetchStats = (player_name,) => {
        axios
          .get(`http://localhost:8080/stats/${id}`)
          .then((response) => {
            setSelectedPlayer(response.data[0]);
          })
          .catch((error) => {
            console.error(error);
          });
      };
    useEffect(()=>{
        fetchStats()
    },[])
  return (
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
  )
}

export default PlayerStats