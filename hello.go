// 声明一个main包（包是对函数进行分组的一种方式，它由同一目录中的所有文件组成）
package main

// 导入流行的 fmt包，其中包含格式化文本的功能，包括打印到控制台。这个包是 你安装 Go 时获得 的标准库包之一
import "fmt"

// 使用外部包rsc.io/quote
import "rsc.io/quote"

import ( 
// 导入项目另一个模块
	"hello-go/greetings" 
	"log"
) 
// 实现一个main功能以将消息打印到控制台。main运行main包 时默认执行一个函数
func main() {
    fmt.Println("Hello, World!")  // 打印字符串
		fmt.Println(quote.Go()) // 打印Go函数的调用

		// 设置预定义 Logger 的属性，包括日志条目前缀和禁用打印的标志、时间、源文件和行号。
		log.SetPrefix("greetings: ")
    log.SetFlags(0)
		
		 // greetings通过调用包的 Hello函数获取问候消息并打印
		 message, err := greetings.Hello("刷我的卡") 
		 // 如果返回错误，将其打印到控制台并退出程序
		 if err != nil {
			log.Fatal(err)  // log包的 Fatal功能 来打印错误并停止程序
		 }
		 fmt.Println(message) 

		 
    // 一个名称切片。
    names := []string{"茄子", "西红柿", "鸡蛋"}
		// 请求名字的问候信息
		messages2, err2 := greetings.Hellos(names)
    // 请求名字的问候信息。
		if err2 != nil {
				log.Fatal(err2)
		}
		fmt.Println(messages2)

		
		 message9, err9 := greetings.Hello("")
		 if err9 != nil {
			log.Fatal(err9)
		 }

		 fmt.Println(message9)

}