/**
 * 当前模块使用的结构体定义
 *
 * @param addTime 2022-05-06
 * @param author  ChengBo
 */
package entity

import "strconv"

//定义加密信息接口体
type EncryptionUserInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

//定义首页需要展示的数据
type UserGameItem struct {
	PsnId        string `json:"psn_id"`
	AppNPCId     string `json:"app_npc_id"`
	UserProgress int    `json:"user_progress"`
	UserPlatinum int    `json:"user_platinum"`
	UserGold     int    `json:"user_gold"`
	UserSilver   int    `json:"user_silver"`
	UserBronze   int    `json:"user_bronze"`
	PlayTime     int    `json:"play_time"`
	GameName     string `json:"app_name"`
	GamePic      string `json:"app_icon"`
}

//定义游戏数据结构体
type GameInfo struct {
	AppId       int    `json:"app_id"`
	AppNPCId    string `json:"app_npc_id"`
	AppName     string `json:"app_name"`
	AppDesc     string `json:"app_desc"`
	AppIcon     string `json:"app_icon"`
	AppPlatinum int    `json:"app_platinum"`
	AppGold     int    `json:"app_gold"`
	AppSilver   int    `json:"app_silver"`
	AppBronze   int    `json:"app_bronze"`
}

//设置配置文件结构体
type Config struct {
	MySQLConfig          []DBConfig `yaml:"MySQL"`
	EncryptionPrivateKey string     `yaml:"EncryptionPrivateKey"`
}

//数据库实例结构
type DBSource struct {
	BBS  DBConfig `yaml:"BBS"`
	Data DBConfig `yaml:"Data"`
}

//MySQL相关配置信息
type DBConfig struct {
	Key      string `yaml:"key"`
	Dirver   string `yaml:"dirver"`
	Host     string `yaml:"host"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	CharSet  string `yaml:"charset"`
	DataBase string `yaml:"database"`
}

//自定义json序列化格式
type EscapeString string

func (esc EscapeString) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(esc))), nil
}
