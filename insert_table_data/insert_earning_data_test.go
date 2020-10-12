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

func TestInsertEarningData(t *testing.T) {



	for i:=0;i<1000;i++{
		log.Printf("=== 开始插入   第 %d 条  数据 ====\n",i)
		id := utils.HuoQuIdZiFu()
		str:=strconv.Itoa(i)

		mid := str
		mtype := "yellow_"+str

		// 获取时间
		nowTime:=utils.TimeNow()
		FindUserId()




		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.Earning,
			consts.ShuJuZhis: map[string]interface{}{
				consts.Id: id,

				consts.CurrentGrossRevenue:   mid,
				consts.NowTime: nowTime,
				consts.UserIdEarning: mtype,
			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}


}
func FindUserId()map[string]interface{}{
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
