import React, { Component } from 'react';
import Container from 'react-bootstrap/Container';
import Table from 'react-bootstrap/Table';
import BootstrapTable from 'react-bootstrap-table-next';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';


export default class Steps extends Component {

    constructor(props) {
        super(props);

        this.rowEvents = {
            onMouseEnter: (e, row, rowIndex) => {
                this.props.markerShownFunc(true)
                this.props.markerPositionFunc(row.Location)
            },
            onMouseLeave: (e, row, rowIndex) => {
                this.props.markerShownFunc(false)
            }
        };
        this.columns = [{
            dataField: 'Instruction',
            text: 'Instruction',
            formatter: this.instructionFormatter
        }, {
            dataField: 'Mode',
            text: 'Mode',
        }];
    }

    instructionFormatter(cell, row) {
        console.log(cell)
        return (
            <div dangerouslySetInnerHTML={{ __html: cell }} />
        );
    }

    instructionFormatter(cell, row) {
        console.log(cell)
        return (
            <div dangerouslySetInnerHTML={{ __html: cell }} />
        );
    }

    render() {
        return (
            <React.Fragment>
                {/* <Container style={{ height: '400px', overflow: 'auto' }}> */}
                <Container>
                    <BootstrapTable bootstrap4 keyField='Instruction' data={this.props.routeInfo.Route} columns={this.columns} rowEvents={this.rowEvents}>
                    </BootstrapTable>
                </Container>
            </React.Fragment>
        );
    }
}