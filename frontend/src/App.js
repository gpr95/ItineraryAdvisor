import _ from "lodash";
import React from "react";

import MapContainer from './components/MapContainer.js';

const enhance = _.identity;

const ReactGoogleMaps = () => [
    <MapContainer />,
];

export default enhance(ReactGoogleMaps);