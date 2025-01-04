package main

import (
	"fmt"
)

// 図形を表す構造体
type Shape struct {
	Name string
	Area float64
}

// 複数の図形を管理する構造体
type ShapeManager struct {
	Shapes []Shape
}

// 図形を追加するメソッド
func (sm *ShapeManager) AddShape(s Shape) {
	sm.Shapes = append(sm.Shapes, s)
}

// 総面積を計算するメソッド
func (sm ShapeManager) TotalArea() float64 {
	var total float64
	for _, s := range sm.Shapes {
		total += s.Area
	}
	return total
}

func main() {
	// ShapeManagerの初期化
	manager := ShapeManager{}

	// 図形を追加
	manager.AddShape(Shape{Name: "Circle", Area: 78.5})
	manager.AddShape(Shape{Name: "Square", Area: 64.0})

	// 総面積を表示
	fmt.Printf("総面積: %.2f\n", manager.TotalArea())
}
