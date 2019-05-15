import React, { Component } from 'react';
import Container from 'react-bootstrap/Container';
import Table from 'react-bootstrap/Table';


export default class RouteInfo extends Component {

    render() {
        let distance;
        let duration;
        let arrival;

        if (this.props.routeInfo.Distance !== undefined && this.props.routeInfo.Distance !== "") {
            distance = <tr>
                    <td>Distance</td>
                    <td>{this.props.routeInfo.Distance}</td>
                </tr>
        }
        if (this.props.routeInfo.Duration !== undefined && this.props.routeInfo.Duration !== "") {
                duration = <tr>
                    <td>Duration</td>
                    <td>{this.props.routeInfo.Duration}</td>
                </tr>
        }
        if (this.props.routeInfo.ArrivalTime !== undefined && this.props.routeInfo.ArrivalTime !== "") {
                arrival = <tr>
                    <td>ArrivalTime</td>
                    <td>{this.props.routeInfo.ArrivalTime}</td>
                </tr>
        }
        
        return (
            <React.Fragment>
                <Container>
                    <Table bordered hover>
                        <thead>
                            <tr>
                                <th colSpan="2">Route Info</th>
                            </tr>
                        </thead>
                        <tbody>
                            {distance}
                            {duration}
                            {arrival}
                        </tbody>
                    </Table>
                </Container>
            </React.Fragment>
        );
    }
}