package tracker

import (
	"context"

	"github.com/speedfs/speedfs/proto"
)

const (
	CmdReply proto.Cmd = 100

	CmdQuit proto.Cmd = 82
	CmdPing proto.Cmd = 111

	CmdTrackerSaveSysFilesBegin proto.Cmd = 61
	CmdTrackerSaveSysFilesEnd   proto.Cmd = 62
	CmdTrackerGetSysFile        proto.Cmd = 63
	CmdTrackerGetStatus         proto.Cmd = 64

	CmdTrackerPingLeader           proto.Cmd = 65
	CmdTrackerNotifyNextLeader     proto.Cmd = 66
	CmdTrackerCommitNextLeader     proto.Cmd = 67
	CmdStorageNotifyReselectLeader proto.Cmd = 68

	CmdStorageUpdateStatus      proto.Cmd = 59
	CmdStorageGetIP             proto.Cmd = 60
	CmdStorageListIDs           proto.Cmd = 69
	CmdStorageGetID             proto.Cmd = 70
	CmdStorageGetStatus         proto.Cmd = 71
	CmdStorageGetTrunkFileID    proto.Cmd = 72
	CmdStorageReportTrunkFileID proto.Cmd = 73
	CmdStorageReportTrunkFree   proto.Cmd = 74
	CmdStorageGetParams         proto.Cmd = 75
	CmdStorageReportStatus      proto.Cmd = 76
	CmdStorageGetChangeLog      proto.Cmd = 77
	CmdStorageUpdateIP          proto.Cmd = 78
	CmdStorageJoin              proto.Cmd = 81
	CmdStorageHeartbeat         proto.Cmd = 83
	CmdStorageReportDiskUsage   proto.Cmd = 84

	CmdGetGroup          proto.Cmd = 90
	CmdListGroups        proto.Cmd = 91
	CmdDeleteGroup       proto.Cmd = 108
	CmdListGroupStorages proto.Cmd = 92
	CmdGetStorage        proto.Cmd = 92
	CmdDeleteStorage     proto.Cmd = 93
	CmdSetTrunkServer    proto.Cmd = 94

	// upload
	CmdQueryStorage       proto.Cmd = 101
	CmdQueryGroupStorage  proto.Cmd = 104 // group name
	CmdQueryStorages      proto.Cmd = 106
	CmdQueryGroupStorages proto.Cmd = 107 // group name

	// download
	CmdQueryDownloadableStorage  proto.Cmd = 102 // group name, file name
	CmdQueryDownloadableStorages proto.Cmd = 105 // group name, file name

	// update
	CmdQueryUpdatableStorage proto.Cmd = 103 // group name, file name
)

type TrackerService interface {
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
