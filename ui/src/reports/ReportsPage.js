import React, { Component } from 'react';
import Toolbar from './Toolbar';
import httpApi from '../utils/http-api';
import ReportCreateView from './report-create/ReportCreateView';
import ReportsView from './report-display/ReportsView';

class ReportsPage extends Component {
    state = {
        showReportCreate: false,
        loading: false,
        report: []
    };

    componentWillMount() {
        this.loadReports();
    }

    loadReports = async () => {
        this.setState({
            loading: true
        });        

        httpApi.getWithErrorHandled(`/api/reports`).then(reports => {                  
            this.setState({                
                reports: reports,
                loading: false
            })            
        })
    }

    render() {                  
        return (
            <div className="column-main tile">
                <Toolbar 
                    count={1}
                    loading={this.state.loading}
                />

                {this.state.showReportCreate ? (
                    <ReportCreateView />
                ) : null }

                {this.state.loading ? (
                    <ReportsView 
                        reports={this.state.reports}                        
                    />
                ) 
                : null }
            </div>
        );
    }
}

export default ReportsPage;