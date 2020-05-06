import React, { Component } from 'react';
import { Route, Redirect, Switch } from 'react-router-dom';
import { LocaleProvider } from 'antd';
import zhCN from 'antd/lib/locale-provider/zh_CN';
import 'moment/locale/zh-cn';
import NotFoundPage from '@/components/notfind';
import MainPage from '@/pages/main';
import './App.css';
// import withRuleRouter from '@/components/utils/withRuleRouter';
class App extends Component {
    
    render() {
        return (
            <div className="App">
                <LocaleProvider locale={zhCN}>
                    <div className="containBox">
                        <Switch>
                            <Route exact
                                path="/"
                                render={()=><Redirect push
                                    to="/main"></Redirect>}  >
                            </Route>
                            <Route path="/main"
                                component={MainPage}/>
                            <Route path="*"
                                component={NotFoundPage}/>
                        </Switch>
                    </div>
                </LocaleProvider>
            </div>
        );
    }
}
export default App;
