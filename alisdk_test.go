// This file is auto-generated, don't edit it. Thanks.
package main

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dyvmsapi20170525 "github.com/alibabacloud-go/dyvmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dyvmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dyvmsapi.aliyuncs.com")
	_result = &dyvmsapi20170525.Client{}
	_result, _err = dyvmsapi20170525.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
	if _err != nil {
		return _err
	}

	addRtcAccountRequest := &dyvmsapi20170525.AddRtcAccountRequest{
		ResourceOwnerAccount: tea.String("your_value"),
		ResourceOwnerId:      tea.Int64(1),
		DeviceId:             tea.String("your_value"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.AddRtcAccount(addRtcAccountRequest)
	if _err != nil {
		return _err
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
