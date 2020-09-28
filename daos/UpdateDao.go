package daos

import (
	"log"
	"minercms/consts"
	"strings"
)

func UpdateData(canShu map[string]interface{}) map[string]interface{} {
	//update ku.biao set key1=?,key2=?,key3=? where col1=?,col2=?
	caoZuoKu := canShu[consts.ShuJuKu].(string)
	caoZuoBiao := canShu[consts.ShuJuBiao].(string)
	caoZuoZhis := canShu[consts.ShuJuZhis].(map[string]interface{}) //把字段拿出来
	tiaoJians := canShu[consts.TiaoJians].(map[string]interface{})  //把条件拿出来

	needSets := []string{}
	values := []interface{}{}
	//默认ZhuJian全都通过snowflakerid得到，但是业务表中主键是从主键表来的。
	for k, v := range caoZuoZhis {
		needSets = append(needSets, k+" = ? ")
		values = append(values, v)
	}
	tkeys := []string{}
	tvalues := []interface{}{}
	for k, v := range tiaoJians {
		tkeys = append(tkeys, k+" = ? ")
		tvalues = append(tvalues, v)
	}

	builder := strings.Builder{}
	builder.WriteString(" UPDATE ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(".")
	builder.WriteString(caoZuoBiao)

	builder.WriteString(" SET ")
	builder.WriteString(strings.Join(needSets, ","))
	builder.WriteString(" WHERE ")
	builder.WriteString(strings.Join(tkeys, ","))

	sqlStr := builder.String()
	db := HuoQuLianJieChi()
	allValues := append(values, tvalues...)
	result, err := db.Exec(sqlStr, allValues...)
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

	log.Println("XiuGai:sqlStr,values,result,err,ret---", sqlStr, values, result, err, ret)
	return ret
}
