import React from 'react';
import './search.css';
import { SearchBar } from 'antd-mobile';
  
class SearchBarExample extends React.Component {
    state = {
        
    };
    componentDidMount() {
    }
    onChange = (value) => {
        this.setState({ value });
    };
    render() {
        return (<div>
            <SearchBar
                value={this.state.value}
                placeholder="Search"
                onSubmit={value => console.log(value, 'onSubmit')}
                onClear={value => console.log(value, 'onClear')}
                onFocus={() => console.log('onFocus')}
                onBlur={() => console.log('onBlur')}
                onCancel={() => console.log('onCancel')}
                showCancelButton
                onChange={this.onChange}
            />
        </div>);
    }
}

export default SearchBarExample;