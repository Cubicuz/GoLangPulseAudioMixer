package gopainterface

/*
#cgo pkg-config: libpulse
#include "pulse/pulseaudio.h"

extern void PaStateCallback(pa_context *pactx);

void paStateCallbackCgo(pa_context *pactx, void *userdata){
	PaStateCallback(pactx);
}
*/
import "C"
