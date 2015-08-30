package filesystem

import (
	"io/ioutil"
	"os"
	"testing"
	"fmt"

	storagedriver "github.com/docker/distribution/registry/storage/driver"
	"github.com/docker/distribution/registry/storage/driver/testsuites"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	fmt.Printf("I AM HERE");
	root, err := ioutil.TempDir("", "driver-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(root)

	testsuites.RegisterSuite(func() (storagedriver.StorageDriver, error) {
		return New(root), nil
	}, testsuites.NeverSkip)
}
