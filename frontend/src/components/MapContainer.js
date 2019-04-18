import React from 'react';
import UserInput from './UserInput.js';
import RouteInfo from './RouteInfo.js';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import MyMapComponent from './Map.js';
import moment from 'moment';

export default class MapContainer extends React.Component {

    componentWillMount() {
        this.setState({ routeInfo: {} });
        this.submit = this.submit.bind(this);
    }

    componentDidMount() {
        // What should happen when container is loaded?
    }

    submit(data) {
        console.log(data);
        fetch('/api/form-submit-url', {
            method: 'POST',
            body: data,
        })
            .then((response) => {
                return response.json()
            })
            .then(responseData => {
                console.log(responseData);
                // Here we set response as path
                if (responseData.DepartureTime === "0001-01-01T00:00:00Z")
                    responseData.DepartureTime = ""
                if (responseData.ArrivalTime === "0001-01-01T00:00:00Z")
                    responseData.ArrivalTime = ""
                responseData.Duration = moment.duration(responseData.Duration / 1000000).humanize()
                this.setState({ routeInfo: responseData });
            });

        ;
    }

    render() {
        return <React.Fragment>
            <Row className="flex-grow-1" >
                <Col md={3}>
                    <UserInput submit={this.submit} />
                </Col>
                <Col md={4}>
                    <RouteInfo routeInfo={this.state.routeInfo} />
                </Col>
                <Col md={5} className="right-col">
                    <MyMapComponent pathCoordinates={this.state.routeInfo.Route} />
                </Col>
            </Row>

        </React.Fragment>;
    }
}
