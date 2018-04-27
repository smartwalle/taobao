package taobao

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////
// TBKGetItemsParam 淘宝客商品查询
// http://open.taobao.com/docs/api.htm?spm=a219a.7629065.0.0.YOvS3k&apiId=24515
type TBKGetItemsParam struct {
	Fields      []string // 必须 需返回的字段列表
	Q           string   // 特殊可选 查询词
	Cats        []string // 特殊可选 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	ItemLoc     string   // 可选 所在地
	Sort        string   // 可选 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi）
	IsTmall     bool     // 可选 是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性
	IsOverseasa bool     // 可选 是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性
	StartPrice  float64  // 可选 折扣价范围下限，单位：元
	EndPrice    float64  // 可选 折扣价范围上限，单位：元
	StartTKRate int      // 可选 淘客佣金比率上限，如：1234表示12.34%
	EndTKRate   int      // 可选 淘客佣金比率下限，如：1234表示12.34%
	Platform    int      // 可选 链接形式：1：PC，2：无线，默认：１
	PageNo      int      // 可选 第几页，默认：１
	PageSize    int      // 可选 页大小，默认20，1~100
}

func (this TBKGetItemsParam) APIName() string {
	return "taobao.tbk.item.get"
}

func (this TBKGetItemsParam) Params() map[string]string {
	var m = make(map[string]string)
	if len(this.Fields) > 0 {
		m["fields"] = strings.Join(this.Fields, ",")
	} else {
		m["fields"] = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick"
	}
	if this.Q != "" {
		m["q"] = this.Q
	}
	if len(this.Cats) > 0 {
		m["cat"] = strings.Join(this.Cats, ",")
	}
	if this.ItemLoc != "" {
		m["itemloc"] = this.ItemLoc
	}
	if this.Sort != "" {
		m["sort"] = this.Sort
	}
	if this.IsTmall {
		m["is_tmall"] = "true"
	} else {
		m["is_tmall"] = "false"
	}
	if this.IsOverseasa {
		m["is_overseas"] = "true"
	} else {
		m["is_overseas"] = "false"
	}
	if this.StartPrice > 0 {
		m["start_price"] = fmt.Sprintf("%.2f", this.StartPrice)
	}
	if this.EndPrice > 0 {
		m["end_price"] = fmt.Sprintf("%.2f", this.StartPrice)
	}
	if this.StartTKRate > 0 {
		m["start_tk_rate"] = fmt.Sprintf("%d", this.StartTKRate)
	}
	if this.EndTKRate > 0 {
		m["end_tk_rate"] = fmt.Sprintf("%d", this.EndTKRate)
	}

	if this.Platform <= 0 {
		this.Platform = 1
	}
	m["platform"] = fmt.Sprintf("%d", this.Platform)

	if this.PageNo > 0 {
		m["page_no"] = fmt.Sprintf("%d", this.PageNo)
	}
	if this.PageSize > 0 {
		m["page_size"] = fmt.Sprintf("%d", this.PageSize)
	}

	return m
}

func (this TBKGetItemsParam) ExtJSONParamName() string {
	return ""
}

func (this TBKGetItemsParam) ExtJSONParamValue() string {
	return ""
}

type TBKItem struct {
	NumIID       int    `json:"num_iid"`        // 商品ID
	Title        string `json:"title"`          // 商品标题
	PictURL      string `json:"pict_url"`       // 商品主图
	SmallImages  string `json:"small_images"`   // 商品小图列表
	ReservePrice string `json:"reserve_price"`  // 商品一口价格
	ZKFinalPrice string `json:"zk_final_price"` // 商品折扣价格
	UserType     int    `json:"user_type"`      // 卖家类型，0表示集市，1表示商城
	Provcity     string `json:"provcity"`       // 宝贝所在地
	ItemURL      string `json:"item_url"`       // 商品地址
	Nick         string `json:"nick"`           // 卖家昵称
	SellerID     string `json:"seller_id"`      // 卖家id
	Volume       int    `json:"volume"`         // 30天销量
}

type tbkGetItemsRsp struct {
	Rsp struct {
		Results struct {
			TBKItems []*TBKItem `json:"n_tbk_item"`
		} `json:"results"`
		TotalResults int `json:"total_results"`
	} `json:"tbk_item_get_response"`
	Err *ErrorResponse `json:"error_response"`
}

type TBKGetItemsRsp struct {
	TBKItems     []*TBKItem `json:"n_tbk_item"`
	TotalResults int        `json:"total_results"`
}
