package order

import (
	"github.com/anvlad11/testapp-20230927/pkg/config"
	"reflect"
	"testing"
)

func TestService_CreateOrder(t *testing.T) {
	type fields struct {
		cfg config.Orders
	}
	type args struct {
		itemsCount int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "TestShouldSucceedWith1Item",
			fields: fields{
				cfg: config.Orders{PackSizes: []int{250, 500, 1000, 2000, 5000}},
			},
			args: args{
				itemsCount: 1,
			},
			want:    map[int]int{250: 1},
			wantErr: false,
		}, {
			name: "TestShouldSucceedWith250Items",
			fields: fields{
				cfg: config.Orders{PackSizes: []int{250, 500, 1000, 2000, 5000}},
			},
			args: args{
				itemsCount: 250,
			},
			want:    map[int]int{250: 1},
			wantErr: false,
		}, {
			name: "TestShouldSucceedWith251Items",
			fields: fields{
				cfg: config.Orders{PackSizes: []int{250, 500, 1000, 2000, 5000}},
			},
			args: args{
				itemsCount: 251,
			},
			want:    map[int]int{500: 1},
			wantErr: false,
		},
		{
			name: "TestShouldSucceedWith501Items",
			fields: fields{
				cfg: config.Orders{PackSizes: []int{250, 500, 1000, 2000, 5000}},
			},
			args: args{
				itemsCount: 501,
			},
			want:    map[int]int{500: 1, 250: 1},
			wantErr: false,
		},
		{
			name: "TestShouldSucceedWith12001Items",
			fields: fields{
				cfg: config.Orders{PackSizes: []int{250, 500, 1000, 2000, 5000}},
			},
			args: args{
				itemsCount: 12001,
			},
			want:    map[int]int{5000: 2, 2000: 1, 250: 1},
			wantErr: false,
		},
		{
			name: "TestShouldFailWith0Items",
			fields: fields{
				cfg: config.Orders{PackSizes: []int{250, 500, 1000, 2000, 5000}},
			},
			args: args{
				itemsCount: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.cfg)
			got, err := s.CreateOrder(tt.args.itemsCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
