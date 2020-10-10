package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)

func TestMysqlPool(t *testing.T) {
	db := daos.HuoQuLianJieChi()

	log.Println("db-----", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.News,
		consts.ZhuJian:   consts.Id,
		consts.SuoYin:    consts.BiaoTi,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.Id,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.BiaoTi,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "200",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.LeiXing,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.ShiJian,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.GengXin,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.SuoLueTu,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "500",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.JianJie,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "300",
				consts.MoRenZhi: "'boyu'",
			},
			map[string]interface{}{
				consts.BianMa:  consts.NeiRong,
				consts.LeiXing: consts.MEDIUMTEXT,
			},
		},
	}
	daos.ChuangJianBiao(canShu)
}

