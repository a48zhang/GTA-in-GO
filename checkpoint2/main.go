package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	url := "http://gtainmuxi.muxixyz.com/api/v1/organization/secret_key"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("code", "0")
	req.Header.Set("Passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMCIsImlhdCI6MTY2OTM1NzQxMiwibmJmIjoxNjY5MzU3NDEyfQ.ez49O-DVgO1YCaOZ40z5wWnXq7RTJB8cnXrclnUQjJo")
	cli := http.Client{}
	res, _ := cli.Do(req)

	fmt.Println(res)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

/* "text":"恭喜你拿到了 passport，现在你可以着手准备骇入银行。\n银行的第一道门是代码安全门，我们计划将错误代码写入此门来破解它。\n
但是这个门具有识别明文代码的功能，所以我们还需要一个密钥加密我们的错误代码，再写入至此门。\n
不需要担心，两者我们都为你提供了，你只需要解析 response 中的密文（在 ExtraInfo 中）来得到它们。\n
你现在的任务：\n解析密文，获取 error_code 和 secret_key\n使用 secret_key 加密 error_code\n然后将加密过的 error_code 放入 请求body
并以 \"正确的请求方法\" 发送至 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate , 同时注意response的信息。",
"extra_info":"c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="}} */
