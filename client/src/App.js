import React from 'react';
import './App.css';
import dealFinder from "./dealFinder"
import {
  BrowserRouter as Router,
  Route, Switch
} from "react-router-dom"

function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component = {dealFinder}/>
      </Switch>
    </Router>
  );
}

export default App;
