import React, { Component } from 'react';
import Container from 'react-bootstrap/Container';
import Table from 'react-bootstrap/Table';


export default class RouteInfo extends Component {

    constructor(props) {
        super(props);

        // this.state = {
        //     routeInfo: {},
        // };
    }

    render() {
        return (
            <React.Fragment>
                <Container>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th colSpan="2">Route Info</th>
                            </tr>
                        </thead>
                        <tbody>
                            {/* {this.props.routeInfo.map((route) => (
                                <tr>
                                <td>{route}</td>
                                <td>{route}</td>    
                                </tr>
                            ))} */}
                            <tr>
                                <td>Distance</td>
                                <td>{this.props.routeInfo.Distance}</td>
                            </tr>
                            <tr>
                                <td>Duration</td>
                                <td>{this.props.routeInfo.Duration}</td>
                            </tr>
                            <tr>
                                <td>DepartureTime</td>
                                <td>{this.props.routeInfo.DepartureTime}</td>
                            </tr>
                            <tr>
                                <td>ArrivalTime</td>
                                <td>{this.props.routeInfo.ArrivalTime}</td>
                            </tr>
                        </tbody>
                        
                    </Table>
                </Container>
            </React.Fragment>
        );
    }
}