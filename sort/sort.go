package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Sortwrap struct {
	value    reflect.Value
	lessfunc func(int, int) bool
}

func NewSortwrap(v interface{}, lessfunc func(int, int) bool) *Sortwrap {
	return &Sortwrap{
		value:    reflect.ValueOf(v),
		lessfunc: lessfunc,
	}
}

func (s *Sortwrap) Len() int {
	return s.value.Len()
}

func (s *Sortwrap) Less(i, j int) bool {
	return s.lessfunc(i, j)
}

func (s *Sortwrap) Swap(i, j int) {
	value := s.value
	v1 := value.Index(i).Interface()
	v2 := value.Index(j).Interface()
	value.Index(j).Set(reflect.ValueOf(v1))
	value.Index(i).Set(reflect.ValueOf(v2))
}

type MyObject struct {
	value string
}

func (m *MyObject) Compare(m2 MyObject) int {
	// fmt.Println([]byte(m.value), []byte(m2.value), hex.EncodeToString([]byte(m.value)))
	return 1
	// return binary.BigEndian.Int64([]byte(m.value)) - binary.BigEndian.Int64([]byte(m2.value))
}

func main() {
	l1 := []int{1, 5, 634, 45, 2345, 62, 4}
	sort.Sort(NewSortwrap(
		l1,
		func(i, j int) bool {
			return l1[i] < l1[j]
		},
	))
	fmt.Println(l1)

	l2 := []MyObject{{"a"}, {"z"}, {"f"}, {"g"}}
	sort.Sort(NewSortwrap(
		l2,
		func(i, j int) bool {
			// return l2[i].Compare(l2[j]) < 0
			return l2[i].value < l2[j].value
		},
	))
	fmt.Println(l2)
}
