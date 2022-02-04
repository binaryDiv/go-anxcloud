package v1

import (
	. "github.com/onsi/ginkgo/v2"
	testutils "go.anx.io/go-anxcloud/pkg/utils/test"

	"go.anx.io/go-anxcloud/pkg/api/types"
)

var _ = Describe("Object Record", func() {
	o := Record{}

	ifaces := make([]interface{}, 0, 3)
	{
		var i types.Object
		ifaces = append(ifaces, &i)
	}
	{
		var i types.ResponseDecodeHook
		ifaces = append(ifaces, &i)
	}
	{
		var i types.PaginationSupportHook
		ifaces = append(ifaces, &i)
	}

	testutils.ObjectTests(&o, ifaces...)
})

var _ = Describe("Object Zone", func() {
	o := Zone{}

	ifaces := make([]interface{}, 0, 5)
	{
		var i types.Object
		ifaces = append(ifaces, &i)
	}
	{
		var i types.RequestFilterHook
		ifaces = append(ifaces, &i)
	}
	{
		var i types.RequestBodyHook
		ifaces = append(ifaces, &i)
	}
	{
		var i types.ResponseFilterHook
		ifaces = append(ifaces, &i)
	}
	{
		var i types.PaginationSupportHook
		ifaces = append(ifaces, &i)
	}

	testutils.ObjectTests(&o, ifaces...)
})