package daos

import (
	"log"
	"minercms/consts"
	"strings"
)

func InsertData(canShu map[string]interface{}) map[string]interface{} {
	//insert into caoZuoKu.caoZuoBiao (columns) values(values)
	caoZuoKu := canShu[consts.ShuJuKu].(string)
	caoZuoBiao := canShu[consts.ShuJuBiao].(string)
	caoZuoZhis := canShu[consts.ShuJuZhis].(map[string]interface{}) //把字段拿出来

	keys := []string{}
	wenHaos := []string{}
	values := []interface{}{}
	//默认ZhuJian全都通过snowflakerid得到，但是业务表中主键是从主键表来的。
	for k, v := range caoZuoZhis {
		keys = append(keys, k)
		values = append(values, v)
		wenHaos = append(wenHaos, "?")
	}

	builder := strings.Builder{}
	builder.WriteString(" INSERT INTO ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(".")
	builder.WriteString(caoZuoBiao)

	builder.WriteString(" ( ")
	builder.WriteString(strings.Join(keys, ","))
	builder.WriteString(" )values( ")
	builder.WriteString(strings.Join(wenHaos, ","))
	builder.WriteString(" ) ")

	sqlStr := builder.String()
	db := HuoQuLianJieChi()
	result, err := db.Exec(sqlStr, values...)
	//如果err不为空则把err封装丢出去，如果为空则把数据原样返回
	ret := map[string]interface{}{}
	if err != nil {
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败：" + err.Error()

	} else {
		ret[consts.ZhuangTai] = consts.ChengGong
		ret[consts.ShuoMing] = consts.ChengGongCN
		ret[consts.ShuJu] = []map[string]interface{}{
			caoZuoZhis,
		}
	}

	log.Println("XinZeng:sqlStr,values,result,err,ret---", sqlStr, values, result, err, ret)
	return ret
}
