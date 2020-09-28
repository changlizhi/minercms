package consts

//表名
const (
	News = "News"
)
const (
	Local         = "Local"
	MoRenShiJian  = "1970-01-01 10:00:00"
	Nyr           = "2006-01-02"
	NyrXhx        = "2006_01_02"
	NyrWu         = "20060102"
	NyrSfm        = "2006-01-02 15:04:05"
	NyrSfmXhx     = "2006_01_02_15_04_05"
	ZhengZeNyr    = "^[1-9]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$"
	ZhengZeSfm    = "^(20|21|22|23|[0-1]\\d):[0-5]\\d:[0-5]\\d$"
	ZhengZeNyrSfm = "^[1-9]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\\s+(20|21|22|23|[0-1]\\d):[0-5]\\d:[0-5]\\d$"
)

//字段名
const (
	FuWuHao   = "FuWuHao"
	ZhuangTai = "ZhuangTai"
	ShuoMing  = "ShuoMing"
	ShuJu     = "ShuJu"
	Id        = "Id"
	BiaoTi    = "BiaoTi"
	LeiXing   = "LeiXing"
	ShiJian   = "ShiJian"
	GengXin   = "GengXin"
	JianJie   = "JianJie"
	SuoLueTu  = "SuoLueTu"
	NeiRong   = "NeiRong"
)

//自定义字段，不在数据库中使用的
const (
	ShuJuZhis = "ShuJuZhis"
	ShuJuKu   = "ShuJuKu"
	ShuJuBiao = "ShuJuBiao"
	ZhuJian   = "ZhuJian"
	SuoYin    = "SuoYin"
	ZiDuans   = "ZiDuans"
	BianMa    = "BianMa"
	ChangDu   = "ChangDu"
	MoRenZhi  = "MoRenZhi"
)

//服务号
const (
	FW001 = "FW001"
	FW002 = "FW002"
	FW003 = "FW003"
	FW004 = "FW004"
	FW005 = "FW005"
)

//固定值
const (
	ShiBai      = "ShiBai"
	ChengGong   = "ChengGong"
	ShiBaiCN    = "失败"
	ChengGongCN = "成功"
	MINERCMS    = "minercms" //表名
)
