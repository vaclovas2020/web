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
	MaxAllowedCount   int    // Maximum allowed access count to resource. 0 - means take action immediately
	CountInterval     int64  // Time interval seconds until MaxAllowedCount would be valid and should increment value by one
	Action            uint8  // block action type
	ActionTimePeriod  int64  // time period (seconds) until action is active after web attack incident occurs. -1 means forever
}
