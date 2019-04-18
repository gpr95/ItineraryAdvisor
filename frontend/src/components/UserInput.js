import React, { Component } from 'react';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';
import Tooltip from 'react-bootstrap/Tooltip';
import TimePicker from 'rc-time-picker';

import 'rc-time-picker/assets/index.css';
import 'bootstrap/dist/css/bootstrap.min.css'

import 'react/umd/react.production.min.js'
// Should be 'umd' but compilation fails
// import 'react-dom/umd/react-dom.production.js'
import 'react-dom/cjs/react-dom.production.min.js'
import 'react-bootstrap/dist/react-bootstrap.min.js'

import "./style.css";
import Table from 'react-bootstrap/Table';

let TRANSPORT_TYPE = ['driving', 'walking', 'bicycling', 'transit']

export default class UserInput extends Component {
    constructor(props) {
        super(props);
        this.state = {
            newWaypointName: '',
            newWaypointTime: '',
            newWaypointValid: false,
            newWaypointNameValid: false,
            newWaypointTimeValid: false,
            waypoints: [],
        };
        this.handleSubmit = this.handleSubmit.bind(this);
        this.validateNewWaypoint = this.validateNewWaypoint.bind(this);
    }

    handleSubmit(event) {
        event.preventDefault();
        const data = new FormData(event.target);

        data.set('arrival', this.refs.arrival.state.value);
        data.set('departure', this.refs.departure.state.value);
        data.set('waypoints', this.state.waypoints.map((w) => w.name).join('|'));
        data.set('waypoints-time', this.state.waypoints.map((w) => w.time).join('|'));
        this.props.submit(data);
    }

    handleAddWaypoint = () => {
        let newelement = { name: this.state.newWaypointName, time: this.state.newWaypointTime }

        this.setState(prevState => ({
            waypoints: [...prevState.waypoints, newelement]
        }));
    };

    handleRemoveWaypoint = idx => () => {
        this.setState({
            waypoints: this.state.waypoints.filter((s, sidx) => idx !== sidx)
        });
    };

    validateNewWaypoint(fieldName, value) {
        let timeValid = this.state.newWaypointTimeValid
        let nameValid = this.state.newWaypointNameValid

        switch (fieldName) {
            case 'newWaypointTime':
                timeValid = /^(\d+h)?[ ]?(\d+m)?$/.test(value);
                this.setState({ newWaypointTimeValid: timeValid });
                break;
            case 'newWaypointName':
                nameValid = value.length >= 2;
                this.setState({ newWaypointNameValid: nameValid });
                break;
            default:
                break;
        }
        this.setState({ newWaypointValid: timeValid && nameValid });
    }

    handleUserWaypointInput(e) {
        const name = e.target.name;
        const value = e.target.value;
        this.setState({ [name]: value });
        this.validateNewWaypoint(name, value);
    }

    render() {
        return (

            <React.Fragment>
                <Container>
                    <Form onSubmit={this.handleSubmit}>
                        <Row>
                            <Form.Label>Origin</Form.Label>
                            <Form.Control defaultValue="Chełm"
                                id="origin"
                                name="origin"
                            />
                        </Row>
                        <Row>
                            <Form.Label>Destination</Form.Label>
                            <Form.Control defaultValue="Warszawa"
                                id="destination"
                                name="destination" />
                        </Row>
                        <Row>
                            <Col>
                                <Form.Group>
                                    <Form.Label>Departure</Form.Label>
                                    <TimePicker ref="departure" />
                                </Form.Group>
                            </Col><Col>
                                <Form.Group>
                                    <Form.Label>Arrival</Form.Label>
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
                                            label={`${type}`} />
                                    ))}
                                </Col>
                            </Form.Group>
                        </Row>
                        <Row>
                            <Button className="find" type="submit">Find!</Button>
                        </Row>
                        <Row>
                            <Form.Label>Waypoints</Form.Label>
                            <InputGroup className="mb-3 waypoint" >
                                <Form.Control
                                    name="newWaypointName"
                                    ref="newWaypointName"
                                    style={{ flexGrow: '2' }}
                                    placeholder='Waypoint name'
                                    value={this.state.newWaypoint}
                                    onChange={(event) => this.handleUserWaypointInput(event)}
                                    isValid={this.state.newWaypointNameValid} />
                                <OverlayTrigger key="top"
                                    placement="top"
                                    overlay={
                                        <Tooltip id={'tooltip-top'}>
                                            Time format: <strong>3h 15m</strong>.
                                        </Tooltip>}>
                                    <Form.Control
                                        name="newWaypointTime"
                                        ref="newWaypointTime"
                                        placeholder='time'
                                        value={this.state.newWaypointTime}
                                        onChange={(event) => this.handleUserWaypointInput(event)}
                                        isValid={this.state.newWaypointTimeValid} />
                                </OverlayTrigger>
                                <InputGroup.Append>
                                    <Button type="button"
                                        onClick={this.state.newWaypointValid ? this.handleAddWaypoint : null}
                                        disabled={!this.state.newWaypointValid}>
                                        Add waypoint
                                    </Button>
                                </InputGroup.Append>
                            </InputGroup>
                        </Row>
                        <Row style={{height: '400px', overflow: 'auto'}}>
                            <Table>
                                <tbody>
                                    {this.state.waypoints.map((waypoint, idx) => (
                                        <tr>
                                            {/* <td>{idx}</td> */}
                                            <td>{waypoint.name}</td>
                                            <td>{waypoint.time}</td>
                                            <td>
                                                <Button type="button"
                                                    onClick={this.handleRemoveWaypoint(idx)}
                                                    size="sm"
                                                    className="remove"
                                                    variant="danger" >
                                                    Remove
                                                </Button>
                                            </td>
                                        </tr>
                                    ))}
                                </tbody>
                            </Table>
                        </Row>
                    </Form>
                </Container>
            </React.Fragment>
        );
    }
}