package config

type Redis struct {
	NetWork string `json:"netWork"`
	Addr    string `json:"addr"`
	Port    string `json:"port"`
	Pwd     string `json:"pwd"`
	Prefix  string `json:"prefix"`
}
