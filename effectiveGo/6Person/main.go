package main

import "fmt"

// Person 结构体包含 Name 和 Age 字段
type Person struct {
    Name string
    Age  int
}

// Employee 结构体嵌入 Person 并添加 EmployeeID 字段
type Employee struct {
    Person
    EmployeeID string
}

// PrintInfo 方法输出员工的信息
func (e Employee) PrintInfo() {
    fmt.Printf("姓名: %s\n", e.Name)
    fmt.Printf("年龄: %d\n", e.Age)
    fmt.Printf("员工ID: %s\n", e.EmployeeID)
}

func main() {
    // 创建 Employee 实例并初始化所有字段
    emp := Employee{
        Person: Person{
            Name: "Alice",
            Age:  30,
        },
        EmployeeID: "E12345",
    }
    
    // 调用方法输出完整的员工信息
    emp.PrintInfo()
}