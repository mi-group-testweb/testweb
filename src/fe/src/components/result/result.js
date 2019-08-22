import React from 'react';
import './result.css';
import { Table } from 'react-bootstrap';

class Result extends React.Component {
    render() {
        return (
            <div className="tab">
                <Table striped bordered hover size="sm">
                    <thead>
                        <tr>
                            <th>URL</th>
                            <th>DNS解析时间</th>
                            <th>TCP连接时间</th>
                            <th>首字节时间</th>
                            <th>下载时间</th>
                            <th>完成请求时间</th>
                            <th>SSL握手时间</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>1</td>
                            <td>Mark</td>
                            <td>Otto</td>
                            <td>@mdo</td>
                            <td> </td>
                            <td></td>
                            <td></td>
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





