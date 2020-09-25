package service

import "minercms/changliangs"

func DelNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	ret[changliangs.ZhuangTai] = changliangs.ChengGong
	ret["ShuoMing"] = changliangs.ChengGongCN
	ret["ShuJu"] = []map[string]interface{}{
		map[string]interface{}{
			"FangWenShuoMing": "访问了Fw002",
		},
	}
	return ret
}
