package daos

import (
	"database/sql"
	"log"
	"minercms/consts"
	"minercms/utils"
	"strconv"
	"strings"
)

func FindData(canShu map[string]interface{}) map[string]interface{} {
	//select ?,?,? from kus.biaos where key1=? and key2=?
	caoZuoKu := canShu[consts.ShuJuKu].(string)
	caoZuoBiao := canShu[consts.ShuJuBiao].(string)
	caoZuoZhis := canShu[consts.ShuJuZhis].(map[string]interface{}) //把需要查询的字段拿出来
	tiaoJians := canShu[consts.TiaoJians].(map[string]interface{})  //把条件拿出来
	dangQianYe := ""
	meiYeTiaoShu := ""

	wenHaos := []string{}
	colNames := []string{}
	values := []interface{}{}
	//默认ZhuJian全都通过snowflakerid得到，但是业务表中主键是从主键表来的。
	for k, v := range caoZuoZhis {
		colNames = append(colNames, k)
		wenHaos = append(wenHaos, k)
		values = append(values, v)
	}
	tkeys := []string{}
	tvalues := []interface{}{}
	for k, v := range tiaoJians {

		if k != consts.DangQianYe && k != consts.MeiYeTiaoShu && v!=""{
			tkeys = append(tkeys, k+" = ? ")
			tvalues = append(tvalues, v)
		}
		if k == consts.DangQianYe{
			dangQianYeStr:=utils.HuoQuZiFuZhi(v)
			dangQianYeInt ,_:= strconv.Atoi(dangQianYeStr)
			tempDangQianYe:=dangQianYeInt-1
			dangQianYe=strconv.Itoa(tempDangQianYe)

		}
		if k==consts.MeiYeTiaoShu{
			meiYeTiaoShu=utils.HuoQuZiFuZhi(v)
		}
	}

	builder := strings.Builder{}
	builder.WriteString(" SELECT ")
	builder.WriteString(strings.Join(wenHaos, " , "))
	builder.WriteString(" FROM ")

	builder.WriteString(caoZuoKu)
	builder.WriteString(".")
	builder.WriteString(caoZuoBiao)


	if len(tvalues) > 0{

		builder.WriteString(" WHERE ")
		builder.WriteString(strings.Join(tkeys, " AND "))
	}

	if dangQianYe != "" && meiYeTiaoShu != "" {
		builder.WriteString(" LIMIT ")
		builder.WriteString(dangQianYe)
		builder.WriteString(" , ")
		builder.WriteString(meiYeTiaoShu)

	} else {
		builder.WriteString(" LIMIT 0,10 ")

	}

	sqlStr := builder.String()
	db := HuoQuLianJieChi()
	allValues := tvalues
	result, err := db.Query(sqlStr, allValues...)
	ret := map[string]interface{}{}
	if err != nil{
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败：" + err.Error()
		return ret
	}
	if result == nil{
		ret[consts.ZhuangTai] = consts.ShiBai
		ret[consts.ShuoMing] = "失败：result为空"
		return ret
	}
	retShuJu := scanRet(colNames, result)
	//如果err不为空则把err封装丢出去，如果为空则把数据原样返回
	ret[consts.ZhuangTai] = consts.ChengGong
	ret[consts.ShuoMing] = consts.ChengGongCN
	ret[consts.ShuJu] = retShuJu

	log.Println("XiuGai:sqlStr,allvalues,result,err,ret---", sqlStr, allValues, result, err, ret)
	return ret
}

func scanRet(lies []string, rows *sql.Rows) []map[string]interface{} {
	//必须保证查到的列顺序和传入的列顺序一致，否则将会匹配失败
	dbCols, err := rows.Columns()
	if err != nil {
		log.Println("查询数据库失败，请注意检查！")
	}
	for i, _ := range dbCols {
		if dbCols[i] != lies[i] {
			log.Println("数据库返回的字段顺序和查询时字段顺序不一致，请注意检查", dbCols[i], lies[i])
		}
	}
	//保证了顺序之后再Scan就没问题了
	ret := []map[string]interface{}{}
	for rows.Next() { //每一行
		tempLie := make([]string, len(lies))
		tempLieRef := make([]interface{}, len(lies))

		for i := 0; i < len(lies); i++ {
			tempLieRef[i] = &tempLie[i]
		}
		rows.Scan(
			tempLieRef...,
		)
		retOne := map[string]interface{}{}
		for xiaBiao, lie := range lies {
			if tempLieRef[xiaBiao] != nil {
				retOne[lie] = *tempLieRef[xiaBiao].(*string)
			}
		}
		ret = append(ret, retOne)
	}
	return ret
}
