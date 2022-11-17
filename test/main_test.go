package test

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	totalItem := 5
	listItem := [][]int{
		{4, 2},
		{8, 8},
	}
	// totalUser := 3
	query := [][]int{
		{4, 5, 2, 3},
	}

	foundItem := 0

	for i := 0; i < totalItem; i++ {
		for j := 0; j < len(listItem); j++ {
			for k := 0; k < len(query); k++ {
				trueCount := 0

				Lj := query[k][0]
				Vi := listItem[j][0]
				Rj := query[k][1]

				if Lj <= Vi && Vi <= Rj {
					trueCount++
				}

				Xj := query[k][2]
				Qj := listItem[j][1]
				Yj := query[k][3]

				if Xj <= Qj && Qj <= Yj {
					trueCount++
				}

				if trueCount == 2 {
					foundItem++
				}
			}
		}
	}

	fmt.Println(foundItem)
}

func TestMinimumSwap(t *testing.T) {
	arr := []int{7, 1, 3, 2, 4, 5, 6}
	cSwap := 0
	for i := 0; i < len(arr); i++ {
		minValue := arr[i]
		indexMinValue := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minValue {
				minValue = arr[j]
				indexMinValue = j
				cSwap++
			}
		}
		arr[indexMinValue] = arr[i]
		arr[i] = minValue
	}

	fmt.Println(cSwap)
}

func TestMultipler(t *testing.T) {
	for i := 1; i <= 100; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		} else if i%15 == 0 {
			fmt.Println("Mampu")
			continue
		} else if i%5 == 0 {
			fmt.Println("Pu")
			continue
		} else if i%3 == 0 {
			fmt.Println("Mam")
			continue
		}
	}
}
