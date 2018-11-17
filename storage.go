package main

import (
	"bytes"
	"fmt"
	"github.com/ZJU-DistributedAI/Common/app"
	"github.com/goadesign/goa"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

var IPFS_HOST = os.Getenv("IPFS_HOST")

// StorageController implements the storage resource.
type StorageController struct {
	*goa.Controller
}

// NewStorageController creates a storage controller.
func NewStorageController(service *goa.Service) *StorageController {
	return &StorageController{Controller: service.NewController("StorageController")}
}

func (c *StorageController) Cat(ctx *app.CatStorageContext) error {
	var url = fmt.Sprintf("http://%s:5001/api/v0/cat?arg=%s", IPFS_HOST, ctx.Address)
	goa.LogInfo(ctx, fmt.Sprintf("Calling IPFS Cat API, at url: %s", url))
	resp, err := http.Get(url)
	if err != nil {
		goa.LogError(ctx, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		goa.LogError(ctx, err.Error())
	}
	return ctx.OK([]byte(body))
}

func (c *StorageController) Add(ctx *app.AddStorageContext) error {
	file, err := ctx.Payload.File.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("", "")
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s:5001/api/v0/add", IPFS_HOST)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return ctx.OK([]byte(responseBody))
}
