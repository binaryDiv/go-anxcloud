package v1

import (
	. "github.com/onsi/ginkgo/v2"
	testutils "go.anx.io/go-anxcloud/pkg/utils/test"

	"go.anx.io/go-anxcloud/pkg/api/types"
)

var _ = Describe("Object VLAN", func() {
	o := VLAN{}

	ifaces := make([]interface{}, 0, 2)
	{
		var i types.Object
		ifaces = append(ifaces, &i)
	}
	{
		var i types.RequestBodyHook
		ifaces = append(ifaces, &i)
	}

	testutils.ObjectTests(&o, ifaces...)
})