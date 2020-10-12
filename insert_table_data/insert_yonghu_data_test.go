package insert_table_data

import (
	"fmt"
	"log"
	"minercms/consts"
	"minercms/daos"
	"minercms/utils"
	"strconv"
	"strings"
	"testing"
)
func tianChong0str(str string,geShu int)string{
	retBuilder:=strings.Builder{}
	for i := 0;i< geShu;i++{
		retBuilder.WriteString("0")
	}
	retBuilder.WriteString(str)
	return retBuilder.String()
}
func TestTianChong(t *testing.T){
	for i:=0;i<1000;i++{
		i2str:=strconv.Itoa(i)
		leni:=len(i2str)
		if leni < 3{
			i2str=tianChong0str(i2str,3-leni)
		}
		log.Println("i2str---",i2str)
	}
}

func FindJueSes()map[string]interface{}{
	canShu:=map[string]interface{}{
		consts.ShuJuKu: consts.MINERCMS,
		consts.ShuJuBiao: consts.JueSes,
		consts.ShuJuZhis: map[string]interface{}{
			consts.JueSeId: consts.JueSeId,
			consts.JueSeBianMa: consts.JueSeBianMa,
			consts.JueSeMingCheng: consts.JueSeMingCheng,
		},
		consts.TiaoJians: map[string]interface{}{
			consts.DangQianYe: "1",
			consts.MeiYeTiaoShu: "3",
		},
	}
	ret := daos.FindData(canShu)
	return ret
}

func TestJueSeFind(t *testing.T){

	js:=FindJueSes()
	js0:=js[consts.ShuJu].([]map[string]interface{})
	log.Println(js0,"----js0")
}
func TestInsertYongHu(t *testing.T){

	js:=FindJueSes()
	log.Println(js,"----js")
	for i:=0;i<1000;i++{
		id := utils.HuoQuIdZiFu()
		i2str:=strconv.Itoa(i)
		leni:=len(i2str)
		if leni < 3{
			i2str=tianChong0str(i2str,3-leni)
		}
		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.YongHus,
			consts.ShuJuZhis: map[string]interface{}{
				consts.YongHuId: id,
				consts.YongHuBianMa: "YongHuBianMa"+i2str,
				consts.SouJiHao: "1112222"+i2str,
				consts.YongHuMing: "YongHuBianMa"+i2str,
				consts.YouXiang: i2str+"@riverstr.com",
				consts.DengLuMiMa: "YongHuBianMa"+i2str,
				consts.ZhiFuMiMa: "YongHuBianMa"+i2str,
				consts.JueSeIdYongHu: "YongHuBianMa"+i2str,
				consts.ZhuCeShiJian: "YongHuBianMa"+i2str,
			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}
}
