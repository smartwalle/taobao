package taobao

func (this *TaoBao) OpenIMGetUsers(param OpenIMGetUsersParam) (result []*OpenIMUserInfo, err error) {
	var r *openIMGetUsersRsp
	err = this.doRequest(param, &r)
	if r != nil {
		if r.Err != nil {
			return nil, r.Err
		}
		return r.Rsp.Userinfos.UserInfos, err
	}
	return result, err
}

func (this *TaoBao) OpenIMAddUsers(param OpenIMAddUsersParam) (result *OpenIMAddUsersRsp, err error) {
	var r *openIMAddUsersRsp
	err = this.doRequest(param, &r)
	if r != nil {
		if r.Err != nil {
			return nil, r.Err
		}
		result = &OpenIMAddUsersRsp{}
		result.SuccUIDList = r.Rsp.UIDSuccList.String
		result.FailUIDList = r.Rsp.UIDFailList.String
		result.FailMSgList = r.Rsp.FailMsgList.String
		return result, nil
	}
	return result, err
}

func (this *TaoBao) OpenIMUpdateUsers(param OpenIMUpdateUsersParam) (result *OpenIMUpdateUsersRsp, err error) {
	var r *openIMUpdateUsersRsp
	err = this.doRequest(param, &r)
	if r != nil {
		if r.Err != nil {
			return nil, r.Err
		}
		result = &OpenIMUpdateUsersRsp{}
		result.SuccUIDList = r.Rsp.UIDSuccList.String
		result.FailUIDList = r.Rsp.UIDFailList.String
		result.FailMSgList = r.Rsp.FailMsgList.String
		return result, nil
	}
	return result, err
}

func (this *TaoBao) OpenIMDeleteUsers(param OpenIMDeleteUsersParam) (result *OpenIMDeleteUsersRsp, err error) {
	var r *openIMDeleteUsersRsp
	err = this.doRequest(param, &r)
	if r != nil {
		if r.Err != nil {
			return nil, r.Err
		}
		result = &OpenIMDeleteUsersRsp{}
		result.Result = r.Rsp.Result.String
		return result, nil
	}
	return result, err
}
