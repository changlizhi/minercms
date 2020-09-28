package daos

import (
	"log"
	"minercms/consts"
	"minercms/utils"
	"strings"
)

func FindData(canShu map[string]interface{}) map[string]interface{} {
	//select ?,?,? from kus.biaos where key1=? and key2=?
	caoZuoKu := canShu[consts.ShuJuKu].(string)
	caoZuoBiao := canShu[consts.ShuJuBiao].(string)
	caoZuoZhis := canShu[consts.ShuJuZhis].(map[string]interface{}) //把需要查询的字段拿出来
	tiaoJians := canShu[consts.TiaoJians].(map[string]interface{})  //把条件拿出来
	dangQianYe := utils.HuoQuZiFuZhi(canShu[consts.DangQianYe])
	meiYeTiaoShu := utils.HuoQuZiFuZhi(canShu[consts.MeiYeTiaoShu])

	needCols := []string{}
	values := []interface{}{}
	//默认ZhuJian全都通过snowflakerid得到，但是业务表中主键是从主键表来的。
	for _, v := range caoZuoZhis {
		needCols = append(needCols, " ? ")
		values = append(values, v)
	}
	tkeys := []string{}
	tvalues := []interface{}{}
	for k, v := range tiaoJians {
		tkeys = append(tkeys, k+" = ? ")
		tvalues = append(tvalues, v)
	}

	builder := strings.Builder{}
	builder.WriteString(" SELECT ")
	builder.WriteString(strings.Join(needCols, " , "))
	builder.WriteString(" FROM ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(".")
	builder.WriteString(caoZuoBiao)

	builder.WriteString(" WHERE ")
	builder.WriteString(strings.Join(tkeys, " AND "))
	if dangQianYe != "" && meiYeTiaoShu != "" {
		builder.WriteString(" LIMIT( ")
		builder.WriteString(dangQianYe)
		builder.WriteString(" , ")
		builder.WriteString(meiYeTiaoShu)
		builder.WriteString(" ) ")

	} else {
		builder.WriteString(" LIMIT (0,10)")

	}

	sqlStr := builder.String()
	db := HuoQuLianJieChi()
	allValues := append(values, tvalues...)
	result, err := db.Query(sqlStr, allValues...)
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

	log.Println("XiuGai:sqlStr,allvalues,result,err,ret---", sqlStr, allValues, result, err, ret)
	return ret
}
