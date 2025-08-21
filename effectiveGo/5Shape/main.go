package main

import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建矩形实例
	rectangle := Rectangle{Width: 5, Height: 4}
	
	// 创建圆形实例
	circle := Circle{Radius: 3}
	
	// 通过接口变量调用方法（多态）
	var shape1 Shape = rectangle
	var shape2 Shape = circle
	
	// 输出矩形结果
	fmt.Printf("矩形 Area: %.2f\n", shape1.Area())
	fmt.Printf("矩形 Perimeter: %.2f\n", shape1.Perimeter())
	
	// 输出圆形结果
	fmt.Printf("圆形 Area: %.16f\n", shape2.Area())
	fmt.Printf("圆形 Perimeter: %.16f\n", shape2.Perimeter())
}