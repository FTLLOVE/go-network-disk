package model

type FileMeta struct {
	FileSha1 string `json:"file_sha_1"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FileAddr string `json:"file_addr"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
	Status   int8   `json:"status"`
	Ext1     int    `json:"ext_1"`
	Ext2     string `json:"ext_2"`
}
