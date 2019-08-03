package helloworldtemplatepanel

/*

	Panel name: HelloWorldTemplatePanel

*/

// panelCaller communicates with the main process via an asynchrounous connection.
type panelCaller struct {
	uniqueID     uint64
	group        *panelGroup
	presenter    *panelPresenter
	controler    *panelControler
	unspawn      func() error
	UnSpawningCh chan struct{}

	/* NOTE TO DEVELOPER. Step 1 of 4.

	// 1.1: Declare your panelCaller members.

	// example:

	state uint64

	*/
}

/* NOTE TO DEVELOPER. Step 2 of 4.

// 2.1: Define your funcs which send a message to the main process.
// 2.2: Define your funcs which receive a message from the main process.

// example:

// import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc/message"
// import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/store/record"

// Add Customer.

func (caller *panelCaller) addCustomer(r *record.Customer) {
	msg := &message.AddCustomerRendererToMainProcess{
		UniqueID: caller.uniqueID,
		Record:   r,
	}
	sendCh <- msg
}

func (caller *panelCaller) addCustomerRX(msg *message.AddCustomerMainProcessToRenderer) {
	if msg.UniqueID == caller.uniqueID {
		if msg.Error {
			tools.Error(msg.ErrorMessage)
			return
		}
		// no errors
		tools.Success("Customer Added.")
	}
}

*/

// listen listens for messages from the main process.
// It stops when it receives on the eoj channel.
func (caller *panelCaller) listen() {
	go func() {
		for {
			select {
			case <-eojCh:
				return
			case <-caller.UnSpawningCh:
				return
			case msg := <-receiveCh:
				// A message sent from the main process to the renderer.
				switch msg := msg.(type) {

				/* NOTE TO DEVELOPER. Step 3 of 4.

				// 3.1:   Remove the default statement below.
				// 3.2.a: Add a case for each of the messages
				//          that you are expecting from the main process.
				// 3.2.b: In that case statement, pass the message to your message receiver func.

				// example:

				// import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc/message"

				case *message.AddCustomerMainProcessToRenderer:
					caller.addCustomerRX(msg)

				*/

				default:
					_ = msg
				}
			}
		}
	}()

	return
}

// initialCalls makes the first calls to the main process.
func (caller *panelCaller) initialCalls() {

	/* NOTE TO DEVELOPER. Step 4 of 4.

	//4.1: Make any initial calls to the main process that must be made when the app starts.

	// example:

	// import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/data/loglevels"
	// import "github.com/josephbudd/kickwasm/examples/spawntabs/domain/lpc/message"

	msg := &message.LogRendererToMainProcess{
		Level:   loglevels.LogLevelInfo,
		Message: "Started",
	}
	sendCh <- msg

	*/

}
