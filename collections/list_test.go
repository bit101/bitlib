// Package collections contains custom collection types.
package collections

import (
	"testing"
)

func TestListCreation(t *testing.T) {
	list := NewList[string]()

	length := list.Length()
	if length != 0 {
		t.Errorf("expected length %d, got %d", 0, length)
	}
}

func TestListAppendPrepend(t *testing.T) {
	list := NewList[string]()

	val := "foo"
	list.Append(val)

	length := list.Length()
	exp := 1
	if length != exp {
		t.Errorf("expected length %d, got %d", exp, length)
	}
	ret := list.GetFirst()
	if ret != val {
		t.Errorf("expected value %s, got %s", val, ret)
	}
	ret = list.GetLast()
	if ret != val {
		t.Errorf("expected value %s, got %s", val, ret)
	}

	val2 := "bar"
	list.Prepend(val2)

	length = list.Length()
	exp = 2
	if length != exp {
		t.Errorf("expected length %d, got %d", exp, length)
	}
	ret = list.GetFirst()
	if ret != val2 {
		t.Errorf("expected value %s, got %s", val2, ret)
	}
	ret = list.GetLast()
	if ret != val {
		t.Errorf("expected value %s, got %s", val, ret)
	}

	ret, err := list.Get(0)
	if err != nil {
		t.Error(err)
	}
	if ret != val2 {
		t.Errorf("expected value %s, got %s", val2, ret)
	}

	ret, err = list.Get(1)
	if err != nil {
		t.Error(err)
	}
	if ret != val {
		t.Errorf("expected value %s, got %s", val, ret)
	}

	b := list.Contains(val)
	expB := true
	if b != expB {
		t.Error("expected value to be in list")
	}

	b = list.Contains(val2)
	expB = true
	if b != expB {
		t.Error("expected value to be in list")
	}

	b = list.Contains("nope")
	expB = false
	if b != expB {
		t.Error("expected value to NOT be in list")
	}

	index := list.IndexOf(val)
	expIndex := 1
	if index != expIndex {
		t.Errorf("expected index %d, got %d", expIndex, index)
	}

	index = list.IndexOf(val2)
	expIndex = 0
	if index != expIndex {
		t.Errorf("expected index %d, got %d", expIndex, index)
	}
}

func TestListInsert(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	list.Insert(6, 999)
	val, err := list.Get(6)
	exp := 999

	if err != nil {
		t.Error(err)
	}
	if val != exp {
		t.Errorf("expected %d, got %d", exp, val)
	}
	length := list.Length()
	exp = 11
	if length != exp {
		t.Errorf("expected %d, got %d", exp, length)
	}

	err = list.Insert(100, 999)
	if err == nil {
		t.Error("expected error, got none")
	}
	err = list.Insert(-1, 999)
	if err == nil {
		t.Error("expected error, got none")
	}
}

func TestListRemove(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}

	val, err := list.Remove(5)
	exp := 5

	if err != nil {
		t.Error(err)
	}
	if exp != val {
		t.Errorf("expected %d, got %d", exp, val)
	}
	length := list.Length()
	exp = 9
	if exp != length {
		t.Error(list)
		t.Errorf("expected %d, got %d", exp, length)
	}

	val, err = list.Remove(100)
	exp = 0
	if err == nil {
		t.Error("expected error, got none")
	}
	if exp != val {
		t.Errorf("expected %d, got %d", exp, val)
	}

	val, err = list.Remove(-100)
	exp = 0
	if err == nil {
		t.Error("expected error, got none")
	}
	if exp != val {
		t.Errorf("expected %d, got %d", exp, val)
	}

	val = list.RemoveFirst()
	exp = 0
	if exp != val {
		t.Errorf("expected %d, got %d", exp, val)
	}

	val = list.RemoveLast()
	exp = 9
	if exp != val {
		t.Errorf("expected %d, got %d", exp, val)
	}
}

func TestListMisc(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i * 10)
	}

	index := list.IndexOf(50)
	exp := 5
	if exp != index {
		t.Errorf("expected %d, got %d", exp, index)
	}
}

func TestReverse(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	list.Reverse()
	for i := 0; i < 10; i++ {
		val := list.RemoveLast()
		if val != i {
			t.Errorf("expected %d, got %d", i, val)
		}
	}
}

func TestListShuffleSort(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	str1 := list.String()

	list.Shuffle()

	str2 := list.String()

	if str1 == str2 {
		t.Errorf("expected elements to be shuffled, got %s", str2)
	}

	list.Sort()
	str2 = list.String()
	if str1 != str2 {
		t.Errorf("expected elements to be sorted, got %s", str2)
	}
}

func TestListFilter(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}

	even := func(val int) bool {
		return val%2 == 0
	}

	list2 := list.Filter(even)

	str := list2.String()
	exp := "[0 2 4 6 8]"
	if exp != str {
		t.Errorf("expeced %s, got %s", exp, str)
	}

	high := func(val int) bool {
		return val > 4
	}
	list2 = list.Filter(high)

	str = list2.String()
	exp = "[5 6 7 8 9]"
	if exp != str {
		t.Errorf("expeced %s, got %s", exp, str)
	}
}

func TestListMap(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}

	double := func(val int) int {
		return val * 2
	}

	list2 := list.Map(double)

	str := list2.String()
	exp := "[0 2 4 6 8 10 12 14 16 18]"
	if exp != str {
		t.Errorf("expeced %s, got %s", exp, str)
	}

	strlist := NewList[string]()
	strlist.Append("tom")
	strlist.Append("dick")
	strlist.Append("harry")

	hello := func(value string) string {
		return "hello, " + value
	}
	strlist2 := strlist.Map(hello)

	val := strlist2.GetFirst()
	exp = "hello, tom"
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	val, _ = strlist2.Get(1)
	exp = "hello, dick"
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}

	val = strlist2.GetLast()
	exp = "hello, harry"
	if exp != val {
		t.Errorf("expected %s, got %s", exp, val)
	}
}

func TestListSlice(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i * 10)
	}

	for k, i := range list.Slice() {
		if i != k*10 {
			t.Errorf("expected %d, got %d", k*10, i)
		}
	}

	// mutating slice should not affect next call to slice.
	s := list.Slice()
	s[0] = 999
	s[2] = 999
	s[4] = 999
	s[6] = 999
	s[8] = 999
	for k, i := range list.Slice() {
		if i != k*10 {
			t.Errorf("expected %d, got %d", k*10, i)
		}
	}
}

func TestListClear(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 10; i++ {
		list.Append(i * 10)
	}
	list.Clear()

	length := list.Length()
	exp := 0
	if exp != length {
		t.Errorf("expected %d, got %d", exp, length)
	}
}
