package lists

import "bytes"

func ByteListListContains(itemList [][]byte, target []byte) bool {
	for _, item := range itemList {
		if bytes.Equal(item, target) {
			return true
		}
	}
	return false
}

type Equaler interface {
	Equals(Equaler) bool
}

func EqualerListContains(itemList []Equaler, target Equaler) bool {
	for _, item := range itemList {
		if item.Equals(target) {
			return true
		}
	}
	return false
}

func StringListContains(itemList []string, target string) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func IntListContains(itemList []int, target int) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func Int8ListContains(itemList []int8, target int8) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func Int16ListContains(itemList []int16, target int16) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func Int32ListContains(itemList []int32, target int32) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func Int64ListContains(itemList []int64, target int64) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func Float64ListContains(itemList []float64, target float64) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func ByteListContains(itemList []byte, target byte) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}

func BoolListContains(itemList []bool, target bool) bool {
	for _, item := range itemList {
		if item == target {
			return true
		}
	}
	return false
}
