# API接入 {docsify-ignore-all}

## 嗅探版本API {docsify-ignore}

`GET https://xx.wps.cn/api/v1/project/:pid/ver`


### 接口说明

1. 客户端定时轮询该嗅探接口来判断项目版本是否已更新
2. 如果接口返回的版本号与上一次客户端维护的版本号`一致`，则表示配置`未`更新，此时进入下一轮嗅探。
3. 如果接口返回的版本号与上一次客户端维护的版本号`不一致`，则表示配置`已`更新，此时可进行项目/模块组/模块级别的配置拉取。

### 请求参数 Path Params参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| pid  | Int |   项目ID   | 是|

### 返回示例
```json
{
    "ver": 2,
    "md5": "5e7f985cf1e282657e1a8c85"
}
```
#### 响应数据说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |  服务端项目版本号    |
| md5  | String |   服务端项目数据MD5   |

## 获取项目级别的配置API

`POST https://xx.wps.cn/api/v1/project/:pid/config`


### 接口说明

?> 客户端通过该接口可获取到项目级别的配置数据，如返回“在线参数项目”、“应用页项目”等所有的配置数据。项目下可有多个模块组，模块组下可有多个模块。

### 请求参数 Path Params参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| pid  | Int |   项目id   | 是|


### 请求参数 Body参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| md5  | string |    客户端上次请求服务端返回的md5，第一次请求传空  | 是|
| ver  | int |   客户端上次请求服务端返回的版本号，第一次请求传空  | 是|
| filters  | map[string]string |  客户端请求过滤参数，服务端根据该过滤条件返回过滤后的ids | 否|


### 请求示例

?> 请求项目ID为1的配置数据，服务端判断客户端上次拉取的项目版本为10，如果服务端的该项目版本已发布到11，客户端则会拉取到版本为11的数据。

```json
https://xx.wps.cn/api/v1/project/1/config
{
  "md5": "69e2d0410fe805dfeaef461220d43a52",
  "ver": 10,
  "filters": {
    "app_version": "13.0",
	"deviceid": "11111",
	"platform": "1",
  }
}
```
### 返回示例

#### 返回示例一：全量动态

```json
{ 
	"ver": 11,                                                                   
	"md5": "5e5781c06253f12682cd5009",    
	"ret_type": 0, 
	"ids": [13],                                                   
	"data": {
		"full": [{                                                   
			"mg_id": 1,                                                
			"mg_name": "稍后阅读",                           
			"sub": [],                                                  
			"modules": [{                                              
				"id": 13,                                                        
				"name": "屏蔽OEM渠道",
				"expire_time" : 1590748456,
				"effective_time" : 1590748456,
				"value": {                                               
					"timeInterval": 5
				}
			}]
		}]
	}

}
```

####  返回参数说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |   当前拉取配置的项目版本号  |
| md5  | string |  当前拉取配置的项目MD5 |
| ret_type  | int | 拉取策略, 0:动态全量, 1: 动态diff, 2:cdn全量, 3:cdn diff |
| ids  | array | 根据客户端参数过滤后的ID列表，客户端根据此ID列表解析下面的full全量数据 |
| data.full  | array | 该项目下的所有模块组的配置 |
| full.mg_id  | int | 模块组ID |
| full.mg_name  | string | 模块组名称 |
| full.sub  | array | 子模块组，如果该模块组有配置子模块组，则子模块组的具体配置在该层级 |
| full.modules  | array | 该模块组下的具体模块配置  |
| modules.id  | int | 模块ID  |
| modules.name  | int | 模块名称  |
| modules.value  | object | 客户端需要的具体json参数，可自定义配置  |


#### 返回示例二：全量CDN
```json
{
    "ret_type": 1,
    "ver": 11,
    "ids": [3,2],
    "md5": "75b5bf51913174aacb786c662c39126b",
	"data":{
	    "url": "fpccomb.wpscdn.cn/api/v1/file/project-all-1-10-75b5bf51913174aacb786c662c39126b"
	}
}

```

#### 返回参数说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |   当前拉取配置的项目版本号  |
| md5  | string |  当前拉取配置的项目MD5 |
| ret_type  | int | 拉取策略, 0:动态全量, 1: 动态diff, 2:cdn全量, 3:cdn diff |
| ids  | array | 根据客户端参数过滤后的ID列表，客户端根据此ID列表解析下面的full全量数据 |
| data.url  | string | 全量数据的cdn下载地址，请求该地址可直接下载全量数据 |

#### 返回示例三：DIFF动态
```json
{
    "ret_type": 2,
    "ver": 11,
	"ids": [3,2],
	"md5": "6522058b28147d5d871fc44cec93b158",
    "data":{
	    "diff": `@@ -148,9 +148,9 @@\n id%22:\n-3\n+2\n ,%22ta\n@@ -163,16 +163,11 @@\n %87%E7%AD%BE\n-%E4%BA%8C%E6%96%B011\n+%E4%B8%80\n %22,%22e\n@@ -195,25 +195,25 @@\n null%7D,%7B%22id%22:\n-2\n+3\n ,%22tag%22:%22%E6%A0%87%E7\n@@ -206,35 +206,40 @@\n :3,%22tag%22:%22%E6%A0%87%E7%AD%BE\n-%E4%B8%80\n+%E4%BA%8C%E6%96%B011\n %22,%22enabled%22:%22on%22\n`
	}
}
```

#### 返回参数说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |   当前拉取配置的项目版本号  |
| md5  | string |  当前拉取配置的项目MD5 |
| ret_type  | int | 拉取策略, 0:动态全量, 1: 动态diff, 2:cdn全量, 3:cdn diff |
| ids  | array | 根据客户端参数过滤后的ID列表，客户端根据此ID列表解析下面的full全量数据 |
| data.diff  | string | diff数据 |

#### 返回示例三：DIFF CDN
```json
{
    "ret_type": 3,
    "ver": 11,
    "ids": [3,2],
    "md5": "ad6dbd5e3cd2c3fb49c8034381040754",
    "data":{
	    "url": "fpccomb.wpscdn.cn/api/v1/file/project-result-1-10-ad6dbd5e3cd2c3fb49c8034381040754",
        "diff_url": "fpccomb.wpscdn.cn/api/v1/file/project-result-1-10-ad6dbd5e3cd2c3fb49c8034381040754"
	}
}
```

#### 返回参数说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |   当前拉取配置的项目版本号  |
| md5  | string |  当前拉取配置的项目MD5 |
| ret_type  | int | 拉取策略, 0:动态全量, 1: 动态diff, 2:cdn全量, 3:cdn diff |
| ids  | array | 根据客户端参数过滤后的ID列表，客户端根据此ID列表解析下面的full全量数据 |
| data.url  | string | 全量数据的cdn下载地址，请求该地址可直接下载全量数据 |
| data.diff_url  | string | diff的cdn下载地址，请求该地址可直接下载diff数据 |


## 获取模块组级别的配置API

`POST https://xx.wps.cn/api/v1/project/:pid/group/:mgid/config`


### 接口说明

?> 客户端通过该接口可获取到模块组级别的配置数据，如返回“在线参数”项目下的“稍后阅读”模块组下的所有模块数据。模块组下可有多个模块。

### 请求参数 Path Params参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| pid  | Int |   项目ID   | 是|
| mgid  | Int |   模块组ID   | 是|


### 请求参数 Body参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| md5  | string |    客户端上次请求服务端返回的md5，第一次请求传空  | 是|
| ver  | int |   客户端上次请求服务端返回的版本号，第一次请求传空  | 是|
| filters  | map[string]string |  客户端请求过滤参数，服务端根据该过滤条件返回过滤后的ids | 否|


### 请求示例

?> 请求模块组ID为1的配置数据，服务端判断客户端上次拉取的项目版本为10，如果服务端的该项目版本已发布到11，客户端则会拉取到版本为11的数据。

```json
https://xx.wps.cn/api/v1/project/1/config
{
  "md5": "69e2d0410fe805dfeaef461220d43a52",
  "ver": 10,
  "filters": {
    "app_version": "13.0",
	"deviceid": "11111",
	"platform": "1",
  }
}
```
### 返回示例

#### 返回示例一：全量动态

```json
{ 
	"ver": 11,                                                                   
	"md5": "5e5781c06253f12682cd5009",    
	"ret_type": 0, 
	"ids": [13],                                                   
	"data": {
		"full": {                                                   
			"mg_id": 1,                                                
			"mg_name": "稍后阅读",                           
			"sub": [],                                                  
			"modules": [{                                              
				"id": 13,                                                        
				"name": "屏蔽OEM渠道",
				"expire_time" : 1590748456,
				"effective_time" : 1590748456,
				"value": {                                               
					"timeInterval": 5
				}
			}]
		}
	}
}
```

#### 返回参数说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |   当前拉取配置的项目版本号  |
| md5  | string |  当前拉取配置的项目MD5 |
| ret_type  | int | 拉取策略, 0:动态全量, 1: 动态diff, 2:cdn全量, 3:cdn diff |
| ids  | array | 根据客户端参数过滤后的ID列表，客户端根据此ID列表解析下面的full全量数据 |
| data.full  | object | 该模块组下的配置 |
| full.mg_id  | int | 模块组ID |
| full.mg_name  | string | 模块组名称 |
| full.sub  | array | 子模块组，如果该模块组有配置子模块组，则子模块组的具体配置在该层级 |
| full.modules  | array | 该模块组下的具体模块配置  |
| modules.id  | int | 模块ID  |
| modules.name  | int | 模块名称  |
| modules.value  | object | 客户端需要的具体json参数，可自定义配置  |


#### 返回示例二：全量CDN（同上）
#### 返回示例三：DIFF动态（同上）
#### 返回示例三：DIFF CDN（同上）

## 获取模块级别的配置API

`POST https://xx.wps.cn/api/v1/project/:pid/group/:mgid/module/:mid/config`


### 接口说明

?> 客户端通过该接口可获取到模块级别的配置数据，如返回“在线参数”项目下的“稍后阅读”模块组下的“屏蔽OEM渠道”模块。

### 请求参数 Path Params参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| pid  | Int |   项目ID   | 是|
| mgid  | Int |   模块组ID   | 是|
| mid  | Int |   模块ID   | 是|


### 请求参数 Body参数
|  字段   | 类型  | 描述  |  是否必传 |
|  ----  | ----  | ---- |  ---- |
| md5  | string |    客户端上次请求服务端返回的md5，第一次请求传空  | 是|
| ver  | int |   客户端上次请求服务端返回的版本号，第一次请求传空  | 是|
| filters  | map[string]string |  客户端请求过滤参数，服务端根据该过滤条件返回过滤后的ids | 否|


### 请求示例

?> 请求模块ID为1的配置数据，服务端判断客户端上次拉取的项目版本为10，如果服务端的该项目版本已发布到11，客户端请则会拉取到版本为11的数据。

```json
https://xx.wps.cn/api/v1/module/1/config
{
  "md5": "69e2d0410fe805dfeaef461220d43a52",
  "ver": 10,
  "filters": {
    "app_version": "13.0",
	"deviceid": "11111",
	"platform": "1",
  }
}
```
### 返回示例

#### 返回示例一：全量动态

```json
{ 
	"ver": 11,
	"md5": "5e5781c06253f12682cd5009",
	"ret_type": 0,
	"ids": [13],
	"data": {
		"full": {
				"id": 13,
				"name": "屏蔽OEM渠道",
				"expire_time" : 1590748456,
				"effective_time" : 1590748456,
				"value": {
					"timeInterval": 5
				}
		}
	}
}
```

#### 返回参数说明
|  字段   | 类型  | 描述  |
|  ----  | ----  | ---- |
| ver  | int |   当前拉取配置的项目版本号  |
| md5  | string |  当前拉取配置的项目MD5 |
| ret_type  | int | 拉取策略, 0:动态全量, 1: 动态diff, 2:cdn全量, 3:cdn diff |
| ids  | array | 根据客户端参数过滤后的ID列表，客户端根据此ID列表解析下面的full全量数据 |
| data.full  | object | 该模块下的全量配置 |
| full.id  | int | 模块ID |
| full.name  | string | 模块名称 |
| full.value  | object | 客户端需要的具体json参数，可自定义配置  |


#### 返回示例二：全量CDN（同上）
#### 返回示例三：DIFF动态（同上）
#### 返回示例三：DIFF CDN（同上）


