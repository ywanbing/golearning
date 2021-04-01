package goloang_c_script

import "testing"

func TestExecCmd(t *testing.T) {
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test ls",
			args: args{
				name: "ls",
				args: []string{"-l", "-a"},
			},
			wantErr: false,
		},
		{
			name: "test sh",
			args: args{
				name: "./test.sh",
				args: []string{""},
			},
			wantErr: false,
		},
		{
			name: "test sh args",
			args: args{
				name: "./test.sh",
				args: []string{"1", "2", "3"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExecCmd(tt.args.name, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("ExecCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
