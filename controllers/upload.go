package controllers

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beegae"
)

type UploadController struct {
	beegae.Controller
}

func (c *UploadController) Get() {
	fmt.Println("This happened")
	c.Data["json"] = "This is great place to be"
	c.ServeJson()
}

func (c *UploadController) Post() {

	file, header, err := c.GetFile("file")
	_ = header

	defer file.Close()

	if err != nil {
		c.Data["json"] = err
		c.AppEngineCtx.Errorf("upload err: %v", err)
		return
	}

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		c.Data["json"] = err
		c.AppEngineCtx.Errorf("create err: %v", err)
		return
	}

	c.Data["json"] = struct {
		StatusText string
		Content    string
	}{
		"success", string(bs),
	}
}

/*func (this *MainController) Get() {
	todos := []models.Todo{}
	ks, err := datastore.NewQuery("Todo").Ancestor(models.DefaultTodoList(this.AppEngineCtx)).Order("Created").GetAll(this.AppEngineCtx, &todos)
	if err != nil {
		this.Data["json"] = err
		return
	}
	for i := 0; i < len(todos); i++ {
		todos[i].Id = ks[i].IntID()
	}
	this.Data["json"] = todos
}

func (this *MainController) Post() {
	todo, err := decodeTodo(this.Ctx.Input.Request.Body)
	if err != nil {
		this.Data["json"] = err
		return
	}
	t, err := todo.Save(this.AppEngineCtx)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = &t
	}
}

func (this *MainController) Delete() {
	err := datastore.RunInTransaction(this.AppEngineCtx, func(c appengine.Context) error {
		ks, err := datastore.NewQuery("Todo").KeysOnly().Ancestor(models.DefaultTodoList(c)).Filter("Done=", true).GetAll(c, nil)
		if err != nil {
			return err
		}
		return datastore.DeleteMulti(c, ks)
	}, nil)

	if err == nil {
		this.Data["json"] = nil
	} else {
		this.Data["json"] = err
	}
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
}*/

func (c *UploadController) Render() error {
	c.ServeJson()
	return nil
}
