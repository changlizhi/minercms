package createtable

import (
	"minercms/consts"
	"minercms/daos"
	"minercms/log"
	"testing"
)

func TestCreateKuangChi(t *testing.T){
	db := daos.HuoQuLianJieChi()

	log.Info.Println("db-----", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.KuangChis,
		consts.ZhuJian:   consts.KuangChiId,
		consts.SuoYin:    consts.KuangChiBianHao,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.KuangChiId,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.KuangChiBianHao,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.KuangJiMingCheng,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.KuangJiLeiXing,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'feitian'",
			},
		},
	}
	daos.ChuangJianBiao(canShu)
}