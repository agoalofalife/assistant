package helpers

// List converts [][]string in []string where leave only regex group
func GroupExclude(foundString [][]string) []string{
	listFromGroup := make([]string, len(foundString))
	for position,value := range foundString{
		listFromGroup[position] = value[1]

	}
	return listFromGroup
}
