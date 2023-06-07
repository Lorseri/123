package indexparamcheck

import (
	"strconv"
	"testing"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/stretchr/testify/assert"
)

func Test_rHnswPQChecker_CheckTrain(t *testing.T) {
	validParams := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         L2,
	}

	invalidMatricParams := copyParams(validParams)
	invalidMatricParams[Metric] = JACCARD

	invalidEfParamsMin := copyParams(validParams)
	invalidEfParamsMin[EFConstruction] = strconv.Itoa(HNSWMinEfConstruction - 1)

	invalidEfParamsMax := copyParams(validParams)
	invalidEfParamsMax[EFConstruction] = strconv.Itoa(HNSWMaxEfConstruction + 1)

	invalidMParamsMin := copyParams(validParams)
	invalidMParamsMin[HNSWM] = strconv.Itoa(HNSWMinM - 1)

	invalidMParamsMax := copyParams(validParams)
	invalidMParamsMax[HNSWM] = strconv.Itoa(HNSWMaxM + 1)

	invalidParamsWithoutPQM := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		EFConstruction: strconv.Itoa(200),
		Metric:         L2,
	}

	invalidParamsPQM := copyParams(validParams)
	invalidParamsPQM[PQM] = "NAN"

	invalidParamsPQMZero := copyParams(validParams)
	invalidParamsPQMZero[PQM] = "0"

	p1 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         L2,
	}
	p2 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         IP,
	}

	//p3 := map[string]string{
	//	DIM:            strconv.Itoa(128),
	//	HNSWM:          strconv.Itoa(16),
	//	PQM:            strconv.Itoa(8),
	//	EFConstruction: strconv.Itoa(200),
	//	Metric:         COSINE,
	//}

	p4 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         HAMMING,
	}
	p5 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         JACCARD,
	}
	p6 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         TANIMOTO,
	}
	p7 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         SUBSTRUCTURE,
	}
	p8 := map[string]string{
		DIM:            strconv.Itoa(128),
		HNSWM:          strconv.Itoa(16),
		PQM:            strconv.Itoa(8),
		EFConstruction: strconv.Itoa(200),
		Metric:         SUPERSTRUCTURE,
	}

	cases := []struct {
		params   map[string]string
		errIsNil bool
	}{
		{validParams, true},
		{invalidMatricParams, false},
		{invalidEfParamsMin, false},
		{invalidEfParamsMax, false},
		{invalidMParamsMin, false},
		{invalidMParamsMax, false},
		{invalidParamsWithoutPQM, false},
		{invalidParamsPQM, false},
		{invalidParamsPQMZero, false},
		{p1, true},
		{p2, true},
		//{p3, true},
		{p4, false},
		{p5, false},
		{p6, false},
		{p7, false},
		{p8, false},
	}

	c := newRHnswPQChecker()
	for _, test := range cases {
		err := c.CheckTrain(test.params)
		if test.errIsNil {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func Test_rHnswPQChecker_CheckValidDataType(t *testing.T) {

	cases := []struct {
		dType    schemapb.DataType
		errIsNil bool
	}{
		{
			dType:    schemapb.DataType_Bool,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_Int8,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_Int16,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_Int32,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_Int64,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_Float,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_Double,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_String,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_VarChar,
			errIsNil: false,
		},
		{
			dType:    schemapb.DataType_FloatVector,
			errIsNil: true,
		},
		{
			dType:    schemapb.DataType_BinaryVector,
			errIsNil: false,
		},
	}

	c := newRHnswPQChecker()
	for _, test := range cases {
		err := c.CheckValidDataType(test.dType)
		if test.errIsNil {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
