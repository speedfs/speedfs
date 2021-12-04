package storage

import (
	"context"

	"github.com/speedfs/speedfs/proto"
	"github.com/speedfs/speedfs/rpc"
)

const (
	MaxFileExtLen = 6
)

const (
	CmdReply rpc.Cmd = 100

	CmdQuit rpc.Cmd = 82
	CmdPing rpc.Cmd = 111

	CmdReportStorageID    rpc.Cmd = 9
	CmdUploadFile         rpc.Cmd = 11
	CmdUploadAppenderFile rpc.Cmd = 21
	CmdDeleteFile         rpc.Cmd = 12
	CmdSetMetadata        rpc.Cmd = 13
	CmdGetMetadata        rpc.Cmd = 15
	CmdDownloadFile       rpc.Cmd = 14
	CmdGetFileInfo        rpc.Cmd = 22
	CmdAppendFile         rpc.Cmd = 24
	CmdModifyFile         rpc.Cmd = 34
	CmdTruncateFile       rpc.Cmd = 36
	CmdRenameFile         rpc.Cmd = 38
)

type Service interface {
	UploadFile(ctx context.Context, cmd *UploadFileCommand) (*UploadFileReply, error)
	UploadAppenderFile(ctx context.Context, cmd *UploadAppenderFileCommand) (*UploadAppenderFileReply, error)
	SetMetadata(ctx context.Context, cmd *SetMetadataCommand) (*SetMetadataReply, error)
	GetMetadata(ctx context.Context, cmd *GetMetadataCommand) (*GetMetadataReply, error)
	DownloadFile(ctx context.Context, cmd *DownloadFileCommand) (*DownloadFileReply, error)
	AppendFile(ctx context.Context, cmd *AppendFileCommand) (*AppendFileReply, error)
	ModifyFile(ctx context.Context, cmd *ModifyFileCommand) (*ModifyFileReply, error)
	TruncateFile(ctx context.Context, cmd *TruncateFileCommand) (*TruncateFileReply, error)
	RenameFile(ctx context.Context, cmd *RenameFileCommand) (*RenameFileReply, error)
}

// ReportIDCommand
type ReportStorageIDCommand struct {
	StorageID [proto.MaxStorageIDLen]byte
}

// ReportStorageIDReply
type ReportStorageIDReply struct {
}

// UploadFileCommand
type UploadFileCommand struct {
	StorePathIndex uint8
	MetadataLen    uint64
	Size           uint64
	Ext            [MaxFileExtLen]byte
	Metadata       []byte `proto:"len:MetadataLen"`
	Content        []byte `proto:"len:Size"`
}

// UploadFileReply
type UploadFileReply struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// UploadAppenderFileCommand
type UploadAppenderFileCommand struct {
	StorePathIndex uint8
	MetadataLen    uint64
	Size           uint64
	Ext            [MaxFileExtLen]byte
	Metadata       []byte `proto:"len:MetadataLen"`
	Content        []byte `proto:"len:Size"`
}

// UploadAppenderFileReply
type UploadAppenderFileReply struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// DeleteFileCommand
type DeleteFileCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// DeleteFileReply
type DeleteFileReply struct {
}

// SetMetadataCommand
type SetMetadataCommand struct {
	FilenameLen uint64
	MetadataLen uint64
	Flag        byte
	GroupName   [proto.MaxGroupNameLen]byte
	Filename    string `proto:"len:FilenameLen"`
	Metadata    string `proto:"len:MetadataLen"`
}

// SetMetadataReply
type SetMetadataReply struct {
}

// GetMetadataCommand
type GetMetadataCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// GetMetadataReply
type GetMetadataReply struct {
	Metadata string
}

// DownloadFileCommand
type DownloadFileCommand struct {
	Offset    uint64
	Size      uint64
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// DownloadFileReply
type DownloadFileReply struct {
	Content []byte
}

// GetFileInfoCommand
type GetFileInfoCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// GetFileInfoReply
type GetFileInfoReply struct {
	Size       uint64
	CreateTime uint64
	CRC32      uint64
	SourceIP   [proto.MaxIPLen]byte
}

// AppendFileCommand
type AppendFileCommand struct {
	FilenameLen uint64
	Size        uint64
	Filename    string `proto:"len:FilenameLen"`
	Content     []byte `proto:"len:Size"`
}

// AppendFileReply
type AppendFileReply struct {
}

// ModifyFileCommand
type ModifyFileCommand struct {
	FilenameLen uint64
	Offset      uint64
	Size        uint64
	Filename    string `proto:"len:FilenameLen"`
	Content     []byte `proto:"len:Size"`
}

// ModifyFileReply
type ModifyFileReply struct {
}

// TruncateFileCommand
type TruncateFileCommand struct {
	FilenameLen uint64
	Size        uint64
	Filename    string `proto:"len:FilenameLen"`
}

// TruncateFileReply
type TruncateFileReply struct {
}

// RenameFileCommand
type RenameFileCommand struct {
	Filename string
}

// RenameFileReply
type RenameFileReply struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}
