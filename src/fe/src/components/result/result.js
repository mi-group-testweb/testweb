import React from 'react';
import './result.css';
import { Table } from 'react-bootstrap';



class Result extends React.Component {



    render() {
        return (

            <div className="tab">
                <div className="left">
                    <h3>监测结果</h3>
                </div>

                <p>  {new Date().toLocaleTimeString()}</p>
               
    
    



                <Table striped bordered hover size="sm">
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
        )
    }
}
export default Result;





