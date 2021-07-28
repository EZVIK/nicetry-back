package main

import (
	"fmt"
	"strings"
)

func main() {

	const s = `本期登场人物：@realEnjolras @文森特动物园 @特梨西 @cbvivi

本期编辑：@特梨西

本期你可能会听到：

小 E 自己安装了Carplay

文森特看了乐夏，马上喜欢了几个乐队，而且觉得乐夏时长很良心

文森特看了东京爱情故事

文森特不清楚自己的快乐排名

文森特手机放 Radiohead 放到没电，但仍念不来高中最爱曲目

文森特哥哥的同学算是镇上比较前卫的小孩

中午十二点，打开县里的电台，我会给你点一首 B 安的《真的爱你》

cbvivi 买了新表带

cbvivi 的 5000 万英镑之夜

trtr 完美一周 x2

受 Apple Watch 之邀，呼吸

上周挑战：玩一个没玩过的游戏

cbvivi：FIFA 21

小 E：极乐迪士科

文森特： 《森喜刚：热带冻结》

trtr：Untitled Goose Game

本周挑战：每天睡够7小时`

	const ss = `Random Names
Ianis Esparza
Jed Brady
Chace Dunne
Wiktor Morgan
Tomas Holmes
Emmett Mcclain
Cherise Lee
Sanjay Vo
Mai Mcpherson
Nida Medrano`

	//fmt.Printf(strings.Join(strings.Split(ss, "\n"), `\n`))
	ggg := strings.Split(ss, "\n")
	fmt.Println(ggg)
}
