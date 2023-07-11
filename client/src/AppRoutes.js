import React, { Component, Suspense, lazy } from "react";
import { Switch, Route, Redirect } from "react-router-dom";
import Spinner from "./shared/Spinner";
const Login = lazy(() => import("./screen/Login"));
const Dashboard = lazy(()=> import('./screen/Dashboard'));
const Players = lazy(()=> import('./screen/Players'))
const Stats = lazy(()=> import('./screen/Stats'))

const AppRoutes = () => {
  return (
    <Suspense fallback={<Spinner />}>
  
      <Switch>
        {/* <Route path="/login" component={Login} /> */}
        <Route path="/dashboard" component = {Dashboard}/>
        <Route path="/players" component={Players}/>
        <Route path="/Stats" component={Stats}/>
        <Redirect to="/dashboard" />
      </Switch>

  </Suspense>
  )
}

export default AppRoutes