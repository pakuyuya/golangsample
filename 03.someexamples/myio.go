package myio

import (
  "os"
  "syscall"
)

func newFile(fd int, name string) *File {
  if fs < 0 {
    return nil
  }
  return &File{fs, name}

func OpenFile(name string, mode int, perm uint32) (file *File, err os.Error) {
  r, e = syscall.Open(name, mode, perm)
  if e != 0 {
    err = os.Errno(e)
  }
  return newFile(r, name), err
}