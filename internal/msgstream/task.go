package msgstream

import (
	"github.com/golang/protobuf/proto"
	internalPb "github.com/zilliztech/milvus-distributed/internal/proto/internalpb"
)

type MsgType = internalPb.MsgType

type TsMsg interface {
	BeginTs() Timestamp
	EndTs() Timestamp
	Type() MsgType
	HashKeys() []int32
	Marshal(*TsMsg) ([]byte, error)
	Unmarshal([]byte) (*TsMsg, error)
}

type BaseMsg struct {
	BeginTimestamp Timestamp
	EndTimestamp   Timestamp
	HashValues     []int32
}

func (bm *BaseMsg) BeginTs() Timestamp {
	return bm.BeginTimestamp
}

func (bm *BaseMsg) EndTs() Timestamp {
	return bm.EndTimestamp
}

func (bm *BaseMsg) HashKeys() []int32 {
	return bm.HashValues
}

/////////////////////////////////////////Insert//////////////////////////////////////////
type InsertMsg struct {
	BaseMsg
	internalPb.InsertRequest
}

func (it *InsertMsg) Type() MsgType {
	return it.MsgType
}

func (it *InsertMsg) Marshal(input *TsMsg) ([]byte, error) {
	insertMsg := (*input).(*InsertMsg)
	insertRequest := &insertMsg.InsertRequest
	mb, err := proto.Marshal(insertRequest)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (it *InsertMsg) Unmarshal(input []byte) (*TsMsg, error) {
	insertRequest := internalPb.InsertRequest{}
	err := proto.Unmarshal(input, &insertRequest)
	if err != nil {
		return nil, err
	}
	insertMsg := &InsertMsg{InsertRequest: insertRequest}
	for _, timestamp := range insertMsg.Timestamps {
		insertMsg.BeginTimestamp = timestamp
		insertMsg.EndTimestamp = timestamp
		break
	}
	for _, timestamp := range insertMsg.Timestamps {
		if timestamp > insertMsg.EndTimestamp {
			insertMsg.EndTimestamp = timestamp
		}
		if timestamp < insertMsg.BeginTimestamp {
			insertMsg.BeginTimestamp = timestamp
		}
	}

	var tsMsg TsMsg = insertMsg
	return &tsMsg, nil
}

/////////////////////////////////////////Delete//////////////////////////////////////////
type DeleteMsg struct {
	BaseMsg
	internalPb.DeleteRequest
}

func (dt *DeleteMsg) Type() MsgType {
	return dt.MsgType
}

func (dt *DeleteMsg) Marshal(input *TsMsg) ([]byte, error) {
	deleteTask := (*input).(*DeleteMsg)
	deleteRequest := &deleteTask.DeleteRequest
	mb, err := proto.Marshal(deleteRequest)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (dt *DeleteMsg) Unmarshal(input []byte) (*TsMsg, error) {
	deleteRequest := internalPb.DeleteRequest{}
	err := proto.Unmarshal(input, &deleteRequest)
	if err != nil {
		return nil, err
	}
	deleteMsg := &DeleteMsg{DeleteRequest: deleteRequest}
	for _, timestamp := range deleteMsg.Timestamps {
		deleteMsg.BeginTimestamp = timestamp
		deleteMsg.EndTimestamp = timestamp
		break
	}
	for _, timestamp := range deleteMsg.Timestamps {
		if timestamp > deleteMsg.EndTimestamp {
			deleteMsg.EndTimestamp = timestamp
		}
		if timestamp < deleteMsg.BeginTimestamp {
			deleteMsg.BeginTimestamp = timestamp
		}
	}

	var tsMsg TsMsg = deleteMsg
	return &tsMsg, nil
}

/////////////////////////////////////////Search//////////////////////////////////////////
type SearchMsg struct {
	BaseMsg
	internalPb.SearchRequest
}

func (st *SearchMsg) Type() MsgType {
	return st.MsgType
}

func (st *SearchMsg) Marshal(input *TsMsg) ([]byte, error) {
	searchTask := (*input).(*SearchMsg)
	searchRequest := &searchTask.SearchRequest
	mb, err := proto.Marshal(searchRequest)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (st *SearchMsg) Unmarshal(input []byte) (*TsMsg, error) {
	searchRequest := internalPb.SearchRequest{}
	err := proto.Unmarshal(input, &searchRequest)
	if err != nil {
		return nil, err
	}
	searchMsg := &SearchMsg{SearchRequest: searchRequest}
	searchMsg.BeginTimestamp = searchMsg.Timestamp
	searchMsg.EndTimestamp = searchMsg.Timestamp

	var tsMsg TsMsg = searchMsg
	return &tsMsg, nil
}

/////////////////////////////////////////SearchResult//////////////////////////////////////////
type SearchResultMsg struct {
	BaseMsg
	internalPb.SearchResult
}

func (srt *SearchResultMsg) Type() MsgType {
	return srt.MsgType
}

func (srt *SearchResultMsg) Marshal(input *TsMsg) ([]byte, error) {
	searchResultTask := (*input).(*SearchResultMsg)
	searchResultRequest := &searchResultTask.SearchResult
	mb, err := proto.Marshal(searchResultRequest)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (srt *SearchResultMsg) Unmarshal(input []byte) (*TsMsg, error) {
	searchResultRequest := internalPb.SearchResult{}
	err := proto.Unmarshal(input, &searchResultRequest)
	if err != nil {
		return nil, err
	}
	searchResultMsg := &SearchResultMsg{SearchResult: searchResultRequest}
	searchResultMsg.BeginTimestamp = searchResultMsg.Timestamp
	searchResultMsg.EndTimestamp = searchResultMsg.Timestamp

	var tsMsg TsMsg = searchResultMsg
	return &tsMsg, nil
}

/////////////////////////////////////////TimeTick//////////////////////////////////////////
type TimeTickMsg struct {
	BaseMsg
	internalPb.TimeTickMsg
}

func (tst *TimeTickMsg) Type() MsgType {
	return tst.MsgType
}

func (tst *TimeTickMsg) Marshal(input *TsMsg) ([]byte, error) {
	timeTickTask := (*input).(*TimeTickMsg)
	timeTick := &timeTickTask.TimeTickMsg
	mb, err := proto.Marshal(timeTick)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (tst *TimeTickMsg) Unmarshal(input []byte) (*TsMsg, error) {
	timeTickMsg := internalPb.TimeTickMsg{}
	err := proto.Unmarshal(input, &timeTickMsg)
	if err != nil {
		return nil, err
	}
	timeTick := &TimeTickMsg{TimeTickMsg: timeTickMsg}
	timeTick.BeginTimestamp = timeTick.Timestamp
	timeTick.EndTimestamp = timeTick.Timestamp

	var tsMsg TsMsg = timeTick
	return &tsMsg, nil
}

/////////////////////////////////////////QueryNodeSegStats//////////////////////////////////////////
type QueryNodeSegStatsMsg struct {
	BaseMsg
	internalPb.QueryNodeSegStats
}

func (qs *QueryNodeSegStatsMsg) Type() MsgType {
	return qs.MsgType
}

func (qs *QueryNodeSegStatsMsg) Marshal(input *TsMsg) ([]byte, error) {
	queryNodeSegStatsTask := (*input).(*QueryNodeSegStatsMsg)
	queryNodeSegStats := &queryNodeSegStatsTask.QueryNodeSegStats
	mb, err := proto.Marshal(queryNodeSegStats)
	if err != nil {
		return nil, err
	}
	return mb, nil
}

func (qs *QueryNodeSegStatsMsg) Unmarshal(input []byte) (*TsMsg, error) {
	queryNodeSegStats := internalPb.QueryNodeSegStats{}
	err := proto.Unmarshal(input, &queryNodeSegStats)
	if err != nil {
		return nil, err
	}
	queryNodeSegStatsMsg := &QueryNodeSegStatsMsg{QueryNodeSegStats: queryNodeSegStats}

	var tsMsg TsMsg = queryNodeSegStatsMsg
	return &tsMsg, nil
}

///////////////////////////////////////////Key2Seg//////////////////////////////////////////
//type Key2SegMsg struct {
//	BaseMsg
//	internalPb.Key2SegMsg
//}
//
//func (k2st *Key2SegMsg) Type() MsgType {
//	return
//}
