/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

/* class name regexp */
const regExpClassName string = "[[:alpha:]]\\w*"
const regExpNamespaceName string = "(" + regExpClassName + "[\\\\]{0,1})+"
const regExpNamespaceStart string = "^\\s*(namespace)\\s+"
const regExpNamespace string = regExpNamespaceStart + regExpClassName

const regExpComments string = "(\\/\\*([^*]|[\r\n]|(\\*+([^*\\/]|[\r\n])))*\\*+\\/)|(\\/\\/.*)"
