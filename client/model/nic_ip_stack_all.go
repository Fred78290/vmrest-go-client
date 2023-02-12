/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Information about all NICs
type NicIpStackAll struct {
	Common

	Nics   []NicIpStack `json:"nics,omitempty"`
	Routes []RouteEntry `json:"routes,omitempty"`
	// Global DNS configuration
	Dns *DnsConfig `json:"dns,omitempty"`
	// Global WINS configuration
	Wins *WinsConfig `json:"wins,omitempty"`
	// Global DHCPv4 configuration
	Dhcpv4 *DhcpConfig `json:"dhcpv4,omitempty"`
	// Global DHCPv6 configuration
	Dhcpv6 *DhcpConfig `json:"dhcpv6,omitempty"`
}
