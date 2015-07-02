package main
/**
 * 数独解法，返回结果，类似c风格
 */
import (
	"fmt"
)

func main() {
	arr := [9][9]int{
		{0, 6, 1, 0, 3, 0, 0, 2, 0},
		{0, 5, 0, 0, 0, 8, 1, 0, 7},
		{0, 0, 0, 0, 0, 7, 0, 3, 4},
		{0, 0, 9, 0, 0, 6, 0, 7, 8},
		{0, 0, 3, 2, 0, 9, 5, 0, 0},
		{5, 7, 0, 3, 0, 0, 9, 0, 0},
		{1, 9, 0, 7, 0, 0, 0, 0, 0},
		{8, 0, 2, 4, 0, 0, 0, 6, 0},
		{0, 4, 0, 0, 1, 0, 2, 5, 0},
	}
	var result [9][9]int
	fill(&arr, &result, 0)
	output(result)
}

func fill(a, result *[9][9]int,  count int){
	if count == 81 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				result[i][j] = a[i][j]
			}
		}
		return 
	}
	
	row := int(count/9)
	col := int(count%9)
				
	//不需要填充数字，就跳到下一个
	if a[row][col] > 0 {
		fill(a, result, count + 1)
		return
	}
		
	for v := 1; v <= 9; v++ {
		if canFill(a, count, v) {
			a[row][col] = v
			fill(a,result, count + 1)
			a[row][col] = 0 //如果上面单位无解，回溯
		}
	}
}

func canFill(a *[9][9]int, count,value int) bool {
	var i, j, row, col int
	row = int(count / 9)
	col = int(count % 9)
	//横行|竖行检查
	for i = 0; i < 9; i++ {
		if (row != i && a[row][i] == value) || (col != i && a[i][col] == value) {
			return false
		}
	}

	//区块检查
	p := row - row % 3
	q := col - col % 3
	for i = p; i < p + 3; i++ {
		for j = q; j < q + 3; j++ {
			if a[i][j] == value {
				return false
			}
		}
	}
	return true	
}

func output(a [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println("")
	}
}

