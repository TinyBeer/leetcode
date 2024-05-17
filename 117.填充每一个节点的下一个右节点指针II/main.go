package main

/*
给定一个二叉树：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL 。
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层序遍历
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	q := []*Node{root}
// 	for len(q) != 0 {
// 		var pre *Node
// 		for size := len(q); size > 0; size-- {
// 			cur := q[0]
// 			q = q[1:]
// 			if pre != nil {
// 				pre.Next = cur
// 			}
// 			pre = cur
// 			if cur.Left != nil {
// 				q = append(q, cur.Left)
// 			}
// 			if cur.Right != nil {
// 				q = append(q, cur.Right)
// 			}
// 		}
// 	}
// 	return root
// }

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
