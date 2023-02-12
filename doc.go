/*
Copyright 2023 Fred78290.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package vmrestgoglient is a structured vmrest client api.
//
// The simple way to use it:
//
//	package main
//
//	import "github.com/Fred78290/vmrest-go-client/client"
//
//	func main() {
//	  cfg := client.NewConfiguration("username", "password", 8697, true)
//
//	  if client, err := client.NewAPIClient(cfg); err == nil {
//	    vms, _ := client.GetAllVMs()
//	  }
//	}
package vmrestgoglient

import (
	_ "github.com/Fred78290/vmrest-go-client/client"
	_ "github.com/Fred78290/vmrest-go-client/client/api"
	_ "github.com/Fred78290/vmrest-go-client/client/context"
	_ "github.com/Fred78290/vmrest-go-client/client/model"
)
