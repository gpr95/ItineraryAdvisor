import React, { Component } from 'react';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';

import TimePicker from 'rc-time-picker';
import 'rc-time-picker/assets/index.css';

import "./style.css";

export default class UserInput extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: '',
            waypoints: [{name:""}],
        };
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange = ({ target }) => {
        this.setState({ [target.name]: target.value });
    };

    handleSubmit(event) {
        event.preventDefault();
        const data = new FormData(event.target);
        console.log(data);

        console.log(this.refs.departure);

        data.set('departure', this.refs.departure.state.value);
        data.set('arrival', this.refs.arrival.state.value);

        let wayPointsParsed;
        wayPointsParsed = '';
        for (let waypoint of this.state.waypoints) {
            wayPointsParsed = wayPointsParsed.concat(waypoint.name);
            wayPointsParsed = wayPointsParsed.concat('|')
        }
        if(wayPointsParsed.length !== 0)
            wayPointsParsed = wayPointsParsed.substring(0, wayPointsParsed.length - 1);

        data.set('waypoints', wayPointsParsed);

        this.props.submit(data);
    }

    handleWaypointNameChange = idx => evt => {
        const newWaypoints = this.state.waypoints.map((waypoint, sidx) => {
            if (idx !== sidx) return waypoint;
            return { ...waypoint, name: evt.target.value };
        });

        this.setState({ waypoints: newWaypoints });
    };

    handleAddWaypoint = () => {
        this.setState({
            waypoints: this.state.waypoints.concat([{ name: "" }])
        });
    };

    handleRemoveWaypoint = idx => () => {
        this.setState({
            waypoints: this.state.waypoints.filter((s, sidx) => idx !== sidx)
        });
    };


    render() {
        return (

            <React.Fragment>
                <Container>

                    <link
                        rel="stylesheet"
                        href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
                        integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
                        crossOrigin="anonymous"
                    />

                    <script src="https://unpkg.com/react/umd/react.production.js" crossOrigin="anonymous" />

                    <script
                        src="https://unpkg.com/react-dom/umd/react-dom.production.js"
                        crossOrigin="anonymous"
                    />

                    <script
                        src="https://unpkg.com/react-bootstrap@next/dist/react-bootstrap.min.js"
                        crossOrigin="anonymous"
                    />

                    <script>var Alert = ReactBootstrap.Alert;</script>

                    <Form onSubmit={this.handleSubmit}>
                    <Row>
                    <Col md={8}>
                        <Row>
                            <Col>
                                <Form.Label>Origin</Form.Label>
                                <Form.Control defaultValue="Chełm"
                                    id="origin"
                                    name="origin"
                                />
                            </Col>
                            <Col>
                                <Form.Label>Destination</Form.Label>
                                <Form.Control defaultValue="Warszawa"
                                    id="destination"
                                    name="destination" />
                            </Col>

                        </Row>
                        <Row>
                            <Col>
                                <Form.Group controlId="exampleForm.ControlSelect1">
                                    <Form.Label>Lookup mode</Form.Label>
                                    <Form.Control as="select"
                                        name="mode">
                                        <option>driving</option>
                                        <option>walking</option>
                                        <option>bicycling</option>
                                        <option>transit</option>
                                        <option>ignore</option>
                                    </Form.Control>
                                </Form.Group>
                            </Col>
                        </Row>

                        <Row>
                        <Col>

                            </Col>
                        </Row>

                        <Row>
                            <Form.Group>
                                <Col>

                                    <Form.Label>Departure</Form.Label>
                                </Col><Col>
                                    <TimePicker ref="departure" />

                                </Col>
                            </Form.Group>
                            <Form.Group>
                                <Col>

                                    <Form.Label>Arrival</Form.Label>
                                </Col><Col>
                                    <TimePicker ref="arrival" />

                                </Col>
                            </Form.Group>
                            {/* </Row>
                        <Row> */}
                            <Col>
                                {/* <Form.Group> */}
                                    <Button className="find" type="submit">Find!</Button>
                                {/* </Form.Group> */}
                            </Col>
                        </Row>
                        </Col>
                        <Col md={4}>
                        <Row>
                        <Form.Label>Waypoints</Form.Label>
                        {this.state.waypoints.map((waypoint, idx) => (
                                <InputGroup className="mb-3 waypoint" >
                                    <Form.Control
                                        type="text"
                                        placeholder={`Waypoint #${idx + 1} name`}
                                        value={waypoint.name}
                                        onChange={this.handleWaypointNameChange(idx)}
                                    />
                                    <InputGroup.Append>
                                    <Button 
                                        type="button"
                                        onClick={this.handleRemoveWaypoint(idx)}
                                        className="small"
                                    >
                                        -
                                    </Button>
                                    </InputGroup.Append>
                                </InputGroup>
                            ))}
                            </Row>
                            <Row>
                            <Button
                                type="button"
                                onClick={this.handleAddWaypoint}
                                className="small"
                                size="sm"
                            >
                                Add Waypoint
                            </Button>
                            </Row>
                        </Col>
                        </Row>
                    </Form>

                    {/* <h3>Your username is: {this.state.username}:{this.state.password}</h3> */}
                </Container>
            </React.Fragment>
        );
    }
}