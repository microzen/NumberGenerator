package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	num := 0
	isSort := false
	unsigned := false
	numbertype := map[string]int{
		"byte":  int(math.Pow(2, 8)),
		"word":  int(math.Pow(2, 16)),
		"dword": int(math.Pow(2, 32)),
	}
	size := numbertype["byte"]
	for i, arg := range args {
		if arg == "-n" || arg == "--number" {
			val, err := strconv.Atoi(args[i+1])
			if err != nil {
				println(err.Error())
				return
			}
			num = val
		} else if arg == "--sort" {
			isSort = true
		} else if arg == "--size" {
			size = numbertype[args[i+1]]
		} else if arg == "--unsigned" {
			unsigned = true
		} else if arg == "-s" || arg == "--shuffle" {
			pr(shuffle(args[i+1]))
			return
		}
	}
	if !unsigned {
		size = size/2 - 1
	}
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = rand.Intn(size)
		if rand.Int()%2 == 0 && !unsigned {
			arr[i] *= (-1)
		}
	}
	if isSort {
		sort.Ints(arr)
	}
	pr(arr)

}
func shuffle(input string) []int {
	parts := strings.Split(input, ",")
	nums := make([]int, len(parts))
	for i, part := range parts {

		num, err := strconv.Atoi(strings.Trim(part, " "))
		if err != nil {
			println(err.Error())
			return nil
		}
		nums[i] = num
	}
	for i := 0; i < len(nums); i++ {
		r := rand.Intn(len(nums))
		nums[r], nums[i] = nums[i], nums[r]
	}
	return nums
}
func pr(arr []int) {
	for i, n := range arr {
		if i == len(arr)-1 {
			fmt.Printf("%d\n", n)
		} else {
			fmt.Printf("%d, ", n)
		}
	}
}
