package service

import (
	"log"
	"minercms/consts"
	"minercms/daosnews"
	"minercms/utils"
)

var allNewsLeiXing = []string{
	"IPFS",
	"HangYeDongTai",
	"GongSiFaBu",
}

//先校验必填项，数据格式，长度
//再查询是否已经有相同的标题，如果有则返回提示已经有相同的标题，请重新输入
//通过这两个校验之后直接入库，如果入库没报错则返回入库成功后的数据
//id和shijian由此时生成
func JiaoYanFw001(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	biaoTi := utils.HuoQuZiFuZhi(canShu[consts.BiaoTi])
	leiXing := utils.HuoQuZiFuZhi(canShu[consts.LeiXing])
	log.Println("leiXing---", leiXing)
	if utils.ShiFouKongZiFu(biaoTi) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:标题不能为空"
		return ret
	}
	if utils.ShiFouKongZiFu(leiXing) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:类型不能为空"
		return ret
	}
	if utils.ShiFouChaoChang(biaoTi, 100) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:标题不能为空"
		return ret
	}
	if utils.ShiFouChaoChang(leiXing, 50) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:类型不能为空"
		return ret
	}
	if !utils.ShiFouBaoHan(allNewsLeiXing, leiXing) {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败:LeiXing不在系统定义范围"
		return ret
	}

	return ret
}

func AddNewsService(canShu map[string]interface{}) map[string]interface{} {
	retJiaoYan := JiaoYanFw001(canShu)
	if retJiaoYan[consts.ZhuangTai] == consts.ShiBai {
		return retJiaoYan
	}
	//可以设置表的标题字段为唯一索引就不需要在插入之前校验了
	ret := daosnews.InsertNews(canShu)
	return ret
}
