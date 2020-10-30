package dao

import (
	"database/sql"
	"go-network-disk/db"
	"go-network-disk/model"
	"log"
)

// OnFileUploadFinished: 文件上传到DB
func OnFileUploadFinished(meta *model.FileMeta) bool {
	stmtIn, err := db.DB.Prepare(`insert into tbl_file (file_sha1,file_name,file_size,file_addr,status) values (?,?,?,?,1)`)
	if err != nil {
		log.Printf("DB.Prepare fail,err: %v \n", err.Error())
		return false
	}
	defer stmtIn.Close()
	result, err := stmtIn.Exec(meta.FileSha1, meta.FileName, meta.FileSize, meta.FileAddr)
	if err != nil {
		log.Printf("stmtIn.Exec fail,err: %v \n", err.Error())
		return false
	}
	if rowsAffected, err := result.RowsAffected(); err == nil {
		if rowsAffected <= 0 {
			return false
		}
		return true
	} else {
		log.Printf("result.RowsAffected fail,err: %v \n", err.Error())
	}
	return false
}

// GetFileMeta: 从DB中获取文件元信息
func GetFileMeta(fileHash string) (*model.FileMeta, error) {
	fileMeta := &model.FileMeta{}
	stmtOut, err := db.DB.Prepare(`select file_sha1,file_addr,file_name,file_size,update_at from tbl_file where file_sha1=? and status=1 limit 1`)
	if err != nil {
		log.Printf("DB.Prepare fail,err: %v \n", err.Error())
		return nil, err
	}
	defer stmtOut.Close()
	if err = stmtOut.QueryRow(fileHash).Scan(&fileMeta.FileSha1, &fileMeta.FileAddr,
		&fileMeta.FileName, &fileMeta.FileSize, &fileMeta.UpdateAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return fileMeta, nil
}

// UpdateFileLocation: 从DB中更新location
func UpdateFileLocation(meta *model.FileMeta) bool {
	stmt, err := db.DB.Prepare("update tbl_file set`file_addr`=? where  `file_sha1`=? limit 1")
	if err != nil {
		log.Printf("DB.Prepare fail,err: %v \n", err.Error())
		return false
	}
	defer stmt.Close()
	if result, err := stmt.Exec(meta.FileAddr, meta.FileSha1); err != nil {
		log.Printf("stmt.Exec fail,err: %v \n", err.Error())
		return false
	} else {
		if ret, err := result.RowsAffected(); err == nil {
			if ret <= 0 {
				log.Printf("sql: no rows in result set")
			}
			return true
		}
		return false
	}
}
