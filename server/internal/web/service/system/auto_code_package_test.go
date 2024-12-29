package system

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"reflect"
	"testing"
)

func Test_autoCodePackage_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		info *system2.SysAutoCodePackageCreate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试 package",
			args: args{
				ctx: context.Background(),
				info: &system2.SysAutoCodePackageCreate{
					Template:    "package",
					PackageName: "gva",
				},
			},
			wantErr: false,
		},
		{
			name: "测试 plugin",
			args: args{
				ctx: context.Background(),
				info: &system2.SysAutoCodePackageCreate{
					Template:    "plugin",
					PackageName: "gva",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AutoCodePackageService{}
			if err := a.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_autoCodePackage_templates(t *testing.T) {
	type args struct {
		ctx    context.Context
		entity system.SysAutoCodePackage
		info   system2.AutoCode
	}
	tests := []struct {
		name      string
		args      args
		wantCode  map[string]string
		wantEnter map[string]map[string]string
		wantErr   bool
	}{
		{
			name: "测试1",
			args: args{
				ctx: context.Background(),
				entity: system.SysAutoCodePackage{
					Desc:        "描述",
					Label:       "展示名",
					Template:    "plugin",
					PackageName: "preview",
				},
				info: system2.AutoCode{
					Abbreviation:    "user",
					HumpPackageName: "user",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AutoCodePackageService{}
			gotCode, gotEnter, gotCreates, err := s.templates(tt.args.ctx, tt.args.entity, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("templates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for key, value := range gotCode {
				t.Logf("\n")
				t.Logf(key)
				t.Logf(value)
				t.Logf("\n")
			}
			t.Log(gotCreates)
			if !reflect.DeepEqual(gotEnter, tt.wantEnter) {
				t.Errorf("templates() gotEnter = %v, want %v", gotEnter, tt.wantEnter)
			}
		})
	}
}
