import React, { Component } from 'react';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import TimePicker from 'rc-time-picker';

import 'rc-time-picker/assets/index.css';
import 'bootstrap/dist/css/bootstrap.min.css'

import 'react/umd/react.production.min.js'
// Should be 'umd' but compilation fails
// import 'react-dom/umd/react-dom.production.js'
import 'react-dom/cjs/react-dom.production.min.js'
import 'react-bootstrap/dist/react-bootstrap.min.js'

import "./style.css";

let TRANSPORT_TYPE = ['driving', 'walking', 'bicycling', 'transit']

export default class UserInput extends Component {
    constructor(props) {
        super(props);
        this.state = {
            checkedModes: [],
        };
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleCheckboxChange = this.handleCheckboxChange.bind(this);
    }

    handleSubmit(event) {
        event.preventDefault();
        const data = new FormData(event.target);

        data.set('arrival', this.refs.arrival.state.value);
        data.set('departure', this.refs.departure.state.value);
        // data.set('waypoints', this.props.waypoints.map((w) => w.name).join('|'));
        // data.set('waypoints-time', this.props.waypoints.map((w) => w.time).join('|'));
        data.set('waypoints', JSON.stringify([...this.props.waypoints]));
        data.set('lookup-mode', JSON.stringify([...this.state.checkedModes]));
        this.props.submit(data);
    }

    handleCheckboxChange(e) {
        const id = e.target.id;
        const isChecked = e.target.checked;
        this.setState(prevState => ({ 
            checkedModes: [...prevState.checkedModes.filter(obj => {return obj.name !== id}), { name: id, used: isChecked }],
        }));
    }

    render() {
        return (

            <React.Fragment>
                <Container>
                    <Form onSubmit={this.handleSubmit}>
                        <Row>
                            <Form.Label>Startin point</Form.Label>
                            <Form.Control defaultValue="Pałac kultury i Nauki"
                                id="origin"
                                name="origin"
                            />
                        </Row>
                        <br/>
                        <Row>
                            <Col>
                                <Form.Group>
                                    <Form.Label>Departure</Form.Label>{'\u00A0'}
                                    <TimePicker ref="departure" />
                                </Form.Group>
                            </Col><Col>
                                <Form.Group>
                                    <Form.Label>Arrival</Form.Label>{'\u00A0'}
                                    <TimePicker ref="arrival" />
                                </Form.Group>
                            </Col>
                        </Row>
                        <Row>
                            <Form.Group>
                                <Form.Label>Lookup mode</Form.Label>
                                <Col>
                                    {TRANSPORT_TYPE.map(type => (
                                        <Form.Check type='checkbox'
                                            id={`${type}`}
                                            label={type.charAt(0).toUpperCase() + type.slice(1)}
                                            checked={
                                                this.state.checkedModes.filter(obj => {return obj.name === type})[0] ? 
                                                this.state.checkedModes.filter(obj => {return obj.name === type})[0].used :
                                                undefined
                                            } 
                                            onChange={this.handleCheckboxChange} />
                                    ))}
                                </Col>
                            </Form.Group>
                        </Row>
                        <Row>
                            <Button className="find" type="submit">Find!</Button>
                        </Row>
                    </Form>
                </Container>
            </React.Fragment>
        );
    }
}