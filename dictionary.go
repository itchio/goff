package goff

//#cgo pkg-config: libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
import "C"

type Dictionary = C.struct_AVDictionary
type DictionaryEntry = C.struct_AVDictionaryEntry
type DictionaryFlags = C.int

var (
	DictionaryFlags_MATCH_CASE      DictionaryFlags = C.AV_DICT_MATCH_CASE
	DictionaryFlags_IGNORE_SUFFIX   DictionaryFlags = C.AV_DICT_IGNORE_SUFFIX
	DictionaryFlags_DONT_STRDUP_KEY DictionaryFlags = C.AV_DICT_DONT_STRDUP_KEY
	DictionaryFlags_DONT_STRDUP_VAL DictionaryFlags = C.AV_DICT_DONT_STRDUP_VAL
	DictionaryFlags_DONT_OVERWRITE  DictionaryFlags = C.AV_DICT_DONT_OVERWRITE
	DictionaryFlags_APPEND          DictionaryFlags = C.AV_DICT_APPEND
	DictionaryFlags_MULTIKEY        DictionaryFlags = C.AV_DICT_MULTIKEY
)

// Copy entries from one AVDictionary struct into another.
//
// Parameters
//   dst	pointer to a pointer to a AVDictionary struct. If *dst is NULL, this function will allocate a struct for you and put it in *dst
//   src	pointer to source AVDictionary struct
//   flags	flags to use when setting entries in *dst
//
// Note
//   metadata is read using the AV_DICT_IGNORE_SUFFIX flag
//
// Returns
//   0 on success, negative AVERROR code on failure. If dst was allocated by this function, callers should free the associated memory.
func (d *Dictionary) Copy(flags DictionaryFlags) (*Dictionary, error) {
	var dst *Dictionary
	err := CheckErr(C.av_dict_copy(&dst, d, C.int(flags)))
	return dst, err
}

func (d *Dictionary) AsMap() map[string]string {
	m := make(map[string]string)
	for _, k := range d.Keys() {
		if v, ok := d.Get(k); ok {
			m[k] = v
		}
	}
	return m
}

func (d *Dictionary) Keys() []string {
	var keys []string

	// gotcha: can't use our CString() wrapper
	// because that'll give us a nil pointer
	key_ := C.CString("")
	defer FreeString(key_)

	var entry *DictionaryEntry = nil
	for {
		entry = C.av_dict_get(d, key_, entry, DictionaryFlags_IGNORE_SUFFIX)
		if entry == nil {
			break
		}
		keys = append(keys, entry.Key())
	}
	return keys
}

func (d *Dictionary) Get(key string) (string, bool) {
	key_ := CString(key)
	defer FreeString(key_)

	var entry *DictionaryEntry
	entry = C.av_dict_get(d, key_, nil, 0)

	if entry == nil {
		return "", false
	}
	return entry.Value(), true
}

// Set the given entry in *pm, overwriting an existing entry.
//
// Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set, these
// arguments will be freed on error.
//
// Warning: Adding a new entry to a dictionary invalidates all existing entries
// previously returned with av_dict_get.
//
// Parameters
//     pm pointer to a pointer to a dictionary struct. If *pm is NULL a dictionary
//     struct is allocated and put in *pm.
//     key entry key to add to *pm (will either be av_strduped or added as a new key
//     depending on flags)
//     value entry value to add to *pm (will be av_strduped or added as a new key
//     depending on flags). Passing a NULL value will cause an existing entry to be
//     deleted.
//
// Returns
//     >= 0 on success otherwise an error code <0
func (d *Dictionary) Set(key string, value string) (*Dictionary, error) {
	key_ := CString(key)
	defer FreeString(key_)

	value_ := CString(value)
	defer FreeString(value_)

	err := CheckErr(C.av_dict_set(&d, key_, value_, 0))
	return d, err
}

// Free all the memory allocated for an AVDictionary struct and all keys and
// values.
func (d *Dictionary) Free() {
	C.av_dict_free(&d)
}

func (e *DictionaryEntry) Key() string {
	return C.GoString(e.key)
}

func (e *DictionaryEntry) Value() string {
	return C.GoString(e.value)
}
