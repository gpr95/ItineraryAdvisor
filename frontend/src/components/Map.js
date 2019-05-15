import React from "react";
import { compose, withProps, lifecycle } from "recompose";
import {
    withScriptjs,
    withGoogleMap,
    GoogleMap,
    Marker,
    Polyline
} from "react-google-maps";
import * as config from '../config/config.secret.json'

const MyMapComponent = compose(
    withProps({
        googleMapURL:
            "https://maps.googleapis.com/maps/api/js?key=" + config.GoogleAPIKey + "&v=3.exp&libraries=geometry,drawing,places",
        loadingElement: <div style={{ height: '100%' }} />,
        containerElement: <div style={{ height: '100%' }} />,
        mapElement: <div style={{ height: '100%' }} />,
    }),
    withScriptjs,
    withGoogleMap,
    lifecycle({
        componentWillMount() {
            const refs = {}

            this.setState({
                bounds: null,
                center: {
                    lat: 52.237, lng: 21.018
                },
                markers: [],
                onMapMounted: ref => {
                    refs.map = ref;
                },
                onBoundsChanged: () => {
                    this.props.setCurrentViewFunc(refs.map.getBounds())
                },
            })
        },
    }),
)(props => (
    <GoogleMap defaultZoom={14} ref={props.onMapMounted} center={props.center}
        onBoundsChanged={props.onBoundsChanged}>
        <Polyline
            path={props.overviewPolyline}
            geodesic={true}
            options={{
                strokeColor: "#ff2527",
                strokeOpacity: 0.75,
                strokeWeight: 2,
                icons: [
                    {
                        offset: "0",
                        repeat: "20px"
                    }
                ]
            }}
        />
        {props.isMarkerShown && <Marker position={props.markerPosition} />}
    </GoogleMap>
));

export default MyMapComponent;