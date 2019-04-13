import React, { Component } from 'react';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import TimePicker from 'rc-time-picker';
import 'rc-time-picker/assets/index.css';

export default class UserInput extends Component {

    constructor(props) {
        super(props);
        this.state = {
            username: '',
            password: ''
        };
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange = ({ target }) => {
        this.setState({ [target.name]: target.value });
    };

    handleSubmit(event) {
        event.preventDefault();
        const data = new FormData(event.target);
        console.log(data)

        console.log(this.refs.departure)

        data.set('departure', this.refs.departure.state.value)
        data.set('arrival', this.refs.arrival.state.value)

        this.props.submit(data);
    }

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
                                    <Button type="submit">Find!</Button>
                                {/* </Form.Group> */}
                            </Col>
                        </Row>

                    </Form>

                    {/* <h3>Your username is: {this.state.username}:{this.state.password}</h3> */}
                </Container>
            </React.Fragment>
        );
    }
}