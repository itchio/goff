
#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <stdint.h>
#include <string.h>
#include <libavutil/avutil.h>

#include "_cgo_export.h"

#define _GOFF_LOG_BUFFER_SIZE 1024

void goff_log_callback_trampoline(void *ptr, int level, const char *fmt, va_list vl) {
  if (!goff_should_send_log(level)) {
    return;
  }

  char line[_GOFF_LOG_BUFFER_SIZE];
  static int print_prefix = 1;
  vsnprintf(line, _GOFF_LOG_BUFFER_SIZE, fmt, vl);
  goff_send_log_to_go((uintptr_t)(ptr), level, line, print_prefix);
}

int goff_reader_read_packet_trampoline(void *opaque, uint8_t *buf, int buf_size) {
  return goff_reader_read_packet(opaque, buf, buf_size);
}

int64_t goff_reader_seek_trampoline(void *opaque, int64_t offset, int whence) {
  return goff_reader_seek(opaque, offset, whence);
}


