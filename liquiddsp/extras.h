#include "liquid.h"
// #include "_cgo_export.h"
// #include "cgo_helpers.h"

#pragma once

int frame_sync_callback_helper(
    unsigned char *  _header,
    int              _header_valid,
    unsigned char *  _payload,
    unsigned int     _payload_len,
    int              _payload_valid,
    framesyncstats_s _stats,
    void *           _userdata
);

liquid_float_complex make_complex_float(float a, float b);