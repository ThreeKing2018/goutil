# go 常用hash函数

## 示例

```
package main

import (
	"fmt"

	"github.com/ThreeKing2018/goutil/hash"
)

func main() {
	/* 字符串 */
	// md5
	md5 := hash.Md5String("111111")
	fmt.Println(md5)
	// sha1
	sha1 := hash.Sha1String("111111")
	fmt.Println(sha1)

	/* 字节 */
	// md5
	md5 = hash.Md5Byte([]byte("111111"))
	fmt.Println(md5)
	// sha1
	sha1 = hash.Sha1Byte([]byte("111111"))
	fmt.Println(sha1)

	/* 文件 */
	md5, err := hash.Md5File("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(md5)
}

```