package main

import (
	"fmt"
	"sync"
	"time"
)

var timerNameDelete string
var timerNameSleep string
var wg sync.WaitGroup

func Timer(na string, hour int, fun string) {
	name := na
	if name == timerNameSleep {
		time.Sleep(time.Hour * 24)
		fmt.Println("取消下一次提醒成功")

	} //直接休眠24小时以取消下一次提醒
	var is bool

	for {
		if time.Now().Hour() == hour {
			is = true
			break
		}
	}
	if is {
		ticker := time.NewTicker(time.Hour * 24)
		if name == timerNameDelete {
			ticker.Stop()
			fmt.Println("删除成功")

		} //删除
		for {
			select {
			case <-ticker.C:
				fmt.Println("晚上悄悄学习，惊艳所有人")
			}
		}
	}

}
func Timer3() {
	name := "timer3"
	if name == timerNameSleep {
		time.Sleep(time.Hour * 24)
		fmt.Println("取消下一次提醒成功")

	} //直接休眠24小时以取消下一次提醒
	ticker := time.NewTicker(time.Second * 30)
	if name == timerNameDelete {
		ticker.Stop()
		fmt.Println("删除成功")

	} //删除
	for {
		select {
		case <-ticker.C:
			fmt.Printf("我还能继续卷")
		}
	}

}
func add(c int) {
	fmt.Println("请输入你要自定义的闹钟名字")
	var name, fun string
	var getTime int
	fmt.Scanln(&name)
	fmt.Println("请输入你要自定义的闹钟启动时间为某某时")
	fmt.Scanln(&getTime)
	fmt.Println("请输入你要自定义的闹钟作用")
	fmt.Scanln(&fun)
	var is bool

	if name == timerNameSleep {
		time.Sleep(time.Hour * 24)
		fmt.Println("取消下一次提醒成功")

	} //直接休眠24小时以取消下一次提醒
	for {
		if time.Now().Hour() == getTime {
			is = true
			break
		}
	}
	if is {
		ticker := time.NewTicker(time.Hour * 24)
		if name == timerNameDelete {
			ticker.Stop()
			fmt.Println("删除成功")

		} //删除
		for {
			select {
			case <-ticker.C:
				fmt.Println(fun)
				if c == 1 {
					ticker.Stop() //直接关了
				}
			}
		}
	}

}
func main() {
	fmt.Println("请做出你的选择")
	fmt.Println("1，开始")
	fmt.Println("2，退出")
	var or int
	fmt.Scanln(&or)
	if or == 2 {
		return
	}
	if or == 1 {
		fmt.Println("请做出你的选择")
		fmt.Println("1.我要自定义我的一次性闹钟")
		fmt.Println("2.我要自定义我的重复性闹钟")
		fmt.Println("3.我要删除某闹钟")
		fmt.Println("4.我要取消某闹钟下一次的提醒")
		var choose int
		fmt.Scanln(&choose)
		if choose == 1 || choose == 2 {
			go add(choose)
		}
		if choose == 3 {
			fmt.Println("请输入你要删除闹钟的名字")
			fmt.Scanln(&timerNameDelete)
			fmt.Println("你输入的闹钟名字为", timerNameDelete)
		}
		if choose == 4 {
			fmt.Println("请输入你要取消下一次提醒闹钟的名字")
			fmt.Scanln(&timerNameSleep)
		}
		go Timer("timer1", 0, "晚上悄悄学习，惊艳所有人")
		go Timer("timer2", 6, "起床开卷")
		go Timer3()
		time.Sleep(time.Second * 3)
		main()
	}
}
