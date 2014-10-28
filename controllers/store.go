package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beegae"
	"github.com/golang/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"

	"appengine"
)

type StoreController struct {
	beegae.Controller
}

func (this StoreController) Get() {

	name := this.GetString("name")
	if name == "" {
		this.List()
		return
	}
	c := this.AppEngineCtx

	conf := google.NewAppEngineConfig(
		c, storage.ScopeFullControl)

	ctx := cloud.NewContext(appengine.AppID(c),
		&http.Client{Transport: conf.NewTransport()})

	rc, err := storage.NewReader(
		ctx,
		BucketName,
		name)
	if err != nil {
		log.Fatal(err)
	}
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Fatal(err)
	}
	rc.Close()

	fmt.Fprintf(this.Ctx.ResponseWriter, "%s", slurp)

}

func (this StoreController) List() {
	ctx := cloud.NewContext("project-id", &http.Client{Transport: nil})

	var query *storage.Query
	for {
		// If you are using this package on App Engine Managed VMs runtime,
		// you can init a bucket client with your app's default bucket name.
		// See http://godoc.org/google.golang.org/appengine/file#DefaultBucketName.
		objects, err := storage.ListObjects(ctx, "bucketname", query)
		if err != nil {
			log.Fatal(err)
		}
		for _, obj := range objects.Results {
			log.Printf("object name: %s, size: %v", obj.Name, obj.Size)
		}
		// if there are more results, objects.Next
		// will be non-nil.
		query = objects.Next
		if query == nil {
			break
		}
	}

}
