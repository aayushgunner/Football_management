import React, {useState, useEffect} from "react";
import { Dropdown } from "react-bootstrap";
import { Link, } from "react-router-dom";

// import { Trans } from "react-i18next";
// import './navbar.scss'
const Navbar = () => {
//   const location = useLocation();
  const [pageTitle, setPageTitle] = useState("");
//   useEffect(() => {
//     // switch (location.pathname) {
//       case "/":
//         setPageTitle("Dashboard");
//         break;
//       case "/all-users":
//         setPageTitle("All users");
//         break;
//       case "/companies":
//         setPageTitle("All companies");
//         break;
//       case "/error-codes":
//         setPageTitle("All error codes");
//         break;
//       case "/subscription-plans":
//         setPageTitle("All subscription plans");
//         break;
//       case "/coupons":
//         setPageTitle("All subscription coupons");
//         break;
//       case "/compliance-forms":
//         setPageTitle("All Forms");
//         break;
//       case "/manage-roles":
//         setPageTitle("All roles");
//         break;
//       default:
//         setPageTitle("");
//         break;

//     }
//   }, [location.pathname]);
  const toggleOffcanvas = () => {
    document.querySelector(".sidebar-offcanvas").classList.toggle("active");
  };
  return (
    <nav className="navbar default-layout-navbar col-lg-12 col-12 p-0 fixed-top d-flex flex-row">
      <div className="text-center navbar-brand-wrapper d-flex align-items-center justify-content-center">
        <Link className="navbar-brand brand-logo" to="/">
          <img src={require("../assets/images/premierleague.png")} alt="logo" style={{height: '55px'}}/>
        </Link>
        <Link className="navbar-brand brand-logo-mini" to="/">
          {/* <img src={require("../../assets/images/logo-mini.svg")} alt="logo" /> */}
        </Link>
      </div>
      <div className="navbar-menu-wrapper d-flex align-items-stretch">
        <button
          className="navbar-toggler navbar-toggler align-self-center"
          type="button"
          onClick={() => document.body.classList.toggle("sidebar-icon-only")}
        >
          <span className="mdi mdi-menu"></span>
        </button>
        <div className="top_header_title">
          <h1>{pageTitle}</h1>
        </div>

        <ul className="navbar-nav navbar-nav-right">
          <li className="nav-item nav-profile">
            <Dropdown alignRight>
              <Dropdown.Toggle className="nav-link">
                <div className="nav-profile-img">
                  {/* <img
                    // src={require("../../assets/images/faces/face1.jpg")}
                    alt="user"
                  /> */}
                  <span className="availability-status online"></span>
                </div>
                <div className="nav-profile-text">
                  {/* {user && (
                    <p className="mb-1 text-black">
                      <Trans>{user.firstName}</Trans>
                    </p>
                  )} */}
                </div>
              </Dropdown.Toggle>

              {/* <Dropdown.Menu className="navbar-dropdown">
                <Dropdown.Item href="!#" onClick={(evt) => logoutClicked(evt)}>
                  <i className="mdi mdi-logout mr-2 text-primary"></i>
                  <Trans>Signout</Trans>
                </Dropdown.Item>
              </Dropdown.Menu> */}
            </Dropdown>
          </li>
        </ul>
        <button
          className="navbar-toggler navbar-toggler-right d-lg-none align-self-center"
          type="button"
          onClick={toggleOffcanvas}
        >
          <span className="mdi mdi-menu"></span>
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
