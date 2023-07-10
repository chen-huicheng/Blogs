package main

import (
	"fmt"
	"regexp"
)

func main() {
	TestSetValTemplate()
}
func TestSetValTemplate() {
	tpl := "您好！您有一份来自${enterprise}的签署文件：https://xxx.net/${sign_url}，请于收到短信后的48小时内点击以上链接完成相关合同签署。"
	params := map[string]string{"enterprise": "字节跳动", "sign_url": "xSlj"}
	res, err := SetValTemplate(tpl, params)
	if err != nil {
		fmt.Println("[Error]:%s", err)
		return
	}
	fmt.Println(res)
}

func SetValTemplate(tpl string, params map[string]string) (string, error) {
	var FieldRegex = regexp.MustCompile(`\${[a-zA-Z0-9-_.]+}`)
	lackParam := false
	res := FieldRegex.ReplaceAllStringFunc(tpl, func(str string) (field string) {
		str = str[2 : len(str)-1]
		if v, ok := params[str]; ok {
			return v
		}
		lackParam = true
		return ""
	})
	if lackParam {
		return "", fmt.Errorf("缺少参数")
	}
	return res, nil
}
