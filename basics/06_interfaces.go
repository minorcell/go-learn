package main

import (
	"fmt"
	"math"
)

/*
06_interfaces.go - Go语言基础：接口
学习内容：
1. 接口定义和实现
2. 接口组合
3. 空接口
4. 类型断言
5. 多态性
*/

// 定义几何图形接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义可描述接口
type Describable interface {
	Description() string
}

// 组合接口
type GeometricShape interface {
	Shape       // 嵌入Shape接口
	Describable // 嵌入Describable接口
}

// 圆形实现Shape接口
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Description() string {
	return fmt.Sprintf("这是一个半径为%.2f的圆形", c.Radius)
}

// 矩形实现Shape接口
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Description() string {
	return fmt.Sprintf("这是一个宽%.2f高%.2f的矩形", r.Width, r.Height)
}

// 动物接口示例
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + "说: 汪汪!"
}

func (d Dog) Move() string {
	return d.Name + "在跑步"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Cat + "说: 喵喵!"
}

func (c Cat) Move() string {
	return c.Name + "在悄悄走路"
}

// 工具函数：打印图形信息
func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

// 工具函数：打印完整图形信息
func printGeometricShapeInfo(gs GeometricShape) {
	fmt.Printf("%s\n", gs.Description())
	fmt.Printf("面积: %.2f, 周长: %.2f\n", gs.Area(), gs.Perimeter())
}

// 工具函数：让动物表演
func performAnimal(a Animal) {
	fmt.Printf("%s\n", a.Speak())
	fmt.Printf("%s\n", a.Move())
}

func main() {
	fmt.Println("=== Go语言基础：接口 ===")

	// 1. 基本接口使用
	fmt.Println("\n1. 基本接口使用：")

	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	fmt.Println("圆形信息:")
	printShapeInfo(circle)

	fmt.Println("矩形信息:")
	printShapeInfo(rectangle)

	// 2. 接口切片
	fmt.Println("\n2. 接口切片：")

	shapes := []Shape{
		Circle{Radius: 3.0},
		Rectangle{Width: 5.0, Height: 4.0},
		Circle{Radius: 2.5},
		Rectangle{Width: 3.0, Height: 3.0},
	}

	totalArea := 0.0
	for i, shape := range shapes {
		fmt.Printf("图形%d - ", i+1)
		printShapeInfo(shape)
		totalArea += shape.Area()
	}
	fmt.Printf("总面积: %.2f\n", totalArea)

	// 3. 组合接口
	fmt.Println("\n3. 组合接口：")

	fmt.Println("圆形完整信息:")
	printGeometricShapeInfo(circle)

	fmt.Println("矩形完整信息:")
	printGeometricShapeInfo(rectangle)

	// 4. 动物接口示例
	fmt.Println("\n4. 动物接口示例：")

	dog := Dog{Name: "小白"}
	cat := Cat{Name: "小黑"}

	fmt.Println("狗的表演:")
	performAnimal(dog)

	fmt.Println("猫的表演:")
	performAnimal(cat)

	// 动物园
	animals := []Animal{
		Dog{Name: "大黄"},
		Cat{Name: "小花"},
		Dog{Name: "阿旺"},
	}

	fmt.Println("动物园表演:")
	for i, animal := range animals {
		fmt.Printf("第%d个动物:\n", i+1)
		performAnimal(animal)
		fmt.Println()
	}

	// 5. 空接口
	fmt.Println("\n5. 空接口：")

	var anything interface{}

	anything = 42
	fmt.Printf("存储整数: %v (类型: %T)\n", anything, anything)

	anything = "Hello"
	fmt.Printf("存储字符串: %v (类型: %T)\n", anything, anything)

	anything = []int{1, 2, 3}
	fmt.Printf("存储切片: %v (类型: %T)\n", anything, anything)

	anything = Circle{Radius: 1.0}
	fmt.Printf("存储圆形: %v (类型: %T)\n", anything, anything)

	// 空接口切片
	mixed := []interface{}{
		42,
		"字符串",
		true,
		Circle{Radius: 2.0},
		[]int{1, 2, 3},
	}

	fmt.Println("混合类型切片:")
	for i, item := range mixed {
		fmt.Printf("  [%d] %v (类型: %T)\n", i, item, item)
	}

	// 6. 类型断言
	fmt.Println("\n6. 类型断言：")

	var shape Shape = Circle{Radius: 3.0}

	// 类型断言 - 安全方式
	if circle, ok := shape.(Circle); ok {
		fmt.Printf("这是一个圆形，半径: %.2f\n", circle.Radius)
	} else {
		fmt.Println("这不是圆形")
	}

	// 类型断言 - 检查是否为矩形
	if rectangle, ok := shape.(Rectangle); ok {
		fmt.Printf("这是一个矩形，宽: %.2f, 高: %.2f\n", rectangle.Width, rectangle.Height)
	} else {
		fmt.Println("这不是矩形")
	}

	// 对空接口进行类型断言
	fmt.Println("\n空接口类型断言:")
	processValue := func(value interface{}) {
		switch v := value.(type) {
		case int:
			fmt.Printf("整数: %d, 平方: %d\n", v, v*v)
		case string:
			fmt.Printf("字符串: %s, 长度: %d\n", v, len(v))
		case bool:
			fmt.Printf("布尔值: %t\n", v)
		case Circle:
			fmt.Printf("圆形: 半径=%.2f, 面积=%.2f\n", v.Radius, v.Area())
		case Rectangle:
			fmt.Printf("矩形: 宽=%.2f, 高=%.2f, 面积=%.2f\n", v.Width, v.Height, v.Area())
		default:
			fmt.Printf("未知类型: %T, 值: %v\n", v, v)
		}
	}

	testValues := []interface{}{
		42,
		"Go语言",
		true,
		Circle{Radius: 2.5},
		Rectangle{Width: 3.0, Height: 4.0},
		[]int{1, 2, 3},
	}

	for _, value := range testValues {
		processValue(value)
	}

	// 7. 接口nil检查
	fmt.Println("\n7. 接口nil检查：")

	var nilShape Shape
	fmt.Printf("nilShape == nil: %t\n", nilShape == nil)

	var nilAnimal *Dog
	var animalInterface Animal = nilAnimal
	fmt.Printf("animalInterface == nil: %t\n", animalInterface == nil)
	fmt.Printf("animalInterface 的值: %v\n", animalInterface)

	// 8. 接口的多态性
	fmt.Println("\n8. 接口的多态性：")

	// 计算总面积的通用函数
	calculateTotalArea := func(shapes []Shape) float64 {
		total := 0.0
		for _, shape := range shapes {
			total += shape.Area()
		}
		return total
	}

	differentShapes := []Shape{
		Circle{Radius: 1.0},
		Rectangle{Width: 2.0, Height: 3.0},
		Circle{Radius: 1.5},
	}

	total := calculateTotalArea(differentShapes)
	fmt.Printf("不同图形的总面积: %.2f\n", total)

	// 9. 接口赋值和转换
	fmt.Println("\n9. 接口赋值和转换：")

	var shape1 Shape
	var shape2 GeometricShape

	c := Circle{Radius: 4.0}

	shape1 = c // Circle实现了Shape接口
	shape2 = c // Circle也实现了GeometricShape接口

	fmt.Printf("shape1面积: %.2f\n", shape1.Area())
	fmt.Printf("shape2描述: %s\n", shape2.Description())

	// 从GeometricShape转换为Shape
	shape1 = shape2
	fmt.Printf("转换后shape1面积: %.2f\n", shape1.Area())
}
