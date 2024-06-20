package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
    "strconv"
)


// Calc is a calculator.
type Calc struct {
    p *os.Process   // Underlying process
    w io.Writer     // Process stdin
    r *bufio.Reader // Process stdout & stderr
}



// NewCalc creates a new calculator
func NewCalc() (*Calc, error) {
    // -l adds mathlib, -q mean no banner
    cmd := exec.Command("bc", "-lq")

    w, err := cmd.StdinPipe()
    if err != nil {
        return nil, err
    }

    r, err := cmd.StdoutPipe()
    if err != nil {
        return nil, err
    }
    cmd.Stderr = cmd.Stdout // redirect standard error to standard output

    if err := cmd.Start(); err != nil {
        return nil, err
    }

    c := &Calc{
        p: cmd.Process,
        w: w,
        r: bufio.NewReader(r),
    }
    return c, nil
}


// Eval evaluates a math expression such as "3 / 7".
func (c *Calc) Eval(expr string) (float64, error) {
    if _, err := fmt.Fprintf(c.w, "%s\n", expr); err != nil {
        return 0, err
    }

    line, err := c.r.ReadString('\n')
    if err != nil {
        return 0, err
    }

    line = line[:len(line)-1] // trim \n

    val, err := strconv.ParseFloat(line, 64)
    if err != nil {
        return 0, err
    }

    return val, nil
}


// Close closes the calculator.
func (c *Calc) Close() error {
    return c.p.Kill()
}


func main() {
    c, err := NewCalc()
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer c.Close()

    exprs := []string{
        "7 / 3",
        "sqrt(2)",
        "3/0", // error
        "-4/77",
    }


    for _, expr := range exprs {
        out, err := c.Eval(expr)
        if err != nil {
            fmt.Println(expr, "error:", err)
            continue
        }
        fmt.Println(expr, "→", out)
    }

}

