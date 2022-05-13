package main

import (
	. "byx-script-go/interpreter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type runRequest struct {
	Script string `json:"script"`
}

type runResponse struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
}

func getOutput(executable func()) string {
	file, _ := ioutil.TempFile("", "tmp")
	oldStdout := os.Stdout
	os.Stdout = file

	defer func() {
		os.Stdout = oldStdout
	}()

	executable()

	bytes, _ := ioutil.ReadFile(file.Name())
	return string(bytes)
}

func handler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		// 解码请求json字符串
		decoder := json.NewDecoder(req.Body)
		var reqData runRequest
		err := decoder.Decode(&reqData)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		defer func() {
			r := recover()
			if r != nil {
				resp.Header().Set("Access-Control-Allow-Origin", "*")
				resp.Header().Set("Access-Control-Allow-Methods", "*")
				resp.Header().Add("Access-Control-Allow-Headers", "*")
				resp.Header().Set("content-type", "application/json")

				respData := runResponse{false, fmt.Sprintf("%v", r)}
				bytes, err := json.Marshal(respData)
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
					return
				}
				_, err = resp.Write(bytes)
				if err != nil {
					resp.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}()

		// 执行脚本，并获取标准输出内容
		output := getOutput(func() {
			RunScript(reqData.Script, RunConfig{})
		})

		// 构造响应数据
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Methods", "*")
		resp.Header().Add("Access-Control-Allow-Headers", "*")
		resp.Header().Set("content-type", "application/json")

		bytes, err := json.Marshal(runResponse{true, output})
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = resp.Write(bytes)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodOptions:
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Methods", "*")
		resp.Header().Add("Access-Control-Allow-Headers", "*")
	default:
		resp.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ByxScript在线运行服务端
func main() {
	http.HandleFunc("/run", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
