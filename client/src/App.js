import React from 'react';
import { Route, Switch, BrowserRouter } from 'react-router-dom'
import { Home, Register, Verify, NotFound, Login } from './components';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/register" component={Register} />
        <Route exact path="/login" component={Login} />
        <Route exact path="/verify/:userID" component={Verify} />
        <Route component={NotFound} />
      </Switch>
    </BrowserRouter>
  )
}

export default App;
