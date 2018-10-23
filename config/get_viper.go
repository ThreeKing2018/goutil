package config

import (
	"github.com/spf13/cast"
	"strings"
	"time"
)

type Getvalue interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
	GetIntSlice(key string) []int
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
}

func (v *viper) GetString(key string) string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToString(v.get(key))
}

func (v *viper) GetBool(key string) bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToBool(v.get(key))
}

func (v *viper) GetInt(key string) int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToInt(v.get(key))
}

func (v *viper) GetInt32(key string) int32 {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToInt32(v.get(key))
}

func (v *viper) GetInt64(key string) int64 {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToInt64(v.get(key))
}

func (v *viper) GetFloat64(key string) float64 {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToFloat64(v.get(key))
}

func (v *viper) GetTime(key string) time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToTime(v.get(key))
}

func (v *viper) GetDuration(key string) time.Duration {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToDuration(v.get(key))
}

func (v *viper) GetStringSlice(key string) []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToStringSlice(v.get(key))
}

func (v *viper) GetIntSlice(key string) []int {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToIntSlice(v.get(key))
}

func (v *viper) GetStringMap(key string) map[string]interface{} {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToStringMap(v.get(key))
}

func (v *viper) GetStringMapString(key string) map[string]string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToStringMapString(v.get(key))
}

func (v *viper) GetStringMapStringSlice(key string) map[string][]string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return cast.ToStringMapStringSlice(v.get(key))
}

// 有远程配置中心， 下来配置到 --> config  保存到本地
// 没有远程配置中心 读取本地配置 --> config

// Get 将按照 flag  > config > defaults
func (v *viper) get(key string) interface{} {
	val := v.find(key)
	if val == nil {
		return nil
	}

	return val
}

// find 将按照 flag > config > defaults
// 这个版本 写完成 config > defaults
func (v *viper) find(lcaseKey string) interface{} {
	var (
		val    interface{}
		path   = strings.Split(lcaseKey, v.keyDelim)
		nested = len(path) > 1
	)

	// Config file next
	val = v.searchMapWithPathPrefixes(v.config, path)
	if val != nil {
		return val
	}
	if nested && v.isPathShadowedInDeepMap(path, v.config) != "" {
		return nil
	}

	// Default next
	val = v.searchMap(v.defaults, path)
	if val != nil {
		return val
	}
	if nested && v.isPathShadowedInDeepMap(path, v.defaults) != "" {
		return nil
	}
	return nil

}

// searchMapWithPathPrefixes recursively searches for a value for path in source map.
//
// While searchMap() considers each path element as a single map key, this
// function searches for, and prioritizes, merged path elements.
// e.g., if in the source, "foo" is defined with a sub-key "bar", and "foo.bar"
// is also defined, this latter value is returned for path ["foo", "bar"].
//
// This should be useful only at config level (other maps may not contain dots
// in their keys).
//
// Note: This assumes that the path entries and map keys are lower cased.
func (v *viper) searchMapWithPathPrefixes(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	// search for path prefixes, starting from the longest one
	for i := len(path); i > 0; i-- {
		prefixKey := strings.ToLower(strings.Join(path[0:i], v.keyDelim))

		next, ok := source[prefixKey]
		if ok {
			// Fast path
			if i == len(path) {
				return next
			}

			// Nested case
			var val interface{}
			switch next.(type) {
			case map[interface{}]interface{}:
				val = v.searchMapWithPathPrefixes(cast.ToStringMap(next), path[i:])
			case map[string]interface{}:
				// Type assertion is safe here since it is only reached
				// if the type of `next` is the same as the type being asserted
				val = v.searchMapWithPathPrefixes(next.(map[string]interface{}), path[i:])
			default:
				// got a value but nested key expected, do nothing and look for next prefix
			}
			if val != nil {
				return val
			}
		}
	}

	// not found
	return nil
}

// isPathShadowedInDeepMap makes sure the given path is not shadowed somewhere
// on its path in the map.
// e.g., if "foo.bar" has a value in the given map, it “shadows”
//       "foo.bar.baz" in a lower-priority map
func (v *viper) isPathShadowedInDeepMap(path []string, m map[string]interface{}) string {
	var parentVal interface{}
	for i := 1; i < len(path); i++ {
		parentVal = v.searchMap(m, path[0:i])
		if parentVal == nil {
			// not found, no need to add more path elements
			return ""
		}
		switch parentVal.(type) {
		case map[interface{}]interface{}:
			continue
		case map[string]interface{}:
			continue
		default:
			// parentVal is a regular value which shadows "path"
			return strings.Join(path[0:i], v.keyDelim)
		}
	}
	return ""
}

// searchMap recursively searches for a value for path in source map.
// Returns nil if not found.
// Note: This assumes that the path entries and map keys are lower cased.
func (v *viper) searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	next, ok := source[path[0]]
	if ok {
		// Fast path
		if len(path) == 1 {
			return next
		}

		// Nested case
		switch next.(type) {
		case map[interface{}]interface{}:
			return v.searchMap(cast.ToStringMap(next), path[1:])
		case map[string]interface{}:
			// Type assertion is safe here since it is only reached
			// if the type of `next` is the same as the type being asserted
			return v.searchMap(next.(map[string]interface{}), path[1:])
		default:
			// got a value but nested key expected, return "nil" for not found
			return nil
		}
	}
	return nil
}
