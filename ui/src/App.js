import React, { Component } from 'react';
import { Route, Redirect } from 'react-router-dom';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import Home from './home/Home';
import Nav from './header/Nav';
import ReportsPage from './reports/ReportsPage';

class App extends Component {
    constructor(props) {
        super(props);        
    }    

    render() {        
        return (
            <div className="container">
                <ToastContainer />
                <Nav />
                <div className="body">
                    <Route 
                        path="/"
                        exact
                        render={props => <Home {...props} />}
                    />
                   
                    <Route
                        path="/reports"
                        render={props =>
                            <ReportsPage auth={this.auth} {...props} />
                        }
                    />
                </div>
            </div>
        );
    }
}

export default App;