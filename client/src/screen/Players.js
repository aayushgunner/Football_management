// import React, { useState } from 'react'
// import Spinner from '../shared/Spinner';
// import BootstrapTable from "react-bootstrap-table-next";
// import paginationFactory from "react-bootstrap-table2-paginator";
// import ToolkitProvider, { Search } from "react-bootstrap-table2-toolkit";
// const { SearchBar } = Search;
// const Players = () => {
//     const [showLoader, setShowLoader] = useState(false)
//     const users = [
//         {
//             firstName:'helloworld',
//             email: 'pasang@gmail.com',
//             status:'Active',
//             isConfirmed: 'Open'
//         },
//         {
//             firstName:'helloworld1',
//             email: 'pasang1@gmail.com',
//             status:'Active',
//             isConfirmed: 'Open'
//         },
//         {
//             firstName:'helloworld2',
//             email: 'pasang2@gmail.com',
//             status:'Active',
//             isConfirmed: 'Open'
//         }
//     ]
//     const   columns = [
//         {
//           dataField: "firstName",
//           text: "Name",
//           sort: true,
//         },
//         {
//           dataField: "email",
//           text: "Email",
//           sort: true,
//         },
//         {
//           dataField: "status",
//           text: "Status",
//           sort: true,
//           formatter: (cellContent, row) => {
//             if (cellContent === 0) {
//               return <label className="badge badge-danger">In-active</label>;
//             } else if (cellContent === 1) {
//               return <label className="badge badge-success">Active</label>;
//             } else if (cellContent === "Closed") {
//               return <label className="badge badge-success">Closed</label>;
//             } else if (cellContent === "Open") {
//               return <label className="badge badge-warning">Open</label>;
//             }
//           },
//         },
//         {
//           dataField: "isConfirmed",
//           text: "Verified",
//           sort: true,
//           formatter: (cellContent, row) => {
//             if (cellContent === false) {
//               return <label className="badge badge-danger">No</label>;
//             } else if (cellContent === true) {
//               return <label className="badge badge-success">Yes</label>;
//             } else if (cellContent === "Closed") {
//               return <label className="badge badge-success">Closed</label>;
//             } else if (cellContent === "Open") {
//               return <label className="badge badge-warning">Open</label>;
//             }
//           },
//         },
//         {
//           dataField: "action",
//           text: "Action",
//           sort: false,
//           formatter: (cell, row) => {
//             return (
//               <div className="d-flex">
//                 <button
//                   className="btn btn-primary py-1 px-2 btn-icon-text btn-rounded"
//                   onClick={() => {
//                     // assignUserCompany(row._id)
//                     //   .then((res) => {
//                     //     console.log(res);
//                     //   })
//                     //   .catch((err) => {
//                     //     console.log(err);
//                     //   });
//                   }}
//                 >
//                   {/*<i className="mdi mdi-pencil btn-icon-prepend"></i>*/}Login
//                 </button>

//                 <button
//                   className="btn btn-info py-1 px-2 btn-icon-text btn-rounded"
//                 //   onClick={() => this.viewCompaniesList(row)}
//                 >
//                   <i className="mdi mdi-eye btn-icon-prepend"></i>
//                   Companies list
//                 </button>
//                 <button
//                   className="btn btn-primary py-1 px-2 btn-icon-text btn-rounded pen_btn"
//                   onClick={() => this.editUser(row)}
//                 >
//                   <img src={require("../assets/images/pen.svg")} />
//                 </button>
//                 <button
//                   className="btn btn-danger py-1 px-2 btn-icon-text btn-rounded delete_btn"
//                 //   onClick={() =>
//                 //     // confirmDialog({
//                 //     //   message: "Do you want to delete this record?",
//                 //     //   header: "Delete Confirmation",
//                 //     //   icon: "pi pi-info-circle",
//                 //     //   acceptClassName: "p-button-danger",
//                 //     //   accept: () => {
//                 //     //     this.deleteUser(row);
//                 //     //   },
//                 //     // })
//                 //   }
//                 >
//                   <img src={require("../assets/images/trash.svg")} />
//                 </button>
//               </div>
//             );
//           },
//         },
//       ];
//   return (
//     <div className="row">
//     <div className="col-12">
//       {showLoader && <Spinner />}
//       {!showLoader && users.length > 0 && (
//         <div className="card">
//           <div className="card-body">
//             <div className="row">
//               <div className="col-12">
//                 <ToolkitProvider
//                   keyField="_id"
//                   bootstrap4
//                   data={users}
//                   columns={columns}
//                   search
//                 >
//                   {(props) => {
//                     const data = props.baseProps.data.reverse(); // Invertir el orden de los datos

//                     return (
//                       <div>
//                         <div className="d-flex align-items-center">
//                           <div className="card flex flex-wrap justify-content-center-top gap-3 padding_bot">
//                             {/* <button
//                               type="button"
//                               onClick={() => {
//                                 // openNewUserModal();
//                               }}
//                               className="btn btn-gradient-primary btn-rounded btn-fw btn-icon-text"
//                             >
//                               <i className="mdi mdi-account-plus btn-icon-prepend"></i>
//                               Add User
//                             </button> */}
//                             {/* <button
//                               type="button"
//                               onClick={() => {
//                                 // _this.openGenerateLinkModal();
//                               }}
//                               className="btn btn-gradient-primary btn-rounded btn-fw btn-icon-text"
//                             >
//                               Generate link
//                             </button> */}
//                             <span className="p-input-icon-left">
//                               <i className="pi pi-search" />
//                               <SearchBar {...props.searchProps} />
//                             </span>
//                           </div>
//                         </div>
//                         <div className="talbe_redesign">
//                           <BootstrapTable
//                             pagination={paginationFactory()}
//                             {...props.baseProps}
//                             wrapperClasses="table-responsive"
//                             columns={[
//                               {
//                                 dataField: "",
//                                 text: "#",
//                                 formatter: (cell, row, rowIndex) =>
//                                   data.length - data.indexOf(row),
//                                 headerAlign: "center",
//                                 align: "center",
//                               },
//                               ...props.baseProps.columns,
//                             ]}
//                             data={data}
//                           />
//                         </div>
//                       </div>
//                     );
//                   }}
//                 </ToolkitProvider>
//               </div>
//             </div>
//           </div>
//         </div>
//       )}
//       {!showLoader && users.length === 0 && (
//         <div className="text-center">
//           <p className="text-primary">No members found.</p>
//         </div>
//       )}
//     </div>
//   </div>
//   )
// }

// export default Players

import React, { useState, useEffect } from "react";
import axios from "axios";
import "../assets/styles/table.css";
const Players = () => {
  const [data, setData] = useState([]);
  const [teams, setTeams] = useState([]);
  const [players, setPlayers] = useState([]);
  const [showtable, setShowTable] = useState([]);
  const [selectedteamPhoto, setSelectedTeamPhoto] = useState("");


  useEffect(() => {
    fetchTeams(); // Fetch teams when the component mounts
  }, []);
  const clearData = () => {
    setData([]);
  };

  const fetchById = (id) => {
    axios
      .get(`http://localhost:8000/records/${id}`)
      .then((response) => {
        console.log(response.data);
        // console.log('teamssss')
        setPlayers(response.data);
        setShowTable(true);
        setSelectedTeamPhoto(response.data.photo);
        // setTeams(teamName);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  const fetchTeams = () => {
    axios
      .get(`http://localhost:8000/teams/`)
      .then((response) => {
        console.log(response.data);
        // console.log('teamssss')
        setTeams(response.data);
        setShowTable(true);
        // setTeams(teamName);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  

  return (
    <>
    <body>
      <div >{/* <img src={require("../assets/images/saka.png")}></img> */}</div>
      <table>
        <thead>
          <tr>
            <th>Team Name</th>
          </tr>
        </thead>
        <tbody>
          {teams.map((item, index) => (
            <tr key={index} onClick={()=> fetchById(item.team_id)}>
              <td>{item.team_name}</td>
              {/* <td>{item.last_name}</td> */}
            </tr>
          ))}
        </tbody>
      </table>
    
    {players.length > 0 && 
      <table>
        <thead>
          <tr>
            <th>First Name</th>
            <th>Last Name</th>
          </tr>
        </thead>
        <tbody>
          {players.map((item, index) => (
            <tr key={index} >
              <td>{item.first_name}</td>
              <td>{item.last_name}</td>
            </tr>
          ))}
        </tbody>
      </table>
    }
      </body>
    </>
  );
};

export default Players;
