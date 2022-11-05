package main

import (
	"fmt"
	"strings"
)

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
func mySkill(newM map[string]string, s *string) {
	var myskillNames, myskillFun string
	for {
		fmt.Println("请输入你的自定义技能名称")
		fmt.Scanln(&myskillNames)
		if !senWord(myskillNames) {
			break
		}
	}
	for {
		fmt.Println("请输入你的自定义技能作用")
		fmt.Scanln(&myskillFun)
		if !senWord(myskillFun) {
			break
		}
	}
	newM[myskillNames] = myskillFun
	*s = myskillNames
}
func senWord(word string) bool {
	wordHouse := "国家政治暴力黄色自杀"
	var ret bool

	ret = strings.Contains(wordHouse, word)
	if ret {
		fmt.Println("抱歉你输入了敏感词，请重新输入")

	}

	return ret
}
func main() {

	fmt.Println("请输入你准备释放的技能")
	fmt.Println("龟派气功")
	fmt.Println("元气弹")
	fmt.Println("气圆斩")
	fmt.Println("界王拳")
	fmt.Println("魔封波")
	fmt.Println("自在极意功")
	fmt.Println("不满意？输入“我要玩别的”开启你的自定义技能")
	fmt.Println("输入“我不玩了”即可退出")
	m := map[string]string{
		"龟派气功":  "释放龟派气功",
		"元气弹":   "释放元气弹",
		"气圆斩":   "释放气圆斩",
		"界王拳":   "释放界王拳",
		"魔封波":   "释放魔封波",
		"自在极意功": "释放自在极意功",
	}
	var skillNames string
	_, err := fmt.Scanln(&skillNames)
	if err != nil {
		fmt.Printf("输入错误=%v", err)
	}

	if skillNames == "我要玩别的" {
		mySkill(m, &skillNames)
	}
	if skillNames == "我不玩了" {
		fmt.Println("拜拜您嘞")
		return
	}
	ReleaseSkill(skillNames, func(skillNames string) {
		fmt.Println(m[skillNames])
	})
}
