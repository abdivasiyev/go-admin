package iris

import (
	"net/http"
	"testing"

	"github.com/abdivasiyev/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestIris(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
