import React, { Component } from 'react';
import moment from 'moment';

class Nav extends Component {
    render() {
        return (
            <div>
                <header>
                    <h1>Taskpad</h1>
                </header>
                <nav>
                    <ul className="site-nav">
                        <li>
                            <a href="/">Tasks</a>
                        </li>
                        <li>
                            <a href="/reports">Reports</a>
                        </li>
                        <li className="nav-date">{moment().format('DD MMM YYYY')}</li>
                        <li className="nav-username">{'administrator'}</li>
                    </ul>
                </nav>
            </div>
        );
    }
}

export default Nav;