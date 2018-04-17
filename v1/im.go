package taobao

import (
	"strings"
)

type OpenIMUserInfo struct {
	UserId   string `json:"userid"`             // 必须 im用户名
	Password string `json:"password,omitempty"` // 必须 im密码

	Nick     string `json:"nick,omitempty"`     // 可选 昵称
	IconURL  string `json:"icon_url,omitempty"` // 可选 头像url
	Email    string `json:"email,omitempty"`    // 可选 email地址
	Mobile   string `json:"mobile,omitempty"`   // 可选 手机号码
	TaoBaoId string `json:"taobaoid,omitempty"` // 可选 淘宝账号
	Remark   string `json:"remark,omitempty"`   // 可选 remark
	Extra    string `json:"extra,omitempty"`    // 可选 扩展字段（Json）
	Career   string `json:"career,omitempty"`   // 可选 职位
	VIP      string `json:"vip,omitempty"`      // 可选 vip（Json）
	Address  string `json:"address,omitempty"`  // 可选 地址
	Name     string `json:"name,omitempty"`     // 可选 名字
	Age      int    `json:"age,omitempty"`      // 可选 年龄
	Gender   string `json:"gender,omitempty"`   // 可选 性别。M: 男。 F：女
	WeChat   string `json:"wechat,omitempty"`   // 可选 微信
	QQ       string `json:"qq,omitempty"`       // 可选 qq
	WeiBo    string `json:"weibo,omitempty"`    // 可选 微博
}

type OpenIMUser struct {
	UserId        string `json:"uid"`            // 可选 用户id
	TaoBaoAccount bool   `json:"taobao_account"` // 可选 是否为淘宝账号
	AppKey        string `json:"app_key"`        // 可选 账户appkey
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMAddUsersParam 添加 IM 用户
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.pMpmGG&apiId=24164&docType=
type OpenIMAddUsersParam struct {
	userInfoList []*OpenIMUserInfo `json:"userinfos"` // 必须  用户信息列表
}

func (this OpenIMAddUsersParam) APIName() string {
	return "taobao.openim.users.add"
}

func (this OpenIMAddUsersParam) Params() map[string]string {
	return nil
}

func (this OpenIMAddUsersParam) ExtJSONParamName() string {
	return "userinfos"
}

func (this OpenIMAddUsersParam) ExtJSONParamValue() string {
	return marshal(this.userInfoList)
}

func (this *OpenIMAddUsersParam) AddOpenIMUser(user *OpenIMUserInfo) {
	if user == nil {
		return
	}
	if this.userInfoList == nil {
		this.userInfoList = make([]*OpenIMUserInfo, 0, 0)
	}
	this.userInfoList = append(this.userInfoList, user)
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMGetUsersParam 批量获取 IM 用户信息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.JOoCKr&apiId=24157
type OpenIMGetUsersParam struct {
	UserIds []string // 必须 多个用户用半角逗号分隔，最多一次获取100个用户
}

func (this OpenIMGetUsersParam) APIName() string {
	return "taobao.openim.users.get"
}

func (this OpenIMGetUsersParam) Params() map[string]string {
	var m = make(map[string]string)
	m["userids"] = strings.Join(this.UserIds, ",")
	return m
}

func (this OpenIMGetUsersParam) ExtJSONParamName() string {
	return ""
}

func (this OpenIMGetUsersParam) ExtJSONParamValue() string {
	return ""
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMDeleteUsersParam 批量删除 IM 用户信息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.z7KKvy&apiId=24160
type OpenIMDeleteUsersParam struct {
	OpenIMGetUsersParam
}

func (this OpenIMDeleteUsersParam) APIName() string {
	return "taobao.openim.users.delete"
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMUpdateUsers 批量更新用户信息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.JcWhN0&apiId=24161
type OpenIMUpdateUsersParam struct {
	OpenIMAddUsersParam
}

func (this OpenIMUpdateUsersParam) APIName() string {
	return "taobao.openim.users.update"
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMPushMsgParam 发送标准 IM 消息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.f7Uzbf&apiId=25731
type OpenIMPushMsgParam struct {
	FromUser  string   `json:"from_user,omitempty"`  // 可选 消息发送者
	ToUsers   []string `json:"to_users,omitempty"`   // 可选 消息接受者
	MsgType   int      `json:"msg_type,omitempty"`   // 可选 消息类型。0:文本消息。1:图片消息，只支持jpg、gif。2:语音消息，只支持amr。8:地理位置信息。
	Context   string   `json:"context,omitempty"`    // 可选 发送的消息内容。根据不同消息类型，传不同的值。0(文本消息):填消息内容字符串。1(图片):base64编码的jpg或gif文件。2(语音):base64编码的amr文件。8(地理位置):经纬度，格式如 111,222
	ToAppKey  string   `json:"to_appkey,omitempty"`  // 可选 接收方appkey，默认本app，跨app发送时需要用到
	MediaAttr string   `json:"media_attr,omitempty"` // 可选 json map，媒体信息属性。根据msgtype变化。0(文本):填空即可。 1(图片):需要图片格式，{"type":"jpg"}或{"type":"gif"}。 2(语音): 需要文件格式和语音长度信息{"type":"amr","playtime":5}
}

func (this OpenIMPushMsgParam) APIName() string {
	return "taobao.openim.immsg.push"
}

func (this OpenIMPushMsgParam) Params() map[string]string {
	return nil
}

func (this OpenIMPushMsgParam) ExtJSONParamName() string {
	return "immsg"
}

func (this OpenIMPushMsgParam) ExtJSONParamValue() string {
	return marshal(this)
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMPushCustomMsgParam 发送自定义 IM 消息
// http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.FWTE56&apiId=25185
type OpenIMPushCustomMsgParam struct {
	FromUser string                 `json:"from_user,omitempty"`  // 可选 消息发送者
	ToUsers  []string               `json:"to_users,omitempty"`   // 必须 消息接受者
	ToAppKey string                 `json:"to_appkey,omitempty"`  // 可选 接收方appkey，默认本app，跨app发送时需要用到
	Summary  string                 `json:"summary,omitempty"`    // 必须 客户端最近消息里面显示的消息摘要
	Data     string                 `json:"data,omitempty"`       // 可选 发送的自定义数据，sdk默认无法解析消息，该数据需要客户端自己解析
	Aps      map[string]interface{} `json:"aps,omitempty"`        // 可选 apns推送时，里面的aps结构体json字符串，aps.alert为必填字段。本字段为可选，若为空，则表示不进行apns推送
	ApsParam map[string]interface{} `json:"apns_param,omitempty"` // 可选 apns推送的附带数据。
}

func (this OpenIMPushCustomMsgParam) APIName() string {
	return "taobao.openim.custmsg.push"
}

func (this OpenIMPushCustomMsgParam) Params() map[string]string {
	return nil
}

func (this OpenIMPushCustomMsgParam) ExtJSONParamName() string {
	return "custmsg"
}

func (this OpenIMPushCustomMsgParam) ExtJSONParamValue() string {
	return marshal(this)
}

func (this *OpenIMPushCustomMsgParam) setAps(key string, value interface{}) {
	if this.Aps == nil {
		this.Aps = make(map[string]interface{})
	}
	this.Aps[key] = value
}

func (this *OpenIMPushCustomMsgParam) SetApsAlert(alert string) {
	this.setAps("alert", alert)
}

func (this *OpenIMPushCustomMsgParam) SetApsBadge(badge int) {
	this.setAps("badge", badge)
}

func (this *OpenIMPushCustomMsgParam) SetApsSound(sound string) {
	this.setAps("sound", sound)
}

func (this *OpenIMPushCustomMsgParam) AddApsParam(key string, value interface{}) {
	if this.ApsParam == nil {
		this.ApsParam = make(map[string]interface{})
	}
	this.ApsParam[key] = value
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMCreateTribeParam 创建群
// http://open.taobao.com/doc2/apiDetail.htm?spm=a219a.7395905.0.0.JQ2jX6&apiId=25570
type OpenIMCreateTribeParam struct {
	User      *OpenIMUser   `json:"user"`       // 必须 用户信息
	TribeName string        `json:"tribe_name"` // 必须 群名称
	Notice    string        `json:"notice"`     // 必须 群公告
	TribeType string        `json:"tribe_type"` // 必须 群类型有两种tribe_type = 0 普通群 普通群有管理员角色，对成员加入有权限控制tribe_type = 1 讨论组 讨论组没有管理员，不能解散
	Members   []*OpenIMUser `json:"members"`    // 可选 创建群时候拉入群的成员tribe_type = 1（即为讨论组类型)时 该参数为必选; tribe_type = 0 (即为普通群类型)时，改参数无效，可不填
}

func (this OpenIMCreateTribeParam) APIName() string {
	return "taobao.openim.tribe.create"
}

func (this OpenIMCreateTribeParam) Params() map[string]string {
	var m = make(map[string]string)
	m["tribe_name"] = this.TribeName
	m["notice"] = this.Notice
	m["tribe_type"] = this.TribeType
	m["user"] = marshal(this.User)

	if this.Members != nil {
		m["members"] = marshal(this.Members)
	}
	return m
}

func (this OpenIMCreateTribeParam) ExtJSONParamName() string {
	return ""
}

func (this OpenIMCreateTribeParam) ExtJSONParamValue() string {
	return ""
}

func (this OpenIMCreateTribeParam) AddMember(m *OpenIMUser) {
	if m == nil {
		return
	}
	if this.Members == nil {
		this.Members = make([]*OpenIMUser, 0, 0)
	}
	this.Members = append(this.Members, m)
}

////////////////////////////////////////////////////////////////////////////////
// OpenIMGetTribeInfoParam 获取群信息
// http://open.taobao.com/doc2/apiDetail.htm?spm=a219a.7395905.0.0.9IDIj5&apiId=25571
type OpenIMGetTribeInfoParam struct {
	User    *OpenIMUser `json:"user"`     // 必须 用户信息
	TribeId string      `json:"tribe_id"` // 必须 群id
}

func (this OpenIMGetTribeInfoParam) APIName() string {
	return "taobao.openim.tribe.gettribeinfo"
}

func (this OpenIMGetTribeInfoParam) Params() map[string]string {
	var m = make(map[string]string)
	m["tribe_id"] = this.TribeId
	m["user"] = marshal(this.User)
	return m
}

func (this OpenIMGetTribeInfoParam) ExtJSONParamName() string {
	return ""
}

func (this OpenIMGetTribeInfoParam) ExtJSONParamValue() string {
	return ""
}
