package main

import (
	"io"
	"log"
	"os"
)

type PathTransformFunc func(string) string

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

var DdefaultPathTransformFunc = func(key string) string {
	return key
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathName := s.PathTransformFunc(key)
	if err:= os.MkdirAll(pathName, os.ModePerm); err != nil {
		return nil
	}
	fileName := "someFile"
	filePath := pathName+"/"+fileName
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	n, err := io.Copy(f,r)
	if err != nil {
		return err
	}
	log.Print("written ", n, " bytes to the disk ",filePath)
	return nil
}