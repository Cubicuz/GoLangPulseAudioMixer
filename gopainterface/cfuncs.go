package gopainterface

/*
#cgo pkg-config: libpulse
#include "pulse/pulseaudio.h"

extern void paStateCallback(pa_context *pactx, void *userdata);

void paStateCallbackCgo(pa_context *pactx, void *userdata){
	paStateCallback(pactx, userdata);
}


extern void paSubscribeCallback(pa_context *pactx, pa_subscription_event_type_t t, uint32_t idx, void *userdata);

void paSubscribeCallbackCgo(pa_context *pactx, pa_subscription_event_type_t t, uint32_t idx, void *userdata){
	paSubscribeCallback(pactx, t, idx, userdata);
}

*/
import "C"
