package config

// Config 写入参数配置
type Config struct {
	// http://host:port
	URL          string   `toml:"url"`
	DataBase     string   `toml:"database"`
	Measurements []string `toml:"measurements"`
	Concurrent   int      `toml:"concurrent"`
	Time         int      `toml:"time"`
	BatchSize    int      `toml:"batch-size"`
}
