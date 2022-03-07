// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package demo provides an app that shows how to use the github package on
// Google App Engine.
package appengine

// 无法正常使用

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"google.golang.org/appengine/log"
)

func init() {
	fmt.Println("init")
	http.HandleFunc("/git", handler)
	http.HandleFunc("/", handler2)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	http.NotFound(w, r)
	//	return
	//}
	fmt.Println("index")
	//ctx := appengine.NewContext(r)
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_AUTH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	commits, _, err := client.Repositories.ListCommits(ctx, "google", "go-github", nil)
	if err != nil {
		log.Errorf(ctx, "ListCommits: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, commit := range commits {
		fmt.Fprintln(w, commit.GetHTMLURL())
	}
}
func handler2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("index")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	buff := bytes.Buffer{}
	buff.Write([]byte("Hello world"))
	fmt.Fprintln(w, buff.String())

}
