// 声明一个greetings包来收集相关功能
package greetings 

import (
	"fmt" 
	"errors"

	"math/rand"
	"time"
)

// 实现一个Hello函数来指定人员的问候语。返回两个值string和error
// 在 Go 中，名称以大写字母开头的函数可以被不在同一个包中的函数调用。
func Hello(name string) (string, error) { 
	    // 如果没有给出名称，则返回带有消息的错误。
			if name == "" { 
        return "", errors.New("empty name") 
    }

    // 返回在消息中嵌入名称的问候语。
		//声明和初始化变量的快捷方式
    // message := fmt.Sprintf("Hi, %v. Welcome!", name) 
		 // 使用随机格式创建消息。
    message := fmt.Sprintf(randomFormat(), name) 
		// // 更改greetings.Hello函数以使其不再包含名称
		// message := fmt.Sprint(randomFormat())
    // 如果接收到名称，则返回一个嵌入该名称的值；nil（意味着没有错误）作为成功返回的第二个值
    return message, nil
}


// Hellos 返回一个地图，将每个命名的人与问候消息相关联。 其参数是名称切片而不是单个名称
func Hellos(names []string) (map[string]string, error) {
	// 将名称与消息关联的映射。
	messages := make(map[string]string) 
	// 循环接收到的名称片段，调用Hello 函数来获取每个名称的消息。
	for _, name := range names { 
		// Hellos函数调用现有 Hello函数
			message, err := Hello(name) 
			if err != nil { 
					return nil, err 
			} 
			// 创建一个messages映射以将每个接收到的名称（作为键）与生成的消息（作为值）相关联
			messages[name] = message 
	}
	return messages, nil
}

// init 为函数中使用的变量设置初始值。
func init() { 
	rand.Seed(time.Now().UnixNano()) 
} 

// 该函数返回随机选择的问候消息格式
func randomFormat() string { 
	// 一段消息格式。声明formats具有三种消息格式的切片
	formats := []string{ 
			"Hi, %v. Welcome!", 
			"Great to see you, %v!", 
			"Hail, %v! Well met!", 
	} 

	// 使用 math/rand包 生成一个随机数，用于从切片中选择一个项目
	return formats[rand.Intn(len(formats))]
}