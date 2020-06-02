import React from 'react';
import { Route, Switch, BrowserRouter } from 'react-router-dom'
import { Home, Register, Verify, NotFound, Login, Profile } from './components';
import { UserProfileProvider } from "./context/user"

function App() {
  return (
    <UserProfileProvider>
      <BrowserRouter>
        <Switch>
          <Route exact path="/" component={Home} />
          <Route exact path="/register" component={Register} />
          <Route exact path="/login" component={Login} />
          <Route exact path="/verify/:userID" component={Verify} />
          <Route exact path="/profile" component={Profile} />
          <Route component={NotFound} />
        </Switch>
      </BrowserRouter>
    </UserProfileProvider>
  )
}

export default App;
