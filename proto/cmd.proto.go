// Code generated by speedfs-rpcgen. DO NOT EDIT.

package proto

// EncodeTo
func (x *QuitCommand) EncodeTo(enc *Encoder) {
}

// DecodeFrom
func (x *QuitCommand) DecodeFrom(dec *Decoder) error {

	return nil
}

// Cmd
func (x *QuitCommand) Cmd() Cmd {
	return CmdQuit
}

// EncodeTo
func (x *Reply) EncodeTo(enc *Encoder) {
}

// DecodeFrom
func (x *Reply) DecodeFrom(dec *Decoder) error {

	return nil
}

// EncodeTo
func (x *PingCommand) EncodeTo(enc *Encoder) {
}

// DecodeFrom
func (x *PingCommand) DecodeFrom(dec *Decoder) error {

	return nil
}

// Cmd
func (x *PingCommand) Cmd() Cmd {
	return CmdPing
}

// EncodeTo
func (x *PingReply) EncodeTo(enc *Encoder) {
}

// DecodeFrom
func (x *PingReply) DecodeFrom(dec *Decoder) error {

	return nil
}

// Cmd
func (x *PingReply) Cmd() Cmd {
	return CmdReply
}
