import React from 'react'
import { Link, withRouter, useLocation } from "react-router-dom";
import { Trans } from "react-i18next";

const Sidebar = () => {
    const location = useLocation()
    const isPathActive = (path) => {
        return location.pathname.startsWith(path);;
    }
  return (
    <nav className="sidebar sidebar-offcanvas" id="sidebar">
<ul className="nav">
      <li
        className={
          isPathActive("/dashboard") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/dashboard">
          <span className="menu-title">
            <Trans>Dashboard</Trans>
          </span>
          <i className="mdi mdi-home menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/players") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/players">
          <span className="menu-title">Player Management</span>
          <i className="mdi mdi-account-multiple menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/teams") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/teams">
          <span className="menu-title">Teams</span>
          <i className="mdi mdi-briefcase menu-icon"></i>
        </Link>
      </li>
    </ul>
  </nav>
  );
}

export default Sidebar