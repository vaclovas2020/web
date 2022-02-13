/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package instructionset

/* Weblang Instruction type: printLog(format, args) built-in function to print log message. Works with all class types */
const InstructionSetType_PrintLog uint16 = 0x0001

/* Weblang Instruction type: declare new variable in current scope */
const InstructionSetType_DeclareVar uint16 = 0x0002

/* Weblang Instruction type: update variable value only if reachable in current scope */
const InstructionSetType_UpdateVarValue uint16 = 0x0003

/* WebLang Instruction type: if sentence main code block */
const InstructionSetType_IfMain uint16 = 0x0004

/* WebLang Instruction type: if sentence else code block */
const InstructionSetType_IfElse uint16 = 0x0005

/* Weblang Instruction type: loadView(model) built-in function to load html template and parse with html/template package. Works only with controller class type */
const InstructionSetType_LoadViewTemplate uint16 = 0x0006

/* Weblang Instruction type: getCurrentRequestUrl() buil-in function to get current request url. Works only in controller class */
const InstructionSetType_GetCurrentRequestUrl uint16 = 0x0007

/* Weblang Instruction type: getCurrentRequestParam(name) buil-in function to get current request parameter by name. Works only in controller class */
const InstructionSetType_GetCurrentRequestParam uint16 = 0x0008

/* Weblang Instruction type: call class method */
const InstructionSetType_CallClassMethod uint16 = 0x0009

/* Weblang Instruction type: get binded model class. Works only with repository and controller class types */
const InstructionSetType_GetModel uint16 = 0x000a

/* Weblang Instruction type: set binded model class. Works only with repository class type */
const InstructionSetType_SetModel uint16 = 0x000b

/* Weblang Instruction type: return data model. Works only with repository class type */
const InstructionSetType_ReturnDataModel uint16 = 0x000c

// This file is not finished. There would be more instruction set types in the future
// Copyright(c) 2022 Vaclovas Lapinskis. All rights reserved. License: BSD 3-Clause License
