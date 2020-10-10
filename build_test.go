package main

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"minercms/utils"
	"reflect"
	"testing"
	"time"
)

func TestString2Time(t *testing.T) {
	ret := utils.String2Time("2020-03-16 20:20:20", consts.NyrSfm)
	log.Println(ret)
}
func TestString2TimeNyrSfm(t *testing.T) {
	ret := utils.String2TimeNyrSfm("2020-03-16 20:20:20")
	log.Println(ret)
}

func TestHuoQuIdInt(t *testing.T) {
	ret := utils.HuoQuId()
	log.Println(ret)
}
func TestHuoQuIdZiFu(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go func(i int) {
			ret := utils.HuoQuIdZiFu()
			log.Println("TestHuoQuIdZiFu---", i, ret, len(ret))
		}(i)
	}
	time.Sleep(time.Duration(5) * time.Hour)

}

func TestReflect(t *testing.T) {
	ziFu := "123"
	shuZi := 123
	shuZu := []string{"123"}
	mapStr := map[string]string{"123": "123"}
	log.Println(
		reflect.TypeOf(ziFu).String(),
		reflect.TypeOf(ziFu).Kind(),
		reflect.TypeOf(shuZi).String(),
		reflect.TypeOf(shuZi).Kind(),
		reflect.TypeOf(shuZu).String(),
		reflect.TypeOf(shuZu).Kind(),
		reflect.TypeOf(mapStr).String(),
		reflect.TypeOf(mapStr).Kind(),
	)
}
func TestBaoHan(t *testing.T) {
	ziFu := "123"
	shuZi := 123
	mapStr := map[string]string{"123": "123"}
	var allNewsLeiXing = []string{
		"IPFS",
		"HangYeDongTai",
		"GongSiFaBu",
	}
	log.Println(
		utils.ShiFouBaoHan("12345", ziFu),
		utils.ShiFouBaoHan([]string{"123", "345"}, ziFu),
		utils.ShiFouBaoHan(12245, shuZi),
		utils.ShiFouBaoHan(mapStr, ziFu),
		utils.ShiFouBaoHan(allNewsLeiXing, "IPFSs"),
	)

}
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
func TestMysqlPageData(t *testing.T) {
	db := daos.HuoQuLianJieChi()

	log.Println("== create table PageData ==", db)
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
	//======================================
	//创建 收益表
	log.Println("=== 开始创建收益表 ===")
	log.Println("== create table PageData ==", db)
	dataEarning := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.Earning,
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
				consts.BianMa:   consts.CurrentGrossRevenue,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "100",
				consts.MoRenZhi: "'星河飞天'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.NowTime,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "50",
				consts.MoRenZhi: "'IPFS-PageData'",
			},
			map[string]interface{}{
				consts.BianMa:   consts.UserIdEarning,
				consts.LeiXing:  consts.VARCHAR,
				consts.ChangDu:  "20",
				consts.MoRenZhi: "'星河飞天'",
			},
		},
	}
	daos.ChuangJianBiao(dataEarning)
	log.Println("=== 创建收益表完成 ===")


}
