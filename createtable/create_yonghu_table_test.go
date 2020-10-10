package createtable

import (
	"minercms/consts"
	"minercms/daos"
	"minercms/log"
	"testing"
)

func TestCreateYongHu(t *testing.T){
	db := daos.HuoQuLianJieChi()

	log.Info.Println("db-----", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.YongHus,
		consts.ZhuJian:   consts.YongHuId,
		consts.SuoYin:    consts.YongHuBianMa,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.YongHuId,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.YongHuBianMa,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.SouJiHao,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.YongHuMing,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.YouXiang,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.DengLuMiMa,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.ZhiFuMiMa,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.JueSeIdYongHu,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},

		},
	}
	daos.ChuangJianBiao(canShu)
}