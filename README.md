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
>* 请求参数: cb=1&ca=2&ba=3&a=4&_JiaSrc=1
>* 加密参数： _JiaSrc=123&a=4&ba=3&ca=2&cb=1232020-01-01key
>* MD5后：03479689f0782dd83a03d9d4177af085
>* 最终请求参数:cb=1&ca=2&ba=3&a=4&_JiaSrc=1&_JiaSign=03479689f0782dd83a03d9d4177af085
>* 鉴权参数：_JiaSrc 和 _JiaSign 只参与签名，不会参与最终后端接口的请求

+ 签名加密方式：MD5 32位小写
+ 签名需要增加传递参数：_JiaSrc（参与签名） 和 _JiaSign（签名后的md5），这两个参数不会参与到后端接口的请求中
+ 请求有效期：当天 ，具体请查看 TABLE.md 中的鉴权表说明。
+ 签名方式：所有参数（包含:_JiaSrc）按参数名的 ASCII 码从小到大排序，末尾添加当天的日期 YYYY-MM-dd

### 请求网关地址
<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/gw1.jpg"></img>

---
### 转发路径
<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/gw2.jpg"></img>

---

<img src="https://raw.githubusercontent.com/jiashaokun/doc/master/txt/pay.jpg" width="300" heigth="300">
