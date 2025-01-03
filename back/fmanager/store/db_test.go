package store

import (
	"testing"

	"gotest.tools/v3/assert"
)

type Version struct {
	Release string
}

func TestGetDb(t *testing.T) {
	tests := []struct {
		name string
		open bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDb := NewStore().db; gotDb != nil {
				err := gotDb.AutoMigrate(&Version{})
				assert.NilError(t, err)
				// err = gotDb.Migrator().DropTable(&Version{})
				assert.NilError(t, err)
			}
		})
	}
}
