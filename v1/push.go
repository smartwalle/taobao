package taobao

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
// CloudPushParam 向所有平台设备推送通知的参数
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.CxuRVy&apiId=25040
type CloudPushParam struct {
	Title                   string  // 必须 推送的标题内容.
	Body                    string  // 必须 推送内容
	DeviceType              int     // 必须 设备类型,取值范围为:0~3云推送支持多种设备,各种设备类型编号如下: iOS设备:deviceType=0; Andriod设备:deviceType=1;如果存在此字段,则向指定的设备类型推送消息。 默认为全部(3);
	Target                  string  // 必须 推送目标: device:推送给设备; account:推送给指定帐号,tag:推送给自定义标签; all: 推送给全部
	TargetValue             string  // 必须 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔.(帐号与设备有一次最多100个的限制)
	Remind                  bool    // 必须 当APP不在线时候，是否通过通知提醒. 针对不同设备，处理逻辑不同。 该参数只针对iOS设备生效， (remind=true & 发送消息的话(type=0)). 当你的目标设备不在线(既长连接通道不通, 我们会将这条消息的标题，通过苹果的apns通道再送达一次。发apns是发送生产环境的apns，需要在云推送配置的app的iOS生产证书和密码需要正确，否则也发送不了。 (remind=false & 并且是发送消息的话(type=0))，那么设备不在线，则不会再走苹果apns发送了。
	StoreOffline            bool    // 必须 离线消息是否保存,若保存, 在推送时候，用户即使不在线，下一次上线则会收到

	AndroidActivity         string  // 可选 Android对应的activity,仅仅当androidOpenType=2有效
	AndroidMusic            string  // 可选 android通知声音
	AndroidOpenType         string  // 可选 点击通知后动作,1:打开应用 2: 打开应用Activity 3:打开 url
	AndroidOpenUrl          string  // 可选 Android收到推送后打开对应的url,仅仅当androidOpenType=3有效
	AntiHarassDuration      string  // 可选 防打扰时长,取值范围为1~23
	AntiHarassStartTime     string  // 可选 防打扰开始时间点,取值范围为0~23
	BatchNumber             string  // 可选 批次编号,用于活动效果统计
	IOSBadge                string  // 可选 iOS应用图标右上角角标
	IOSMusic                string  // 可选 iOS通知声音
	Summery                 string  // 可选 通知的摘要
	Timeout                 int     // 可选 离线消息保存时长,取值范围为1~72,若不填,则表示不保存离线消息
	Type                    int     // 可选 0:表示消息(默认为0),1:表示通知

	androidExtParameters    map[string]string  // 可选 自定义的kv结构,开发者扩展用 针对android
	iOSExtParameters        map[string]string  // 可选 自定义的kv结构,开发者扩展用 针对iOS设备
}

func (this CloudPushParam) APIName() string {
	return "taobao.cloudpush.push"
}

func (this CloudPushParam) Params() map[string]string {
	var m = make(map[string]string)
	m["target"]                 = this.Target
	m["target_value"]           = this.TargetValue
	if len(this.AndroidActivity) > 0 {
		m["android_activity"]       = this.AndroidActivity
	}
	if len(this.androidExtParameters) > 0 {
		m["android_ext_parameters"] = marshal(this.androidExtParameters)
	}
	if len(this.AndroidMusic) > 0 {
		m["android_music"]          = this.AndroidMusic
	}
	if len(this.AndroidOpenType) > 0 {
		m["android_open_type"]      = this.AndroidOpenType
	}
	if len(this.AndroidOpenUrl) > 0 {
		m["android_open_url"]       = this.AndroidOpenUrl
	}
	if len(this.AntiHarassDuration) > 0 {
		m["anti_harass_duration"]   = this.AntiHarassDuration
	}
	if len(this.AntiHarassStartTime) > 0 {
		m["anti_harass_start_time"] = this.AntiHarassStartTime
	}
	if len(this.BatchNumber) > 0 {
		m["batch_number"]           = this.BatchNumber
	}
	m["body"]                   = this.Body
	m["device_type"]            = fmt.Sprintf("%d", this.DeviceType)
	if len(this.IOSBadge) > 0 {
		m["ios_badge"]              = this.IOSBadge
	}
	if len(this.iOSExtParameters) > 0 {
		m["ios_ext_parameters"] = marshal(this.iOSExtParameters)
	}
	if len(this.IOSMusic) > 0 {
		m["ios_music"]              = this.IOSMusic
	}
	if this.Remind {
		m["remind"]             = "true"
	} else {
		m["remind"]             = "false"
	}
	if this.StoreOffline {
		m["store_offline"]      = "true"
	} else {
		m["store_offline"]      = "false"
	}

	if len(this.Summery) > 0 {
		m["summery"]            = this.Summery
	} else {
		m["summery"]            = this.Title
	}

	if this.Timeout > 0 {
		m["timeout"]            = fmt.Sprintf("%d", this.Timeout)
	}
	m["title"]                  = this.Title
	m["type"]                   = fmt.Sprintf("%d", this.Type)
	return m
}

func (this CloudPushParam) ExtJSONParamName() string {
	return ""
}

func (this CloudPushParam) ExtJSONParamValue() string {
	return ""
}

func (this *CloudPushParam) AddAndroidExtParam(key, value string) {
	if this.androidExtParameters == nil {
		this.androidExtParameters = make(map[string]string)
	}
	this.androidExtParameters[key] = value
}

func (this *CloudPushParam) AddiOSExtParam(key, value string) {
	if this.iOSExtParameters == nil {
		this.iOSExtParameters = make(map[string]string)
	}
	this.iOSExtParameters[key] = value
}

////////////////////////////////////////////////////////////////////////////////
// CloudPushNoticeToIOSParam 向 iOS 平台设备推送 APNS 通知的参数
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.DXAaM1&apiId=25041
type CloudPushNoticeToIOSParam struct {
	Summary     string                  // 必须  通知摘要
	Target      string                  // 必须  推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	TargetValue string                  // 必须  根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	Environment string                  // 必须  iOS的通知是通过APNS中心来发送的，需要填写对应的环境信息. DEV:表示开发环境, PRODUCT: 表示生产环境.
	Ext         map[string]interface{}  // 可选  提供给IOS通知的扩展属性，如角标或者声音等,注意：参数值为json
}

func (this CloudPushNoticeToIOSParam) APIName() string {
	return "taobao.cloudpush.notice.ios"
}

func (this CloudPushNoticeToIOSParam) Params() map[string]string {
	var m = make(map[string]string)
	m["summary"] = this.Summary
	m["target"]  = this.Target
	m["target_value"] = this.TargetValue
	m["env"] = this.Environment
	return m
}

func (this CloudPushNoticeToIOSParam) ExtJSONParamName() string {
	return "ext"
}

func (this CloudPushNoticeToIOSParam) ExtJSONParamValue() string {
	return marshal(this.Ext)
}

func (this *CloudPushNoticeToIOSParam) AddParam(key string, value interface{}) {
	if this.Ext == nil {
		this.Ext = make(map[string]interface{})
	}
	this.Ext[key] = value
}

////////////////////////////////////////////////////////////////////////////////
// CloudPushNoticeToAndroidParam 向 Android 平台设备推送通知消息的参数
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.4lW0QW&apiId=25039
type CloudPushNoticeToAndroidParam struct {
	Title       string  // 必须  通知的标题.
	Summary     string  // 必须  通知摘要
	Target      string  // 必须  推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	TargetValue string  // 必须  根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
}

func (this CloudPushNoticeToAndroidParam) APIName() string {
	return "taobao.cloudpush.notice.android"
}

func (this CloudPushNoticeToAndroidParam) Params() map[string]string {
	var m = make(map[string]string)
	m["summary"] = this.Summary
	m["target"]  = this.Target
	m["target_value"] = this.TargetValue
	m["title"] = this.Title
	return m
}

func (this CloudPushNoticeToAndroidParam) ExtJSONParamName() string {
	return ""
}

func (this CloudPushNoticeToAndroidParam) ExtJSONParamValue() string {
	return ""
}

////////////////////////////////////////////////////////////////////////////////
// CloudPushMessageToIOSParam 向 iOS 平台设备推送消息的参数（非 APNS 通知）
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.w7TlXB&apiId=25037
type CloudPushMessageToIOSParam struct {
	Body        string  // 必须  发送的消息内容.
	Target      string  // 必须  推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	TargetValue string  // 必须  根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
}

func (this CloudPushMessageToIOSParam) APIName() string {
	return "taobao.cloudpush.message.ios"
}

func (this CloudPushMessageToIOSParam) Params() map[string]string {
	var m = make(map[string]string)
	m["body"] = this.Body
	m["target"]  = this.Target
	m["target_value"] = this.TargetValue
	return m
}

func (this CloudPushMessageToIOSParam) ExtJSONParamName() string {
	return ""
}

func (this CloudPushMessageToIOSParam) ExtJSONParamValue() string {
	return ""
}

func (this *CloudPushMessageToIOSParam) AddParam(key string, value interface{}) {
}

////////////////////////////////////////////////////////////////////////////////
// CloudPushMessageToAndroidParam 向 Android 平台设备推送消息的参数
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.z3SYCA&apiId=25038
type CloudPushMessageToAndroidParam struct {
	CloudPushMessageToIOSParam
}

func (this CloudPushMessageToAndroidParam) APIName() string {
	return "taobao.cloudpush.message.android"
}
