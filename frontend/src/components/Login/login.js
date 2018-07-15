import React, { Component } from 'react';
import { Redirect } from 'react-router-dom'
import './login.css'
import axios from 'axios'
axios.defaults.withCredentials = true;
export default class Login extends Component {
    constructor(props){
        super(props);
        this.state = {
            authorized  :   false
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
            <div class="login">
                <a class="button-primary button" href="http://localhost:5000/v1/api/auth/google">Login with google</a>
                <a class="button-primary button" href="http://localhost:5000/v1/api/auth/github">Login with github</a>
            </div>
        )
    }
}
