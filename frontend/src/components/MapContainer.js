import React from 'react';
import UserInput from './UserInput.js';
import MyMapComponent from './Map.js';

const route = [
    { lat: 51.1431268, lng: 23.4712059 },
    { lat: 51.1434216, lng: 23.4712741 },
    { lat: 51.1434216, lng: 23.4712741 },
    { lat: 51.1464649, lng: 23.4370398 },
    { lat: 51.1464649, lng: 23.4370398 },
    { lat: 51.1480887, lng: 22.8609503 },
    { lat: 51.1480887, lng: 22.8609503 },
    { lat: 51.2889916, lng: 22.4611911 },
    { lat: 51.2889916, lng: 22.4611911 },
    { lat: 51.2869669, lng: 22.4568567 },
    { lat: 51.2869669, lng: 22.4568567 },
    { lat: 51.2882191, lng: 22.4419692 },
    { lat: 51.2882191, lng: 22.4419692 },
    { lat: 51.4336224, lng: 22.1332215 },
    { lat: 51.4336224, lng: 22.1332215 },
    { lat: 51.43433510000001, lng: 22.1199593 },
    { lat: 51.43433510000001, lng: 22.1199593 },
    { lat: 51.4444916, lng: 22.0149385 },
    { lat: 51.4444916, lng: 22.0149385 },
    { lat: 51.4376349, lng: 21.950906 },
    { lat: 51.4376349, lng: 21.950906 },
    { lat: 51.438193, lng: 21.9483589 },
    { lat: 51.438193, lng: 21.9483589 },
    { lat: 51.5606812, lng: 21.8452964 },
    { lat: 51.5606812, lng: 21.8452964 },
    { lat: 51.56098, lng: 21.8445582 },
    { lat: 51.56098, lng: 21.8445582 },
    { lat: 51.5636601, lng: 21.8344754 },
    { lat: 51.5636601, lng: 21.8344754 },
    { lat: 52.1554853, lng: 21.1824691 },
    { lat: 52.1554853, lng: 21.1824691 },
    { lat: 52.1728445, lng: 21.1543993 },
    { lat: 52.1728445, lng: 21.1543993 },
    { lat: 52.2370126, lng: 21.044299 },
    { lat: 52.2370126, lng: 21.044299 },
    { lat: 52.2371881, lng: 21.0468107 },
    { lat: 52.2371881, lng: 21.0468107 },
    { lat: 52.2319724, lng: 21.0210484 },
    { lat: 52.2319724, lng: 21.0210484 },
    { lat: 52.2301315, lng: 21.0119668 },
    { lat: 52.2301315, lng: 21.0119668 },
    { lat: 52.2278977, lng: 21.01273 },
    { lat: 52.2278977, lng: 21.01273 },
    { lat: 52.2284848, lng: 21.0156871 },
    { lat: 52.2284848, lng: 21.0156871 },
    { lat: 52.2291168, lng: 21.015462 }
];

export default class MapContainer extends React.Component {

      componentWillMount(){
        this.setState({pathCoordinates : []});
        this.submit = this.submit.bind(this);
      }

      componentDidMount() {
        // What should happen when container is loaded?
      }

    submit(data) {
        console.log(data);
        fetch('/api/form-submit-url', {
            method: 'POST',
            body: data,
        });
        // Here we set response as path
        this.setState({ pathCoordinates : route });
    }

    render() {
        return <div>
            <UserInput submit={this.submit}/>
            <MyMapComponent pathCoordinates={this.state.pathCoordinates}/>
        </div>;
    }
}
