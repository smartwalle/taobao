package taobao

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/smartwalle/ngx"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	k_TAO_BAO_OPEN_SANDBOX_API_URL    = "https://gw.api.tbsandbox.com/router/rest"
	k_TAO_BAO_OPEN_PRODUCTION_API_URL = "https://eco.taobao.com/router/rest"
)

type TaoBao struct {
	appKey       string
	appSecret    string
	isProduction bool
	apiDomain    string
}

func New(appKey, appSecret string, isProduction bool) (client *TaoBao) {
	client = &TaoBao{}
	client.appKey = appKey
	client.appSecret = appSecret
	client.isProduction = isProduction
	if isProduction {
		client.apiDomain = k_TAO_BAO_OPEN_PRODUCTION_API_URL
	} else {
		client.apiDomain = k_TAO_BAO_OPEN_SANDBOX_API_URL
	}
	return client
}

func (this *TaoBao) URLValues(param TaoBaoParam) (value url.Values, err error) {
	var p = url.Values{}
	var keys = make([]string, 6, 6)

	p.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	p.Add("format", "json")
	p.Add("v", "2.0")
	p.Add("sign_method", "md5")
	p.Add("app_key", this.appKey)
	p.Add("method", param.APIName())

	keys[0] = "timestamp"
	keys[1] = "format"
	keys[2] = "v"
	keys[3] = "sign_method"
	keys[4] = "app_key"
	keys[5] = "method"

	if len(param.ExtJSONParamName()) > 0 {
		p.Add(param.ExtJSONParamName(), param.ExtJSONParamValue())
		keys = append(keys, param.ExtJSONParamName())
	}

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	sign, err := sign(this.appSecret, keys, p)
	if err != nil {
		return nil, err
	}
	p.Add("sign", sign)
	return p, nil
}

func (this *TaoBao) doRequest(param TaoBaoParam, result interface{}) (err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return err
	}
	var req = ngx.NewRequest("POST", this.apiDomain)
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.SetParams(p)

	var rsp = req.Exec()
	fmt.Println(rsp.String())
	err = rsp.UnmarshalJSON(result)
	return err
}

func (this *TaoBao) DoRequest(param TaoBaoParam, result interface{}) (err error) {
	return this.doRequest(param, result)
}

func sign(appSecret string, keys []string, param url.Values) (s string, err error) {
	for _, key := range keys {
		s = s + key + param.Get(key)
	}
	s = fmt.Sprintf("%s%s%s", appSecret, s, appSecret)
	var m = md5.New()
	if _, err = m.Write([]byte(s)); err != nil {
		return "", err
	}
	s = strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
	return s, nil
}
