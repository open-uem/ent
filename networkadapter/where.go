// Code generated by ent, DO NOT EDIT.

package networkadapter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-uem/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldName, v))
}

// MACAddress applies equality check predicate on the "mac_address" field. It's identical to MACAddressEQ.
func MACAddress(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldMACAddress, v))
}

// Addresses applies equality check predicate on the "addresses" field. It's identical to AddressesEQ.
func Addresses(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldAddresses, v))
}

// Subnet applies equality check predicate on the "subnet" field. It's identical to SubnetEQ.
func Subnet(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldSubnet, v))
}

// DefaultGateway applies equality check predicate on the "default_gateway" field. It's identical to DefaultGatewayEQ.
func DefaultGateway(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDefaultGateway, v))
}

// DNSServers applies equality check predicate on the "dns_servers" field. It's identical to DNSServersEQ.
func DNSServers(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDNSServers, v))
}

// DNSDomain applies equality check predicate on the "dns_domain" field. It's identical to DNSDomainEQ.
func DNSDomain(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDNSDomain, v))
}

// DhcpEnabled applies equality check predicate on the "dhcp_enabled" field. It's identical to DhcpEnabledEQ.
func DhcpEnabled(v bool) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDhcpEnabled, v))
}

// DhcpLeaseObtained applies equality check predicate on the "dhcp_lease_obtained" field. It's identical to DhcpLeaseObtainedEQ.
func DhcpLeaseObtained(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseExpired applies equality check predicate on the "dhcp_lease_expired" field. It's identical to DhcpLeaseExpiredEQ.
func DhcpLeaseExpired(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDhcpLeaseExpired, v))
}

// Speed applies equality check predicate on the "speed" field. It's identical to SpeedEQ.
func Speed(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldSpeed, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldName, v))
}

// MACAddressEQ applies the EQ predicate on the "mac_address" field.
func MACAddressEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldMACAddress, v))
}

// MACAddressNEQ applies the NEQ predicate on the "mac_address" field.
func MACAddressNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldMACAddress, v))
}

// MACAddressIn applies the In predicate on the "mac_address" field.
func MACAddressIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldMACAddress, vs...))
}

// MACAddressNotIn applies the NotIn predicate on the "mac_address" field.
func MACAddressNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldMACAddress, vs...))
}

// MACAddressGT applies the GT predicate on the "mac_address" field.
func MACAddressGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldMACAddress, v))
}

// MACAddressGTE applies the GTE predicate on the "mac_address" field.
func MACAddressGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldMACAddress, v))
}

// MACAddressLT applies the LT predicate on the "mac_address" field.
func MACAddressLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldMACAddress, v))
}

// MACAddressLTE applies the LTE predicate on the "mac_address" field.
func MACAddressLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldMACAddress, v))
}

// MACAddressContains applies the Contains predicate on the "mac_address" field.
func MACAddressContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldMACAddress, v))
}

// MACAddressHasPrefix applies the HasPrefix predicate on the "mac_address" field.
func MACAddressHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldMACAddress, v))
}

// MACAddressHasSuffix applies the HasSuffix predicate on the "mac_address" field.
func MACAddressHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldMACAddress, v))
}

// MACAddressEqualFold applies the EqualFold predicate on the "mac_address" field.
func MACAddressEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldMACAddress, v))
}

// MACAddressContainsFold applies the ContainsFold predicate on the "mac_address" field.
func MACAddressContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldMACAddress, v))
}

// AddressesEQ applies the EQ predicate on the "addresses" field.
func AddressesEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldAddresses, v))
}

// AddressesNEQ applies the NEQ predicate on the "addresses" field.
func AddressesNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldAddresses, v))
}

// AddressesIn applies the In predicate on the "addresses" field.
func AddressesIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldAddresses, vs...))
}

// AddressesNotIn applies the NotIn predicate on the "addresses" field.
func AddressesNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldAddresses, vs...))
}

// AddressesGT applies the GT predicate on the "addresses" field.
func AddressesGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldAddresses, v))
}

// AddressesGTE applies the GTE predicate on the "addresses" field.
func AddressesGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldAddresses, v))
}

// AddressesLT applies the LT predicate on the "addresses" field.
func AddressesLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldAddresses, v))
}

// AddressesLTE applies the LTE predicate on the "addresses" field.
func AddressesLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldAddresses, v))
}

// AddressesContains applies the Contains predicate on the "addresses" field.
func AddressesContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldAddresses, v))
}

// AddressesHasPrefix applies the HasPrefix predicate on the "addresses" field.
func AddressesHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldAddresses, v))
}

// AddressesHasSuffix applies the HasSuffix predicate on the "addresses" field.
func AddressesHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldAddresses, v))
}

// AddressesEqualFold applies the EqualFold predicate on the "addresses" field.
func AddressesEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldAddresses, v))
}

// AddressesContainsFold applies the ContainsFold predicate on the "addresses" field.
func AddressesContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldAddresses, v))
}

// SubnetEQ applies the EQ predicate on the "subnet" field.
func SubnetEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldSubnet, v))
}

// SubnetNEQ applies the NEQ predicate on the "subnet" field.
func SubnetNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldSubnet, v))
}

// SubnetIn applies the In predicate on the "subnet" field.
func SubnetIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldSubnet, vs...))
}

// SubnetNotIn applies the NotIn predicate on the "subnet" field.
func SubnetNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldSubnet, vs...))
}

// SubnetGT applies the GT predicate on the "subnet" field.
func SubnetGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldSubnet, v))
}

// SubnetGTE applies the GTE predicate on the "subnet" field.
func SubnetGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldSubnet, v))
}

// SubnetLT applies the LT predicate on the "subnet" field.
func SubnetLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldSubnet, v))
}

// SubnetLTE applies the LTE predicate on the "subnet" field.
func SubnetLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldSubnet, v))
}

// SubnetContains applies the Contains predicate on the "subnet" field.
func SubnetContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldSubnet, v))
}

// SubnetHasPrefix applies the HasPrefix predicate on the "subnet" field.
func SubnetHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldSubnet, v))
}

// SubnetHasSuffix applies the HasSuffix predicate on the "subnet" field.
func SubnetHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldSubnet, v))
}

// SubnetIsNil applies the IsNil predicate on the "subnet" field.
func SubnetIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldSubnet))
}

// SubnetNotNil applies the NotNil predicate on the "subnet" field.
func SubnetNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldSubnet))
}

// SubnetEqualFold applies the EqualFold predicate on the "subnet" field.
func SubnetEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldSubnet, v))
}

// SubnetContainsFold applies the ContainsFold predicate on the "subnet" field.
func SubnetContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldSubnet, v))
}

// DefaultGatewayEQ applies the EQ predicate on the "default_gateway" field.
func DefaultGatewayEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDefaultGateway, v))
}

// DefaultGatewayNEQ applies the NEQ predicate on the "default_gateway" field.
func DefaultGatewayNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldDefaultGateway, v))
}

// DefaultGatewayIn applies the In predicate on the "default_gateway" field.
func DefaultGatewayIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldDefaultGateway, vs...))
}

// DefaultGatewayNotIn applies the NotIn predicate on the "default_gateway" field.
func DefaultGatewayNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldDefaultGateway, vs...))
}

// DefaultGatewayGT applies the GT predicate on the "default_gateway" field.
func DefaultGatewayGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldDefaultGateway, v))
}

// DefaultGatewayGTE applies the GTE predicate on the "default_gateway" field.
func DefaultGatewayGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldDefaultGateway, v))
}

// DefaultGatewayLT applies the LT predicate on the "default_gateway" field.
func DefaultGatewayLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldDefaultGateway, v))
}

// DefaultGatewayLTE applies the LTE predicate on the "default_gateway" field.
func DefaultGatewayLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldDefaultGateway, v))
}

// DefaultGatewayContains applies the Contains predicate on the "default_gateway" field.
func DefaultGatewayContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldDefaultGateway, v))
}

// DefaultGatewayHasPrefix applies the HasPrefix predicate on the "default_gateway" field.
func DefaultGatewayHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldDefaultGateway, v))
}

// DefaultGatewayHasSuffix applies the HasSuffix predicate on the "default_gateway" field.
func DefaultGatewayHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldDefaultGateway, v))
}

// DefaultGatewayIsNil applies the IsNil predicate on the "default_gateway" field.
func DefaultGatewayIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldDefaultGateway))
}

// DefaultGatewayNotNil applies the NotNil predicate on the "default_gateway" field.
func DefaultGatewayNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldDefaultGateway))
}

// DefaultGatewayEqualFold applies the EqualFold predicate on the "default_gateway" field.
func DefaultGatewayEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldDefaultGateway, v))
}

// DefaultGatewayContainsFold applies the ContainsFold predicate on the "default_gateway" field.
func DefaultGatewayContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldDefaultGateway, v))
}

// DNSServersEQ applies the EQ predicate on the "dns_servers" field.
func DNSServersEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDNSServers, v))
}

// DNSServersNEQ applies the NEQ predicate on the "dns_servers" field.
func DNSServersNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldDNSServers, v))
}

// DNSServersIn applies the In predicate on the "dns_servers" field.
func DNSServersIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldDNSServers, vs...))
}

// DNSServersNotIn applies the NotIn predicate on the "dns_servers" field.
func DNSServersNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldDNSServers, vs...))
}

// DNSServersGT applies the GT predicate on the "dns_servers" field.
func DNSServersGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldDNSServers, v))
}

// DNSServersGTE applies the GTE predicate on the "dns_servers" field.
func DNSServersGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldDNSServers, v))
}

// DNSServersLT applies the LT predicate on the "dns_servers" field.
func DNSServersLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldDNSServers, v))
}

// DNSServersLTE applies the LTE predicate on the "dns_servers" field.
func DNSServersLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldDNSServers, v))
}

// DNSServersContains applies the Contains predicate on the "dns_servers" field.
func DNSServersContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldDNSServers, v))
}

// DNSServersHasPrefix applies the HasPrefix predicate on the "dns_servers" field.
func DNSServersHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldDNSServers, v))
}

// DNSServersHasSuffix applies the HasSuffix predicate on the "dns_servers" field.
func DNSServersHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldDNSServers, v))
}

// DNSServersIsNil applies the IsNil predicate on the "dns_servers" field.
func DNSServersIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldDNSServers))
}

// DNSServersNotNil applies the NotNil predicate on the "dns_servers" field.
func DNSServersNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldDNSServers))
}

// DNSServersEqualFold applies the EqualFold predicate on the "dns_servers" field.
func DNSServersEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldDNSServers, v))
}

// DNSServersContainsFold applies the ContainsFold predicate on the "dns_servers" field.
func DNSServersContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldDNSServers, v))
}

// DNSDomainEQ applies the EQ predicate on the "dns_domain" field.
func DNSDomainEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDNSDomain, v))
}

// DNSDomainNEQ applies the NEQ predicate on the "dns_domain" field.
func DNSDomainNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldDNSDomain, v))
}

// DNSDomainIn applies the In predicate on the "dns_domain" field.
func DNSDomainIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldDNSDomain, vs...))
}

// DNSDomainNotIn applies the NotIn predicate on the "dns_domain" field.
func DNSDomainNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldDNSDomain, vs...))
}

// DNSDomainGT applies the GT predicate on the "dns_domain" field.
func DNSDomainGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldDNSDomain, v))
}

// DNSDomainGTE applies the GTE predicate on the "dns_domain" field.
func DNSDomainGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldDNSDomain, v))
}

// DNSDomainLT applies the LT predicate on the "dns_domain" field.
func DNSDomainLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldDNSDomain, v))
}

// DNSDomainLTE applies the LTE predicate on the "dns_domain" field.
func DNSDomainLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldDNSDomain, v))
}

// DNSDomainContains applies the Contains predicate on the "dns_domain" field.
func DNSDomainContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldDNSDomain, v))
}

// DNSDomainHasPrefix applies the HasPrefix predicate on the "dns_domain" field.
func DNSDomainHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldDNSDomain, v))
}

// DNSDomainHasSuffix applies the HasSuffix predicate on the "dns_domain" field.
func DNSDomainHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldDNSDomain, v))
}

// DNSDomainIsNil applies the IsNil predicate on the "dns_domain" field.
func DNSDomainIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldDNSDomain))
}

// DNSDomainNotNil applies the NotNil predicate on the "dns_domain" field.
func DNSDomainNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldDNSDomain))
}

// DNSDomainEqualFold applies the EqualFold predicate on the "dns_domain" field.
func DNSDomainEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldDNSDomain, v))
}

// DNSDomainContainsFold applies the ContainsFold predicate on the "dns_domain" field.
func DNSDomainContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldDNSDomain, v))
}

// DhcpEnabledEQ applies the EQ predicate on the "dhcp_enabled" field.
func DhcpEnabledEQ(v bool) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDhcpEnabled, v))
}

// DhcpEnabledNEQ applies the NEQ predicate on the "dhcp_enabled" field.
func DhcpEnabledNEQ(v bool) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldDhcpEnabled, v))
}

// DhcpEnabledIsNil applies the IsNil predicate on the "dhcp_enabled" field.
func DhcpEnabledIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldDhcpEnabled))
}

// DhcpEnabledNotNil applies the NotNil predicate on the "dhcp_enabled" field.
func DhcpEnabledNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldDhcpEnabled))
}

// DhcpLeaseObtainedEQ applies the EQ predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedEQ(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseObtainedNEQ applies the NEQ predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedNEQ(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseObtainedIn applies the In predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedIn(vs ...time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldDhcpLeaseObtained, vs...))
}

// DhcpLeaseObtainedNotIn applies the NotIn predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedNotIn(vs ...time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldDhcpLeaseObtained, vs...))
}

// DhcpLeaseObtainedGT applies the GT predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedGT(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseObtainedGTE applies the GTE predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedGTE(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseObtainedLT applies the LT predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedLT(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseObtainedLTE applies the LTE predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedLTE(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldDhcpLeaseObtained, v))
}

// DhcpLeaseObtainedIsNil applies the IsNil predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldDhcpLeaseObtained))
}

// DhcpLeaseObtainedNotNil applies the NotNil predicate on the "dhcp_lease_obtained" field.
func DhcpLeaseObtainedNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldDhcpLeaseObtained))
}

// DhcpLeaseExpiredEQ applies the EQ predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredEQ(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldDhcpLeaseExpired, v))
}

// DhcpLeaseExpiredNEQ applies the NEQ predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredNEQ(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldDhcpLeaseExpired, v))
}

// DhcpLeaseExpiredIn applies the In predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredIn(vs ...time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldDhcpLeaseExpired, vs...))
}

// DhcpLeaseExpiredNotIn applies the NotIn predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredNotIn(vs ...time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldDhcpLeaseExpired, vs...))
}

// DhcpLeaseExpiredGT applies the GT predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredGT(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldDhcpLeaseExpired, v))
}

// DhcpLeaseExpiredGTE applies the GTE predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredGTE(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldDhcpLeaseExpired, v))
}

// DhcpLeaseExpiredLT applies the LT predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredLT(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldDhcpLeaseExpired, v))
}

// DhcpLeaseExpiredLTE applies the LTE predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredLTE(v time.Time) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldDhcpLeaseExpired, v))
}

// DhcpLeaseExpiredIsNil applies the IsNil predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredIsNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIsNull(FieldDhcpLeaseExpired))
}

// DhcpLeaseExpiredNotNil applies the NotNil predicate on the "dhcp_lease_expired" field.
func DhcpLeaseExpiredNotNil() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotNull(FieldDhcpLeaseExpired))
}

// SpeedEQ applies the EQ predicate on the "speed" field.
func SpeedEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEQ(FieldSpeed, v))
}

// SpeedNEQ applies the NEQ predicate on the "speed" field.
func SpeedNEQ(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNEQ(FieldSpeed, v))
}

// SpeedIn applies the In predicate on the "speed" field.
func SpeedIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldIn(FieldSpeed, vs...))
}

// SpeedNotIn applies the NotIn predicate on the "speed" field.
func SpeedNotIn(vs ...string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldNotIn(FieldSpeed, vs...))
}

// SpeedGT applies the GT predicate on the "speed" field.
func SpeedGT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGT(FieldSpeed, v))
}

// SpeedGTE applies the GTE predicate on the "speed" field.
func SpeedGTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldGTE(FieldSpeed, v))
}

// SpeedLT applies the LT predicate on the "speed" field.
func SpeedLT(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLT(FieldSpeed, v))
}

// SpeedLTE applies the LTE predicate on the "speed" field.
func SpeedLTE(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldLTE(FieldSpeed, v))
}

// SpeedContains applies the Contains predicate on the "speed" field.
func SpeedContains(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContains(FieldSpeed, v))
}

// SpeedHasPrefix applies the HasPrefix predicate on the "speed" field.
func SpeedHasPrefix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasPrefix(FieldSpeed, v))
}

// SpeedHasSuffix applies the HasSuffix predicate on the "speed" field.
func SpeedHasSuffix(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldHasSuffix(FieldSpeed, v))
}

// SpeedEqualFold applies the EqualFold predicate on the "speed" field.
func SpeedEqualFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldEqualFold(FieldSpeed, v))
}

// SpeedContainsFold applies the ContainsFold predicate on the "speed" field.
func SpeedContainsFold(v string) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.FieldContainsFold(FieldSpeed, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.NetworkAdapter {
	return predicate.NetworkAdapter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Agent) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NetworkAdapter) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NetworkAdapter) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NetworkAdapter) predicate.NetworkAdapter {
	return predicate.NetworkAdapter(sql.NotPredicates(p))
}
