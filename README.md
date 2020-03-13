![](http://cdn7.okayapi.com/CEE4B8A091578B252AC4C92FB4E893C3_20190304213902_63f85e982adc8419feffd862e883581e.jpeg)

# yesapi-Golang-sdk
小白接口（YesApi.cn）Golang SDK包

## 使用示例

```go
package main

import (
    "fmt"
    "yesapi-go-sdk/YesApiCilent"
)

func main() {
	param := map[string]string{"s": "App.Hello.World" , "name" : "HXH"}

	rs, err := YesApiCilent.DoRequest(param)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("code：", rs.Code)
		fmt.Println("data：", rs.Data)
		fmt.Println("msg：", rs.Msg)
	}

    // 第二版改进
	yesapiClient := YesApiCilent.NewYesApiClient()
	// 可以做账号的切换，也可以在NewYesApiClient()方法写死配置
	yesapiClient.SetYesapiHost("").SetYesapiAppKey("").SetYesapiAppSecrect("")
	res, err := yesapiClient.DoRequest(param)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("code：", res.Code)
	fmt.Println("data：", res.Data)
	fmt.Println("msg：", res.Msg)

}
```

## 返回示例
```
code： 200
data： map[err_code:0 err_msg: title:Hi HXH，欢迎使用小白接口！]
msg： 当前小白接口：App.Hello.World
```
