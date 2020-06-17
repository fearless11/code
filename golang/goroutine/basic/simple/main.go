// 创建goroutine以及调度器行为

///////////// 两个goroutine并发运行，显示字母表3次
package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
运行：go run main.go
output：
start Goroutines
waiting to finish
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z
Terminating program

第一个goroutine完成时间太短，以至于在调度器切换到第二个goroutine前就完成所有任务。
这是为什么会先看到所有大写字母后才输出小写字母。
*/

func main() {
	// 分配一个逻辑处理器
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Goroutines")

	// 匿名函数
	go func() {
		// 函数退出时调用Done通知main函数工作已完成
		defer wg.Done()
		// 显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("waiting to finish")
	// 等待goroutine结束
	wg.Wait()

	fmt.Println("\nTerminating program")
}
