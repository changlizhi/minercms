package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"minercms/log"
	"minercms/service"
	"minercms/utils"
	"net/http"
	"os"
	"encoding/json"
)

func BoyuCmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	if r.Method == "OPTIONS" {
		log.Error.Println("OPTIONS嗅探了")
		return
	}
	if r.Method != "POST" {
		log.Error.Println("非POST请求，不处理")
		w.Write(utils.CanShuTiShi("请使用POST进行请求"))
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	datab, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error.Println("Data byte 读取错误", err, r)
		w.Write(utils.CanShuTiShi("数据读取错误，请重试！"))
		return
	}

	dataObj := map[string]interface{}{}
	//从具体操作来说，拿到数据直接丢给分发应该没多大问题，但是如果参数没传上来那就不对应该直接返回错误信息提示
	err = json.Unmarshal(datab, &dataObj)
	if err != nil {
		log.Error.Println("datab解析错误", err)
		w.Write(utils.CanShuTiShi("参数错误，请注意约定好的参数格式！"))
		return
	}
	fmt.Println("获取  json 值 ==dataObj",dataObj)

	ret := service.FenFaYeWu(dataObj)
	log.Error.Println("分发后返回的数据直接转成json发给前端ret---", ret)
	retByte, err := json.Marshal(ret)
	if err != nil {
		log.Error.Println("返回数据组装json错误，这个错误不应该发生", err)
		w.Write(utils.CanShuTiShi("内部错误，请稍后重试！"))
		return
	}
	w.Write(retByte)
	return
}


func StartAPI() {
	http.HandleFunc("/cms/addNews", BoyuCmsHandler)
	http.HandleFunc("/cms/delNews", BoyuCmsHandler)
	http.HandleFunc("/cms/updateNews", BoyuCmsHandler)
	http.HandleFunc("/cms/queryNews", BoyuCmsHandler)
	http.HandleFunc("/cms/getOneNews", BoyuCmsHandler)

	// 浏览器接口
	//获取页面数据 Page Data
	http.HandleFunc("/cms/getPageData", BoyuCmsHandler)




	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Error.Println("服务端报错：", err.Error())
	}

}
func main() {
	log.Info.Println("-===================  启动  程序  ================")
	StartAPI()
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		log.Error.Println("main.go,YongHuHandler:err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		log.Error.Println("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
		if part.FileName() == "" { // this is FormData
			bodyByte, err := ioutil.ReadAll(part)
			log.Error.Println("main.go,35---", string(bodyByte)) //传文件的话字符没法单独传了
			if err != nil {
				w.Write([]byte("error,0:参数读取错误" + err.Error()))
				return
			}
			bodyStrArr := []string{}
			err = json.Unmarshal(bodyByte, &bodyStrArr)
			if err != nil {
				w.Write([]byte("error,1:参数格式错误"))
				return
			}
			if len(bodyStrArr) == 0 {
				w.Write([]byte("error,2:请勿传入空数组！"))
				return
			}
		} else { // This is FileData
			dst, _ := os.Create("./" + part.FileName() + ".upload")
			defer dst.Close()
			io.Copy(dst, part)
		}
	}

}
