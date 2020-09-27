package service

import "minercms/consts"

func UpdateNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	ret[consts.ZhuangTai] = consts.ChengGong
	ret[consts.ShuoMing] = consts.ChengGongCN
	ret[consts.ShuJu] = []map[string]interface{}{
		map[string]interface{}{
			"FangWenShuoMing": "访问了Fw003",
		},
	}
	return ret
}
