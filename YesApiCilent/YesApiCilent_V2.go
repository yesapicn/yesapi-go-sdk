package YesApiCilent

/**
 * YesApi客户端SDK包-V2（Go版）
 * @package   YesApiCilent
 * @author    sHuXnHs <610087273@qq.com>
 */

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// 返回结构体
type APIResponse struct {
	Code int64       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type YesApiClient struct {
	YesapiHost       string
	YesapiAppKey     string
	YesapiAppSecrect string
}

func NewYesApiClient() *YesApiClient {
	return &YesApiClient{
		YesapiHost:       "", // TODO:填写你的小白接口域名，可以在http://open.yesapi.cn/?r=App/Mine查看
		YesapiAppKey:     "", // TODO:填写你的app_key
		YesapiAppSecrect: "", // TODO:填写你的app_secrect
	}
}

func (y *YesApiClient) SetYesapiHost(yesapiHost string) *YesApiClient {
	y.YesapiHost = yesapiHost
	return y
}

func (y *YesApiClient) SetYesapiAppKey(YesapiAppKey string) *YesApiClient {
	y.YesapiAppKey = YesapiAppKey
	return y
}

func (y *YesApiClient) SetYesapiAppSecrect(YesapiAppSecrect string) *YesApiClient {
	y.YesapiAppSecrect = YesapiAppSecrect
	return y
}

func (y *YesApiClient) DoRequest(param map[string]string) (APIResponse, error) {
	param["app_key"] = y.YesapiAppKey

	sign := y.encryptAppkey(param)

	str := y.makeParams(param)

	url := y.YesapiHost + "?sign=" + sign + str

	apiResponse := &APIResponse{}

	resp, err := http.Get(url)
	if err != nil {
		return *apiResponse, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	bytes, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(bytes, apiResponse)
	return *apiResponse, nil
}

// generate sign
func (y *YesApiClient) encryptAppkey(param map[string]string) string {
	var strs []string
	for k := range param {
		strs = append(strs, k)
	}
	sort.Strings(strs) // sort
	var paramstr string
	for _, k := range strs {
		paramstr += param[k]
	}
	paramstr += y.YesapiAppSecrect // add app_secrect
	h := md5.New()
	h.Write([]byte(paramstr))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

//拼接参数
func (y *YesApiClient) makeParams(params map[string]string) (str string) {
	if params == nil {
		return ""
	} else {
		for k := range params {
			str += "&" + k + "=" + params[k]
		}
		return str
	}
}
