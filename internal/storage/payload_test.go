// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"math"
	"math/rand"
	"testing"

	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/pkg/util/testutils"
)

func TestPayload_ReaderAndWriter(t *testing.T) {
	t.Run("TestBool", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, false, false, false})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]bool{false, false, false, false})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 8, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Bool, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 8)
		bools, err := r.GetBoolFromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []bool{false, false, false, false, false, false, false, false}, bools)
		ibools, _, err := r.GetDataFromPayload()
		bools = ibools.([]bool)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []bool{false, false, false, false, false, false, false, false}, bools)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestInt8", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int8)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt8ToPayload([]int8{1, 2, 3})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]int8{4, 5, 6})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int8, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)

		int8s, err := r.GetInt8FromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int8{1, 2, 3, 4, 5, 6}, int8s)

		iint8s, _, err := r.GetDataFromPayload()
		int8s = iint8s.([]int8)
		assert.NoError(t, err)

		assert.ElementsMatch(t, []int8{1, 2, 3, 4, 5, 6}, int8s)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestInt16", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int16)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt16ToPayload([]int16{1, 2, 3})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]int16{1, 2, 3})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int16, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)
		int16s, err := r.GetInt16FromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int16{1, 2, 3, 1, 2, 3}, int16s)

		iint16s, _, err := r.GetDataFromPayload()
		int16s = iint16s.([]int16)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int16{1, 2, 3, 1, 2, 3}, int16s)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestInt32", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int32)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt32ToPayload([]int32{1, 2, 3})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]int32{1, 2, 3})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int32, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)

		int32s, err := r.GetInt32FromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int32{1, 2, 3, 1, 2, 3}, int32s)

		iint32s, _, err := r.GetDataFromPayload()
		int32s = iint32s.([]int32)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int32{1, 2, 3, 1, 2, 3}, int32s)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestInt64", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int64)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt64ToPayload([]int64{1, 2, 3})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]int64{1, 2, 3})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int64, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)

		int64s, err := r.GetInt64FromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int64{1, 2, 3, 1, 2, 3}, int64s)

		iint64s, _, err := r.GetDataFromPayload()
		int64s = iint64s.([]int64)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int64{1, 2, 3, 1, 2, 3}, int64s)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestFloat32", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Float)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddFloatToPayload([]float32{1.0, 2.0, 3.0})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]float32{1.0, 2.0, 3.0})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Float, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)

		float32s, err := r.GetFloatFromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []float32{1.0, 2.0, 3.0, 1.0, 2.0, 3.0}, float32s)

		ifloat32s, _, err := r.GetDataFromPayload()
		float32s = ifloat32s.([]float32)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []float32{1.0, 2.0, 3.0, 1.0, 2.0, 3.0}, float32s)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestDouble", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Double)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddDoubleToPayload([]float64{1.0, 2.0, 3.0})
		assert.NoError(t, err)
		err = w.AddDataToPayload([]float64{1.0, 2.0, 3.0})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Double, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)

		float64s, err := r.GetDoubleFromPayload()
		assert.NoError(t, err)
		assert.ElementsMatch(t, []float64{1.0, 2.0, 3.0, 1.0, 2.0, 3.0}, float64s)

		ifloat64s, _, err := r.GetDataFromPayload()
		float64s = ifloat64s.([]float64)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []float64{1.0, 2.0, 3.0, 1.0, 2.0, 3.0}, float64s)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestAddString", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_String)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddOneStringToPayload("hello0")
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("hello1")
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("hello2")
		assert.NoError(t, err)
		err = w.AddDataToPayload("hello3")
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)
		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_String, buffer)
		assert.NoError(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)

		str, err := r.GetStringFromPayload()
		assert.NoError(t, err)

		assert.Equal(t, str[0], "hello0")
		assert.Equal(t, str[1], "hello1")
		assert.Equal(t, str[2], "hello2")
		assert.Equal(t, str[3], "hello3")

		istr, _, err := r.GetDataFromPayload()
		strArray := istr.([]string)
		assert.NoError(t, err)
		assert.Equal(t, strArray[0], "hello0")
		assert.Equal(t, strArray[1], "hello1")
		assert.Equal(t, strArray[2], "hello2")
		assert.Equal(t, strArray[3], "hello3")
		r.ReleasePayloadReader()
		w.ReleasePayloadWriter()
	})

	t.Run("TestAddArray", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Array)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddOneArrayToPayload(&schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: []int32{1, 2},
				},
			},
		})
		assert.NoError(t, err)
		err = w.AddOneArrayToPayload(&schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: []int32{3, 4},
				},
			},
		})
		assert.NoError(t, err)
		err = w.AddOneArrayToPayload(&schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: []int32{5, 6},
				},
			},
		})
		assert.NoError(t, err)
		err = w.AddDataToPayload(&schemapb.ScalarField{
			Data: &schemapb.ScalarField_IntData{
				IntData: &schemapb.IntArray{
					Data: []int32{7, 8},
				},
			},
		})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)
		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Array, buffer)
		assert.NoError(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)

		arrayList, err := r.GetArrayFromPayload()
		assert.NoError(t, err)

		assert.EqualValues(t, []int32{1, 2}, arrayList[0].GetIntData().GetData())
		assert.EqualValues(t, []int32{3, 4}, arrayList[1].GetIntData().GetData())
		assert.EqualValues(t, []int32{5, 6}, arrayList[2].GetIntData().GetData())
		assert.EqualValues(t, []int32{7, 8}, arrayList[3].GetIntData().GetData())

		iArrayList, _, err := r.GetDataFromPayload()
		arrayList = iArrayList.([]*schemapb.ScalarField)
		assert.NoError(t, err)
		assert.EqualValues(t, []int32{1, 2}, arrayList[0].GetIntData().GetData())
		assert.EqualValues(t, []int32{3, 4}, arrayList[1].GetIntData().GetData())
		assert.EqualValues(t, []int32{5, 6}, arrayList[2].GetIntData().GetData())
		assert.EqualValues(t, []int32{7, 8}, arrayList[3].GetIntData().GetData())
		r.ReleasePayloadReader()
		w.ReleasePayloadWriter()
	})

	t.Run("TestAddJSON", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_JSON)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddOneJSONToPayload([]byte(`{"1":"1"}`))
		assert.NoError(t, err)
		err = w.AddOneJSONToPayload([]byte(`{"2":"2"}`))
		assert.NoError(t, err)
		err = w.AddOneJSONToPayload([]byte(`{"3":"3"}`))
		assert.NoError(t, err)
		err = w.AddDataToPayload([]byte(`{"4":"4"}`))
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)
		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_JSON, buffer)
		assert.NoError(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)

		json, err := r.GetJSONFromPayload()
		assert.NoError(t, err)

		assert.EqualValues(t, []byte(`{"1":"1"}`), json[0])
		assert.EqualValues(t, []byte(`{"2":"2"}`), json[1])
		assert.EqualValues(t, []byte(`{"3":"3"}`), json[2])
		assert.EqualValues(t, []byte(`{"4":"4"}`), json[3])

		iJSON, _, err := r.GetDataFromPayload()
		json = iJSON.([][]byte)
		assert.NoError(t, err)
		assert.EqualValues(t, []byte(`{"1":"1"}`), json[0])
		assert.EqualValues(t, []byte(`{"2":"2"}`), json[1])
		assert.EqualValues(t, []byte(`{"3":"3"}`), json[2])
		assert.EqualValues(t, []byte(`{"4":"4"}`), json[3])
		r.ReleasePayloadReader()
		w.ReleasePayloadWriter()
	})

	t.Run("TestBinaryVector", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_BinaryVector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)

		in := make([]byte, 16)
		for i := 0; i < 16; i++ {
			in[i] = 1
		}
		in2 := make([]byte, 8)
		for i := 0; i < 8; i++ {
			in2[i] = 1
		}

		err = w.AddBinaryVectorToPayload(in, 8)
		assert.NoError(t, err)
		err = w.AddDataToPayload(in2, 8)
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 24, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_BinaryVector, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 24)

		binVecs, dim, err := r.GetBinaryVectorFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, 8, dim)
		assert.Equal(t, 24, len(binVecs))

		ibinVecs, dim, err := r.GetDataFromPayload()
		assert.NoError(t, err)
		binVecs = ibinVecs.([]byte)
		assert.Equal(t, 8, dim)
		assert.Equal(t, 24, len(binVecs))
		defer r.ReleasePayloadReader()
	})

	t.Run("TestFloatVector", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_FloatVector, 1)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddFloatVectorToPayload([]float32{1.0, 2.0}, 1)
		assert.NoError(t, err)
		err = w.AddDataToPayload([]float32{3.0, 4.0}, 1)
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 4, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_FloatVector, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 4)

		floatVecs, dim, err := r.GetFloatVectorFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, 1, dim)
		assert.Equal(t, 4, len(floatVecs))
		assert.ElementsMatch(t, []float32{1.0, 2.0, 3.0, 4.0}, floatVecs)

		ifloatVecs, dim, err := r.GetDataFromPayload()
		assert.NoError(t, err)
		floatVecs = ifloatVecs.([]float32)
		assert.Equal(t, 1, dim)
		assert.Equal(t, 4, len(floatVecs))
		assert.ElementsMatch(t, []float32{1.0, 2.0, 3.0, 4.0}, floatVecs)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestFloat16Vector", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Float16Vector, 1)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddFloat16VectorToPayload([]byte{1, 2}, 1)
		assert.NoError(t, err)
		err = w.AddDataToPayload([]byte{3, 4}, 1)
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 2, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Float16Vector, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 2)

		float16Vecs, dim, err := r.GetFloat16VectorFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, 1, dim)
		assert.Equal(t, 4, len(float16Vecs))
		assert.ElementsMatch(t, []byte{1, 2, 3, 4}, float16Vecs)

		ifloat16Vecs, dim, err := r.GetDataFromPayload()
		assert.NoError(t, err)
		float16Vecs = ifloat16Vecs.([]byte)
		assert.Equal(t, 1, dim)
		assert.Equal(t, 4, len(float16Vecs))
		assert.ElementsMatch(t, []byte{1, 2, 3, 4}, float16Vecs)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestBFloat16Vector", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_BFloat16Vector, 1)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBFloat16VectorToPayload([]byte{1, 2}, 1)
		assert.NoError(t, err)
		err = w.AddDataToPayload([]byte{3, 4}, 1)
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 2, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_BFloat16Vector, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 2)

		bfloat16Vecs, dim, err := r.GetBFloat16VectorFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, 1, dim)
		assert.Equal(t, 4, len(bfloat16Vecs))
		assert.ElementsMatch(t, []byte{1, 2, 3, 4}, bfloat16Vecs)

		ibfloat16Vecs, dim, err := r.GetDataFromPayload()
		assert.NoError(t, err)
		bfloat16Vecs = ibfloat16Vecs.([]byte)
		assert.Equal(t, 1, dim)
		assert.Equal(t, 4, len(bfloat16Vecs))
		assert.ElementsMatch(t, []byte{1, 2, 3, 4}, bfloat16Vecs)
		defer r.ReleasePayloadReader()
	})

	t.Run("TestSparseFloatVector", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_SparseFloatVector)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddSparseFloatVectorToPayload(&SparseFloatVectorFieldData{
			SparseFloatArray: schemapb.SparseFloatArray{
				Dim: 600,
				Contents: [][]byte{
					testutils.CreateSparseFloatRow([]uint32{0, 1, 2}, []float32{1.1, 1.2, 1.3}),
					testutils.CreateSparseFloatRow([]uint32{10, 20, 30}, []float32{2.1, 2.2, 2.3}),
					testutils.CreateSparseFloatRow([]uint32{100, 200, 599}, []float32{3.1, 3.2, 3.3}),
				},
			},
		})
		assert.NoError(t, err)
		err = w.AddSparseFloatVectorToPayload(&SparseFloatVectorFieldData{
			SparseFloatArray: schemapb.SparseFloatArray{
				Dim: 600,
				Contents: [][]byte{
					testutils.CreateSparseFloatRow([]uint32{30, 41, 52}, []float32{1.1, 1.2, 1.3}),
					testutils.CreateSparseFloatRow([]uint32{60, 80, 230}, []float32{2.1, 2.2, 2.3}),
					testutils.CreateSparseFloatRow([]uint32{170, 300, 579}, []float32{3.1, 3.2, 3.3}),
				},
			},
		})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 6, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_SparseFloatVector, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 6)

		floatVecs, dim, err := r.GetSparseFloatVectorFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, 600, dim)
		assert.Equal(t, 6, len(floatVecs.Contents))
		assert.Equal(t, schemapb.SparseFloatArray{
			// merged dim should be max of all dims
			Dim: 600,
			Contents: [][]byte{
				testutils.CreateSparseFloatRow([]uint32{0, 1, 2}, []float32{1.1, 1.2, 1.3}),
				testutils.CreateSparseFloatRow([]uint32{10, 20, 30}, []float32{2.1, 2.2, 2.3}),
				testutils.CreateSparseFloatRow([]uint32{100, 200, 599}, []float32{3.1, 3.2, 3.3}),
				testutils.CreateSparseFloatRow([]uint32{30, 41, 52}, []float32{1.1, 1.2, 1.3}),
				testutils.CreateSparseFloatRow([]uint32{60, 80, 230}, []float32{2.1, 2.2, 2.3}),
				testutils.CreateSparseFloatRow([]uint32{170, 300, 579}, []float32{3.1, 3.2, 3.3}),
			},
		}, floatVecs.SparseFloatArray)

		ifloatVecs, dim, err := r.GetDataFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, floatVecs, ifloatVecs.(*SparseFloatVectorFieldData))
		assert.Equal(t, 600, dim)
		defer r.ReleasePayloadReader()
	})

	testSparse_OneBatch := func(t *testing.T, rows [][]byte, actualDim int) {
		w, err := NewPayloadWriter(schemapb.DataType_SparseFloatVector)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddSparseFloatVectorToPayload(&SparseFloatVectorFieldData{
			SparseFloatArray: schemapb.SparseFloatArray{
				Dim:      int64(actualDim),
				Contents: rows,
			},
		})
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 3, length)
		defer w.ReleasePayloadWriter()

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_SparseFloatVector, buffer)
		require.Nil(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, length, 3)

		floatVecs, dim, err := r.GetSparseFloatVectorFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, actualDim, dim)
		assert.Equal(t, 3, len(floatVecs.Contents))
		assert.Equal(t, schemapb.SparseFloatArray{
			Dim:      int64(dim),
			Contents: rows,
		}, floatVecs.SparseFloatArray)

		ifloatVecs, dim, err := r.GetDataFromPayload()
		assert.NoError(t, err)
		assert.Equal(t, floatVecs, ifloatVecs.(*SparseFloatVectorFieldData))
		assert.Equal(t, actualDim, dim)
		defer r.ReleasePayloadReader()
	}

	t.Run("TestSparseFloatVector_emptyRow", func(t *testing.T) {
		testSparse_OneBatch(t, [][]byte{
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
			testutils.CreateSparseFloatRow([]uint32{10, 20, 30}, []float32{2.1, 2.2, 2.3}),
			testutils.CreateSparseFloatRow([]uint32{100, 200, 599}, []float32{3.1, 3.2, 3.3}),
		}, 600)
		testSparse_OneBatch(t, [][]byte{
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
		}, 0)
	})

	t.Run("TestSparseFloatVector_largeRow", func(t *testing.T) {
		nnz := 100000
		// generate an int slice with nnz random sorted elements
		indices := make([]uint32, nnz)
		values := make([]float32, nnz)
		for i := 0; i < nnz; i++ {
			indices[i] = uint32(i * 6)
			values[i] = float32(i)
		}
		dim := int(indices[nnz-1]) + 1
		testSparse_OneBatch(t, [][]byte{
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
			testutils.CreateSparseFloatRow([]uint32{10, 20, 30}, []float32{2.1, 2.2, 2.3}),
			testutils.CreateSparseFloatRow(indices, values),
		}, dim)
	})

	t.Run("TestSparseFloatVector_negativeValues", func(t *testing.T) {
		testSparse_OneBatch(t, [][]byte{
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
			testutils.CreateSparseFloatRow([]uint32{10, 20, 30}, []float32{-2.1, 2.2, -2.3}),
			testutils.CreateSparseFloatRow([]uint32{100, 200, 599}, []float32{3.1, -3.2, 3.3}),
		}, 600)
	})

	// even though SPARSE_INVERTED_INDEX and SPARSE_WAND index do not support
	// arbitrarily large dimensions, HNSW does, so we still need to test it.
	// Dimension range we support is 0 to positive int32 max - 1(to leave room
	// for dim).
	t.Run("TestSparseFloatVector_largeIndex", func(t *testing.T) {
		int32Max := uint32(math.MaxInt32)
		testSparse_OneBatch(t, [][]byte{
			testutils.CreateSparseFloatRow([]uint32{}, []float32{}),
			testutils.CreateSparseFloatRow([]uint32{10, 20, 30}, []float32{-2.1, 2.2, -2.3}),
			testutils.CreateSparseFloatRow([]uint32{100, int32Max / 2, int32Max - 1}, []float32{3.1, -3.2, 3.3}),
		}, int(int32Max))
	})

	// t.Run("TestAddDataToPayload", func(t *testing.T) {
	// 	w, err := NewPayloadWriter(schemapb.DataType_Bool)
	// 	w.colType = 999
	// 	require.Nil(t, err)
	// 	require.NotNil(t, w)

	// 	err = w.AddDataToPayload([]bool{false, false, false, false})
	// 	assert.NotNil(t, err)

	// 	err = w.AddDataToPayload([]bool{false, false, false, false}, 0)
	// 	assert.NotNil(t, err)

	// 	err = w.AddDataToPayload([]bool{false, false, false, false}, 0, 0)
	// 	assert.NotNil(t, err)

	// 	err = w.AddBoolToPayload([]bool{})
	// 	assert.NotNil(t, err)
	// 	err = w.FinishPayloadWriter()
	// 	assert.Nil(t, err)
	// 	err = w.AddBoolToPayload([]bool{false})
	// 	assert.NotNil(t, err)
	// })

	t.Run("TestAddBoolAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddBoolToPayload([]bool{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddBoolToPayload([]bool{false})
		assert.Error(t, err)
	})

	t.Run("TestAddInt8AfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int8)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddInt8ToPayload([]int8{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddInt8ToPayload([]int8{0})
		assert.Error(t, err)
	})
	t.Run("TestAddInt16AfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int16)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddInt16ToPayload([]int16{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddInt16ToPayload([]int16{0})
		assert.Error(t, err)
	})
	t.Run("TestAddInt32AfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int32)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddInt32ToPayload([]int32{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddInt32ToPayload([]int32{0})
		assert.Error(t, err)
	})
	t.Run("TestAddInt64AfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int64)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddInt64ToPayload([]int64{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddInt64ToPayload([]int64{0})
		assert.Error(t, err)
	})
	t.Run("TestAddFloatAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Float)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddFloatToPayload([]float32{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddFloatToPayload([]float32{0.0})
		assert.Error(t, err)
	})
	t.Run("TestAddDoubleAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Double)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddDoubleToPayload([]float64{})
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddDoubleToPayload([]float64{0.0})
		assert.Error(t, err)
	})
	t.Run("TestAddOneStringAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_String)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.AddOneStringToPayload("")
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("c")
		assert.Error(t, err)
	})
	t.Run("TestAddBinVectorAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_BinaryVector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		_, err = w.GetPayloadBufferFromWriter()
		assert.Error(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		err = w.AddBinaryVectorToPayload([]byte{}, 8)
		assert.Error(t, err)
		err = w.AddBinaryVectorToPayload([]byte{1}, 0)
		assert.Error(t, err)

		err = w.AddBinaryVectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.Error(t, err)
		err = w.AddBinaryVectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.Error(t, err)
	})
	t.Run("TestAddFloatVectorAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_FloatVector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		err = w.AddFloatVectorToPayload([]float32{}, 8)
		assert.Error(t, err)
		err = w.AddFloatVectorToPayload([]float32{1.0}, 0)
		assert.Error(t, err)

		err = w.AddFloatVectorToPayload([]float32{1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}, 8)
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.Error(t, err)
		err = w.AddFloatVectorToPayload([]float32{1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}, 8)
		assert.Error(t, err)
	})
	t.Run("TestAddFloat16VectorAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Float16Vector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		err = w.AddFloat16VectorToPayload([]byte{}, 8)
		assert.Error(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		err = w.AddFloat16VectorToPayload([]byte{}, 8)
		assert.Error(t, err)
		err = w.AddFloat16VectorToPayload([]byte{1}, 0)
		assert.Error(t, err)

		err = w.AddFloat16VectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.Error(t, err)
		err = w.AddFloat16VectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.Error(t, err)
	})
	t.Run("TestAddBFloat16VectorAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_BFloat16Vector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		err = w.AddBFloat16VectorToPayload([]byte{}, 8)
		assert.Error(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		err = w.AddBFloat16VectorToPayload([]byte{}, 8)
		assert.Error(t, err)
		err = w.AddBFloat16VectorToPayload([]byte{1}, 0)
		assert.Error(t, err)

		err = w.AddBFloat16VectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.Error(t, err)
		err = w.FinishPayloadWriter()
		assert.Error(t, err)
		err = w.AddBFloat16VectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.Error(t, err)
	})
	t.Run("TestAddSparseFloatVectorAfterFinish", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_SparseFloatVector)
		require.Nil(t, err)
		require.NotNil(t, w)
		defer w.Close()

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		err = w.AddSparseFloatVectorToPayload(&SparseFloatVectorFieldData{
			SparseFloatArray: schemapb.SparseFloatArray{
				Dim: 53,
				Contents: [][]byte{
					testutils.CreateSparseFloatRow([]uint32{30, 41, 52}, []float32{1.1, 1.2, 1.3}),
				},
			},
		})
		assert.Error(t, err)
		err = w.AddSparseFloatVectorToPayload(&SparseFloatVectorFieldData{
			SparseFloatArray: schemapb.SparseFloatArray{
				Dim: 600,
				Contents: [][]byte{
					testutils.CreateSparseFloatRow([]uint32{30, 41, 52}, []float32{1.1, 1.2, 1.3}),
				},
			},
		})
		assert.Error(t, err)

		err = w.FinishPayloadWriter()
		assert.Error(t, err)
	})
	t.Run("TestNewReadError", func(t *testing.T) {
		buffer := []byte{0}
		r, err := NewPayloadReader(999, buffer)
		assert.Error(t, err)
		assert.Nil(t, r)
	})
	t.Run("TestGetDataError", func(t *testing.T) {
		r := PayloadReader{}
		r.colType = 999

		_, _, err := r.GetDataFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetBoolError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int8)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt8ToPayload([]int8{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Bool, buffer)
		assert.NoError(t, err)

		_, err = r.GetBoolFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetBoolFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetBoolError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{true, false, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Bool, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetBoolFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt8Error", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int8, buffer)
		assert.NoError(t, err)

		_, err = r.GetInt8FromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetInt8FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt8Error2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int8)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt8ToPayload([]int8{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int8, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetInt8FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt16Error", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int16, buffer)
		assert.NoError(t, err)

		_, err = r.GetInt16FromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetInt16FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt16Error2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int16)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt16ToPayload([]int16{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int16, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetInt16FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt32Error", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int32, buffer)
		assert.NoError(t, err)

		_, err = r.GetInt32FromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetInt32FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt32Error2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int32)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt32ToPayload([]int32{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int32, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetInt32FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt64Error", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int64, buffer)
		assert.NoError(t, err)

		_, err = r.GetInt64FromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetInt64FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetInt64Error2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Int64)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddInt64ToPayload([]int64{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Int64, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetInt64FromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetFloatError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Float, buffer)
		assert.NoError(t, err)

		_, err = r.GetFloatFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetFloatFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetFloatError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Float)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddFloatToPayload([]float32{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Float, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetFloatFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetDoubleError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Double, buffer)
		assert.NoError(t, err)

		_, err = r.GetDoubleFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetDoubleFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetDoubleError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Double)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddDoubleToPayload([]float64{1, 2, 3})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Double, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetDoubleFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetStringError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_String, buffer)
		assert.NoError(t, err)

		_, err = r.GetStringFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetStringFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetStringError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_String)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddOneStringToPayload("hello0")
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("hello1")
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("hello2")
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_String, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, err = r.GetStringFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetArrayError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_Array, buffer)
		assert.NoError(t, err)

		_, err = r.GetArrayFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, err = r.GetArrayFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetBinaryVectorError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_BinaryVector, buffer)
		assert.NoError(t, err)

		_, _, err = r.GetBinaryVectorFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, _, err = r.GetBinaryVectorFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetBinaryVectorError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_BinaryVector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBinaryVectorToPayload([]byte{1, 0, 0, 0, 0, 0, 0, 0}, 8)
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_BinaryVector, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, _, err = r.GetBinaryVectorFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetFloatVectorError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_FloatVector, buffer)
		assert.NoError(t, err)

		_, _, err = r.GetFloatVectorFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, _, err = r.GetFloatVectorFromPayload()
		assert.Error(t, err)
	})
	t.Run("TestGetFloatVectorError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_FloatVector, 8)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddFloatVectorToPayload([]float32{1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}, 8)
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_FloatVector, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, _, err = r.GetFloatVectorFromPayload()
		assert.Error(t, err)
	})

	t.Run("TestByteArrayDatasetError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_String)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddOneStringToPayload("hello0")
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_FloatVector, buffer)
		assert.NoError(t, err)

		r.colType = 99
		_, err = r.GetByteArrayDataSet()
		assert.Error(t, err)

		r.colType = schemapb.DataType_String
		dataset, err := r.GetByteArrayDataSet()
		assert.NoError(t, err)

		dataset.columnIdx = math.MaxInt
		_, err = dataset.NextBatch(100)
		assert.Error(t, err)

		dataset.groupID = math.MaxInt
		assert.Error(t, err)
	})

	t.Run("TestGetSparseFloatVectorError", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_Bool)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddBoolToPayload([]bool{false, true, true})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_SparseFloatVector, buffer)
		assert.NoError(t, err)

		_, _, err = r.GetSparseFloatVectorFromPayload()
		assert.Error(t, err)

		r.colType = 999
		_, _, err = r.GetSparseFloatVectorFromPayload()
		assert.Error(t, err)
	})

	t.Run("TestGetSparseFloatVectorError2", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_SparseFloatVector)
		require.Nil(t, err)
		require.NotNil(t, w)

		err = w.AddSparseFloatVectorToPayload(&SparseFloatVectorFieldData{
			SparseFloatArray: schemapb.SparseFloatArray{
				Dim: 53,
				Contents: [][]byte{
					testutils.CreateSparseFloatRow([]uint32{30, 41, 52}, []float32{1.1, 1.2, 1.3}),
				},
			},
		})
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_SparseFloatVector, buffer)
		assert.NoError(t, err)

		r.numRows = 99
		_, _, err = r.GetSparseFloatVectorFromPayload()
		assert.Error(t, err)
	})

	t.Run("TestWriteLargeSizeData", func(t *testing.T) {
		t.Skip("Large data skip for online ut")
		size := 1 << 29 // 512M
		var vec []float32
		for i := 0; i < size/4; i++ {
			vec = append(vec, 1)
		}

		w, err := NewPayloadWriter(schemapb.DataType_FloatVector)
		assert.NoError(t, err)

		err = w.AddFloatVectorToPayload(vec, 128)
		assert.NoError(t, err)

		err = w.FinishPayloadWriter()
		assert.NoError(t, err)

		_, err = w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		w.ReleasePayloadWriter()
	})
}

func TestArrowRecordReader(t *testing.T) {
	t.Run("TestArrowRecordReader", func(t *testing.T) {
		w, err := NewPayloadWriter(schemapb.DataType_String)
		assert.NoError(t, err)
		defer w.Close()

		err = w.AddOneStringToPayload("hello0")
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("hello1")
		assert.NoError(t, err)
		err = w.AddOneStringToPayload("hello2")
		assert.NoError(t, err)
		err = w.FinishPayloadWriter()
		assert.NoError(t, err)
		length, err := w.GetPayloadLengthFromWriter()
		assert.NoError(t, err)
		assert.Equal(t, 3, length)
		buffer, err := w.GetPayloadBufferFromWriter()
		assert.NoError(t, err)

		r, err := NewPayloadReader(schemapb.DataType_String, buffer)
		assert.NoError(t, err)
		length, err = r.GetPayloadLengthFromReader()
		assert.NoError(t, err)
		assert.Equal(t, 3, length)

		rr, err := r.GetArrowRecordReader()
		assert.NoError(t, err)

		for rr.Next() {
			rec := rr.Record()
			arr := rec.Column(0).(*array.String)
			defer rec.Release()

			assert.Equal(t, "hello0", arr.Value(0))
			assert.Equal(t, "hello1", arr.Value(1))
			assert.Equal(t, "hello2", arr.Value(2))
		}
	})
}

func dataGen(size int) ([]byte, error) {
	w, err := NewPayloadWriter(schemapb.DataType_String)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i := 0; i < size; i++ {
		b := make([]rune, 20)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		w.AddOneStringToPayload(string(b))
	}
	err = w.FinishPayloadWriter()
	if err != nil {
		return nil, err
	}
	buffer, err := w.GetPayloadBufferFromWriter()
	if err != nil {
		return nil, err
	}
	return buffer, err
}

func BenchmarkDefaultReader(b *testing.B) {
	size := 1000000
	buffer, err := dataGen(size)
	assert.NoError(b, err)

	b.ResetTimer()
	r, err := NewPayloadReader(schemapb.DataType_String, buffer)
	require.Nil(b, err)
	defer r.ReleasePayloadReader()

	length, err := r.GetPayloadLengthFromReader()
	assert.NoError(b, err)
	assert.Equal(b, length, size)

	d, err := r.GetStringFromPayload()
	assert.NoError(b, err)
	for i := 0; i < 100; i++ {
		for _, de := range d {
			assert.Equal(b, 20, len(de))
		}
	}
}

func BenchmarkDataSetReader(b *testing.B) {
	size := 1000000
	buffer, err := dataGen(size)
	assert.NoError(b, err)

	b.ResetTimer()
	r, err := NewPayloadReader(schemapb.DataType_String, buffer)
	require.Nil(b, err)
	defer r.ReleasePayloadReader()

	length, err := r.GetPayloadLengthFromReader()
	assert.NoError(b, err)
	assert.Equal(b, length, size)

	ds, err := r.GetByteArrayDataSet()
	assert.NoError(b, err)

	for i := 0; i < 100; i++ {
		for ds.HasNext() {
			stringArray, err := ds.NextBatch(1024)
			assert.NoError(b, err)
			for _, de := range stringArray {
				assert.Equal(b, 20, len(string(de)))
			}
		}
	}
}

func BenchmarkArrowRecordReader(b *testing.B) {
	size := 1000000
	buffer, err := dataGen(size)
	assert.NoError(b, err)

	b.ResetTimer()
	r, err := NewPayloadReader(schemapb.DataType_String, buffer)
	require.Nil(b, err)
	defer r.ReleasePayloadReader()

	length, err := r.GetPayloadLengthFromReader()
	assert.NoError(b, err)
	assert.Equal(b, length, size)

	rr, err := r.GetArrowRecordReader()
	assert.NoError(b, err)
	defer rr.Release()

	for i := 0; i < 100; i++ {
		for rr.Next() {
			rec := rr.Record()
			arr := rec.Column(0).(*array.String)
			defer rec.Release()
			for i := 0; i < arr.Len(); i++ {
				assert.Equal(b, 20, len(arr.Value(i)))
			}
		}
	}
}
