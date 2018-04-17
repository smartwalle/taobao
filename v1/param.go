package taobao

import (
	"encoding/json"
)

////////////////////////////////////////////////////////////////////////////////
type ITaoBaoParam interface{
	// 用于提供访问的 method
	APIName() string

	// 返回参数列表
	Params() map[string]string

	// 返回扩展 JSON 参数的字段名称
	ExtJSONParamName() string
	// 返回扩展 JSON 参数的字段值
	ExtJSONParamValue() string
}

////////////////////////////////////////////////////////////////////////////////
// 示例，别无它用
type TaoBaoParam map[string]interface{}

func (this TaoBaoParam) APIName() string {
	return "taobao.open.sms.sendmsg"
}

func (this TaoBaoParam) Params() map[string]string {
	return nil
}

func (this TaoBaoParam) ExtJSONParamName() string {
	return "send_message_request"
}

func (this TaoBaoParam) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (this TaoBaoParam) AddParam(key string, value interface{}) {
	this[key] = value
}


func marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
