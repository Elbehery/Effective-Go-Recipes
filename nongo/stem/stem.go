package main

import (
    "fmt"
    "log"
    "unsafe"
)


/*
#include <libstemmer.h>
#include <stdlib.h>

#cgo LDFLAGS: -lstemmer
*/
import "C"



// Stemmer stems words for a specific language.
type Stemmer struct {
    st *C.struct_sb_stemmer
}



// NewStemmer creates a new stemmer for a language.
func NewStemmer(lang string) (*Stemmer, error) {
    cLang := C.CString(lang) // Go -> C string <label id="code.stem.cstr"/>
    st := C.sb_stemmer_new(cLang, nil)
    C.free(unsafe.Pointer(cLang)) 

    if st == nil {
        return nil, fmt.Errorf("can't create stemmer for %q", lang)
    }

    return &Stemmer{st}, nil
}



// Close closes the stemmer, freeing allocated memory.
func (s *Stemmer) Close() {
    if s.st != nil {
        C.sb_stemmer_delete(s.st)
        s.st = nil
    }
}



// Stem will stem a word.
// In case of error will return the empty string
func (s *Stemmer) Stem(word string) string {

    cWord := C.CBytes([]byte(word))
    size := C.int(len(word))
    sym := C.sb_stemmer_stem(s.st, (*C.uchar)(cWord), size)
    if sym == nil {
        return ""
    }

    // Convert to Go string
    i := C.sb_stemmer_length(s.st)
    data := C.GoBytes(unsafe.Pointer(sym), i)

    // No need to free data, managed by C stemmer
    return string(data)
}

func main() {
    s, err := NewStemmer("en")
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer s.Close()

    for _, w := range []string{"works", "working", "worked"} {
        stem := s.Stem(w)
        fmt.Printf("%-8s → %s\n", w, stem)
    }
}
