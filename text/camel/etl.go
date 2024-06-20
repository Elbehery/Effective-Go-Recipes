package main

import (
    "fmt"
    "regexp"
    "unicode"
)

var (
    // Example: match "rN" in "userName"
    camelRe = regexp.MustCompile(`[a-z][A-Z]`)
)


// fixCase gets a string in the format "aB" and return "a_b"
func fixCase(s string) string {
    return fmt.Sprintf("%c_%c", s[0], unicode.ToLower(rune(s[1])))
}


// camelToLower turns "camelCase" to "camel_case"
func camelToLower(s string) string {
    return camelRe.ReplaceAllStringFunc(s, fixCase)
}


// Small test
func main() {
    values := []string{
        "camelCase",
        "camel",
        "",
        "ABC",
        "aB",
        "Ba",
    }

    for _, s := range values {
        fmt.Printf("%q -> %q\n", s, camelToLower(s))
    }
}
