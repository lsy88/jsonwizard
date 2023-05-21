package utils

import (
	"testing"
)

func TestEncryptJSON(t *testing.T) {
	jsonString := `{
		"name":"json加密测试",
			"url":"https://www.baidu.com",
			"address":{
			"city":"新乡",
				"country":"中国"
		},
		"arrayBrowser":[{
	"name":"Google",
	"url":"http://www.google.com"
	},
	{
	"name":"Baidu",
	"url":"http://www.baidu.com"
	},
	{
	"name":"SoSo",
	"url":"http://www.SoSo.com"
	}]
	}`
	json, err := EncryptJSON(jsonString, "list123111111111list123111111111")
	if err != nil {
		t.Error(err)
	}
	decryptJSON, err := DecryptJSON(json, "list123111111111list123111111111")
	if err != nil {
		t.Error(err)
	}
	if decryptJSON != jsonString {
		t.Error("failed")
	} else {
		t.Log("success")
	}
}
