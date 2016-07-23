package data_structures
import "fmt"
type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	mappedVals := make([]int32,len(vals), len(vals))
	for index, value := range vals {
		mappedVals[index] = op(value)
	}
	return mappedVals

}

func filterInts(op filterOperation, vals []int32) []int32 {
	filterVals := make([]int32,1,1)
	for _, value := range vals {
		if(op(value)){
			filterVals = append(filterVals, value) 
		}
	}
	return filterVals
}

func concatenate(dest []string, newValues ...string) []string {
	newSlice := make([]string,0,0)
		newSlice = append(newSlice, dest...)
		newSlice = append(newSlice, newValues...)
fmt.Println("%+v",newSlice)
	return newSlice
}

func equals(list1 []string, list2 []string) bool {
	for index,value := range list1{
		if(value != list2[index]){
			return false
		}
	}
	return true
}

func partialReverse(src []int, from, to int) []int {
	newSlice := make([]int,0,0)
	for i:=to; i>=from; i-- {
		newSlice = append(newSlice, src[i])
	}
	return newSlice
}
