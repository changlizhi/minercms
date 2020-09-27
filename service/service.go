package service

import (
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
		return QueryNewsService(canShu)
	} else if fuWuHao == consts.FW005 {
		return GetOneNewsService(canShu)
	} else {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败：不支持的服务号"
		return ret
	}
}
