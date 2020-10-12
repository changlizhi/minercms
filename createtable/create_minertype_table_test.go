package createtable

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"testing"
)


func TestMysqlTypeTable(t *testing.T) {
	db := daos.HuoQuLianJieChi()
	log.Println("=== db ===", db)


	//创建 类型
	log.Println("=== 开始创建类型表 ===")
	datapower := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.MinerTypes,
		consts.ZhuJian:   consts.MinerTypeId,
		consts.SuoYin:    consts.TypeEncode,
		consts.ZiDuans: []interface{}{
			map[string]interface{}{
				consts.BianMa:   consts.MinerTypeId,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.TypeEncode,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'leixing1'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.Mainboard,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "100",
				consts.MoRenZhi: "'Supermicro'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.TypeName,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.GPU,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.CPU,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.SSD,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.HDD,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
		},
	}
	daos.ChuangJianBiao(datapower)
	log.Println("=== 创建类型表完成 ===")
}