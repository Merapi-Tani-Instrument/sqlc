// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package authors

import (
	"regexp"
	"strings"
)

func appendQuery(q string, l int) (string, error) {
	if l <= 1 {
		return q, nil
	}
	regex, err := regexp.Compile(`\(([\s\?\,]*?)\)`)
	if err != nil {
		return "", err
	}
	res := regex.FindAllString(q, 1)[0]
	var join []string
	for i := 1; i < l; i++ {
		join = append(join, res)
	}
	return q + "," + strings.Join(join, ","), nil
}