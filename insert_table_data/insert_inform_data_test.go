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

func TestInsertInformData(t *testing.T) {

	for i := 0; i < 1000; i++ {
		log.Printf("=== 开始插入   第 %d 条  数据 ====\n", i)
		id := utils.HuoQuIdZiFu()
		str := strconv.Itoa(i)

		tel := "188000" + str
		email := "stariver" + str

		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.Inform,
			consts.ShuJuZhis: map[string]interface{}{
				consts.Id: id,

				consts.Tel:   tel,
				consts.Email: email,

				consts.NoInformTel:   "0",
				consts.NoInformEmail: "0",
				consts.UserIDInform:  str,
			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n", i)

	}

}
