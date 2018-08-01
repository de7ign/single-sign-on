import React, { Component } from 'react'
import { Redirect } from 'react-router-dom'
import './dashboard.css'
import axios from 'axios'
axios.defaults.withCredentials = true
export default class Dashboard extends Component {
    constructor(props){
        super(props);
        this.state = {
            name    :   'Loading name',
            avatar :   'Loading avatar',
            email   :   'Loading email',
            authorized  :   true
        }

        axios.post('http://localhost:5000/v1/api/userinfo')
        .then((res)=>{
            this.setState({
                name    :   res.data.Name,
                email   :   res.data.Email,
                avatar  :   res.data.AvatarURL
            })
        })
        .catch((err)=>{
            this.setState({
                authorized  :   false
            })
        })
    }

    handleLogout = (event) => {
        axios.post('http://localhost:5000/v1/api/logout')
        .then((res)=>{
            this.setState({
                authorized  :   false
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
        if (!authorized) {
            return (
                <Redirect to="/" />
            )
        }
        return(
            <div class="dashboard">
                <h2>User data</h2>
                <hr />
                <img src={this.state.avatar} alt="avatar" height="300" width="300"/>
                <h4>{this.state.name}</h4>
                <h4>{this.state.email}</h4>
                <br /><br /><br /><br />
                <input type="button" value="Logout" class="button-primary button" onClick={this.handleLogout} />
            </div>
        )
    }
}