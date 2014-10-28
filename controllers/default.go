// Controllers routes requests to the approriate data and
// views for rending that data.
package controllers

import (
	"encoding/json"
	"io"
	"log"

	"models"

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
	//this.TplNames = "views/form.tpl"

}

func (this *MainController) Render() error {
	if _, ok := this.Data["json"].(error); ok {
		this.AppEngineCtx.Errorf("todo error: %v", this.Data["json"])
	}
	this.ServeJson()
	return nil
}

func decodeTodo(r io.ReadCloser) (*models.Todo, error) {
	defer r.Close()
	var todo models.Todo
	err := json.NewDecoder(r).Decode(&todo)
	return &todo, err
}
