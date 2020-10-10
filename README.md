# gw 网关

>GW（是一个轻量级网关项目）包含请求转发，请求超时设置，请求限制，该项目对外提供数据增加的接口，暂无页面显示，具体接口文档将跟随项目 README 更新

# 项目依赖
- gin (所用框架)
- MongoDB (数据存储)
- Redis (缓存及限流等使用)

# 功能包含
- [动态路由（新增涵盖在组内接口动态添加不用重新部署项目）](#dtly)
- [请求转发 (目前只包含:GET/POST)](#qqzf)
- [多 IP/域名 配置](#dns)
- [请求 dns](#dns)
- [限流](#xl)
- [数据缓存](#hc) 
- [容错 （请求递减）](#rc)
- [统一鉴权](#jq)

# 如何使用
>* DOC.md 中提供了目前写入数据的接口文档
>* [TABLE.md](https://github.com/jiashaokun/gw/blob/master/TABLE.md) 中提供了MongoDB表设计文档

1. 启动项目
 ```shell
 go run main.go
 ```
2. 通过访问接口(http://127.0.0.1:1323/req/add/group) 或 自行添加MongoDB表（group）数据（路由组数据添加后，需要重新部署项目才可生效）。
3. 通过访问接口(http://127.0.0.1:1323/req/add/api) 或 自行添加MondoDB表 (wg) 数据（该数据的增加无需重新部署项目）。
4. 开始访问，比如在 group 中添加的是 /user/ 在 wg 中添加的是 /user/info 那么访问的地址就是 http://127.0.0.1:1323/user/info?id=1...

<h3 id='dtly'>动态路由</h3>

>* gw 提供动态路由功能，该功能避免增加接口后反复构建项目上线的操作，具体描述如下。

1. Mongo 中 group 表设置的是分组用途，比如拿瓜子二手车来讲：group 字段为 /car/ ，name（group字段的中文描述）字段为 车辆信息分组。
2. 举例要访问瓜子 car_id 为 999 的车辆信息，Mongo 中表 wg 中的字段 to 值应该为：http://api.guazi.xin.com。
3. 依据第 1 点举例，在 Mongo 表 wg 中添加的所有该分组路由都可以直接转发，但path规则为(举例)：/car/info?car_id=999...(参数列表)。

<h3 id='qqzf'>请求转发</h3>

>* gw 目前只提供 GET 和 POST 两种转发，后续会增加 RPC 调用。

1. Mongo 中 wg 表中字段 method 为请求方式，目前只支持 GET 和 POST。
2. 请求转发不影响前端调用，所有在 Mongo 表 group 中添加的分组，均提供多种调用方式比如：GET/POST/PUT/PATCH 等。
3. wg 表中的 method 的值为最终发送请求的方式。

<h3 id='dns'>DNS</h3>

1. 在多个IP配置时，gw 提供相应的 DNS 服务，通过 Mongo 中表 wg 中的字段 dns 0/1 来设置 关闭/开启。
2. 多IP配置方式为 Mongo 中标 wg 中字段 to 使用英文逗号分隔（末尾不加）举例：http://api1.guazi.com,http://127.0.0.1:1,http://127.0.0.1:2 。
3. 在字段 dns 为 1 的前提下，对请求提供轮训访问模式。

<h3 id='xl'>限流</h3>

1. 在 Mongo 中表 wg 中字段 flow 大于 0 时，将开启限流模式。
2. 限流模式将会最终落在每个 域名/IP 上，举例 flow:10 则 to 字段中的每个 域名/IP 都会限制 10 次请求。

<h3 id='hc'>数据缓存</h3>

1. 在 Mongo 中标 wg 中字段 cache_ime 大于 0 时，则代表缓存开启，单位: 秒。
2. 缓存模式是基于 Redis 的，所以请确保 conf 包下的 Cache 配置正确。

<h3 id='rc'>容错</h3>

1. 在 Mongo 中标 wg 中 decay 字段为 0/1 代表容错的 关闭/开启。
2. 当 decay 为 1 时，若请求遇到错误（http 500 或者 超时）时，将返回上一次请求的数据。
3. Mongo 中 wg 表中的字段 decay_time 为再次请求时间，当 decay 为 1 并且请求遇到错误时，间隔 decay_time 秒后，请求再次渗透，decay_time 秒内一直返回上一次的正确请求。

<h3 id='jq'>鉴权说明</h3>

1. 签名包含在Header中。
2. 请求参数示例：bc=1&ca=2&ba=3&ac=4 。
3. Content-Src: 1 (来源)。
4. Content-Md5: 03479689f0782dd83a03d9d4177af085 (签名值)。
5. 签名表结构请查看 TABLE.md 中的鉴权表说明。
6. 签名加密方式：MD5 32位小写。
7. 请求有效期：当天。

签名说明：加入请求参数是 bc=1&ca=2&ba=3&ac=4 ，那加密方式则是请求参数经过 ASCII 码有小到大排序后(参数排序 + 签名串 + 时间 到天) ab=3&ac=1&cd=2&sn=asf&src=1d4ac86d2-dc91-4675-942e-3aaa8bd4b5c12020-01-07 并对其进行MD5后为 a1f39e2e203a389a82974d1a6862e411 。

### 请求网关地址
<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/gw1.jpg"></img>

---
### 转发路径
<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/gw2.jpg"></img>

---

<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/pay.jpg" width="300" heigth="300">

