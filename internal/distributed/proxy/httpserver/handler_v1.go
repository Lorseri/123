package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/milvus-io/milvus/pkg/util/merr"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/common"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/proxy"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

func (h *Handlers) checkDatabase(c *gin.Context, dbName string) bool {
	if dbName == DefaultDbName {
		return true
	}
	resp, err := h.proxy.ListDatabases(c, &milvuspb.ListDatabasesRequest{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
		return false
	} else if resp.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(resp.Status.ErrorCode)), HTTPReturnMessage: resp.Status.Reason})
		return false
	}
	for _, db := range resp.DbNames {
		if db == dbName {
			return true
		}
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrDatabaseNotfound), HTTPReturnMessage: merr.ErrDatabaseNotfound.Error()})
	return false
}

func (h *Handlers) describeCollection(c *gin.Context, dbName string, collectionName string, needAuth bool) (*milvuspb.DescribeCollectionResponse, error) {
	req := milvuspb.DescribeCollectionRequest{
		DbName:         dbName,
		CollectionName: collectionName,
	}
	if needAuth && proxy.Params.CommonCfg.AuthorizationEnabled {
		username, ok := c.Get(ContextUsername)
		if !ok {
			c.JSON(http.StatusProxyAuthRequired, gin.H{HTTPReturnCode: Code(merr.ErrNeedAuthenticate), HTTPReturnMessage: merr.ErrNeedAuthenticate.Error()})
			return nil, merr.ErrNeedAuthenticate
		}
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return nil, authErr
		}
	}
	response, err := h.proxy.DescribeCollection(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
		return nil, err
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
		return nil, errors.New(response.Status.Reason)
	}
	primaryField, ok := getPrimaryField(response.Schema)
	if ok && primaryField.AutoID && !response.Schema.AutoID {
		log.Warn("primary filed autoID VS schema autoID", zap.String("collectionName", collectionName), zap.Bool("primary Field", primaryField.AutoID), zap.Bool("schema", response.Schema.AutoID))
		response.Schema.AutoID = EnableAutoID
	}
	return response, nil
}

func (h *Handlers) hasCollection(c *gin.Context, dbName string, collectionName string) (bool, error) {
	req := milvuspb.HasCollectionRequest{
		DbName:         dbName,
		CollectionName: collectionName,
	}
	response, err := h.proxy.HasCollection(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
		return false, err
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
		return false, errors.New(response.Status.Reason)
	} else {
		return response.Value, nil
	}
}

func (h *Handlers) RegisterRoutesToV1(router gin.IRouter) {
	router.GET(VectorCollectionsPath, h.listCollections)
	router.POST(VectorCollectionsCreatePath, h.createCollection)
	router.GET(VectorCollectionsDescribePath, h.getCollectionDetails)
	router.POST(VectorCollectionsDropPath, h.dropCollection)
	router.POST(VectorQueryPath, h.query)
	router.POST(VectorGetPath, h.get)
	router.POST(VectorDeletePath, h.delete)
	router.POST(VectorInsertPath, h.insert)
	router.POST(VectorSearchPath, h.search)
}

func (h *Handlers) listCollections(c *gin.Context) {
	dbName := c.DefaultQuery(HTTPDbName, DefaultDbName)
	req := milvuspb.ShowCollectionsRequest{
		DbName: dbName,
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, dbName) {
		return
	}
	response, err := h.proxy.ShowCollections(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
	} else {
		var collections []string
		if response.CollectionNames != nil {
			collections = response.CollectionNames
		} else {
			collections = []string{}
		}
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: collections})
	}
}

func (h *Handlers) createCollection(c *gin.Context) {
	httpReq := CreateCollectionReq{
		DbName:       DefaultDbName,
		MetricType:   DefaultMetricType,
		PrimaryField: DefaultPrimaryFieldName,
		VectorField:  DefaultVectorFieldName,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of create collection is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" || httpReq.Dimension == 0 {
		log.Warn("high level restful api, create collection require parameters: [collectionName, dimension], but miss", zap.Any("request", httpReq))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters.Error()})
		return
	}
	schema, err := proto.Marshal(&schemapb.CollectionSchema{
		Name:        httpReq.CollectionName,
		Description: httpReq.Description,
		AutoID:      EnableAutoID,
		Fields: []*schemapb.FieldSchema{
			{
				FieldID:      common.StartOfUserFieldID,
				Name:         httpReq.PrimaryField,
				IsPrimaryKey: true,
				DataType:     schemapb.DataType_Int64,
				AutoID:       EnableAutoID,
			}, {
				FieldID:      common.StartOfUserFieldID + 1,
				Name:         httpReq.VectorField,
				IsPrimaryKey: false,
				DataType:     schemapb.DataType_FloatVector,
				TypeParams: []*commonpb.KeyValuePair{
					{
						Key:   Dim,
						Value: strconv.FormatInt(int64(httpReq.Dimension), 10),
					},
				},
				AutoID: DisableAutoID,
			},
		},
		EnableDynamicField: EnableDynamic,
	})
	if err != nil {
		log.Warn("high level restful api, marshal collection schema fail", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMarshalCollectionSchema), HTTPReturnMessage: merr.ErrMarshalCollectionSchema.Error()})
		return
	}
	req := milvuspb.CreateCollectionRequest{
		DbName:           httpReq.DbName,
		CollectionName:   httpReq.CollectionName,
		Schema:           schema,
		ShardsNum:        ShardNumDefault,
		ConsistencyLevel: commonpb.ConsistencyLevel_Bounded,
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	resp, err := h.proxy.CreateCollection(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
		return
	} else if resp.ErrorCode != commonpb.ErrorCode_Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(resp.ErrorCode)), HTTPReturnMessage: resp.Reason})
		return
	}

	resp, err = h.proxy.CreateIndex(c, &milvuspb.CreateIndexRequest{
		DbName:         httpReq.DbName,
		CollectionName: httpReq.CollectionName,
		FieldName:      httpReq.VectorField,
		ExtraParams:    []*commonpb.KeyValuePair{{Key: common.MetricTypeKey, Value: httpReq.MetricType}},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
		return
	} else if resp.ErrorCode != commonpb.ErrorCode_Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(resp.ErrorCode)), HTTPReturnMessage: resp.Reason})
		return
	}
	resp, err = h.proxy.LoadCollection(c, &milvuspb.LoadCollectionRequest{
		DbName:         httpReq.DbName,
		CollectionName: httpReq.CollectionName,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
		return
	} else if resp.ErrorCode != commonpb.ErrorCode_Success {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(resp.ErrorCode)), HTTPReturnMessage: resp.Reason})
		return
	}
	c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: gin.H{}})
}

func (h *Handlers) getCollectionDetails(c *gin.Context) {
	collectionName := c.Query(HTTPCollectionName)
	if collectionName == "" {
		log.Warn("high level restful api, desc collection require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters.Error()})
		return
	}
	dbName := c.DefaultQuery(HTTPDbName, DefaultDbName)
	if !h.checkDatabase(c, dbName) {
		return
	}
	coll, err := h.describeCollection(c, dbName, collectionName, true)
	if err != nil || coll == nil {
		return
	}
	stateResp, stateErr := h.proxy.GetLoadState(c, &milvuspb.GetLoadStateRequest{
		DbName:         dbName,
		CollectionName: collectionName,
	})
	collLoadState := ""
	if stateErr != nil {
		log.Warn("get collection load state fail", zap.String("collection", collectionName), zap.String("err", stateErr.Error()))
	} else if stateResp.Status.ErrorCode != commonpb.ErrorCode_Success {
		log.Warn("get collection load state fail", zap.String("collection", collectionName), zap.String("err", stateResp.Status.Reason))
	} else {
		collLoadState = stateResp.State.String()
	}
	vectorField := ""
	for _, field := range coll.Schema.Fields {
		if field.DataType == schemapb.DataType_BinaryVector || field.DataType == schemapb.DataType_FloatVector {
			vectorField = field.Name
			break
		}
	}
	indexResp, indexErr := h.proxy.DescribeIndex(c, &milvuspb.DescribeIndexRequest{
		DbName:         dbName,
		CollectionName: collectionName,
		FieldName:      vectorField,
	})
	var indexDesc []gin.H
	if indexErr != nil {
		indexDesc = []gin.H{}
		log.Warn("get indexes description fail", zap.String("collection", collectionName), zap.String("vectorField", vectorField), zap.String("err", indexErr.Error()))
	} else if indexResp.Status.ErrorCode != commonpb.ErrorCode_Success {
		indexDesc = []gin.H{}
		log.Warn("get indexes description fail", zap.String("collection", collectionName), zap.String("vectorField", vectorField), zap.String("err", indexResp.Status.Reason))
	} else {
		indexDesc = printIndexes(indexResp.IndexDescriptions)
	}
	c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: gin.H{
		HTTPCollectionName:    coll.CollectionName,
		HTTPReturnDescription: coll.Schema.Description,
		"fields":              printFields(coll.Schema.Fields),
		"indexes":             indexDesc,
		"load":                collLoadState,
		"shardsNum":           coll.ShardsNum,
		"enableDynamic":       coll.Schema.EnableDynamicField,
	}})
}

func (h *Handlers) dropCollection(c *gin.Context) {
	httpReq := DropCollectionReq{
		DbName: DefaultDbName,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of drop collection is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" {
		log.Warn("high level restful api, drop collection require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters.Error()})
		return
	}
	req := milvuspb.DropCollectionRequest{
		DbName:         httpReq.DbName,
		CollectionName: httpReq.CollectionName,
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	has, err := h.hasCollection(c, httpReq.DbName, httpReq.CollectionName)
	if err != nil {
		return
	}
	if !has {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrCollectionNotFound), HTTPReturnMessage: merr.ErrCollectionNotFound.Error()})
		return
	}
	response, err := h.proxy.DropCollection(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.ErrorCode)), HTTPReturnMessage: response.Reason})
	} else {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: gin.H{}})
	}
}

func (h *Handlers) query(c *gin.Context) {
	httpReq := QueryReq{
		DbName: DefaultDbName,
		Limit:  100,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of query is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" {
		log.Warn("high level restful api, query require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters.Error()})
		return
	}
	req := milvuspb.QueryRequest{
		DbName:             httpReq.DbName,
		CollectionName:     httpReq.CollectionName,
		Expr:               httpReq.Filter,
		OutputFields:       httpReq.OutputFields,
		GuaranteeTimestamp: BoundedTimestamp,
		QueryParams:        []*commonpb.KeyValuePair{},
	}
	if httpReq.Offset > 0 {
		req.QueryParams = append(req.QueryParams, &commonpb.KeyValuePair{Key: ParamOffset, Value: strconv.FormatInt(int64(httpReq.Offset), 10)})
	}
	if httpReq.Limit > 0 {
		req.QueryParams = append(req.QueryParams, &commonpb.KeyValuePair{Key: ParamLimit, Value: strconv.FormatInt(int64(httpReq.Limit), 10)})
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	if req.OutputFields == nil {
		req.OutputFields = []string{DefaultOutputFields}
		coll, err := h.describeCollection(c, httpReq.DbName, httpReq.CollectionName, false)
		if err != nil || coll == nil {
			return
		}
		for _, field := range coll.Schema.Fields {
			if field.DataType == schemapb.DataType_BinaryVector || field.DataType == schemapb.DataType_FloatVector {
				req.OutputFields = append(req.OutputFields, field.Name)
			}
		}
	}
	response, err := h.proxy.Query(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
	} else {
		outputData, err := buildQueryResp(int64(0), response.OutputFields, response.FieldsData, nil, nil)
		if err != nil {
			log.Warn("high level restful api, fail to deal with query result", zap.Any("response", response), zap.Error(err))
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrInvalidSearchResult), HTTPReturnMessage: merr.ErrInvalidSearchResult.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: outputData})
		}
	}
}

func (h *Handlers) get(c *gin.Context) {
	httpReq := GetReq{
		DbName: DefaultDbName,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of get is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" {
		log.Warn("high level restful api, get require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters})
		return
	}
	req := milvuspb.QueryRequest{
		DbName:             httpReq.DbName,
		CollectionName:     httpReq.CollectionName,
		OutputFields:       httpReq.OutputFields,
		GuaranteeTimestamp: BoundedTimestamp,
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	coll, err := h.describeCollection(c, httpReq.DbName, httpReq.CollectionName, false)
	if err != nil || coll == nil {
		return
	}
	body, _ := c.Get(gin.BodyBytesKey)
	filter, err := checkGetPrimaryKey(coll.Schema, gjson.Get(string(body.([]byte)), DefaultPrimaryFieldName))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrCheckPrimaryKey), HTTPReturnMessage: merr.ErrCheckPrimaryKey.Error()})
		return
	}
	req.Expr = filter
	if req.OutputFields == nil {
		req.OutputFields = []string{DefaultOutputFields}
		for _, field := range coll.Schema.Fields {
			if field.DataType == schemapb.DataType_BinaryVector || field.DataType == schemapb.DataType_FloatVector {
				req.OutputFields = append(req.OutputFields, field.Name)
			}
		}
	}
	response, err := h.proxy.Query(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
	} else {
		outputData, err := buildQueryResp(int64(0), response.OutputFields, response.FieldsData, nil, nil)
		if err != nil {
			log.Warn("high level restful api, fail to deal with get result", zap.Any("response", response), zap.Error(err))
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrInvalidSearchResult), HTTPReturnMessage: merr.ErrInvalidSearchResult.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: outputData})
			log.Error("get resultIS: ", zap.Any("res", outputData))
		}
	}
}

func (h *Handlers) delete(c *gin.Context) {
	httpReq := DeleteReq{
		DbName: DefaultDbName,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of delete is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" {
		log.Warn("high level restful api, delete require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters.Error()})
		return
	}
	req := milvuspb.DeleteRequest{
		DbName:         httpReq.DbName,
		CollectionName: httpReq.CollectionName,
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	coll, err := h.describeCollection(c, httpReq.DbName, httpReq.CollectionName, false)
	if err != nil || coll == nil {
		return
	}
	body, _ := c.Get(gin.BodyBytesKey)
	filter, err := checkGetPrimaryKey(coll.Schema, gjson.Get(string(body.([]byte)), DefaultPrimaryFieldName))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrCheckPrimaryKey), HTTPReturnMessage: merr.ErrCheckPrimaryKey.Error()})
		return
	}
	req.Expr = filter
	response, err := h.proxy.Delete(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
	} else {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: gin.H{}})
	}
}

func (h *Handlers) insert(c *gin.Context) {
	httpReq := InsertReq{
		DbName: DefaultDbName,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of insert is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" {
		log.Warn("high level restful api, insert require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters.Error()})
		return
	}
	req := milvuspb.InsertRequest{
		DbName:         httpReq.DbName,
		CollectionName: httpReq.CollectionName,
		PartitionName:  "_default",
		NumRows:        uint32(len(httpReq.Data)),
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	coll, err := h.describeCollection(c, httpReq.DbName, httpReq.CollectionName, false)
	if err != nil || coll == nil {
		return
	}
	body, _ := c.Get(gin.BodyBytesKey)
	err = checkAndSetData(string(body.([]byte)), coll, &httpReq)
	if err != nil {
		log.Warn("high level restful api, fail to deal with insert data", zap.Any("body", body), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrInvalidInsertData), HTTPReturnMessage: merr.ErrInvalidInsertData.Error()})
		return
	}
	req.FieldsData, err = anyToColumns(httpReq.Data, coll.Schema)
	if err != nil {
		log.Warn("high level restful api, fail to deal with insert data", zap.Any("data", httpReq.Data), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrInvalidInsertData), HTTPReturnMessage: merr.ErrInvalidInsertData.Error()})
		return
	}
	response, err := h.proxy.Insert(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
	} else {
		switch response.IDs.GetIdField().(type) {
		case *schemapb.IDs_IntId:
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: gin.H{"insertCount": response.InsertCnt, "insertIds": response.IDs.IdField.(*schemapb.IDs_IntId).IntId.Data}})
		case *schemapb.IDs_StrId:
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: gin.H{"insertCount": response.InsertCnt, "insertIds": response.IDs.IdField.(*schemapb.IDs_StrId).StrId.Data}})
		default:
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrCheckPrimaryKey), HTTPReturnMessage: merr.ErrCheckPrimaryKey.Error()})
		}
	}
}

func (h *Handlers) search(c *gin.Context) {
	httpReq := SearchReq{
		DbName: DefaultDbName,
		Limit:  100,
	}
	if err := c.ShouldBindBodyWith(&httpReq, binding.JSON); err != nil {
		log.Warn("high level restful api, the parameter of search is incorrect", zap.Any("request", httpReq), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrIncorrectParameterFormat), HTTPReturnMessage: merr.ErrIncorrectParameterFormat.Error()})
		return
	}
	if httpReq.CollectionName == "" {
		log.Warn("high level restful api, search require parameter: [collectionName], but miss")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrMissingRequiredParameters), HTTPReturnMessage: merr.ErrMissingRequiredParameters})
		return
	}

	params := map[string]interface{}{ //auto generated mapping
		"level": int(commonpb.ConsistencyLevel_Bounded),
	}
	bs, _ := json.Marshal(params)
	searchParams := []*commonpb.KeyValuePair{
		{Key: common.TopKKey, Value: strconv.FormatInt(int64(httpReq.Limit), 10)},
		{Key: Params, Value: string(bs)},
		{Key: ParamRoundDecimal, Value: "-1"},
		{Key: ParamOffset, Value: strconv.FormatInt(int64(httpReq.Offset), 10)},
	}
	req := milvuspb.SearchRequest{
		DbName:             httpReq.DbName,
		CollectionName:     httpReq.CollectionName,
		Dsl:                httpReq.Filter,
		PlaceholderGroup:   vector2PlaceholderGroupBytes(httpReq.Vector),
		DslType:            commonpb.DslType_BoolExprV1,
		OutputFields:       httpReq.OutputFields,
		SearchParams:       searchParams,
		GuaranteeTimestamp: BoundedTimestamp,
		Nq:                 int64(1),
	}
	if proxy.Params.CommonCfg.AuthorizationEnabled {
		username, _ := c.Get(ContextUsername)
		_, authErr := proxy.PrivilegeInterceptorWithUsername(c, username.(string), &req)
		if authErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{HTTPReturnCode: Code(authErr), HTTPReturnMessage: authErr.Error()})
			return
		}
	}
	if !h.checkDatabase(c, req.DbName) {
		return
	}
	response, err := h.proxy.Search(c, &req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(err), HTTPReturnMessage: err.Error()})
	} else if response.Status.ErrorCode != commonpb.ErrorCode_Success {
		c.JSON(http.StatusOK, gin.H{HTTPReturnCode: code(int32(response.Status.ErrorCode)), HTTPReturnMessage: response.Status.Reason})
	} else {
		if response.Results.TopK == int64(0) {
			c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: []interface{}{}})
		} else {
			outputData, err := buildQueryResp(response.Results.TopK, response.Results.OutputFields, response.Results.FieldsData, response.Results.Ids, response.Results.Scores)
			if err != nil {
				log.Warn("high level restful api, fail to deal with search result", zap.Any("result", response.Results), zap.Error(err))
				c.JSON(http.StatusOK, gin.H{HTTPReturnCode: Code(merr.ErrInvalidSearchResult), HTTPReturnMessage: merr.ErrInvalidSearchResult.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{HTTPReturnCode: http.StatusOK, HTTPReturnData: outputData})
			}
		}
	}
}

func code(code int32) int32 {
	return merr.RootReasonCodeMask & code
}
func Code(err error) int32 {
	return code(merr.Code(err))
}
