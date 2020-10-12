package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)


func TestMysqlInformTable(t *testing.T) {
	db := daos.HuoQuLianJieChi()
	log.Println("=== db ===", db)


	//创建 类型
	log.Println("=== 开始创建通知表 ===")
	datapower := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.Inform,
		consts.ZhuJian:   consts.Id,
		consts.SuoYin:    "",
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.Id,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Tel,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "100",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Email,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'IPFS-Earnings'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.NoInformEmail,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.NoInformTel,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.UserIDInform,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},

		},
	}
	daos.ChuangJianBiao(datapower)
	log.Println("=== 创建通知表完成 ===")
}