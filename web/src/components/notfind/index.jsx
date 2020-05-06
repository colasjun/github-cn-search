import React from 'react';
import './index.styl';
class NotFoundPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
        };
    }
    render(){
        return (
            <div className="notFoundBox">
                <img alt="无数据"
                    src={require('@/images/404.svg')}
                    width="40%"/>
                <h1 className="notStatus">404</h1>
                <p className="fontStyle">抱歉,你访问的页面不存在!</p>
            </div>
        )
    }
}
export default NotFoundPage;