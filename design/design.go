package design

// The goa API design language is a DSL implemented in Go and is not Go.
// The generated code or any of the actual Go code in goa does not make
// use of “dot imports”. Using this technique for the DSL results in far
// cleaner looking code.
//
// For more details refer to https://goa.design/design/overview/

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Common service APIs", func() {
	Title("Common service APIs documentation")
	Description("This API includes a list of common utilities which can be used by any participants in our system")
	Host("localhost:3001")
	Scheme("http")
	BasePath("api/v0")
})

var _ = Resource("Storage", func() {
	Action("cat", func() {
		Routing(GET("storage/:address"))
		Description("Cat the file in IPFS at :address")
		Params(func() {
			Param("address", String, "IPFS address")
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
	})

	Action("add", func() {
		Routing(POST("storage"))
		Description("Upload file to IPFS using multipart post")
		Payload(FilePayload)
		MultipartForm()
		Response(OK)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})

var FilePayload = Type("FilePayload", func() {
	Attribute("file", File, "file")
	Required("file")
})

var _ = Resource("Public", func() {
	// CORS policy that applies to all actions and file servers of "public" resource
	Origin("*", func() {
		Methods("GET")
	})
	Files("/schema/*filepath", "public/schema/")
	Files("/swagger/*filepath", "public/swagger/")
})
