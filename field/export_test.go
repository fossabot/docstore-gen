package field

import (
	"reflect"
	"testing"
	"time"

	"gocloud.dev/docstore"
)

func TestExpr_Build(t *testing.T) {

	var (
		table    = "user"
		column   = "password"
		testTime = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	)

	tests := []struct {
		name          string
		expr          Expr
		wantFieldPath docstore.FieldPath
		wantOp        string
		wantValue     interface{}
	}{
		// ======================== generic ===================================
		{
			name:          "Field-Eq",
			expr:          NewField(table, column).Eq("123"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     "123",
		},
		{
			name:          "Field-Gt",
			expr:          NewField(table, column).Gt(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     123,
		},
		{
			name:          "Field-Gte",
			expr:          NewField(table, column).Gte(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     123,
		},
		{
			name:          "Field-Lt",
			expr:          NewField(table, column).Lt(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     123,
		},
		{
			name:          "Field-Lte",
			expr:          NewField(table, column).Lte(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     123,
		},
		{
			name:          "Field-In",
			expr:          NewField(table, column).In(123, 456, 789),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{123, 456, 789},
		},
		{
			name:          "Field-NotIn",
			expr:          NewField(table, column).NotIn(123, 456, 789),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{123, 456, 789},
		},

		// ======================== string ====================================
		{
			name:          "Field-Eq-String",
			expr:          NewString(table, column).Eq("abc"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     "abc",
		},
		{
			name:          "Field-Gt-String",
			expr:          NewString(table, column).Gt("abc"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     "abc",
		},
		{
			name:          "Field-Gte-String",
			expr:          NewString(table, column).Gte("abc"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     "abc",
		},
		{
			name:          "Field-Lt-String",
			expr:          NewString(table, column).Lt("abc"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     "abc",
		},
		{
			name:          "Field-Lte-String",
			expr:          NewString(table, column).Lte("abc"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     "abc",
		},
		{
			name:          "Field-In-String",
			expr:          NewString(table, column).In("abc", "def", "ghi"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{"abc", "def", "ghi"},
		},
		{
			name:          "Field-NotIn-String",
			expr:          NewString(table, column).NotIn("abc", "def", "ghi"),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{"abc", "def", "ghi"},
		},
		// ======================== bytes =====================================
		{
			name:          "Field-Eq-Bytes",
			expr:          NewBytes(table, column).Eq([]byte{1, 2, 3}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     []byte{1, 2, 3},
		},
		{
			name:          "Field-Gt-Bytes",
			expr:          NewBytes(table, column).Gt([]byte{1, 2, 3}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     []byte{1, 2, 3},
		},
		{
			name:          "Field-Gte-Bytes",
			expr:          NewBytes(table, column).Gte([]byte{1, 2, 3}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     []byte{1, 2, 3},
		},
		{
			name:          "Field-Lt-Bytes",
			expr:          NewBytes(table, column).Lt([]byte{1, 2, 3}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     []byte{1, 2, 3},
		},
		{
			name:          "Field-Lte-Bytes",
			expr:          NewBytes(table, column).Lte([]byte{1, 2, 3}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     []byte{1, 2, 3},
		},
		{
			name:          "Field-In-Bytes",
			expr:          NewBytes(table, column).In([]byte{1, 2, 3}, []byte{4, 5, 6}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{[]byte{1, 2, 3}, []byte{4, 5, 6}},
		},
		{
			name:          "Field-NotIn-Bytes",
			expr:          NewBytes(table, column).NotIn([]byte{1, 2, 3}, []byte{4, 5, 6}),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{[]byte{1, 2, 3}, []byte{4, 5, 6}},
		},

		// ======================== int =======================================
		{
			name:          "Field-Eq-Int",
			expr:          NewInt(table, column).Eq(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     123,
		},
		{
			name:          "Field-Gt-Int",
			expr:          NewInt(table, column).Gt(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     123,
		},
		{
			name:          "Field-Gte-Int",
			expr:          NewInt(table, column).Gte(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     123,
		},
		{
			name:          "Field-Lt-Int",
			expr:          NewInt(table, column).Lt(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     123,
		},
		{
			name:          "Field-Lte-Int",
			expr:          NewInt(table, column).Lte(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     123,
		},
		{
			name:          "Field-In-Int",
			expr:          NewInt(table, column).In(123, 456, 789),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{123, 456, 789},
		},
		{
			name:          "Field-NotIn-Int",
			expr:          NewInt(table, column).NotIn(123, 456, 789),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{123, 456, 789},
		},

		// ======================== int8 ======================================
		{
			name:          "Field-Eq-Int8",
			expr:          NewInt8(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     int8(12),
		},
		{
			name:          "Field-Gt-Int8",
			expr:          NewInt8(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     int8(12),
		},
		{
			name:          "Field-Gte-Int8",
			expr:          NewInt8(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     int8(12),
		},
		{
			name:          "Field-Lt-Int8",
			expr:          NewInt8(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     int8(12),
		},
		{
			name:          "Field-Lte-Int8",
			expr:          NewInt8(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     int8(12),
		},
		{
			name:          "Field-In-Int8",
			expr:          NewInt8(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{int8(12), int8(34), int8(56)},
		},
		{
			name:          "Field-NotIn-Int8",
			expr:          NewInt8(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{int8(12), int8(34), int8(56)},
		},

		// ======================== int16 =====================================
		{
			name:          "Field-Eq-Int16",
			expr:          NewInt16(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     int16(12),
		},
		{
			name:          "Field-Gt-Int16",
			expr:          NewInt16(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     int16(12),
		},
		{
			name:          "Field-Gte-Int16",
			expr:          NewInt16(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     int16(12),
		},
		{
			name:          "Field-Lt-Int16",
			expr:          NewInt16(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     int16(12),
		},
		{
			name:          "Field-Lte-Int16",
			expr:          NewInt16(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     int16(12),
		},
		{
			name:          "Field-In-Int16",
			expr:          NewInt16(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{int16(12), int16(34), int16(56)},
		},
		{
			name:          "Field-NotIn-Int16",
			expr:          NewInt16(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{int16(12), int16(34), int16(56)},
		},
		// ======================== int32 =====================================
		{
			name:          "Field-Eq-Int32",
			expr:          NewInt32(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     int32(12),
		},
		{
			name:          "Field-Gt-Int32",
			expr:          NewInt32(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     int32(12),
		},
		{
			name:          "Field-Gte-Int32",
			expr:          NewInt32(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     int32(12),
		},
		{
			name:          "Field-Lt-Int32",
			expr:          NewInt32(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     int32(12),
		},
		{
			name:          "Field-Lte-Int32",
			expr:          NewInt32(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     int32(12),
		},
		{
			name:          "Field-In-Int32",
			expr:          NewInt32(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{int32(12), int32(34), int32(56)},
		},
		{
			name:          "Field-NotIn-Int32",
			expr:          NewInt32(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{int32(12), int32(34), int32(56)},
		},
		// ======================== int64 =====================================
		{
			name:          "Field-Eq-Int64",
			expr:          NewInt64(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     int64(12),
		},
		{
			name:          "Field-Gt-Int64",
			expr:          NewInt64(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     int64(12),
		},
		{
			name:          "Field-Gte-Int64",
			expr:          NewInt64(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     int64(12),
		},
		{
			name:          "Field-Lt-Int64",
			expr:          NewInt64(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     int64(12),
		},
		{
			name:          "Field-Lte-Int64",
			expr:          NewInt64(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     int64(12),
		},
		{
			name:          "Field-In-Int64",
			expr:          NewInt64(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{int64(12), int64(34), int64(56)},
		},
		{
			name:          "Field-NotIn-Int64",
			expr:          NewInt64(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{int64(12), int64(34), int64(56)},
		},
		// ======================== uint ======================================
		{
			name:          "Field-Eq-Uint",
			expr:          NewUint(table, column).Eq(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     uint(123),
		},
		{
			name:          "Field-Gt-Uint",
			expr:          NewUint(table, column).Gt(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     uint(123),
		},
		{
			name:          "Field-Gte-Uint",
			expr:          NewUint(table, column).Gte(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     uint(123),
		},
		{
			name:          "Field-Lt-Uint",
			expr:          NewUint(table, column).Lt(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     uint(123),
		},
		{
			name:          "Field-Lte-Uint",
			expr:          NewUint(table, column).Lte(123),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     uint(123),
		},
		{
			name:          "Field-In-Uint",
			expr:          NewUint(table, column).In(123, 456, 789),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{uint(123), uint(456), uint(789)},
		},
		{
			name:          "Field-NotIn-Uint",
			expr:          NewUint(table, column).NotIn(123, 456, 789),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{uint(123), uint(456), uint(789)},
		},
		// ======================== uint8 =====================================
		{
			name:          "Field-Eq-Uint8",
			expr:          NewUint8(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     uint8(12),
		},
		{
			name:          "Field-Gt-Uint8",
			expr:          NewUint8(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     uint8(12),
		},
		{
			name:          "Field-Gte-Uint8",
			expr:          NewUint8(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     uint8(12),
		},
		{
			name:          "Field-Lt-Uint8",
			expr:          NewUint8(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     uint8(12),
		},
		{
			name:          "Field-Lte-Uint8",
			expr:          NewUint8(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     uint8(12),
		},
		{
			name:          "Field-In-Uint8",
			expr:          NewUint8(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{uint8(12), uint8(34), uint8(56)},
		},
		{
			name:          "Field-NotIn-Uint8",
			expr:          NewUint8(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{uint8(12), uint8(34), uint8(56)},
		},

		// ======================== uint16 ====================================
		{
			name:          "Field-Eq-Uint16",
			expr:          NewUint16(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     uint16(12),
		},
		{
			name:          "Field-Gt-Uint16",
			expr:          NewUint16(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     uint16(12),
		},
		{
			name:          "Field-Gte-Uint16",
			expr:          NewUint16(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     uint16(12),
		},
		{
			name:          "Field-Lt-Uint16",
			expr:          NewUint16(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     uint16(12),
		},
		{
			name:          "Field-Lte-Uint16",
			expr:          NewUint16(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     uint16(12),
		},
		{
			name:          "Field-In-Uint16",
			expr:          NewUint16(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{uint16(12), uint16(34), uint16(56)},
		},
		{
			name:          "Field-NotIn-Uint16",
			expr:          NewUint16(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{uint16(12), uint16(34), uint16(56)},
		},
		// ======================== uint32 ====================================
		{
			name:          "Field-Eq-Uint32",
			expr:          NewUint32(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     uint32(12),
		},
		{
			name:          "Field-Gt-Uint32",
			expr:          NewUint32(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     uint32(12),
		},
		{
			name:          "Field-Gte-Uint32",
			expr:          NewUint32(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     uint32(12),
		},
		{
			name:          "Field-Lt-Uint32",
			expr:          NewUint32(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     uint32(12),
		},
		{
			name:          "Field-Lte-Uint32",
			expr:          NewUint32(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     uint32(12),
		},
		{
			name:          "Field-In-Uint32",
			expr:          NewUint32(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{uint32(12), uint32(34), uint32(56)},
		},
		{
			name:          "Field-NotIn-Uint32",
			expr:          NewUint32(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{uint32(12), uint32(34), uint32(56)},
		},
		// ======================== uint64 ====================================
		{
			name:          "Field-Eq-Uint64",
			expr:          NewUint64(table, column).Eq(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     uint64(12),
		},
		{
			name:          "Field-Gt-Uint64",
			expr:          NewUint64(table, column).Gt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     uint64(12),
		},
		{
			name:          "Field-Gte-Uint64",
			expr:          NewUint64(table, column).Gte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     uint64(12),
		},
		{
			name:          "Field-Lt-Uint64",
			expr:          NewUint64(table, column).Lt(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     uint64(12),
		},
		{
			name:          "Field-Lte-Uint64",
			expr:          NewUint64(table, column).Lte(12),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     uint64(12),
		},
		{
			name:          "Field-In-Uint64",
			expr:          NewUint64(table, column).In(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{uint64(12), uint64(34), uint64(56)},
		},
		{
			name:          "Field-NotIn-Uint64",
			expr:          NewUint64(table, column).NotIn(12, 34, 56),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{uint64(12), uint64(34), uint64(56)},
		},
		// ======================== float32 ===================================
		{
			name:          "Field-Eq-Float32",
			expr:          NewFloat32(table, column).Eq(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     float32(12.34),
		},
		{
			name:          "Field-Gt-Float32",
			expr:          NewFloat32(table, column).Gt(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     float32(12.34),
		},
		{
			name:          "Field-Gte-Float32",
			expr:          NewFloat32(table, column).Gte(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     float32(12.34),
		},
		{
			name:          "Field-Lt-Float32",
			expr:          NewFloat32(table, column).Lt(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     float32(12.34),
		},
		{
			name:          "Field-Lte-Float32",
			expr:          NewFloat32(table, column).Lte(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     float32(12.34),
		},
		{
			name:          "Field-In-Float32",
			expr:          NewFloat32(table, column).In(12.34, 56.78, 91.01),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{float32(12.34), float32(56.78), float32(91.01)},
		},
		{
			name:          "Field-NotIn-Float32",
			expr:          NewFloat32(table, column).NotIn(12.34, 56.78, 91.01),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{float32(12.34), float32(56.78), float32(91.01)},
		},

		// ======================== float64 ===================================
		{
			name:          "Field-Eq-Float64",
			expr:          NewFloat64(table, column).Eq(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     float64(12.34),
		},
		{
			name:          "Field-Gt-Float64",
			expr:          NewFloat64(table, column).Gt(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     float64(12.34),
		},
		{
			name:          "Field-Gte-Float64",
			expr:          NewFloat64(table, column).Gte(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     float64(12.34),
		},
		{
			name:          "Field-Lt-Float64",
			expr:          NewFloat64(table, column).Lt(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     float64(12.34),
		},
		{
			name:          "Field-Lte-Float64",
			expr:          NewFloat64(table, column).Lte(12.34),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     float64(12.34),
		},
		{
			name:          "Field-In-Float64",
			expr:          NewFloat64(table, column).In(12.34, 56.78, 91.01),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{float64(12.34), float64(56.78), float64(91.01)},
		},
		{
			name:          "Field-NotIn-Float64",
			expr:          NewFloat64(table, column).NotIn(12.34, 56.78, 91.01),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{float64(12.34), float64(56.78), float64(91.01)},
		},
		// ======================== bool ======================================
		{
			name:          "Field-Eq-Bool",
			expr:          NewBool(table, column).Eq(true),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     true,
		},
		// ======================== time ======================================
		{
			name:          "Field-Eq-Time",
			expr:          NewTime(table, column).Eq(testTime),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "=",
			wantValue:     testTime,
		},
		{
			name:          "Field-Gt-Time",
			expr:          NewTime(table, column).Gt(testTime),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">",
			wantValue:     testTime,
		},
		{
			name:          "Field-Gte-Time",
			expr:          NewTime(table, column).Gte(testTime),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        ">=",
			wantValue:     testTime,
		},
		{
			name:          "Field-Lt-Time",
			expr:          NewTime(table, column).Lt(testTime),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<",
			wantValue:     testTime,
		},
		{
			name:          "Field-Lte-Time",
			expr:          NewTime(table, column).Lte(testTime),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "<=",
			wantValue:     testTime,
		},
		{
			name:          "Field-In-Time",
			expr:          NewTime(table, column).In(testTime, testTime.Add(1*time.Hour), testTime.Add(2*time.Hour)),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "in",
			wantValue:     []interface{}{testTime, testTime.Add(1 * time.Hour), testTime.Add(2 * time.Hour)},
		},
		{
			name:          "Field-NotIn-Time",
			expr:          NewTime(table, column).NotIn(testTime, testTime.Add(1*time.Hour), testTime.Add(2*time.Hour)),
			wantFieldPath: docstore.FieldPath(column),
			wantOp:        "not-in",
			wantValue:     []interface{}{testTime, testTime.Add(1 * time.Hour), testTime.Add(2 * time.Hour)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFieldPath, gotOp, gotValue := tt.expr.Build()

			if tt.wantFieldPath != gotFieldPath {
				t.Errorf("FieldPath: want %v, got %v", tt.wantFieldPath, gotFieldPath)
			}
			if tt.wantOp != gotOp {
				t.Errorf("Op: want %v, got %v", tt.wantOp, gotOp)
			}
			if !reflect.DeepEqual(tt.wantValue, gotValue) {
				t.Errorf("Value: want %v, got %v", tt.wantValue, gotValue)
			}
		})
	}
}

func TestFieldExpr_BuildOrderBy(t *testing.T) {
	table := "table"
	column := "column"
	tests := []struct {
		name          string
		expr          OrderByExpression
		wantFieldPath docstore.FieldPath
		wantDirection string
	}{
		{
			name:          "Field-Asc",
			expr:          NewString(table, column).Asc(),
			wantFieldPath: docstore.FieldPath(column),
			wantDirection: docstore.Ascending,
		},
		{
			name:          "Field-Desc",
			expr:          NewString(table, column).Desc(),
			wantFieldPath: docstore.FieldPath(column),
			wantDirection: docstore.Descending,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFieldPath, gotAsc := tt.expr.BuildOrderBy()

			if tt.wantFieldPath != docstore.FieldPath(gotFieldPath) {
				t.Errorf("FieldPath: want %v, got %v", tt.wantFieldPath, gotFieldPath)
			}
			if tt.wantDirection != gotAsc {
				t.Errorf("Direction: want %v, got %v", tt.wantDirection, gotAsc)
			}
		})
	}
}
