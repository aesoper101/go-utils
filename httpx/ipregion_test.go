package httpx

import "testing"

func TestGetRegionFromIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test TestGetRegionFromIP 1.2.3.4 Ok",
			args:    args{ip: "1.2.3.4"},
			want:    "美国|0|华盛顿|0|谷歌",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegionFromIP(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegionFromIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRegionFromIP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
