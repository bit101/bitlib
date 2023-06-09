// Package bitmap creates bitmap images.
package bitmap

import (
	"bytes"
	"encoding/binary"
	"os"
)

// EncodeBmp encodes a bitmap.
func EncodeBmp(pixelData []float64, w, h int, filepath string) {
	var data []byte
	buff := bytes.NewBuffer(data)
	size := uint32(54 + len(pixelData))

	// bmp header - 14 bytes
	buff.Write([]byte("BM"))                              // bitmap signature
	binary.Write(buff, binary.LittleEndian, uint32(size)) // bmp size
	binary.Write(buff, binary.LittleEndian, uint32(0))    // reserved
	binary.Write(buff, binary.LittleEndian, uint32(54))   // offset for headers

	// bmp info header - 40 bytes
	binary.Write(buff, binary.LittleEndian, uint32(40))  // info header size
	binary.Write(buff, binary.LittleEndian, int32(w))    // width
	binary.Write(buff, binary.LittleEndian, int32(h))    // height
	binary.Write(buff, binary.LittleEndian, uint16(1))   // color planes
	binary.Write(buff, binary.LittleEndian, uint16(24))  // color depth
	binary.Write(buff, binary.LittleEndian, uint32(0))   // compression method
	binary.Write(buff, binary.LittleEndian, uint32(0))   // raw size (ignored)
	binary.Write(buff, binary.LittleEndian, int32(3780)) // vert res - pix per meter
	binary.Write(buff, binary.LittleEndian, int32(3780)) // vert res - pix per meter
	binary.Write(buff, binary.LittleEndian, uint32(0))   // color table entries
	binary.Write(buff, binary.LittleEndian, uint32(0))   // important colors

	for _, p := range pixelData {
		binary.Write(buff, binary.LittleEndian, uint8(p*255))
	}
	os.WriteFile(filepath, buff.Bytes(), 0777)
}
