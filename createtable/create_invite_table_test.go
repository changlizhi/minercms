package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)


func TestMysqlInviteTable(t *testing.T) {
	db := daos.HuoQuLianJieChi()
	log.Println("=== db ===", db)


	//创建 类型
	log.Println("=== 开始创建邀请表 ===")
	datapower := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.Invite,
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
				consts.BianMa:   consts.UserIdInvite,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "100",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Reference,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'IPFS-Earnings'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.VipInviteCode,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
		},
	}
	daos.ChuangJianBiao(datapower)
	log.Println("=== 创建邀请表完成 ===")
}