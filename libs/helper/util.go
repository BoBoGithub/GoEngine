package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"time"
)

/**
 * 检查字符串 是否是数字格式
 */
func CheckInt(str string) bool {
	//正则验证参数
	result, err := regexp.MatchString("^\\d+$", str)
	if err != nil {
		return false
	}

	return result
}

//数字字符串转int
func StringToInt(num string) int {
	//转换类型
	intNum, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}

	return intNum
}

//检查字符串 是否是手机号格式
func CheckMobile(str string) bool {
	//正则验证参数
	result, err := regexp.MatchString("^[1]([3-9])[0-9]{9}$", str)
	if err != nil {
		return false
	}

	return result
}

//结构体转
func StructToMap(obj interface{}, isPureStruct bool) map[string]interface{} {
	//如果是纯转结构体
	if isPureStruct {
		//反射获取结构体类型和值
		t := reflect.TypeOf(obj)
		v := reflect.ValueOf(obj)

		//转换
		data := map[string]interface{}{}
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}

		return data
	}

	//以结构体的json格式名称转map
	m := make(map[string]interface{})
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)
	return m
}

//排序map
func GoKSort(param map[string]interface{}) map[string]interface{} {
	//检查参数
	if len(param) == 0 {
		return param
	}

	//提取map的key
	keys := []string{}
	for k, _ := range param {
		keys = append(keys, k)
	}

	//排序key
	sort.Strings(keys)

	//设置返回结果
	sortRet := map[string]interface{}{}
	for _, k := range keys {
		sortRet[k] = param[k]
	}

	return sortRet
}

//Md5加密字符串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//获取当前时间
func GetDateTimeFromTimeStamp(timestamp int64, formate string) string {
	return time.Unix(timestamp, 0).Format(formate)
}

//AES - CBC模式 - 加密函数
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

//AES - CBC模式 - 解密函数
func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 检查数据长度
	if len(crytedByte)%blockSize != 0 {
		return ""
	}

	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
