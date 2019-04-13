import React from 'react';
import UserInput from './UserInput.js';
import MyMapComponent from './Map.js';

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
        })
        .then((response) => {
            return response.json()
        })
        .then(responseData => {
            console.log(responseData);
            // Here we set response as path
            this.setState({ pathCoordinates : responseData });
        });
        
        ;
    }

    render() {
        return <div>
            <UserInput submit={this.submit}/>
            <MyMapComponent pathCoordinates={this.state.pathCoordinates}/>
        </div>;
    }
}
