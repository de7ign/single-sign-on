import React, { Component } from 'react';
import './App.css';
import Login from './components/Login/login.js'

class App extends Component {
  render() {
    return (
      <div class="container">
        <div class="row">
          <div class="one-third column">
            &nbsp;
          </div>
          <div class="one-third column">
            <Login />
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
