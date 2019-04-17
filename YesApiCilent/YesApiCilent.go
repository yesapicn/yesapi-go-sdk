package YesApiCilent

/**
 * YesApi客户端SDK包（Go版）
 * @package   YesApiCilent
 * @author    sHuXnHs <610087273@qq.com>
 */

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const YESAPI_Host string = ""			//TODO:填写你的小白接口域名，可以在http://open.yesapi.cn/?r=App/Mine查看
const YESAPI_APP_KEY string = ""		//TODO:填写你的app_key
const YESAPI_APP_SECRECT string = ""	//TODO:填写你的app_secrect

type result struct {
	Code int64       `json:"ret"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func DoRequest(param map[string]string) (*result, error) {

	//添加app_key
	param["app_key"] = YESAPI_APP_KEY

	//生成client 超时时间为10s
	client := http.Client{Timeout: 10 * time.Second}

	//生成签名
	sign := encryptAppkey(param)

	str := MakeParams(param)

	url := YESAPI_Host + "?sign=" + sign + str

	//提交请求
	reqest, err :=  client.Get(url)

	if err != nil {
		return nil, err
	} else {
		defer reqest.Body.Close()
		return dealResult(reqest)
	}

}

//生成签名
func encryptAppkey(param map[string]string) string{

	var strs []string
	for k := range param {
		strs = append(strs, k)
	}
	sort.Strings(strs)	//排序
	//paramstr为参数值排序后得字符串
	var paramstr string
	for _, k := range strs {
		paramstr += param[k]
	}
	//追加app_secrect
	paramstr += YESAPI_APP_SECRECT
	h := md5.New()
	h.Write([]byte(paramstr))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

//拼接参数
func MakeParams(params map[string]string) (str string) {
	if params == nil {
		return ""
	}else {
		for k := range params {
			str += "&"+ k + "=" + params[k]
		}
		return str
	}
}

//返回处理结果
func dealResult(response *http.Response) (*result, error) {
	if response.Status == "200 OK" {
		ret := new(result)
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, ret)
		return ret, nil
	} else {
		return nil, errors.New(response.Status)
	}
}

