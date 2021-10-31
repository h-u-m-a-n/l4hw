package tetris

import (
	"os"
	"regexp"
	"sort"
	"strings"
)

type Asd int8

var typeSizes = map[string]int{
	"bool":    1,
	"int8":    1,
	"int16":   2,
	"int32":   4,
	"float32": 4,
	"uint32":  4,
	"int64":   8,
	"uint64":  8,
	"float64": 8,
	"int":     8,
	"uint":    8,
	"float":   8,
	"*":       8,
	"string":  16,
}

func customSort(s []string) {
	sort.Slice(s, func(i, j int) bool {
		s1 := strings.Split(s[i], " ")
		s2 := strings.Split(s[j], " ")
		t1 := s1[len(s1)-1]
		t2 := s2[len(s2)-1]
		if t1[0] == '*' {
			t1 = t1[0:1]
		}
		if t2[0] == '*' {
			t2 = t2[0:1]
		}
		return typeSizes[t1] > typeSizes[t2]
	})

}

func SortStruct(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile(`\{([\s\S]*?)\}`)
	text := string(reg.Find(file))

	text = text[2 : len(text)-2]

	params := strings.Split(text, "\n")
	customSort(params)
	changed := make([]byte, 0)
	changed = append(changed, '{', '\n')
	for _, v := range params {
		temp := []byte(v)
		temp = append(temp, '\n')
		changed = append(changed, temp...)
	}
	changed = append(changed, '}')
	overwrite := reg.ReplaceAll(file, changed)
	if err := os.WriteFile(path, overwrite, 0); err != nil {
		return "", err
	}

	return string(changed), nil
}

func Get3Top(path string) ([][]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	reg := regexp.MustCompile(`\{([\s\S]*?)\}`)
	text := string(reg.Find(file))

	text = text[2 : len(text)-2]

	params := strings.Split(text, "\n")
	customSort(params)
	all := AllPerm(len(params))
	result := make([][]string, 0, 3)
	max, _ := findSizeOfType(params[0])
	optimalSize := getSizeOfStruct(params, all[0], max)
	for _, arr := range all {
		size := getSizeOfStruct(params, arr, max)
		if size == optimalSize {
			temp := make([]string, len(params))
			customCopy(temp, params, arr)
			result = append(result, temp)
		}
		if len(result) == cap(result) {
			break
		}
	}
	return result, nil
}

func customCopy(dest, src []string, order []int) {
	for originI, i := range order {
		dest[i] = src[originI]
	}
}

func getSizeOfStruct(s []string, a []int, max int) int {
	numberOfBlocs := 0
	remainingSize := 0
	numberOfString := 0
	for _, i := range a {
		temp, ok := findSizeOfType(s[i])
		if ok {
			numberOfString++
		}
		if remainingSize < temp {
			numberOfBlocs++
			remainingSize = max
		}
		remainingSize -= temp
	}
	return numberOfBlocs*max + numberOfString*8
}

func findSizeOfType(parameter string) (int, bool) {
	s := strings.Split(parameter, " ")
	t := s[len(s)-1]
	if t[0] == '*' {
		t = t[0:1]
	} else if t == "string" {
		return 8, true
	}
	return typeSizes[t], false
}

func AllPerm(n int) [][]int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	res := make([][]int, 0)
	temp := make([]int, len(arr))
	copy(temp, arr)
	res = append(res, temp)
	for nextSet(arr) {
		temp = make([]int, len(arr))
		copy(temp, arr)
		res = append(res, temp)
	}
	return res
}

func nextSet(arr []int) bool {
	n := len(arr)
	j := n - 2
	for j != -1 && arr[j] > arr[j+1] {
		j--
	}
	if j == -1 {
		return false
	}
	k := n - 1
	for arr[j] >= arr[k] {
		k--
	}
	arr[j], arr[k] = arr[k], arr[j]
	l := j + 1
	r := n - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	return true
}
