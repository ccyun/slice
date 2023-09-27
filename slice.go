package slice

func toMap[T comparable](sliceData []T) map[T]bool {
	mapData := make(map[T]bool)
	for _, v := range sliceData {
		mapData[v] = true
	}
	return mapData
}

func toSlice[T comparable](mData map[T]bool) []T {
	data := []T{}
	for item := range mData {
		data = append(data, item)
	}
	return data
}

// InSlice 判断切片是否存在某个元素
func InSlice[T comparable](value T, sliceData []T) bool {
	for _, item := range sliceData {
		if value == item {
			return true
		}
	}
	return false
}

// SliceUnique 切片去重复
func SliceUnique[T comparable](sliceData []T) []T {
	var data []T
	for item := range toMap(sliceData) {
		data = append(data, item)
	}
	return data
}

// SliceDiff 返回切片差集
func SliceDiff[T comparable](slices ...[]T) []T {
	snum := len(slices)
	switch snum {
	case 0:
		return []T{}
	case 1:
		return slices[0]
	}
	data := toMap(slices[0])
	for i := 1; i < snum; i++ {
		for _, item := range slices[i] {
			if _, ok := data[item]; ok {
				delete(data, item)
			}
		}
	}
	return toSlice(data)
}

// SliceMerge 多个切片的并集
func SliceMerge[T comparable](slices ...[]T) []T {
	data := []T{}
	for _, item := range slices {
		data = append(data, item...)
	}
	return SliceUnique(data)
}

// SliceIntersect 返回切片交集
func SliceIntersect[T comparable](slices ...[]T) []T {
	snum := len(slices)
	switch snum {
	case 0:
		return []T{}
	case 1:
		return slices[0]
	}

	tempData := make(map[T]int)
	for i := 0; i < snum; i++ {
		for v := range toMap(slices[i]) {
			tempData[v]++
		}
	}
	data := []T{}
	for k, v := range tempData {
		if v == snum {
			data = append(data, k)
		}
	}
	return data
}
