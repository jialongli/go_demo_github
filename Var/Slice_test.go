package Var

import (
	"testing"
)

/**
切片的初始化,注意:append的时候,需要重新赋值回给原来的切片.
为什么????
原因: !!!!!如果扩容的话,go中,是将原来的数据拷贝到一个新的数组,再把指针指向新的数组!!!!!
*/
func TestSlice(t *testing.T) {
	//直接创建了1,2
	slice1 := []int{1, 2}
	t.Log(len(slice1), cap(slice1))

	slice2 := make([]int, 3, 4)
	t.Log(len(slice2), cap(slice2))
	for i := 0; i < 5; i++ {
		slice2 = append(slice2, 1)
		t.Log(len(slice2), cap(slice2))
	}
}

func TestSliceSplit(t *testing.T) {
	sliceOld := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //len:9 cap:9

	//3:5代表要获取的起始和终止的索引,前包含,后不包含
	//注意: 切片底层是指向的一个连续数组,所以slice1指向的和sliceOld是同一个数组空间,只不过slice1的可视空间是4,5  长度为[4,5,6,7,8,9]=6
	slice1 := sliceOld[3:5]     //5,6
	t.Log(slice1[0], slice1[1]) //打印4,5
	//t.Log(slice1[0],slice1[1],slice1[2])  //抛异常,slice[2]越界
	t.Log(len(slice1), cap(slice1)) //打印len=2,cap=6

	slice2 := sliceOld[3:7]                           //4,5,6,7  lenth为6 4,5,6,7,8,9
	t.Log(slice2[0], slice2[1], slice2[2], slice2[3]) //打印4,5,6,7
	t.Log(len(slice2), cap(slice2))                   //打印len=2,cap=6

	//注意这里,我从slice2([4,5,6,7],8,9)截取了 [5,6,7,8],9     []里的为数据长度,就是可以获取到的 len/  整个为cap,即容量
	slice3 := slice2[1:5]
	t.Log(slice3[0], slice3[1], slice3[2], slice3[3]) //打印5,6,7,8
	t.Log(len(slice3), cap(slice3))                   //打印len=4,cap=5

	slice4 := sliceOld[3:5]         //[4,5],6,7,8,9  len:2 cap:6
	t.Log(len(slice4), cap(slice4)) //打印len=2,cap=6
	//!!!!注意这里append了10,那么就会将slice4的可见的元素的下一个元素设置为10,即 [4,5,10],7,8,9 对应的原来的数组,也都会变化
	slice4 = append(slice4, 10)
	t.Log(slice4[0], slice4[1], slice4[2]) //打印4,5,10
	t.Log(len(slice4), cap(slice4))        //打印len=3,cap=6

	t.Log(len(sliceOld), cap(sliceOld))                                                                                        //打印len=9,cap=9 由于slice4自作主张将6改为了10,会影响所有的引用
	t.Log(sliceOld[0], sliceOld[1], sliceOld[2], sliceOld[3], sliceOld[4], sliceOld[5], sliceOld[6], sliceOld[7], sliceOld[8]) // 1 2 3 4 5 10 7 8 9

	//测试当扩容时,会赋值,并将引用修改
	slice5 := sliceOld[3:5]     //[4,5],10,7,8,9  len:2 cap:6
	slice5 = append(slice5, 55) //[4,5,55],7,8,9
	slice5 = append(slice5, 55) //[4,5,55,55],8,9
	slice5 = append(slice5, 55) //[4,5,55,55,55],9
	slice5 = append(slice5, 55) //[4,5,55,55,55,55]
	//到这里,slice5指向的数组已经被修改了,其他的也是这样
	t.Log(slice4[0], slice4[1], slice4[2]) //打印4,5,55

	slice5 = append(slice5, 55)            //[4,5,55,55,55,55,55]  此时,发生扩容,会新建一个数组,并修改指针
	t.Log(slice5[0], slice5[1], slice5[6]) //打印4,5,55
	//再次修改,不会影响原有数组
	slice5[0] = 12
	t.Log(slice5[0], slice5[1], slice5[2]) //打印12,5,55
	t.Log(slice4[0], slice4[1], slice4[2]) //打印4,5,55 没有被影响

}

func Testcompare(t *testing.T) {

}
