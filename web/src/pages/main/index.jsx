import React, { Component } from 'react';
import './index.styl';
import { Form, Icon, Input, Button, notification, Row, Col, Tag, Select, Drawer } from 'antd';
import axios from '@/common/axios';
import { connect } from 'react-redux';
const FormItem = Form.Item;
const { Option } = Select;
class MainPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loading:false,
            tagOptions:[],
            checkTag:[],
            formData:{},
            dataDetail:[]
        };
    }
    componentDidMount(){
        this.getTag()
    }
    searchInfo(value){
        this.state.formData[value.key]= value.type==='select'?value.value :value.value.target.value;
        let newData ={};
        for( let i in this.state.formData){
            if(this.state.formData[i]){
                newData[i]=this.state.formData[i];
            }
        }
        this.setState({formData:newData })
        
        
    }
    searchInfoBtn(){
        if(!Object.keys(this.state.formData).length) return; 
        axios({
            method:'post',
            url:'search',
            data:this.state.formData
        }).then((data)=>{
            if(data.code !==200) return;
            const {searchItems, pageData} = data.Data;
            this.setState({
                dataDetail:searchItems,
                visible:true
            })
        })
    }
    getTag(){
        axios({
            method:'get',
            url:'menu'
        }).then((data)=>{
            if(+data.code !==200) return;
            const { unitData } = data.data;
            let formData = {};
            const newUnitData=unitData.map(item=>{
                formData[item.unitEN]='';
                return {
                    ...item,
                    check:false
                }
            })
            this.setState({tagOptions:newUnitData,formData })
        })
    }
    checkTagChange(data){
        const { tagOptions } = this.state;
        this.setState({tagOptions:[]});
        this.state.formData[data.unitEN] = '';
        const newUnitData=tagOptions.map(item=>{
            return {
                ...item,
                check:data.unitEN===item.unitEN? !item.check:item.check
            }
        })
        this.setState({tagOptions:newUnitData})
    }
    onClose(){
        this.setState({
            visible: false,
            dataDetail:[]
        });
    };
    render() {
        const { getFieldDecorator } = this.props.form;
        const { loading, tagOptions, formData, visible, dataDetail  } = this.state;
        return (
            <div className="index-search">
                <Row className="flex-row" gutter={16}>
                    <Col className="flex-col"  span={10}>
                        <div className="gutter-row input-box">
                            {tagOptions.map((item)=>{
                                if(!item.check) return<span></span>;
                                if(item.type==='text')
                                    return  <Input className="search-input"
                                        allowClear
                                        onChange={($event)=>this.searchInfo({ key: item.unitEN,
                                            value:  $event })}
                                        placeholder={item.unitCN}  />
                                if(item.type==='select')
                                    return  <Select
                                                showSearch
                                                allowClear
                                                onChange={($event)=>this.searchInfo({ key: item.unitEN,
                                                                                      value:  $event, type:'select' })}
                                                className="search-input"
                                                placeholder={item.unitCN}
                                            >
                                            {
                                                item.unitValue.map(itemValue=>{
                                                    return <Option value={itemValue}>{itemValue}</Option>
                                                })
                                            }
                                        </Select>
                            })}
                        </div>
                    </Col>
                    <Col className="flex-col" span={14}>
                        <div className="gutter-row tags">
                            {tagOptions.map((item)=>{
                                return  <Tag
                                    color={
                                        item.check?'#cd201f':'#87d068'}
                                        key={item.unitEN}
                                        onClick={()=>this.checkTagChange(item)}>
                                        {item.unitCN}
                                    </Tag>
                            })}
                            <Button size='small' type="primary" disabled={!Object.values(formData).find(items=>items)} onClick={this.searchInfoBtn.bind(this)}>
                            查询
                            </Button>
                        </div>
                    </Col>
                </Row>
                <Drawer
                    title="详细数据"
                    placement="right"
                    closable={false}
                    width={600}
                    onClose={this.onClose.bind(this)}
                    visible={visible}
                    >
                    <div>
                    {/* name: "freeCodeCamp/freeCodeCamp"
description: "freeCodeCamp.org's open source codebase and curriculum. Learn to code at home."
stars: 311000
labels: ["curriculum", "certification", "react", "nodejs", "javascript", "d3", "teachers", "community",…]
language: "JavaScript" */}
                        {
                            dataDetail&&dataDetail[0]?dataDetail.map(item=>{
                                return(
                                    <div style={{marginBottom:10,borderBottom: '1px dashed #ddd', padding:'10px 0'}}>
                                        <p style={{fontWeight:600,fontSize: 16 }}>{item.name}</p>
                                        <p>{item.description}</p>
                                        <div>{
                                           item.labels&&item.labels[0]? item.labels.map((labelItem=>{
                                                return(
                                                    <Tag
                                                    color='#87d068'
                                                    style={{marginBottom:10}}
                                                        key={labelItem}
                                                    >
                                                        {labelItem}
                                                    </Tag>  
                                                )
                                            })):null
                                        }
                                        </div>
                                        <div>
                                        <Tag color="#55acee">
                                            {item.stars}
                                        </Tag>
                                        <Tag color="#3b5999">
                                            {item.language||'未知'}
                                        </Tag> 
                                            {/* <span>{item.stars}</span>
                                            <span>{item.language}</span> */}
                                        </div>
                                    </div>
                                )
                            }):null
                        }
                    </div>
                    {/* <p>Some contents...</p>
                    <p>Some contents...</p>
                    <p>Some contents...</p> */}
                </Drawer>
            </div>
        );
    }
}
const WrappedNormalLoginForm = Form.create()(MainPage);
export default connect()(WrappedNormalLoginForm);
