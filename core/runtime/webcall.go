/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* This is core package for Weblang language. It's implements runtime environment. */
package runtime

import (
	"fmt"

	"webimizer.dev/web/base"
)

/* WebCall struct for communication between diffrent WebLang VM runtime objects. Use to call receiver's function and send arguments to receiver's FunctionHandler */
type WebCall struct {
	CallerObj    *base.Object           // Caller object pointer
	ReceiverObj  *base.Object           // Receiver object pointer
	ReceiverFunc *base.Function         // Receiver's function pointer
	Args         map[string]interface{} //Arguments to send to receiver's FunctionHandler
}

/* Invoke receiver's FunctionHandler and send arguments */
func (wcall *WebCall) Call() error {
	if wcall.ReceiverFunc != nil && wcall.ReceiverFunc.Handler != nil {
		return wcall.ReceiverFunc.Handler.Invoke(wcall.Args, wcall.ReceiverFunc, wcall.CallerObj, wcall.ReceiverObj)
	}
	return fmt.Errorf("receiver function struct or receiver function handler is nil")
}
