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

func Find3JueSes()map[string]interface{}{
	canShu:=map[string]interface{}{
		consts.ShuJuKu: consts.MINERCMS,
		consts.ShuJuBiao: consts.YongHus,
		consts.ShuJuZhis: map[string]interface{}{
			consts.YongHuId: consts.YongHuId,
			//consts.JueSeBianMa: consts.JueSeBianMa,
		//	consts.JueSeMingCheng: consts.JueSeMingCheng,
		},
		consts.TiaoJians: map[string]interface{}{
			consts.DangQianYe: "1",
			consts.MeiYeTiaoShu: "2",
		},
	}
	ret := daos.FindData(canShu)
	return ret
}
func TestMod(t *testing.T){
	str := "12"
	for i := 0;i< 1000;i ++{
		if i%3==0{
			str = "3"
			//log.Println("i%100==0",i)
		}else if i%2==0{
			str = "2"
			//log.Println("i%100==0",i)

		}else {
			log.Println("i%101==0",i)
			str = "1"

		}
		log.Println("str---",str)
	}
}
func TestJueSeFind(t *testing.T){


	js:= Find3JueSes()
	jss:=js[consts.ShuJu].([]map[string]interface{})
	log.Println(jss[0],"----jss[0]")
	log.Println(jss[1],"----jss[1]")
	log.Println(jss[2],"----jss[2]")

}
func TestInsertYongHu(t *testing.T){
	js:= Find3JueSes()
	jss:=js[consts.ShuJu].([]map[string]interface{})
	js0:=jss[0]
	js1:=jss[1]
	js2:=jss[2]
	jueSeId:=""
	for i:=0;i<900;i++{
		id := utils.HuoQuIdZiFu()
		i2str:=strconv.Itoa(i)
		leni:=len(i2str)
		if leni < 3{
			i2str=tianChong0str(i2str,3-leni)
		}
		if i%3==0{
			jueSeId=js2[consts.JueSeId].(string)
			log.Println("jueSeId---i%3==0",jueSeId)
		}else if i %2==0{
			jueSeId=js1[consts.JueSeId].(string)
			log.Println("jueSeId---i%2==0",jueSeId)
		}else if i %1==0{
			jueSeId=js0[consts.JueSeId].(string)
			log.Println("jueSeId---i%1==0",jueSeId)
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
				consts.JueSeIdYongHu: jueSeId,
				consts.ZhuCeShiJian: "YongHuBianMa"+i2str,
			},
		}

		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}
}
