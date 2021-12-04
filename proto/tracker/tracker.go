package tracker

import (
	"context"

	"github.com/speedfs/speedfs/proto"
	"github.com/speedfs/speedfs/rpc"
)

const (
	CmdReply rpc.Cmd = 100

	CmdQuit rpc.Cmd = 82
	CmdPing rpc.Cmd = 111

	CmdTrackerSaveSysFilesBegin rpc.Cmd = 61
	CmdTrackerSaveSysFilesEnd   rpc.Cmd = 62
	CmdTrackerGetSysFile        rpc.Cmd = 63
	CmdTrackerGetStatus         rpc.Cmd = 64

	CmdTrackerPingLeader           rpc.Cmd = 65
	CmdTrackerNotifyNextLeader     rpc.Cmd = 66
	CmdTrackerCommitNextLeader     rpc.Cmd = 67
	CmdStorageNotifyReselectLeader rpc.Cmd = 68

	CmdStorageUpdateStatus      rpc.Cmd = 59
	CmdStorageGetIP             rpc.Cmd = 60
	CmdStorageListIDs           rpc.Cmd = 69
	CmdStorageGetID             rpc.Cmd = 70
	CmdStorageGetStatus         rpc.Cmd = 71
	CmdStorageGetTrunkFileID    rpc.Cmd = 72
	CmdStorageReportTrunkFileID rpc.Cmd = 73
	CmdStorageReportTrunkFree   rpc.Cmd = 74
	CmdStorageGetParams         rpc.Cmd = 75
	CmdStorageReportStatus      rpc.Cmd = 76
	CmdStorageGetChangeLog      rpc.Cmd = 77
	CmdStorageUpdateIP          rpc.Cmd = 78
	CmdStorageJoin              rpc.Cmd = 81
	CmdStorageHeartbeat         rpc.Cmd = 83
	CmdStorageReportDiskUsage   rpc.Cmd = 84

	CmdGetGroup          rpc.Cmd = 90
	CmdListGroups        rpc.Cmd = 91
	CmdDeleteGroup       rpc.Cmd = 108
	CmdListGroupStorages rpc.Cmd = 92
	CmdGetStorage        rpc.Cmd = 92
	CmdDeleteStorage     rpc.Cmd = 93
	CmdSetTrunkServer    rpc.Cmd = 94

	// upload
	CmdQueryStorage       rpc.Cmd = 101
	CmdQueryGroupStorage  rpc.Cmd = 104 // group name
	CmdQueryStorages      rpc.Cmd = 106
	CmdQueryGroupStorages rpc.Cmd = 107 // group name

	// download
	CmdQueryDownloadableStorage  rpc.Cmd = 102 // group name, file name
	CmdQueryDownloadableStorages rpc.Cmd = 105 // group name, file name

	// update
	CmdQueryUpdatableStorage rpc.Cmd = 103 // group name, file name
)

type Service interface {
	QueryStorage(ctx context.Context, cmd *QueryStorageCommand) (*QueryStorageReply, error)
}

// GroupStat
type GroupStat struct {
	GroupName          [proto.MaxGroupNameLen + 1]byte
	TotalMiB           uint64
	FreeMiB            uint64
	TrunkFreeMiB       uint64
	ServerCount        uint64
	StoragePort        uint64
	StorageHTTPPort    uint64
	ActiveServerCount  uint64
	CurrentWriteServer uint64
	StorePathCount     uint64
	SubdirCountPerPath uint64
	CurrentTrunkFileID uint64
}

// GetGroupCommand
type GetGroupCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
}

// GetGroupReply
type GetGroupReply struct {
	GroupStat GroupStat
}

// ListGroupsCommand
type ListGroupsCommand struct {
}

// ListGroupsReply
type ListGroupsReply struct {
	GroupStats []GroupStat
}

// DeleteGroupCommand
type DeleteGroupCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
}

// DeleteGroupReply
type DeleteGroupReply struct {
}

// ListGroupStoragesCommand
type ListGroupStoragesCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
}

// ListGroupStoragesReply
type ListGroupStoragesReply struct {
}

// GetStorageCommand
type GetStorageCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	StorageID [proto.MaxStorageIDLen]byte
}

// GetStorageReply
type GetStorageReply struct {
}

// DeleteStorageCommand
type DeleteStorageCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	StorageID [proto.MaxStorageIDLen]byte
}

// DeleteStorageReply
type DeleteStorageReply struct {
}

// SetTrunkServerCommand
type SetTrunkServerCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	StorageID [proto.MaxStorageIDLen]byte
}

// SetTrunkServerReply
type SetTrunkServerReply struct {
}

// Addr
type Addr struct {
	IP   [proto.MaxIPLen - 1]byte
	Port uint64
}

// QueryStorageCommand
type QueryStorageCommand struct {
}

// QueryStorageReply
type QueryStorageReply struct {
	GroupName      [proto.MaxGroupNameLen]byte
	Addr           Addr
	StorePathIndex uint8
}

// QueryGroupStorageCommand
type QueryGroupStorageCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
}

// QueryGroupStorageReply
type QueryGroupStorageReply struct {
	GroupName      [proto.MaxGroupNameLen]byte
	Addr           Addr
	StorePathIndex uint8
}

// QueryStoragesCommand
type QueryStoragesCommand struct {
}

// QueryStoragesReply
type QueryStoragesReply struct {
	GroupName      [proto.MaxGroupNameLen]byte
	Addr           []Addr
	StorePathIndex uint8
}

// QueryGroupStoragesCommand
type QueryGroupStoragesCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
}

// QueryGroupStoragesReply
type QueryGroupStoragesReply struct {
	GroupName      [proto.MaxGroupNameLen]byte
	Addr           []Addr
	StorePathIndex uint8
}

// QueryDownloadableStorageCommand
type QueryDownloadableStorageCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// QueryDownloadableStorageReply
type QueryDownloadableStorageReply struct {
	GroupName [proto.MaxGroupNameLen]byte
	Addr      Addr
}

// QueryDownloadableStoragesCommand
type QueryDownloadableStoragesCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// QueryDownloadableStoragesReply
type QueryDownloadableStoragesReply struct {
	GroupName [proto.MaxGroupNameLen]byte
	Port      uint8
	IPs       [][proto.MaxIPLen - 1]byte
}

// QueryUpdatableStorageCommand
type QueryUpdatableStorageCommand struct {
	GroupName [proto.MaxGroupNameLen]byte
	Filename  string
}

// QueryUpdatableStorageReply
type QueryUpdatableStorageReply struct {
	GroupName [proto.MaxGroupNameLen]byte
	Addr      Addr
}
