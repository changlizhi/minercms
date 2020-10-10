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

func TestInsertInviteData(t *testing.T) {

	for i := 0; i < 1000; i++ {
		log.Printf("=== 开始插入   第 %d 条  数据 ====\n", i)
		id := utils.HuoQuIdZiFu()
		str := strconv.Itoa(i)

		userid :=   str
		reference := "nick" + str
        vipCode:="ABCD"+str
		insertNeedCanShu := map[string]interface{}{
			consts.ShuJuKu:   consts.MINERCMS,
			consts.ShuJuBiao: consts.Invite,
			consts.ShuJuZhis: map[string]interface{}{
				consts.Id: id,

				consts.UserIdInvite:   userid,
				consts.Reference: reference,
				consts.VipInviteCode:   vipCode,

			},
		}
		ret := daos.InsertData(insertNeedCanShu)
		fmt.Println("ret=", ret)
		log.Printf("=== 插入   第 %d 条  数据  完成====\n", i)

	}

}
