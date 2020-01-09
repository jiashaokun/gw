package conf

// 是否开鉴权 建议内网访问，不建议开启，开启后，所有访问接口（admin接口同样适用）都会进行鉴权
const AccessToken = 0

//app 主动断开时间
const RequestTimeOut = 30

const MongoPoll = 50

//Cache 缓存 mongoDB 整体数据缓存时间
const CacheMongoAllTime = 300

//Cache 全量数据存储Key
const CacheMongoAllData = "getCacheMongoAllData"

//Cache 最多连接数
const CachePoll = 100

//Cache 最少保持空闲
const CacheMinIdleConns = 10

const ReqMaxConnsPerHost = 200
