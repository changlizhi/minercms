package service

import (
	"log"
	"minercms/changliangs"
)

func FenFaYeWu(canShu map[string]interface{}) map[string]interface{} {
	log.Println(canShu, "----canshu")
	ret := map[string]interface{}{}
	yeWuHaoInterface := canShu[changliangs.FuWuHao]
	if yeWuHaoInterface == nil || yeWuHaoInterface.(string) == "" {
		ret["ZhuangTai"] = "ShiBai"
		ret["ShuoMing"] = "失败：未提供服务号"
		return ret
	}
	fuWuHao := yeWuHaoInterface.(string)
	if fuWuHao == changliangs.FW001 {
		ret["ZhuangTai"] = "ChengGong"
		ret["ShuoMing"] = "成功"
		ret["ShuJu"] = []map[string]interface{}{
			map[string]interface{}{
				"FangWenShuoMing": "访问了Fw001",
			},
		}
	} else {
		ret = canShu
	}
	return ret

}
