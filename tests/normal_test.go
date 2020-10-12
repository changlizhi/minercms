package tests

import (
	"log"
	"minercms/consts"
	"minercms/utils"
	"reflect"
	"testing"
	"time"
)

func TestString2Time(t *testing.T) {
	ret := utils.String2Time("2020-03-16 20:20:20", consts.NyrSfm)
	log.Println(ret)
}
func TestString2TimeNyrSfm(t *testing.T) {
	ret := utils.String2TimeNyrSfm("2020-03-16 20:20:20")
	log.Println(ret)
}

func TestHuoQuIdInt(t *testing.T) {
	ret := utils.HuoQuId()
	log.Println(ret)
}
func TestHuoQuIdZiFu(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go func(i int) {
			ret := utils.HuoQuIdZiFu()
			log.Println("TestHuoQuIdZiFu---", i, ret, len(ret))
		}(i)
	}
	time.Sleep(time.Duration(5) * time.Hour)

}

func TestReflect(t *testing.T) {
	ziFu := "123"
	shuZi := 123
	shuZu := []string{"123"}
	mapStr := map[string]string{"123": "123"}
	log.Println(
		reflect.TypeOf(ziFu).String(),
		reflect.TypeOf(ziFu).Kind(),
		reflect.TypeOf(shuZi).String(),
		reflect.TypeOf(shuZi).Kind(),
		reflect.TypeOf(shuZu).String(),
		reflect.TypeOf(shuZu).Kind(),
		reflect.TypeOf(mapStr).String(),
		reflect.TypeOf(mapStr).Kind(),
	)
}
func TestBaoHan(t *testing.T) {
	ziFu := "123"
	shuZi := 123
	mapStr := map[string]string{"123": "123"}
	var allNewsLeiXing = []string{
		"IPFS",
		"HangYeDongTai",
		"GongSiFaBu",
	}
	log.Println(
		utils.ShiFouBaoHan("12345", ziFu),
		utils.ShiFouBaoHan([]string{"123", "345"}, ziFu),
		utils.ShiFouBaoHan(12245, shuZi),
		utils.ShiFouBaoHan(mapStr, ziFu),
		utils.ShiFouBaoHan(allNewsLeiXing, "IPFSs"),
	)

}