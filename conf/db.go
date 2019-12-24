package conf

var MongoDB = map[string]string{
	"host":   "127.0.0.1",
	"user":   "admin",
	"pwd":    "",
	"dbname": "gw",
	"port":   "27017",
}

var Cache = map[string]string{
	"host": "127.0.0.1",
	"port": "6379",
}
