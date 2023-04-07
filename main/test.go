package main

import (
	"algor_lib/charpool"
)

func main() {
	pool := `车这忽看记北壮道座窝胆微从当信地故窗鸟铅孤乡往棵赶应乐想香笔飞学快说思床支友京夜边毛战架面士因开走很单偏好非太运打在匹敌
秋该音直爸鱼叫告送居外乱动册阳给怎羽戏种诉多独后暖篮身散自以校跳淹会叽井忙为捉色常招他像温忘蓝玩听觉静喜安江跟前声试没得许再喊敢晚笑游来
辆然绳讲举是排块急鹊足村要低原广场机席主望吃发观少只呼树有就街金喳球连死疑勇事踢水河光睡颜哥亲尝甜住行都邻也`
	newChar := ``
	charpool.NewCharPool(pool).AddChars(newChar)
}
