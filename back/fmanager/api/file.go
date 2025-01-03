package api

import (
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/icehardstone/fmanager/middleware"
	"github.com/icehardstone/fmanager/serror"
	"github.com/icehardstone/fmanager/store"
	"github.com/icehardstone/fmanager/utils"
)

type FileController struct {
	Controller
	fileManager store.IFileManager
}

// @Description File List Stat
type FileListStat struct {
	utils.BaseTime
	// Name
	Name string `json:"name"`
	// ID
	ID int `json:"id"`
	// Size
	Size int64 `json:"size"`
	// Type
	Type string `json:"type"`
	// Tags
	Tags []string `json:"tags"`
}

type GetFileListArgs struct {
	// page
	PageArgs
	// search key word
	Key string `json:"key" form:"key"`
}

// @Description get struct array
// @Accept  json
// @Produce  json
// @Tags File
// @Param   pageSize   query    int     true        "pageSize"
// @Param   pageIndex  query    int     true        "pageIndex"
// @Param   key        query    string  false       "search key word"
// @Success 200 {object} Result{data=[]FileListStat} "ok"
// @Failure 400 {object} serror.APIError "We need args!!"
// @Failure 404 {object} serror.APIError "Can not find data"
// @Router  /api/file [get]
// @Security ApiKeyAuth
func (c *FileController) GetFileList(ctx *gin.Context) {

	var args GetFileListArgs
	if err := ctx.ShouldBindQuery(&args); err != nil {
		ctx.JSON(http.StatusBadRequest, &serror.APIError{
			ErrorMessage: "We need args!",
		})
		return
	}

	files, err := c.Store.Queryfile(&args.Key, args.OffSet(), args.Take())
	if err != nil {
		log.Printf("query files by keyword '%s' get a error, %s", args.Key, err)
	}
	var rets []FileListStat
	for _, f := range *files {
		rets = append(rets, FileListStat{
			Size: f.Size,
			Tags: strings.Split(*f.Tags, ";"),
			Type: f.Type,
			Name: f.Name,
			ID:   int(f.ID),
			BaseTime: utils.BaseTime{
				CreateTime: &utils.DTime{Time: f.CreatedAt},
			},
		})
	}
	ctx.JSON(http.StatusOK, ResultOk(&rets))
}

// @Description file stat
type FileStat struct {
	// create time
	utils.BaseTime
	// file's Name
	Name string `json:"name"`
	// file's ID
	ID int `json:"id"`
	/// file's Size
	Size int64 `json:"size"`
	// file's Type
	Type string `json:"type"`
	// file's Ticket
	Ticket string `json:"ticket"`
	// file's Tags
	Tags []string `json:"tags"`
}

// @Description get file by ID
type GetFileArgs struct {
	// ID
	ID int `json:"id" uri:"id" binding:"required"`
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Tags File
// @Param   id     path    int     true        "ID"
// @Success 200 {object} Result{data=FileStat} "{}"
// @Failure 400 {object} serror.APIError "We need ID!!"
// @Failure 404 {object} serror.APIError "Can not find ID"
// @Router /api/file/{id} [get]
// @Security ApiKeyAuth
func (c *FileController) GetFile(cxt *gin.Context) {

	var args GetFileArgs

	if err := cxt.ShouldBindUri(&args); err != nil {
		cxt.JSON(http.StatusBadRequest, &serror.APIError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "We need ID!",
		})
		return
	}

	f, err := c.Store.FindfileById(args.ID)

	if err != nil {
		c.NotFound(cxt)
		return
	}

	filestat := &FileStat{
		Name: f.Name,
		ID:   int(f.ID),
		Size: f.Size,
		Tags: strings.Split(*f.Tags, ";"),
	}

	filestat.CreateTime = &utils.DTime{Time: f.CreatedAt}
	cxt.JSON(http.StatusOK, ResultOk(&filestat))
}

// @Description upload file input args
type PutFileArgs struct {
	// file
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Description upload file response
type PutFileRet struct {
	FileStat
}

// @Description post a file
// @Accept  multipart/form-data
// @Produce  json
// @Tags File
// @Param   file 	  formData 		file  		true 		"file"
// @Success 200 {object} Result{data=PutFileRet} "ok"
// @Failure 400 {object} serror.APIError "We need ID!!"
// @Failure 404 {object} serror.APIError "Can not find ID"
// @Router /api/file [post]
func (c *FileController) PutFile(cxt *gin.Context) {

	var args PutFileArgs

	if err := cxt.ShouldBind(&args); err != nil {
		cxt.JSON(http.StatusBadRequest, &serror.APIError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "We need file!",
		})
		return
	}

	if args.File == nil {
		log.Printf("file is null")
		c.BadRequest(cxt)
	}

	storage_name, err := c.fileManager.Save(args.File)
	if err != nil {
		log.Printf("storage file %s accure error, %s", args.File.Filename, err)
		c.BadRequest(cxt)
		return
	}

	f := store.File{
		Name:        args.File.Filename,
		Size:        args.File.Size,
		Type:        filepath.Ext(args.File.Filename),
		Tags:        new(string),
		Ticket:      uuid.NewString(),
		StorageName: *storage_name,
	}

	err = c.Store.Creatfile(f)

	if err != nil {
		c.BadRequest(cxt)
	}

	cxt.JSON(http.StatusOK, ResultOk(&PutFileRet{
		FileStat: FileStat{},
	}))
}

// @Description delete by ID
// @Accept  json
// @Produce  json
// @Tags File
// @Param   id     path    int     true        "ID"
// @success 200 {object} Result{data=integer} "desc"
// @Failure 400 {object} serror.APIError "We need ID!!"
// @Failure 404 {object} serror.APIError "Can not find ID"
// @Router /api/file/{id} [delete]
func (c *FileController) DeleteFile(cxt *gin.Context) {

	id, err := strconv.Atoi(cxt.Params.ByName("id"))

	if err != nil {
		c.BadRequest(cxt)
		return
	}

	err = c.Store.Deletefile(id)

	if err != nil {
		c.BadRequest(cxt)
		return
	}

	cxt.JSON(http.StatusOK, ResultOk(id))
}

// @Description download file by id
// @Accept  json
// @Tags File
// @Param   id     path    int     true        "ID"
// @success 200 {object} any "desc"
// @Failure 400 {object} serror.APIError "We need ID!!"
// @Failure 404 {object} serror.APIError "Can not find ID"
// @Router /api/file/{id}/download [get]
func (c *FileController) Download(cxt *gin.Context) {
	id, err := strconv.Atoi(cxt.Params.ByName("id"))

	if err != nil {
		c.BadRequest(cxt)
		return
	}
	file, err := c.Store.FindfileById(id)

	if err != nil {
		log.Printf("find file by id %d accure error, %s", id, err)
		c.BadRequest(cxt)
		return
	}

	err = c.fileManager.Load(cxt, *file)

	if err != nil {
		log.Printf("copy file to strem accur error, %s", err)
		c.BadRequest(cxt)
	}
}

// construct a new file controller
func NewFileController() *FileController {
	return &FileController{
		Controller:  *NewController(),
		fileManager: &store.DiskFileManager{},
	}
}

// create、get、get list、delete file
func RegisterFile(router *gin.Engine) {
	fg := router.Group("/api/file")
	fg.Use(middleware.OAuth2Middleware)
	c := NewFileController()

	fg.POST("", c.PutFile)
	fg.GET(":id", c.GetFile)
	fg.GET("", c.GetFileList)
	fg.DELETE(":id", c.DeleteFile)
	fg.GET(":id/download", c.Download)
}
