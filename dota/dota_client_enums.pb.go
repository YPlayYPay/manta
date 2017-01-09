// Code generated by protoc-gen-go.
// source: dota_client_enums.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ETournamentTemplate int32

const (
	ETournamentTemplate_k_ETournamentTemplate_None          ETournamentTemplate = 0
	ETournamentTemplate_k_ETournamentTemplate_AutomatedWin3 ETournamentTemplate = 1
)

var ETournamentTemplate_name = map[int32]string{
	0: "k_ETournamentTemplate_None",
	1: "k_ETournamentTemplate_AutomatedWin3",
}
var ETournamentTemplate_value = map[string]int32{
	"k_ETournamentTemplate_None":          0,
	"k_ETournamentTemplate_AutomatedWin3": 1,
}

func (x ETournamentTemplate) Enum() *ETournamentTemplate {
	p := new(ETournamentTemplate)
	*p = x
	return p
}
func (x ETournamentTemplate) String() string {
	return proto.EnumName(ETournamentTemplate_name, int32(x))
}
func (x *ETournamentTemplate) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentTemplate_value, data, "ETournamentTemplate")
	if err != nil {
		return err
	}
	*x = ETournamentTemplate(value)
	return nil
}
func (ETournamentTemplate) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type ETournamentGameState int32

const (
	ETournamentGameState_k_ETournamentGameState_Unknown              ETournamentGameState = 0
	ETournamentGameState_k_ETournamentGameState_Canceled             ETournamentGameState = 1
	ETournamentGameState_k_ETournamentGameState_Scheduled            ETournamentGameState = 2
	ETournamentGameState_k_ETournamentGameState_Active               ETournamentGameState = 3
	ETournamentGameState_k_ETournamentGameState_RadVictory           ETournamentGameState = 20
	ETournamentGameState_k_ETournamentGameState_DireVictory          ETournamentGameState = 21
	ETournamentGameState_k_ETournamentGameState_RadVictoryByForfeit  ETournamentGameState = 22
	ETournamentGameState_k_ETournamentGameState_DireVictoryByForfeit ETournamentGameState = 23
	ETournamentGameState_k_ETournamentGameState_ServerFailure        ETournamentGameState = 40
	ETournamentGameState_k_ETournamentGameState_NotNeeded            ETournamentGameState = 41
)

var ETournamentGameState_name = map[int32]string{
	0:  "k_ETournamentGameState_Unknown",
	1:  "k_ETournamentGameState_Canceled",
	2:  "k_ETournamentGameState_Scheduled",
	3:  "k_ETournamentGameState_Active",
	20: "k_ETournamentGameState_RadVictory",
	21: "k_ETournamentGameState_DireVictory",
	22: "k_ETournamentGameState_RadVictoryByForfeit",
	23: "k_ETournamentGameState_DireVictoryByForfeit",
	40: "k_ETournamentGameState_ServerFailure",
	41: "k_ETournamentGameState_NotNeeded",
}
var ETournamentGameState_value = map[string]int32{
	"k_ETournamentGameState_Unknown":              0,
	"k_ETournamentGameState_Canceled":             1,
	"k_ETournamentGameState_Scheduled":            2,
	"k_ETournamentGameState_Active":               3,
	"k_ETournamentGameState_RadVictory":           20,
	"k_ETournamentGameState_DireVictory":          21,
	"k_ETournamentGameState_RadVictoryByForfeit":  22,
	"k_ETournamentGameState_DireVictoryByForfeit": 23,
	"k_ETournamentGameState_ServerFailure":        40,
	"k_ETournamentGameState_NotNeeded":            41,
}

func (x ETournamentGameState) Enum() *ETournamentGameState {
	p := new(ETournamentGameState)
	*p = x
	return p
}
func (x ETournamentGameState) String() string {
	return proto.EnumName(ETournamentGameState_name, int32(x))
}
func (x *ETournamentGameState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentGameState_value, data, "ETournamentGameState")
	if err != nil {
		return err
	}
	*x = ETournamentGameState(value)
	return nil
}
func (ETournamentGameState) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

type ETournamentTeamState int32

const (
	ETournamentTeamState_k_ETournamentTeamState_Unknown      ETournamentTeamState = 0
	ETournamentTeamState_k_ETournamentTeamState_Node1        ETournamentTeamState = 1
	ETournamentTeamState_k_ETournamentTeamState_NodeMax      ETournamentTeamState = 1024
	ETournamentTeamState_k_ETournamentTeamState_Eliminated   ETournamentTeamState = 14003
	ETournamentTeamState_k_ETournamentTeamState_Forfeited    ETournamentTeamState = 14004
	ETournamentTeamState_k_ETournamentTeamState_Finished1st  ETournamentTeamState = 15001
	ETournamentTeamState_k_ETournamentTeamState_Finished2nd  ETournamentTeamState = 15002
	ETournamentTeamState_k_ETournamentTeamState_Finished3rd  ETournamentTeamState = 15003
	ETournamentTeamState_k_ETournamentTeamState_Finished4th  ETournamentTeamState = 15004
	ETournamentTeamState_k_ETournamentTeamState_Finished5th  ETournamentTeamState = 15005
	ETournamentTeamState_k_ETournamentTeamState_Finished6th  ETournamentTeamState = 15006
	ETournamentTeamState_k_ETournamentTeamState_Finished7th  ETournamentTeamState = 15007
	ETournamentTeamState_k_ETournamentTeamState_Finished8th  ETournamentTeamState = 15008
	ETournamentTeamState_k_ETournamentTeamState_Finished9th  ETournamentTeamState = 15009
	ETournamentTeamState_k_ETournamentTeamState_Finished10th ETournamentTeamState = 15010
	ETournamentTeamState_k_ETournamentTeamState_Finished11th ETournamentTeamState = 15011
	ETournamentTeamState_k_ETournamentTeamState_Finished12th ETournamentTeamState = 15012
	ETournamentTeamState_k_ETournamentTeamState_Finished13th ETournamentTeamState = 15013
	ETournamentTeamState_k_ETournamentTeamState_Finished14th ETournamentTeamState = 15014
	ETournamentTeamState_k_ETournamentTeamState_Finished15th ETournamentTeamState = 15015
	ETournamentTeamState_k_ETournamentTeamState_Finished16th ETournamentTeamState = 15016
)

var ETournamentTeamState_name = map[int32]string{
	0:     "k_ETournamentTeamState_Unknown",
	1:     "k_ETournamentTeamState_Node1",
	1024:  "k_ETournamentTeamState_NodeMax",
	14003: "k_ETournamentTeamState_Eliminated",
	14004: "k_ETournamentTeamState_Forfeited",
	15001: "k_ETournamentTeamState_Finished1st",
	15002: "k_ETournamentTeamState_Finished2nd",
	15003: "k_ETournamentTeamState_Finished3rd",
	15004: "k_ETournamentTeamState_Finished4th",
	15005: "k_ETournamentTeamState_Finished5th",
	15006: "k_ETournamentTeamState_Finished6th",
	15007: "k_ETournamentTeamState_Finished7th",
	15008: "k_ETournamentTeamState_Finished8th",
	15009: "k_ETournamentTeamState_Finished9th",
	15010: "k_ETournamentTeamState_Finished10th",
	15011: "k_ETournamentTeamState_Finished11th",
	15012: "k_ETournamentTeamState_Finished12th",
	15013: "k_ETournamentTeamState_Finished13th",
	15014: "k_ETournamentTeamState_Finished14th",
	15015: "k_ETournamentTeamState_Finished15th",
	15016: "k_ETournamentTeamState_Finished16th",
}
var ETournamentTeamState_value = map[string]int32{
	"k_ETournamentTeamState_Unknown":      0,
	"k_ETournamentTeamState_Node1":        1,
	"k_ETournamentTeamState_NodeMax":      1024,
	"k_ETournamentTeamState_Eliminated":   14003,
	"k_ETournamentTeamState_Forfeited":    14004,
	"k_ETournamentTeamState_Finished1st":  15001,
	"k_ETournamentTeamState_Finished2nd":  15002,
	"k_ETournamentTeamState_Finished3rd":  15003,
	"k_ETournamentTeamState_Finished4th":  15004,
	"k_ETournamentTeamState_Finished5th":  15005,
	"k_ETournamentTeamState_Finished6th":  15006,
	"k_ETournamentTeamState_Finished7th":  15007,
	"k_ETournamentTeamState_Finished8th":  15008,
	"k_ETournamentTeamState_Finished9th":  15009,
	"k_ETournamentTeamState_Finished10th": 15010,
	"k_ETournamentTeamState_Finished11th": 15011,
	"k_ETournamentTeamState_Finished12th": 15012,
	"k_ETournamentTeamState_Finished13th": 15013,
	"k_ETournamentTeamState_Finished14th": 15014,
	"k_ETournamentTeamState_Finished15th": 15015,
	"k_ETournamentTeamState_Finished16th": 15016,
}

func (x ETournamentTeamState) Enum() *ETournamentTeamState {
	p := new(ETournamentTeamState)
	*p = x
	return p
}
func (x ETournamentTeamState) String() string {
	return proto.EnumName(ETournamentTeamState_name, int32(x))
}
func (x *ETournamentTeamState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentTeamState_value, data, "ETournamentTeamState")
	if err != nil {
		return err
	}
	*x = ETournamentTeamState(value)
	return nil
}
func (ETournamentTeamState) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

type ETournamentState int32

const (
	ETournamentState_k_ETournamentState_Unknown            ETournamentState = 0
	ETournamentState_k_ETournamentState_CanceledByAdmin    ETournamentState = 1
	ETournamentState_k_ETournamentState_Completed          ETournamentState = 2
	ETournamentState_k_ETournamentState_Merged             ETournamentState = 3
	ETournamentState_k_ETournamentState_ServerFailure      ETournamentState = 4
	ETournamentState_k_ETournamentState_TeamAbandoned      ETournamentState = 5
	ETournamentState_k_ETournamentState_TeamTimeoutForfeit ETournamentState = 6
	ETournamentState_k_ETournamentState_TeamTimeoutRefund  ETournamentState = 7
	ETournamentState_k_ETournamentState_InProgress         ETournamentState = 100
	ETournamentState_k_ETournamentState_WaitingToMerge     ETournamentState = 101
)

var ETournamentState_name = map[int32]string{
	0:   "k_ETournamentState_Unknown",
	1:   "k_ETournamentState_CanceledByAdmin",
	2:   "k_ETournamentState_Completed",
	3:   "k_ETournamentState_Merged",
	4:   "k_ETournamentState_ServerFailure",
	5:   "k_ETournamentState_TeamAbandoned",
	6:   "k_ETournamentState_TeamTimeoutForfeit",
	7:   "k_ETournamentState_TeamTimeoutRefund",
	100: "k_ETournamentState_InProgress",
	101: "k_ETournamentState_WaitingToMerge",
}
var ETournamentState_value = map[string]int32{
	"k_ETournamentState_Unknown":            0,
	"k_ETournamentState_CanceledByAdmin":    1,
	"k_ETournamentState_Completed":          2,
	"k_ETournamentState_Merged":             3,
	"k_ETournamentState_ServerFailure":      4,
	"k_ETournamentState_TeamAbandoned":      5,
	"k_ETournamentState_TeamTimeoutForfeit": 6,
	"k_ETournamentState_TeamTimeoutRefund":  7,
	"k_ETournamentState_InProgress":         100,
	"k_ETournamentState_WaitingToMerge":     101,
}

func (x ETournamentState) Enum() *ETournamentState {
	p := new(ETournamentState)
	*p = x
	return p
}
func (x ETournamentState) String() string {
	return proto.EnumName(ETournamentState_name, int32(x))
}
func (x *ETournamentState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentState_value, data, "ETournamentState")
	if err != nil {
		return err
	}
	*x = ETournamentState(value)
	return nil
}
func (ETournamentState) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

type ETournamentNodeState int32

const (
	ETournamentNodeState_k_ETournamentNodeState_Unknown             ETournamentNodeState = 0
	ETournamentNodeState_k_ETournamentNodeState_Canceled            ETournamentNodeState = 1
	ETournamentNodeState_k_ETournamentNodeState_TeamsNotYetAssigned ETournamentNodeState = 2
	ETournamentNodeState_k_ETournamentNodeState_InBetweenGames      ETournamentNodeState = 3
	ETournamentNodeState_k_ETournamentNodeState_GameInProgress      ETournamentNodeState = 4
	ETournamentNodeState_k_ETournamentNodeState_A_Won               ETournamentNodeState = 5
	ETournamentNodeState_k_ETournamentNodeState_B_Won               ETournamentNodeState = 6
	ETournamentNodeState_k_ETournamentNodeState_A_WonByForfeit      ETournamentNodeState = 7
	ETournamentNodeState_k_ETournamentNodeState_B_WonByForfeit      ETournamentNodeState = 8
	ETournamentNodeState_k_ETournamentNodeState_A_Bye               ETournamentNodeState = 9
	ETournamentNodeState_k_ETournamentNodeState_A_Abandoned         ETournamentNodeState = 10
	ETournamentNodeState_k_ETournamentNodeState_ServerFailure       ETournamentNodeState = 11
	ETournamentNodeState_k_ETournamentNodeState_A_TimeoutForfeit    ETournamentNodeState = 12
	ETournamentNodeState_k_ETournamentNodeState_A_TimeoutRefund     ETournamentNodeState = 13
)

var ETournamentNodeState_name = map[int32]string{
	0:  "k_ETournamentNodeState_Unknown",
	1:  "k_ETournamentNodeState_Canceled",
	2:  "k_ETournamentNodeState_TeamsNotYetAssigned",
	3:  "k_ETournamentNodeState_InBetweenGames",
	4:  "k_ETournamentNodeState_GameInProgress",
	5:  "k_ETournamentNodeState_A_Won",
	6:  "k_ETournamentNodeState_B_Won",
	7:  "k_ETournamentNodeState_A_WonByForfeit",
	8:  "k_ETournamentNodeState_B_WonByForfeit",
	9:  "k_ETournamentNodeState_A_Bye",
	10: "k_ETournamentNodeState_A_Abandoned",
	11: "k_ETournamentNodeState_ServerFailure",
	12: "k_ETournamentNodeState_A_TimeoutForfeit",
	13: "k_ETournamentNodeState_A_TimeoutRefund",
}
var ETournamentNodeState_value = map[string]int32{
	"k_ETournamentNodeState_Unknown":             0,
	"k_ETournamentNodeState_Canceled":            1,
	"k_ETournamentNodeState_TeamsNotYetAssigned": 2,
	"k_ETournamentNodeState_InBetweenGames":      3,
	"k_ETournamentNodeState_GameInProgress":      4,
	"k_ETournamentNodeState_A_Won":               5,
	"k_ETournamentNodeState_B_Won":               6,
	"k_ETournamentNodeState_A_WonByForfeit":      7,
	"k_ETournamentNodeState_B_WonByForfeit":      8,
	"k_ETournamentNodeState_A_Bye":               9,
	"k_ETournamentNodeState_A_Abandoned":         10,
	"k_ETournamentNodeState_ServerFailure":       11,
	"k_ETournamentNodeState_A_TimeoutForfeit":    12,
	"k_ETournamentNodeState_A_TimeoutRefund":     13,
}

func (x ETournamentNodeState) Enum() *ETournamentNodeState {
	p := new(ETournamentNodeState)
	*p = x
	return p
}
func (x ETournamentNodeState) String() string {
	return proto.EnumName(ETournamentNodeState_name, int32(x))
}
func (x *ETournamentNodeState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETournamentNodeState_value, data, "ETournamentNodeState")
	if err != nil {
		return err
	}
	*x = ETournamentNodeState(value)
	return nil
}
func (ETournamentNodeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

type EDOTAGroupMergeResult int32

const (
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_OK                   EDOTAGroupMergeResult = 0
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_FAILED_GENERIC       EDOTAGroupMergeResult = 1
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_NOT_LEADER           EDOTAGroupMergeResult = 2
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_TOO_MANY_PLAYERS     EDOTAGroupMergeResult = 3
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_TOO_MANY_COACHES     EDOTAGroupMergeResult = 4
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_ENGINE_MISMATCH      EDOTAGroupMergeResult = 5
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_NO_SUCH_GROUP        EDOTAGroupMergeResult = 6
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_OTHER_GROUP_NOT_OPEN EDOTAGroupMergeResult = 7
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_ALREADY_INVITED      EDOTAGroupMergeResult = 8
	EDOTAGroupMergeResult_k_EDOTAGroupMergeResult_NOT_INVITED          EDOTAGroupMergeResult = 9
)

var EDOTAGroupMergeResult_name = map[int32]string{
	0: "k_EDOTAGroupMergeResult_OK",
	1: "k_EDOTAGroupMergeResult_FAILED_GENERIC",
	2: "k_EDOTAGroupMergeResult_NOT_LEADER",
	3: "k_EDOTAGroupMergeResult_TOO_MANY_PLAYERS",
	4: "k_EDOTAGroupMergeResult_TOO_MANY_COACHES",
	5: "k_EDOTAGroupMergeResult_ENGINE_MISMATCH",
	6: "k_EDOTAGroupMergeResult_NO_SUCH_GROUP",
	7: "k_EDOTAGroupMergeResult_OTHER_GROUP_NOT_OPEN",
	8: "k_EDOTAGroupMergeResult_ALREADY_INVITED",
	9: "k_EDOTAGroupMergeResult_NOT_INVITED",
}
var EDOTAGroupMergeResult_value = map[string]int32{
	"k_EDOTAGroupMergeResult_OK":                   0,
	"k_EDOTAGroupMergeResult_FAILED_GENERIC":       1,
	"k_EDOTAGroupMergeResult_NOT_LEADER":           2,
	"k_EDOTAGroupMergeResult_TOO_MANY_PLAYERS":     3,
	"k_EDOTAGroupMergeResult_TOO_MANY_COACHES":     4,
	"k_EDOTAGroupMergeResult_ENGINE_MISMATCH":      5,
	"k_EDOTAGroupMergeResult_NO_SUCH_GROUP":        6,
	"k_EDOTAGroupMergeResult_OTHER_GROUP_NOT_OPEN": 7,
	"k_EDOTAGroupMergeResult_ALREADY_INVITED":      8,
	"k_EDOTAGroupMergeResult_NOT_INVITED":          9,
}

func (x EDOTAGroupMergeResult) Enum() *EDOTAGroupMergeResult {
	p := new(EDOTAGroupMergeResult)
	*p = x
	return p
}
func (x EDOTAGroupMergeResult) String() string {
	return proto.EnumName(EDOTAGroupMergeResult_name, int32(x))
}
func (x *EDOTAGroupMergeResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDOTAGroupMergeResult_value, data, "EDOTAGroupMergeResult")
	if err != nil {
		return err
	}
	*x = EDOTAGroupMergeResult(value)
	return nil
}
func (EDOTAGroupMergeResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func init() {
	proto.RegisterEnum("dota.ETournamentTemplate", ETournamentTemplate_name, ETournamentTemplate_value)
	proto.RegisterEnum("dota.ETournamentGameState", ETournamentGameState_name, ETournamentGameState_value)
	proto.RegisterEnum("dota.ETournamentTeamState", ETournamentTeamState_name, ETournamentTeamState_value)
	proto.RegisterEnum("dota.ETournamentState", ETournamentState_name, ETournamentState_value)
	proto.RegisterEnum("dota.ETournamentNodeState", ETournamentNodeState_name, ETournamentNodeState_value)
	proto.RegisterEnum("dota.EDOTAGroupMergeResult", EDOTAGroupMergeResult_name, EDOTAGroupMergeResult_value)
}

func init() { proto.RegisterFile("dota_client_enums.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 884 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd6, 0x49, 0x6f, 0xdb, 0x46,
	0x14, 0x07, 0x70, 0x2b, 0x92, 0x97, 0x4c, 0x5b, 0xe0, 0x61, 0x9a, 0x34, 0x68, 0xd1, 0xa4, 0x49,
	0x1c, 0x2f, 0x51, 0x82, 0x34, 0x8e, 0xbb, 0x1e, 0x29, 0x69, 0x2c, 0x09, 0xb5, 0x28, 0x83, 0xa2,
	0x63, 0xf8, 0x52, 0x82, 0xd5, 0xbc, 0x98, 0x83, 0x88, 0x33, 0x06, 0x39, 0x4c, 0xea, 0x5b, 0x4e,
	0xfd, 0x0e, 0xdd, 0xf7, 0xd6, 0xf7, 0xf6, 0x3b, 0xf4, 0xd2, 0xef, 0xd1, 0xaf, 0x51, 0x50, 0x96,
	0x4d, 0x89, 0x8b, 0xcc, 0xf3, 0xfc, 0x34, 0xe4, 0xbc, 0xf7, 0xf8, 0x1f, 0x91, 0x1b, 0x5c, 0x69,
	0xd7, 0x19, 0x8e, 0x04, 0x4a, 0xed, 0xa0, 0x8c, 0xfc, 0xf0, 0xd1, 0x71, 0xa0, 0xb4, 0xa2, 0xb5,
	0x78, 0xa1, 0xfe, 0x39, 0x79, 0x93, 0xd9, 0x2a, 0x0a, 0xa4, 0xeb, 0xa3, 0xd4, 0x36, 0xfa, 0xc7,
	0x23, 0x57, 0x23, 0xbd, 0x45, 0xde, 0x79, 0xee, 0xe4, 0x2c, 0x38, 0xa6, 0x92, 0x08, 0x0b, 0x74,
	0x83, 0xac, 0xe6, 0xaf, 0x1b, 0x91, 0x56, 0xbe, 0xab, 0x91, 0x1f, 0x08, 0xb9, 0x0d, 0x95, 0xfa,
	0x69, 0x95, 0x5c, 0x9b, 0x72, 0x6d, 0xd7, 0xc7, 0x81, 0x8e, 0x9f, 0x70, 0x97, 0xdc, 0x9a, 0xd9,
	0xe1, 0x62, 0xc5, 0xd9, 0x97, 0xcf, 0xa5, 0x7a, 0x29, 0x61, 0x81, 0xae, 0x92, 0xf7, 0x0a, 0x4c,
	0xd3, 0x95, 0x43, 0x1c, 0x21, 0x87, 0x0a, 0xbd, 0x47, 0x6e, 0x17, 0xa0, 0xc1, 0xd0, 0x43, 0x1e,
	0xc5, 0xea, 0x0a, 0xbd, 0x43, 0x6e, 0x16, 0x28, 0x63, 0xa8, 0xc5, 0x0b, 0x84, 0x2a, 0x5d, 0x23,
	0x77, 0x0a, 0x88, 0xe5, 0xf2, 0xa7, 0x62, 0xa8, 0x55, 0x70, 0x02, 0xd7, 0xe8, 0x3a, 0xb9, 0x5b,
	0xc0, 0x5a, 0x22, 0xc0, 0x73, 0x77, 0x9d, 0x3e, 0x22, 0xf5, 0x4b, 0xb7, 0x6b, 0x9c, 0xec, 0xa8,
	0xe0, 0x19, 0x0a, 0x0d, 0x6f, 0xd1, 0xf7, 0xc9, 0x83, 0xcb, 0xf7, 0x4d, 0x7e, 0x70, 0x83, 0x6e,
	0x92, 0x7b, 0x45, 0x07, 0xc7, 0xe0, 0x05, 0x06, 0x3b, 0xae, 0x18, 0x45, 0x01, 0xc2, 0xe6, 0x9c,
	0x12, 0x99, 0x4a, 0x9b, 0x88, 0x1c, 0x39, 0xdc, 0xaf, 0xff, 0xb7, 0x34, 0xd3, 0x2a, 0x1b, 0x5d,
	0x3f, 0xbf, 0x55, 0x17, 0x2b, 0x53, 0xad, 0xba, 0x4d, 0xde, 0x2d, 0x30, 0xa6, 0xe2, 0xb8, 0x05,
	0x15, 0xba, 0x5a, 0xb8, 0x4b, 0x2c, 0x7a, 0xee, 0x97, 0xf0, 0x6a, 0x85, 0xae, 0xa7, 0x7a, 0x90,
	0x20, 0x36, 0x12, 0xbe, 0x90, 0xf1, 0x64, 0xc1, 0x5f, 0x3e, 0x5d, 0x4b, 0x9d, 0x28, 0x71, 0x93,
	0xfa, 0x20, 0x87, 0xbf, 0x7d, 0xba, 0x91, 0xea, 0xd5, 0x14, 0x13, 0x52, 0x84, 0x1e, 0xf2, 0xad,
	0x50, 0xc3, 0xd7, 0x51, 0x09, 0xf8, 0x44, 0x72, 0xf8, 0xa6, 0x0c, 0xdc, 0x0e, 0x38, 0x7c, 0x5b,
	0x06, 0x7e, 0xa0, 0x3d, 0xf8, 0xae, 0x0c, 0xfc, 0x50, 0x7b, 0xf0, 0x7d, 0x19, 0xf8, 0x91, 0xf6,
	0xe0, 0x87, 0x32, 0xf0, 0x63, 0xed, 0xc1, 0x8f, 0x65, 0xe0, 0x27, 0xda, 0x83, 0x9f, 0xca, 0xc0,
	0x4f, 0xb5, 0x07, 0x3f, 0x47, 0x74, 0x33, 0x93, 0x0b, 0x99, 0x82, 0x3f, 0xd6, 0x1e, 0xfc, 0x52,
	0x4a, 0x6e, 0x69, 0x0f, 0x7e, 0x2d, 0x25, 0x9f, 0x68, 0x0f, 0x7e, 0x2b, 0x25, 0xb7, 0xb5, 0x07,
	0xbf, 0x97, 0x92, 0x71, 0x7b, 0xfe, 0x28, 0x25, 0xe3, 0xfe, 0xfc, 0x59, 0x4a, 0xc6, 0x0d, 0x3a,
	0x8d, 0xea, 0x5f, 0x55, 0x09, 0x4c, 0xc1, 0xb3, 0xaf, 0x2c, 0x1d, 0xb9, 0xe9, 0x2f, 0x2c, 0x9d,
	0x3b, 0xb3, 0x41, 0xd8, 0x38, 0x31, 0xb8, 0x2f, 0x24, 0x54, 0x32, 0x5f, 0xe2, 0xc4, 0x29, 0xff,
	0x78, 0x84, 0x7a, 0x9c, 0x85, 0x37, 0xc9, 0xdb, 0x39, 0xa2, 0x87, 0xc1, 0x11, 0x72, 0xa8, 0x66,
	0xd2, 0x22, 0x2f, 0x53, 0x6a, 0x05, 0x2a, 0x3e, 0xb3, 0xf1, 0x85, 0x2b, 0xb9, 0x92, 0xc8, 0x61,
	0x91, 0xde, 0x27, 0x6b, 0x05, 0xca, 0x16, 0x3e, 0xaa, 0x48, 0x9f, 0xc7, 0xd9, 0x52, 0x26, 0xce,
	0x32, 0xd4, 0xc2, 0x67, 0x91, 0xe4, 0xb0, 0x9c, 0xc9, 0xf2, 0x33, 0xd9, 0x95, 0x7b, 0x81, 0x3a,
	0x0a, 0x30, 0x0c, 0x81, 0x67, 0xb2, 0xfc, 0x8c, 0x1c, 0xb8, 0x42, 0x0b, 0x79, 0x64, 0xab, 0xf1,
	0x59, 0x01, 0xeb, 0xff, 0xd6, 0x66, 0x22, 0x2f, 0x0e, 0xa2, 0xfc, 0xc8, 0xbb, 0x58, 0x99, 0x73,
	0x3b, 0x25, 0x66, 0xea, 0x76, 0x4a, 0xdf, 0x02, 0x09, 0x8a, 0x4f, 0x16, 0x9a, 0x4a, 0x1f, 0xa2,
	0x36, 0xc2, 0x50, 0x1c, 0xc9, 0x71, 0x6f, 0xd2, 0x05, 0x4b, 0x7c, 0x57, 0x36, 0x50, 0xbf, 0x44,
	0x94, 0x71, 0x7a, 0x87, 0x50, 0x9d, 0x43, 0x63, 0x31, 0x55, 0x8e, 0x5a, 0x66, 0x26, 0x12, 0x6a,
	0x38, 0x07, 0x4a, 0xc2, 0xe2, 0x1c, 0xd1, 0x18, 0x8b, 0xa5, 0x39, 0x8f, 0x1b, 0xef, 0x91, 0xdc,
	0x4c, 0xcb, 0x73, 0x68, 0x63, 0x96, 0xae, 0xcc, 0x7d, 0xb3, 0xc6, 0x09, 0xc2, 0xd5, 0xcc, 0xdc,
	0x4f, 0x8b, 0x64, 0xd4, 0x48, 0x66, 0x7e, 0x12, 0x37, 0x3b, 0xba, 0xaf, 0xd1, 0x07, 0x64, 0xa3,
	0x70, 0xc7, 0xd4, 0x58, 0xbe, 0x4e, 0xeb, 0x64, 0xfd, 0x32, 0x3c, 0x19, 0xcc, 0x37, 0xea, 0xff,
	0x54, 0xc9, 0x75, 0xd6, 0xea, 0xdb, 0x46, 0x3b, 0x50, 0xd1, 0xf1, 0x78, 0xc8, 0x2c, 0x0c, 0xa3,
	0x91, 0x9e, 0x7c, 0xdc, 0x79, 0x4b, 0x4e, 0xff, 0x33, 0x58, 0x98, 0x3c, 0x25, 0x77, 0x7d, 0xc7,
	0xe8, 0xee, 0xb2, 0x96, 0xd3, 0x66, 0x26, 0xb3, 0xba, 0x4d, 0xa8, 0x4c, 0x0a, 0x92, 0x6b, 0xcd,
	0xbe, 0xed, 0xec, 0x32, 0xa3, 0xc5, 0x2c, 0xb8, 0x42, 0x1f, 0x92, 0xcd, 0x22, 0x67, 0xf7, 0xfb,
	0x4e, 0xcf, 0x30, 0x0f, 0x9d, 0xbd, 0x5d, 0xe3, 0x90, 0x59, 0x03, 0xa8, 0x96, 0xd2, 0xcd, 0xbe,
	0xd1, 0xec, 0xb0, 0x01, 0xd4, 0x26, 0x25, 0xcc, 0xd5, 0xcc, 0x6c, 0x77, 0x4d, 0xe6, 0xf4, 0xba,
	0x83, 0x9e, 0x61, 0x37, 0x3b, 0x17, 0x21, 0x50, 0xf0, 0xc2, 0xce, 0x60, 0xbf, 0xd9, 0x71, 0xda,
	0x56, 0x7f, 0x7f, 0x0f, 0x96, 0xe8, 0x63, 0xf2, 0xb0, 0xb0, 0x4e, 0x76, 0x87, 0x59, 0x67, 0x70,
	0x7c, 0xce, 0xfe, 0x1e, 0x33, 0x61, 0x79, 0xde, 0x9b, 0x18, 0xbb, 0x16, 0x33, 0x5a, 0x87, 0x4e,
	0xd7, 0x7c, 0xda, 0xb5, 0x59, 0x0b, 0x56, 0x26, 0x7f, 0x5b, 0x0b, 0x4b, 0x77, 0x0e, 0xaf, 0x36,
	0x16, 0x3b, 0x95, 0x57, 0x95, 0x85, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53, 0xe1, 0x77, 0x3d,
	0x3d, 0x0b, 0x00, 0x00,
}
