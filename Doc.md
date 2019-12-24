Admin Api

### 增加配置接口
字段名 | 类型 | 备注 | 试例
---|---|---|---
name | string | 接口中文名称 | 获取详情接口
path | string | 访问地址路由 | /get/info (与api请求uri一致)
to | string | 转向地址 | http://baidu.com/get/info,http://baidu2.com/get/info
cache_time | int | 配置缓存时间（秒） | 200
dns | int | 是否开启路由dns 0:no 1:yes | 1
decay| int | 是否开启请求衰减 0:no 1:yes | 1
decay_time | int | 请求衰减时间间隔 （秒）| 300