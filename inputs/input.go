package inputs

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func Load(day int) string {
	f := generateFilename(day)
	contents, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func generateFilename(day int) string {
	_, d, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not determine current caller")
	}
	return fmt.Sprintf("%s/day_%d", filepath.Dir(d), day)
}
