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


  import React, { useState , useEffect} from 'react';
  import axios from 'axios';
  // import './App.css';
  import './navi.css'
  import './App.scss'
import Navbar from './shared/Navbar';
import Sidebar from './shared/Sidebar';
import AppRoutes from "./AppRoutes";
import 'primereact/resources/themes/saga-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
  function App() {
    const [data, setData] = useState([]);
    const [showTable, setShowTable] = useState(false);
    const [coach, setCoach] = useState('');
    const [routes, setRoutes] = useState('');
    const [team , setTeam] = useState('');

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
      // console.log(teamName)
      // if (showTable && team ===teamName) {
      //   console.log('this')
      //   setShowTable(false);
      //   setTeam('');
      // }
      //   else {
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
      

    
    return (
      <>
       <div className="container-scroller">
      <Navbar/>
      <div className="container-fluid page-body-wrapper">
          <Sidebar/>
          <div className="main-panel">
            <div className="content-wrapper">
              <AppRoutes />
            </div>
          </div>
        </div>
      </div>
      </>
    );
  }

  export default App;