#include <complex.h>
#include "_cgo_export.h"
#include "liquid.h"

int frame_sync_callback_helper(
    unsigned char *  _header,
    int              _header_valid,
    unsigned char *  _payload,
    unsigned int     _payload_len,
    int              _payload_valid,
    framesyncstats_s _stats,
    void *           _userdata
) {
    return frameSyncCb(_header, _header_valid, _payload, _payload_valid, _payload_len, _stats, _userdata);
}

liquid_float_complex make_complex_float(float a, float b) {
    return a + _Complex_I * b;
}
