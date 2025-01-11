package test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"ticket/util"
)

func TestFile(t *testing.T) {
	ociCat := util.SongPath + util.OciCat
	dir, _ := ioutil.ReadDir(ociCat)
	for _, d := range dir {
		_ = os.RemoveAll(path.Join([]string{ociCat, d.Name()}...))
	}
}
