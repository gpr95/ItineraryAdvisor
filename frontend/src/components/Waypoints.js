import React, { Component } from 'react';
import Table from 'react-bootstrap/Table';
import Row from 'react-bootstrap/Row';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';
import Tooltip from 'react-bootstrap/Tooltip';
import Container from 'react-bootstrap/Container';
import { Typeahead } from 'react-bootstrap-typeahead';

import 'react-bootstrap-typeahead/css/Typeahead.css';

const SUPPORTED_PLACES_CODES = ['amusement_park', 'aquarium', 'art_gallery', 'bar', 'beauty_salon', 'book_store', 'bowling_alley', 'cafe', 'casino', 'church', 'city_hall', 'clothing_store', 'gym', 'hair_care', 'laundry', 'meal_takeaway', 'movie_rental', 'movie_theater', 'museum', 'night_club', 'park', 'pharmacy', 'store', 'supermarket', 'travel_agency', 'zoo']

export default class Waypoints extends Component {

    constructor(props) {
        super(props);
        this.state = {
            selectedWaypoint: [],
            newWaypointName: '',
            newWaypointTime: '',
            newWaypointValid: false,
            newWaypointNameValid: false,
            newWaypointTimeValid: false,
            waypoints: [],
        };
        this.validateNewWaypoint = this.validateNewWaypoint.bind(this);
    }

    handleAddWaypoint = () => {
        let newelement = { name: this.state.newWaypointName, openingHours: this.state.newWaypointOpeningHours, time: this.state.newWaypointTime }
        console.log(this.state)
        this.typeahead.getInstance().clear()
        this.setState(prevState => ({
            waypoints: [...prevState.waypoints, newelement],
            selectedWaypoint: [],
            newWaypointName: '',
            newWaypointTime: '',
            newWaypointOpeningHours: '',
            newWaypointValid: false,
            newWaypointNameValid: false,
            newWaypointTimeValid: false,
            newWaypointOpeningHoursValid: false,
        }));
        // setState is asynchronous - not updated immediately
        // https://stackoverflow.com/questions/38558200/react-setstate-not-updating-immediately
        this.props.waypointsFunc([...this.state.waypoints, newelement])
    };

    handleRemoveWaypoint = idx => () => {
        this.setState({
            waypoints: this.state.waypoints.filter((s, sidx) => idx !== sidx)
        }, function () {
            this.props.waypointsFunc(this.state.waypoints);
        }.bind(this));
    };

    validateNewWaypoint(fieldName, value) {
        let timeValid = this.state.newWaypointTimeValid
        let openingHoursValid = this.state.newWaypointOpeningHoursValid
        let nameValid = this.state.newWaypointNameValid
        let result = false
        switch (fieldName) {
            case 'newWaypointTime':
                timeValid = /^(\d+h)?[ ]?(\d+m)?$/.test(value);
                timeValid = timeValid && value.length > 1;
                result = timeValid
                this.setState({ newWaypointTimeValid: timeValid });
                break;
            case 'newWaypointOpeningHours':
                openingHoursValid = /^(\d{2}:\d{2})-(\d{2}:\d{2})$/.test(value);
                result = openingHoursValid
                this.setState({ newWaypointOpeningHoursValid: openingHoursValid });
                break;
            case 'newWaypointName':
                nameValid = value.length >= 2;
                result = nameValid
                this.setState({ newWaypointNameValid: nameValid });
                break;
            default:
                break;
        }
        this.setState({ newWaypointValid: timeValid && nameValid && openingHoursValid });
        return result
    }

    handleUserWaypointInput(e) {
        console.log(e)
        if (e.target === undefined) {

            let selectedWaypoint = this.props.places.find(o => { return o.Name === e[0] })
            let isNewWaypointValid = false
            if (selectedWaypoint === undefined) {
                this.setState({ newWaypointName: e });
                this.validateNewWaypoint('newWaypointName', e);
                return
            }
            isNewWaypointValid = this.validateNewWaypoint('newWaypointName', selectedWaypoint.Name);
            isNewWaypointValid = isNewWaypointValid && this.validateNewWaypoint('newWaypointTime', selectedWaypoint.Time);
            isNewWaypointValid = isNewWaypointValid && this.validateNewWaypoint('newWaypointOpeningHours', selectedWaypoint.OpeningHours);
            this.setState({
                newWaypointName: selectedWaypoint.Name,
                newWaypointTime: selectedWaypoint.Time,
                newWaypointOpeningHours: selectedWaypoint.OpeningHours,
                newWaypointValid: isNewWaypointValid,
            });
            return
        }
        const name = e.target.name;
        const value = e.target.value;
        this.setState({ [name]: value });
        this.validateNewWaypoint(name, value);
    }

    renderCheckbox(name) {
        let handle = (ev) => {
            let checked = this.props.selectedPlacesTypes.slice()
            if (ev.target.checked) {
                checked.push(name);
            } else {
                let index = checked.indexOf(name);
                checked.splice(index, 1);
            }
            this.props.placesFunc(checked)
        }
        return [
            <input type="checkbox" key={name} id={name} name={name} value={name} onClick={handle} />,
            <label key={name + "_label"} htmlFor={name}>{name}</label>,
            <br />
        ];
    }

    render() {

        return (
            <React.Fragment>
                <Row>
                    <div style={{ height: '200px', overflow: 'auto' }} className="fullWidth">
                        {SUPPORTED_PLACES_CODES.map(place_code => this.renderCheckbox(place_code))}
                    </div>
                </Row>
                <Row>
                    <Container>
                        <Button className="fullWidth" onClick={this.props.getWaypoints}>Fetch Waypoints for this area</Button>
                    </Container>
                </Row>
                <Row>
                    <Form.Label>Waypoints</Form.Label>
                    <InputGroup className="mb-3 waypoint" >
                        <Typeahead
                            id="typeahead"
                            key="typeahead"
                            name="newWaypointName"
                            ref={(typeahead) => this.typeahead = typeahead}
                            options={this.props.places.map(o => { return o.Name })}
                            style={{ flexGrow: '3' }}
                            placeholder='Waypoint name'
                            value={this.state.newWaypointName}
                            selected={this.state.selectedWaypoint}
                            onChange={(event) => this.handleUserWaypointInput(event)}
                            onInputChange={(event) => this.handleUserWaypointInput(event)}
                            isValid={this.state.newWaypointNameValid} />
                        <OverlayTrigger key="tooltipOpeningHours"
                            placement="top"
                            overlay={
                                <Tooltip id={'tooltip-top'}>
                                    Time format: <strong>10:00-15:30</strong>.
                                        </Tooltip>}>
                            <Form.Control
                                name="newWaypointOpeningHours"
                                ref="newWaypointOpeningHours"
                                key="newWaypointOpeningHours"
                                style={{ flexGrow: '2' }}
                                placeholder='opening hours name'
                                value={this.state.newWaypointOpeningHours}
                                onChange={(event) => this.handleUserWaypointInput(event)}
                                isValid={this.state.newWaypointOpeningHoursValid} />
                        </OverlayTrigger>
                        <OverlayTrigger key="topWaypointTime"
                            placement="top"
                            overlay={
                                <Tooltip id={'tooltip-top'}>
                                    Time format: <strong>3h 15m</strong>.
                                        </Tooltip>}>
                            <Form.Control
                                name="newWaypointTime"
                                ref="newWaypointTime"
                                key="newWaypointTime"
                                placeholder='time'
                                value={this.state.newWaypointTime}
                                onChange={(event) => this.handleUserWaypointInput(event)}
                                isValid={this.state.newWaypointTimeValid} />
                        </OverlayTrigger>
                        <InputGroup.Append>
                            <Button type="button"
                                key="addWaypointButton"
                                onClick={this.state.newWaypointValid ? this.handleAddWaypoint : null}
                                disabled={!this.state.newWaypointValid}>
                                Add waypoint
                                    </Button>
                        </InputGroup.Append>
                    </InputGroup>
                </Row>
                <Row style={{ height: '200px', overflow: 'auto' }}>
                    <Table>
                        <tbody>
                            {this.state.waypoints.map((waypoint, idx) => (
                                <tr>
                                    {/* <td>{idx}</td> */}
                                    <td>{waypoint.name}</td>
                                    <td>{waypoint.openingHours}</td>
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
            </React.Fragment>
        );
    }
}