package main

import (
	"./YesApiCilent"
	"fmt"
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
}

//  返回示例
//	code： 200
//	data： map[err_code:0 err_msg: title:Hi HXH，欢迎使用小白接口！]
//	msg： 当前小白接口：App.Hello.World
