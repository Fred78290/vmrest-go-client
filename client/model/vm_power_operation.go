/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VmPowerOperation string

// List of VMPowerOperation
const (
	VM_ON       VmPowerOperation = "on"
	VM_OFF      VmPowerOperation = "off"
	VM_SHUTDOWN VmPowerOperation = "shutdown"
	VM_SUSPEND  VmPowerOperation = "suspend"
	VM_PAUSE    VmPowerOperation = "pause"
	VM_UNPAUSE  VmPowerOperation = "unpause"
)