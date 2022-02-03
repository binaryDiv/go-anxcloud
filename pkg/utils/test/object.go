package test

import (
	"context"
	"fmt"
	"net/url"
	"reflect"

	"go.anx.io/go-anxcloud/pkg/api"
	"go.anx.io/go-anxcloud/pkg/api/types"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

type hookErrorCheck func(context.Context) error

// ObjectTests contains the logic to test any Object implementation with the hooks it implements, checking
// * if the Object actually implements the interface of the hook
// * if the Object has an identifier field
// * calling the hook function with incomplete contexts gives none or the correct error (meaning the error is handled correctly)
//
func ObjectTests(o types.Object, hooks ...interface{}) {
	ginkgo.It("has an identifier", func() {
		_, err := api.GetObjectIdentifier(o, false)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	})

	for _, hook := range hooks {
		name := reflect.TypeOf(hook).Elem().Name()

		supportedHooks := map[string]hookErrorCheck{
			"Object": func(ctx context.Context) error {
				_, err := o.EndpointURL(ctx)
				return err
			},
			"PaginationSupportHook": func(ctx context.Context) error {
				_, err := o.(types.PaginationSupportHook).HasPagination(ctx)
				return err
			},
		}

		if errorCheck, ok := supportedHooks[name]; ok {
			ginkgo.Context(fmt.Sprintf("implementing %v", name), func() {
				ginkgo.It("actually implements the interface", func() {
					implementsHook := reflect.TypeOf(o).Implements(reflect.TypeOf(hook).Elem())
					gomega.Expect(implementsHook).To(gomega.BeTrue())
				})

				args := []interface{}{
					func(ctx context.Context) {
						err := errorCheck(ctx)

						// It can be fine with incomplete context, but if it fails, than with the error indicating
						// it checked for it - the one OperationFromContext and co. return
						if err != nil {
							gomega.Expect(err).To(
								gomega.MatchError(types.ErrContextKeyNotSet),
							)
						}
					},

					// technically we have to omit "url" for checking EndpointURL, as it is not set there, yet
					// EndpointURL using URL from context should crash in other tests though, so we take the simplicity of only adding a test case
					// for others, instead of generating different sets of test cases.
					ginkgo.Entry("missing options", makeTestContext("operation", "url")),
					ginkgo.Entry("missing operation", makeTestContext("options", "url")),
				}

				if name != "Object" {
					args = append(args,
						ginkgo.Entry("missing url", makeTestContext("options", "operation")),
					)
				}

				ginkgo.DescribeTable("handles being called with context", args...)
			})
		}
	}
}

func makeTestContext(elems ...string) context.Context {
	ctx := context.TODO()

	for _, elem := range elems {
		switch elem {
		case "operation":
			ctx = types.ContextWithOperation(ctx, types.OperationList)
		case "options":
			ctx = types.ContextWithOptions(ctx, &types.ListOptions{})
		case "url":
			u, _ := url.Parse("http://localhost:1312")
			ctx = types.ContextWithURL(ctx, *u)
		}
	}

	return ctx
}
