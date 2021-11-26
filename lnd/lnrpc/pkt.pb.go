// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkt.proto

package lnrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NeutrinoBan struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	EndTime              string   `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	BanScore             int32    `protobuf:"varint,4,opt,name=ban_score,json=banScore,proto3" json:"ban_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NeutrinoBan) Reset()         { *m = NeutrinoBan{} }
func (m *NeutrinoBan) String() string { return proto.CompactTextString(m) }
func (*NeutrinoBan) ProtoMessage()    {}
func (*NeutrinoBan) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f63a845a51abc, []int{0}
}

func (m *NeutrinoBan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeutrinoBan.Unmarshal(m, b)
}
func (m *NeutrinoBan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeutrinoBan.Marshal(b, m, deterministic)
}
func (m *NeutrinoBan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeutrinoBan.Merge(m, src)
}
func (m *NeutrinoBan) XXX_Size() int {
	return xxx_messageInfo_NeutrinoBan.Size(m)
}
func (m *NeutrinoBan) XXX_DiscardUnknown() {
	xxx_messageInfo_NeutrinoBan.DiscardUnknown(m)
}

var xxx_messageInfo_NeutrinoBan proto.InternalMessageInfo

func (m *NeutrinoBan) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *NeutrinoBan) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *NeutrinoBan) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *NeutrinoBan) GetBanScore() int32 {
	if m != nil {
		return m.BanScore
	}
	return 0
}

type NeutrinoQuery struct {
	Peer                 string   `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	ReqNum               uint32   `protobuf:"varint,3,opt,name=req_num,json=reqNum,proto3" json:"req_num,omitempty"`
	CreateTime           uint32   `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	LastRequestTime      uint32   `protobuf:"varint,5,opt,name=last_request_time,json=lastRequestTime,proto3" json:"last_request_time,omitempty"`
	LastResponseTime     uint32   `protobuf:"varint,6,opt,name=last_response_time,json=lastResponseTime,proto3" json:"last_response_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NeutrinoQuery) Reset()         { *m = NeutrinoQuery{} }
func (m *NeutrinoQuery) String() string { return proto.CompactTextString(m) }
func (*NeutrinoQuery) ProtoMessage()    {}
func (*NeutrinoQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f63a845a51abc, []int{1}
}

func (m *NeutrinoQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeutrinoQuery.Unmarshal(m, b)
}
func (m *NeutrinoQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeutrinoQuery.Marshal(b, m, deterministic)
}
func (m *NeutrinoQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeutrinoQuery.Merge(m, src)
}
func (m *NeutrinoQuery) XXX_Size() int {
	return xxx_messageInfo_NeutrinoQuery.Size(m)
}
func (m *NeutrinoQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_NeutrinoQuery.DiscardUnknown(m)
}

var xxx_messageInfo_NeutrinoQuery proto.InternalMessageInfo

func (m *NeutrinoQuery) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *NeutrinoQuery) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *NeutrinoQuery) GetReqNum() uint32 {
	if m != nil {
		return m.ReqNum
	}
	return 0
}

func (m *NeutrinoQuery) GetCreateTime() uint32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *NeutrinoQuery) GetLastRequestTime() uint32 {
	if m != nil {
		return m.LastRequestTime
	}
	return 0
}

func (m *NeutrinoQuery) GetLastResponseTime() uint32 {
	if m != nil {
		return m.LastResponseTime
	}
	return 0
}

type NeutrinoInfo struct {
	Peers                []*PeerDesc      `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	Bans                 []*NeutrinoBan   `protobuf:"bytes,2,rep,name=bans,proto3" json:"bans,omitempty"`
	Queries              []*NeutrinoQuery `protobuf:"bytes,3,rep,name=queries,proto3" json:"queries,omitempty"`
	BlockHash            string           `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Height               int32            `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	BlockTimestamp       string           `protobuf:"bytes,6,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	IsSyncing            bool             `protobuf:"varint,7,opt,name=is_syncing,json=isSyncing,proto3" json:"is_syncing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NeutrinoInfo) Reset()         { *m = NeutrinoInfo{} }
func (m *NeutrinoInfo) String() string { return proto.CompactTextString(m) }
func (*NeutrinoInfo) ProtoMessage()    {}
func (*NeutrinoInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f63a845a51abc, []int{2}
}

func (m *NeutrinoInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeutrinoInfo.Unmarshal(m, b)
}
func (m *NeutrinoInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeutrinoInfo.Marshal(b, m, deterministic)
}
func (m *NeutrinoInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeutrinoInfo.Merge(m, src)
}
func (m *NeutrinoInfo) XXX_Size() int {
	return xxx_messageInfo_NeutrinoInfo.Size(m)
}
func (m *NeutrinoInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NeutrinoInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NeutrinoInfo proto.InternalMessageInfo

func (m *NeutrinoInfo) GetPeers() []*PeerDesc {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *NeutrinoInfo) GetBans() []*NeutrinoBan {
	if m != nil {
		return m.Bans
	}
	return nil
}

func (m *NeutrinoInfo) GetQueries() []*NeutrinoQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *NeutrinoInfo) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *NeutrinoInfo) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *NeutrinoInfo) GetBlockTimestamp() string {
	if m != nil {
		return m.BlockTimestamp
	}
	return ""
}

func (m *NeutrinoInfo) GetIsSyncing() bool {
	if m != nil {
		return m.IsSyncing
	}
	return false
}

type WalletInfo struct {
	CurrentBlockHash      string       `protobuf:"bytes,1,opt,name=current_block_hash,json=currentBlockHash,proto3" json:"current_block_hash,omitempty"`
	CurrentHeight         int32        `protobuf:"varint,2,opt,name=current_height,json=currentHeight,proto3" json:"current_height,omitempty"`
	CurrentBlockTimestamp string       `protobuf:"bytes,3,opt,name=current_block_timestamp,json=currentBlockTimestamp,proto3" json:"current_block_timestamp,omitempty"`
	WalletVersion         int32        `protobuf:"varint,4,opt,name=wallet_version,json=walletVersion,proto3" json:"wallet_version,omitempty"`
	WalletStats           *WalletStats `protobuf:"bytes,5,opt,name=wallet_stats,json=walletStats,proto3" json:"wallet_stats,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *WalletInfo) Reset()         { *m = WalletInfo{} }
func (m *WalletInfo) String() string { return proto.CompactTextString(m) }
func (*WalletInfo) ProtoMessage()    {}
func (*WalletInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f63a845a51abc, []int{3}
}

func (m *WalletInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletInfo.Unmarshal(m, b)
}
func (m *WalletInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletInfo.Marshal(b, m, deterministic)
}
func (m *WalletInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletInfo.Merge(m, src)
}
func (m *WalletInfo) XXX_Size() int {
	return xxx_messageInfo_WalletInfo.Size(m)
}
func (m *WalletInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WalletInfo proto.InternalMessageInfo

func (m *WalletInfo) GetCurrentBlockHash() string {
	if m != nil {
		return m.CurrentBlockHash
	}
	return ""
}

func (m *WalletInfo) GetCurrentHeight() int32 {
	if m != nil {
		return m.CurrentHeight
	}
	return 0
}

func (m *WalletInfo) GetCurrentBlockTimestamp() string {
	if m != nil {
		return m.CurrentBlockTimestamp
	}
	return ""
}

func (m *WalletInfo) GetWalletVersion() int32 {
	if m != nil {
		return m.WalletVersion
	}
	return 0
}

func (m *WalletInfo) GetWalletStats() *WalletStats {
	if m != nil {
		return m.WalletStats
	}
	return nil
}

type PeerDesc struct {
	BytesReceived        uint64   `protobuf:"varint,1,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	BytesSent            uint64   `protobuf:"varint,2,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	LastRecv             string   `protobuf:"bytes,3,opt,name=last_recv,json=lastRecv,proto3" json:"last_recv,omitempty"`
	LastSend             string   `protobuf:"bytes,4,opt,name=last_send,json=lastSend,proto3" json:"last_send,omitempty"`
	Connected            bool     `protobuf:"varint,5,opt,name=connected,proto3" json:"connected,omitempty"`
	Addr                 string   `protobuf:"bytes,6,opt,name=addr,proto3" json:"addr,omitempty"`
	Inbound              bool     `protobuf:"varint,7,opt,name=inbound,proto3" json:"inbound,omitempty"`
	Na                   string   `protobuf:"bytes,8,opt,name=na,proto3" json:"na,omitempty"`
	Id                   int32    `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	UserAgent            string   `protobuf:"bytes,10,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Services             string   `protobuf:"bytes,11,opt,name=services,proto3" json:"services,omitempty"`
	VersionKnown         bool     `protobuf:"varint,12,opt,name=version_known,json=versionKnown,proto3" json:"version_known,omitempty"`
	AdvertisedProtoVer   uint32   `protobuf:"varint,13,opt,name=advertised_proto_ver,json=advertisedProtoVer,proto3" json:"advertised_proto_ver,omitempty"`
	ProtocolVersion      uint32   `protobuf:"varint,14,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	SendHeadersPreferred bool     `protobuf:"varint,15,opt,name=send_headers_preferred,json=sendHeadersPreferred,proto3" json:"send_headers_preferred,omitempty"`
	VerAckReceived       bool     `protobuf:"varint,16,opt,name=ver_ack_received,json=verAckReceived,proto3" json:"ver_ack_received,omitempty"`
	WitnessEnabled       bool     `protobuf:"varint,17,opt,name=witness_enabled,json=witnessEnabled,proto3" json:"witness_enabled,omitempty"`
	WireEncoding         string   `protobuf:"bytes,18,opt,name=wire_encoding,json=wireEncoding,proto3" json:"wire_encoding,omitempty"`
	TimeOffset           int64    `protobuf:"varint,19,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	TimeConnected        string   `protobuf:"bytes,20,opt,name=time_connected,json=timeConnected,proto3" json:"time_connected,omitempty"`
	StartingHeight       int32    `protobuf:"varint,21,opt,name=starting_height,json=startingHeight,proto3" json:"starting_height,omitempty"`
	LastBlock            int32    `protobuf:"varint,22,opt,name=last_block,json=lastBlock,proto3" json:"last_block,omitempty"`
	LastAnnouncedBlock   []byte   `protobuf:"bytes,23,opt,name=last_announced_block,json=lastAnnouncedBlock,proto3" json:"last_announced_block,omitempty"`
	LastPingNonce        uint64   `protobuf:"varint,24,opt,name=last_ping_nonce,json=lastPingNonce,proto3" json:"last_ping_nonce,omitempty"`
	LastPingTime         string   `protobuf:"bytes,25,opt,name=last_ping_time,json=lastPingTime,proto3" json:"last_ping_time,omitempty"`
	LastPingMicros       int64    `protobuf:"varint,26,opt,name=last_ping_micros,json=lastPingMicros,proto3" json:"last_ping_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerDesc) Reset()         { *m = PeerDesc{} }
func (m *PeerDesc) String() string { return proto.CompactTextString(m) }
func (*PeerDesc) ProtoMessage()    {}
func (*PeerDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f63a845a51abc, []int{4}
}

func (m *PeerDesc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerDesc.Unmarshal(m, b)
}
func (m *PeerDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerDesc.Marshal(b, m, deterministic)
}
func (m *PeerDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerDesc.Merge(m, src)
}
func (m *PeerDesc) XXX_Size() int {
	return xxx_messageInfo_PeerDesc.Size(m)
}
func (m *PeerDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerDesc.DiscardUnknown(m)
}

var xxx_messageInfo_PeerDesc proto.InternalMessageInfo

func (m *PeerDesc) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *PeerDesc) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *PeerDesc) GetLastRecv() string {
	if m != nil {
		return m.LastRecv
	}
	return ""
}

func (m *PeerDesc) GetLastSend() string {
	if m != nil {
		return m.LastSend
	}
	return ""
}

func (m *PeerDesc) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

func (m *PeerDesc) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PeerDesc) GetInbound() bool {
	if m != nil {
		return m.Inbound
	}
	return false
}

func (m *PeerDesc) GetNa() string {
	if m != nil {
		return m.Na
	}
	return ""
}

func (m *PeerDesc) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PeerDesc) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *PeerDesc) GetServices() string {
	if m != nil {
		return m.Services
	}
	return ""
}

func (m *PeerDesc) GetVersionKnown() bool {
	if m != nil {
		return m.VersionKnown
	}
	return false
}

func (m *PeerDesc) GetAdvertisedProtoVer() uint32 {
	if m != nil {
		return m.AdvertisedProtoVer
	}
	return 0
}

func (m *PeerDesc) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func (m *PeerDesc) GetSendHeadersPreferred() bool {
	if m != nil {
		return m.SendHeadersPreferred
	}
	return false
}

func (m *PeerDesc) GetVerAckReceived() bool {
	if m != nil {
		return m.VerAckReceived
	}
	return false
}

func (m *PeerDesc) GetWitnessEnabled() bool {
	if m != nil {
		return m.WitnessEnabled
	}
	return false
}

func (m *PeerDesc) GetWireEncoding() string {
	if m != nil {
		return m.WireEncoding
	}
	return ""
}

func (m *PeerDesc) GetTimeOffset() int64 {
	if m != nil {
		return m.TimeOffset
	}
	return 0
}

func (m *PeerDesc) GetTimeConnected() string {
	if m != nil {
		return m.TimeConnected
	}
	return ""
}

func (m *PeerDesc) GetStartingHeight() int32 {
	if m != nil {
		return m.StartingHeight
	}
	return 0
}

func (m *PeerDesc) GetLastBlock() int32 {
	if m != nil {
		return m.LastBlock
	}
	return 0
}

func (m *PeerDesc) GetLastAnnouncedBlock() []byte {
	if m != nil {
		return m.LastAnnouncedBlock
	}
	return nil
}

func (m *PeerDesc) GetLastPingNonce() uint64 {
	if m != nil {
		return m.LastPingNonce
	}
	return 0
}

func (m *PeerDesc) GetLastPingTime() string {
	if m != nil {
		return m.LastPingTime
	}
	return ""
}

func (m *PeerDesc) GetLastPingMicros() int64 {
	if m != nil {
		return m.LastPingMicros
	}
	return 0
}

type WalletStats struct {
	MaintenanceInProgress       bool     `protobuf:"varint,1,opt,name=maintenance_in_progress,json=maintenanceInProgress,proto3" json:"maintenance_in_progress,omitempty"`
	MaintenanceName             string   `protobuf:"bytes,2,opt,name=maintenance_name,json=maintenanceName,proto3" json:"maintenance_name,omitempty"`
	MaintenanceCycles           int32    `protobuf:"varint,3,opt,name=maintenance_cycles,json=maintenanceCycles,proto3" json:"maintenance_cycles,omitempty"`
	MaintenanceLastBlockVisited int32    `protobuf:"varint,4,opt,name=maintenance_last_block_visited,json=maintenanceLastBlockVisited,proto3" json:"maintenance_last_block_visited,omitempty"`
	TimeOfLastMaintenance       string   `protobuf:"bytes,5,opt,name=time_of_last_maintenance,json=timeOfLastMaintenance,proto3" json:"time_of_last_maintenance,omitempty"`
	Syncing                     bool     `protobuf:"varint,6,opt,name=syncing,proto3" json:"syncing,omitempty"`
	SyncStarted                 string   `protobuf:"bytes,7,opt,name=sync_started,json=syncStarted,proto3" json:"sync_started,omitempty"`
	SyncRemainingSeconds        int64    `protobuf:"varint,8,opt,name=sync_remaining_seconds,json=syncRemainingSeconds,proto3" json:"sync_remaining_seconds,omitempty"`
	SyncCurrentBlock            int32    `protobuf:"varint,9,opt,name=sync_current_block,json=syncCurrentBlock,proto3" json:"sync_current_block,omitempty"`
	SyncFrom                    int32    `protobuf:"varint,10,opt,name=sync_from,json=syncFrom,proto3" json:"sync_from,omitempty"`
	SyncTo                      int32    `protobuf:"varint,11,opt,name=sync_to,json=syncTo,proto3" json:"sync_to,omitempty"`
	BirthdayBlock               int32    `protobuf:"varint,12,opt,name=birthday_block,json=birthdayBlock,proto3" json:"birthday_block,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *WalletStats) Reset()         { *m = WalletStats{} }
func (m *WalletStats) String() string { return proto.CompactTextString(m) }
func (*WalletStats) ProtoMessage()    {}
func (*WalletStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c5f63a845a51abc, []int{5}
}

func (m *WalletStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletStats.Unmarshal(m, b)
}
func (m *WalletStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletStats.Marshal(b, m, deterministic)
}
func (m *WalletStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletStats.Merge(m, src)
}
func (m *WalletStats) XXX_Size() int {
	return xxx_messageInfo_WalletStats.Size(m)
}
func (m *WalletStats) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletStats.DiscardUnknown(m)
}

var xxx_messageInfo_WalletStats proto.InternalMessageInfo

func (m *WalletStats) GetMaintenanceInProgress() bool {
	if m != nil {
		return m.MaintenanceInProgress
	}
	return false
}

func (m *WalletStats) GetMaintenanceName() string {
	if m != nil {
		return m.MaintenanceName
	}
	return ""
}

func (m *WalletStats) GetMaintenanceCycles() int32 {
	if m != nil {
		return m.MaintenanceCycles
	}
	return 0
}

func (m *WalletStats) GetMaintenanceLastBlockVisited() int32 {
	if m != nil {
		return m.MaintenanceLastBlockVisited
	}
	return 0
}

func (m *WalletStats) GetTimeOfLastMaintenance() string {
	if m != nil {
		return m.TimeOfLastMaintenance
	}
	return ""
}

func (m *WalletStats) GetSyncing() bool {
	if m != nil {
		return m.Syncing
	}
	return false
}

func (m *WalletStats) GetSyncStarted() string {
	if m != nil {
		return m.SyncStarted
	}
	return ""
}

func (m *WalletStats) GetSyncRemainingSeconds() int64 {
	if m != nil {
		return m.SyncRemainingSeconds
	}
	return 0
}

func (m *WalletStats) GetSyncCurrentBlock() int32 {
	if m != nil {
		return m.SyncCurrentBlock
	}
	return 0
}

func (m *WalletStats) GetSyncFrom() int32 {
	if m != nil {
		return m.SyncFrom
	}
	return 0
}

func (m *WalletStats) GetSyncTo() int32 {
	if m != nil {
		return m.SyncTo
	}
	return 0
}

func (m *WalletStats) GetBirthdayBlock() int32 {
	if m != nil {
		return m.BirthdayBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*NeutrinoBan)(nil), "lnrpc.NeutrinoBan")
	proto.RegisterType((*NeutrinoQuery)(nil), "lnrpc.NeutrinoQuery")
	proto.RegisterType((*NeutrinoInfo)(nil), "lnrpc.NeutrinoInfo")
	proto.RegisterType((*WalletInfo)(nil), "lnrpc.WalletInfo")
	proto.RegisterType((*PeerDesc)(nil), "lnrpc.PeerDesc")
	proto.RegisterType((*WalletStats)(nil), "lnrpc.WalletStats")
}

func init() { proto.RegisterFile("pkt.proto", fileDescriptor_3c5f63a845a51abc) }

var fileDescriptor_3c5f63a845a51abc = []byte{
	// 1178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x56, 0xdb, 0x6e, 0x1b, 0x37,
	0x10, 0x85, 0x64, 0xc9, 0x92, 0x46, 0xd7, 0xb0, 0x76, 0xb2, 0x49, 0x9a, 0xd6, 0x55, 0x73, 0x51,
	0x8b, 0xc4, 0x29, 0x7a, 0x7d, 0x4e, 0xdc, 0x14, 0x09, 0xda, 0xb8, 0xee, 0x2a, 0x48, 0x81, 0xbe,
	0x2c, 0x28, 0xee, 0x58, 0x5a, 0x48, 0xcb, 0x95, 0x49, 0x4a, 0x86, 0xff, 0xa1, 0x7f, 0x90, 0xaf,
	0xea, 0xd7, 0xf4, 0xb5, 0x98, 0x21, 0x57, 0xda, 0xf4, 0xc1, 0xc0, 0xce, 0x39, 0x87, 0xe4, 0xdc,
	0x65, 0xe8, 0xac, 0x97, 0xee, 0x74, 0x6d, 0x0a, 0x57, 0x88, 0xe6, 0x4a, 0x9b, 0xb5, 0x1a, 0x5f,
	0x41, 0xf7, 0x1c, 0x37, 0xce, 0x64, 0xba, 0x78, 0x29, 0xb5, 0x10, 0xd0, 0x90, 0x69, 0x6a, 0xa2,
	0xda, 0x49, 0x6d, 0xd2, 0x89, 0xf9, 0x5b, 0xdc, 0x86, 0x43, 0x83, 0xd2, 0x16, 0x3a, 0xaa, 0x33,
	0x1a, 0x2c, 0x71, 0x17, 0xda, 0xa8, 0xd3, 0xc4, 0x65, 0x39, 0x46, 0x07, 0xcc, 0xb4, 0x50, 0xa7,
	0xef, 0xb2, 0x1c, 0xc5, 0x7d, 0xe8, 0xcc, 0xa4, 0x4e, 0xac, 0x2a, 0x0c, 0x46, 0x8d, 0x93, 0xda,
	0xa4, 0x19, 0xb7, 0x67, 0x52, 0x4f, 0xc9, 0x1e, 0xff, 0x53, 0x83, 0x7e, 0xf9, 0xe6, 0x1f, 0x1b,
	0x34, 0x37, 0xf4, 0xea, 0x1a, 0x71, 0xf7, 0x2a, 0x7d, 0x8b, 0x08, 0x5a, 0xaa, 0xc8, 0x73, 0xa9,
	0xd3, 0xf0, 0x6c, 0x69, 0x8a, 0x3b, 0xd0, 0x32, 0x78, 0x95, 0xe8, 0x4d, 0xce, 0xcf, 0xf6, 0xc9,
	0xa1, 0xab, 0xf3, 0x4d, 0x2e, 0x3e, 0x87, 0xae, 0x32, 0x28, 0x1d, 0x7a, 0x9f, 0x1a, 0x4c, 0x82,
	0x87, 0xd8, 0xad, 0xaf, 0xe1, 0xd6, 0x4a, 0x5a, 0x97, 0x18, 0xbc, 0xda, 0xa0, 0x75, 0x5e, 0xd6,
	0x64, 0xd9, 0x90, 0x88, 0xd8, 0xe3, 0xac, 0x7d, 0x0a, 0x22, 0x68, 0xed, 0xba, 0xd0, 0x36, 0xdc,
	0x79, 0xc8, 0xe2, 0x91, 0x17, 0x7b, 0x82, 0xd4, 0xe3, 0xbf, 0xeb, 0xd0, 0x2b, 0x63, 0x7a, 0xa3,
	0x2f, 0x0b, 0xf1, 0x08, 0x9a, 0x14, 0x86, 0x8d, 0x6a, 0x27, 0x07, 0x93, 0xee, 0xb7, 0xc3, 0x53,
	0x4e, 0xf7, 0xe9, 0x05, 0xa2, 0xf9, 0x19, 0xad, 0x8a, 0x3d, 0x2b, 0x1e, 0x43, 0x63, 0x26, 0xb5,
	0x8d, 0xea, 0xac, 0x12, 0x41, 0x55, 0xa9, 0x48, 0xcc, 0xbc, 0x38, 0x85, 0xd6, 0xd5, 0x06, 0x4d,
	0x86, 0x36, 0x3a, 0x60, 0xe9, 0xd1, 0xff, 0xa4, 0x9c, 0xc8, 0xb8, 0x14, 0x89, 0x07, 0x00, 0xb3,
	0x55, 0xa1, 0x96, 0xc9, 0x42, 0xda, 0x05, 0x67, 0xa2, 0x13, 0x77, 0x18, 0x79, 0x2d, 0xed, 0x82,
	0x4a, 0xba, 0xc0, 0x6c, 0xbe, 0x70, 0x1c, 0x7d, 0x33, 0x0e, 0x96, 0x78, 0x02, 0x43, 0x7f, 0x8c,
	0x82, 0xb5, 0x4e, 0xe6, 0x6b, 0x8e, 0xb8, 0x13, 0x0f, 0x18, 0x7e, 0x57, 0xa2, 0x74, 0x7f, 0x66,
	0x13, 0x7b, 0xa3, 0x55, 0xa6, 0xe7, 0x51, 0xeb, 0xa4, 0x36, 0x69, 0xc7, 0x9d, 0xcc, 0x4e, 0x3d,
	0x30, 0xfe, 0xb7, 0x06, 0xf0, 0xa7, 0x5c, 0xad, 0xd0, 0x71, 0x32, 0x9e, 0x82, 0x50, 0x1b, 0x63,
	0x50, 0xbb, 0xa4, 0xe2, 0x95, 0xaf, 0xf6, 0x28, 0x30, 0x2f, 0x77, 0xce, 0x3d, 0x82, 0x41, 0xa9,
	0x0e, 0x4e, 0xd6, 0xd9, 0xc9, 0x7e, 0x40, 0x5f, 0x7b, 0x5f, 0x7f, 0x84, 0x3b, 0x1f, 0x5f, 0xba,
	0xf7, 0xd9, 0x77, 0xe3, 0x71, 0xf5, 0xe6, 0xbd, 0xeb, 0x8f, 0x60, 0x70, 0xcd, 0xae, 0x25, 0x5b,
	0x34, 0x36, 0x2b, 0x74, 0x68, 0xd0, 0xbe, 0x47, 0xdf, 0x7b, 0x50, 0xfc, 0x00, 0xbd, 0x20, 0xb3,
	0x4e, 0x3a, 0xcb, 0x89, 0xda, 0x57, 0xc8, 0x07, 0x37, 0x25, 0x26, 0xee, 0x5e, 0xef, 0x8d, 0xf1,
	0x87, 0x16, 0xb4, 0xcb, 0x22, 0xd3, 0x53, 0xb3, 0x1b, 0x87, 0x36, 0x31, 0xa8, 0x30, 0xdb, 0x62,
	0xca, 0x31, 0x37, 0xe2, 0x3e, 0xa3, 0x71, 0x00, 0xb9, 0x58, 0x2c, 0xb3, 0xa8, 0x7d, 0xb0, 0x8d,
	0xb8, 0xc3, 0xc8, 0x14, 0xb5, 0xa3, 0x61, 0x0a, 0x9d, 0xa8, 0xb6, 0x21, 0xb4, 0xb6, 0x6f, 0x40,
	0xb5, 0xdd, 0x91, 0x16, 0x75, 0x1a, 0xea, 0xcc, 0xe4, 0x14, 0x75, 0x2a, 0x3e, 0x85, 0x8e, 0x2a,
	0xb4, 0x46, 0xe5, 0x30, 0xe5, 0x00, 0xda, 0xf1, 0x1e, 0xd8, 0xcd, 0xfa, 0x61, 0x65, 0xd6, 0x23,
	0x68, 0x65, 0x7a, 0x56, 0x6c, 0x74, 0x1a, 0x8a, 0x5a, 0x9a, 0x62, 0x00, 0x75, 0x2d, 0xa3, 0x36,
	0x6b, 0xeb, 0x5a, 0x92, 0x9d, 0xa5, 0x51, 0x87, 0x53, 0x57, 0xcf, 0x38, 0x88, 0x8d, 0x45, 0x93,
	0xc8, 0x39, 0x05, 0x01, 0xbe, 0xe3, 0x08, 0x79, 0x41, 0x80, 0xb8, 0x07, 0x6d, 0x8b, 0x66, 0x9b,
	0x29, 0xb4, 0x51, 0xd7, 0xbb, 0x59, 0xda, 0xe2, 0x4b, 0xe8, 0x87, 0x52, 0x24, 0x4b, 0x5d, 0x5c,
	0xeb, 0xa8, 0xc7, 0x4f, 0xf7, 0x02, 0xf8, 0x2b, 0x61, 0xe2, 0x1b, 0x38, 0x92, 0xe9, 0x16, 0x8d,
	0xcb, 0x2c, 0xa6, 0x09, 0xef, 0x30, 0x2a, 0x60, 0xd4, 0xe7, 0x89, 0x14, 0x7b, 0xee, 0x82, 0xa8,
	0xf7, 0x68, 0xc4, 0x57, 0x30, 0x62, 0x99, 0x2a, 0x56, 0xbb, 0x52, 0x0f, 0xfc, 0xb0, 0x97, 0x78,
	0x59, 0xec, 0xef, 0xe1, 0x36, 0x25, 0x30, 0x59, 0xa0, 0x4c, 0xd1, 0xd8, 0x64, 0x6d, 0xf0, 0x12,
	0x8d, 0xc1, 0x34, 0x1a, 0xb2, 0x2b, 0x47, 0xc4, 0xbe, 0xf6, 0xe4, 0x45, 0xc9, 0x89, 0x09, 0x8c,
	0xb6, 0x14, 0xb1, 0x5a, 0xee, 0x0b, 0x3c, 0x62, 0xfd, 0x60, 0x8b, 0xe6, 0x85, 0x5a, 0xee, 0x2a,
	0xfc, 0x04, 0x86, 0xd7, 0x99, 0xd3, 0x68, 0x6d, 0x82, 0x5a, 0xce, 0x56, 0x98, 0x46, 0xb7, 0xbc,
	0x30, 0xc0, 0xaf, 0x3c, 0x4a, 0xa9, 0xb8, 0xce, 0x0c, 0x26, 0xa8, 0x55, 0x91, 0xd2, 0x68, 0x09,
	0xce, 0x55, 0x8f, 0xc0, 0x57, 0x01, 0xa3, 0x3d, 0x47, 0xbd, 0x9e, 0x14, 0x97, 0x97, 0x16, 0x5d,
	0xf4, 0xc9, 0x49, 0x6d, 0x72, 0x10, 0x03, 0x41, 0xbf, 0x33, 0x42, 0x7d, 0xc7, 0x82, 0x7d, 0xf1,
	0x8f, 0xf8, 0x9a, 0x3e, 0xa1, 0x67, 0xbb, 0x06, 0x78, 0x02, 0x43, 0xeb, 0xa4, 0x71, 0x99, 0x9e,
	0x97, 0x93, 0x76, 0xcc, 0xf5, 0x1c, 0x94, 0x70, 0x18, 0xb5, 0x07, 0x00, 0xdc, 0x64, 0x3c, 0x67,
	0xd1, 0x6d, 0xd6, 0x70, 0xdb, 0xf1, 0x68, 0x51, 0x69, 0x98, 0x96, 0x5a, 0x17, 0x1b, 0xad, 0x30,
	0x0d, 0xc2, 0x3b, 0x27, 0xb5, 0x49, 0x2f, 0xe6, 0x35, 0xfa, 0xa2, 0xa4, 0xfc, 0x89, 0xc7, 0xc0,
	0xfb, 0x36, 0x59, 0xd3, 0xd3, 0xba, 0xd0, 0x0a, 0xa3, 0xc8, 0x4f, 0x06, 0xc1, 0x17, 0x99, 0x9e,
	0x9f, 0x13, 0x28, 0x1e, 0xc2, 0x60, 0xaf, 0xe3, 0x05, 0x7c, 0xd7, 0xe7, 0xa3, 0x94, 0xf1, 0xaa,
	0x9e, 0xc0, 0x68, 0xaf, 0xca, 0x33, 0x65, 0x0a, 0x1b, 0xdd, 0xe3, 0xa4, 0x0c, 0x4a, 0xdd, 0x5b,
	0x46, 0xc7, 0x1f, 0x1a, 0xd0, 0xad, 0x8c, 0x2e, 0xed, 0x90, 0x5c, 0x66, 0xda, 0xa1, 0x96, 0x5a,
	0x61, 0x92, 0x69, 0x6a, 0xac, 0xb9, 0x41, 0x6b, 0x79, 0x52, 0xdb, 0xf1, 0x71, 0x85, 0x7e, 0xa3,
	0x2f, 0x02, 0x49, 0xad, 0x55, 0x3d, 0xa7, 0x65, 0x8e, 0xe1, 0x57, 0x6a, 0x58, 0xc1, 0xcf, 0x65,
	0x8e, 0xe2, 0x19, 0x88, 0xaa, 0x54, 0xdd, 0xa8, 0x15, 0x2f, 0x71, 0xca, 0xe1, 0xad, 0x0a, 0x73,
	0xc6, 0x84, 0x38, 0x83, 0xcf, 0xaa, 0xf2, 0x7d, 0xda, 0x93, 0x6d, 0x66, 0x33, 0x2a, 0xa5, 0xdf,
	0x56, 0xf7, 0x2b, 0xaa, 0xdf, 0xca, 0x4a, 0xbc, 0xf7, 0x12, 0xf1, 0x13, 0x44, 0xa1, 0x41, 0xfc,
	0x05, 0x15, 0x2d, 0xaf, 0x81, 0x4e, 0x7c, 0xec, 0xbb, 0x85, 0x4e, 0xbe, 0xdd, 0x93, 0x34, 0xfe,
	0xe5, 0x4e, 0x3f, 0xf4, 0xe3, 0x1f, 0x4c, 0xf1, 0x05, 0xf4, 0xe8, 0x33, 0xe1, 0xce, 0x40, 0xbf,
	0x1d, 0x3a, 0x71, 0x97, 0xb0, 0xa9, 0x87, 0x78, 0x88, 0x48, 0x62, 0x90, 0xde, 0xa3, 0x5a, 0x58,
	0x54, 0x85, 0x4e, 0x2d, 0x6f, 0x8d, 0x83, 0xf8, 0x88, 0xd8, 0xb8, 0x24, 0xa7, 0x9e, 0xa3, 0xdf,
	0x06, 0x3e, 0xf5, 0xd1, 0x2e, 0x0f, 0x7b, 0x65, 0x44, 0xcc, 0x59, 0x65, 0x8b, 0xd3, 0xba, 0x63,
	0xf5, 0xa5, 0x29, 0x72, 0x5e, 0x32, 0xcd, 0xb8, 0x4d, 0xc0, 0x2f, 0xa6, 0xc8, 0xe9, 0x1f, 0x03,
	0x26, 0x5d, 0xc1, 0x2b, 0xa6, 0x19, 0x1f, 0x92, 0xf9, 0xae, 0xe0, 0x3d, 0x9c, 0x19, 0xb7, 0x48,
	0xe5, 0x4d, 0xb8, 0xbf, 0xe7, 0x57, 0x7e, 0x89, 0xf2, 0xe5, 0x2f, 0x1f, 0xfe, 0x35, 0x9e, 0x67,
	0x6e, 0xb1, 0x99, 0x9d, 0xaa, 0x22, 0x7f, 0xbe, 0x5e, 0xba, 0x67, 0x4a, 0xda, 0x05, 0x7d, 0xa4,
	0xcf, 0x57, 0x9a, 0xfe, 0xcc, 0x5a, 0xcd, 0x0e, 0x79, 0x79, 0x7c, 0xf7, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x57, 0xdd, 0x22, 0x4c, 0x09, 0x00, 0x00,
}
