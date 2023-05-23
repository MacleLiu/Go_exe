package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

//String方法展示tree的值序列
func (t *tree) String() string {
	res := ""
	if t == nil {
		return res
	}
	res += t.left.String()
	res = fmt.Sprintf("%s %d", res, t.value)
	res += t.right.String()
	return res
}

//buildTree函数创建一棵树
func buildTree(data []int) *tree {
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	return root
}

// 就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues将元素按照顺序追加到values里面，然后返回结果slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		//等价于返回&tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	num := []int{23, 4, 5, 6, 21, 67, 6, 234, 6, 234, 78, 24}
	t := buildTree(num)
	Sort(num)
	fmt.Println(num)
	fmt.Println(t.String())
}
