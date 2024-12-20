package cache

import "sync"

type InstanceType struct {
	M map[string]interface{}
}

var ins = &InstanceType{M: make(map[string]any)}

var once sync.Once

func Instance() *InstanceType {
	//once.Do(func() {
	//	ins = &InstanceType{M: make(map[string]any)}
	//})
	return ins
}

func (i *InstanceType) SetVal(key string, val any) {
	i.M[key] = val
}
func (i *InstanceType) GetVal(key string) any {
	return i.M[key]
}
