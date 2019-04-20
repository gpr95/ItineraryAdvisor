import React from "react";
import { compose, withProps } from "recompose";
import {
    withScriptjs,
    withGoogleMap,
    GoogleMap,
    // Marker,
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
    withGoogleMap
)(props => (
    <GoogleMap defaultZoom={8} defaultCenter={{ lat: 52.237, lng: 21.018 }}>
        <Polyline
            // path={props.pathCoordinates}
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
    </GoogleMap>
));

export default MyMapComponent;