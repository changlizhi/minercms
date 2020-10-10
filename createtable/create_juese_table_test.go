package createtable

import (
	"minercms/consts"
	"minercms/daos"
	"minercms/log"
	"testing"
)

func TestCreateJueSe(t *testing.T){
	db := daos.HuoQuLianJieChi()

	log.Info.Println("db-----", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.JueSes,
		consts.ZhuJian:   consts.JueSeId,
		consts.SuoYin:    consts.JueSeBianMa,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.JueSeId,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.JueSeBianMa,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.JueSeMingCheng,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
		},
	}
	daos.ChuangJianBiao(canShu)
}