package structs

import (
	"fmt"
	"math"
)

type Bag struct {
	items []int
}

// 指针式接收器,可以影响到外面内容

// 将一个物品放入背包的过程
func insert(b *Bag, item int)  {
	b.items = append(b.items, item)
}


// 比较喜欢这一种
func (b *Bag) insert1(item int) {
	b.items = append(b.items, item)
}

// 带有返回值 此中写法多余的,
// 已经对b *Bag 指针式,使用,可不必要返回整个指针类型
func (b *Bag) insert2(item int) *Bag {
	b.items = append(b.items, item)
	return b
}

// 接收器,可以非指针式的
// 传入的值,不会影响原来的内容,只是在函数内部有效
func (b Bag) add(item int) Bag {
	b.items = append(b.items, item)
	fmt.Println(b)
	return b
}

func Insert()  {
	bag := new(Bag)
	insert(bag, 1001)
	insert(bag, 1002)
	insert(bag, 1003)
	insert(bag, 1004)
	insert(bag, 1005)
	fmt.Println(bag)
	bag.insert1(10006)
	bag = bag.insert2(10007)
	bag.add(10008)
	for _, item := range bag.items{
		fmt.Println(item)
	}
}

/*
	实现二维矢量结构
*/


type Vec2 struct {
	X, Y float32
}

// 加
func (v Vec2) Add(other Vec2) Vec2 {

	ve := Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
	//fmt.Println(ve)
	return ve

}

// 减
func (v Vec2) Sub(other Vec2) Vec2 {

	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 乘
func (v Vec2) Scale(s float32) Vec2 {

	return Vec2{v.X * s, v.Y * s}
}

// 距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 插值
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}

func TwoDimensionalVector()  {
	// 二维失量
	vec := &Vec2{
		X: 1.0,
		Y: 1.0,
	}
	vec.Add(Vec2{X:1.1, Y:2.2})
	fmt.Println(vec)
}

/*
	 游戏玩家移动
*/

type Player struct {
	currPos   Vec2    // 当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// 移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {

	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 是否到达
func (p *Player) IsArrived() bool {

	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 逻辑更新
func (p *Player) Update() {

	if !p.IsArrived() {

		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()

		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))

		// 移动完成后，更新当前位置
		p.currPos = newPos
	}else {
		fmt.Println("更新失败", p.IsArrived())
	}

}

// 创建新玩家  使用构造函数
func NewPlayer(speed float32) *Player {

	p := &Player{
		speed: speed,
	}
	// targetPos
	p.MoveTo(Vec2{1.0, 2.0})
	p.Update()
	currPos := p.Pos()
	fmt.Println(currPos)
	return p
}