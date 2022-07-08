/**
 * 首页服务实现
 *
 * @param addTime 2022-05-06
 * @param author  BoBo
 */
package service

import (
	"GoEngine/entity"
	"GoEngine/libs/helper"
	"GoEngine/libs/response"
	"GoEngine/model"
)

//提取首页展示数据
func GetIndexShowData() map[string]interface{} {
	//返回 首页数据配置
	return map[string]interface{}{
		"title": "GoEngine",
		"name":  "GoEngine",
		"desc":  "玩游戏是平衡生活中一个有趣的部分。",
	}
}

//提取用户游戏数据
func GetUserGameData(userId string) []entity.UserGameItem {
	//检查请求参数
	if !helper.CheckInt(userId) {
		//抛出用户信息错误异常
		response.ThrowException("2000 user hasn't login in")
	}

	//查询用户的游戏信息
	psnGameList := model.GetUserPsnGameByUid(userId)
	if len(psnGameList) == 0 {
		return []entity.UserGameItem{}
	}

	//遍历游戏数据提取 游戏的NPCID
	npcId := []string{}
	for _, game := range psnGameList {
		//提取游戏的npcid
		npcId = append(npcId, game.AppNPCId)
	}

	//提取游戏名称信息
	gameInfos := GetGameInfoByNPCId(npcId)

	//设置游戏名称及图标
	for key, game := range psnGameList {
		//设置游戏信息
		if gameInfo, ok := gameInfos[game.AppNPCId]; ok {
			//设置游戏名称
			psnGameList[key].GameName = gameInfo.AppName

			//设置游戏图标
			psnGameList[key].GamePic = gameInfo.AppIcon
		}
	}

	//返回用户的游戏数据
	return psnGameList
}

//通过NPCID提取游戏信息
func GetGameInfoByNPCId(npcIds []string) map[string]entity.GameInfo {
	//定义返回结果变量
	gameInfos := map[string]entity.GameInfo{}

	//检查参数
	if len(npcIds) == 0 {
		return gameInfos
	}

	//查询游戏数据
	gameListData := model.GetGameInfoByNpcIds(npcIds)

	//提取返回值数据
	for _, info := range gameListData {
		//设置返回值
		gameInfos[info.AppNPCId] = info
	}

	//返回游戏信息
	return gameInfos
}
