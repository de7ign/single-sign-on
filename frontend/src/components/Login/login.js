import React, { Component } from 'react';
import { Redirect } from 'react-router-dom'
import './login.css'
import axios from 'axios'
axios.defaults.withCredentials = true;
export default class Login extends Component {
    constructor(props){
        super(props);
        this.state = {
            authorized  :   false,
            HostURI :   window.location.hostname
        }

        axios.post('http://localhost:5000/v1/api/userinfo')
        .then((res)=>{
            this.setState({
                authorized  :   true
            })
        })
        .catch((err)=>{
            this.setState({
                authorized  :   false
            })
        })
    }

    render() {
        let authorized = this.state.authorized;
        if (authorized) {
            return(
                <Redirect to="/dashboard" />
            )
        }
        return (
            <div class="container login">
                <div class="row">
                    <h1>
                        Welcome to seminar demo
                    </h1>
                    <h2>
                        You can try yourself by going to {this.state.HostURI}
                    </h2>
                </div>
                <br /><br />
                <div className="row">
                    <div className="one-third column">
                        &nbsp;
                    </div>
                    <div className="one-third column">
                        <a class="button-primary button u-full-width" href="http://localhost:5000/v1/api/auth/google">Login with google</a>
                        <a class="button-primary button u-full-width" href="http://localhost:5000/v1/api/auth/github">Login with github</a>
                    </div>
                    <div className="one-third column">
                        &nbsp;
                    </div>
                </div>
            </div>

        )
    }
}
