
[![Build Status](https://travis-ci.org/itchio/goff.svg?branch=master)](https://travis-ci.org/itchio/goff)
[![Go Report Card](https://goreportcard.com/badge/github.com/itchio/goff)](https://goreportcard.com/report/github.com/itchio/butler)
![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)

# goff

goff binds a subset of the FFmpeg libraries (avformat, avcodec, avutil, swscale, etc.)

### Mission statement

At the time of this writing, there's a dozen takes on Go/FFmpeg, most of
which are forks of some other projects.

After careful examination, it appears worthwhile to contribute a fresh take
to the ecosystem that emphasizes correctness, forgoes deprecated functions
and types, tries to stick to a friendly naming convention, and does not
attempt to find universal high-level wrappers (besides type aliases, see
below).

### Contributing

If the mission statement above resonates with you, and you would like to add
missing types/constants/functions to goff, feel free to submit small pull
requests, that add only a handful of related type/functions at a time.

Try to write contributions in the style of the original codebase - if you
have questions, open a GitHub issue to get an answer.

### Prerequisites

goff is currently based on FFmpeg 4.1.

The `vendor-all.sh` script is provided to download, extract, configure 
(with an opinionated set of build flags), and install ffmpeg and x264
into `vendor_c`.

That script is enough to pass the integration tests, but it might be missing
features you might need :)

After that, `PKG_CONFIG_PATH` still needs to be set properly, which
`source vendor-env.sh` achieves (you could put that in your `.envrc`,
using <https://direnv.net/> is a good idea here).

To run binaries compiled against the libs installed by `vendor-all.sh`,
`LD_LIBRARY_PATH` needs to be set. Again, `source vendor-env.sh` takes care
of that.

### Design: Packages

Technically, avformat, avcodec, avutil are separate libraries - they
have separate headers and separate binaries.

However, due to how cgo works, and because some types are shared
across all those libraries, goff exports only a single package.

### Design: Types

Go types are used to prevent misuse whenever possible.

For example, enums aren't `int`, but `SampleFormat`.

Type aliases are used to define functions on them - for example,
`SampleFormat` implements `String()`.

Whenever possible, functions are written as methods on a type. For example,
the `CodecID` has a `FindDecoder()` method, whereas

Getters and setters are made available for types based on FFmpeg C structs,
see `frame.Width()` and `frame.SetWidth()`, for example. Hopefully those
get inlined by the go compiler!

### Design: Comments

Whenever a Doxygen comment is available for a type, constant, or function,
it is copy/pasted almost as-is to Godoc.

Functions/type names mentioned in comments are not rewritten, and are in C style.

### Design: Timings (pts/dts)

Timings (presentation timestamps, durations) in FFmpeg are expressed
as int64, in relation to a fractional timebase.

In goff, timebases are of type `Rational`, and `Timing` values can
be converted to an (approximate) `time.Duration`. Which is useful
for human display purposes, but not for internal computation.

### Design: Frame (Data and Linesize)

In FFmpeg, video frames may have as many as eight planes. RGBA video
frames typically have only one plane, with interleaved Red, Green, Blue, and Alpha value. Whereas planar formats, like YUV420P, have three planes: one for Y, one for U, one for V.

Technically, the type of `Frame.data` is `[8]*uint8`. However, this
is incredibly hard to use and surprisingly easy to mess up. So, goff
ships with the types `Planes` and `Plane`. These are used in the bindings
for swscale functions as well, preventing incorrect usage.

Additionally, `Frame` comes with a `PlaneData(i int)` getter, which directly
returns a slice (`[]uint8`) of the right capacity & length. This is the most
practical way to access raw pixel data.

### Design: Memory management

Users of goff should take care of allocating, initializing, and deallocating everything
properly. No finalizers are set, so, if you don't free something, you're leaking memory!

### License

goff is MIT-licensed, but FFmpeg has LGPL and GPL components.

See <https://ffmpeg.org/legal.html> for details.
