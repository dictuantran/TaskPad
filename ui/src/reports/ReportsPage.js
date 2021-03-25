import React, { Component } from 'react';
import Toolbar from './Toolbar';
import httpApi from '../utils/http-api';
import ReportCreateView from './report-create/ReportCreateView';
import ReportView from './report-display/ReportView';

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
            console.log('api test')     
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
                    <ReportCreateView/>
                ) : null }

                {this.state.loading ? (
                    <ReportView 
                        reports={this.state.reports}                        
                    />
                ) 
                : null }
            </div>
        );
    }
}

export default ReportsPage;