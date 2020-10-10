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

func TestMysqlTypeTable(t *testing.T) {



	for i:=0;i<1000;i++{
		log.Printf("=== 开始插入   第 %d 条  数据 ====\n",i)
		id := utils.HuoQuIdZiFu()
		str:=strconv.Itoa(i)

		mid := str
		mname := "lily_"+str
		mtype := "yellow_"+str

		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.Power,
			consts.ShuJuZhis: map[string]interface{}{
				consts.Id: id,

				consts.MachineId:   mid,
				consts.MachineName: mname,
				consts.MachineType: mtype,
				//	consts.ShiJian:     shiJian,
				//	consts.GengXin:     gengXin,
			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}


}
