package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	aSlice := make([]string, 0)
	for key,value := range data{
		if contains(value, interest) {
			aSlice = append(aSlice, key)
		}
	}
	return aSlice
}

func contains(src []string, elem string) bool {
	for _,value := range src{
		if(value == elem){
			return true
		}
	}
	return false
}
