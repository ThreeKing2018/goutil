package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

func (v *viper) SetSplit(key string, value []interface{}) {
	// If alias passed in, then set the proper override
	//value = toCaseInsensitiveValue(value)

	path := strings.Split(key, v.keyDelim)
	lastKey := path[len(path)-1]
	//fmt.Println("path:",path,"lastkey:",lastKey)
	deepestMap := deepSearch(v.config, path[0:len(path)-1])
	//fmt.Println("deepestMap:",deepestMap[lastKey])

	// set innermost value
	deepestMap[lastKey] = value
}

func (v *viper) Set(key string, value interface{}) error {
	if _, ok := value.(string); ok {
		if strings.HasPrefix(value.(string), "[") && strings.HasSuffix(value.(string), "]") {

			var a1 = make([]interface{}, 0)
			data := []byte(value.(string))
			err := json.Unmarshal(data, &a1)
			if err != nil {
				return err
			}
			//fmt.Println("jsonerr", err)
			//fmt.Println("a1", len(a1))
			//v.config[a.Key] =a1
			//var v1 = make([]interface{},0)
			//jsonStringToObject(a.Value,&v1)

			//fmt.Println("len",len(v1))
			v.set(key, a1)
			return nil
		}
		if strings.HasPrefix(value.(string), "{") && strings.HasSuffix(value.(string), "}") {
			value = cast.ToStringMap(value)
			v.set(key, value)
			return nil
		}
	}

	err := v.fn(key, value)

	if err != nil {
		fmt.Println(err)
		return err
	}
	v.set(key, value)
	return nil
	//
	//switch  {
	//case strings.HasPrefix(value.(string),"[") && strings.HasSuffix(value.(string),"]"):
	//	var a1  = make([]interface{},0)
	//	data := []byte(value.(string))
	//	err:=json.Unmarshal(data, &a1)
	//	fmt.Println("jsonerr",err)
	//	fmt.Println("a1",len(a1))
	//	//v.config[a.Key] =a1
	//	//var v1 = make([]interface{},0)
	//	//jsonStringToObject(a.Value,&v1)
	//
	//	//fmt.Println("len",len(v1))
	//	v.set(key,a1)
	//case strings.HasPrefix(value.(string),"{") && strings.HasSuffix(value.(string),"}"):
	//	value=cast.ToStringMap(value)
	//
	//	v.set(key,value)
	//default:
	//
	//	err := v.fn(key,value)
	//	//err = remoteCfg.Fn(a)
	//	if err != nil {
	//		fmt.Println(err)
	//		return err
	//	}
	//	//key value 都是字符串
	//	v.set(key,value)
	//}
	//return nil

}

func (v *viper) set(key string, value interface{}) {
	// If alias passed in, then set the proper override
	//value = toCaseInsensitiveValue(value)

	v.mu.Lock()
	defer v.mu.Unlock()

	path := strings.Split(key, v.keyDelim)
	lastKey := path[len(path)-1]
	//fmt.Println("path:",path,"lastkey:",lastKey)
	deepestMap := deepSearch(v.config, path[0:len(path)-1])
	//fmt.Println("deepestMap:",deepestMap[lastKey])

	// set innermost value
	deepestMap[lastKey] = value
}

func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		// continue search from here
		m = m3
	}
	return m
}
