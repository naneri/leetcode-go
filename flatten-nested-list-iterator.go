// https://leetcode.com/problems/flatten-nested-list-iterator/

type NestedIterator struct {
	Pointer int
	Ints []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iterator := NestedIterator{Ints: make([]int, 0)}

	for _, val := range nestedList {
		if val.IsInteger() {
			iterator.Ints = append(iterator.Ints, val.GetInteger())
		} else {
			iterator.Ints = append(iterator.Ints, processIntegerList(val.GetList())...)	
		}
	}
	
	return &iterator
}

func processIntegerList(nestedList []*NestedInteger) []int {
	result := make([]int, 0)
	for _, val := range nestedList {
		if val.IsInteger() {
			result = append(result, val.GetInteger())
		} else {
			result = append(result, processIntegerList(val.GetList())...)
		}
	}
	
	return result
}


func (this *NestedIterator) Next() int {
	val := this.Ints[this.Pointer]
	this.Pointer++
	return val
}

func (this *NestedIterator) HasNext() bool {
	if this.Pointer >= len(this.Ints) {
		return false
	}
	return true
}
