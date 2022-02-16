/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang firewall package. Web defense mechanism implementation */
package firewall

/* Web application firewall struct */
type Firewall struct {
	Rules []FirewallRule // Active Firewall rules for web application
}

/* FirewallRule to take action against dangerous web attacks base of RequestUrlPattern */
type FirewallRule struct {
	RequestUrlPattern string // request pattern to block
	Action            uint8  // block action type
}
