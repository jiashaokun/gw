Admin Api

### 增加配置接口 (admin)
#### http://127.0.0.1:1323/req/add

字段名 | 类型 | 是否必选 | 备注 | 试例
---|---|---|---|---
name | string | 是 | 接口中文名称 | 获取详情接口
path | string | 是 | 访问地址路由 | /get/info (与api请求uri一致)
to | string | 是 | 转向地址 | http://baidu.com/get/info,http://baidu2.com/get/info
flow | int | 是 | 单机限流
method | string | 是 | 请求方式目前只包含 GET/POST
cache_time | int | 是 | 配置缓存时间（秒） | 200
timeout| int | 是 | 请求超时时间（秒）| 20
dns | int | 否 | 是否开启路由dns 0:no 1:yes | 1
decay| int | 否 | 是否开启请求衰减 0:no 1:yes | 1
decay_time | int | 否 | 请求衰减时间间隔 （秒）| 300

### 增加配置分组接口
#### http://127.0.0.1:1323/req/group

字段名 | 类型 | 是否必选 | 备注 | 试例
---|---|---|---|---
name | string | 是 | 接口中文名称 | 用户信息相关分组
group | string | 是 | 分组路由信息 | /user/ (字符串首尾有 '/' )

### group 分组列表接口
#### http://127.0.0.1:1323/req/list/group
