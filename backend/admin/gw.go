package admin

type Add struct {
	Id         string `json:"id"`
	Name       string `json:"name" form:"name" binding:"required"`
	Path       string `json:"path" form:"path" binding:"required"`
	To         string `json:"to" form:"to" binding:"required"`
	Method     string `json:"method" form:"method" binding:"oneof=GET POST get post"`
	Dns        int    `json:"dns" form:"dns" binding:"oneof=0 1"`
	CacheTime  int    `json:"cache_time" form:"cache_time" binding:"required"`
	Timeout    int    `json:"timeout" form:"timeout" binding:"required"`
	Decay      int    `json:"decay" form:"decay" binding:"oneof=0 1"`
	DecayTime  int    `json:"decay_time" form:"decay_time"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
