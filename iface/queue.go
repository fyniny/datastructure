package iface

// Queue 队列 --- 先进先出(FIFO first in first out)的线性表
// 队列允许从一端进行插入操作，而在另一端进行删除操作的线性表
type Queue interface {
	// Init 初始化队列，list 为队列的存储引擎，Destroy() 会销毁该引擎
	Init(list LinearList)

	// Destroy 销毁队列, 若需要重用队列需要重新初始化
	Destroy()

	// Clear 清空栈
	Clear()

	// Head 返回队列的队头元素
	Head() interface{}

	// Push 插入队列元素
	Push(elem interface{}) error

	// Pop 弹出队头元素
	Delete() interface{}

	// Length 返回队长度
	Length() int
}