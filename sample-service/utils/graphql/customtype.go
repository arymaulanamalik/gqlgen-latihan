package graphql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

func MarshalDecimal(b decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, fmt.Sprintf("%v", b.InexactFloat64()))
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch v := v.(type) {
	case string:
		return decimal.NewFromString(v)
	case float64:
		return decimal.NewFromFloat(v), nil
	case float32:
		return decimal.NewFromFloat32(v), nil
	case int64:
		return decimal.NewFromInt(v), nil
	case int32:
		return decimal.NewFromInt32(v), nil
	default:
		return decimal.NewFromInt(0), fmt.Errorf("%T is not supported for decimal", v)
	}
}

func MarshalFloat32(b float32) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, fmt.Sprintf("%v", float64(b)))
	})
}

func UnmarshalFloat32(v interface{}) (float32, error) {
	switch v := v.(type) {
	case string:
		value, err := strconv.ParseFloat(v, 32)
		return float32(value), err
	case float64:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case int32:
		return float32(v), nil
	default:
		return float32(0), fmt.Errorf("%T is not supported for float", v)
	}
}
