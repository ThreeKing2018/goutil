package ini

import "io"

type conf struct {
	configfile string
}

func NewConf() *conf {
	return &conf{}
}

func (cfg *conf) SetConfigFile(configfile string) {
	cfg.configfile = configfile
}

func (cfg *conf) ReadConfig(in io.Reader, c map[string]interface{}) error {
	return nil
}

func (cfg *conf) WriteConfig(v interface{}) error {
	return nil
}
