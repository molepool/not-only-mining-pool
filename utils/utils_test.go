package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

func TestSerializeNumber(t *testing.T) {
	if bytes.Compare(SerializeNumber(100), []byte{0x01, 0x64}) != 0 {
		fmt.Println(SerializeNumber(100))
		t.Fail()
	}

	if hex.EncodeToString(SerializeNumber(1<<31-1)) != "04ffffff7f" {
		t.Fail()
	}
}

func TestSerializeString(t *testing.T) {
	if bytes.Compare(SerializeString("HelloWorld"), []byte{0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64}) != 0 {
		fmt.Println(SerializeString("HelloWorld"))
		t.Fail()
	}
}

func TestReverseByteOrder(t *testing.T) {
	hash := Sha256([]byte("0000"))
	fmt.Println(hash)
	fmt.Println(ReverseByteOrder(hash))
}

func TestReverseBytes(t *testing.T) {
	hash := Sha256([]byte("0000"))
	fmt.Println(hash)
	fmt.Println(ReverseBytes(hash))
	fmt.Println(hex.EncodeToString(hash))
	fmt.Println(hex.EncodeToString(ReverseBytes(hash)))
}

func TestUint256BytesFromHash(t *testing.T) {
	result, _ := hex.DecodeString("691938264876d1078051da4e30ec0643262e8b93fca661f525fe7122b38d5f18")
	if bytes.Compare(Uint256BytesFromHash(hex.EncodeToString(Sha256([]byte("Hello")))), result) != 0 {
		t.Fail()
	}
}

func TestVarIntBytes(t *testing.T) {
	if hex.EncodeToString(VarIntBytes(uint64(23333))) != "fd255b" {
		t.Fail()
	}

	t.Log(int64(1<<31 - 1))
	if hex.EncodeToString(VarIntBytes(uint64(1<<31-1))) != "feffffff7f" {
		t.Fail()
	}
}

func TestVarStringBytes(t *testing.T) {
	t.Log(hex.EncodeToString(VarStringBytes("Hello")))
}

func TestRange(t *testing.T) {
	if !reflect.DeepEqual(Range(0, 8, 2), []int{0, 2, 4, 6}) {
		t.Fail()
	}

	if !reflect.DeepEqual(Range(2, 0, 2), []int{}) {
		t.Fail()
	}
}

func TestSha256d(t *testing.T) {
	hexStr := "01000000" +
		"81cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000" +
		"e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122b" +
		"c7f5d74d" +
		"f2b9441a" +
		"42a14695"
	raw, _ := hex.DecodeString(hexStr)
	if hex.EncodeToString(Sha256d(raw)) != "1dbd981fe6985776b644b173a4d0385ddc1aa2a829688d1e0000000000000000" {
		t.Fail()
	}
}
