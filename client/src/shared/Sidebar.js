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
            <Trans>Home</Trans>
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
          <span className="menu-title">Fixtures</span>
          <i className="mdi mdi-account-multiple menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/companies") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/companies">
          <span className="menu-title">Results</span>
          <i className="mdi mdi-briefcase menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/error-codes") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/error-codes">
          <span className="menu-title">Tables</span>
          <i className="mdi mdi-information-variant menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/subscription-plans")
            ? "nav-item active"
            : "nav-item"
        }
      >
        <Link className="nav-link" to="/subscription-plans">
          <span className="menu-title">Transfers</span>
          <i className="mdi mdi-currency-gbp menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/coupons") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/coupons">
          <span className="menu-title">Stats</span>
          <i className="mdi mdi-currency-gbp menu-icon"></i>
        </Link>
      </li>
     
    </ul>
  </nav>
  );
}

export default Sidebar