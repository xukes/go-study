package cache

type instanceType struct {
	M map[string]interface{}
}

var Ins *instanceType

func init() {
	Ins = &instanceType{M: make(map[string]any)}
}

func (i *instanceType) SetVal(key string, val any) {
	i.M[key] = val
}
func (i *instanceType) GetVal(key string) any {
	return i.M[key]
}
