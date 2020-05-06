import moment from 'moment';
import axios, { AUTHAPI }  from './axios';
const getPubKey = () => {
    return axios({
        method: 'get',
        url: `${AUTHAPI}pubKey`
        // url: 'pubKey'
    }).then(data => {
        if (data.code === '000')
            return data.data.pubkey;
        else
            return null;
    });
};
const disabledDate =(current)=> {
    // Can not select days before today and today
    return current && current > moment(new Date() );
}
//获取昨天
const yesterday =()=> {
    let Time =   new Date();
    Time.setTime(new Date().getTime()-24*60*60*1000);
    return  Time
}
export {
    disabledDate,
    yesterday,
    getPubKey
}