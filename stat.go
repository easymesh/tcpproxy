package main

import (
	"fmt"
	"log"
	"sync/atomic"
)

var upload_size int64
var download_size int64

func calcUnit(cnt int64) string {
	if cnt < 1024 {
		return fmt.Sprintf("%d B", cnt)
	} else if cnt < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float32(cnt)/1024)
	} else if cnt < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float32(cnt)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f GB", float32(cnt)/(1024*1024*1024))
	}
}

func StatUpdate(up int64, down int64) {
	atomic.AddInt64(&upload_size, up)
	atomic.AddInt64(&download_size, down)
	log.SetPrefix(fmt.Sprintf("[UP:%s DOWN:%s] ", calcUnit(upload_size), calcUnit(download_size)))
}
