import React, { Component } from 'react';
import Toolbar from './Toolbar';

class ReportsPage extends Component {
    state = {
        showReportCreate: false,
        loading: false,
        report: []
    };

    render() {
        return (
            <div className="column-main tile">
                <Toolbar 
                    count={1}
                    loading={this.state.loading}
                />
            </div>
        );
    }
}

export default ReportsPage;