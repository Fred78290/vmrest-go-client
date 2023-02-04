/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VmRegisterParameter struct {
	// Register VM name
	Name string `json:"name,omitempty"`
	// Register VM path
	Path string `json:"path,omitempty"`
}