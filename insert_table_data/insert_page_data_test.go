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

func TestInsertPageData(t *testing.T) {

	for i := 0; i < 1000; i++ {
		log.Printf("=== 开始插入   第 %d 条  数据 ====\n", i)
		id := utils.HuoQuIdZiFu()
		str := strconv.Itoa(i)

		tp := "星河飞天_标题_" + str
		info := "信息说明_" + str
		data := "天下武功，唯快不破。" + str
		intro := "倚天不出，谁与争锋。" + str
		details:= "天若有情天亦老" + str
		thu := "stariver" + str
		nowTime:=utils.TimeNow()

		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.PageData,
			consts.ShuJuZhis: map[string]interface{}{
				consts.Id: id,

				consts.TitleName:   tp,
				consts.DataType: info,

				consts.TimeInfo:   nowTime,
				consts.Data: data,
				consts.Thumbnail:  thu,
				consts.Introduce: intro,
				consts.Details:  details,



			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n", i)

	}

}
