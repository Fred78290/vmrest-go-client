/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VmRegistrationInformation struct {
	// Registered VM name id
	Id string `json:"id,omitempty"`
	// Registered VM path
	Path string `json:"path,omitempty"`
}