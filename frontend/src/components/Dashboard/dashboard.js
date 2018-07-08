import React, { Component } from 'react'
import './dashboard.css'

export default class Dashboard extends Component {
    constructor(props){
        super(props);
        this.state = {
            name    :   'Nihal Murmu',
            picture :   'https://avatars0.githubusercontent.com/u/31739586?s=460&v=4',
            email   :   'nihal@mail.com'
        }
    }
    render() {
        return(
            <div class="dashboard">
                <h2>User data</h2>
                <hr />
                <img src={this.state.picture} alt="acatar" height="300" width="300"/>
                <h4>{this.state.name}</h4>
                <h4>{this.state.email}</h4>
            </div>
        )
    }
}