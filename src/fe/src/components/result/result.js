import React from 'react';
import './result.css';
import { Table } from 'react-bootstrap';
import ReactEcharts from 'echarts-for-react';


class Result extends React.Component {

    render() {
        let option = {
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'shadow'
                }
            },
            legend: {
                data: ['DNS解析', 'TCP链接', '首字节', '下载时间', 'SSL握手'],
                textStyle: {
                    fontSize: 14,
                    color: '#ffffff'
                },
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: {
                type: 'value',
                axisLabel: {
                    color: '#fff',
                    fontSize: 16,
                }
            },
            yAxis: {
                type: 'category',
                data: ['网站性能'],
                axisLabel: {
                    color: '#fff',
                    fontSize: 16,
                }
            },
            series: [
                {
                    name: 'DNS解析',
                    type: 'bar',
                    stack: '总量',
                    label: {
                        normal: {
                            show: true,
                            position: 'insideRight'
                        }
                    },
                    data: [320]
                },
                {
                    name: 'TCP链接',
                    type: 'bar',
                    stack: '总量',
                    label: {
                        normal: {
                            show: true,
                            position: 'insideRight'
                        }
                    },
                    data: [120]
                },
                {
                    name: '首字节',
                    type: 'bar',
                    stack: '总量',
                    label: {
                        normal: {
                            show: true,
                            position: 'insideRight'
                        }
                    },
                    data: [220]
                },
                {
                    name: '下载时间',
                    type: 'bar',
                    stack: '总量',
                    label: {
                        normal: {
                            show: true,
                            position: 'insideRight'
                        }
                    },
                    data: [150]
                },
                {
                    name: 'SSL握手',
                    type: 'bar',
                    stack: '总量',
                    label: {
                        normal: {
                            show: true,
                            position: 'insideRight'
                        }
                    },
                    data: [82]
                }
            ]
        }
        return (
            <div className="tab">
                <div className="left">
                    <span className="resuleName">监测结果</span>
                    <span>  {new Date().toLocaleTimeString()}</span>
                </div>
                <div className='echart_c'>
                    <h5 style={{ fontSize: '16px', textAlign: 'left', marginTop: '35px', width: '85%' }}>性能总览</h5>
                    <ReactEcharts
                        option={option}
                        style={{ height: '400px', width: '85%', backgroundColor: 'rgba(40,44,52,1)', padding: '30px 15px' }}
                        className='react_for_echarts' />
                </div>
                <div className='table_c'>
                    <h5 style={{ fontSize: '16px', textAlign: 'left', marginTop: '35px', width: '85%' }}>监测说明</h5>
                    <Table striped bordered hover size="sm" style={{ width: '85%' }}>
                        <thead>
                            <tr>
                                <td>站点的URL</td>
                                <td>DNS解析时间</td>
                                <td>TCP连接时间</td>
                                <td>首字节时间</td>
                                <td>下载时间</td>
                                <td>完成请求时间</td>
                                <td>SSL握手时间</td>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>account.mioffice.com</td>
                                <td>0.023</td>
                                <td>0.012</td>
                                <td>0.0112</td>
                                <td> 0.0234</td>
                                <td>0.024</td>
                                <td>0.045</td>
                            </tr>
                            <tr>
                                <td>2</td>
                                <td>Mark</td>
                                <td>Otto</td>
                                <td>@mdo</td>
                                <td> </td>
                                <td></td>
                                <td></td>
                            </tr>
                        </tbody>
                    </Table>
                </div>
            </div>
        )
    }
}
export default Result;





