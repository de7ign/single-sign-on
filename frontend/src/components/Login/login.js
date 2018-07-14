import React, { Component } from 'react';
import './login.css'
export default class Login extends Component {
    render() {
        return (
            <div class="login">
                <a class="button-primary button" href="http://localhost:5000/v1/api/auth/google">Login with google</a>
                <a class="button-primary button" href="http://localhost:5000/v1/api/auth/github">Login with github</a>
            </div>
        )
    }
}
