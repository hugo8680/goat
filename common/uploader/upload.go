package uploader

import (
	"encoding/base64"
	"errors"
	"github.com/hugo8680/goat/common/serializer/datetime"
	"github.com/hugo8680/goat/framework/config"
	"math/rand"
	"net/textproto"
	"os"
	"strings"
	"time"
)

// Uploader 上传文件
type Uploader struct {
	Config *UploadConfig
	File   *File
}

var (
	UploadLocalDriver = "local"
	UploadOssDriver   = "oss"
)

type UploadOption func(*UploadConfig)

// UploadConfig 上传配置
type UploadConfig struct {
	Driver     string   // 上传驱动
	SavePath   string   // 保存路径
	UrlPath    string   // 访问地址路径
	LimitSize  int      // 限制文件大小
	LimitType  []string // 限制文件类型
	RandomName bool     // 使用随机文件名
}

// File 文件信息
type File struct {
	FileName    string               // 文件名
	FileSize    int                  // 文件大小
	FileType    string               // 文件类型
	FileHeader  textproto.MIMEHeader // 文件头
	FileContent []byte               // 文件内容
}

// Result 返回结果
type Result struct {
	OriginalName string `json:"originalName"`
	FileName     string `json:"fileName"`
	FileSize     int    `json:"fileSize"`
	FileType     string `json:"fileType"`
	SavePath     string `json:"savePath"`
	UrlPath      string `json:"urlPath"`
	Url          string `json:"url"`
}

var setting = config.GetSetting()

// NewUploader 初始化上传对象
func NewUploader(options ...UploadOption) *Uploader {
	todayPath := time.Now().Format(datetime.DATE_FORMAT1) + "/"
	// 配置默认驱动
	c := &UploadConfig{
		Driver:     UploadLocalDriver,
		UrlPath:    setting.System.UploadPath + todayPath,
		SavePath:   setting.System.UploadPath + todayPath,
		RandomName: false,
	}
	for _, option := range options {
		option(c)
	}
	return &Uploader{
		Config: c,
	}
}

// SetDriver 设置上传驱动
func SetDriver(driver string) UploadOption {
	return func(config *UploadConfig) {
		config.Driver = driver
	}
}

// SetSavePath 设置保存路径
func SetSavePath(savePath string) UploadOption {
	return func(config *UploadConfig) {
		config.SavePath = savePath
	}
}

// SetUrlPath 设置访问地址路径
func SetUrlPath(urlPath string) UploadOption {

	return func(config *UploadConfig) {
		config.UrlPath = urlPath
	}
}

// SetLimitSize 设置限制文件大小
func SetLimitSize(limitSize int) UploadOption {
	return func(config *UploadConfig) {
		config.LimitSize = limitSize
	}
}

// SetLimitType 设置限制文件类型
func SetLimitType(limitType []string) UploadOption {
	return func(config *UploadConfig) {
		config.LimitType = limitType
	}
}

// SetRandomName 使用随机文件名
func SetRandomName(isRandomName bool) UploadOption {
	return func(config *UploadConfig) {
		config.RandomName = isRandomName
	}
}

// SetFile 设置上传文件
func (u *Uploader) SetFile(file *File) *Uploader {
	u.File = file
	return u
}

// Save 保存文件
func (u *Uploader) Save() (*Result, error) {
	var err error
	if setting.System.Host == "" {
		return nil, errors.New("未找到域名，无法生成访问地址")
	}
	domain := setting.System.Host
	if u.File == nil || len(u.File.FileContent) <= 0 {
		return nil, errors.New("上传文件数据不全，无法保存")
	}
	// 获取文件后缀并且生成hash文件名
	fileName := strings.Split(u.File.FileName, ".")
	if len(fileName) != 2 {
		return nil, errors.New("文件缺少后缀")
	}
	// 拼接随机文件名
	randomName := u.File.FileName
	if u.Config.RandomName {
		randomName = u.generateRandomName() + "." + fileName[1]
	}
	if err = u.checkLimitSize(); err != nil {
		return nil, err
	}
	if err = u.checkLimitType(); err != nil {
		return nil, err
	}
	switch u.Config.Driver {
	case UploadLocalDriver:
		err = u.saveToLocal(randomName)
	case UploadOssDriver:
		err = u.saveToOss()
	default:
		err = u.saveToLocal(randomName)
	}
	if err != nil {
		return nil, err
	}
	return &Result{
		OriginalName: u.File.FileName,
		FileName:     randomName,
		FileSize:     u.File.FileSize,
		FileType:     u.File.FileType,
		SavePath:     u.Config.SavePath,
		UrlPath:      u.Config.UrlPath,
		Url:          domain + "/" + u.Config.UrlPath + randomName,
	}, err
}

// 检查文件大小
func (u *Uploader) checkLimitSize() error {
	if u.Config.LimitSize > 0 && u.File.FileSize > 0 && u.Config.LimitSize < u.File.FileSize {
		return errors.New("文件大小超出限制")
	}
	return nil
}

// 检查文件类型
func (u *Uploader) checkLimitType() error {
	if len(u.Config.LimitType) <= 0 || u.File.FileType == "" {
		return nil
	}
	for _, limitType := range u.Config.LimitType {
		if limitType == u.File.FileType {
			return nil
		}
	}
	return errors.New("文件格式不合法")
}

// 生成随机字符串
func (u *Uploader) generateRandomName() string {
	// 创建一个新的随机数生成器实例
	r := rand.New(rand.NewSource(int64(len(base64.StdEncoding.EncodeToString([]byte(u.File.FileName))))))
	// 定义可能的字符集，包括字母和数字
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 生成随机字符串
	var randomName string
	for i := 0; i < 64; i++ {
		// 从字符集中随机选择一个字符
		randomChar := chars[r.Intn(len(chars))]
		randomName = randomName + string(randomChar)
	}
	return randomName
}

// 保存到本地
func (u *Uploader) saveToLocal(randomName string) error {
	if _, err := os.Stat(u.Config.SavePath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(u.Config.SavePath, 0644); err != nil {
				return err
			}
		}
	}
	return os.WriteFile(u.Config.SavePath+randomName, u.File.FileContent, 0644)
}

// 保存到Oss
func (u *Uploader) saveToOss() error {

	// TODO

	return nil
}
