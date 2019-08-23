import React from 'react';
import './search.css';
import { SearchBar } from 'antd-mobile';
// import {ButtonGroup,Button,DropdownButton} from 'react-bootstrap/ButtonGroup'
// var Dropdown = require("react-bootstrap-dropdown");
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
            {/* <ButtonGroup vertical>
                <Button>Button</Button>
                <Button>Button</Button>
                <DropdownButton as={ButtonGroup} title="Dropdown" id="bg-vertical-dropdown-1">
                    <Dropdown.Item eventKey="1">Dropdown link</Dropdown.Item>
                    <Dropdown.Item eventKey="2">Dropdown link</Dropdown.Item>
                </DropdownButton>
                <Button>Button</Button>
                <Button>Button</Button>
                <DropdownButton as={ButtonGroup} title="Dropdown" id="bg-vertical-dropdown-2">
                    <Dropdown.Item eventKey="1">Dropdown link</Dropdown.Item>
                    <Dropdown.Item eventKey="2">Dropdown link</Dropdown.Item>
                </DropdownButton>
                <DropdownButton as={ButtonGroup} title="Dropdown" id="bg-vertical-dropdown-3">
                    <Dropdown.Item eventKey="1">Dropdown link</Dropdown.Item>
                    <Dropdown.Item eventKey="2">Dropdown link</Dropdown.Item>
                </DropdownButton>
            </ButtonGroup> */}
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