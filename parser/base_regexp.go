/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

/* class name regexp */
const regExpClassName string = "[[:alpha:]]\\w*"

const regExpComments string = "(\\/\\*([^*]|[\r\n]|(\\*+([^*\\/]|[\r\n])))*\\*+\\/)|(\\/\\/.*)"
