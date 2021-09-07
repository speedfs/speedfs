// Code generated by speedfs-gen. DO NOT EDIT.

package storage

import (
	"github.com/speedfs/speedfs/internal/rpc"
)

func (x *ReportStorageIDCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.StorageID[:])
}

func (x *ReportStorageIDReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *UploadFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint8(x.StorePathIndex)
	enc.EncodeUint64(x.MetadataLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeBytes(x.Ext[:])
	enc.EncodeBytes(x.Metadata[:])
	enc.EncodeBytes(x.Content[:])
}

func (x *UploadFileReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *UploadAppenderFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint8(x.StorePathIndex)
	enc.EncodeUint64(x.MetadataLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeBytes(x.Ext[:])
	enc.EncodeBytes(x.Metadata[:])
	enc.EncodeBytes(x.Content[:])
}

func (x *UploadAppenderFileReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *DeleteFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *DeleteFileReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *SetMetadataCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.MetadataLen)
	enc.EncodeByte(x.Flag)
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
	enc.EncodeString(x.Metadata)
}

func (x *SetMetadataReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *GetMetadataCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *GetMetadataReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeString(x.Metadata)
}

func (x *DownloadFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint64(x.Offset)
	enc.EncodeUint64(x.Size)
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *DownloadFileReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.Content[:])
}

func (x *GetFileInfoCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *GetFileInfoReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint64(x.Size)
	enc.EncodeUint64(x.CreateTime)
	enc.EncodeUint64(x.CRC32)
	enc.EncodeBytes(x.SourceIP[:])
}

func (x *AppendFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeString(x.Filename)
	enc.EncodeBytes(x.Content[:])
}

func (x *AppendFileReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *ModifyFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.Offset)
	enc.EncodeUint64(x.Size)
	enc.EncodeString(x.Filename)
	enc.EncodeBytes(x.Content[:])
}

func (x *ModifyFileReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *TruncateFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeString(x.Filename)
}

func (x *TruncateFileReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *RenameFileCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeString(x.Filename)
}

func (x *RenameFileReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}
