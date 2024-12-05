package utils

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func GetMiddleVal(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)/2]
}
