/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

/* class name regexp */
const regExpClassName string = "[[:alpha:]]\\w*"
const regExpNamespaceName string = "(" + regExpClassName + "[\\])+"
const regExpNamespace string = "^\\s*(namespace)\\s+" + regExpClassName

const regExpComments string = "(\\/\\*([^*]|[\r\n]|(\\*+([^*\\/]|[\r\n])))*\\*+\\/)|(\\/\\/.*)"
