/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package firewall

/* Defense action: Block client's IP address and prevent access to web application. Good for prevent brute force attacks. */
const Action_BlockClientIp uint8 = 0x01

/*  Defense action: return Http status 404 Not found to prevent access to web resource (recommended action) */
const Action_Fake404Status uint8 = 0x02

/*  Defense action: return http status 403 to prevent access to web resource (not recommended action, because hackers often try to hack when they see that web resource exists) */
const Action_403Status uint8 = 0x03
