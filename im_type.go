package taobao

import "strings"

type OpenIMUserInfo struct {
	UserId   string `json:"userid"`             // 必须 im用户名
	Password string `json:"password,omitempty"` // 必须 im密码

	Nick     string `json:"nick,omitempty"`     // 可选 昵称
	Status   int    `json:"status,omitempty"`   // 可选 用户激活状态，0表示未激活，1表示激活
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

type openIMGetUsersRsp struct {
	Rsp struct {
		Userinfos struct {
			UserInfos []*OpenIMUserInfo `json:"userinfos"`
		} `json:"userinfos"`
	} `json:"openim_users_get_response"`
	Err *ErrorResponse `json:"error_response"`
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

type openIMAddUsersRsp struct {
	Rsp struct {
		UIDSuccList struct {
			String []string `json:"string"`
		} `json:"uid_succ"`
		UIDFailList struct {
			String []string `json:"string"`
		} `json:"uid_fail"`
		FailMsgList struct {
			String []string `json:"string"`
		} `json:"fail_msg"`
	} `json:"openim_users_add_response"`
	Err *ErrorResponse `json:"error_response"`
}

type OpenIMAddUsersRsp struct {
	SuccUIDList []string
	FailUIDList []string
	FailMSgList []string
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

type openIMUpdateUsersRsp struct {
	Rsp struct {
		UIDSuccList struct {
			String []string `json:"string"`
		} `json:"uid_succ"`
		UIDFailList struct {
			String []string `json:"string"`
		} `json:"uid_fail"`
		FailMsgList struct {
			String []string `json:"string"`
		} `json:"fail_msg"`
	} `json:"openim_users_update_response"`
	Err *ErrorResponse `json:"error_response"`
}

type OpenIMUpdateUsersRsp struct {
	SuccUIDList []string
	FailUIDList []string
	FailMSgList []string
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

type openIMDeleteUsersRsp struct {
	Rsp struct {
		Result struct {
			String []string `json:"string"`
		} `json:"result"`
	} `json:"openim_users_delete_response"`
	Err *ErrorResponse `json:"error_response"`
}

type OpenIMDeleteUsersRsp struct {
	Result []string
}
