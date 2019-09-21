package YesApiCilent

import (
	"testing"
)

/**
 * YesApi客户端单元测试,性能测试
 * @package   YesApiCilent
 * @author    sHuXnHs <610087273@qq.com>
 */

func TestDoRequest(t *testing.T) {
	param := map[string]string{"s": "App.Hello.World", "name": "HXH"}
	rs, err := DoRequest(param)
	if err != nil {
		t.Log("请求失败")
		t.FailNow()
	} else {
		t.Log("code：", rs.Code)
		t.Log("data：", rs.Data)
		t.Log("msg：", rs.Msg)
	}
}

func BenchmarkDoRequest(b *testing.B) {
	param := map[string]string{"s": "App.Hello.World", "name": "HXH"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = DoRequest(param)
	}
}
