import React from 'react';
import './App.css';
import Content from './module/content'
import Header from './module/header'

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  render() {
    return (
      <div className="App">
        <Header/>
        <Content/>
        <footer className="App-footer">
          &copy;by Group 8
        </footer>
      </div>
    );
  }
}

export default App;
