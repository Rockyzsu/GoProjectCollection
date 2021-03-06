// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The tokenauth command demonstrates using the oauth2.StaticTokenSource.
// You can test out a GitHub Personal Access Token using this simple example.
// You can generate them here: https://github.com/settings/tokens
package main

// 可以用
import (
	"context"
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {
	//fmt.Print("GitHub Token: ")
	//byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	//println()
	//token := string(byteToken)
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	fmt.Println(token)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v\n", resp.Rate)

	// If a Token Expiration has been set, it will be displayed.
	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
