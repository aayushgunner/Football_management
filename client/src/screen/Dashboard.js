import React, { useState } from 'react'
import Spinner from '../shared/Spinner';
import BootstrapTable from "react-bootstrap-table-next";
import paginationFactory from "react-bootstrap-table2-paginator";
import ToolkitProvider, { Search } from "react-bootstrap-table2-toolkit";
const { SearchBar } = Search;
const Dashboard = () => {
  // const [table, setTable] = useState([]);
  const [showLoader, setShowLoader] = useState(false)
  const   columns = [
    {
      dataField: "team_name",
      text: "club",
      sort: false,
    },
    {
      dataField: "played",
      text: "Played",
      sort: false,
    },
    {
      dataField: "won",
      text: "Won",
      sort: false,
    },
    {
      dataField: "drawn",
      text: "Drawn",
      sort: false,
    },
    {
      dataField: "lost",
      text: "Lost",
      sort: false,
    },
    {
      dataField: "gf",
      text: "GF",
      sort: false,
    },
    {
      dataField: "ga",
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
  const table = [
    {
      position: 1,
      team_name: 'Arsenal',
      played:100,
      won: 100,
      drawn: 0,
      lost:0,
      gf: 1200,
      ga:0,
      gd: 1200,
      points: 300
    },
    {
      position: 2,
      team_name: 'Manchestar city',
      played:100,
      won: 90,
      drawn:0,
      lost:10,
      gf: 100,
      ga:10,
      gd: 90,
      points: 270
    }
  ]
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
                                text: "Position",
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