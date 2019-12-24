# Mongo 表结构设计


### wg

字段名 | 类型 | 解释 | 备注
---|---|---|---
id|string|uuid (唯一标识)| f2509627-08e5-4575-818c-f8b27b26c631
name|string|接口名称|获取用户信息
path|string|访问网关路径|/user/info
to|string|请求专项的地址/IP列表，英文逗号分隔 | http://0.0.0.1:1,http://0.0.0.1:2 或 http://a1.com,http://a2.com
dns|int32|是否需要dns | 0：不需要；1：需要 (若需要，则轮训访问字段 to 中的 IP 或 域名)
cacheTime|int32|配置缓存时间|单位：秒
timeout|int32|请求超时时间|单位：秒
decay|int32|请求衰减|0：不开启,1：开启 若开启超时或500时则逐渐递减请求，请求结果将是最后一次访问的正确结果
decayTime|int32|请求衰减间隔时间|单位：秒，decay=1时生效
createTime|string|创建时间|Y/m/d H:i:s
updateTime|string|更新时间|Y/m/d H:i:s