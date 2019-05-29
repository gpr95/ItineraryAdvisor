import React, { Component } from 'react';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';

import 'react-bootstrap-typeahead/css/Typeahead.css';

const SUPPORTED_PLACES_CODES = ['amusement_park', 'aquarium', 'art_gallery', 'bar', 'beauty_salon', 'book_store', 'bowling_alley', 'cafe', 'casino', 'church', 'city_hall', 'clothing_store', 'gym', 'hair_care', 'laundry', 'meal_takeaway', 'movie_rental', 'movie_theater', 'museum', 'night_club', 'park', 'pharmacy', 'store', 'supermarket', 'travel_agency', 'zoo']

export default class WaypointChooser extends Component {

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
        return (
            <Col sm={4}>
                <input type="checkbox" key={name} id={name} name={name} value={name} onClick={handle} style={{ marginRight: '5px' }} />
                <label key={name + "_label"} htmlFor={name}>{name}</label>
            </Col>
        );
    }

    render() {

        return (
            <React.Fragment>
                <Row style={{ height: '200px', overflow: 'auto', whiteSpace: 'nowrap' }} className="fullWidth">
                    {SUPPORTED_PLACES_CODES.map(place_code => this.renderCheckbox(place_code))}
                </Row>
                <Row>
                    <Container>
                        <Button className="fullWidth" onClick={this.props.getWaypoints}>Fetch Waypoints for this area</Button>
                    </Container>
                </Row>
            </React.Fragment>
        );
    }
}