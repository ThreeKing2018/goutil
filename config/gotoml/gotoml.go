package gotoml

import (
	"bytes"
	"fmt"
	"github.com/pelletier/go-toml"
	"io"
	"io/ioutil"
	"os"
)

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
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	tree, err := toml.LoadReader(buf)
	if err != nil {
		return err
	}
	tmap := tree.ToMap()
	for k, v := range tmap {
		c[k] = v
		//return nil
	}
	fmt.Println(c)
	return nil
}

func (cfg *conf) WriteConfig(v interface{}) error {
	saveData, err :=  toml.TreeFromMap(v.(map[string]interface{}))
	if err != nil {
		return err
	}


	temp, err := ioutil.TempFile(".", ".tmp")
	if err != nil {
		return err
	}

	_, err = temp.WriteString(saveData.String())
	if err != nil {
		return err
	}

	temp.Close()
	//defer os.Remove(temp.Name())
	err = os.Rename(temp.Name(), cfg.configfile)

	if err != nil {
		return err
	}
	return nil
}
