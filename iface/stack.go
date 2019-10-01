package iface

// Stack 栈 --- 后进先出(LIFO last in first out)的线性表
// 栈是限定仅在表尾进行插入和删除操作的现线表
// 栈顶： 允许插入和删除的一端
// 栈底： 不允许插入和删除的一端
// 空栈： 不含如何数据结构的栈
type Stack interface {
	// Init 初始化栈， list 为栈的存储引擎
	Init(list LinearList)

	// Destroy 若栈存在则消除, 该操作会释放栈的存储引擎，若需要复用需要重新初始化
	// this method would release stack storage engineer, reset storage engineer if reused.
	Destroy()

	// Clear 清空栈
	Clear()

	// Empty 是否为空
	Empty() bool

	// Top 返回栈顶元素
	Top() interface{}

	// Push 将元素压栈
	Push(elem interface{})

	// Pop 弹出栈顶元素
	Pop() interface{}

	// Length 长度
	Length() int
}
