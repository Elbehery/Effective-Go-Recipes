package main

import (
    "fmt"
    "unicode/utf8"
)


/*
#include <sys/ioctl.h>
#include <stdio.h>
#include <unistd.h>

int term_width() {
    struct winsize w;
    ioctl(STDOUT_FILENO, TIOCGWINSZ, &w);
    return w.ws_col;
}
*/
import "C"


func Center(text string) string {
    width := int(C.term_width()) // Convert C.int to Go int

    textLen := utf8.RuneCountInString(text)
    lpad := (width - textLen) / 2 // left padding
    if lpad <= 0 {                // text >= screen width
        return text
    }
    rpad := width - len(text) - lpad // right padding

    return fmt.Sprintf("%*s%s%*s", lpad, " ", text, rpad, " ")
}

func main() {
    fmt.Println(Center("Hello Gophers!"))
}

