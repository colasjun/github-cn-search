import axios from 'axios';
import { notification } from 'antd';
// import cookie from 'react-cookies';
// export const API = process.env.REACT_APP_URL;
axios.defaults.headers.post["Content-Type"] =
  "application/x-www-form-urlencoded";
axios.defaults.baseURL = 'http://47.108.24.233:8090/';

axios.interceptors.request.use(
    function (config) {
        // if(config.url.includes('file/template/download')){
        //     config.headers.get['Content-Type']= 'application/excel'
        //     config.responseType = 'arraybuffer'
        // }
        // if(cookie.load('innerapp')&&!config.url.includes('/open/pubKey'))
        //     config.headers.Authorization = `Bearer ${cookie.load('innerapp')}`
        return config;
    },
  function(error) {
    notification.error({
        message:'错误',
        description: '操作失败！',
        duration: 2
    });
    return Promise.reject(error);
  }
);
axios.interceptors.response.use(
    function(response) {
    // 对响应数据做点什么
    if(response.config.url.includes('file/template/download')){
        return response;
    }
    if ( !['000','0','033','034',200].includes(response.data.code) )
        if(['rp9901','au1002'].includes(response.data.code)){
            notification.error({
                message:'错误',
                description: 'token过期，请重新登录',
                duration: 2
            });
            // cookie.remove('reportToken');
            setTimeout(()=>{
                window.location.href = '/';
            },2000)
        }else
            notification.warn({
                message:'警告',
                description: response.data.msg||response.data.message||'操作失败！',
                duration: 2
            });
        return response.data;
    },
    function(error) {
    // 对响应错误做点什么
        if (error.response && error.response.status === 401) {
            notification.error({
                message:'错误',
                description: 'token过期，请重新登录',
                duration: 2
            });
            // cookie.remove('reportToken');
            setTimeout(()=>{
                window.location.href = '/';
            },2000)
        } else if (!navigator.onLine) {
            notification.error({
                message:'错误',
                description: '请检查网络！',
                duration: 2
            });
        } else {
            notification.error({
                message:'提示',
                description: '操作失败！',
                duration: 2
            });
        }
        return Promise.reject(error);
    }
);
export default axios;
