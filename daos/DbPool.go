package daos

import (
	"database/sql"
	"log"
	"minercms/consts"
	"minercms/utils"
	"os"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	chuShiHuaChi() //初始化时用test作为初始化链接
}

var (
	suoShiLi sync.Once
	chi      *sql.DB
)

func chuShiHuaChi() {
	suoShiLi.Do(func() { //这里需要把已存在的都纳入进来，所以需要新建一个配置文件，这个配置文件用go写成
		chuangJianChi(consts.MINERCMS)
	})
}

func chuangJianChi(shuJuKuMing string) *sql.DB {
	lianJieChi, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/"+shuJuKuMing) //进入mysql数据库，然后
	if err != nil {
		log.Println("建立链接失败", err)
		os.Exit(1)
	}
	lianJieChi.SetMaxOpenConns(50)
	lianJieChi.SetMaxIdleConns(5)
	chi = lianJieChi
	log.Println("chuangJianChi:lianJieChi---", lianJieChi)
	return lianJieChi
}

func HuoQuLianJieChi() *sql.DB {
	return chi
}

func ChuangJianBiao(canShu map[string]interface{}) map[string]interface{} {
	//CREATE TABLE caoZuoKu.BiaoMing (
	// `ZhuJian` BIGINT(20) NOT NULL,
	// `MingCheng` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// `BianMa` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// `ZhuJianZhiBiao` VARCHAR(50) NOT NULL DEFAULT 'hfx',
	// PRIMARY KEY (`ZhuJian`)
	//  )
	//  ENGINE=InnoDB
	// 这里要求传入的ShuJu第一个必须是表名字符串，
	// 第二个必须是主键名指定
	// 第三个必须是字段列表
	// 后续强化这个校验
	// 要重点命名清楚每个中间变量的意义，ShuJuKu是需要访问的数据库，BiaoMing是需要访问的数据库下的表，字段中的ShuJuKu是值

	caoZuoKu := utils.HuoQuZiFuZhi(canShu[consts.ShuJuKu])
	caoZuoBiao := utils.HuoQuZiFuZhi(canShu[consts.ShuJuBiao])
	zhuJian := utils.HuoQuZiFuZhi(canShu[consts.ZhuJian])
	suoYin := utils.HuoQuZiFuZhi(canShu[consts.SuoYin])
	ziDuans := canShu[consts.ZiDuans].([]interface{}) //把字段拿出来

	builder := strings.Builder{}

	builder.WriteString("CREATE TABLE ")
	builder.WriteString(caoZuoKu)
	builder.WriteString(".")
	builder.WriteString(caoZuoBiao)
	builder.WriteString(" (")
	for _, interv := range ziDuans {
		//`MingCheng` VARCHAR(50) NOT NULL DEFAULT 'hfx',
		v := interv.(map[string]interface{})
		builder.WriteString(utils.HuoQuZiFuZhi(v[consts.BianMa]))
		builder.WriteString(" ")
		ziDuanLeiXing := utils.HuoQuZiFuZhi(v[consts.LeiXing])
		builder.WriteString(ziDuanLeiXing)
		if ziDuanLeiXing != consts.TINYTEXT &&
			ziDuanLeiXing != consts.TEXT &&
			ziDuanLeiXing != consts.MEDIUMTEXT &&
			ziDuanLeiXing != consts.FULLTEXT &&
			ziDuanLeiXing != consts.LONGTEXT {

			builder.WriteString("(")
			builder.WriteString(utils.HuoQuZiFuZhi(v[consts.ChangDu]))
			builder.WriteString(") ")
			builder.WriteString("NOT NULL DEFAULT ")
			builder.WriteString(utils.HuoQuZiFuZhi(v[consts.MoRenZhi]))
		}
		builder.WriteString(",")
	}

	builder.WriteString("PRIMARY KEY (")
	builder.WriteString(zhuJian)
	builder.WriteString(")")
	if suoYin != "" {
		builder.WriteString(",UNIQUE INDEX ")
		builder.WriteString(suoYin)
		builder.WriteString(" (")
		builder.WriteString(suoYin)
		builder.WriteString(")")
	}
	builder.WriteString(")COLLATE='utf8mb4_general_ci' ENGINE=InnoDB")

	sqlStr := builder.String()

	result, err := chi.Exec(sqlStr)
	if err!=nil{
		log.Println("打印 sqlStr ，result :  ", sqlStr, result)

		log.Println("创建 数据库 表 失败 ，原因是 : ", sqlStr, result, err)

		return canShu
	}

	log.Println("创建 数据库 表  成功  ", sqlStr, result, err)
	return canShu
}
