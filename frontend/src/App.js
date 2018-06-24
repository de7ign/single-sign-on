import React, { Component } from 'react';
import './App.css';

class App extends Component {
  render() {
    return (
      <div class="container">
        <div class="row">
          <div class="one-third column">
            &nbsp;
          </div>
          <div class="one-third column">
            <a class="button-primary button" href="http://localhost:5000/v1/api/auth/google">Login with google</a>
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
