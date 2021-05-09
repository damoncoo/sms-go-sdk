package main

import (
	"github.com/damoncoo/sms-go-sdk/sms"
	"fmt"
)

func main() {

	params := make(map[string]interface{}, 0)
	params["code"] = 151713

	client := sms.NewClient()
	client.SetAppId("hw_10128")
	client.SetSecretKey("8a0db761fb7ac11a9d489c2c7d1cf138")

	request := sms.NewRequest()
	request.SetMethod("sms.message.send")
	request.SetBizContent(sms.TemplateMessage{
		Mobile:     []string{"13570276759"},
		Type:       0,
		Sign:       "易知",
		TemplateId: "ST_2020101100000007",
		SendTime:   "",
		Params:     params,
	})

	buf, err := client.Execute(request)
	fmt.Println(buf, err)

}
