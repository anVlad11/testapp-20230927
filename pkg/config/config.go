package config

type App struct {
	HTTPServer HTTPServer `mapstructure:"http_server"`
	Orders     Orders     `mapstructure:"orders"`
}

type HTTPServer struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
	Listen  bool   `mapstructure:"listen"`
}

type Orders struct {
	PackSizes []int `mapstructure:"pack_sizes"`
}
