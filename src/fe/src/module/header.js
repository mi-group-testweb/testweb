import React from 'react';
import './module.css'
import logo from '../logo.png'

class Header extends React.Component {
    render() {
        return (
            <div className="header-c">
                <div className="log-c">
                    <img src={logo} className="App-logo" alt="logo" />
                    <span className="App-name">Pmonitor</span>
                </div>
                <span className="explane-c">
                    <span className="explane">
                        网站说明
                    </span>
                    <span className="explane">
                        联系我们
                    </span>
                    <span className="explane">
                        意见反馈
                    </span>
                </span>
            </div>
        )
    }
}
export default Header;
