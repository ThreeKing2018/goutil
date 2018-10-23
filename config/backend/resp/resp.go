package resp

//为了解决循环导入的问题，才单独建立一个包

type Response struct {
	Action string
	Key    string
	Value  interface{}
	Error  error
}
