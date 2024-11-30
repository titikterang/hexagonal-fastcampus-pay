package common

import (
	"github.com/shopspring/decimal"
	"google.golang.org/genproto/googleapis/type/money"
	"reflect"
	"testing"
)

func TestMoneyToString(t *testing.T) {
	var data money.Money
	data.CurrencyCode = "IDR"
	data.Nanos = 758009000
	data.Units = 30000

	type args struct {
		data *money.Money
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_convert_to_string",
			args: args{
				data: &data,
			},
			want: "30000.758009",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoneyToString(tt.args.data); got != tt.want {
				t.Errorf("MoneyToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimalToMoney(t *testing.T) {
	data, _ := decimal.NewFromString("30000.758009")

	type args struct {
		data decimal.Decimal
	}
	tests := []struct {
		name      string
		args      args
		wantUnit  int64
		wantNanos int32
	}{
		{
			name: "test_decimal_to_money",
			args: args{
				data: data,
			},
			wantUnit:  30000,
			wantNanos: 758009000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnits := DecimalToMoney(tt.args.data).Units
			gotNanos := DecimalToMoney(tt.args.data).Nanos
			if !reflect.DeepEqual(gotUnits, tt.wantUnit) || !reflect.DeepEqual(gotNanos, tt.wantNanos) {
				t.Errorf("DecimalToMoney() gotUnits = %v, wantUnit %v , gotNanos = %v, wantNanos %v ,", gotUnits, tt.wantUnit, gotNanos, tt.wantNanos)
			}
		})
	}
}
