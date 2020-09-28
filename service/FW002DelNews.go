package service

import (
	"log"
	"minercms/consts"
	"minercms/daosnews"
	"minercms/utils"
)

func JiaoYanFw002(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	id := utils.HuoQuZiFuZhi(canShu[consts.Id])
	log.Println("id-----", id)
	if utils.ShiFouKongZiFu(id) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:Id主键不能为空"
		return ret
	}

	return ret
}
func DelNewsService(canShu map[string]interface{}) map[string]interface{} {
	retJiaoYan := JiaoYanFw002(canShu)
	if retJiaoYan[consts.ZhuangTai] == consts.ShiBai {
		return retJiaoYan
	}
	//可以设置表的标题字段为唯一索引就不需要在插入之前校验了
	ret := daosnews.DeleteNews(canShu)
	return ret
}
