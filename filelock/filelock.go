package filelock

import (
	"fmt"
	"os"
	"syscall"
)

// 文件锁
type FileLock struct {
	fpath string
	f     *os.File
}

func New(fpath string) (*FileLock, error) {
	if _, err := os.Stat(fpath); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(fpath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &FileLock{
		fpath: fpath,
	}, nil
}

// 加锁
func (l *FileLock) Lock() error {
	f, err := os.Open(l.fpath)
	if err != nil {
		return err
	}
	l.f = f
	// LOCK_EX 排他锁, 只有一个进程可以获得锁
	// LOCK_NB 表示当前获取锁的模式是非阻塞模式
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.fpath, err)
	}
	return nil
}

// 释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}
