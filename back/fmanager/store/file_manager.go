package store

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var save_path string = "/tmp"

type IFileManager interface {
	Save(file *multipart.FileHeader) (*string, error)
	Load(cxt *gin.Context, f File) error
}

type DiskFileManager struct {
}

func (disk *DiskFileManager) Load(cxt *gin.Context, f File) error {

	cxt.Header("Content-Disposition", "attachment; filename="+f.Name)
	cxt.File(f.StorageName)

	return nil
}
func (disk *DiskFileManager) Save(file *multipart.FileHeader) (*string, error) {

	src, err := file.Open()

	if err != nil {
		return nil, err
	}

	defer src.Close()

	if runtime.GOOS != "linux" {
		save_path = ""
	}
	dst := filepath.Join(save_path, uuid.NewString())

	out, err := os.Create(dst)

	if err != nil {
		return nil, err
	}

	defer out.Close()

	_, err = io.Copy(out, src)
	return &dst, err
}
