package main

import (
	"fmt"
)

func main() {
	// 1. 基本的 append 操作 - 向切片添加单个或多个元素
	// 创建一个初始切片
	slice1 := []int{1, 2, 3}
	fmt.Println("初始切片:", slice1)

	// 添加单个元素
	slice1 = append(slice1, 4)
	fmt.Println("添加单个元素后:", slice1)

	// 添加多个元素
	slice1 = append(slice1, 5, 6, 7)
	fmt.Println("添加多个元素后:", slice1)

	// 2. append 两个切片（使用 ... 展开第二个切片）
	slice2 := []int{8, 9, 10}
	// 测试1: 使用正确的 ... 语法（必须这样用）
	// slice1 = append(slice1, slice2...)
	slice1 = append(slice2, slice1...)
	fmt.Println("测试1 - 使用 ... 展开切片后:", slice1)

	// 注意：在Go中，... 操作符是必须的，用于将切片展开为单独的元素
	// 以下是错误示例，为了演示我们先注释掉
	// 测试2: 不使用 ... 语法（这会导致编译错误）
	// newSlice := []int{11, 12, 13}
	// slice1 = append(slice1, newSlice) // 编译错误: 不能将 []int 类型直接作为 int 参数传入
	// fmt.Println("不使用 ... 时的结果:", slice1)

	// 为什么刚才使用多个点(........)能运行？
	// 这是因为Go编译器会将多个连续的点视为三个点(...)的语法糖
	// 但这是不规范的写法，应该始终使用三个点(...)

	// 3. append 时的容量变化
	// 创建一个容量为3的切片
	slice3 := make([]int, 0, 3)
	fmt.Printf("初始 - 长度: %d, 容量: %d, 切片: %v\n", len(slice3), cap(slice3), slice3)

	// 添加元素并观察容量变化
	slice3 = append(slice3, 1)
	fmt.Printf("添加1个元素 - 长度: %d, 容量: %d, 切片: %v\n", len(slice3), cap(slice3), slice3)

	slice3 = append(slice3, 2, 3)
	fmt.Printf("添加到容量上限 - 长度: %d, 容量: %d, 切片: %v\n", len(slice3), cap(slice3), slice3)

	// 超过容量时，Go 会自动分配更大的内存
	slice3 = append(slice3, 4)
	fmt.Printf("超过容量 - 长度: %d, 容量: %d, 切片: %v\n", len(slice3), cap(slice3), slice3)

	// 4. 使用 append 删除切片元素
	// 删除中间元素（索引为2的元素）
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", slice4)
	slice4 = append(slice4[:2], slice4[3:]...)
	fmt.Println("删除索引2的元素后:", slice4)

	// 删除第一个元素
	slice4 = slice4[1:]
	fmt.Println("删除第一个元素后:", slice4)

	// 删除最后一个元素
	slice4 = slice4[:len(slice4)-1]
	fmt.Println("删除最后一个元素后:", slice4)

	// 5. append 与 nil 切片
	var slice5 []int // nil 切片
	fmt.Printf("nil切片 - 长度: %d, 容量: %d, 切片: %v\n", len(slice5), cap(slice5), slice5)

	// 可以直接向 nil 切片添加元素
	slice5 = append(slice5, 1, 2, 3)
	fmt.Printf("向nil切片添加元素后 - 长度: %d, 容量: %d, 切片: %v\n", len(slice5), cap(slice5), slice5)

	// 6. 在 for range 循环中批量删除或添加元素
	fmt.Println("\n=== 在 for range 循环中批量操作元素 ===")

	// 6.1 批量删除符合条件的元素（正确方法）
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原始数组:", numbers)

	// 方法1: 创建新切片
	var evenNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Println("方法1 - 保留偶数后的数组:", evenNumbers)

	// 方法2: 原地删除（适用于删除元素较少的情况）
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 重新初始化
	newIdx := 0
	for _, num := range numbers {
		if num%2 == 0 { // 保留偶数
			numbers[newIdx] = num
			newIdx++
		}
	}
	numbers = numbers[:newIdx]
	fmt.Println("方法2 - 原地保留偶数后的数组:", numbers)

	// 6.2 在循环中添加元素（注意：不能直接遍历原切片并修改）
	numbers = []int{1, 2, 3}
	fmt.Println("添加元素前:", numbers)

	// 错误方法：直接在for range中添加会导致死循环或遗漏元素
	// 正确方法1: 遍历副本
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	for _, num := range numbersCopy {
		if num < 5 {
			numbers = append(numbers, num*10) // 为小于5的元素添加一个10倍的新元素
		}
	}
	fmt.Println("方法1 - 遍历副本并添加元素后:", numbers)

	// 正确方法2: 从后向前遍历
	numbers = []int{1, 2, 3}
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] < 5 {
			numbers = append(numbers, numbers[i]*10)
		}
	}
	fmt.Println("方法2 - 从后向前遍历并添加元素后:", numbers)

	// 6.3 批量替换元素
	numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("替换前:", numbers)
	for i, num := range numbers {
		numbers[i] = num * 2 // 将每个元素乘以2
	}
	fmt.Println("批量替换后:", numbers)

}
