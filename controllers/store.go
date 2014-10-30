package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	context "code.google.com/p/go.net/context"
	"github.com/astaxie/beegae"
	"github.com/golang/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"

	"appengine"
)

type StoreController struct {
	beegae.Controller
}

// Get if passed a name looks for that image and shows it
// Otherwise, all images in GCS are rendered on the page.
func (this *StoreController) Get() {

	c := this.AppEngineCtx
	// This should be global in the Controller
	conf := google.NewAppEngineConfig(
		c, storage.ScopeFullControl)

	ctx := cloud.NewContext(appengine.AppID(c),
		&http.Client{Transport: conf.NewTransport()})

	name := this.GetString("name")
	this.Data["name"] = "boom"
	if name == "" {
		this.List(ctx)
		return
	}

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

func (this *StoreController) List(ctx context.Context) {

	var query *storage.Query
	var names []string
	for {
		// If you are using this package on App Engine Managed VMs runtime,
		// you can init a bucket client with your app's default bucket name.
		// See http://godoc.org/google.golang.org/appengine/file#DefaultBucketName.
		objects, err := storage.ListObjects(ctx, BucketName, query)
		if err != nil {
			log.Fatal(err)
		}
		for _, obj := range objects.Results {
			log.Printf("object name: %s, size: %v", obj.Name, obj.Size)
			names = append(names, obj.Name)
		}
		// if there are more results, objects.Next
		// will be non-nil.
		query = objects.Next
		if query == nil {
			break
		}
	}

	this.Data["Names"] = &names
	this.Layout = "layout.tpl"
	this.TplNames = "storecontroller/list.tpl"

}
