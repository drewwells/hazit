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

	c := this.AppEngineCtx

	bucketName := "hazzzit.appspot.com"

	conf := google.NewAppEngineConfig(
		c, storage.ScopeFullControl)

	ctx := cloud.NewContext(appengine.AppID(c),
		&http.Client{Transport: conf.NewTransport()})

	rc, err := storage.NewReader(
		ctx,
		bucketName,
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
