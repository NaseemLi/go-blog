package utils

func InList[T comparable](key T, list []T) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// 切片去重升级版 泛型参数 利用map的key不能重复的特性 + append函数 一次for循环搞定
func Unique[T comparable](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}

	newSlices := make([]T, 0) // 这里新建一个切片，大小为0，因为我们不知道有几个非重复数据
	m1 := make(map[T]bool)    // 用于判断是否重复
	for _, v := range ss {
		if _, ok := m1[v]; !ok { // 如果数据不在map中，放入
			m1[v] = true                     // 保存到map中，用于下次判断
			newSlices = append(newSlices, v) // 将数据放入新的切片中
		}
	}
	return newSlices
}
