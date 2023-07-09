import React from 'react'
import { Link, withRouter } from "react-router-dom";
import { Trans } from "react-i18next";
const Sidebar = () => {
    const isPathActive = (path) => {
        return false;
    }
  return (
    <nav className="sidebar sidebar-offcanvas" id="sidebar">
<ul className="nav">
      <li className="nav-item nav-profile">
        <a
          href="!#"
          className="nav-link"
          onClick={(evt) => evt.preventDefault()}
        >
          <div className="nav-profile-image">
            {/* <img
              src={require("../../assets/images/faces/face1.jpg")}
              alt="profile"
            /> */}
            <span className="login-status online"></span>{" "}
            {/* change to offline or busy as needed */}
          </div>
     
            <div className="nav-profile-text">
              <span className="font-weight-bold mb-2">
                <Trans>Hello User</Trans>
              </span>
              <span className="text-secondary text-small">
                <Trans>Admin</Trans>
              </span>
            </div>
          <i className="mdi mdi-bookmark-check text-success nav-profile-badge"></i>
        </a>
      </li>
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
          isPathActive("/all-users") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/all-users">
          <span className="menu-title">User Management</span>
          <i className="mdi mdi-account-multiple menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/companies") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/companies">
          <span className="menu-title">Companies</span>
          <i className="mdi mdi-briefcase menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/error-codes") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/error-codes">
          <span className="menu-title">Error Codes</span>
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
          <span className="menu-title">Subscription Plans</span>
          <i className="mdi mdi-currency-gbp menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/coupons") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/coupons">
          <span className="menu-title">Coupon Codes</span>
          <i className="mdi mdi-currency-gbp menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/compliance-forms")
            ? "nav-item active"
            : "nav-item"
        }
      >
        <Link className="nav-link" to="/compliance-forms">
          <span className="menu-title">Compliance Forms</span>
          <i className="mdi mdi-currency-gbp menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/manage-roles") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/manage-roles">
          <span className="menu-title">Manage Roles</span>
          <i className="mdi mdi-account-multiple menu-icon"></i>
        </Link>
      </li>
      <li
        className={
          isPathActive("/email-sequencing") ? "nav-item active" : "nav-item"
        }
      >
        <Link className="nav-link" to="/email-sequencing">
          <span className="menu-title">Email sequencing</span>
          <i className="mdi mdi-briefcase menu-icon"></i>
        </Link>
      </li>
     
    </ul>
  </nav>
  );
}

export default Sidebar