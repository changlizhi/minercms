package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)

func TestCreateJueSe(t *testing.T){
	db := daos.HuoQuLianJieChi()

	log.Println("db-----", db)
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