/* Weblang language implementation main package */
package web

/* EventHandler function for server and VM events */
type EventHandler func(e interface{}) error

/* Event before static files server handler initialization */
func (vm *VM) BeforeStaticFilesInit(event EventHandler) {
	vm.events["BeforeStaticFilesInit"] = event
}

func (vm *VM) loadBeforeStaticFilesInitEvent() error {
	if event, exists := vm.events["BeforeStaticFilesInit"]; exists {
		err := event(*vm.server)
		if err != nil {
			return err
		}
	}
	return nil
}
