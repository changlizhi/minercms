package daosnews

import (
	"log"
	"minercms/consts"
	"minercms/daos"
	"minercms/utils"
)

func QueryNewsList(canShu map[string]interface{}) map[string]interface{} {
	//查询列表时不需要返回内容
	//参数进来之后分辨出从第几条开始查，查多少条然后拼接参数查询返回即可。
	//kaiShiShiJian := utils.HuoQuZiFuZhi(canShu[consts.KaiShiShiJian])
	//jieShuShiJian := utils.HuoQuZiFuZhi(canShu[consts.JieShuShiJian])
	dangQianYe := utils.HuoQuZiFuZhi(canShu[consts.DangQianYe])
	meiYeTiaoShu := utils.HuoQuZiFuZhi(canShu[consts.MeiYeTiaoShu])
	biaoTi := utils.HuoQuZiFuZhi(canShu[consts.BiaoTi])
	leiXing := utils.HuoQuZiFuZhi(canShu[consts.LeiXing])
	log.Println("标题等于空=",biaoTi=="")
	log.Println("leiXing题等于空=",leiXing=="")


	needCanShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.News,
		consts.ShuJuZhis: map[string]interface{}{
			consts.Id:       consts.Id,
			consts.BiaoTi:   consts.BiaoTi,
			consts.LeiXing:  consts.LeiXing,
			consts.JianJie:  consts.JianJie,
			consts.ShiJian:  consts.ShiJian,
			consts.GengXin:  consts.GengXin,
			consts.SuoLueTu: consts.SuoLueTu,
		},
		consts.TiaoJians: map[string]interface{}{
			//consts.KaiShiShiJian: kaiShiShiJian,
			//consts.JieShuShiJian: jieShuShiJian,
			consts.BiaoTi:       biaoTi,
			consts.LeiXing:      leiXing,
			consts.DangQianYe:   dangQianYe,
			consts.MeiYeTiaoShu: meiYeTiaoShu,
		},
	}
	log.Println("QueryNewsList---needCanShu",needCanShu)
	ret := daos.FindData(needCanShu)
	return ret
}

func QueryOneNews(canShu map[string]interface{}) map[string]interface{} {
	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
	needCanShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.News,
		consts.ShuJuZhis: map[string]interface{}{
			consts.Id:       consts.Id,
			consts.BiaoTi:   consts.BiaoTi,
			consts.LeiXing:  consts.LeiXing,
			consts.JianJie:  consts.JianJie,
			consts.ShiJian:  consts.ShiJian,
			consts.GengXin:  consts.GengXin,
			consts.SuoLueTu: consts.SuoLueTu,
			consts.NeiRong:  consts.NeiRong,
		},
		consts.TiaoJians: map[string]interface{}{
			consts.Id: id,
		},
	}
	ret := daos.FindData(needCanShu)
	return ret
}
func QueryOnePageDtaNews(canShu map[string]interface{}) map[string]interface{} {
	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
	needCanShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.PageData,
		consts.ShuJuZhis: map[string]interface{}{
			consts.Id:        consts.Id,
			consts.TitleName: consts.TitleName,
			consts.DataType:  consts.DataType,
			consts.Introduce: consts.Introduce,
			consts.TimeInfo:  consts.TimeInfo,
			consts.Thumbnail: consts.Thumbnail,
			consts.Details:   consts.Details,
		},
		consts.TiaoJians: map[string]interface{}{
			consts.Id: id,
		},
	}
	ret := daos.FindData(needCanShu)

	return ret
}

//

func PageDataQueryNewsList(canShu map[string]interface{}) map[string]interface{} {
	//查询列表时不需要返回内容
	//参数进来之后分辨出从第几条开始查，查多少条然后拼接参数查询返回即可。
	//kaiShiShiJian := utils.HuoQuZiFuZhi(canShu[consts.KaiShiShiJian])
	//jieShuShiJian := utils.HuoQuZiFuZhi(canShu[consts.JieShuShiJian])
	dangQianYe := utils.HuoQuZiFuZhi(canShu[consts.DangQianYe])
	meiYeTiaoShu := utils.HuoQuZiFuZhi(canShu[consts.MeiYeTiaoShu])
	biaoTi := utils.HuoQuZiFuZhi(canShu[consts.TitleName])
	leiXing := utils.HuoQuZiFuZhi(canShu[consts.DataType])
	needCanShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.PageData,
		consts.ShuJuZhis: map[string]interface{}{
			consts.Id:        consts.Id,
			consts.TitleName: consts.TitleName,
			consts.DataType:  consts.DataType,
			consts.Introduce: consts.Introduce,
			consts.TimeInfo:  consts.TimeInfo,
			consts.Thumbnail: consts.Thumbnail,
			consts.Details:   consts.Details,
		},
		consts.TiaoJians: map[string]interface{}{
			//consts.KaiShiShiJian: kaiShiShiJian,
			//consts.JieShuShiJian: jieShuShiJian,
			consts.BiaoTi:       biaoTi,
			consts.LeiXing:      leiXing,
			consts.DangQianYe:   dangQianYe,
			consts.MeiYeTiaoShu: meiYeTiaoShu,
		},
	}
	ret := daos.FindData(needCanShu)
	return ret
}
