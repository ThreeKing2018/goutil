package check

import (
	"errors"
	"github.com/ThreeKing2018/goutil/array"
	"path"
)

type Check struct {
}
//检查是否包含,不包含则报错
func (c *Check) InArrayString(layout, msg string, in []string) (err error) {
	if array.InArray(layout, in) == false {
		err = errors.New(msg)
		return
	}
	return
}
//检查参数是否为空
func (c *Check) Require(layout, msg string) (err error) {
	if layout == "" {
		err = errors.New(msg)
		return
	}
	return
}
//检查参数是否为空
func (c *Check) RequireInt(layout int, msg string) (err error) {
	if layout == 0 {
		err = errors.New(msg)
		return
	}
	return
}
//不相等则报错
func (c *Check) RequireNe(layout, msg, in string) (err error) {
	if layout == "" {
		err = errors.New(msg)
		return
	}
	if layout != in {
		err = errors.New(msg)
		return
	}
	return
}
//相等则报错
func (c *Check) RequireEq(layout, msg, in string) (err error) {
	if layout == "" {
		err = errors.New(msg)
		return
	}
	if layout == in {
		err = errors.New(msg)
		return
	}
	return
}
//检查图片是否上传
func (c *Check) Image(layout string, msg string) (err error) {
	if err = c.Require(layout, msg); err != nil {
		return
	}
	if isOk := c.CheckImageUrlAndExtension(layout); isOk == false {
		err = errors.New(msg)
		return
	}
	return
}
//检查图片
func (c *Check) CheckImageUrlAndExtension(image string, extension ...string) bool {
	if image == "" {
		return false
	}
	if len(extension) == 0 {
		extension = []string{"jpg", "jpeg", "png", "gif"}
	}
	ext := path.Ext(image)
	for _, v := range extension {
		if v == ext {
			return true
		}
	}

	return false
}
//检查图片是否上传 支持数组
func (c *Check) Images(layouts []string, msg string) (err error) {
	for _, image := range layouts {
		if err = c.Image(image, msg); err != nil {
			return
		}
	}
	return
}
//判断map key是否存在,不存在则false
func (c *Check) MapKey(layout map[string]interface{}, key string) bool {
	if _, ok := layout[key];ok{
		return true
	}
	return false
}