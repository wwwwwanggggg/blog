package controller

import (
	"blog/common"
	"blog/config"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type FileName struct {
	FileName string `form:"filename" binding:"required"`
}

type DirInfo struct {
	CurrentDir string `json:"curr_dir" binding:"required"`
	OpDir      string `json:"op_dir" binding:"required"`
}

type FileManagement struct {
}

// 上传文件
func (fm *FileManagement) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	mf, e := c.MultipartForm()
	if e != nil {
		c.Error(common.ErrNew(e, common.ParamErr))
	}
	fmt.Printf("\nmf:%v", mf)
	if err != nil {
		c.Error(common.ErrNew(errors.New("接受文件上传失败"), common.ParamErr))
	}
	path, ok := c.GetQuery("path")
	if !ok {
		c.Error(common.ErrNew(errors.New("请传入文件存储路径"), common.ParamErr))
		return
	}
	c.SaveUploadedFile(file, path)
	// err = srv.FileManagement.UploadFile(file)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, nil, "上传成功"))
}

// 获取文件
func (fm *FileManagement) GetFile(c *gin.Context) {
	var info FileName
	if err := c.ShouldBindQuery(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	contentType, err := srv.FileManagement.GetFile(info.FileName)
	if err != nil {
		c.Error(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("正在传文件")
		c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", info.FileName))
		c.Header("Content-Type", contentType)
		c.File(config.Config.Root + info.FileName)
		fmt.Println("传完了")
	}()
	fmt.Println(c.GetHeader("Content-Type"))
	wg.Wait()
}

// 新建目录
func (fm *FileManagement) NewDir(c *gin.Context) {
	info := DirInfo{}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	err := srv.FileManagement.NewDir(info.CurrentDir, info.OpDir)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, nil, "创建成功"))
}

// 删除目录
func (fm *FileManagement) DeleteDir(c *gin.Context) {
	var info DirInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.Error(common.ErrNew(err, common.ParamErr))
		return
	}
	err := srv.FileManagement.DeleteDir(info.CurrentDir, info.OpDir)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, nil, "删除成功"))
}
