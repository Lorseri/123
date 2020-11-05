package main

import (
	"encoding/json"
	"fmt"
	"milvus_go_test/utils"
	"testing"

	"github.com/milvus-io/milvus-sdk-go/milvus"
	"github.com/stretchr/testify/assert"
)
var autoId bool = true
var segmentRowLimit int = utils.DefaultSegmentRowLimit
var fieldFloatName string = utils.DefaultFieldFloatName
var fieldFloatVectorName string = utils.DefaultFieldFloatVectorName

// TODO issue: failed sometimes
func TestCreateCollection(t *testing.T) {
	client, name := Collection(false, milvus.VECTORFLOAT)
	value, status, _ := client.HasCollection(name)
	t.Log(value)
	t.Log(status)
	assert.Equal(t, status.Ok(), true)
	assert.Equal(t, value, true)
}

func TestCreateCollectionBinary(t *testing.T) {
	client, name := Collection(false, milvus.VECTORBINARY)
	value, status, _ := client.HasCollection(name)
	t.Log(value)
	t.Log(status)
	assert.Equal(t, status.Ok(), true)
	assert.Equal(t, value, true)
}

func TestCreateCollectionWithoutConnect(t *testing.T) {
	client := GenDisconnectClient()
	name := utils.RandString(8)
	mapping := GenCollectionParams(name, false, segmentRowLimit)
	assert.Panics(t, func() {
		client.CreateCollection(mapping)
	})
}

// TODO
func TestCreateCollectionWithInvalidName(t *testing.T) {
	client := GetClient()
	for _, name := range utils.GenInvalidStrs() {
		mapping := GenCollectionParams(name, autoId, segmentRowLimit)
		t.Log(mapping)
		status, _ := client.CreateCollection(mapping)
		t.Log(status)
		assert.False(t, status.Ok())
		isHas, _, _ := client.HasCollection(name)
		assert.False(t, isHas)
	}
}

func TestCreateCollectionInvalidDimension(t *testing.T) {
	client := GetClient()
	var dimension int = 0
	dimParams := map[string]interface{}{
		"dim": dimension,
	}
	extraParams, _ := json.Marshal(dimParams)
	fields := []milvus.Field{
		{
			fieldFloatName,
			milvus.FLOAT,
			"",
			"",
		},
		{
			fieldFloatVectorName,
			milvus.VECTORFLOAT,
			"",
			string(extraParams),
		},
	}
	name := utils.RandString(8)
	params := map[string]interface{}{
		"auto_id":           autoId,
		"segment_row_limit": segmentRowLimit,
	}
	paramsStr, _ := json.Marshal(params)
	mapping := milvus.Mapping{CollectionName: name, Fields: fields, ExtraParams: string(paramsStr)}
	status, _ := client.CreateCollection(mapping)
	assert.Equal(t, status.Ok(), false)
}

func TestShowCollections(t *testing.T) {
	client := GetClient()
	originCollections := make([]string, 10)
	for i := 0; i < 10; i++ {
		name := utils.RandString(8)
		mapping := GenCollectionParams(name, autoId, segmentRowLimit)
		status, _ := client.CreateCollection(mapping)
		assert.Equal(t, status.Ok(), true)
		originCollections[i] = name
	}
	listCollections, listStatus, _ := client.ListCollections()
	assert.True(t, listStatus.Ok(), true)
	for i := 0; i < len(originCollections); i++ {
		assert.Contains(t, listCollections, originCollections[i])
	}
}

func TestDropCollections(t *testing.T) {
	client, name := Collection(false, milvus.VECTORFLOAT)
	status, _ := client.DropCollection(name)
	assert.True(t, status.Ok())
	isHas, _, _ := client.HasCollection(name)
	assert.False(t, isHas)
	listCollections, _, _ := client.ListCollections()
	assert.NotContains(t, listCollections, name)
}

// #4131
func TestDropCollectionNotExisted(t *testing.T) {
	client := GetClient()
	name := utils.RandString(8)
	status, error := client.DropCollection(name)
	assert.False(t, status.Ok())
	fmt.Println(error)
	t.Log(error)
}

func TestDropCollectionWithoutConnect(t *testing.T) {
	client, name := Collection(false, milvus.VECTORFLOAT)
	isHas, _, _ := client.HasCollection(name)
	assert.True(t, isHas)
	client.Disconnect()
	assert.Panics(t, func() {
		client.DropCollection(name)
	})
}

func TestHasCollectionNotExisted(t *testing.T) {
	client := GetClient()
	name := utils.RandString(8)
	isHas, _, _ := client.HasCollection(name)
	assert.False(t, isHas)
}

func TestHasCollection(t *testing.T) {
	client, name := Collection(false, milvus.VECTORFLOAT)
	isHas, status, _ := client.HasCollection(name)
	assert.True(t, status.Ok())
	assert.True(t, isHas)
}

func TestDescribeCollectionNotExisted(t *testing.T) {
	client := GetClient()
	name := utils.RandString(8)
	mapping, status, error := client.GetCollectionInfo(name)
	assert.False(t, status.Ok())
	fmt.Println(mapping)
	fmt.Println(error)
	t.Log(error)
}

 // #4130
func TestDescribeCollection(t *testing.T) {
	client, name := Collection(false, milvus.VECTORFLOAT)
	mapping, status, _ := client.GetCollectionInfo(name)
	assert.True(t, status.Ok())
	assert.Equal(t, mapping.CollectionName, name)
	t.Log(mapping)
	//assert.Equal(t, utils.DefaultSegmentRowLimit, mapping.)
	for i :=0; i<len(mapping.Fields); i++ {
		if mapping.Fields[i].Type == milvus.VECTORFLOAT {
			var dat map[string]interface{}
			json.Unmarshal([]byte(mapping.Fields[i].ExtraParams), &dat)
			var dim map[string]interface{}
			json.Unmarshal([]byte(dat["params"].(string)), &dim)
			assert.Equal(t, utils.DefaultDimension, int(dim["dim"].(float64)))
		}
		var datMapping map[string]interface{}
		json.Unmarshal([]byte(mapping.ExtraParams), &datMapping)
		var params map[string]interface{}
		json.Unmarshal([]byte(datMapping["params"].(string)), &params)
		assert.Equal(t, utils.DefaultSegmentRowLimit, int(params["segment_row_limit"].(float64)))
	}
}
