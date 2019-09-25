// Code generated by protoc-gen-go. DO NOT EDIT.
// source: league.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	league.proto

It has these top-level messages:
	Player
	Match
	Season
	Roster
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Match_Type int32

const (
	Match_SEASON   Match_Type = 0
	Match_PLAYOFFS Match_Type = 1
)

var Match_Type_name = map[int32]string{
	0: "SEASON",
	1: "PLAYOFFS",
}
var Match_Type_value = map[string]int32{
	"SEASON":   0,
	"PLAYOFFS": 1,
}

func (x Match_Type) String() string {
	return proto1.EnumName(Match_Type_name, int32(x))
}
func (Match_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Match_Offset int32

const (
	Match_MAKEUP    Match_Offset = 0
	Match_PLAYAHEAD Match_Offset = 1
)

var Match_Offset_name = map[int32]string{
	0: "MAKEUP",
	1: "PLAYAHEAD",
}
var Match_Offset_value = map[string]int32{
	"MAKEUP":    0,
	"PLAYAHEAD": 1,
}

func (x Match_Offset) String() string {
	return proto1.EnumName(Match_Offset_name, int32(x))
}
func (Match_Offset) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type Match_Forfeit int32

const (
	Match_PLAYER1 Match_Forfeit = 0
	Match_PLAYER2 Match_Forfeit = 1
)

var Match_Forfeit_name = map[int32]string{
	0: "PLAYER1",
	1: "PLAYER2",
}
var Match_Forfeit_value = map[string]int32{
	"PLAYER1": 0,
	"PLAYER2": 1,
}

func (x Match_Forfeit) String() string {
	return proto1.EnumName(Match_Forfeit_name, int32(x))
}
func (Match_Forfeit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

type Player struct {
	PlayerId  int32   `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	FirstName string  `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string  `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	NickName  string  `protobuf:"bytes,4,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Email     string  `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Phone     string  `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	Mu        float64 `protobuf:"fixed64,7,opt,name=mu" json:"mu,omitempty"`
	Sigma     float64 `protobuf:"fixed64,8,opt,name=sigma" json:"sigma,omitempty"`
	//  repeated Match matches = 9;
	WaitList bool `protobuf:"varint,10,opt,name=wait_list,json=waitList" json:"wait_list,omitempty"`
	Active   bool `protobuf:"varint,11,opt,name=active" json:"active,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto1.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Player) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Player) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Player) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Player) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Player) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Player) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Player) GetMu() float64 {
	if m != nil {
		return m.Mu
	}
	return 0
}

func (m *Player) GetSigma() float64 {
	if m != nil {
		return m.Sigma
	}
	return 0
}

func (m *Player) GetWaitList() bool {
	if m != nil {
		return m.WaitList
	}
	return false
}

func (m *Player) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Match struct {
	P1Name    string        `protobuf:"bytes,1,opt,name=p1_name,json=p1Name" json:"p1_name,omitempty"`
	P2Name    string        `protobuf:"bytes,2,opt,name=p2_name,json=p2Name" json:"p2_name,omitempty"`
	P1Needs   float64       `protobuf:"fixed64,4,opt,name=p1_needs,json=p1Needs" json:"p1_needs,omitempty"`
	P2Needs   float64       `protobuf:"fixed64,5,opt,name=p2_needs,json=p2Needs" json:"p2_needs,omitempty"`
	P1Got     float64       `protobuf:"fixed64,6,opt,name=p1_got,json=p1Got" json:"p1_got,omitempty"`
	P2Got     float64       `protobuf:"fixed64,7,opt,name=p2_got,json=p2Got" json:"p2_got,omitempty"`
	P1Skill   float64       `protobuf:"fixed64,8,opt,name=p1_skill,json=p1Skill" json:"p1_skill,omitempty"`
	P2Skill   float64       `protobuf:"fixed64,9,opt,name=p2_skill,json=p2Skill" json:"p2_skill,omitempty"`
	Date      string        `protobuf:"bytes,10,opt,name=date" json:"date,omitempty"`
	MatchType Match_Type    `protobuf:"varint,11,opt,name=match_type,json=matchType,enum=proto.Match_Type" json:"match_type,omitempty"`
	Forfeit   Match_Forfeit `protobuf:"varint,12,opt,name=forfeit,enum=proto.Match_Forfeit" json:"forfeit,omitempty"`
	Offset    Match_Offset  `protobuf:"varint,13,opt,name=offset,enum=proto.Match_Offset" json:"offset,omitempty"`
}

func (m *Match) Reset()                    { *m = Match{} }
func (m *Match) String() string            { return proto1.CompactTextString(m) }
func (*Match) ProtoMessage()               {}
func (*Match) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Match) GetP1Name() string {
	if m != nil {
		return m.P1Name
	}
	return ""
}

func (m *Match) GetP2Name() string {
	if m != nil {
		return m.P2Name
	}
	return ""
}

func (m *Match) GetP1Needs() float64 {
	if m != nil {
		return m.P1Needs
	}
	return 0
}

func (m *Match) GetP2Needs() float64 {
	if m != nil {
		return m.P2Needs
	}
	return 0
}

func (m *Match) GetP1Got() float64 {
	if m != nil {
		return m.P1Got
	}
	return 0
}

func (m *Match) GetP2Got() float64 {
	if m != nil {
		return m.P2Got
	}
	return 0
}

func (m *Match) GetP1Skill() float64 {
	if m != nil {
		return m.P1Skill
	}
	return 0
}

func (m *Match) GetP2Skill() float64 {
	if m != nil {
		return m.P2Skill
	}
	return 0
}

func (m *Match) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Match) GetMatchType() Match_Type {
	if m != nil {
		return m.MatchType
	}
	return Match_SEASON
}

func (m *Match) GetForfeit() Match_Forfeit {
	if m != nil {
		return m.Forfeit
	}
	return Match_PLAYER1
}

func (m *Match) GetOffset() Match_Offset {
	if m != nil {
		return m.Offset
	}
	return Match_MAKEUP
}

type Season struct {
	StartDate string   `protobuf:"bytes,1,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	EndDate   string   `protobuf:"bytes,2,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
	TimeSpec  string   `protobuf:"bytes,3,opt,name=time_spec,json=timeSpec" json:"time_spec,omitempty"`
	Matches   []*Match `protobuf:"bytes,4,rep,name=matches" json:"matches,omitempty"`
}

func (m *Season) Reset()                    { *m = Season{} }
func (m *Season) String() string            { return proto1.CompactTextString(m) }
func (*Season) ProtoMessage()               {}
func (*Season) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Season) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Season) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *Season) GetTimeSpec() string {
	if m != nil {
		return m.TimeSpec
	}
	return ""
}

func (m *Season) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

type Roster struct {
	Players []*Player `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (m *Roster) Reset()                    { *m = Roster{} }
func (m *Roster) String() string            { return proto1.CompactTextString(m) }
func (*Roster) ProtoMessage()               {}
func (*Roster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Roster) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto1.RegisterType((*Player)(nil), "proto.Player")
	proto1.RegisterType((*Match)(nil), "proto.Match")
	proto1.RegisterType((*Season)(nil), "proto.Season")
	proto1.RegisterType((*Roster)(nil), "proto.Roster")
	proto1.RegisterEnum("proto.Match_Type", Match_Type_name, Match_Type_value)
	proto1.RegisterEnum("proto.Match_Offset", Match_Offset_name, Match_Offset_value)
	proto1.RegisterEnum("proto.Match_Forfeit", Match_Forfeit_name, Match_Forfeit_value)
}

func init() { proto1.RegisterFile("league.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdd, 0x6e, 0xd4, 0x3c,
	0x10, 0x86, 0xeb, 0x6d, 0xe3, 0x24, 0xd3, 0x6d, 0xb5, 0x9f, 0xbf, 0x02, 0xa9, 0x2a, 0xa4, 0x55,
	0x90, 0x60, 0x25, 0xa4, 0x15, 0x09, 0x57, 0xb0, 0x52, 0xb7, 0x80, 0xe8, 0x9f, 0x1c, 0x38, 0xe0,
	0x28, 0x32, 0x1b, 0x6f, 0x6b, 0x35, 0x3f, 0x56, 0xec, 0x82, 0x7a, 0x01, 0x88, 0xfb, 0xe4, 0x4a,
	0x90, 0xc7, 0x89, 0x68, 0x8f, 0xe2, 0x79, 0x9f, 0x77, 0x34, 0x9e, 0xc9, 0x18, 0xa6, 0xb5, 0x14,
	0x37, 0xf7, 0x72, 0xa9, 0xfb, 0xce, 0x76, 0x2c, 0xc0, 0x4f, 0xfa, 0x6b, 0x02, 0xf4, 0xba, 0x16,
	0x0f, 0xb2, 0x67, 0x27, 0x10, 0x6b, 0x3c, 0x95, 0xaa, 0x4a, 0xc8, 0x9c, 0x2c, 0x02, 0x1e, 0x79,
	0xe1, 0x53, 0xc5, 0x5e, 0x02, 0x6c, 0x55, 0x6f, 0x6c, 0xd9, 0x8a, 0x46, 0x26, 0x93, 0x39, 0x59,
	0xc4, 0x3c, 0x46, 0xe5, 0x52, 0x34, 0xd2, 0xe5, 0xd6, 0x62, 0xa4, 0xbb, 0x48, 0x23, 0x27, 0x8c,
	0xb0, 0x55, 0x9b, 0x3b, 0x0f, 0xf7, 0x3c, 0x74, 0x02, 0xc2, 0x23, 0x08, 0x64, 0x23, 0x54, 0x9d,
	0x04, 0x08, 0x7c, 0xe0, 0x54, 0x7d, 0xdb, 0xb5, 0x32, 0xa1, 0x5e, 0xc5, 0x80, 0x1d, 0xc2, 0xa4,
	0xb9, 0x4f, 0xc2, 0x39, 0x59, 0x10, 0x3e, 0x69, 0xee, 0x9d, 0xcb, 0xa8, 0x9b, 0x46, 0x24, 0x11,
	0x4a, 0x3e, 0x70, 0xe5, 0x7e, 0x0a, 0x65, 0xcb, 0x5a, 0x19, 0x9b, 0xc0, 0x9c, 0x2c, 0x22, 0x1e,
	0x39, 0xe1, 0x5c, 0x19, 0xcb, 0x9e, 0x03, 0x15, 0x1b, 0xab, 0x7e, 0xc8, 0x64, 0x1f, 0xc9, 0x10,
	0xa5, 0x7f, 0x76, 0x21, 0xb8, 0x10, 0x76, 0x73, 0xcb, 0x5e, 0x40, 0xa8, 0x33, 0x7f, 0x57, 0x82,
	0xc5, 0xa9, 0xce, 0xf0, 0xa6, 0x0e, 0xe4, 0x8f, 0xfb, 0xa7, 0x3a, 0x47, 0x70, 0x0c, 0x91, 0xcb,
	0x90, 0xb2, 0x32, 0xd8, 0x1e, 0xe1, 0xa1, 0xce, 0x2e, 0x5d, 0x88, 0x28, 0x1f, 0x50, 0x30, 0xa0,
	0xdc, 0xa3, 0x67, 0x40, 0x75, 0x56, 0xde, 0x74, 0x16, 0x7b, 0x24, 0x3c, 0xd0, 0xd9, 0x87, 0xce,
	0xa2, 0x9c, 0xa3, 0x1c, 0x0e, 0x72, 0xee, 0x64, 0x5f, 0xc3, 0xdc, 0xa9, 0xba, 0x1e, 0xba, 0x0d,
	0x75, 0x56, 0xb8, 0x70, 0xa8, 0xe1, 0x51, 0x3c, 0xd6, 0xf0, 0x88, 0xc1, 0x5e, 0x25, 0xac, 0xc4,
	0x29, 0xc4, 0x1c, 0xcf, 0xec, 0x1d, 0x40, 0xe3, 0x1a, 0x2d, 0xed, 0x83, 0xf6, 0x53, 0x38, 0xcc,
	0xff, 0xf3, 0x4b, 0xb1, 0xc4, 0x09, 0x2c, 0xbf, 0x3c, 0x68, 0xc9, 0x63, 0x34, 0xb9, 0x23, 0x5b,
	0x42, 0xb8, 0xed, 0xfa, 0xad, 0x54, 0x36, 0x99, 0xa2, 0xfd, 0xe8, 0x89, 0xfd, 0xcc, 0x33, 0x3e,
	0x9a, 0xd8, 0x5b, 0xa0, 0xdd, 0x76, 0x6b, 0xa4, 0x4d, 0x0e, 0xd0, 0xfe, 0xff, 0x13, 0xfb, 0x15,
	0x22, 0x3e, 0x58, 0xd2, 0x39, 0xec, 0x61, 0x11, 0x00, 0x5a, 0xac, 0x57, 0xc5, 0xd5, 0xe5, 0x6c,
	0x87, 0x4d, 0x21, 0xba, 0x3e, 0x5f, 0x7d, 0xbb, 0x3a, 0x3b, 0x2b, 0x66, 0x24, 0x7d, 0x05, 0xd4,
	0xe7, 0x38, 0xcf, 0xc5, 0xea, 0xf3, 0xfa, 0xeb, 0xf5, 0x6c, 0x87, 0x1d, 0x40, 0xec, 0x3c, 0xab,
	0x8f, 0xeb, 0xd5, 0x29, 0x9a, 0xc2, 0xe1, 0x1e, 0x6c, 0x1f, 0x42, 0x47, 0xd6, 0x3c, 0x9b, 0xed,
	0xfc, 0x0b, 0xf2, 0x19, 0x49, 0x7f, 0x13, 0xa0, 0x85, 0x14, 0xa6, 0x6b, 0xdd, 0x3e, 0x1b, 0x2b,
	0x7a, 0x5b, 0xe2, 0x7c, 0xfc, 0x8f, 0x8e, 0x51, 0x39, 0x75, 0x43, 0x3a, 0x86, 0x48, 0xb6, 0x95,
	0x87, 0xfe, 0x67, 0x87, 0xb2, 0xad, 0x10, 0x9d, 0x40, 0x6c, 0x55, 0x23, 0x4b, 0xa3, 0xe5, 0x66,
	0x5c, 0x75, 0x27, 0x14, 0x5a, 0x6e, 0xd8, 0x6b, 0x08, 0x71, 0x6e, 0xd2, 0x6d, 0xc2, 0xee, 0x62,
	0x3f, 0x9f, 0x3e, 0xee, 0x9d, 0x8f, 0x30, 0xcd, 0x80, 0xf2, 0xce, 0x58, 0xd9, 0xb3, 0x37, 0x10,
	0xfa, 0x47, 0x66, 0x12, 0x82, 0x19, 0x07, 0x43, 0x86, 0x7f, 0x95, 0x7c, 0xa4, 0xdf, 0x29, 0xca,
	0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x98, 0x86, 0x26, 0x83, 0xc7, 0x03, 0x00, 0x00,
}
