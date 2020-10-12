package insert_table_data

import (
	"fmt"
	"log"
	"minercms/consts"
	"minercms/daos"
	"minercms/utils"
	"strconv"
	"testing"
)
func Find3MinerType(){

}
func TestInsertKuangChi(t *testing.T){
	js:= Find3JueSes()
	jss:=js[consts.ShuJu].([]map[string]interface{})
	js0:=jss[0]
	js1:=jss[1]
	js2:=jss[2]
	jueSeId:=""
	for i:=0;i<20;i++{
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
			consts.ShuJuBiao: consts.CaiDans,
			consts.ShuJuZhis: map[string]interface{}{
				consts.CaiDanId: id,
				consts.CaiDanBianMa: "1112222"+i2str,
				consts.CaiDanMingCheng: "YongHuBianMa"+i2str,
				consts.CaiDanUrl: "http://starriver.com/"+i2str,
				consts.JueSeIdCaiDan: jueSeId,
			},
		}

		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}
}
