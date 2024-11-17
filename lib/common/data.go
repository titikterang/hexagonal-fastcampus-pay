package common

import (
	"github.com/shopspring/decimal"
	money "google.golang.org/genproto/googleapis/type/money"
	"math"
	"strings"
)

func MoneyToString(data *money.Money) string {
	if data == nil {
		return "0"
	}
	nanos := data.GetNanos()
	val := data.GetUnits()

	decimals := float64(nanos) * (math.Pow(10.0, -9.0))
	mainPart := decimal.NewFromInt(val)
	decimalPart := decimal.NewFromFloat(decimals)

	result := mainPart.Add(decimalPart)
	return result.String()
}

func DecimalToMoney(data decimal.Decimal) *money.Money {
	result := strings.Split(data.String(), ".")
	if len(result) == 1 {
		return &money.Money{
			CurrencyCode: "IDR",
			Units:        data.IntPart(),
			Nanos:        0,
		}
	}

	dataResult := &money.Money{
		CurrencyCode: "IDR",
		Units:        data.IntPart(),
	}

	unitPart, _ := decimal.NewFromString(result[1])
	multiplier := math.Pow(10.0, 9.0-float64(len(unitPart.String())))
	nanos := unitPart.Mul(decimal.NewFromFloat(multiplier)).IntPart()
	dataResult.Nanos = int32(nanos)

	return dataResult
}
