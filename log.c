
#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <stdint.h>
#include <string.h>
#include <libavutil/avutil.h>

#include "_cgo_export.h"

void goff_log_callback(void *ptr, int level, const char *fmt, va_list vl) {
  char line[1024];
  static int print_prefix = 1;
  av_log_format_line(ptr, level, fmt, vl, line, 1024, &print_prefix);
  goff_send_log_to_go(line);
}
