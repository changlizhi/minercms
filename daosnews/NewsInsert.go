package daosnews

import (
	"minercms/consts"
	"minercms/daos"
	"minercms/utils"
	"time"
)

//新增新闻时数据库操作的顺序
//所有有关新闻新增的数据库操作都在这里集中编写

func InsertNews(canShu map[string]interface{}) map[string]interface{} {
	id := utils.HuoQuIdZiFu()
	biaoTi := utils.HuoQuZiFuZhi(canShu[consts.BiaoTi])
	leiXing := utils.HuoQuZiFuZhi(canShu[consts.LeiXing])
	jianJie := utils.HuoQuZiFuZhi(canShu[consts.JianJie])
	suoLueTu := utils.HuoQuZiFuZhi(canShu[consts.SuoLueTu])
	neiRong := utils.HuoQuZiFuZhi(canShu[consts.NeiRong])
	shiJian := utils.Time2StringNyrSfm(time.Now())
	gengXin := utils.Time2StringNyrSfm(time.Now())
	insertNeedCanShu := map[string]interface{}{
		consts.ShuJuKu:   consts.MINERCMS,
		consts.ShuJuBiao: consts.News,
		consts.ShuJuZhis: map[string]interface{}{
			consts.Id:       id,
			consts.BiaoTi:   biaoTi,
			consts.LeiXing:  leiXing,
			consts.JianJie:  jianJie,
			consts.SuoLueTu: suoLueTu,
			consts.NeiRong:  neiRong,
			consts.ShiJian:  shiJian,
			consts.GengXin:  gengXin,
		},
	}
	ret := daos.InsertData(insertNeedCanShu)
	return ret
}
