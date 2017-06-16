package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
)

const (
	projectID = "data-importer-168420"
)

func main() {
	pkey, err := ioutil.ReadFile("pkey.pem")

	bucket := "data-importer-storage"
	filename := "csvexample - Sheet1"
	method := "PUT"
	expires := time.Now().Add(time.Second * 3000)

	url, err := storage.SignedURL(bucket, filename, &storage.SignedURLOptions{
		GoogleAccessID: accessId,
		PrivateKey:     pkey,
		Method:         method,
		Expires:        expires,
	})
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	fmt.Println("URL = " + url)
}
