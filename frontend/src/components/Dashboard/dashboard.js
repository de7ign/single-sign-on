import React, { Component } from 'react'
import './dashboard.css'
import axios from 'axios'
axios.defaults.withCredentials = true
export default class Dashboard extends Component {
    constructor(props){
        super(props);
        this.state = {
            name    :   'Loading name',
            avatar :   'Loading avatar',
            email   :   'Loading email'
        }

        axios.post('http://localhost:5000/v1/api/userinfo')
        .then((res)=>{
            console.log(res.data)
            this.setState({
                name    :   res.data.Name,
                email   :   res.data.Email,
                avatar  :   res.data.AvatarURL
            })
        })
        .catch((err)=>{
            console.log(err)
        })
    }



    render() {
        return(
            <div class="dashboard">
                <h2>User data</h2>
                <hr />
                <img src={this.state.avatar} alt="avatar" height="300" width="300"/>
                <h4>{this.state.name}</h4>
                <h4>{this.state.email}</h4>
            </div>
        )
    }
}