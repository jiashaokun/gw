package backend

type MongoInfo struct {
	Id         string `id`
	Name       string `name`
	Path       string `path`
	To         string `to`
	Dns        int    `dns`
	CacheTime  int    `cache_time`
	Timeout    int    `timeout`
	Decay      int    `decay`
	DecayTime  int    `decay_time`
	CreateTime string `create_time`
	UpdateTime string `update_time`
}

type MongoList struct {
	List []MongoInfo
}
