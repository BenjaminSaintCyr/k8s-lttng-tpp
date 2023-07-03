#undef TRACEPOINT_PROVIDER
#define TRACEPOINT_PROVIDER k8s_ust

#undef TRACEPOINT_INCLUDE
#define TRACEPOINT_INCLUDE "./k8s-tp.h"

#if !defined(_TP_H) || defined(TRACEPOINT_HEADER_MULTI_READ)
#define _TP_H

#include <lttng/tracepoint.h>

TRACEPOINT_EVENT(
    k8s_ust,
    event,

    /* Input arguments */
    TP_ARGS(
        const char*, o_name,
        int, name_length,
        const char*, o_ctx,
        int, ctx_length
    ),

    /* Output event fields */
    TP_FIELDS(
        ctf_sequence_text(char, op_name, o_name, unsigned int, name_length)
        ctf_sequence_text(char, op_ctx, o_ctx, unsigned int, ctx_length)
    )
)

#endif /* _TP_H */

#include <lttng/tracepoint-event.h>
