package service

import (
	"minercms/changliangs"
)

func FenFaYeWu(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	yeWuHaoInterface := canShu[changliangs.FuWuHao]
	if yeWuHaoInterface == nil || yeWuHaoInterface.(string) == "" {
		ret[changliangs.ZhuangTai] = changliangs.ShiBai
		ret[changliangs.ShuoMing] = "失败：未提供服务号"
		return ret
	}
	fuWuHao := yeWuHaoInterface.(string)
	if fuWuHao == changliangs.FW001 {
		return AddNewsService(canShu)
	} else if fuWuHao == changliangs.FW002 {
		return DelNewsService(canShu)
	} else if fuWuHao == changliangs.FW003 {
		return UpdateNewsService(canShu)
	} else if fuWuHao == changliangs.FW004 {
		return QueryNewsService(canShu)
	} else if fuWuHao == changliangs.FW005 {
		return GetOneNewsService(canShu)
	} else {
		ret[changliangs.ZhuangTai] = changliangs.ShiBai
		ret[changliangs.ShuoMing] = "失败：不支持的服务号"
		return ret
	}
}
