package main

import "fmt"

type ProductOfNumbers struct {
	val []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{val: make([]int, 0, 256)}
}

func (this *ProductOfNumbers) Add(num int) {

	this.val = append(this.val, num)

}

func (this *ProductOfNumbers) GetProduct(k int) int {

	return func(num []int) int {

		n := num[len(num)-1]

		for in := 0; in < len(num)-1; in++ {
			n *= num[in]
		}
		return n
	}(this.val[len(this.val)-k:])
}

func main() {
	obj := Constructor()

	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)
	fmt.Println(obj)
	fmt.Println(obj.GetProduct(1))
}
