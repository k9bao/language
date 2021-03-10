package mapg

//Key类型只要能支持==和!=操作符，即可以做为Key，当两个值==时，则认为是同一个Key。
//MapUtil
type MapUtil struct {
	m map[string]string
}

//Create return new MapUtil
func Create() *MapUtil {
	return &MapUtil{
		m: make(map[string]string),
	}
}

//String to print
func (self *MapUtil) String() string {
	str := ""
	for k, v := range self.m {
		str += k + ":" + v + "\n"
	}
	return str
}

//Insert if exist overwrite, else insert
func (self *MapUtil) Insert(k, v string) {
	self.m[k] = v
}

//Delete return key
func (self *MapUtil) Delete(k string) string {
	if v, ok := self.m[k]; ok {
		delete(self.m, k)
		return v
	}
	return ""
}

//Update must exist,if not exist return false
func (self *MapUtil) Update(k, v string) (string, bool) {
	if v0, ok := self.m[k]; ok {
		self.m[k] = v
		return v0, ok
	}
	return "", false
}

//GetV return V,exist
func (self *MapUtil) GetV(k string) (string, bool) {
	if v, ok := self.m[k]; ok {
		return v, ok
	}
	return "", false
}

//GetKeys return keys
func (self *MapUtil) GetKeys() []string {
	keys := []string{}
	for k := range self.m {
		keys = append(keys, k)
	}
	return keys
}
