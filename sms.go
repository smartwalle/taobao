package taobao

func (this *TaoBao) OpenSMSSendMsg(param *OpenSMSSendMsgParam) (result *OpenSMSSendMsgRsp, err error) {
	var r *openSMSSendMsgRsp
	err = this.doRequest(param, &r)
	if r != nil {
		if r.Err != nil {
			return nil, r.Err
		}
		return r.Rsp.Result, nil
	}
	return nil, err
}