package main

import (
	"testing"
	"unsafe"
)

func TestStructLayout(t *testing.T) {
	var info OrderInfo
	if size := unsafe.Sizeof(info); size != 56 {
		t.Fatalf("OrderInfo should be 56 bytes, got %d", size)
	}
	if off := unsafe.Offsetof(info.OrderCode); off != 0 {
		t.Fatalf("unexpected OrderCode offset: %d", off)
	}
	if off := unsafe.Offsetof(info.Amount); off != 8 {
		t.Fatalf("unexpected Amount offset: %d", off)
	}
	if off := unsafe.Offsetof(info.OrderNumber); off != 16 {
		t.Fatalf("unexpected OrderNumber offset: %d", off)
	}
	if off := unsafe.Offsetof(info.Items); off != 24 {
		t.Fatalf("unexpected Items offset: %d", off)
	}
	if off := unsafe.Offsetof(info.IsReady); off != 48 {
		t.Fatalf("unexpected IsReady offset: %d", off)
	}
	if size := unsafe.Sizeof(SmallOrderInfo{}); size >= unsafe.Sizeof(info) {
		t.Fatalf("SmallOrderInfo should use less memory than OrderInfo; got %d vs %d", size, unsafe.Sizeof(info))
	}
}
