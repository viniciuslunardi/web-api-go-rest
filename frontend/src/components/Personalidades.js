import React, { Component } from 'react';
import axios from 'axios';
import '../components/Personalidades.css';

export default class Personalidades extends Component {
    state = {
        personalidades: []
    }

    componentDidMount() {
        axios.get('http://localhost:8000/api/people')
            .then(res => {
                const personalidades = res.data;
                this.setState({ personalidades })
            })
    }

    render() {
        return (
            <div>
                {this.state.personalidades.map((p, id )=>
                    <div className="CardPersonalidades" key={id}>
                        <h3>{p.name}</h3>
                        <p>{p.history}</p>
                    </div>)}
            </div>
        );
    }
}