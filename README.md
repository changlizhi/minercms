# 公司新闻发布及展示

**说明**:

* 表格中Header里中文为说明，字段名用拼音全拼
* 表格中Header字段不注明则该字段前端必须传即使没有值，
* 若标记为NOTNEED////////为后端自动生成，不需要前端传入该字段，
* 必填以NOTNULL表示
* 数据值中Arr开头的数据说名是一个列表，在json中表示为一个[]用字段名的方式获取Arr是说明子表的名称的

参考网页:https://www.ipfsmain.cn/news/952

## 数据接口

* BaseUrl:http://192.168.1.xxx:8888/cms
* method:POST

### 添加新闻

**标题格式参考**:【太空竞赛】官宣:IPFSMain全球奖励排名第一

* uri:/addNews
* 说明:添加新闻标题，只是作为一个新闻主题存在，可以在右侧列表中展示，提供简介内容

* 入参:

```json
{
	"BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
	"BiaoQian":"太空竞赛",//NOTNULL/标签
	"LeiXing":"官宣",//NOTNULL/类型
	"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
	"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
	"JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。"//NOTNULL/简介
}

```

* 成功返回:

```json

{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",
	"ShuJu":[{
		"BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
		"BiaoQian":"太空竞赛",//NOTNULL/标签
		"LeiXing":"官宣",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
		"NeiRongs":[]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	}]
}

```

* 失败返回:

```json

{
	ZhuangTai:"ShiBai",
	ShuoMing:"失败:xxx"
}

```

### 添加新闻内容

* uri:/addNewsContent
* 说明:添加新闻内容，根据传入的顺序跟新闻主题时多对一的方式，查询时会根据录入顺序以子array的方式给前端，前端根据此列表里定义的顺序进行循环展示内容，当然也可以一段传入富文本内容
* 入参:

```json

{
	"Id":"001",//NOTNEED/主键
	"NewsId":"001",//NOTNULL/新闻主键
	"ShunXu":"1",//NOTNEED/顺序
	"XiaoBiaoTi":"生态现状",//小标题
	"WenZiNeiRong":"根据介绍，目前已。。。",//文字内容
	"TuPianDiZhi":"https://image.ipfsmain.cn/news/202009/23/87383e3d9a80983bdd2545dbf444889b.png",//图片地址
}

```

* 成功返回:返回后数据可以直接预览

```json

{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",
	"ShuJu":[{
		"BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
		"BiaoQian":"太空竞赛",//NOTNULL/标签
		"LeiXing":"官宣",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
		"NeiRongs":[
			{
				"Id":"001",
				"NewsId":"001",//NOTNULL/新闻主键
				"ShunXu":"1",//NOTNEED/////////顺序
				"XiaoBiaoTi":"生态现状",//小标题
				"WenZiNeiRong":"根据介绍，目前已。。。",//文字内容
				"TuPianDiZhi":"https://image.ipfsmain.cn/news/202009/23/87383e3d9a80983bdd2545dbf444889b.png",//图片地址
			}
		]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	}]
}

```

### 删除新闻

* uri:/delNews
* 说明：如果要删除一条新闻则删除包括新闻主题和新闻内容列表
* 入参：

```json

{
	"Id":"001"
}

```

* 成功返回:

```json
{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",//如果有简单介绍说明则会放在这里
	"ShuJu":[]//后端必然返回这个数据，删除时如果有业务数据会在此字段提供，
}
```

* 失败返回:

```json
{
	"ZhuangTai":"ShiBai",
	"ShuoMing":"失败:xxx原因",//如果有简单介绍说明则会放在这里
}
```

### 修改新闻

* uri:/updateNews
* 说明：修改新闻则修改的时新闻新建时那些字段进行修改
* 入参：

```json
{
	"Id":"001",
	"BiaoTi":"IPFSMain全球奖励排名第一2",//NOTNULL/标题
	"BiaoQian":"太空竞赛2",//NOTNULL/标签
	"LeiXing":"官宣2",//NOTNULL/类型
	"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
	"GengXin":"2020-09-24 14:18:04",//NOTNEED/时间
	"JianJie":"北京时间9月22日上午8点，太空竞赛成果、Space Race 2进行分享。"//NOTNULL/简介
}

```

* 成功返回：

```json

{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",
	"ShuJu":[{
		"Id":"001",
		"BiaoTi":"IPFSMain全球奖励排名第一2",//NOTNULL/标题
		"BiaoQian":"太空竞赛2",//NOTNULL/标签
		"LeiXing":"官宣2",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-24 14:18:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，太空竞赛成果、Space Race 2进行分享。"//NOTNULL/简介
		"NeiRongs":[
			{
				"Id":"001",
				"NewsId":"001",//NOTNULL/新闻主键
				"ShunXu":"1",//NOTNEED/顺序
				"XiaoBiaoTi":"生态现状",//小标题
				"WenZiNeiRong":"根据介绍，目前已。。。",//文字内容
				"TuPianDiZhi":"https://image.ipfsmain.cn/news/202009/23/87383e3d9a80983bdd2545dbf444889b.png",//图片地址
			}
		]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	}]
}

```

### 修改新闻内容

* uri:/updateNewsContent
* 说明：修改新闻内容只是修改某一段内容，前端要控制可以点选某一段进行修改，当然也可以一段传入富文本内容
* 入参：

```json

{
	"Id":"001",//NOTNEED/主键
	"NewsId":"001",//NOTNULL/新闻主键
	"ShunXu":"1",//NOTNEED/顺序
	"XiaoBiaoTi":"生态现状2",//小标题
	"WenZiNeiRong":"根据介绍，目前已2",//文字内容
	"TuPianDiZhi":"https://image.ipfxxx.png",//图片地址
}

```


* 成功返回:返回后数据可以直接预览

```json

{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",
	"ShuJu":[{
		"BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
		"BiaoQian":"太空竞赛",//NOTNULL/标签
		"LeiXing":"官宣",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
		"NeiRongs":[
			{
				"Id":"001",//NOTNEED/主键
				"NewsId":"001",//NOTNULL/新闻主键
				"ShunXu":"1",//NOTNEED/顺序
				"XiaoBiaoTi":"生态现状2",//小标题
				"WenZiNeiRong":"根据介绍，目前已2",//文字内容
				"TuPianDiZhi":"https://image.ipfxxx.png",//图片地址
			}
		]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	}]
}

```

* 失败返回:

```json
{
	"ZhuangTai":"ShiBai",
	"ShuoMing":"失败:xxx原因，请刷新页面后重试",//如果有简单介绍说明则会放在这里
}
```

### 查询新闻列表

* uri:/queryNewsList
* 说明：在右侧显示新闻列表
* 入参：

```json

{
	"DangQianYe":"1",//当前页
	"MeiYeTiaoShu":"10",//每页条数
	"KaiShiShiJian":"2020-09-20",//开始时间
	"JieShuShiJian":"2020-09-24"//结束时间
}

```


* 成功返回:

```json

{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",
	"ShuJu":[{
		"Id":"001",
		"BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
		"BiaoQian":"太空竞赛",//NOTNULL/标签
		"LeiXing":"官宣",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
		"NeiRongs":[]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	},{
		"Id":"002",
		"BiaoTi":"IPFSMain全球奖励排名第一2",//NOTNULL/标题
		"BiaoQian":"太空竞赛2",//NOTNULL/标签
		"LeiXing":"官宣2",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
		"JianJie":"北京时间922月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
		"NeiRongs":[]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	},{
		"Id":"003",
		"BiaoTi":"IPFSMain全球奖励排名第一3",//NOTNULL/标题
		"BiaoQian":"太空竞赛3",//NOTNULL/标签
		"LeiXing":"官宣3",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
		"NeiRongs":[]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	}]
}

```


* 失败返回:

```json
{
	"ZhuangTai":"ShiBai",
	"ShuoMing":"失败:xxx原因，请刷新页面后重试",//如果有简单介绍说明则会放在这里
}
```

### 查询新闻明细

* uri:/getOneNews
* 说明：点选某个新闻时中间显示新闻内容
* 入参：

```json

{
	"NewsId":"001"//新闻主键
}

```

* 成功返回：
```json

{
	"ZhuangTai":"ChengGong",
	"ShuoMing":"成功",
	"ShuJu":[{
		"Id":"001",
		"BiaoTi":"IPFSMain全球奖励排名第一2",//NOTNULL/标题
		"BiaoQian":"太空竞赛2",//NOTNULL/标签
		"LeiXing":"官宣2",//NOTNULL/类型
		"ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
		"GengXin":"2020-09-24 14:18:04",//NOTNEED/时间
		"JianJie":"北京时间9月22日上午8点，太空竞赛成果、Space Race 2进行分享。"//NOTNULL/简介
		"NeiRongs":[
			{
				"Id":"001",
				"NewsId":"001",//NOTNULL/新闻主键
				"ShunXu":"1",//NOTNEED/顺序
				"XiaoBiaoTi":"生态现状",//小标题
				"WenZiNeiRong":"根据介绍，目前已。。。",//文字内容
				"TuPianDiZhi":"https://image.ipfsmain.cn/news/202009/23/87383e3d9a80983bdd2545dbf444889b.png",//图片地址
			}
		]//刚创建时没有NeiRongs但是后端会返回一个占位字段，后端必然返回，前端直接判断lenth即可。
	}]
}

```


* 失败返回:

```json
{
	"ZhuangTai":"ShiBai",
	"ShuoMing":"失败:xxx原因，请刷新页面后重试",//如果有简单介绍说明则会放在这里
}
```

