// Code generated by speedfs-gen. DO NOT EDIT.

package tracker

import (
	"github.com/speedfs/speedfs/internal/rpc"
)

func (x *GroupStat) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeUint64(x.TotalMiB)
	enc.EncodeUint64(x.FreeMiB)
	enc.EncodeUint64(x.TrunkFreeMiB)
	enc.EncodeUint64(x.ServerCount)
	enc.EncodeUint64(x.StoragePort)
	enc.EncodeUint64(x.StorageHTTPPort)
	enc.EncodeUint64(x.ActiveServerCount)
	enc.EncodeUint64(x.CurrentWriteServer)
	enc.EncodeUint64(x.StorePathCount)
	enc.EncodeUint64(x.SubdirCountPerPath)
	enc.EncodeUint64(x.CurrentTrunkFileID)
}

func (x *GetGroupCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
}

func (x *GetGroupReply) EncodeTo(enc *rpc.Encoder) {
	x.GroupStat.EncodeTo(enc)
}

func (x *ListGroupsCommand) EncodeTo(enc *rpc.Encoder) {
}

func (x *ListGroupsReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *DeleteGroupCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
}

func (x *DeleteGroupReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *ListGroupStoragesCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
}

func (x *ListGroupStoragesReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *GetStorageCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeBytes(x.StorageID[:])
}

func (x *GetStorageReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *DeleteStorageCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeBytes(x.StorageID[:])
}

func (x *DeleteStorageReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *SetTrunkServerCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeBytes(x.StorageID[:])
}

func (x *SetTrunkServerReply) EncodeTo(enc *rpc.Encoder) {
}

func (x *Addr) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.IP[:])
	enc.EncodeUint64(x.Port)
}

func (x *QueryStorageCommand) EncodeTo(enc *rpc.Encoder) {
}

func (x *QueryStorageReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	x.Addr.EncodeTo(enc)
	enc.EncodeUint8(x.StorePathIndex)
}

func (x *QueryGroupStorageCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
}

func (x *QueryGroupStorageReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	x.Addr.EncodeTo(enc)
	enc.EncodeUint8(x.StorePathIndex)
}

func (x *QueryStoragesCommand) EncodeTo(enc *rpc.Encoder) {
}

func (x *QueryStoragesReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeUint8(x.StorePathIndex)
}

func (x *QueryGroupStoragesCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
}

func (x *QueryGroupStoragesReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeUint8(x.StorePathIndex)
}

func (x *QueryDownloadableStorageCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *QueryDownloadableStorageReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	x.Addr.EncodeTo(enc)
}

func (x *QueryDownloadableStoragesCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *QueryDownloadableStoragesReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeUint8(x.Port)
}

func (x *QueryUpdatableStorageCommand) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	enc.EncodeString(x.Filename)
}

func (x *QueryUpdatableStorageReply) EncodeTo(enc *rpc.Encoder) {
	enc.EncodeBytes(x.GroupName[:])
	x.Addr.EncodeTo(enc)
}