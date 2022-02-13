/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package method

/* Method type public: publically accessible class method */
const MethodType_Public uint8 = 0x01

/* Method type private: private accessible class method */
const MethodType_Private uint8 = 0x02

/* Method type protected: protected accessible class method */
const MethodType_Protected uint8 = 0x03

/* Method type external public: external (use with vm.DefineFunc() ) and public accessible class method */
const MethodType_ExternalPublic uint8 = 0x04

/* Method type external public: external (use with vm.DefineFunc() ) and private accessible class method */
const MethodType_ExternalPrivate uint8 = 0x05

/* Method type external public: external (use with vm.DefineFunc() ) and protected accessible class method */
const MethodType_ExternalProtected uint8 = 0x06
