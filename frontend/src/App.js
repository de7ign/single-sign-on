import React, { Component } from 'react';
import './App.css';
import { BrowserRouter, Route } from 'react-router-dom'
import Login from './components/Login/login.js'
import Dashboard from './components/Dashboard/dashboard.js'

class App extends Component {
  render() {
    return (
      <div class="container">
        <div class="row">
          <div class="one-third column">
            &nbsp;
          </div>
          <div class="one-third column">
            <BrowserRouter>
              <div>
                <Route exact path="/" component={Login} />
                <Route exact path="/dashboard" component={Dashboard} />
              </div>
            </BrowserRouter>
          </div>
          <div class="one-third column">
            &nbsp;
          </div>
        </div>
      </div>
    );
  }
}

export default App;
