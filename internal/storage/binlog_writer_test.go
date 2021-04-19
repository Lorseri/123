package storage

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
)

func TestBinlogWriterReader(t *testing.T) {
	binlogWriter, err := NewInsertBinlogWriter(schemapb.DataType_INT32)
	defer binlogWriter.Close()
	assert.Nil(t, err)
	eventWriter, err := binlogWriter.NextInsertEventWriter()
	assert.Nil(t, err)
	err = eventWriter.AddInt32ToPayload([]int32{1, 2, 3})
	assert.Nil(t, err)
	assert.Nil(t, nil, binlogWriter.GetBuffer())
	err = binlogWriter.Close()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, binlogWriter.GetEventNums())
	nums, err := binlogWriter.GetRowNums()
	assert.Nil(t, err)
	assert.EqualValues(t, 3, nums)
	err = eventWriter.AddInt32ToPayload([]int32{1, 2, 3})
	assert.NotNil(t, err)
	nums, err = binlogWriter.GetRowNums()
	assert.Nil(t, err)
	assert.EqualValues(t, 3, nums)

	buffer := binlogWriter.GetBuffer()
	fmt.Println("reader offset : " + strconv.Itoa(len(buffer)))

	binlogReader, err := NewBinlogReader(buffer)
	assert.Nil(t, err)
	eventReader, err := binlogReader.NextEventReader()
	assert.Nil(t, err)
	_, err = eventReader.GetInt8FromPayload()
	assert.NotNil(t, err)
	payload, err := eventReader.GetInt32FromPayload()
	assert.Nil(t, err)
	assert.EqualValues(t, 3, len(payload))
	assert.EqualValues(t, 1, payload[0])
	assert.EqualValues(t, 2, payload[1])
	assert.EqualValues(t, 3, payload[2])

	reader, err := binlogReader.NextEventReader()
	assert.Nil(t, err)
	fmt.Println("reader offset : " + strconv.Itoa(int(binlogReader.currentOffset)))
	assert.Nil(t, reader)
}
