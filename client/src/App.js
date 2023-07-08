// import React, { useEffect, useState } from 'react';
// import axios from 'axios';
// import './App.css';
//   function App() {
//     const [data, setData] = useState([]);
//     const [showTable, setShowTable] = useState(false);
  

//   const fetchData = () => {
//     if (showTable) {
//       setShowTable(false);
//     }
//     else {
//     axios
//       .get('http://localhost:8000/records')
//       .then((response) => {
//         setData(response.data);
//         setShowTable(true);
//       })
//       .catch((error) => {
//         console.error(error);
//       });
//     }
//   };
//   return (
//     <div className="background_lol">
//       <h1>Player Information</h1>
//       <button onClick={fetchData}>
//         {showTable ? 'Hide Table': 'Show Table'}
//         </button>
//       {showTable && (
//         <table>
//           <thead>
//             <tr>
//               <th>First Name</th>
//               <th>Last Name</th>
//             </tr>
//           </thead>
//           <tbody>
//             {data.map((item, index) => (
//               <tr key={index}>
//                 <td>{item.first_name}</td>
//                 <td>{item.last_name}</td>
//               </tr>
//             ))}
//           </tbody>
//         </table>
//       )}
//     </div>
//   );
//             }
  

  
// //}

// export default App;


import React, { useState } from 'react';
import axios from 'axios';
import './App.css';
import './navi.css'
function App() {
  const [data, setData] = useState([]);
  const [showTable, setShowTable] = useState(false);
  const [coach, setCoach] = useState('');
  const [routes, setRoutes] = useState('');

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

  return (
    <div className="haha">
      <h1>Football League</h1>
      <div className='Nav-disp'>
        <a href  ="#home" onClick={()=>{setRoutes('home')}}>Home</a>
        <a href ="#fixtures" onClick={()=>{setRoutes('fixtures')}}>Fixtures</a>
        <a href ="#Results" onClick={()=>{setRoutes('results')}}>Resuts</a>
        <a href ="#Stats" onClick={()=>{setRoutes('stats')}}>Stats</a>
      </div>
      {(routes === 'stats') && 
      (<>
      <button onClick={() => fetchData('Mikel')}>
        {showTable && coach === 'Mikel' ? 'Hide Table' : 'Show Mikel'}
      </button>
      <button onClick={() => fetchData('Pep')}>
        {showTable && coach === 'Pep' ? 'Hide Table' : 'Show Pep'}
      </button>
      {showTable && (
        <table>
          <thead>
            <tr>
              <th>First Name</th>
              <th>Last Name</th>
            </tr>
          </thead>
          <tbody>
            {data.map((item, index) => (
              <tr key={index}>
                <td>{item.first_name}</td>
                <td>{item.last_name}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}</>
      )} 
    </div>
  );
}

export default App;