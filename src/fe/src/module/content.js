import React from 'react';
import './module.css'
import Search from '../components/search/search'
import Result from '../components/result/result';

class Content extends React.Component {
  render() {
    return (
      <div className="App-content">
        <div className="ctent"></div>
        <div className="ctent2">
          <Search/>
          <Result/>
        </div>
        <div className="ctent"></div>
      </div>
    )
  }
}
export default Content;
