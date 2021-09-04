package storage

import (
	"github.com/speedfs/speedfs/proto"
)

const (
	MaxFileExtLen = 6
)

const (
	CmdReportStorageID    proto.Cmd = 9
	CmdUploadFile         proto.Cmd = 11
	CmdUploadAppenderFile proto.Cmd = 21
	CmdDeleteFile         proto.Cmd = 12
	CmdSetMetadata        proto.Cmd = 13
	CmdGetMetadata        proto.Cmd = 15
	CmdDownloadFile       proto.Cmd = 14
	CmdGetFileInfo        proto.Cmd = 22
	CmdAppendFile         proto.Cmd = 24
	CmdModifyFile         proto.Cmd = 34
	CmdTruncateFile       proto.Cmd = 36
	CmdRenameFile         proto.Cmd = 38
)

// ReportIDCommand
type ReportStorageIDCommand struct {
	proto.Header
	StorageID [proto.MaxStorageIDLen]byte
}

// ReportStorageIDReply
type ReportStorageIDReply struct {
	proto.Header
}

// UploadFileCommand
type UploadFileCommand struct {
	proto.Header
	StorePathIndex uint8
	MetadataLen    uint64
	Size           uint64
	Ext            [MaxFileExtLen]byte
	Metadata       []byte
	Content        []byte
}

// UploadFileReply
type UploadFileReply struct {
	proto.Header
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// UploadAppenderFileCommand
type UploadAppenderFileCommand struct {
	proto.Header
	StorePathIndex uint8
	MetadataLen    uint64
	Size           uint64
	Ext            [MaxFileExtLen]byte
	Metadata       []byte
	Content        []byte
}

// UploadAppenderFileReply
type UploadAppenderFileReply struct {
	proto.Header
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// DeleteFileCommand
type DeleteFileCommand struct {
	proto.Header
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// DeleteFileReply
type DeleteFileReply struct {
	proto.Header
}

// SetMetadataCommand
type SetMetadataCommand struct {
	proto.Header
	FilenameLen uint64
	MetadataLen uint64
	Flag        byte
	GroupName   [proto.MaxGroupNameLen]byte
	Filename    string
	Metadata    string
}

// SetMetadataReply
type SetMetadataReply struct {
	proto.Header
}

// GetMetadataCommand
type GetMetadataCommand struct {
	proto.Header
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// GetMetadataReply
type GetMetadataReply struct {
	proto.Header
	Metadata string
}

// DownloadFileCommand
type DownloadFileCommand struct {
	proto.Header
	Offset    uint64
	Size      uint64
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// DownloadFileReply
type DownloadFileReply struct {
	proto.Header
	Content []byte
}

// GetFileInfoCommand
type GetFileInfoCommand struct {
	proto.Header
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// GetFileInfoReply
type GetFileInfoReply struct {
	proto.Header
	Size       uint64
	CreateTime uint64
	CRC32      uint64
	SourceIP   [proto.MaxIPLen]byte
}

// AppendFileCommand
type AppendFileCommand struct {
	proto.Header
	FilenameLen uint64
	Size        uint64
	Filename    string
	Content     []byte
}

// AppendFileReply
type AppendFileReply struct {
	proto.Header
}

// ModifyFileCommand
type ModifyFileCommand struct {
	proto.Header
	FilenameLen uint64
	Offset      uint64
	Size        uint64
	Filename    string
	Content     []byte
}

// ModifyFileReply
type ModifyFileReply struct {
	proto.Header
}

// TruncateFileCommand
type TruncateFileCommand struct {
	proto.Header
	FilenameLen uint64
	Size        uint64
	Filename    string
}

// TruncateFileReply
type TruncateFileReply struct {
	proto.Header
}

// RenameFileCommand
type RenameFileCommand struct {
	proto.Header
	Filename string
}

// RenameFileReply
type RenameFileReply struct {
	proto.Header
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}
