package meta

import (
	"go-network-disk/dao"
	"go-network-disk/model"
	"log"
)

// FileMeta: 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]*FileMeta

func init() {
	fileMetas = make(map[string]*FileMeta)
}

// UpdateFileMeta: 新增/更新文件元信息
func (f *FileMeta) UpdateFileMeta() {
	fileMetas[f.FileSha1] = f
}

// UpdateFileMetaDB: 新增/更新文件元信息DB
func (f *FileMeta) UpdateFileMetaDB() bool {
	fm := &model.FileMeta{
		FileSha1: f.FileSha1,
		FileName: f.FileName,
		FileSize: f.FileSize,
		FileAddr: f.Location,
	}
	return dao.OnFileUploadFinished(fm)
}

// GetFileMeta: 通过sha1获取文件元信息对象
func (f *FileMeta) GetFileMeta(fileSha1 string) *FileMeta {
	return fileMetas[fileSha1]
}

// GetFileMetaDB: 通过sha1从DB获取文件元信息对象
func (f *FileMeta) GetFileMetaDB() (*FileMeta, error) {
	fileMeta, err := dao.GetFileMeta(f.FileSha1)
	if err != nil {
		log.Printf("db.GetFileMeta fail,err: %v \n", err.Error())
		return nil, err
	}
	fm := &FileMeta{
		FileSha1: fileMeta.FileSha1,
		FileName: fileMeta.FileName,
		FileSize: fileMeta.FileSize,
		Location: fileMeta.FileAddr,
		UploadAt: fileMeta.UpdateAt,
	}
	return fm, nil
}

// DeleteFileMeta: 通过sha1删除文件元信息
func (f *FileMeta) DeleteFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
