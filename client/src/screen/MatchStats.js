import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom/cjs/react-router-dom';
import axios from 'axios';
const MatchStats = () => {
    const {id} = useParams()
    const [data, setData] = useState(null)
    const getMatchStats = () => {
        axios
          .get(`http://localhost:8080/fixtures/${id}`)
          .then((response) => {
            setData(response.data);
          })
          .catch((error) => {
            console.error(error);
          });
      };
      useEffect(() => {
        getMatchStats();
      }, []);
  return (
    <div className='MatchStats'>
        <h3>Match Stats</h3>
        <div className='MatchStats_Box'>
            <div className='MatchStats_Box_Heading'>
                <div className='MatchStats_Box_Heading_Left'>
                    <p>{data?.homeFixtures[0].team_home_name}</p>
                    
                </div>
                <div className='MatchStats_Box_Heading_Box'></div>
                <div className='MatchStats_Box_Heading_Right'>
                    <p>{data?.awayFixtures[0].team_away_name}</p>
                </div>
            </div>
            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].possession}</p>
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Possession %</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].possession}</p>
                </div>
            </div>
          	
            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].Shots_on_goal}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Shots On Goal</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].Shots_on_goal}</p>
                </div>
            </div>

            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].total_shots}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Total Shots</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].total_shots}</p>
                </div>
            </div>

            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].total_passes}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Total Passes</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].total_passes}</p>
                </div>
            </div>

            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].corner_kicks}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Corner Kicks</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].corner_kicks}</p>
                </div>
            </div>

            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].offsides}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Offsides</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].offsides}</p>
                </div>
            </div>

            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].yellow_cards}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Yellow Cards</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].yellow_cards}</p>
                </div>
            </div>

            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].red_cards}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Red Cards</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].red_cards}</p>
                </div>
            </div>
            <div className='MatchStats_Box_Bottom'>
                <div className='MatchStats_Box_Bottom_Left'>
                <p>{data?.homeFixtures[0].fouls}</p>
                    
                </div>
                <div className='MatchStats_Box_Bottom_Box'>
                    <p>Fouls</p>
                </div>
                <div className='MatchStats_Box_Bottom_Right'>
                <p>{data?.awayFixtures[0].fouls}</p>
                </div>
            </div>
          
        </div>
    </div>
  )
}

export default MatchStats