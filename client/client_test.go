package client

import (
	"testing"
	"time"

	"github.com/bemyth/influx-stress/point"
)

type MockPoint struct {
}

func (m *MockPoint) Marshal() string {
	return "cpu,location=en used=98.0 1573108345000000000"
}
func Test_client_Write(t *testing.T) {
	type args struct {
		pt point.Point
	}
	tests := []struct {
		name    string
		c       *client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			c:       NewClient("", "", "", ""),
			args:    args{pt: &MockPoint{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Write(tt.args.pt); (err != nil) != tt.wantErr {
				t.Errorf("client.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_WritePoint(t *testing.T) {
	c := NewClient("", "8086", "", "")
	go c.run()
	c.Write(&MockPoint{})
	time.Sleep(10 * time.Second)
}
