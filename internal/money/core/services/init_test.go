package services

import (
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		cfg        *config.Config
		repository ports.MoneyRepositoryAdapter
	}
	tests := []struct {
		name string
		args args
		want *MoneyService
	}{
		{
			name: "test_init_money-service",
			args: args{
				cfg:        nil,
				repository: nil,
			},
			want: &MoneyService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.cfg, tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
