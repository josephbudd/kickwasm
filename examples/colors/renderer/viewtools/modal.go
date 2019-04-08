package viewtools

import (
	"fmt"
	"syscall/js"
)

/*
	WARNING:

	DO NOT EDIT THIS FILE.

*/

type modalViewData struct {
	title   string
	message string
	cb      func()
}

// GoModal adds a title and message and call back to the modalQueue.
func (tools *Tools) GoModal(message, title string, callback func()) {
	tools.queueModal(
		&modalViewData{
			title:   title,
			message: fmt.Sprintf("<p>%s</p>", message),
			cb:      callback,
		},
	)
}

// GoModalHTML adds a title and html message and call back to the modalQueue.
// Param message is html.
// Param title is plain text.
func (tools *Tools) GoModalHTML(htmlMessage, title string, callback func()) {
	tools.queueModal(
		&modalViewData{
			title:   title,
			message: htmlMessage,
			cb:      callback,
		},
	)
}

func (tools *Tools) beModal() {
	wasModal := tools.beingModal
	m := tools.unQueueModal()
	if tools.beingModal = m != nil; !tools.beingModal {
		return
	}
	notJS := tools.notJS
	notJS.SetInnerText(tools.modalMasterViewH1, m.title)
	notJS.SetInnerHTML(tools.modalMasterViewMessage, m.message)
	tools.modalCallBack = m.cb
	tools.ElementShow(tools.modalMasterView)
	if !wasModal {
		tools.SizeApp()
	}
}

func (tools *Tools) beNotModal() {
	if tools.modalQueueLastIndex >= 0 {
		tools.beModal()
		return
	}
	tools.ElementHide(tools.modalMasterView)
	tools.SizeApp()
	tools.beingModal = false
}

func (tools *Tools) queueModal(m *modalViewData) {
	if tools.modalQueueLastIndex < 4 {
		tools.modalQueueLastIndex++
		tools.modalQueue[tools.modalQueueLastIndex] = m
	}
	if !tools.beingModal {
		tools.beModal()
	}
}

func (tools *Tools) unQueueModal() *modalViewData {
	if tools.modalQueueLastIndex < 0 {
		return nil
	}
	m := tools.modalQueue[0]
	for i := 0; i < tools.modalQueueLastIndex; i++ {
		tools.modalQueue[i] = tools.modalQueue[i+1]
	}
	tools.modalQueue[tools.modalQueueLastIndex] = nil
	tools.modalQueueLastIndex--
	return m
}

func (tools *Tools) handleModalMasterViewClose(event js.Value) interface{} {
	if tools.modalCallBack != nil {
		tools.modalCallBack()
		tools.modalCallBack = nil
	}
	tools.beNotModal()
	return nil
}
