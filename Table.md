# Mongo 表结构设计

### group 分组表
---
>* 路由分组，wg 中的path需要在该分组内，使用分组是为了避免添加接口反复重启项目的问题
>* 分钟中的path是模糊搜索的，所以不要重复比如 /user/* 和 /user/info/ 只需要 /user/ 即可
>* 只有分组中的请求能被访问到

字段名 | 类型 | 解释 | 备注
---|---|---|---
id|string|uuid (唯一标识)| f2509627-08e5-4575-818c-f8b27b26c631
name|string|分组名称|车辆基本信息组
group|string|分组信息|例：/user/  请完全按照格式填写
create_time|string|创建时间|Y-m-d H:i:s
update_time|string|更新时间|Y-m-d H:i:s


### wg

字段名 | 类型 | 解释 | 备注
---|---|---|---
id|string|uuid (唯一标识)| f2509627-08e5-4575-818c-f8b27b26c631
name|string|接口名称|获取用户信息
path|string|访问网关路径|/user/info（/user必须在Group表中）
to|string|请求专项的地址/IP列表，英文逗号分隔 | http://0.0.0.1:1,http://0.0.0.1:2 或 http://a1.com,http://a2.com
method|string|请求方式 GET/POST | 目前只支持 GET/POST 
dns|int|是否需要dns | 0：不需要；1：需要 (若需要，则轮训访问字段 to 中的 IP 或 域名)
flow|int|单机请求流量限制|0:代表不限流,该请求轮训字段to中的地址
cache_ime|int|数据缓存时间,单位：秒|0 不开启缓存
timeout|int|请求超时时间|单位：秒
decay|int|请求衰减|0：不开启,1：开启 若开启超时或500时则逐渐递减请求，请求结果将是最后一次访问的正确结果
decay_time|int|请求衰减间隔时间|单位：秒，decay=1时生效
create_time|string|创建时间|Y-m-d H:i:s
update_time|string|更新时间|Y-m-d H:i:s