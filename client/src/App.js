import React from 'react';
import { Redirect, Route, Switch, BrowserRouter } from 'react-router-dom'
import { Register, NotFound } from './components';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Register} />
        <Route component={NotFound} />
      </Switch>
    </BrowserRouter>
  )
}

export default App;
