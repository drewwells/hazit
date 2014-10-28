// EXAMPLE FROM: https://github.com/GoogleCloudPlatform/appengine-angular-gotodos
//
// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
package controllers

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"unicode"

	"github.com/astaxie/beegae"
	"github.com/golang/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"

	"appengine"

	//"appengine/file"
)

type UploadController struct {
	beegae.Controller
}

func fileName() string {
	n := 10
	g := big.NewInt(0)
	max := big.NewInt(130)
	bs := make([]byte, n)

	for i, _ := range bs {
		g, _ = rand.Int(rand.Reader, max)
		r := rune(g.Int64())
		for !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			g, _ = rand.Int(rand.Reader, max)
			r = rune(g.Int64())
		}
		bs[i] = byte(g.Int64())
	}
	return string(bs)
}

// Render a form for submitting file data
func (this *UploadController) Get() {
	this.TplNames = "form.tpl"
}

// Post takes as input a form with file data.  This is stored
// in the configured GCS bucket.
func (this *UploadController) Post() {
	var bs bytes.Buffer
	upload, header, err := this.GetFile("file")
	_ = header
	c := this.AppEngineCtx
	bs.ReadFrom(upload)

	if err != nil {
		c.Errorf("File error: %s", err)
	}

	conf := google.NewAppEngineConfig(
		c, storage.ScopeFullControl)

	ctx := cloud.NewContext(appengine.AppID(c),
		&http.Client{Transport: conf.NewTransport()})
	name := fileName()

	err = storage.PutDefaultACLRule(ctx, BucketName, "allUsers", storage.RoleReader)
	if err != nil {
		c.Errorf("%s", err)
	}
	wc := storage.NewWriter(ctx,
		BucketName,
		name,
		&storage.Object{
			ContentType: "text/plain",
		})

	if _, err := wc.Write(bs.Bytes()); err != nil {
		c.Errorf("%v", err)
	}

	if err := wc.Close(); err != nil {
		c.Errorf("Write Error: %s", err)
	}
	_, err = wc.Object()
	if err != nil {
		c.Errorf("Object Creation %s", err)
	}

	this.Data["Name"] = name
}

func (this *UploadController) dumpStats(obj *storage.Object) {
	r := this.Ctx.ResponseWriter
	fmt.Fprintf(r, "(filename: /%v/%v, ", obj.Bucket, obj.Name)
	fmt.Fprintf(r, "ContentType: %q, ", obj.ContentType)
	fmt.Fprintf(r, "ACL: %#v, ", obj.ACL)
	fmt.Fprintf(r, "Owner: %v, ", obj.Owner)
	fmt.Fprintf(r, "ContentEncoding: %q, ", obj.ContentEncoding)
	fmt.Fprintf(r, "Size: %v, ", obj.Size)
	fmt.Fprintf(r, "MD5: %q, ", obj.MD5)
	fmt.Fprintf(r, "CRC32C: %q, ", obj.CRC32C)
	fmt.Fprintf(r, "Metadata: %#v, ", obj.Metadata)
	fmt.Fprintf(r, "MediaLink: %q, ", obj.MediaLink)
	fmt.Fprintf(r, "StorageClass: %q, ", obj.StorageClass)
	if !obj.Deleted.IsZero() {
		fmt.Fprintf(r, "Deleted: %v, ", obj.Deleted)
	}
	fmt.Fprintf(r, "Updated: %v)\n", obj.Updated)
}

// // statFile reads the stats of the named file in Google Cloud Storage.
// func (this *UploadController) statFile(fileName string) {
// 	d := this.Ctx
// 	io.WriteString(d.w, "\nFile stat:\n")

// 	obj, err := storage.StatObject(d.ctx, d.bucket, fileName)
// 	if err != nil {
// 		d.errorf("statFile: unable to stat file from bucket %q, file %q: %v", d.bucket, fileName, err)
// 		return
// 	}

// 	d.dumpStats(obj)
// }
