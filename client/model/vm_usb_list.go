/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VmUsbList struct {
	Common

	Num        int           `json:"num,omitempty"`
	UsbDevices []VmUsbDevice `json:"usbDevices,omitempty"`
}
