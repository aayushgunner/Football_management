import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
  const [data, setData] = useState([]);

  useEffect(() => {
    // Fetch data from the API
    axios.get('http://localhost:8000/records')
      .then(response => {
        setData(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);

  return (
    <div className="container">
      <h1>JSON Data Table</h1>
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
    </div>
  );
}

export default App;
