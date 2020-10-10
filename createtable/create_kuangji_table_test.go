package createtable

import (
	"minercms/consts"
	"minercms/daos"
	"minercms/log"
	"testing"
)

func TestCreateKuangJi(t *testing.T){
	db := daos.HuoQuLianJieChi()

	log.Info.Println("db-----", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.KuangJis,
		consts.ZhuJian:   consts.KuangJiId,
		consts.SuoYin:    consts.KuangJiBianHaoKuangJi,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.KuangJiId,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.KuangJiBianHaoKuangJi,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.YongHuIdKuangJi,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
		},
	}
	daos.ChuangJianBiao(canShu)
}