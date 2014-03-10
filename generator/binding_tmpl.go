package generator

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func generator_tmpl_binding_tmpl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x59,
		0x5d, 0x4f, 0xdb, 0x48, 0x17, 0xbe, 0xcf, 0xaf, 0x38, 0x45, 0xa8, 0xb2,
		0x69, 0xde, 0xb4, 0xd7, 0x79, 0x97, 0x0b, 0x9a, 0xa6, 0xd4, 0x52, 0x36,
		0x50, 0x62, 0x40, 0x2b, 0x84, 0xaa, 0x89, 0x33, 0x01, 0x0b, 0x67, 0x6c,
		0xec, 0x09, 0x5d, 0x64, 0xe5, 0xbf, 0xef, 0x99, 0xf1, 0x38, 0x1e, 0x8f,
		0xc7, 0x69, 0x02, 0x45, 0x2b, 0xad, 0xe2, 0x9b, 0xd8, 0x33, 0xe7, 0xfb,
		0x39, 0x67, 0x3e, 0x4e, 0x3e, 0x7e, 0x04, 0xff, 0x9b, 0x37, 0x81, 0xaf,
		0xde, 0x68, 0x08, 0xd7, 0x27, 0x13, 0x38, 0xb9, 0xf4, 0xcf, 0x4e, 0x87,
		0xe3, 0xe1, 0xc5, 0x89, 0x3f, 0xfc, 0x02, 0xff, 0x83, 0x93, 0xf1, 0x5f,
		0x30, 0xfc, 0xe2, 0xf9, 0x13, 0xf0, 0xcf, 0x0a, 0xd2, 0x6b, 0x6f, 0x34,
		0x82, 0xcf, 0x43, 0x18, 0x9d, 0x4d, 0x7c, 0xb8, 0xfe, 0x36, 0x1c, 0x83,
		0xe7, 0x03, 0x8e, 0x5f, 0x0c, 0xd7, 0x7c, 0x1d, 0x14, 0x5b, 0x09, 0x39,
		0xf1, 0x21, 0xcf, 0xa1, 0x77, 0x9e, 0xc6, 0x4f, 0x94, 0x11, 0x16, 0xd0,
		0x9e, 0x1f, 0x2e, 0x68, 0xc6, 0xc9, 0x22, 0x81, 0xd5, 0x0a, 0x2e, 0x27,
		0xde, 0xf8, 0x14, 0x82, 0xc7, 0x28, 0x80, 0xab, 0xe1, 0xc5, 0xc4, 0x3b,
		0x1b, 0x9b, 0xe4, 0x57, 0x34, 0xcd, 0xc2, 0x98, 0x21, 0x71, 0xa7, 0x83,
		0x53, 0x87, 0xd9, 0xf3, 0x62, 0x1a, 0x47, 0x19, 0xf4, 0x8f, 0xa1, 0x77,
		0x96, 0x70, 0x9c, 0xca, 0x7a, 0x13, 0x35, 0x26, 0x68, 0x12, 0x12, 0x3c,
		0x90, 0x3b, 0x2a, 0xc5, 0x94, 0xf3, 0xe7, 0x6a, 0x4c, 0xcc, 0x87, 0x8b,
		0x24, 0x4e, 0x39, 0x38, 0x1d, 0xc0, 0x27, 0xcf, 0x53, 0xc2, 0x70, 0xe2,
		0xf0, 0x47, 0x17, 0x0e, 0x13, 0xc2, 0xef, 0xa5, 0x58, 0x4f, 0x92, 0x64,
		0x48, 0x0d, 0xea, 0x39, 0xc8, 0x73, 0x39, 0xbd, 0x5a, 0x1d, 0x28, 0x3e,
		0xca, 0x66, 0x38, 0xef, 0x76, 0x3a, 0x01, 0x2a, 0x28, 0xc5, 0x0d, 0xbe,
		0x8f, 0x06, 0x3f, 0x4a, 0x3f, 0x8e, 0x05, 0x57, 0x8b, 0x2b, 0x07, 0x82,
		0xb3, 0xa6, 0x3c, 0x98, 0x4b, 0xd5, 0x83, 0x38, 0x5a, 0x2e, 0xd8, 0x57,
		0xb2, 0x08, 0xa3, 0x90, 0x0a, 0x0b, 0x94, 0x3a, 0x38, 0x9c, 0xf0, 0x74,
		0x19, 0x70, 0xff, 0x39, 0xa1, 0x82, 0x30, 0x63, 0xe4, 0x81, 0xfa, 0xf1,
		0x80, 0x2c, 0x68, 0x24, 0x98, 0x7b, 0x63, 0x7c, 0x83, 0x8a, 0x5e, 0x97,
		0x1c, 0x47, 0x82, 0x43, 0x10, 0x15, 0xd2, 0x75, 0xc7, 0x84, 0x64, 0x1c,
		0xdd, 0x28, 0x3c, 0x8e, 0xd6, 0xd2, 0x75, 0xb6, 0xef, 0x4b, 0x12, 0x85,
		0xf3, 0x90, 0xce, 0x9a, 0xfc, 0x49, 0x1a, 0x32, 0x5e, 0x33, 0xd9, 0x50,
		0xa2, 0x89, 0xe2, 0xe2, 0x1b, 0xc3, 0x6b, 0x17, 0xb7, 0x5a, 0x15, 0x36,
		0x43, 0x26, 0x87, 0x72, 0xe5, 0xa0, 0x78, 0xe6, 0x4b, 0x16, 0x80, 0x33,
		0x85, 0xa3, 0x2d, 0xb8, 0x5d, 0x28, 0x5e, 0x84, 0x1f, 0x8e, 0x2b, 0x84,
		0x85, 0xec, 0x0e, 0xf2, 0xb5, 0x2c, 0xf1, 0xa4, 0x94, 0x2f, 0x53, 0x26,
		0xa1, 0x2e, 0x5d, 0x2e, 0xe1, 0x16, 0xcf, 0x8b, 0x35, 0xfb, 0xb1, 0xf3,
		0x44, 0xa2, 0x25, 0x85, 0xa3, 0x3c, 0x97, 0x2f, 0x45, 0x40, 0x50, 0xc7,
		0x6a, 0xe5, 0xca, 0xd4, 0x57, 0xb8, 0x7c, 0x0e, 0xd9, 0xac, 0xd5, 0xac,
		0x26, 0x5d, 0x5e, 0x7c, 0xf5, 0x61, 0xda, 0x85, 0x2b, 0x21, 0xb7, 0x0f,
		0x52, 0xfc, 0xca, 0x66, 0x32, 0x02, 0x16, 0xce, 0x21, 0xcc, 0x06, 0xf1,
		0x92, 0x71, 0x9a, 0x2a, 0xe3, 0x64, 0x6e, 0x68, 0x58, 0xbc, 0x20, 0xac,
		0x84, 0x79, 0x2c, 0x48, 0xe9, 0x82, 0x32, 0x8e, 0x81, 0x9d, 0xc6, 0x28,
		0xb0, 0x6e, 0xbf, 0xe6, 0x03, 0x0a, 0xa0, 0xb5, 0xb9, 0x5a, 0x42, 0xd1,
		0x28, 0xa3, 0xa6, 0x31, 0x85, 0xd9, 0x12, 0x8f, 0xc9, 0x32, 0x91, 0x55,
		0x79, 0x4e, 0x52, 0x1e, 0x8a, 0xb2, 0x16, 0xa1, 0x32, 0xc8, 0x77, 0xb7,
		0x7f, 0xf8, 0xa8, 0xc0, 0xd9, 0x80, 0x0d, 0x46, 0x5b, 0xe8, 0xb3, 0xf8,
		0x25, 0x9e, 0xa0, 0x10, 0x85, 0x69, 0xff, 0xfe, 0x97, 0x0a, 0xf3, 0xa6,
		0xb9, 0xe2, 0x99, 0x2a, 0xe0, 0x51, 0xc6, 0x06, 0x98, 0x0b, 0x45, 0x6d,
		0x58, 0x5b, 0xc2, 0x5d, 0x37, 0x3f, 0x57, 0xf2, 0xfa, 0xa5, 0xba, 0x2e,
		0x9c, 0xa7, 0x74, 0x16, 0x06, 0x84, 0xa3, 0x34, 0x49, 0x3b, 0x7c, 0x5c,
		0x8f, 0x34, 0x05, 0x37, 0x47, 0x34, 0x70, 0xbc, 0x6c, 0x44, 0x32, 0x3e,
		0x88, 0x71, 0xe1, 0x64, 0x98, 0x09, 0x36, 0x5c, 0x76, 0xc7, 0xc6, 0x63,
		0x0a, 0x9b, 0x5e, 0xaf, 0xf7, 0x62, 0x78, 0x7e, 0x13, 0x44, 0x6f, 0x01,
		0xd3, 0x6b, 0xa0, 0xf2, 0xd8, 0x06, 0xa8, 0x5a, 0xe1, 0xc2, 0x4d, 0xcb,
		0x52, 0x61, 0xf6, 0x51, 0xb3, 0xee, 0x06, 0xd1, 0x32, 0xc3, 0x95, 0x63,
		0x5f, 0x75, 0xfb, 0xaa, 0xdb, 0x57, 0xdd, 0x6b, 0xab, 0x6e, 0x77, 0x58,
		0x4e, 0xf9, 0xbe, 0x64, 0xd6, 0x38, 0x9c, 0xf2, 0x9d, 0x4a, 0x66, 0xc7,
		0x48, 0xd3, 0x7d, 0xa4, 0xab, 0x48, 0xd3, 0x37, 0x8c, 0xf4, 0x68, 0x9f,
		0xd3, 0x55, 0xa4, 0x47, 0x6f, 0x99, 0xd3, 0xa3, 0x7d, 0x4e, 0x6b, 0x91,
		0xde, 0x3e, 0xa7, 0x9b, 0xab, 0x77, 0x35, 0x52, 0x6b, 0x04, 0xc8, 0x8f,
		0xf2, 0x2a, 0xab, 0x47, 0x44, 0x5d, 0x5c, 0xb5, 0xa8, 0x6e, 0x7f, 0x39,
		0x2f, 0xa8, 0xed, 0xb7, 0x71, 0x94, 0xdc, 0xc4, 0x52, 0xd3, 0x51, 0x98,
		0x55, 0x78, 0xd4, 0x6a, 0xdc, 0x17, 0x3a, 0x7f, 0xbd, 0x7d, 0x3c, 0xbe,
		0x4c, 0x12, 0x9a, 0x1a, 0xa6, 0x15, 0x30, 0x6e, 0x69, 0x9b, 0x4c, 0x66,
		0x01, 0x9c, 0x61, 0x9f, 0x13, 0xe2, 0x89, 0x13, 0x8e, 0xee, 0x62, 0xc4,
		0xad, 0xe7, 0xe1, 0xbb, 0x0b, 0xce, 0xcd, 0xad, 0x41, 0xd4, 0x05, 0x9a,
		0xa6, 0x31, 0x4e, 0x55, 0x2e, 0x90, 0x34, 0x25, 0xcf, 0xc2, 0xee, 0x05,
		0x86, 0xce, 0xc6, 0xf1, 0xc9, 0x5d, 0xd3, 0x22, 0xb3, 0xa0, 0xfc, 0x93,
		0x24, 0x36, 0xe5, 0x5d, 0x69, 0x9b, 0xc3, 0xcd, 0xc8, 0xa1, 0x21, 0xe2,
		0xae, 0x6b, 0xd1, 0x5d, 0xe9, 0x3f, 0x06, 0x82, 0x81, 0x61, 0x33, 0x47,
		0x7e, 0x76, 0x81, 0xbb, 0xb6, 0x8b, 0xbd, 0xb8, 0x14, 0x77, 0x81, 0x85,
		0x51, 0x75, 0x6d, 0xaf, 0xe8, 0x14, 0x8d, 0x12, 0x80, 0xba, 0x1a, 0x51,
		0x6b, 0xb1, 0x5b, 0x0f, 0x5a, 0x17, 0x02, 0x12, 0x45, 0x53, 0x12, 0x3c,
		0x6c, 0xe7, 0x8c, 0x5b, 0xfc, 0x6a, 0x3e, 0x15, 0x60, 0xca, 0x9e, 0x9b,
		0x10, 0x5e, 0xa6, 0x82, 0xa3, 0x19, 0x1a, 0xff, 0xd4, 0xe2, 0x1d, 0x8a,
		0x1e, 0xc3, 0x9c, 0x04, 0x34, 0xc7, 0x58, 0x47, 0x94, 0x39, 0x4a, 0x80,
		0xeb, 0x6a, 0x4d, 0x94, 0x9a, 0x06, 0x99, 0xa1, 0x42, 0x82, 0x61, 0x9a,
		0xde, 0xf0, 0x29, 0xb9, 0x42, 0x41, 0xf7, 0xe9, 0xff, 0xf8, 0xfb, 0x47,
		0x4d, 0x38, 0x8e, 0x7c, 0xf8, 0x60, 0x59, 0xbe, 0xb2, 0x9f, 0x21, 0x0f,
		0xee, 0x4b, 0x27, 0x6e, 0xc2, 0xdb, 0xa2, 0x9d, 0xd5, 0x24, 0xdc, 0x2d,
		0xf3, 0xd7, 0xc1, 0x21, 0x19, 0x35, 0xbb, 0x46, 0x7d, 0x11, 0x10, 0xd4,
		0x84, 0x39, 0xf0, 0x9e, 0xf7, 0xda, 0x6b, 0xd8, 0x62, 0x42, 0x55, 0x1c,
		0xfa, 0x33, 0xa3, 0x73, 0xb2, 0x8c, 0x78, 0xdf, 0x6a, 0x41, 0x14, 0xdf,
		0xf5, 0xbe, 0x12, 0x4e, 0x22, 0xe7, 0x60, 0xc9, 0xee, 0x09, 0x9b, 0x45,
		0x74, 0xa6, 0xdc, 0xed, 0xc3, 0x41, 0xd7, 0xf4, 0xdc, 0xfd, 0xc5, 0xa2,
		0x57, 0xff, 0xc2, 0xcb, 0xc6, 0x3b, 0x09, 0xfa, 0x24, 0x20, 0xcc, 0x41,
		0xb7, 0xf0, 0x5e, 0x60, 0xe6, 0xbb, 0x78, 0xa6, 0x29, 0x25, 0x0f, 0x86,
		0x1c, 0x23, 0xdf, 0xc9, 0x6c, 0x4c, 0xff, 0xe6, 0xdd, 0xb2, 0xe4, 0xca,
		0xac, 0x74, 0x8c, 0xc2, 0x40, 0x8d, 0x82, 0xe2, 0xdd, 0xb1, 0xa8, 0x8a,
		0xf6, 0x6e, 0x52, 0x59, 0x0d, 0xed, 0x66, 0x97, 0x1a, 0xdb, 0x65, 0xe8,
		0x65, 0x57, 0x97, 0xa1, 0x19, 0x6f, 0xd0, 0x56, 0x3d, 0x59, 0xd4, 0x81,
		0x77, 0x30, 0xd5, 0x52, 0xd3, 0x37, 0x89, 0x62, 0x73, 0xce, 0x8a, 0xcd,
		0xd9, 0x5c, 0x70, 0xf1, 0x8a, 0x55, 0xf2, 0xf8, 0x64, 0x1a, 0x51, 0x7b,
		0xdb, 0xcc, 0xd6, 0x32, 0x5b, 0x95, 0x8a, 0x45, 0xab, 0x6c, 0x7b, 0x75,
		0xe5, 0x0d, 0xfe, 0x32, 0xc9, 0x68, 0xda, 0xd2, 0xa5, 0xdb, 0xac, 0xae,
		0xda, 0x02, 0x3b, 0x5b, 0xe8, 0x93, 0x5e, 0xb5, 0x75, 0x5a, 0xf5, 0x2e,
		0xeb, 0xbc, 0xd6, 0x64, 0xd5, 0x57, 0xb5, 0x76, 0xe1, 0x62, 0x97, 0x70,
		0x9e, 0x9a, 0x2b, 0x98, 0xdc, 0xd9, 0xa5, 0xe6, 0x66, 0x1b, 0x35, 0x50,
		0xff, 0x17, 0xdc, 0xdc, 0x5a, 0xce, 0x1d, 0x2f, 0xde, 0xf7, 0x5e, 0xd4,
		0x38, 0x2f, 0x59, 0x7f, 0x53, 0xf3, 0x5c, 0x7a, 0xd7, 0x7e, 0x96, 0x7a,
		0xbf, 0x41, 0x55, 0x75, 0x58, 0xab, 0xce, 0x5a, 0x1b, 0x56, 0xaa, 0x55,
		0xd7, 0xba, 0x81, 0x57, 0x79, 0xa2, 0x81, 0xdb, 0xc0, 0x22, 0x97, 0x1f,
		0xd2, 0x1e, 0x13, 0x52, 0xa1, 0x5e, 0x45, 0x58, 0x9e, 0xfd, 0xb2, 0xe6,
		0xc1, 0xa0, 0x3d, 0x19, 0x0a, 0x46, 0x7c, 0x0f, 0x99, 0x3c, 0xfc, 0xe1,
		0x8e, 0x54, 0x07, 0xb9, 0x99, 0x78, 0xb5, 0xe9, 0xdc, 0xc0, 0x65, 0xd7,
		0xe5, 0xff, 0x15, 0x49, 0xf0, 0x06, 0x89, 0x20, 0x9e, 0x2d, 0x11, 0x37,
		0xfc, 0xb6, 0xe1, 0xa9, 0x23, 0x60, 0x04, 0x5f, 0xa2, 0x89, 0x51, 0xc7,
		0x68, 0x1f, 0x59, 0x70, 0x69, 0xc4, 0xdc, 0x06, 0xfb, 0x7f, 0x2e, 0xf2,
		0xf6, 0x83, 0xf0, 0xd6, 0x35, 0xb8, 0x3d, 0x22, 0xaa, 0x01, 0x98, 0x55,
		0xff, 0x79, 0x96, 0x64, 0x4f, 0x24, 0xad, 0x99, 0x51, 0xae, 0xaf, 0xd0,
		0x38, 0x56, 0x55, 0x08, 0x1a, 0x6b, 0x7c, 0x67, 0x37, 0x38, 0xfe, 0xfd,
		0xbf, 0x0f, 0xd7, 0x2b, 0x44, 0x33, 0xc7, 0x5c, 0x8b, 0x79, 0x65, 0xc4,
		0x1d, 0xd7, 0x76, 0x47, 0xb1, 0x6f, 0x8a, 0xdb, 0x21, 0x28, 0x18, 0x1a,
		0x5b, 0xa7, 0xba, 0x1c, 0xaa, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
		0x49, 0x97, 0x28, 0x44, 0x70, 0x1f, 0x00, 0x00,
		},
		"generator/tmpl/binding.tmpl",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"generator/tmpl/binding.tmpl": generator_tmpl_binding_tmpl,

}
