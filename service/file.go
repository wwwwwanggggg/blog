package service

import (
	"blog/common"
	"blog/config"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type FileManagement struct{}

// 接受新文件
func (fm *FileManagement) UploadFile(file *multipart.FileHeader) (err error) {
	f, err := file.Open()
	if err != nil {
		return common.ErrNew(errors.New("打开文件失败"), common.SysErr)
	}
	resFile, err := os.Create(fmt.Sprintf("%s/%s", config.Config.Root, file.Filename)) // 这一步已经创建文件了
	defer func() {
		_ = resFile.Close()
	}()
	if err != nil {
		return common.ErrNew(errors.New("服务器创建文件失败"), common.SysErr)
	}
	io.Copy(resFile, f)
	// 这里好像存在一个疑惑,multipart.File 类型的文件只能转存,不能直接存吗
	return nil
}

// 拿文件
func (fm *FileManagement) GetFile(filename string) (contentType string, err error) {
	fmt.Println(config.Config.Root + filename)
	// 拿到文件类型
	slice := strings.Split(filename, ".")
	filetype := slice[len(slice)-1]
	contentType, ok := ContentTypeMap[filetype]
	if !ok {
		return "", common.ErrNew(fmt.Errorf("暂不支持%s后缀的文件", filetype), common.UserErr)
	}
	// 我们拿了一个很大的文件来测试,结果就是响应很慢
	_, err = os.Stat(config.Config.Root + filename)
	if os.IsNotExist(err) {
		return "", common.ErrNew(errors.New("文件不存在"), common.UserErr)
	} else if err != nil {
		return "", common.ErrNew(err, common.SysErr)
	}
	return contentType, err
}

// 创建新目录
func (fm *FileManagement) NewDir(currDir string, newDir string) (err error) {
	err = os.Mkdir(config.Config.Root+currDir+newDir, 0777)
	if err != nil {
		return common.ErrNew(err, common.SysErr)
	}
	return nil
}

// 删除目录或者文件
func (fm *FileManagement) DeleteDir(currDir string, deDir string) (err error) {
	err = os.Remove(config.Config.Root + currDir + deDir)
	return err
}
