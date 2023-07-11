import React, { useState, useEffect } from 'react'
import Spinner from '../shared/Spinner';
import BootstrapTable from "react-bootstrap-table-next";
import paginationFactory from "react-bootstrap-table2-paginator";
import ToolkitProvider, { Search } from "react-bootstrap-table2-toolkit";
  import { getPlayer } from '../apiRoutes';
  
const { SearchBar } = Search;
const Players = () => {
    const [showLoader, setShowLoader] = useState(false)
    const [users, setUsers] = useState([])
    useEffect(()=>{
      getPlayer().then((res)=>{
        setUsers(res)
      })
    },[])
    const   columns = [
        {
          dataField: "first_name",
          text: "FirstName",
          sort: true,
        },
        {
          dataField: "last_name",
          text: "LastName",
          sort: true,
        },
        {
          dataField: "position",
          text: "Position",
          sort: true,
        }
      ];
  return (
    <div className="row">
    <div className="col-12">
      {showLoader && <Spinner />}
      {!showLoader && users.length > 0 && (
        <div className="card">
          <div className="card-body">
            <div className="row">
              <div className="col-12">
                <ToolkitProvider
                  keyField="_id"
                  bootstrap4
                  data={users}
                  columns={columns}
                  search
                >
                  {(props) => {
                    const data = props.baseProps.data

                    return (
                      <div>
                        <div className="d-flex align-items-center">
                          <div className="card flex flex-wrap justify-content-center-top gap-3 padding_bot">
                            {/* <button
                              type="button"
                              onClick={() => {
                                // openNewUserModal();
                              }}
                              className="btn btn-gradient-primary btn-rounded btn-fw btn-icon-text"
                            >
                              <i className="mdi mdi-account-plus btn-icon-prepend"></i>
                              Add User
                            </button> */}
                            {/* <button
                              type="button"
                              onClick={() => {
                                // _this.openGenerateLinkModal();
                              }}
                              className="btn btn-gradient-primary btn-rounded btn-fw btn-icon-text"
                            >
                              Generate link
                            </button> */}
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
                                  data.length - data.indexOf(row),
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
      {!showLoader && users.length === 0 && (
        <div className="text-center">
          <p className="text-primary">No members found.</p>
        </div>
      )}
    </div>
  </div>
  )
}

export default Players