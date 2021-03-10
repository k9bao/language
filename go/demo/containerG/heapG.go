package containerg

//定义一个正方形的结构体
type Rectangle struct {
	Width  int
	Height int
}

//
func (rec *Rectangle) Area() int {
	return rec.Width * rec.Height
}

// 定义一个堆结构体
type RectHeap []Rectangle

// 实现heap.Interface接口
func (rech RectHeap) Len() int {
	return len(rech)
}

// 实现sort.Iterface
func (rech RectHeap) Swap(i, j int) {
	rech[i], rech[j] = rech[j], rech[i]
}
func (rech RectHeap) Less(i, j int) bool {
	return rech[i].Area() < rech[j].Area()
}

// 实现heap.Interface接口定义的额外方法
func (rech *RectHeap) Push(h interface{}) {
	*rech = append(*rech, h.(Rectangle))
}
func (rech *RectHeap) Pop() (x interface{}) {
	n := len(*rech)
	x = (*rech)[n-1]      // 返回删除的元素
	*rech = (*rech)[:n-1] // [n:m]不包括下标为m的元素
	return x
}
