import React from 'react';
import UserInput from './UserInput.js';
import RouteInfo from './RouteInfo.js';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import MyMapComponent from './Map.js';
import Steps from './Steps.js';
import Waypoints from './Waypoints.js';
import moment from 'moment';

export default class MapContainer extends React.Component {

    componentWillMount() {
        this.setState({
            places: [],
            routeInfo: {
                Route: [],
            },
            waypoints: [],
        });
        this.places = []
        this.submit = this.submit.bind(this);
        this.setWaypoints = this.setWaypoints.bind(this);
        this.setMarkerShown = this.setMarkerShown.bind(this);
        this.setMarkerPosition = this.setMarkerPosition.bind(this);
    }

    componentDidMount() {
        // What should happen when container is loaded?
        fetch('/api/places')
        .then((response) => {
            return response.json()
        })
        .then(responseData => {
            console.log(responseData);
            this.setState({ places: responseData });
            // this.places = responseData
        });
        
    }

    submit(data) {
        console.log(data);
        this.setState({
            routeInfo: {
                Route: [],
            },
        });
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
            responseData.OverviewPolyline = window.google.maps.geometry.encoding.decodePath(responseData.OverviewPolyline.points)
            this.setState({ routeInfo: responseData });
        });
    }

    setWaypoints(waypoints) {
        this.setState({ waypoints: waypoints });
    }

    setMarkerShown(isMarkerShown) {
        this.setState({ isMarkerShown: isMarkerShown });
    }

    setMarkerPosition(markerPosition) {
        console.log(markerPosition)
        this.setState({ markerPosition: markerPosition });
    }

    render() {
        return <React.Fragment>
            {/* <script src="https://maps.googleapis.com/maps/api/js?key=YOUR_API_KEY&callback=initMap" async defer></script> */}

            <Row className="flex-grow-1" >
                <Col md={3}>
                    <UserInput submit={this.submit} waypoints={this.state.waypoints}/>
                    <Waypoints waypointsFunc={this.setWaypoints} places={this.state.places}/>
                </Col>
                <Col md={4}>
                    <RouteInfo routeInfo={this.state.routeInfo} />
                    <Steps routeInfo={this.state.routeInfo} markerShownFunc={this.setMarkerShown} markerPositionFunc={this.setMarkerPosition}/>
                </Col>
                <Col md={5} className="right-col">
                    <MyMapComponent overviewPolyline={this.state.routeInfo.OverviewPolyline}
                    isMarkerShown={this.state.isMarkerShown} markerPosition={this.state.markerPosition} />
                </Col>
            </Row>

        </React.Fragment>;
    }
}
