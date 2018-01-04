package lists

import "sort"

// based on https://medium.com/capital-one-developers/closures-are-the-generics-for-go-cb32021fb5b5

type sorter struct {
	len  int
	swap func(i, j int)
	less func(i, j int) bool
}

func (x sorter) Len() int                                      { return x.len }
func (x sorter) Swap(i, j int)                                 { x.swap(i, j) }
func (x sorter) Less(i, j int) bool                            { return x.less(i, j) }
func srt(n int, swap func(i, j int), less func(i, j int) bool) { sort.Sort(sorter{n, swap, less}) }

func ByteListListSort(list [][]byte) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return len(list[i]) < len(list[j]) })
}
func StringListSort(list []string) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func IntListSort(list []int) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func Int8ListSort(list []int8) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func Int16ListSort(list []int16) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func Int32ListSort(list []int32) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func Int64ListSort(list []int64) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func Float64ListSort(list []float64) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
func ByteListSort(list []byte) {
	srt(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] }, func(i, j int) bool { return list[i] < list[j] })
}
