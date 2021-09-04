package proto

const (
	MaxGroupNameLen = 16
	MaxStorageIDLen = 16

	MaxIPLen = 16
)

type Cmd uint8

const (
	CmdQuit  Cmd = 82
	CmdReply Cmd = 100
	CmdPing  Cmd = 111
)

// QuitCommand
type QuitCommand struct {
	Header
}

// Reply
type Reply struct {
	Header
}

// PingCommand
type PingCommand struct {
	Header
}

// PingReply
type PingReply struct {
	Header
}
