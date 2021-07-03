// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package proxy

import (
	"fmt"
	"math"

	ant_ast "github.com/antonmedv/expr/ast"
	ant_parser "github.com/antonmedv/expr/parser"
	"github.com/milvus-io/milvus/internal/proto/planpb"
	"github.com/milvus-io/milvus/internal/proto/schemapb"
	"github.com/milvus-io/milvus/internal/util/typeutil"
)

func parseQueryExpr(schema *typeutil.SchemaHelper, exprStr string) (*planpb.Expr, error) {
	if exprStr == "" {
		return nil, nil
	}

	return parseQueryExprAdvanced(schema, exprStr)
}

type ParserContext struct {
	schema *typeutil.SchemaHelper
}

type optimizer struct {
	applied bool
	err     error
}

func (*optimizer) Enter(*ant_ast.Node) {}

func (optimizer *optimizer) Exit(node *ant_ast.Node) {
	patch := func(newNode ant_ast.Node) {
		optimizer.applied = true
		ant_ast.Patch(node, newNode)
	}

	value := func(node *ant_ast.Node) float64 {
		switch node := (*node).(type) {
		case *ant_ast.FloatNode:
			return node.Value
		case *ant_ast.IntegerNode:
			return float64(node.Value)
		}
		panic("unreachable")
	}

	switch node := (*node).(type) {
	case *ant_ast.UnaryNode:
		switch node.Operator {
		case "-":
			if i, ok := node.Node.(*ant_ast.IntegerNode); ok {
				patch(&ant_ast.IntegerNode{Value: -i.Value})
			}
			if i, ok := node.Node.(*ant_ast.FloatNode); ok {
				patch(&ant_ast.FloatNode{Value: -i.Value})
			}
		case "+":
			if i, ok := node.Node.(*ant_ast.IntegerNode); ok {
				patch(&ant_ast.IntegerNode{Value: i.Value})
			}
			if i, ok := node.Node.(*ant_ast.FloatNode); ok {
				patch(&ant_ast.FloatNode{Value: i.Value})
			}
		}

	case *ant_ast.BinaryNode:
		switch node.Operator {
		case "+":
			_, leftFloat := node.Left.(*ant_ast.FloatNode)
			_, leftInteger := node.Left.(*ant_ast.IntegerNode)
			_, rightFloat := node.Right.(*ant_ast.FloatNode)
			_, rightInteger := node.Right.(*ant_ast.IntegerNode)
			if (leftFloat || leftInteger) && (rightFloat || rightInteger) {
				if leftFloat || rightFloat {
					patch(&ant_ast.FloatNode{Value: value(&node.Left) + value(&node.Right)})
				} else {
					patch(&ant_ast.IntegerNode{Value: int(value(&node.Left) + value(&node.Right))})
				}
			}
		case "-":
			_, leftFloat := node.Left.(*ant_ast.FloatNode)
			_, leftInteger := node.Left.(*ant_ast.IntegerNode)
			_, rightFloat := node.Right.(*ant_ast.FloatNode)
			_, rightInteger := node.Right.(*ant_ast.IntegerNode)
			if (leftFloat || leftInteger) && (rightFloat || rightInteger) {
				if leftFloat || rightFloat {
					patch(&ant_ast.FloatNode{Value: value(&node.Left) - value(&node.Right)})
				} else {
					patch(&ant_ast.IntegerNode{Value: int(value(&node.Left) - value(&node.Right))})
				}
			}
		case "*":
			_, leftFloat := node.Left.(*ant_ast.FloatNode)
			_, leftInteger := node.Left.(*ant_ast.IntegerNode)
			_, rightFloat := node.Right.(*ant_ast.FloatNode)
			_, rightInteger := node.Right.(*ant_ast.IntegerNode)
			if (leftFloat || leftInteger) && (rightFloat || rightInteger) {
				if leftFloat || rightFloat {
					patch(&ant_ast.FloatNode{Value: value(&node.Left) * value(&node.Right)})
				} else {
					patch(&ant_ast.IntegerNode{Value: int(value(&node.Left) * value(&node.Right))})
				}
			}
		case "/":
			_, leftFloat := node.Left.(*ant_ast.FloatNode)
			_, leftInteger := node.Left.(*ant_ast.IntegerNode)
			_, rightFloat := node.Right.(*ant_ast.FloatNode)
			_, rightInteger := node.Right.(*ant_ast.IntegerNode)
			if (leftFloat || leftInteger) && (rightFloat || rightInteger) {
				if value(&node.Right) == 0 {
					optimizer.err = fmt.Errorf("number divide by zero")
					return
				}
				if leftFloat || rightFloat {
					patch(&ant_ast.FloatNode{Value: value(&node.Left) / value(&node.Right)})
				} else {
					patch(&ant_ast.IntegerNode{Value: int(value(&node.Left) / value(&node.Right))})
				}
			}
		case "%":
			left, leftInteger := node.Left.(*ant_ast.IntegerNode)
			right, rightInteger := node.Right.(*ant_ast.IntegerNode)
			if leftInteger && rightInteger {
				patch(&ant_ast.IntegerNode{Value: left.Value % right.Value})
			}
		case "**":
			_, leftFloat := node.Left.(*ant_ast.FloatNode)
			_, leftInteger := node.Left.(*ant_ast.IntegerNode)
			right, rightInteger := node.Right.(*ant_ast.IntegerNode)
			if (leftFloat || leftInteger) && rightInteger {
				patch(&ant_ast.FloatNode{Value: math.Pow(value(&node.Left), float64(right.Value))})
			}
		}
	}
}

func parseQueryExprAdvanced(schema *typeutil.SchemaHelper, exprStr string) (*planpb.Expr, error) {
	ast, err := ant_parser.Parse(exprStr)
	if err != nil {
		return nil, err
	}

	optimizer := &optimizer{}
	ant_ast.Walk(&ast.Node, optimizer)
	if optimizer.err != nil {
		return nil, optimizer.err
	}

	context := ParserContext{schema}

	expr, err := context.handleExpr(&ast.Node)
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func (context *ParserContext) createColumnInfo(field *schemapb.FieldSchema) *planpb.ColumnInfo {
	return &planpb.ColumnInfo{
		FieldId:      field.FieldID,
		DataType:     field.DataType,
		IsPrimaryKey: field.IsPrimaryKey,
	}
}

func isSameOrder(a, b string) bool {
	isLessA := a == "<" || a == "<="
	isLessB := b == "<" || b == "<="
	return isLessA == isLessB
}

func getCompareOpType(opStr string, reverse bool) planpb.RangeExpr_OpType {
	type OpType = planpb.RangeExpr_OpType
	var op planpb.RangeExpr_OpType

	if !reverse {
		switch opStr {
		case "<":
			op = planpb.RangeExpr_LessThan
		case ">":
			op = planpb.RangeExpr_GreaterThan
		case "<=":
			op = planpb.RangeExpr_LessEqual
		case ">=":
			op = planpb.RangeExpr_GreaterEqual
		case "==":
			op = planpb.RangeExpr_Equal
		case "!=":
			op = planpb.RangeExpr_NotEqual
		default:
			op = planpb.RangeExpr_Invalid
		}
	} else {
		switch opStr {
		case ">":
			op = planpb.RangeExpr_LessThan
		case "<":
			op = planpb.RangeExpr_GreaterThan
		case ">=":
			op = planpb.RangeExpr_LessEqual
		case "<=":
			op = planpb.RangeExpr_GreaterEqual
		case "==":
			op = planpb.RangeExpr_Equal
		case "!=":
			op = planpb.RangeExpr_NotEqual
		default:
			op = planpb.RangeExpr_Invalid
		}
	}
	return op
}

func getLogicalOpType(opStr string) planpb.BinaryExpr_BinaryOp {
	switch opStr {
	case "&&", "and":
		return planpb.BinaryExpr_LogicalAnd
	case "||", "or":
		return planpb.BinaryExpr_LogicalOr
	default:
		return planpb.BinaryExpr_Invalid
	}
}

func (context *ParserContext) createCmpExpr(left, right ant_ast.Node, operator string) (*planpb.Expr, error) {
	var idNode *ant_ast.IdentifierNode
	var isReversed bool
	var valueNode *ant_ast.Node
	if idNodeLeft, leftOk := left.(*ant_ast.IdentifierNode); leftOk {
		idNode = idNodeLeft
		isReversed = false
		valueNode = &right
	} else if idNodeRight, rightOk := right.(*ant_ast.IdentifierNode); rightOk {
		idNode = idNodeRight
		isReversed = true
		valueNode = &left
	} else {
		return nil, fmt.Errorf("compare expr has no identifier")
	}

	field, err := context.handleIdentifier(idNode)
	if err != nil {
		return nil, err
	}

	val, err := context.handleLeafValue(valueNode, field.DataType)
	if err != nil {
		return nil, err
	}

	op := getCompareOpType(operator, isReversed)
	if op == planpb.RangeExpr_Invalid {
		return nil, fmt.Errorf("invalid binary operator %s", operator)
	}

	expr := &planpb.Expr{
		Expr: &planpb.Expr_RangeExpr{
			RangeExpr: &planpb.RangeExpr{
				ColumnInfo: context.createColumnInfo(field),
				Ops:        []planpb.RangeExpr_OpType{op},
				Values:     []*planpb.GenericValue{val},
			},
		},
	}
	return expr, nil
}

func (context *ParserContext) handleCmpExpr(node *ant_ast.BinaryNode) (*planpb.Expr, error) {
	return context.createCmpExpr(node.Left, node.Right, node.Operator)
}

func (context *ParserContext) handleLogicalExpr(node *ant_ast.BinaryNode) (*planpb.Expr, error) {
	op := getLogicalOpType(node.Operator)
	if op == planpb.BinaryExpr_Invalid {
		return nil, fmt.Errorf("invalid logical op(%s)", node.Operator)
	}

	leftExpr, err := context.handleExpr(&node.Left)
	if err != nil {
		return nil, err
	}

	rightExpr, err := context.handleExpr(&node.Right)
	if err != nil {
		return nil, err
	}

	expr := &planpb.Expr{
		Expr: &planpb.Expr_BinaryExpr{
			BinaryExpr: &planpb.BinaryExpr{
				Op:    op,
				Left:  leftExpr,
				Right: rightExpr,
			},
		},
	}
	return expr, nil
}

func (context *ParserContext) handleArrayExpr(node *ant_ast.Node, dataType schemapb.DataType) ([]*planpb.GenericValue, error) {
	arrayNode, ok2 := (*node).(*ant_ast.ArrayNode)
	if !ok2 {
		return nil, fmt.Errorf("right operand of the InExpr must be array")
	}
	var arr []*planpb.GenericValue
	for _, element := range arrayNode.Nodes {
		val, err := context.handleLeafValue(&element, dataType)
		if err != nil {
			return nil, err
		}
		arr = append(arr, val)
	}
	return arr, nil
}

func (context *ParserContext) handleInExpr(node *ant_ast.BinaryNode) (*planpb.Expr, error) {
	if node.Operator != "in" && node.Operator != "not in" {
		return nil, fmt.Errorf("invalid Operator(%s)", node.Operator)
	}
	idNode, ok := node.Left.(*ant_ast.IdentifierNode)
	if !ok {
		return nil, fmt.Errorf("left operand of the InExpr must be identifier")
	}
	field, err := context.handleIdentifier(idNode)
	if err != nil {
		return nil, err
	}
	arrayData, err := context.handleArrayExpr(&node.Right, field.DataType)
	if err != nil {
		return nil, err
	}

	expr := &planpb.Expr{
		Expr: &planpb.Expr_TermExpr{
			TermExpr: &planpb.TermExpr{
				ColumnInfo: context.createColumnInfo(field),
				Values:     arrayData,
			},
		},
	}

	if node.Operator == "not in" {
		return context.createNotExpr(expr)
	}
	return expr, nil
}

func (context *ParserContext) handleBinaryExpr(node *ant_ast.BinaryNode) (*planpb.Expr, error) {
	switch node.Operator {
	case "<", "<=", ">", ">=":
		exprs := []*planpb.Expr{}
		curNode := node
		binNodeLeft, LeftOk := curNode.Left.(*ant_ast.BinaryNode)
		for LeftOk {
			if isSameOrder(node.Operator, binNodeLeft.Operator) {
				expr, err := context.createCmpExpr(binNodeLeft.Right, curNode.Right, curNode.Operator)
				if err != nil {
					return nil, err
				}
				exprs = append(exprs, expr)
				curNode = binNodeLeft
			}
			binNodeLeft, LeftOk = curNode.Left.(*ant_ast.BinaryNode)
		}
		combinedExpr, err := context.handleCmpExpr(curNode)
		if err != nil {
			return nil, err
		}

		for i := len(exprs) - 1; i >= 0; i-- {
			expr := exprs[i]
			combinedExpr = &planpb.Expr{
				Expr: &planpb.Expr_BinaryExpr{
					BinaryExpr: &planpb.BinaryExpr{
						Op:    planpb.BinaryExpr_LogicalAnd,
						Left:  combinedExpr,
						Right: expr,
					},
				},
			}
		}
		return combinedExpr, nil
	case "==", "!=":
		return context.handleCmpExpr(node)
	case "and", "or", "&&", "||":
		return context.handleLogicalExpr(node)
	case "in", "not in":
		return context.handleInExpr(node)
	}
	return nil, fmt.Errorf("unsupported binary operator %s", node.Operator)
}

func (context *ParserContext) createNotExpr(childExpr *planpb.Expr) (*planpb.Expr, error) {
	expr := &planpb.Expr{
		Expr: &planpb.Expr_UnaryExpr{
			UnaryExpr: &planpb.UnaryExpr{
				Op:    planpb.UnaryExpr_Not,
				Child: childExpr,
			},
		},
	}
	return expr, nil
}

func (context *ParserContext) handleLeafValue(nodeRaw *ant_ast.Node, dataType schemapb.DataType) (gv *planpb.GenericValue, err error) {
	switch node := (*nodeRaw).(type) {
	case *ant_ast.FloatNode:
		if typeutil.IsFloatingType(dataType) {
			gv = &planpb.GenericValue{
				Val: &planpb.GenericValue_FloatVal{
					FloatVal: node.Value,
				},
			}
		} else {
			return nil, fmt.Errorf("type mismatch")
		}
	case *ant_ast.IntegerNode:
		if typeutil.IsFloatingType(dataType) {
			gv = &planpb.GenericValue{
				Val: &planpb.GenericValue_FloatVal{
					FloatVal: float64(node.Value),
				},
			}
		} else if typeutil.IsIntergerType(dataType) {
			gv = &planpb.GenericValue{
				Val: &planpb.GenericValue_Int64Val{
					Int64Val: int64(node.Value),
				},
			}
		} else {
			return nil, fmt.Errorf("type mismatch")
		}
	case *ant_ast.BoolNode:
		if typeutil.IsFloatingType(dataType) {
			gv = &planpb.GenericValue{
				Val: &planpb.GenericValue_BoolVal{
					BoolVal: node.Value,
				},
			}
		} else {
			return nil, fmt.Errorf("type mismatch")
		}
	default:
		return nil, fmt.Errorf("unsupported leaf node")
	}

	return gv, nil
}

func (context *ParserContext) handleIdentifier(node *ant_ast.IdentifierNode) (*schemapb.FieldSchema, error) {
	fieldName := node.Value
	field, err := context.schema.GetFieldFromName(fieldName)
	return field, err
}

func (context *ParserContext) handleUnaryExpr(node *ant_ast.UnaryNode) (*planpb.Expr, error) {
	switch node.Operator {
	case "!", "not":
		subExpr, err := context.handleExpr(&node.Node)
		if err != nil {
			return nil, err
		}
		return context.createNotExpr(subExpr)
	default:
		return nil, fmt.Errorf("invalid unary operator(%s)", node.Operator)
	}
}

func (context *ParserContext) handleExpr(nodeRaw *ant_ast.Node) (*planpb.Expr, error) {
	switch node := (*nodeRaw).(type) {
	case *ant_ast.IdentifierNode,
		*ant_ast.FloatNode,
		*ant_ast.IntegerNode,
		*ant_ast.BoolNode:
		return nil, fmt.Errorf("scalar expr is not supported yet")
	case *ant_ast.UnaryNode:
		expr, err := context.handleUnaryExpr(node)
		if err != nil {
			return nil, err
		}
		return expr, nil
	case *ant_ast.BinaryNode:
		return context.handleBinaryExpr(node)
	default:
		return nil, fmt.Errorf("unsupported node (%s)", node.Type().String())
	}
}

func CreateQueryPlan(schemaPb *schemapb.CollectionSchema, exprStr string, vectorFieldName string, queryInfo *planpb.QueryInfo) (*planpb.PlanNode, error) {
	schema, err := typeutil.CreateSchemaHelper(schemaPb)
	if err != nil {
		return nil, err
	}

	expr, err := parseQueryExpr(schema, exprStr)
	if err != nil {
		return nil, err
	}
	vectorField, err := schema.GetFieldFromName(vectorFieldName)
	if err != nil {
		return nil, err
	}
	fieldID := vectorField.FieldID
	dataType := vectorField.DataType

	if !typeutil.IsVectorType(dataType) {
		return nil, fmt.Errorf("field (%s) to search is not of vector data type", vectorFieldName)
	}

	planNode := &planpb.PlanNode{
		Node: &planpb.PlanNode_VectorAnns{
			VectorAnns: &planpb.VectorANNS{
				IsBinary:       dataType == schemapb.DataType_BinaryVector,
				Predicates:     expr,
				QueryInfo:      queryInfo,
				PlaceholderTag: "$0",
				FieldId:        fieldID,
			},
		},
	}
	return planNode, nil
}
