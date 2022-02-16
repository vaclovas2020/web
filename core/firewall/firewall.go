/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang firewall package. Web defense mechanism implementation */
package firewall

/* Web application firewall struct */
type Firewall struct {
	Rules []FirewallRule // Active Firewall rules for web application
}

/* FirewallRule to take action against dangerous web attacks base of RequestUrlPattern and HttpMethod */
type FirewallRule struct {
	RequestUrlPattern string // request pattern to block
	HttpMethod        string // Http method: GET, POST and etc.
	MinAccessCount    int    // Minimum allowed access count to resource. 0 - means take action immediately
	Action            uint8  // block action type
}
