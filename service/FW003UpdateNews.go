package service

import (
	"minercms/consts"
	"minercms/daosnews"
	"minercms/utils"
)

func JiaoYanFw003(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
	if utils.ShiFouKongZiFu(id) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:Id主键不能为空"
		return ret
	}

	return ret
}
func UpdateNewsService(canShu map[string]interface{}) map[string]interface{} {
	retJiaoYan := JiaoYanFw003(canShu)
	if retJiaoYan[consts.ZhuangTai] == consts.ShiBai {
		return retJiaoYan
	}
	ret := daosnews.UpdateNews(canShu)
	return ret
}
