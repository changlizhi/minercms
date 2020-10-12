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



	for i:=1;i<=10;i++{
		log.Printf("=== 开始插入   第 %d 条  数据 ====\n",i)
		id := utils.HuoQuIdZiFu()
		str:=strconv.Itoa(i)

		mid := str
		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.MinerTypes,
			consts.ShuJuZhis: map[string]interface{}{
				consts.MinerTypeId: id,
				consts.TypeEncode:  "leixing"+mid,
				consts.Mainboard:   "supermicro"+mid,
				consts.TypeName:    "leixing_CN"+mid,
				consts.GPU:         "NVIDIA"+mid,
				consts.CPU:         "AMD"+mid,
				consts.SSD:         mid+"TiB",
				consts.HDD:         mid+"0TiB",
			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n",i)

	}


}
