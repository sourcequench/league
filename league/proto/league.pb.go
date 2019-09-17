// Code generated by protoc-gen-go. DO NOT EDIT.
// source: league.proto

/*
Package league is a generated protocol buffer package.

It is generated from these files:
	league.proto

It has these top-level messages:
	Player
	Match
	Season
	Roster
*/
package league

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return proto.EnumName(Match_Type_name, int32(x))
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
	return proto.EnumName(Match_Offset_name, int32(x))
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
	return proto.EnumName(Match_Forfeit_name, int32(x))
}
func (Match_Forfeit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

type Player struct {
	PlayerId  int32   `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	FirstName string  `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string  `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	NickName  string  `protobuf:"bytes,4,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	Email     string  `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Phone     string  `protobuf:"bytes,6,opt,name=phone" json:"phone,omitempty"`
	Mu        float32 `protobuf:"fixed32,7,opt,name=mu" json:"mu,omitempty"`
	Sigma     float32 `protobuf:"fixed32,8,opt,name=sigma" json:"sigma,omitempty"`
	//  repeated Match matches = 9;
	WaitList bool `protobuf:"varint,10,opt,name=wait_list,json=waitList" json:"wait_list,omitempty"`
	Active   bool `protobuf:"varint,11,opt,name=active" json:"active,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
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

func (m *Player) GetMu() float32 {
	if m != nil {
		return m.Mu
	}
	return 0
}

func (m *Player) GetSigma() float32 {
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
	P1Id            int32         `protobuf:"varint,1,opt,name=p1_id,json=p1Id" json:"p1_id,omitempty"`
	P2Id            int32         `protobuf:"varint,2,opt,name=p2_id,json=p2Id" json:"p2_id,omitempty"`
	MyGames         int32         `protobuf:"varint,4,opt,name=my_games,json=myGames" json:"my_games,omitempty"`
	YourGames       int32         `protobuf:"varint,5,opt,name=your_games,json=yourGames" json:"your_games,omitempty"`
	MyGamesNeeded   int32         `protobuf:"varint,6,opt,name=my_games_needed,json=myGamesNeeded" json:"my_games_needed,omitempty"`
	YourGamesNeeded int32         `protobuf:"varint,7,opt,name=your_games_needed,json=yourGamesNeeded" json:"your_games_needed,omitempty"`
	MatchType       Match_Type    `protobuf:"varint,8,opt,name=match_type,json=matchType,enum=league.Match_Type" json:"match_type,omitempty"`
	Date            string        `protobuf:"bytes,9,opt,name=date" json:"date,omitempty"`
	Forfeit         Match_Forfeit `protobuf:"varint,10,opt,name=forfeit,enum=league.Match_Forfeit" json:"forfeit,omitempty"`
	Offset          Match_Offset  `protobuf:"varint,11,opt,name=offset,enum=league.Match_Offset" json:"offset,omitempty"`
}

func (m *Match) Reset()                    { *m = Match{} }
func (m *Match) String() string            { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()               {}
func (*Match) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Match) GetP1Id() int32 {
	if m != nil {
		return m.P1Id
	}
	return 0
}

func (m *Match) GetP2Id() int32 {
	if m != nil {
		return m.P2Id
	}
	return 0
}

func (m *Match) GetMyGames() int32 {
	if m != nil {
		return m.MyGames
	}
	return 0
}

func (m *Match) GetYourGames() int32 {
	if m != nil {
		return m.YourGames
	}
	return 0
}

func (m *Match) GetMyGamesNeeded() int32 {
	if m != nil {
		return m.MyGamesNeeded
	}
	return 0
}

func (m *Match) GetYourGamesNeeded() int32 {
	if m != nil {
		return m.YourGamesNeeded
	}
	return 0
}

func (m *Match) GetMatchType() Match_Type {
	if m != nil {
		return m.MatchType
	}
	return Match_SEASON
}

func (m *Match) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
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
func (m *Season) String() string            { return proto.CompactTextString(m) }
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
func (m *Roster) String() string            { return proto.CompactTextString(m) }
func (*Roster) ProtoMessage()               {}
func (*Roster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Roster) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*Player)(nil), "league.Player")
	proto.RegisterType((*Match)(nil), "league.Match")
	proto.RegisterType((*Season)(nil), "league.Season")
	proto.RegisterType((*Roster)(nil), "league.Roster")
	proto.RegisterEnum("league.Match_Type", Match_Type_name, Match_Type_value)
	proto.RegisterEnum("league.Match_Offset", Match_Offset_name, Match_Offset_value)
	proto.RegisterEnum("league.Match_Forfeit", Match_Forfeit_name, Match_Forfeit_value)
}

func init() { proto.RegisterFile("league.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0x6b, 0x37, 0xfe, 0x9b, 0xb6, 0x69, 0xbe, 0xfd, 0x0a, 0x72, 0x55, 0x21, 0x45, 0x46,
	0x02, 0x0b, 0xa1, 0xa0, 0x98, 0x2b, 0x88, 0xd4, 0x04, 0x22, 0xda, 0x24, 0xda, 0xc0, 0x01, 0x47,
	0xd6, 0x12, 0x4f, 0x52, 0x8b, 0xf8, 0x47, 0xf6, 0x06, 0xe4, 0x0b, 0x40, 0xe2, 0xf6, 0xb8, 0x23,
	0xb4, 0xb3, 0x36, 0x51, 0xce, 0x76, 0xde, 0xe7, 0x9d, 0xb1, 0x67, 0x76, 0x16, 0x2e, 0xf7, 0x28,
	0x76, 0x07, 0x1c, 0x95, 0x55, 0x21, 0x0b, 0x66, 0xeb, 0x28, 0xf8, 0x65, 0x82, 0xbd, 0xda, 0x8b,
	0x06, 0x2b, 0x76, 0x07, 0x5e, 0x49, 0xa7, 0x38, 0x4d, 0x7c, 0x63, 0x68, 0x84, 0x16, 0x77, 0xb5,
	0x30, 0x4f, 0xd8, 0x0b, 0x80, 0x6d, 0x5a, 0xd5, 0x32, 0xce, 0x45, 0x86, 0xbe, 0x39, 0x34, 0x42,
	0x8f, 0x7b, 0xa4, 0x2c, 0x44, 0x86, 0x2a, 0x77, 0x2f, 0x3a, 0x7a, 0x4e, 0xd4, 0x55, 0x42, 0x07,
	0xf3, 0x74, 0xf3, 0x5d, 0xc3, 0x9e, 0x86, 0x4a, 0x20, 0x78, 0x03, 0x16, 0x66, 0x22, 0xdd, 0xfb,
	0x16, 0x01, 0x1d, 0x28, 0xb5, 0x7c, 0x2a, 0x72, 0xf4, 0x6d, 0xad, 0x52, 0xc0, 0xfa, 0x60, 0x66,
	0x07, 0xdf, 0x19, 0x1a, 0xa1, 0xc9, 0xcd, 0xec, 0xa0, 0x5c, 0x75, 0xba, 0xcb, 0x84, 0xef, 0x92,
	0xa4, 0x03, 0xf5, 0xb9, 0x9f, 0x22, 0x95, 0xf1, 0x3e, 0xad, 0xa5, 0x0f, 0x43, 0x23, 0x74, 0xb9,
	0xab, 0x84, 0x87, 0xb4, 0x96, 0xec, 0x39, 0xd8, 0x62, 0x23, 0xd3, 0x1f, 0xe8, 0x5f, 0x10, 0x69,
	0xa3, 0xe0, 0xcf, 0x39, 0x58, 0x8f, 0x42, 0x6e, 0x9e, 0xd8, 0xff, 0x60, 0x95, 0xe3, 0xe3, 0x08,
	0x7a, 0xe5, 0x78, 0x9e, 0x90, 0x18, 0x29, 0xd1, 0x6c, 0xc5, 0x68, 0x9e, 0xb0, 0x5b, 0x70, 0xb3,
	0x26, 0xde, 0x89, 0x0c, 0x6b, 0x6a, 0xcb, 0xe2, 0x4e, 0xd6, 0x7c, 0x50, 0xa1, 0x1a, 0x57, 0x53,
	0x1c, 0xaa, 0x16, 0x5a, 0x04, 0x3d, 0xa5, 0x68, 0xfc, 0x0a, 0xae, 0xbb, 0xcc, 0x38, 0x47, 0x4c,
	0x30, 0xa1, 0x46, 0x2d, 0x7e, 0xd5, 0x16, 0x58, 0x90, 0xc8, 0xde, 0xc0, 0x7f, 0xc7, 0x32, 0x9d,
	0xd3, 0x21, 0xe7, 0xf5, 0xbf, 0x6a, 0xad, 0x77, 0x0c, 0x90, 0xa9, 0x06, 0x62, 0xd9, 0x94, 0x48,
	0x13, 0xe9, 0x47, 0x6c, 0xd4, 0x5e, 0x3a, 0xb5, 0x36, 0xfa, 0xdc, 0x94, 0xc8, 0x3d, 0x72, 0xa9,
	0x23, 0x63, 0xd0, 0x4b, 0x84, 0x44, 0xdf, 0xa3, 0x21, 0xd3, 0x99, 0xbd, 0x03, 0x67, 0x5b, 0x54,
	0x5b, 0x4c, 0xf5, 0xec, 0xfa, 0xd1, 0xb3, 0xd3, 0x1a, 0x33, 0x0d, 0x79, 0xe7, 0x62, 0x6f, 0xc1,
	0x2e, 0xb6, 0xdb, 0x1a, 0x25, 0x4d, 0xb4, 0x1f, 0xdd, 0x9c, 0xfa, 0x97, 0xc4, 0x78, 0xeb, 0x09,
	0x86, 0xd0, 0xa3, 0x4f, 0x03, 0xd8, 0xeb, 0xe9, 0x64, 0xbd, 0x5c, 0x0c, 0xce, 0xd8, 0x25, 0xb8,
	0xab, 0x87, 0xc9, 0xd7, 0xe5, 0x6c, 0xb6, 0x1e, 0x18, 0xc1, 0x4b, 0xb0, 0x75, 0x8e, 0xf2, 0x3c,
	0x4e, 0x3e, 0x4d, 0xbf, 0xac, 0x06, 0x67, 0xec, 0x0a, 0x3c, 0xe5, 0x99, 0x7c, 0x9c, 0x4e, 0xee,
	0xc9, 0xe4, 0xb4, 0x3f, 0xc2, 0x2e, 0xc0, 0x51, 0x64, 0xca, 0xc7, 0x83, 0xb3, 0x63, 0x10, 0x0d,
	0x8c, 0xe0, 0xb7, 0x01, 0xf6, 0x1a, 0x45, 0x5d, 0xe4, 0xea, 0x3e, 0x6a, 0x29, 0x2a, 0x19, 0x53,
	0xbf, 0x86, 0x5e, 0x5f, 0x52, 0xee, 0x55, 0xd3, 0xb7, 0xe0, 0x62, 0x9e, 0x68, 0xa8, 0x77, 0xdb,
	0xc1, 0x3c, 0x21, 0x74, 0x07, 0x9e, 0x4c, 0x33, 0x8c, 0xeb, 0x12, 0x37, 0xdd, 0x66, 0x2b, 0x61,
	0x5d, 0xe2, 0x86, 0xbd, 0x06, 0x87, 0xa6, 0x49, 0x0b, 0x70, 0x1e, 0x5e, 0x44, 0x57, 0x27, 0xcd,
	0xf3, 0x8e, 0x06, 0x11, 0xd8, 0xbc, 0xa8, 0x25, 0x56, 0x2c, 0x04, 0x47, 0x3f, 0xaa, 0xda, 0x37,
	0x28, 0xa5, 0xdf, 0xa5, 0xe8, 0x67, 0xc8, 0x3b, 0xfc, 0xcd, 0xa6, 0x97, 0xfa, 0xfe, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0x4c, 0x4a, 0xad, 0xb9, 0x03, 0x00, 0x00,
}