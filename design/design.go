

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("common services", func() {
	Title("The common service APIs")
	Description("This is the common service APIs, which could be used by anyone")
	Host("localhost:3001")
	Scheme("http")
})

var _ = Resource("storage", func() {
	Action("cat", func() {
		Routing(GET("storage/:address"))
		Description("Cat the file in distributed storage by address")
		Params(func() {
			Param("address", String, "IPFS address")
		})
		Response(OK, "text/plain")
	})

	Action("add", func() {
		Routing(POST("storage"))
		Description("Upload file in multipart request")
		Payload(FilePayload)
		MultipartForm()
		Response(OK)
	})
})

var FilePayload = Type("FilePayload", func() {
	Attribute("file", File, "file")
})

var _ = Resource("schema", func() {
	Files("/schema/*filepath", "public/schema/")
})

var _ = Resource("swagger", func() {
	Files("/swagger/*filepath", "public/swagger/")
})