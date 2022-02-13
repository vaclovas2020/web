/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

const serverRegExpStart string = "^(server)\\s+" + regExpClassName + "\\s*[{]\\s*"
const serverRegExpParamName string = "(router|port|host)"
const serverRegExpParamValueStart string = "[(]\\s*[\"]*"
const serverRegExpParamValue string = "(\\w|[.])+"
const serverRegExpParamValueEnd string = "[\"]*\\s*[)]\\s+"
const serverRegExpOneParam string = "([@]" + serverRegExpParamName + serverRegExpParamValueStart + serverRegExpParamValue + serverRegExpParamValueEnd + ")"
const serverRegExpParams string = serverRegExpOneParam + "{3}"
const serverRegExpEnd string = "[}]\\s*$"
const serverRegExpFull string = serverRegExpStart + serverRegExpParams + serverRegExpEnd
