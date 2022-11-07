package main

import (
	"fmt"
	"time"
)

var (
	hour, minute, second       int
	rehour, reminute, resecond int
	timerNameDelete            string
	timerNameSleep             string
	//存放闹钟运行的切片
	alarmClockSlice = make([]alarmClock, 0)
)

// 闹钟
type alarmClock struct {
	//闹钟名字
	name string
	//闹钟运行多少小时
	startTime map[string]int

	//是否重复性
	isRepeat bool
	//是否关闭
	isClose chan bool
}

// 闹钟运行线程
func alarmClockGoroutine(index int) {
	for{
		if time.Now().Hour()==hour&&time.Now().Minute()==minute||time.Now().Second()==second  {
			break
		}
	}
	alarmClock := alarmClockSlice[index]
	fmt.Printf("\n%v号闹钟开始运行，闹钟名字：%v | 闹钟开始时间：%v小时%v分%v秒 | 闹钟是否重复：%v\n", index, alarmClock.startTime["hour"], alarmClock.startTime["minute"], alarmClock.startTime["second"], alarmClock.isRepeat)

	//开始延迟X小时
	for {
		select {
		case isClose := <-alarmClock.isClose:
			fmt.Println(isClose)
			fmt.Printf("\n【%v号闹钟已被手动关闭，闹钟名字：%v】\n", index, alarmClock.name)
			return
		case <-time.After(time.Hour * time.Duration(rehour*3600+reminute*60+resecond)):
			fmt.Printf("\n【%v号闹钟已到期，闹钟名字：%v】\n", index, alarmClock.name)
			return
		}
	}
}

// 分发闹钟
func distributionAlarmClock(ac alarmClock) {
	//把闹钟信息加入闹钟切片中，并获得在切片中位置
	alarmClockSlice = append(alarmClockSlice, ac)
	index := len(alarmClockSlice) - 1
	//开始执行闹钟
	go alarmClockGoroutine(index)
}

// 添加闹钟菜单
func addMenu(choose int) {
	fmt.Printf("请输入你要自定义的闹钟名字：")
	var name string
	// var fun string
	var hour, minute, second int
	fmt.Scanln(&name)
	fmt.Printf("请输入你要自定义的闹钟启动时间秒为某时某分某,用空格隔开")
	fmt.Scan(&hour, &minute, &second)

	var isRepeat bool = false
	if choose == 2 {
		isRepeat = true
		fmt.Printf("请输入你要自定义的闹钟重复间隔时间,用空格隔开")
		fmt.Scan(&rehour, &reminute, &resecond)
	}
	st := map[string]int{
		"hour":   hour,
		"minute": minute,
		"second": second,
	}
	distributionAlarmClock(alarmClock{
		name:      name,
		startTime: st,
		isRepeat:  isRepeat,
		isClose:   make(chan bool),
	})
}

// 关闭闹钟，查找闹钟名字然后关闭
func closeClockSlice(name string) {
	for _, value := range alarmClockSlice {
		if value.name == name {
			value.isClose <- true
		}
	}
}

// 菜单
func menu() {
	for {
		fmt.Println("\n=============================")
		fmt.Println("请做出你的选择")
		fmt.Println("1.我要自定义我的一次性闹钟")
		fmt.Println("2.我要自定义我的重复性闹钟")
		fmt.Println("3.我要删除某闹钟")
		fmt.Println("4.我要取消某闹钟下一次的提醒")
		fmt.Printf("请输入编号(1-4)：")
		var choose int
		fmt.Scanln(&choose)
		if choose == 1 || choose == 2 {
			addMenu(choose)
		}
		if choose == 3 {
			fmt.Printf("请输入你要删除闹钟的名字：")
			fmt.Scanln(&timerNameDelete)
			closeClockSlice(timerNameDelete)
		}

	}
}

func main() {
	for {
		fmt.Println("请做出你的选择")
		fmt.Println(" - 1，开始")
		fmt.Println(" - 2，退出")
		fmt.Printf("请输入编号(1或2)：")
		var or int
		fmt.Scanln(&or)
		if or == 2 {
			return
		}
		if or < 1 || or > 2 {
			fmt.Printf("***输入错误，请重新输入!!!!")
			fmt.Println("\n=============================")
			continue
		}
		menu()
	}

}
