package controller_test

import (
	"github.com/garethjevans/filter-controller/internal/controller"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {

	files, err := controller.ListFiles("../..")
	assert.NoError(t, err)

	t.Logf("Got files %s", files)

	include := `!.git
go.*
internal/**/*.go
api
!.*
!**/*_test.go
!**/tests`

	filtered := controller.FilterFileList(files, include)

	t.Logf("Filtered files %s", filtered)

	assert.True(t, len(filtered) > 0)

	assert.Contains(t, filtered, "go.mod")
	assert.Contains(t, filtered, "go.sum")
	assert.Contains(t, filtered, "internal/controller/filter_controller.go")
	assert.Contains(t, filtered, "api/v1alpha1/filter_types.go")

	assert.NotContains(t, filtered, ".gitignore")
	assert.NotContains(t, filtered, ".github/dependabot.yml")
	assert.NotContains(t, filtered, ".idea/.gitignore")
}

func TestHash(t *testing.T) {
	include := `*.txt`

	files01, err := controller.ListFiles("testdata/dir01")
	assert.NoError(t, err)

	t.Logf("Got files %s", files01)

	filtered01 := controller.FilterFileList(files01, include)

	t.Logf("Filtered files %s", filtered01)

	assert.Equal(t, len(filtered01), 1)
	assert.Contains(t, filtered01, "test.txt")

	hash01, err := controller.HashFiles(filtered01, "testdata/dir01")

	assert.NoError(t, err)
	t.Logf("Hash %s", hash01)

	files02, err := controller.ListFiles("testdata/dir02")
	assert.NoError(t, err)

	t.Logf("Got files %s", files02)

	filtered02 := controller.FilterFileList(files02, include)

	t.Logf("Filtered files %s", filtered02)

	assert.Equal(t, len(filtered02), 1)
	assert.Contains(t, filtered02, "test.txt")

	hash02, err := controller.HashFiles(filtered02, "testdata/dir02")

	assert.NoError(t, err)
	t.Logf("Hash %s", hash02)

	assert.Equal(t, hash01, hash02)
}
