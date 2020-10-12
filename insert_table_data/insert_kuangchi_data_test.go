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

func Find3MinerTypes()map[string]interface{}{
	canShu:=map[string]interface{}{
		consts.ShuJuKu: consts.MINERCMS,
		consts.ShuJuBiao: consts.MinerTypes,
		consts.ShuJuZhis: map[string]interface{}{
			consts.MinerTypeId: consts.MinerTypeId,
			consts.TypeName:    consts.TypeName,
			consts.TypeEncode:  consts.TypeEncode,
		},
		consts.TiaoJians: map[string]interface{}{
			consts.DangQianYe: "1",
			consts.MeiYeTiaoShu: "3",
		},
	}
	ret := daos.FindData(canShu)
	return ret
}
func TestFind3MinerTypes(t *testing.T){
	ret :=Find3MinerTypes()
log.Println("ret---",ret)
}
func TestInsertKuangChi(t *testing.T){
	ts:= Find3MinerTypes()
	tss:=ts[consts.ShuJu].([]map[string]interface{})
	ts0:=tss[0]
	ts1:=tss[1]
	ts2:=tss[2]
	useId:=""
	for i:=0;i<20;i++{
		id := utils.HuoQuIdZiFu()
		i2str:=strconv.Itoa(i)
		leni:=len(i2str)
		if leni < 3{
			i2str=tianChong0str(i2str,3-leni)
		}
		if i%3==0{
			useId=ts2[consts.MinerTypeId].(string)
		}else if i %2==0{
			useId=ts1[consts.MinerTypeId].(string)
		}else if i %1==0{
			useId=ts0[consts.MinerTypeId].(string)
		}
		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.KuangChis,
			consts.ShuJuZhis: map[string]interface{}{
				consts.KuangChiId: id,
				consts.KuangChiBianHao: "bianhao"+i2str,
				consts.KuangJiMingCheng: "http://starriver.com/"+i2str,
				consts.KuangJiLeiXing: useId,
			},
		}

		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}
}
