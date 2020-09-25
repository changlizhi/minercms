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
* 业务号: FW001
* 入参:

```json
{
  "YeWuHao":"FW001",//NOTNULL/业务号
  "BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
  "LeiXing":"IPFS",//NOTNULL/类型
  "ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
  "GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
  "JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
  "SuoLueTu":"http://ssss.png",
  "NeiRong":"xxx"
}

```

* 成功返回:

```json
{
  "ZhuangTai":"ChengGong",
  "ShuoMing":"成功",
  "ShuJu":[
    {
      "Id":"001",
      "BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
      "LeiXing":"IPFS",//NOTNULL/类型
      "ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
      "GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
      "JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
      "SuoLueTu":"http://ssss.png",
      "NeiRong":"xxx"
    }
    
  ],
}

```

* 失败返回:

```json

{
  ZhuangTai:"ShiBai",
  ShuoMing:"失败:xxx",
  ShuJu:[]
}

```

### 删除新闻

* uri:/delNews
* 说明：如果要删除一条新闻则删除包括新闻主题和新闻内容列表
* 业务号: FW002
* 入参：

```json

{
  "YeWuHao":"FW002",//NOTNULL/业务号
  "Id":"001"
}

```

* 成功返回:

```json
{
  "ZhuangTai":"ChengGong",
  "ShuoMing":"成功",
  "ShuJu":[
    {
      "ZhuangTai":"ChengGong",
      "ShuoMing":"成功",//如果有简单介绍说明则会放在这里
      "ShuJu":[]//后端必然返回这个数据，删除时如果有业务数据会在此字段提供，
    }
  ]
}
```

* 失败返回:

```json
{
  ZhuangTai:"ShiBai",
  ShuoMing:"失败:xxx",
  ShuJu:[]
}
```

### 修改新闻

* uri:/updateNews
* 说明:更新新闻，带着id即可
* 业务号: FW003
* 入参:

```json
{
  "YeWuHao":"FW003",//NOTNULL/业务号
  "Id":"001",
  "BiaoTi":"IPFSMain修改",//NOTNULL/标题
  "LeiXing":"咨询",//NOTNULL/类型
  "ShiJian":"2020-09-24 16:08:04",//NOTNEED/时间
  "GengXin":"2020-09-24 16:08:04",//NOTNEED/时间
  "JianJie":"北京时间Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
  "SuoLueTu":"http://ssss.png",
  "NeiRong":"xxx"
}

```

* 成功返回:

```json

{
  "ZhuangTai":"ChengGong",
  "ShuoMing":"成功",
  "ShuJu":[
    {
      "Id":"001",
      "BiaoTi":"IPFSMain修改",//NOTNULL/标题
      "LeiXing":"咨询",//NOTNULL/类型
      "ShiJian":"2020-09-24 16:08:04",//NOTNEED/时间
      "GengXin":"2020-09-24 16:08:04",//NOTNEED/时间
      "JianJie":"北京时间Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
      "SuoLueTu":"http://ssss.png",
      "NeiRong":"xxx"
    }
  ]
}


```

* 失败返回:

```json
{
  ZhuangTai:"ShiBai",
  ShuoMing:"失败:xxx",
  ShuJu:[]
}

```

### 查询新闻列表

* uri:/queryNews
* 说明：在右侧显示新闻列表
* 业务号: FW004
* 入参：

```json

{
  "YeWuHao":"FW004",//NOTNULL/业务号
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
  "ShuJu":[
        {
          "Id":"001",
          "BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
          "LeiXing":"IPFS",//NOTNULL/类型
          "ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
          "GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
          "JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
          "SuoLueTu":"http://ssss.png",
          "NeiRong":"xxx"
        },
        {
          "Id":"002",
          "BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
          "LeiXing":"IPFS",//NOTNULL/类型
          "ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
          "GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
          "JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
          "SuoLueTu":"http://ssss.png",
          "NeiRong":"xxx"
        }
   ]
}

```


* 失败返回:

```json
{
  ZhuangTai:"ShiBai",
  ShuoMing:"失败:xxx",
  ShuJu:[]
}
```
### 查询新闻明细

* uri:/getOneNews
* 说明：点选某个新闻时中间显示新闻内容
* 业务号: FW005
* 入参：

```json

{
  "YeWuHao":"FW005",//NOTNULL/业务号
  "NewsId":"001"//新闻主键
}

```

* 成功返回：
```json

{
  "ZhuangTai":"ChengGong",
  "ShuoMing":"成功",
  "ShuJu":[
        {
          "Id":"001",
          "BiaoTi":"IPFSMain全球奖励排名第一",//NOTNULL/标题
          "LeiXing":"IPFS",//NOTNULL/类型
          "ShiJian":"2020-09-23 16:08:04",//NOTNEED/时间
          "GengXin":"2020-09-23 16:08:04",//NOTNEED/时间
          "JianJie":"北京时间9月22日上午8点，Filecoin协议实验室举办了线上Space Race Celebration。在直播中，官方技术人员围绕存储生态、太空竞赛成果、Space Race 2进行分享。",//NOTNULL/简介
          "SuoLueTu":"http://ssss.png",
          "NeiRong":"xxx"
        }
  ]
}

```


* 失败返回:

```json
{
  ZhuangTai:"ShiBai",
  ShuoMing:"失败:xxx",
  ShuJu:[]
}
```



