import _ from "lodash";
import React from "react";
import { compose, withProps } from "recompose";
import {
    withScriptjs,
    withGoogleMap,
    GoogleMap,
    Marker,
    Polyline
} from "react-google-maps";

import UserInput from './components/UserInput.js';

const pathCoordinates = [
    { lat: 51.1431268,  lng: 23.4712059},
    { lat: 51.1434216,  lng: 23.4712741},
    { lat: 51.1434216,  lng: 23.4712741},
    { lat: 51.1464649,  lng: 23.4370398},
    { lat: 51.1464649,  lng: 23.4370398},
    { lat: 51.1480887,  lng: 22.8609503},
    { lat: 51.1480887,  lng: 22.8609503},
    { lat: 51.2889916,  lng: 22.4611911},
    { lat: 51.2889916,  lng: 22.4611911},
    { lat: 51.2869669,  lng: 22.4568567},
    { lat: 51.2869669,  lng: 22.4568567},
    { lat: 51.2882191,  lng: 22.4419692},
    { lat: 51.2882191,  lng: 22.4419692},
    { lat: 51.4336224,  lng: 22.1332215},
    { lat: 51.4336224,  lng: 22.1332215},
    { lat: 51.43433510000001,  lng: 22.1199593},
    { lat: 51.43433510000001,  lng: 22.1199593},
    { lat: 51.4444916,  lng: 22.0149385},
    { lat: 51.4444916,  lng: 22.0149385},
    { lat: 51.4376349,  lng: 21.950906},
    { lat: 51.4376349,  lng: 21.950906},
    { lat: 51.438193,  lng: 21.9483589},
    { lat: 51.438193,  lng: 21.9483589},
    { lat: 51.5606812,  lng: 21.8452964},
    { lat: 51.5606812,  lng: 21.8452964},
    { lat: 51.56098,  lng: 21.8445582},
    { lat: 51.56098,  lng: 21.8445582},
    { lat: 51.5636601,  lng: 21.8344754},
    { lat: 51.5636601,  lng: 21.8344754},
    { lat: 52.1554853,  lng: 21.1824691},
    { lat: 52.1554853,  lng: 21.1824691},
    { lat: 52.1728445,  lng: 21.1543993},
    { lat: 52.1728445,  lng: 21.1543993},
    { lat: 52.2370126,  lng: 21.044299},
    { lat: 52.2370126,  lng: 21.044299},
    { lat: 52.2371881,  lng: 21.0468107},
    { lat: 52.2371881,  lng: 21.0468107},
    { lat: 52.2319724,  lng: 21.0210484},
    { lat: 52.2319724,  lng: 21.0210484},
    { lat: 52.2301315,  lng: 21.0119668},
    { lat: 52.2301315,  lng: 21.0119668},
    { lat: 52.2278977,  lng: 21.01273},
    { lat: 52.2278977,  lng: 21.01273},
    { lat: 52.2284848,  lng: 21.0156871},
    { lat: 52.2284848,  lng: 21.0156871},
    { lat: 52.2291168,  lng: 21.015462}
    
];

const MyMapComponent = compose(
    withProps({
        googleMapURL:
            "https://maps.googleapis.com/maps/api/js?key=AIzaSyALc9COvc2U9XmwYrCGrAVGBulwdZ5yUcE&v=3.exp&libraries=geometry,drawing,places",
        loadingElement: <div style={{ height: `100%` }} />,
        containerElement: <div style={{ height: `800px` }} />,
        mapElement: <div style={{ height: `100%` }} />
    }),
    withScriptjs,
    withGoogleMap
)(props => (
    <GoogleMap defaultZoom={8} defaultCenter={{ lat: 52.237, lng: 21.018 }}>
<Polyline
                path={pathCoordinates}
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

const enhance = _.identity;

const ReactGoogleMaps = () => [
    <UserInput />,
    <MyMapComponent key="map" />
];

export default enhance(ReactGoogleMaps);