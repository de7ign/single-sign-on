import React, { Component } from 'react';
import './App.css';
import { BrowserRouter, Route } from 'react-router-dom'
import Login from './components/Login/login.js'
import Dashboard from './components/Dashboard/dashboard.js'

class App extends Component {
  render() {
    return (
      <BrowserRouter>
      <div>
        <Route exact path="/" component={Login} />
        <Route exact path="/dashboard" component={Dashboard} />
      </div>
    </BrowserRouter>
    );
  }
}

export default App;
