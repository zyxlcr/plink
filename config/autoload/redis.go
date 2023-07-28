package autoload

type Redis struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	Pass string `mapstructure:"pass" json:"pass" yaml:"pass"`
	Tls  bool   `mapstructure:"tls" json:"tls" yaml:"tls"`
}
