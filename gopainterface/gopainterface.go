package gopainterface

/*

#cgo pkg-config: libpulse
#include "pulse/pulseaudio.h"
#include <stdio.h>
#include <stdlib.h>

void paStateCallbackCgo(pa_context *pactx, void *userdata);
void paSubscribeCallbackCgo(pa_context *pactx, pa_subscription_event_type_t t, uint32_t idx, void *userdata);

*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var pa_ml *C.pa_threaded_mainloop = nil
var pa_api *C.pa_mainloop_api = nil // no need to free
var pa_ctx *C.pa_context = nil

func Initialize() error {

	if pa_ml == nil {
		pa_ml = C.pa_threaded_mainloop_new()
		if pa_ml == nil {
			return errors.New("error: could not create mainloop")
		}
	}

	if pa_api == nil {
		pa_api = C.pa_threaded_mainloop_get_api(pa_ml)
	}

	retVal := C.pa_threaded_mainloop_start(pa_ml)
	if retVal < 0 {
		Deinitialize()
		return errors.New("unable to start mainloop")
	}

	return nil

}

//export paStateCallback
func paStateCallback(pactx *C.pa_context, userdata *C.void) {
	fmt.Println("wohoo paStateCallback callback")
}

//export paSubscribeCallback
func paSubscribeCallback(pactx *C.pa_context, subscriptionEventType C.pa_subscription_event_type_t, idx uint32, userdata *C.void) {
	fmt.Println("wohoo paSubscribeCallback callback")
}

func Connect() (bool, error) {
	if pa_ctx != nil {
		return false, nil
	}

	C.pa_threaded_mainloop_lock(pa_ml)

	sgopamixer := C.CString("gopamixer")
	defer C.free(unsafe.Pointer(sgopamixer))
	saudiocard := C.CString("audio-card")
	defer C.free(unsafe.Pointer(saudiocard))
	sPA_PROP_APPLICATION_NAME := C.CString(C.PA_PROP_APPLICATION_NAME)
	defer C.free(unsafe.Pointer(sPA_PROP_APPLICATION_NAME))
	sPA_PROP_APPLICATION_ID := C.CString(C.PA_PROP_APPLICATION_ID)
	defer C.free(unsafe.Pointer(sPA_PROP_APPLICATION_ID))
	sPA_PROP_APPLICATION_ICON_NAME := C.CString(C.PA_PROP_APPLICATION_ICON_NAME)
	defer C.free(unsafe.Pointer(sPA_PROP_APPLICATION_ICON_NAME))

	var proplist *C.pa_proplist = C.pa_proplist_new()
	C.pa_proplist_sets(proplist, sPA_PROP_APPLICATION_NAME, sgopamixer)
	C.pa_proplist_sets(proplist, sPA_PROP_APPLICATION_ID, sgopamixer)
	C.pa_proplist_sets(proplist, sPA_PROP_APPLICATION_ICON_NAME, saudiocard)
	pa_ctx = C.pa_context_new_with_proplist(pa_api, nil, proplist)
	C.pa_proplist_free(proplist)
	proplist = nil

	if pa_ctx == nil {
		C.pa_threaded_mainloop_unlock(pa_ml)
		return false, errors.New("unable to create context")
	}

	C.pa_context_set_state_callback(pa_ctx, (C.pa_context_notify_cb_t)(unsafe.Pointer(C.paStateCallbackCgo)), nil)
	C.pa_context_set_subscribe_callback(pa_ctx, (C.pa_context_subscribe_cb_t)(unsafe.Pointer(C.paSubscribeCallbackCgo)), nil)

	retval := C.pa_context_connect(pa_ctx, nil, C.PA_CONTEXT_NOFAIL, nil)
	if retval < 0 {
		fmt.Println("error connecting ", retval)
		C.pa_context_disconnect(pa_ctx)
		C.pa_context_unref(pa_ctx)
		C.pa_threaded_mainloop_unlock(pa_ml)
		pa_ctx = nil

		return false, errors.New("error connecting")
	}

	C.pa_threaded_mainloop_unlock(pa_ml)

	return true, nil
}

func Deinitialize() {
	pa_api = nil
	if pa_ml != nil {
		C.pa_threaded_mainloop_stop(pa_ml)

		C.pa_threaded_mainloop_free(pa_ml)
	}
}
