/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VmConnectedDevice struct {
	Index            int    `json:"index,omitempty"`
	StartConnected   string `json:"startConnected,omitempty"`
	ConnectionStatus int    `json:"connectionStatus,omitempty"`
	DevicePath       string `json:"devicePath,omitempty"`
}