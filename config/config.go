package config

// Config 服务配置
type Config struct {
	// http://host:port
	URL        string   `toml:"url"`
	DataBase   string   `toml:"database"`
	Prods      []string `toml:"prods"`
	Concurrent int      `toml:"concurrent"`
	Time       int      `toml:"time"`
	BatchSize  int      `toml:"batch-size"`
}
