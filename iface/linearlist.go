package iface

// linearList 线性表
// 线性表时一种抽像数据结构，他是0个或多个数据元素的有序序列
// 线性表可以有链式存储结构和顺序存储结构
type LinearList interface {
	// Init 初始化线性表
	// initial linear list
	// comparer is used to compare elem, typ specifies the type which can put in
	Init(comparer func(inode1, inode2 interface{}) int, typ interface{})

	// Empty 判断线性表是否为空
	// returns true if linear list is empty
	Empty() bool

	// Clear 清空线性表
	// clear linear list
	Clear()

	// GetElem 返回指定位置的元素
	// returns elem indexes of i
	GetElem(i int) interface{}

	// LocateElem 查找与指定元素e相等的元素
	// returns the index of elem e if find it
	// returns 0 if not found
	// returns negative nums if type unmatched
	LocateElem(elem interface{}) int

	// Insert 在指定位置插入元素elem
	// insert elem into linear list by index i
	Insert(i int, elem interface{}) error

	// Delete 删除位置为i的元素，并返回被删除的元素
	// returns the deleted elem whose index is i
	Delete(i int) interface{}

	// Length 返回线性表的长度
	// returns the length of linear list
	Length() int
}