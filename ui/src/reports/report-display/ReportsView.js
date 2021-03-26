import React, { Component } from 'react';
import ReportView from './ReportView';

class ReportsView extends Component {
    renderReports() {    
        if (this.props.reports != undefined) {
            return this.props.reports.map(report => {
                return (
                    <ReportView 
                        key={report.id}
                        report={report}                    
                    />
                );
            })
        }

        return "";
    }

    render() {
        return <div className="reports-view">{this.renderReports()}</div>;
    }
}

export default ReportsView;