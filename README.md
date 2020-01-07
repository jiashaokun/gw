# gw 网关

GW（网关）包含请求转发，请求超时设置，请求限制，该项目对外提供数据增加的接口，暂无页面显示，具体接口文档将跟随项目 README 更新

# 项目依赖
- gin (所用框架)
- MongoDB (数据存储)
- Redis (缓存及限流等使用)

# 功能包含
- 动态路由（新增涵盖在组内接口动态添加不用重新部署项目）
- 请求转发 (目前只包含:GET/POST)
- 多 IP/域名 配置
- 请求 dns
- 限流
- 数据缓存 
- 容错 （请求递减）
- [统一鉴权](#jq)

# 如何使用
>* DOC.md 中提供了目前写入数据的接口文档
>* TABLE.md 中提供了MongoDB表设计文档

1. 启动项目
 ```shell
 go run main.go
 ```

2. 通过访问接口(http://127.0.0.1:1323/req/add/group) 或 自行添加MongoDB表（group）数据（路由组数据添加后，需要重新部署项目才可生效）。
3. 通过访问接口(http://127.0.0.1:1323/req/add/api) 或 自行添加MondoDB表 (wg) 数据（该数据的增加无需重新部署项目）。
4. 开始访问，比如在 group 中添加的是 /user/ 在 wg 中添加的是 /user/info 那么访问的地址就是 http://127.0.0.1:1323/user/info?id=1...

<h3 id='jq'>鉴权说明</h3>
>* 签名包含在Header中
>* 请求参数示例：bc=1&ca=2&ba=3&ac=4
>* Content-Src: 1 (来源)
>* Content-Md5: 03479689f0782dd83a03d9d4177af085 (签名值)
>* 签名表结构请查看 TABLE.md 中的鉴权表说明
>* 签名加密方式：MD5 32位小写
>* 请求有效期：当天

签名说明：加入请求参数是 bc=1&ca=2&ba=3&ac=4 ，那加密方式则是请求参数经过 ASCII 码有小到大排序后(参数排序 + 签名串 + 时间 到天) ab=3&ac=1&cd=2&sn=asf&src=1d4ac86d2-dc91-4675-942e-3aaa8bd4b5c12020-01-07 并对其进行MD5后为 a1f39e2e203a389a82974d1a6862e411 。

### 请求网关地址
<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/gw1.jpg"></img>

---
### 转发路径
<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/gw2.jpg"></img>

---

<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/pay.jpg" width="300" heigth="300">
