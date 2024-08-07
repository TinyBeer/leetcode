package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
*/

// type Codec struct{}

// func Constructor() (_ Codec) {
//     return
// }

// func (Codec) serialize(root *TreeNode) string {
//     sb := &strings.Builder{}
//     var dfs func(*TreeNode)
//     dfs = func(node *TreeNode) {
//         if node == nil {
//             sb.WriteString("null,")
//             return
//         }
//         sb.WriteString(strconv.Itoa(node.Val))
//         sb.WriteByte(',')
//         dfs(node.Left)
//         dfs(node.Right)
//     }
//     dfs(root)
//     return sb.String()
// }

// func (Codec) deserialize(data string) *TreeNode {
//     sp := strings.Split(data, ",")
//     var build func() *TreeNode
//     build = func() *TreeNode {
//         if sp[0] == "null" {
//             sp = sp[1:]
//             return nil
//         }
//         val, _ := strconv.Atoi(sp[0])
//         sp = sp[1:]
//         return &TreeNode{val, build(), build()}
//     }
//     return build()
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "n"
	}
	return fmt.Sprintf("%s,%s,%s", strconv.Itoa(root.Val), this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	i := 0
	var d func() *TreeNode
	d = func() *TreeNode {
		if nodes[i] == "n" {
			return nil
		}
		val, _ := strconv.Atoi(nodes[i])
		n := &TreeNode{Val: val}
		i++
		n.Left = d()
		i++
		n.Right = d()
		return n
	}
	return d()
}
