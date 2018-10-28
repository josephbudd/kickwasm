package caller

// MainProcesser is the behavior for the main process.
type MainProcesser interface{

	// Process is a main process func which is called to process the params after they are received from the renderer.
	Process(params []byte, callback func(params []byte))
}

// Renderer is the behavior for the renderer.
type Renderer interface{

	// AddCallBack is a renderer func which adds a renderer func to a call back dispatcher.
	// Your panel caller code will use this func to add its funcs to handle the main processes callbacks.
	AddCallBack(func(interface{}))

	// CallMainProcess is a renderer func which receives params and passes them to a call to a func in the main process.
	// Your panel caller code will use this func to send params to the main process.
	CallMainProcess(params interface{})

	// Dispatch a renderer func which is called to receive and dispatche params after they are received from the main process.
	Dispatch(params []byte)

}

