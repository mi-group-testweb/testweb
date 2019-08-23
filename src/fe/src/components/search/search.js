import React from 'react';
import './search.css';
import { Input, Button, Menu, Dropdown, Icon } from 'antd';

const { Search } = Input;

class SearchBarExample extends React.Component {
    state = {
        dataSource: [],
        selectValue: '单站'
    };
    handleMenuClick(e) {
        let self = this;
        if(e.key === '1'){
            self.setState({
                selectValue: '单站'
            });
        }else if(e.key === '2'){
            self.setState({
                selectValue: '对比'
            });
        }
    }

    render() {
        let self = this;
        let seltVal = self.state.selectValue;
        const menu = (
            <Menu onClick={self.handleMenuClick.bind(this)}>
                <Menu.Item key="1" value='单站'>单站</Menu.Item>
                <Menu.Item key="2" value='对比'>对比</Menu.Item>
            </Menu>
        );
        return (
            <div className='dropbtn'>
                <div className='sear-c'>
                    <div className='search1 search'>
                        <Search
                            className="heig-30"
                            placeholder="xiaomi.com"
                            size="large"
                            onSearch={value => console.log(value)}
                        />
                    </div>
                    <div className='search2 search' style={{display: 'none'}}>
                        <Search
                            className="heig-30"
                            placeholder="taobao.com"
                            size="large"
                            onSearch={value => console.log(value)}
                        />
                    </div>
                    <div className='select'>
                        <Dropdown overlay={menu} style={{ height: '40px'}}>
                            <Button style={{ height: '40px'}}>
                                 {seltVal}<Icon type="down" />
                            </Button>
                        </Dropdown>
                    </div>
                    <div className="btn">
                        <Button type="primary" className="heig-30">即时监测</Button>
                    </div>
                </div>
                <div className="btn">
                    <Button type="primary" className="heig-30">持续监测</Button>
                </div>
            </div>
        );
    }
}

export default SearchBarExample;