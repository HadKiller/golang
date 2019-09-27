package main

import (
	"fmt"
	"time"
)

func main() {

	start:=time.Now()
	a:=[]int{34,23,6,8,89,23,9,56,54,43,43,5,5,6}
	//BubbleSort(a)
	InsertSort(a)
	end:=time.Now()
	fmt.Printf("%v/\n",a)
	fmt.Print("时间：",end.Sub(start)/1000)
}

func BubbleSort(a []int)[]int{
	for i:=0;i<len(a); i++{
		for j:=0;j<len(a)-i-1;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	return a
}

func InsertSort(a []int)[]int{
	for i:=1;i<len(a);i++  {
		value:=a[i]
		j:=i-1
		for ;j>=0;j--{
			if a[j]>value{
				a[j+1]=a[j]
			}else {
				break
			}
		}
		a[j+1]=value
	}
	return a
}
// 归并排序算法, A 是数组，n 表示数组大小
//merge_sort(A, n) {
//merge_sort_c(A, 0, n-1)
//}

//// 递归调用函数
//merge_sort_c(A, p, r) {
//// 递归终止条件
//if p >= r  then return
//
//// 取 p 到 r 之间的中间位置 q
//q = (p+r) / 2
//// 分治递归
//merge_sort_c(A, p, q)
//merge_sort_c(A, q+1, r)
//// 将 A[p...q] 和 A[q+1...r] 合并为 A[p...r]
//merge(A[p...r], A[p...q], A[q+1...r])
//}

func MergeSort(a[]int) {
	start:=0
	end:=len(a)
	Merge(a,start,end-1)
}
func Merge(a[] int ,start int ,end int){
	if start>=end{
		return
	}else {
		if a[start]>a[end]{
			a[start],a[end]=a[end],a[start]
		}
	}
}