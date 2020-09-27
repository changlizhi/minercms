package service

import "minercms/consts"

func DelNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	ret[consts.ZhuangTai] = consts.ChengGong
	ret[consts.ShuoMing] = consts.ChengGongCN
	ret[consts.ShuJu] = []map[string]interface{}{
		map[string]interface{}{
			"FangWenShuoMing": "访问了Fw002",
		},
	}
	return ret
}
