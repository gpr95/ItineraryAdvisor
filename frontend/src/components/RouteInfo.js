import React, { Component } from 'react';
import Container from 'react-bootstrap/Container';
import Table from 'react-bootstrap/Table';


export default class RouteInfo extends Component {

    render() {
        let tableContent;
        if (this.props.routeInfo.Distance !== undefined) {
            tableContent = <tbody>
                <tr>
                    <td>Distance</td>
                    <td>{this.props.routeInfo.Distance}</td>
                </tr>
                <tr>
                    <td>Duration</td>
                    <td>{this.props.routeInfo.Duration}</td>
                </tr>
                <tr>
                    <td>ArrivalTime</td>
                    <td>{this.props.routeInfo.ArrivalTime}</td>
                </tr>
            </tbody>
        }
        return (
            <React.Fragment>
                <Container>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th colSpan="2">Route Info</th>
                            </tr>
                        </thead>
                        {tableContent}

                    </Table>
                </Container>
            </React.Fragment>
        );
    }
}