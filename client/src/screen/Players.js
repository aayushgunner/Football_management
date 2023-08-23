import React, { useEffect, useState } from 'react'
import Spinner from '../shared/Spinner';
import BootstrapTable from "react-bootstrap-table-next";
import paginationFactory from "react-bootstrap-table2-paginator";
import ToolkitProvider, { Search } from "react-bootstrap-table2-toolkit";
import axios from 'axios';
import { useHistory } from 'react-router-dom/cjs/react-router-dom.min';
const { SearchBar } = Search;
const Players = () => {
    const [showLoader, setShowLoader] = useState(false)
    const [players, setPlayers] = useState([]);
    const history = useHistory()
    const fetchPlayers = () => {
      axios
        .get(`http://localhost:8080/player`)
        .then((response) => {
          console.log("Players is this ")
          console.log(response.data);
          setPlayers(response.data);
        })
        .catch((error) => {
          console.error(error);
        });
    };
    useEffect(()=>{
      fetchPlayers();
    },[])
    const   columns = [
        {
          dataField: "player_name",
          text: "Player Name",
          sort: true,
        },
        {
          dataField: "team_name",
          text: "Club",
          sort: true,
        },
        {
          dataField: "action",
          text: "Action",
          sort: false,
          formatter: (cell, row) => {
            
            return (
              <div className="d-flex">
                <button
                  className="btn btn-primary py-1 px-2 btn-icon-text btn-rounded"
                  onClick={() => {
                    console.log(row)
            console.log('action')
                    history.push(`/playerStats/${row.player_id}`)
                  }}
                >
                  <i className="mdi mdi-pencil btn-icon-prepend"></i>
                </button>
               
              </div>
            );
          },
        }
      ];
  return (
    <div className="row">
    <div className="col-12">
      {showLoader && <Spinner />}
      {!showLoader && players.length > 0 && (
        <div className="card">
          <div className="card-body">
            <div className="row">
              <div className="col-12">
                <ToolkitProvider
                  keyField="_id"
                  bootstrap4
                  data={players}
                  columns={columns}
                  search
                >
                  {(props) => {
                    const data = props.baseProps.data.reverse()// Invertir el orden de los datos

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
                                text: "#",
                                formatter: (cell, row, rowIndex) =>
                                data.indexOf(row) + 1,
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
      {!showLoader && players.length === 0 && (
        <div className="text-center">
          <p className="text-primary">No players found.</p>
        </div>
      )}
    </div>
  </div>
  )
}

export default Players