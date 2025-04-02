###这是一个测试代码，搜索找到main函数
// math_test.go
package math

import "testing"

// 测试 Add 函数
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"正数相加", 2, 3, 5},
        {"负数相加", -1, -1, -2},
        {"零值相加", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; 期望值 %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// 测试 Subtract 函数
func TestSubtract(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"正数相减", 5, 3, 2},
        {"负数相减", -1, -1, 0},
        {"零值相减", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Subtract(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Subtract(%d, %d) = %d; 期望值 %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// main 函数是程序入口，不接受参数且无返回值， 我现在在做西南540000 ctc的。临时可以拿到一些你之前要的信息。你看你还需求不。如果需要的话，我就想想办法。不过我这次打算用 币 来玩。
func main() {
    fmt.Println("Hello, World!") // 输出内容
    result := add(3, 5)          // 调用其他函数
    fmt.Println("3 + 5 =", result)
    
}
