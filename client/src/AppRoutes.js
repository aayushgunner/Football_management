import React, { Component, Suspense, lazy } from "react";
import { Switch, Route, Redirect } from "react-router-dom";
import Spinner from "./shared/Spinner";
const Login = lazy(() => import("./screen/Login"));
const AppRoutes = () => {
  return (
    <Suspense fallback={<Spinner />}>
  
      <Switch>
        <Route path="/login" component={Login} />
        <Redirect to="/login" />
      </Switch>

  </Suspense>
  )
}

export default AppRoutes