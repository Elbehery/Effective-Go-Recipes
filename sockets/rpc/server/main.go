// Go server for RPC
// go:build ignore
package main

import (
    "encoding/json"
    "errors"
    "flag"
    "io"
    "log"
    "net"
    "os"
)

func main() {
    var socketFile string
    flag.StringVar(&socketFile, "socket", "/tmp/server.sock", "socket file name")
    flag.Parse()

    log.SetPrefix("server: ")

    _, err := os.Stat(socketFile)
    if err != nil && !errors.Is(err, os.ErrNotExist) {
        log.Fatalf("error: %s", err)
    }

    if err == nil {
        log.Printf("deleting %q", socketFile)
        if err := os.Remove(socketFile); err != nil {
            log.Fatalf("error: %s", err)
        }
    }

    ln, err := net.Listen("unix", socketFile)
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer ln.Close()

    log.Printf("server running on %q", socketFile)

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatalf("error: %s", err)
        }
        go handler(conn)
    }
}

type Message struct {
    Text string `json:"text"`
}

func handler(conn net.Conn) {
    defer conn.Close()

    dec := json.NewDecoder(conn)
    enc := json.NewEncoder(conn)
    for {
        var msg Message
        err := dec.Decode(&msg)

        if err == io.EOF {
            return
        }
        if err != nil {
            log.Printf("error: %s", err)
            return
        }

        rev := reverse(msg.Text)
        reply := map[string]string{
            "output": rev,
        }

        if err := enc.Encode(reply); err != nil {
            log.Printf("error: can't reply - %s", err)
        }
    }
}

func reverse(s string) string {
    r := []rune(s)
    size := len(r)
    for i := 0; i < size/2; i++ {
        r[i], r[size-i-1] = r[size-i-1], r[i]
    }
    return string(r)
}

/*
import json
import logging
from os import path, remove
from socketserver import StreamRequestHandler

handler = None  # User handler function


def load_handler(file_name, handler_name):
    """Load handler function from a file"""
    with open(file_name) as fp:
        code = fp.read()
    env = {}
    exec(code, None, env)
    fn = env[handler_name]
    if fn is None:
        raise ValueError('no {handler_name!r} in {file_name!r}')

    if not callable(fn):
        raise TypeError('{file_name}:{handler_name} if not callable')

    return fn


class Handler(StreamRequestHandler):
    def reply(self, obj):
        # FIXME: Check for non JSON-serializable objects
        data = json.dumps(obj)
        data += '\n'
        self.wfile.write(data.encode())
        self.wfile.flush()

    def handle(self):
        while True:
            line = self.rfile.readline()
            try:
                obj = json.loads(line)
            except ValueError as err:
                logging.error('JSON error: %s', err)
                self.reply({'error': str(err)})
                continue

            try:
                out = handler(obj)
            except Exception as err:
                logging.error('handler error: %s', err)
                self.reply({'error': str(err)})
                continue

            logging.info('handler output: %r', out)
            self.reply({'output': out})


def main():
    global handler

    from argparse import ArgumentParser
    from socketserver import UnixStreamServer

    parser = ArgumentParser(description=__doc__)
    parser.add_argument('--socket', help='unix socket file name')
    parser.add_argument('--file', help='path to user code')
    parser.add_argument('--handler', help='name of handler function')

    args = parser.parse_args()

    logging.basicConfig(
        format='%(asctime)s - %(levelname)s - %(message)s',
        level=logging.INFO,
    )

    try:
        handler = load_handler(args.file, args.handler)
    except Exception as err:
        raise SystemExit(f'error: {err}')
    logging.info('%s:%s loaded', args.file, args.handler)

    if path.exists(args.socket):
        logging.info('deleting stale socket: %r', args.socket)
        remove(args.socket)
    server = UnixStreamServer(args.socket, Handler)
    logging.info('server running on %r', args.socket)
    server.serve_forever()


if __name__ == '__main__':
    main()
*/
