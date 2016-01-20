package main

import (
    "hash/crc32"
)

func main() {
    // Read Opus file page into "bytes"
    // Zero checksum bytes
    table := crc32.MakeTable(0x04c11db7)
    hash := crc32.New(table)
}
