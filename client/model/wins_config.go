/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WINS configuration
type WinsConfig struct {
	Common

	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}
