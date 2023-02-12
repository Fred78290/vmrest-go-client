/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VmApplianceView struct {
	Common

	Author        string `json:"author,omitempty"`
	Version       string `json:"version,omitempty"`
	Port          int    `json:"port,omitempty"`
	ShowAtPowerOn string `json:"showAtPowerOn,omitempty"`
}
