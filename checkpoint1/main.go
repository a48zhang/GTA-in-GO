package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	url := "http://gtainmuxi.muxixyz.com/api/v1/organization/code"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("code", "0")
	clt := http.Client{}
	res, _ := clt.Do(req)
	fmt.Println(res)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
	//Passport:
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMSIsImlhdCI6MTY2OTM1NjU1NiwibmJmIjoxNjY5MzU2NTU2fQ.8QdF3lHxZXy6vtMt0d0MXGKNM6KeYZQjq6hSju5nGp8
}

/*
{
    "code": 0,
    "message": "OK",
    "data": {
        "text": "访问成功后，网站会给你返回信息，在header中找到你的passport。\n
		将passport加入到你以后的每次请求头中。\n
		完成上述步骤后，用代码访问 http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key，注意查收其中response的信息。",
        "extra_info": ""
    }
} */
