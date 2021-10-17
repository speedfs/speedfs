// Code generated by speedfs-gen. DO NOT EDIT.

package storage

import (
	"github.com/speedfs/speedfs/proto"
)

// EncodeTo
func (x *ReportStorageIDCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.StorageID[:])
}

// DecodeFrom
func (x *ReportStorageIDCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.StorageID[:]); err != nil {
		return err
	}

	return nil
}

// Cmd
func (x *ReportStorageIDCommand) Cmd() proto.Cmd {
	return CmdReportStorageID
}

// EncodeTo
func (x *ReportStorageIDReply) EncodeTo(enc *proto.Encoder) {
}

// DecodeFrom
func (x *ReportStorageIDReply) DecodeFrom(dec *proto.Decoder) error {

	return nil
}

// Cmd
func (x *ReportStorageIDReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *UploadFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint8(x.StorePathIndex)
	enc.EncodeUint64(x.MetadataLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeBytes(x.Ext[:])
	enc.EncodeBytes(x.Metadata[:])
	enc.EncodeBytes(x.Content[:])
}

// DecodeFrom
func (x *UploadFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.StorePathIndex, err = dec.DecodeUint8(); err != nil {
		return err
	}
	if x.MetadataLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if err = dec.DecodeBytes(x.Ext[:]); err != nil {
		return err
	}
	x.Metadata = make([]byte, x.MetadataLen)
	if err = dec.DecodeBytes(x.Metadata); err != nil {
		return err
	}
	x.Content = make([]byte, x.Size)
	if err = dec.DecodeBytes(x.Content); err != nil {
		return err
	}

	return nil
}

// Cmd
func (x *UploadFileCommand) Cmd() proto.Cmd {
	return CmdUploadFile
}

// EncodeTo
func (x *UploadFileReply) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *UploadFileReply) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *UploadFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *UploadAppenderFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint8(x.StorePathIndex)
	enc.EncodeUint64(x.MetadataLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeBytes(x.Ext[:])
	enc.EncodeBytes(x.Metadata[:])
	enc.EncodeBytes(x.Content[:])
}

// DecodeFrom
func (x *UploadAppenderFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.StorePathIndex, err = dec.DecodeUint8(); err != nil {
		return err
	}
	if x.MetadataLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if err = dec.DecodeBytes(x.Ext[:]); err != nil {
		return err
	}
	x.Metadata = make([]byte, x.MetadataLen)
	if err = dec.DecodeBytes(x.Metadata); err != nil {
		return err
	}
	x.Content = make([]byte, x.Size)
	if err = dec.DecodeBytes(x.Content); err != nil {
		return err
	}

	return nil
}

// Cmd
func (x *UploadAppenderFileCommand) Cmd() proto.Cmd {
	return CmdUploadAppenderFile
}

// EncodeTo
func (x *UploadAppenderFileReply) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *UploadAppenderFileReply) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *UploadAppenderFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *DeleteFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *DeleteFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *DeleteFileCommand) Cmd() proto.Cmd {
	return CmdDeleteFile
}

// EncodeTo
func (x *DeleteFileReply) EncodeTo(enc *proto.Encoder) {
}

// DecodeFrom
func (x *DeleteFileReply) DecodeFrom(dec *proto.Decoder) error {

	return nil
}

// Cmd
func (x *DeleteFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *SetMetadataCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.MetadataLen)
	enc.EncodeByte(x.Flag)
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
	enc.EncodeString(x.Metadata)
}

// DecodeFrom
func (x *SetMetadataCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.FilenameLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.MetadataLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Flag, err = dec.DecodeByte(); err != nil {
		return err
	}
	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	{
		buf := make([]byte, x.FilenameLen)
		if err = dec.DecodeBytes(buf); err != nil {
			return err
		}
		x.Filename = string(buf)
	}
	{
		buf := make([]byte, x.MetadataLen)
		if err = dec.DecodeBytes(buf); err != nil {
			return err
		}
		x.Metadata = string(buf)
	}

	return nil
}

// Cmd
func (x *SetMetadataCommand) Cmd() proto.Cmd {
	return CmdSetMetadata
}

// EncodeTo
func (x *SetMetadataReply) EncodeTo(enc *proto.Encoder) {
}

// DecodeFrom
func (x *SetMetadataReply) DecodeFrom(dec *proto.Decoder) error {

	return nil
}

// Cmd
func (x *SetMetadataReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *GetMetadataCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *GetMetadataCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *GetMetadataCommand) Cmd() proto.Cmd {
	return CmdGetMetadata
}

// EncodeTo
func (x *GetMetadataReply) EncodeTo(enc *proto.Encoder) {
	enc.EncodeString(x.Metadata)
}

// DecodeFrom
func (x *GetMetadataReply) DecodeFrom(dec *proto.Decoder) error {

	x.Metadata = string(dec.Bytes())

	return nil
}

// Cmd
func (x *GetMetadataReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *DownloadFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint64(x.Offset)
	enc.EncodeUint64(x.Size)
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *DownloadFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.Offset, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *DownloadFileCommand) Cmd() proto.Cmd {
	return CmdDownloadFile
}

// EncodeTo
func (x *DownloadFileReply) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.Content[:])
}

// DecodeFrom
func (x *DownloadFileReply) DecodeFrom(dec *proto.Decoder) error {

	x.Content = dec.Bytes()

	return nil
}

// Cmd
func (x *DownloadFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *GetFileInfoCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *GetFileInfoCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *GetFileInfoCommand) Cmd() proto.Cmd {
	return CmdGetFileInfo
}

// EncodeTo
func (x *GetFileInfoReply) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint64(x.Size)
	enc.EncodeUint64(x.CreateTime)
	enc.EncodeUint64(x.CRC32)
	enc.EncodeBytes(x.SourceIP[:])
}

// DecodeFrom
func (x *GetFileInfoReply) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.CreateTime, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.CRC32, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if err = dec.DecodeBytes(x.SourceIP[:]); err != nil {
		return err
	}

	return nil
}

// Cmd
func (x *GetFileInfoReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *AppendFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeString(x.Filename)
	enc.EncodeBytes(x.Content[:])
}

// DecodeFrom
func (x *AppendFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.FilenameLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	{
		buf := make([]byte, x.FilenameLen)
		if err = dec.DecodeBytes(buf); err != nil {
			return err
		}
		x.Filename = string(buf)
	}
	x.Content = make([]byte, x.Size)
	if err = dec.DecodeBytes(x.Content); err != nil {
		return err
	}

	return nil
}

// Cmd
func (x *AppendFileCommand) Cmd() proto.Cmd {
	return CmdAppendFile
}

// EncodeTo
func (x *AppendFileReply) EncodeTo(enc *proto.Encoder) {
}

// DecodeFrom
func (x *AppendFileReply) DecodeFrom(dec *proto.Decoder) error {

	return nil
}

// Cmd
func (x *AppendFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *ModifyFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.Offset)
	enc.EncodeUint64(x.Size)
	enc.EncodeString(x.Filename)
	enc.EncodeBytes(x.Content[:])
}

// DecodeFrom
func (x *ModifyFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.FilenameLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Offset, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	{
		buf := make([]byte, x.FilenameLen)
		if err = dec.DecodeBytes(buf); err != nil {
			return err
		}
		x.Filename = string(buf)
	}
	x.Content = make([]byte, x.Size)
	if err = dec.DecodeBytes(x.Content); err != nil {
		return err
	}

	return nil
}

// Cmd
func (x *ModifyFileCommand) Cmd() proto.Cmd {
	return CmdModifyFile
}

// EncodeTo
func (x *ModifyFileReply) EncodeTo(enc *proto.Encoder) {
}

// DecodeFrom
func (x *ModifyFileReply) DecodeFrom(dec *proto.Decoder) error {

	return nil
}

// Cmd
func (x *ModifyFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *TruncateFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeUint64(x.FilenameLen)
	enc.EncodeUint64(x.Size)
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *TruncateFileCommand) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if x.FilenameLen, err = dec.DecodeUint64(); err != nil {
		return err
	}
	if x.Size, err = dec.DecodeUint64(); err != nil {
		return err
	}
	{
		buf := make([]byte, x.FilenameLen)
		if err = dec.DecodeBytes(buf); err != nil {
			return err
		}
		x.Filename = string(buf)
	}

	return nil
}

// Cmd
func (x *TruncateFileCommand) Cmd() proto.Cmd {
	return CmdTruncateFile
}

// EncodeTo
func (x *TruncateFileReply) EncodeTo(enc *proto.Encoder) {
}

// DecodeFrom
func (x *TruncateFileReply) DecodeFrom(dec *proto.Decoder) error {

	return nil
}

// Cmd
func (x *TruncateFileReply) Cmd() proto.Cmd {
	return CmdReply
}

// EncodeTo
func (x *RenameFileCommand) EncodeTo(enc *proto.Encoder) {
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *RenameFileCommand) DecodeFrom(dec *proto.Decoder) error {

	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *RenameFileCommand) Cmd() proto.Cmd {
	return CmdRenameFile
}

// EncodeTo
func (x *RenameFileReply) EncodeTo(enc *proto.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

// DecodeFrom
func (x *RenameFileReply) DecodeFrom(dec *proto.Decoder) error {

	var err error

	if err = dec.DecodeBytes(x.GroupName[:]); err != nil {
		return err
	}
	x.Filename = string(dec.Bytes())

	return nil
}

// Cmd
func (x *RenameFileReply) Cmd() proto.Cmd {
	return CmdReply
}
