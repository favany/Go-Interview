package main

// https://www.jianshu.com/p/f0c425d9a0a5

import "fmt"

// 定义猫猫的结构体

type Node struct {
	no   int // 猫猫的编号
	name string
	next *Node
}

func InsertNode(head *Node, newNode *Node) {
	// 判断是不是添加第一个节点
	if head.next == nil {
		head.next = newNode
		newNode.next = head
		return
	}
	// 定义一个临时的变量
	flag := head
	// 找到最后一个变量
	for flag.next != head {
		flag = flag.next
	}
	flag.next = newNode
	newNode.next = head
}

// 输出环形链表

func ListLink(head *Node) {
	fmt.Println("环形链表的情况如下:")
	flag := head
	if flag.next == nil {
		fmt.Println("链表为空")
		return
	}
	// 环形链表中循环，直到下一个节点是头节点，代表循环结束♻️
	for {
		fmt.Println("猫的信息为第", flag.no, "只猫，", "名字叫做", flag.name)
		if flag.next != head {
			flag = flag.next
		}
	}
}

func main() {

	// 先初始化一个环形链表的头节点
	head := &Node{}

	// 创建一只猫
	cat1 := &Node{
		no:   1,
		name: "tom",
	}

	InsertNode(head, cat1)

	ListLink(head)
}
