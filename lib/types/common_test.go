package types

import (
	"reflect"
	"testing"
	"time"
)

func TestCurrentTime_CurrentTime(t1 *testing.T) {
	mockTime := time.Date(2023, 12, 1, 1, 1, 1, 1, time.UTC)
	type fields struct {
		Mock bool
		Now  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "test_mock_time",
			fields: fields{
				Mock: true,
				Now:  mockTime,
			},
			want: mockTime,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &CurrentTime{
				Mock: tt.fields.Mock,
				Now:  tt.fields.Now,
			}
			if got := t.CurrentTime(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
