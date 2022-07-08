package model

import (
	"GoEngine/entity"
	"GoEngine/libs/db"
	"fmt"
	"strings"
)

//提取用户游戏数据列表
func GetUserPsnGameByUid(userId string) []entity.UserGameItem {
	//设置返回值
	userGames := []entity.UserGameItem{}

	//设置查询字段
	searchFileds := "psn_id, app_npc_id, user_progress, user_platinum, user_gold, user_silver, user_bronze, play_time"

	//提取数据库链接句柄
	db.GetDbLink("bbs").Table("pw_psn_user_game").Select(searchFileds).Where("psn_id = ?", userId).Scan(&userGames)

	//返回游戏数据
	return userGames
}

//通过NPCId获取游戏信息
func GetGameInfoByNpcIds(npcIds []string) []entity.GameInfo {
	//设置返回结果变量
	gameInfos := []entity.GameInfo{}

	//设置查询字段
	searchFileds := "app_id, app_npc_id, app_name, app_desc, app_icon, app_platinum, app_gold, app_silver, app_bronze"

	//设置查询条件
	where := fmt.Sprintf("app_npc_id IN ('%s')", strings.Join(npcIds, "','"))

	//查询数据
	db.GetDbLink("bbs").Table("pw_psn_game").Select(searchFileds).Where(where).Scan(&gameInfos)

	//返回游戏数据
	return gameInfos
}
