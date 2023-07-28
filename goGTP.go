package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {
		fmt.Println("请输入你的问题::")
		var name string
		fmt.Scanf("%s", &name)
		// 设置随机种子
		rand.Seed(time.Now().UnixNano())
		// 生成随机整数
		randomInt := rand.Intn(4) // 生成0到4之间的随机整数
		switch randomInt {
		case 0:
			fmt.Println("不会!")
		case 1:
			fmt.Println("不知道!")
		case 2:
			fmt.Println("你真好看,我还是不会!")
		default:
			fmt.Println("你人还怪好的嘞,去百度查查!!!")
		}
	}

}
