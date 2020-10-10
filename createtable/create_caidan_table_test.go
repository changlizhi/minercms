package createtable

import (
	"minercms/consts"
	"minercms/daos"
	"minercms/log"
	"testing"
)

func TestCreateCaiDan(t *testing.T){
	db := daos.HuoQuLianJieChi()

	log.Info.Println("db-----", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.CaiDans,
		consts.ZhuJian:   consts.CaiDanId,
		consts.SuoYin:    consts.CaiDanBianMa,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.CaiDanId,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.CaiDanBianMa,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.CaiDanMingCheng,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.CaiDanUrl,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.JueSeIdCaiDan,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
		},
	}
	daos.ChuangJianBiao(canShu)
}