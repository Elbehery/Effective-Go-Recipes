package myjanitor

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"time"
)

// gzCompress compresses src into dest with gzip.
func gzCompress(src, dest string) error {
	input, err := os.Open(src)
	if err != nil {
		return err
	}
	defer input.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	w := gzip.NewWriter(out)
	defer w.Close()

	w.Name = src
	fileStat, err := input.Stat()
	if err == nil {
		w.ModTime = fileStat.ModTime()
	}

	_, err = io.Copy(w, input)
	if err != nil {
		os.Remove(dest)
		return err
	}

	return nil
}

func shouldCompress(path string, maxAge time.Duration) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("warning: could not get file info of [%v]", path)
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return time.Since(fileInfo.ModTime()) >= maxAge
}

func filesToCompress(dir string, maxAge time.Duration) ([]string, error) {
	root := os.DirFS(dir)
	logFiles, err := fs.Glob(root, "*.log")
	if err != nil {
		return nil, err
	}

	var files []string

	for _, f := range logFiles {
		file := path.Join(dir, f)
		if shouldCompress(file, maxAge) {
			files = append(files, file)
		}
	}

	return files, nil
}

func fileSHA1(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file
	if path.Ext(name) == ".gz" {
		r, err = gzip.NewReader(r)
		if err != nil {
			return "", err
		}
	}

	w := sha1.New()
	if _, err = io.Copy(w, r); err != nil {
		return "", err
	}

	sig := fmt.Sprintf("%x", w.Sum(nil))
	return sig, nil
}

func sameSig(f1, f2 string) (bool, error) {
	sig1, err := fileSHA1(f1)
	if err != nil {
		return false, err
	}

	sig2, err := fileSHA1(f2)
	if err != nil {
		return false, err
	}

	return sig1 == sig2, nil
}

func compressFiles(dir string, maxAge time.Duration) error {
	files, err := filesToCompress(dir, maxAge)
	if err != nil {
		return err
	}

	for _, src := range files {
		dest := src + ".gz"
		if err = gzCompress(src, dest); err != nil {
			return fmt.Errorf("error compressing file [%v]: %w", src, err)
		}

		match, err := sameSig(src, dest)
		if err != nil {
			return err
		}

		if !match {
			return fmt.Errorf("%q <-> %q: signature don't match", src, dest)
		}

		if err = os.Remove(src); err != nil {
			log.Printf("warning: %q: can't delete - %s", src, err)
		}
	}
	return nil
}
