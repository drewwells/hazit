// Controllers routes requests to the approriate data and
// views for rending that data.
package controllers

import (
	"log"

	"github.com/astaxie/beegae"
)

var BucketName string // Name of your GCS bucket https://cloud.google.com/appengine/docs/go/googlecloudstorageclient/

func init() {
	BucketName = beegae.AppConfig.String("bucket")
	if BucketName == "" {
		log.Fatal("Set the bucket name in conf/app.conf")
	}
}

type MainController struct {
	beegae.Controller
}

func (this *MainController) Get() {
	this.TplNames = "form.tpl"
}
