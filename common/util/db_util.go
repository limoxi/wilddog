package util

import "strings"

func ParseOrderFields(orderFields []string) string{
	result := make([]string, 0)
	for _, field := range orderFields{
		if strings.HasPrefix(field, "-"){
			result = append(result,  strings.Split(field, "-")[1] + " desc")
		}else{
			if strings.HasPrefix(field, "+") {
				field = strings.Split(field, "+")[1]
			}
			result = append(result,  field)
		}
	}
	return strings.Join(result, ", ")
}
