package daos

import (
	"log"
	"minercms/consts"
	"strings"
)

func DeleteData(canShu map[string]interface{}) map[string]interface{} {
	//delete from ku.biao where 1=1 and key1=? and key2=?
	caoZuoKu := canShu[consts.ShuJuKu].(string)
	caoZuoBiao := canShu[consts.ShuJuBiao].(string)
	tiaoJians := canShu[consts.TiaoJians].(map[string]interface{}) //把字段拿出来

	wheres := []string{}
	values := []interface{}{}
	//默认ZhuJian全都通过snowflakerid得到，但是业务表中主键是从主键表来的。
	for k, v := range tiaoJians {
		wheres = append(wheres, k+" = ?")
		values = append(values, v)
	}

	builder := strings.Builder{}
	builder.WriteString(" DELETE FROM ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(".")
	builder.WriteString(caoZuoBiao)

	builder.WriteString(" WHERE ")
	builder.WriteString(strings.Join(wheres, " and "))

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
			tiaoJians,
		}
	}

	log.Println("ShanChu:sqlStr,values,result,err,ret---", sqlStr, values, result, err, ret)
	return ret
}
