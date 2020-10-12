package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)


func TestMysqlEarningTable(t *testing.T) {
	db := daos.HuoQuLianJieChi()

	log.Println("=== db ===", db)

	//======================================
	//创建 收益表
	log.Println("=== 开始创建收益表 ===")
	dataEarning := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.Earning,
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
				consts.BianMa:   consts.CurrentGrossRevenue,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "100",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.NowTime,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'IPFS-Earnings'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.UserIdEarning,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
		},
	}
	daos.ChuangJianBiao(dataEarning)
	log.Println("=== 创建收益表完成 ===")

}