package service

import (
	"log"
	"minercms/consts"
)

func FenFaYeWu(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	yeWuHaoInterface := canShu[consts.FuWuHao]
	if yeWuHaoInterface == nil || yeWuHaoInterface.(string) == "" {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败：未提供服务号"
		return ret
	}
	fuWuHao := yeWuHaoInterface.(string)
	if fuWuHao == consts.FW001 {
		return AddNewsService(canShu)
	} else if fuWuHao == consts.FW002 {
		return DelNewsService(canShu)
	} else if fuWuHao == consts.FW003 {
		return UpdateNewsService(canShu)
	} else if fuWuHao == consts.FW004 {
		log.Println("FenFaYeWu---canShu[consts.DangQianYe",canShu[consts.DangQianYe])
		return QueryNewsService(canShu)
	} else if fuWuHao == consts.FW005 {
		return GetOneNewsService(canShu)

	} else if fuWuHao == consts.FW006 {
		return GetPageDataNewsService(canShu)
	} else if fuWuHao == consts.FW007 {
		return GetPageDataQueryNewsService(canShu)
	} else {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败：不支持的服务号"
		return ret
	}
}
