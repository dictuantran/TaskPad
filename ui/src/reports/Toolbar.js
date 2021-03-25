import React, { Component } from 'react';
import { BarLoader } from 'react-spinners';

class Toolbar extends Component {
    getReportsCount() {
        if (this.props.loading)
            return null;
        
        const { count } = this.props;
        if (count === 0) return 'No Reports (yet)';
        if (count === 1) return '1 report';

        return `${count} reports`;
    }

    render() {
        return(
            <div className="report-toolbar">
                <strong className="count">{this.getReportsCount()}</strong>
                <div className="sweet-loading loader">
                    <BarLoader
                        sizeUnit={'em'}
                        height={7}
                        width={200}
                        color={'#cc6b5a'}
                        loading={this.props.loading}
                    />                
                </div>
            </div>
        );        
    }
}

export default Toolbar;
