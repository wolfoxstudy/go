package main

import "fmt"
func main(){
    // var count,c int   //定义变量不使用也会报错v
    var count int64
var tmp int64    
var flag bool
    count=1
    //while(count<100) {    //go没有while
    for count < 100000000000000 {
        count++
        flag = true;
        //注意tmp变量  :=
        for tmp=2;tmp<count;tmp++ {
            if count%tmp==0{
                flag = false
            }
        }

        // 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
        if flag == false {
            fmt.Println(count,"合数")
        }else{
        continue
        }
    }
}
