package taobao

func (this *TaoBao) TBKGetItems(param TBKGetItemsParam) (result *TBKGetItemsRsp, err error) {
	var r *tbkGetItemsRsp
	err = this.doRequest(param, &r)
	if r != nil {
		if r.Err != nil {
			return nil, r.Err
		}
		result = &TBKGetItemsRsp{}
		result.TBKItems = r.Rsp.Results.TBKItems
		result.TotalResults = r.Rsp.TotalResults
		return result, err
	}
	return result, err
}
