import React, { useEffect, useState } from 'react'
import Spinner from '../shared/Spinner';
import BootstrapTable from "react-bootstrap-table-next";
import paginationFactory from "react-bootstrap-table2-paginator";
import ToolkitProvider, { Search } from "react-bootstrap-table2-toolkit";
import axios from 'axios';
const { SearchBar } = Search;
const Dashboard = () => {
  const [table, setTable] = useState([]);
  const [showLoader, setShowLoader] = useState(false)
  const fetchTable = () => {
    axios
      .get(`http://localhost:8080/standings`)
      .then((response) => {
        console.log("Players is this ")
        console.log(response.data);
        setTable(response.data);
      })
      .catch((error) => {
        console.error(error);
      });
  };
  useEffect(()=>{
    fetchTable()
  },[])

  const   columns = [
    {
      dataField: "team_name",
      text: "club",
      sort: false,
    },
    {
      dataField: "matches_played",
      text: "Played",
      sort: false,
    },
    {
      dataField: "wins",
      text: "Won",
      sort: false,
    },
    {
      dataField: "draws",
      text: "Drawn",
      sort: false,
    },
    {
      dataField: "losses",
      text: "Lost",
      sort: false,
    },
    {
      dataField: "goals_for",
      text: "GF",
      sort: false,
    },
    {
      dataField: "goals_against",
      text: "GA",
      sort: false,
    },
    {
      dataField: "gd",
      text: "GD",
      sort: false,
    },
    {
      dataField: "points",
      text: "Points",
      sort: false,
    },
  ];

  return (
    <div className="row">
    <div className="col-12">
      {showLoader && <Spinner />}
      {!showLoader && table.length > 0 && (
        <div className="card">
          <div className="card-body">
            <div className="row">
              <div className="col-12">
                <ToolkitProvider
                  keyField="_id"
                  bootstrap4
                  data={table}
                  columns={columns}
                  search
                >
                  {(props) => {
                    // const data = props.baseProps.data// Invertir el orden de los datos
                    let data = props.baseProps.data.map(obj => {
                      console.log("Original Object:", obj);
                      let newObj = {
                        ...obj,
                        'gd': (obj.goals_for - obj.goals_against)
                      };
                      console.log("Updated Object:", newObj);
                      return newObj;
                    });

                    return (
                      <div>
                        <div className="d-flex align-items-center">
                          <div className="card flex flex-wrap justify-content-center-top gap-3 padding_bot">
                            <span className="p-input-icon-left">
                              <i className="pi pi-search" />
                              <SearchBar {...props.searchProps} />
                            </span>
                          </div>
                        </div>
                        <div className="talbe_redesign">
                          <BootstrapTable
                            pagination={paginationFactory()}
                            {...props.baseProps}
                            wrapperClasses="table-responsive"
                            columns={[
                              {
                                dataField: "",
                                text: "Position",
                                formatter: (cell, row, rowIndex) =>
                                data.indexOf(row) + 1,
                                headerAlign: "center",
                                align: "center",
                              },
                              {
                                dataField: "team_logo",
                                text: "",
                                formatter: (cell, row, rowIndex) =>
                                <img className='image_for_team_logo' src={row.team_logo}></img>,
                                headerAlign: "center",
                                align: "center",
                              },
                              ...props.baseProps.columns,
                            ]}
                            data={data}
                          />
                        </div>
                      </div>
                    );
                  }}
                </ToolkitProvider>
              </div>
            </div>
          </div>
        </div>
      )}
      {!showLoader && table.length === 0 && (
        <div className="text-center">
          <p className="text-primary">No table found.</p>
        </div>
      )}
    </div>
  </div>
  )
}

export default Dashboard