package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)


func TestMysqlPageData(t *testing.T) {
	db := daos.HuoQuLianJieChi()

	log.Println("=== 创建 PageData ===", db)
	canShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.PageData,
		consts.ZhuJian:   consts.Id,
		consts.SuoYin:    consts.TitleName,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.Id,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.TitleName,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "100",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.DataType,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'IPFS-PageData'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.TimeInfo,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Data,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "300",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Thumbnail,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "500",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Introduce,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "300",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:  consts.Details,
				consts.LeiXing: consts.MEDIUMTEXT,
			},
		},
	}
	daos.ChuangJianBiao(canShu)

}