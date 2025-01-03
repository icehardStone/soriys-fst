package config

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		key     string
		wantErr bool
	}{
		// TODO: Add test cases.
		{key: "dns", wantErr: false},
	}
	os.Setenv("APP_DATABASE_DNS", "dns1")
	os.Setenv("APP_VERSION", "2.0")
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if err := InitPath("../"); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, true, AppConfig.Version != nil)
			assert.Equal(t, AppConfig.Database.DNS, "dns1")
		})
	}
}
