// Code generated by fileb0x at "2018-10-08 15:32:41.965420773 +0300 MSK m=+0.027198443" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(10b22e201f5f07079d911b1ffe25a9bd.f7d60b70267ef579d9aad25918a6b087)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileSwaggerJSON is "swagger.json"
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x7d\x6b\x6f\xdb\xb8\xd2\xf0\xf7\xfe\x0a\xc2\xef\x0b\xb4\xc5\xa9\x9d\x9e\xdd\x83\x03\x3c\xfb\x2d\xdb\xb4\x5d\x63\x9b\x1e\x23\x97\xc5\x01\x9e\x14\x06\x2d\x8d\x6d\x6e\x25\x52\x25\x29\x27\x69\xd1\xff\xfe\x80\xa4\xee\x77\x4b\x96\xa3\x24\xfa\xd4\x54\xe6\x75\x66\x38\x33\x1c\xce\xe5\xc7\x0b\x84\x26\x16\xa3\xc2\x77\x41\x4c\x7e\x43\xff\xfb\x02\x21\x84\x26\xd8\xf3\x1c\x62\x61\x49\x18\x3d\xf9\x5b\x30\x3a\x79\x81\xd0\x97\x37\xaa\xad\xc7\x99\xed\x5b\xcd\xda\x8a\x5b\xbc\xd9\x00\x9f\xfc\x86\x26\xbf\xcc\xde\x4e\xf4\x37\x42\xd7\x6c\xf2\x1b\xfa\x61\xfa\xda\x20\x2c\x4e\x3c\xd5\x57\xb5\xfa\x8b\x39\xbe\x0b\xd3\x73\x4c\xf1\x06\x38\x12\xc0\x77\xc4\x02\x24\x24\xe3\x20\xd0\x4e\xff\x2a\x10\xa6\xb6\xfe\x84\x37\xa0\xff\xc6\x8e\xc3\x6e\x05\x92\x0c\xb9\xba\x23\x92\x5b\x70\xf5\x6c\x08\x4d\x24\x91\x0e\x14\x8c\x7d\xba\x98\x87\x4d\x76\xc0\x45\xb0\x80\xb7\xb3\xb7\xb3\x7f\xaa\x0d\xfc\x34\x9b\xc5\x72\x2b\xe2\xd5\x9e\x60\xdb\x25\xf4\x84\x62\x17\x84\x87\x2d\x10\x27\x3f\xa8\x58\x12\xfb\xe7\x49\xb0\xb4\x93\x1f\x0e\x5e\x81\xf3\x33\xea\xa2\xc6\xf0\x65\xe2\xbf\x6a\x45\x78\x13\x43\x2f\xf8\x66\x16\x27\x26\xd1\xb7\x2f\x6f\xe2\x0e\xc2\x77\x5d\xcc\xef\xd5\xfa\x2e\x40\x90\xef\x10\x40\x02\xbd\xd2\xeb\x11\x88\x51\xe7\xfe\xf5\x6c\x92\xe8\xc2\x3c\xe0\x1a\x23\x73\x5b\x75\x3b\x55\xed\x4c\x5f\x33\x53\xb2\xad\x87\x39\x76\x41\x02\xcf\xae\xea\x47\xe2\x6f\xb5\xee\x7b\x4f\x03\x52\x48\x4e\xe8\x26\x31\x82\xfe\x75\xcd\xb8\x8b\xd5\x4e\x27\xbe\x4f\xec\xec\xaf\x0a\x64\xea\xb7\xff\x4e\xaf\x05\xf0\xe9\xfc\x2c\xdb\x80\x68\xf0\x6f\x01\xdb\xc0\xb3\xbf\x71\xf8\xe6\x13\x0e\x6a\x27\x92\xfb\x90\xf8\xf1\xe7\x9b\xf2\xe5\x02\xf5\xdd\xcc\x86\xf4\x77\x5f\xe4\x66\x50\x44\xac\x20\x34\x49\x7d\xfd\xf2\x66\x9f\xfd\x67\x76\x78\xc1\x1c\xe8\x7f\x8f\x5d\x50\x92\x39\x7a\x57\x5b\x40\xf3\x33\xc4\xd6\xea\xf4\x84\xff\x28\x58\x21\x46\xd1\x0a\xb6\xd8\x59\xab\xaf\xb7\x5b\x62\x6d\xf5\x6f\x1a\x62\x44\x48\x8e\x25\xe3\xc8\xc2\x14\x79\xc0\xd5\x94\xe6\x57\x4b\x0f\x5c\x02\x24\x35\xee\x34\xbf\x24\x03\xa0\x6f\x3e\xf0\xfb\xc9\xf1\x21\xf0\x39\x3c\xd6\x28\x4f\x9e\xe1\xc2\xf5\x71\x2f\x5e\xb6\xe2\x15\xc7\xc0\x6a\xb8\x14\xcd\x69\xfa\x5e\x4a\x38\xd9\x8a\xd9\xf7\xc5\x73\x15\xfd\x92\x99\x2b\xf3\xab\xb0\xb6\xe0\xe2\x14\x4f\x0c\x7e\xf9\xff\x1c\xd6\x6a\xcc\xff\x77\x62\xc3\x9a\x50\xa2\x30\x23\x4e\x34\xf3\x32\x6c\xcb\xb0\xb0\x0b\xf8\xe6\x83\x90\xe9\xd3\xfa\xf3\x45\xd1\xdf\x49\x36\xca\x41\x78\x8c\x0a\x10\x99\xa9\x27\xbf\xbc\x7d\x9b\x5b\x4d\x96\x38\x02\x8e\xcb\xf5\x02\xec\x52\xea\x9c\xd8\xb0\xc6\xbe\x23\x6b\xc7\xb3\xb6\xc0\xf9\x3d\x02\xce\x59\x8e\x17\xec\x0b\x20\x9c\xe5\xee\x67\x66\x11\xbf\x2b\xd4\x34\x00\x52\xf8\x97\xf9\x37\xd8\x4f\x28\xec\x02\xc9\x96\x94\x68\x1b\x38\xa0\x44\xfb\x08\x52\x09\xf0\x48\xb8\x1b\x99\x56\x2f\xd2\x3e\x82\x3c\x75\x9c\x70\x96\x51\x9c\x8d\xe2\xec\xe9\x8b\xb3\xcc\x1e\x4f\x91\x00\xa9\x76\xb1\x26\x8e\xa2\x77\x24\x40\x11\xbf\x04\x1b\xdd\x12\xb9\x45\x16\x73\x5d\x3c\x43\x97\x00\xe8\x66\x62\x63\x76\x33\x41\x1e\xb6\xbe\x2a\x0d\x79\xcd\x38\x72\x19\x07\xa4\xd4\x71\x05\x56\xc2\xe8\xac\x6c\x7b\x66\xf8\x03\xec\x4e\x41\xd8\xd5\x64\xfc\xb6\x84\x18\x09\x95\xb0\xc9\xcf\x15\xae\xc4\xc3\x1b\x18\xc4\x3a\x80\x2f\xf7\x59\x4b\x5f\xf2\x48\xa0\x70\x90\xe6\x32\x24\xdc\x22\xe6\x1c\xdf\xe7\x39\x07\x91\xe0\x8a\x82\x7e\xa5\xd2\x27\xb8\x50\x64\x9a\xff\x2c\x97\x3b\x4f\x55\x5e\x12\xd7\x63\x5c\x9e\x04\x37\xd3\x94\xc4\xf4\x98\x68\x20\x32\x2f\xc3\x9e\x75\x32\x73\xae\x67\x0a\xef\xc0\xa2\x4a\x4c\x9a\xa6\xd1\xc8\xa3\x9c\x7c\x34\x72\x72\x5f\x9e\xf1\x4b\xed\x19\x0a\xe9\x05\x19\x4a\x05\xbb\xeb\x41\x32\xc4\x75\x11\xb2\xa0\x67\x7b\xe4\x0b\x74\xe4\x66\x27\xbe\xa9\x92\x1c\x1c\xf8\x60\x9e\xfa\xf3\x3e\xaa\xc5\x4f\x49\x2d\x1e\xc2\x25\x3c\xa0\xa8\x4f\xe4\xa0\xf7\xee\x7f\x36\xd6\x73\x46\x96\x75\x20\x96\xe5\x10\x97\x48\x51\x61\xc3\xae\xe6\x62\x99\x0d\x5f\x6e\x99\xef\xd8\x68\x05\xc8\xda\x32\x01\x14\xad\x09\x17\x91\x66\xf2\x06\xdd\x6e\x81\x03\x5a\x73\x00\x64\x4c\x6b\xb1\xad\xde\xe2\x80\x65\x64\xca\xd6\x17\x16\x8f\xb3\x1d\xb1\xc1\x46\x16\xf6\xb0\x45\xe4\x7d\x8a\xd3\x75\x61\xa0\xef\xcc\x64\xa6\x39\xf2\x05\xa1\x1b\x6d\x62\x28\x9e\x29\xc3\x53\xcf\x08\x07\x4b\x9a\x21\x46\xf3\xf9\x68\x6f\x18\xcd\xe7\x83\x36\x9f\x0f\x41\x5c\x1a\x96\x61\x98\x85\x61\x1c\x87\x37\x5a\x37\x15\x9e\x01\xa7\x7d\x1e\x46\xeb\x66\x72\xed\xe0\x16\xec\x68\xda\x26\x3a\xfa\x47\x90\xd1\x61\x19\x35\xf5\x51\xa0\x8c\x02\x65\x00\x02\x65\x34\xd1\x3e\x49\xe9\xf0\xe6\x89\x5f\x65\xd6\x8c\x23\xc5\xe0\xd0\xea\x1e\x5d\x61\x4e\xd6\xeb\x2a\xc9\x33\xde\x61\x46\x91\x33\x8a\x9c\xf1\x0e\xb3\x87\xd4\x18\x6f\x2f\xc7\x91\x4f\x36\x38\x20\xe1\x70\xb7\x92\x33\x3d\x9e\x76\xad\xd9\xeb\x72\x62\xfa\x9d\x3a\xce\x78\x45\x19\xe5\xc5\x28\x2f\x9e\xc7\x15\xc5\x30\x9f\xe7\x6e\x21\x2a\xf2\xde\x3f\xb8\xa5\xc8\xcc\x55\x63\x1e\x1a\x35\xf4\x91\xe3\x8e\x1c\x77\x74\xd2\x3f\x92\x4b\xfb\x9e\xe6\xa9\xe6\x46\xa6\x27\x6f\x54\xea\x2d\xb6\xab\x4a\x40\x8c\x81\x5c\xa3\x8c\x18\x65\xc4\x28\x23\x1e\xa7\x41\x69\x8c\xe1\x7a\xd4\x06\xa5\x7a\xf1\x64\x1a\x8e\xe2\x69\x14\x4f\xa3\x78\x1a\xc5\xd3\x31\x38\xfa\x73\x32\x5f\x15\x85\x17\x35\x32\x52\x35\x8e\x2e\xfa\x08\xd1\xab\x37\x72\x88\x90\x35\xb6\xaa\x31\xb6\x68\x64\xf5\x8f\x8f\xd5\x1f\x9e\x17\x45\xd1\x55\xea\xcc\x3c\x9c\xaf\x4f\x70\x1c\x47\x67\x9f\x03\x72\xc4\xc0\xe5\x26\xc0\x70\xbd\x83\x4d\x88\x83\x91\x23\x8e\x1c\xf1\x29\x28\xbf\x43\x30\x1e\x14\xf2\xb5\x9e\x3d\x50\x42\x2d\xe8\x39\x39\xd0\x87\x62\xec\xe4\x87\xc2\xfa\xfe\x99\xcc\x1a\x33\xd5\x6b\xcf\x6e\xc8\x54\x4d\xcb\x91\xa9\x8e\x4c\x75\x64\xaa\x87\x65\xaa\xa9\x93\x55\x6f\x92\xed\x6c\x4b\xd0\xff\x3e\x84\x29\xa1\x71\x72\x04\xe4\x6b\x90\x8c\xd6\xe1\x7d\xd9\x79\x60\x1e\x6e\xc0\xce\x4d\xcb\x91\x9d\x8f\xec\xfc\xf9\x18\x88\x1f\x11\x03\x7c\x4e\xc6\xd4\x63\x84\x86\x6a\xd2\x6d\x16\x15\xaa\x0e\xec\xe8\x6d\x3d\xf2\xc5\xd1\x9a\x3a\x06\x4f\x3e\x97\xe0\x94\xa6\xec\x31\x8a\x4b\x19\x99\xe4\xc8\x24\x47\x26\xf9\xec\xc2\x37\xa2\x9a\x0c\x89\x39\xe2\xca\x0c\xa7\x96\x05\x42\x7c\x82\x1d\x38\x49\x65\xae\x84\xb2\x26\x77\xd3\x0d\x9b\x06\xb9\x70\xd5\xef\x1b\x22\x67\x16\xa3\x12\x13\x0a\xdc\x77\x67\x14\xe4\x89\xb5\x0d\xd4\xc3\xa9\xa9\x2a\xc1\x4f\x76\x40\x6d\xc6\x4f\x36\x44\x6e\xfd\xd5\xcc\x62\xee\x49\xa2\xcf\xc9\x57\x7f\x05\x53\xcb\x21\x40\xe5\x89\xf7\x75\x73\xe2\x32\x1b\x9c\x49\x4a\xdf\x2c\xcd\x64\x9e\x58\x72\x36\xc7\x6f\x49\x17\x14\x4c\x2d\x50\xcc\xf8\x74\x30\xb8\xb5\xc5\x74\x43\xe8\x26\x8c\x46\xd7\x3e\xc5\x58\x18\x42\x8f\x21\x10\x42\x86\xad\xfe\x06\x2b\x7e\xae\x9c\x78\x5c\xf1\x5d\x49\x32\x04\x39\x09\xc3\xd8\xb3\x64\x5a\x95\x30\x37\x71\x74\x09\x95\xff\xfe\x57\xfa\x47\x8d\x83\xf0\x28\xbd\x0b\x87\xaf\x10\x42\xed\x90\x16\xa1\x42\xa4\x71\xf1\x4e\xc1\x09\x4c\xc8\x7c\x3d\x1a\x0a\x5a\x17\x62\x80\x02\xd8\x60\x6b\x44\x48\xdd\x38\xc2\x47\x7b\xd0\x9b\x81\x96\xc4\x2e\x83\x7d\x01\xdf\x4c\x43\xd7\xac\x7b\x7e\x76\x78\xe8\x56\x1f\x89\x15\x71\x9c\x29\xdc\x49\xe0\x14\x3b\xc5\x58\x60\x74\x4d\x36\x1f\x88\x03\x45\xc7\xb6\x31\x71\xea\x51\x96\xeb\xf4\x30\xfb\xc1\xe8\x33\x4e\x2a\x79\x3f\x93\xc9\x1a\xcc\x10\x2d\x41\xaf\x1a\x1e\x19\xec\x0d\x38\x91\x81\xfb\x39\xf6\xaa\x68\x3e\x6c\x83\xa6\x53\xa4\xc7\x30\x0c\x46\x7f\x46\x2e\xf6\xea\x49\x3a\xa1\x0f\xc4\x1a\x4c\xd6\x08\x32\xb1\xb1\xc4\x21\x8c\xbe\xd4\x23\xdb\x3c\xc6\x2d\x71\x56\x9c\xe5\x84\x99\x6a\x48\x18\x45\xfa\x9d\x89\x50\x74\xf1\xe1\xdd\xaf\xbf\xfe\xfa\x3f\x28\x60\x4b\x6f\x5a\x21\xd4\xbc\xf0\xdb\xa7\xb2\x98\x58\xf4\x5e\x32\xeb\x2a\x96\x8f\x11\x78\xcf\x12\xdb\xcf\x0c\x66\xe4\x7a\xfd\x56\x4d\xc3\xc3\x6e\xd4\x28\xdc\xa5\x1b\x0d\x9a\x1d\xfa\xb0\x45\xd1\x99\x9d\xc6\x36\x23\x14\x4e\xc0\x6e\xa9\xae\x48\xd5\x6e\xf0\xff\xe8\xde\x03\x3e\xcf\x67\x69\xfa\x2b\x3d\xd3\xaa\x5d\xd9\xb9\x46\x9a\x88\x6b\x0f\x37\xb6\x6d\x4d\xca\xd8\x59\x94\x49\xae\x34\x58\x07\x07\x2c\x93\xce\xb6\x01\xb4\x74\xc3\x32\x70\x65\x3c\xbe\x5a\x4a\x2f\x35\x50\x19\x59\xe6\x4d\x19\x25\x66\x8c\x3a\x56\x53\x7e\x43\x48\xf3\xb8\x68\xdf\xc3\x24\x75\xd3\xbc\x1a\x71\xa6\x4d\x0e\x67\xc1\x67\x42\x91\x0d\x9e\xc3\xee\x5d\xa0\xb2\xad\x1c\x23\x6e\xe6\x01\x2b\x2b\xd8\x4c\xae\xdd\x3d\x44\x1b\x73\x5d\x4c\xed\x03\xd0\x41\xf1\xc9\xab\x45\x7b\x30\x7d\x21\xe3\x0c\x94\xac\x9e\xe9\xd4\xe0\x27\x6f\x73\xdb\x97\x5a\x13\x0b\x07\xba\xeb\x6b\xc1\xef\xe9\xae\xe1\x22\x53\x2d\x93\xab\x33\x44\xd4\x56\x1e\xcd\xdd\x94\x4f\xd4\xcf\x3c\xf1\x35\xd2\x45\x2e\x40\x30\x9f\x5b\x70\x54\x51\xef\x31\x2e\xfb\x27\xa5\x05\xe3\xb2\x21\x8e\x16\x7a\x41\x85\x6b\x35\x4c\x6c\xe9\x32\x9f\xca\x81\x91\xbf\x69\x7c\x6e\x56\x36\x60\x76\xad\x11\xd1\x80\x65\xab\x76\x69\xb6\xad\x6b\x1f\x10\x1a\xb3\xef\x43\x5d\x3c\xd4\xc0\xa9\xff\x73\x26\x99\xc5\x9c\xe6\x1c\xbb\xc7\xa3\xd1\x93\xa5\x25\x7d\x1e\x7e\x16\xed\xbe\x11\xcb\x58\x64\x60\x35\x44\x92\x0b\xce\x51\x03\xa2\x0b\xb2\x29\x4e\xa7\x51\x69\xd5\x94\x8a\xf7\x1a\xe9\x93\x0f\xf6\x41\xc8\x50\x8f\xb5\xd4\x0f\xfe\x8d\x09\x4d\x6d\xb3\x35\xa1\x9d\xab\xce\x85\x38\x4f\x2c\xa5\xfd\xe0\x3e\x95\x8b\xc4\x66\x8e\x22\x3c\x84\xbf\xea\xb6\xec\x4b\x7f\x95\x59\xf4\x60\xc8\x37\x99\xa4\xb3\x82\x76\x13\xcd\xd0\x74\xda\xfe\xfe\x61\x02\xfc\xda\xc2\xf1\x93\xee\x5d\x8c\xa3\xc0\xc5\xab\x35\x8a\xb2\x2e\xdf\x29\xab\xdc\x43\x59\x45\x5b\x62\xf5\x2c\xbe\x6e\x94\xe3\x34\x6e\x94\x96\x80\xf1\x5d\x45\xb4\x65\x3a\xd1\x72\x45\xd5\x8d\x85\x83\xae\xe5\xbd\xc7\x9d\x05\x5b\x92\xec\x4a\x71\xbc\x62\xcc\x01\x4c\x2b\xb0\x70\x6a\xfa\x97\x5d\x39\xc2\x35\xf7\xad\x72\x35\xbf\x6b\x84\x2b\x2a\x5e\xf1\xb0\x8d\x93\x03\xb6\x27\xea\x7b\xd0\xd2\xf3\x1d\x67\x29\xc0\xe2\x20\x1f\xe8\x12\xac\x6f\x54\x0b\xdf\x71\x2e\xf5\x2a\xc4\x68\xfd\x2c\xb7\x7e\x16\x31\x8e\x7e\xb4\xd6\x8b\x0c\x5f\xca\x08\x1a\xe6\xf8\x8a\x74\x0b\xa4\x41\x36\xc9\x77\xd0\x12\xcd\xcf\xd0\x2b\x5d\x00\x88\xac\x13\xdc\x15\x11\xfd\x9a\xa7\x8b\xaa\x86\x83\xbe\x6e\x49\xef\xe1\x54\x29\x01\x93\x92\x8e\x58\xfa\x0d\xef\xe7\xb1\x5c\xb8\x34\xbd\x8a\x45\x22\x93\xd8\x59\x5a\x9e\x5f\x03\x04\xdd\x0e\xbd\x5b\x5c\x23\x5f\xe0\x0d\xa0\xd5\xbd\xf6\xc4\x89\xb9\xad\x3a\xf3\x72\x4b\x44\x91\x8d\x6c\x1f\xac\xfa\x75\x68\xbd\x52\x2b\x79\xb7\xb8\xae\xda\x8f\x0b\x2e\xe3\xf7\x8d\xb6\x74\x71\x7a\x3e\x8c\x2d\x9d\x9b\x35\x17\x9b\x12\x80\x0b\xb3\xea\x06\x88\xff\x2b\x68\x3c\x3c\x2d\x35\x47\x92\x4d\xb4\x1a\xd3\x54\xe9\x36\x6a\x06\x4e\x41\x82\x40\xe6\x20\xa8\x23\xb7\x8f\x49\xb6\x54\x1f\xd9\x61\xe2\xe0\x95\x03\xcb\x9e\x59\xd2\x69\x38\x51\x35\x6f\xe2\x80\xed\xfb\x65\xef\xec\x11\xdb\xf7\x75\xeb\x78\x40\x06\xed\xd3\xa3\xa1\xe5\x3a\x9e\xaa\x66\x4d\x26\xd0\xa7\xf7\xf5\x98\x69\x0a\xd6\x32\xbc\x93\xfc\x57\x8e\x35\x95\x1f\xe5\xa0\x6d\xd9\x3d\x05\x05\x6c\x2e\x88\xa7\x6a\x7f\x9e\x4b\xd8\x65\x63\x31\xfc\x18\x38\x68\xdd\x3b\x64\xa6\x65\xe9\xdd\xb0\xe3\x43\x64\xf2\x96\xd9\xd3\x7d\x2b\x71\x0b\x6e\xa6\x8e\x27\xb6\x3e\x44\x14\x96\x96\x79\xab\x40\x66\x59\x1f\xa5\x7c\x62\xc4\x83\xff\x18\xa4\x99\x67\x4b\x7d\x57\x8c\x5d\xfa\x42\x6f\x3e\xf4\xea\x96\xc8\x2d\xf3\x25\x5a\x11\xc7\x21\x74\xf3\xfa\x31\xf9\xf7\xbd\x79\x94\x16\xa8\x9e\xbd\x12\xdf\xa7\xde\x2a\xb3\xa4\xf3\x9e\xee\xb4\xee\x04\xf7\xd3\x1d\x76\x7c\x40\x1e\x26\x5c\x29\x4e\x40\x77\x84\x33\x6a\x18\x2f\xe6\x44\x49\xbf\xd6\x96\x22\x3d\x74\xce\x48\xf4\xd0\x4f\x22\x66\x55\xad\xa5\x80\xee\x3d\x3c\x06\xf2\x9e\x57\xb9\x31\xbc\xe7\xda\x81\x41\x48\x4c\x6d\xcc\x6d\x24\x80\x13\xec\x90\xef\x0a\xbd\xe8\x74\x31\x37\x8e\xda\x37\xf4\x1c\x84\xbe\xf5\x4c\xa7\xea\xc6\xa3\x9a\x4b\xf3\x13\x72\xcd\x2f\xbf\xdd\xd0\x7f\xa0\x9b\x09\xa1\x3b\xec\x10\x5b\x3b\xc4\x2b\xc8\xdc\x4c\xcc\xf7\x6f\x3e\x93\x18\xc1\x9d\xa5\xbd\x57\xc3\xaf\xba\xad\x31\x52\x99\x79\x26\x37\x74\x36\x9b\x81\xb4\x66\xb3\xd9\x0d\x9d\x9f\xa9\xf9\x7c\x4a\xbe\xf9\x10\xcc\x46\x6c\xa0\x92\xac\x89\x65\x7a\x59\xcc\x86\x1b\x7a\x06\x12\x13\x47\x2b\xfd\xcc\x33\xfe\x42\xfa\x5e\x06\x77\x99\x45\x0a\xf4\x95\x50\x1b\x9b\xc9\xd7\x04\x1c\x1b\xbd\x0c\xd5\xa6\x97\xc8\xf5\x85\x44\x2b\x40\x94\xd1\xe9\x77\xe0\x0c\x69\x82\x08\xd7\x4a\x99\x44\x40\x99\xbf\xd9\x22\x49\x36\x5b\xa9\xcb\x72\xad\x01\x6c\xb4\x61\xde\x16\x78\xd8\x8e\x07\x0f\xdd\xe8\xe5\x47\x66\xbf\x44\x36\x03\xf1\x52\x22\xb8\x23\x42\xaa\x26\x1f\xd4\xac\xe9\xa5\x0a\xd0\xd6\x80\xf4\x99\x13\x5d\x84\xac\x06\xc7\x03\x59\xb7\x02\x64\x14\x1f\x31\x0d\xf3\x86\x16\x09\x03\xa9\x12\x53\x9e\xdd\x6c\x8c\xf7\x9c\x97\x19\x47\x02\x92\x68\xff\x28\x15\xf4\xaf\xb0\xbc\x2c\xb7\x52\x7a\x3d\x89\x3b\x73\xcf\xfd\xe3\xea\x6a\x71\x6c\x9e\x63\x22\x38\x72\x4c\x66\x7e\x56\xcd\x66\xcc\x61\xe6\xe0\x71\x10\x5a\x93\x4c\x9d\xeb\x44\xa0\xd3\xde\xf4\xae\xce\x74\x63\x72\xf8\x53\x35\x2e\xc6\xd9\x1e\x44\x75\x79\x7c\x4f\xf7\x12\xb0\xff\x99\xde\x7c\x01\xe0\x55\x8b\x0c\xe8\x15\xc4\xb4\x50\x4f\x45\xe1\x94\x92\x66\xb9\x71\xea\x98\x3b\xbd\xac\xa3\xb0\xcb\x1c\x89\x09\xe0\x3b\x62\x85\xf1\x5a\x25\xbb\x3d\x46\xe4\x4e\xd1\x96\x3e\x64\x79\x61\x76\x4b\xb1\xa8\x48\x6c\x29\x76\x89\x4d\x08\x0c\xc3\x56\x2b\xf6\xf8\x38\x1c\x6b\x8b\xa0\x34\xcf\x38\xaf\xed\xcd\x1d\x3e\xd7\x2b\x8b\x85\xec\xe0\x0a\x6f\x1a\x77\x1b\x8a\xae\x37\x77\x3d\xc6\xe5\x45\x18\xd0\x5c\x4e\x5a\xe9\x86\x86\xc4\x82\xbf\xf1\x5a\x02\x47\xa1\x16\x23\x10\x71\x53\x3e\x4c\x7b\xc3\x7f\x8d\x89\x03\x76\x5f\xb7\xfd\x68\x23\xbe\xd3\xf4\xbe\xff\xc1\x2c\xa8\xe4\x89\x50\x0d\x37\xa8\xe5\xce\xc3\x25\x0d\x9d\xdc\xf4\xb3\x44\x63\x9a\xd3\xad\x15\xe1\x19\x90\xc7\xf4\xb7\x66\x5c\x3f\xac\x04\xbc\x5b\x1c\x82\x99\xd5\x20\xc3\x9c\x96\xc1\x41\x35\x1d\x4e\x5b\x0a\x4e\xdf\x91\x69\x38\xaa\x0f\x0a\x8a\x8c\x42\x74\x8e\xdb\x9f\xdf\x5e\xf5\xe4\x47\xf5\xc8\x3d\x18\x02\xa1\x1b\x0e\xa2\x4a\x73\x08\x5a\xa4\xad\xb8\x24\xe8\x76\x20\xcf\x56\xee\x3b\x20\x9e\x4c\x4c\xdd\x18\x06\xf7\x88\x1c\x41\x34\xe5\xf5\x24\xa1\x2f\x7c\xa7\xa9\x5b\xfc\x45\xf2\x04\x0c\x8f\x3f\x40\xdd\x9b\x4f\xaa\x5d\x21\xaf\x80\xae\xef\x3d\xd1\x38\xbd\x29\x54\x01\x53\x6b\xa8\x4b\x65\x5b\x0f\x06\x67\xe7\x58\x08\xb2\x83\x6b\x2a\xfc\x95\x42\xd2\xaa\x71\xde\x80\x9a\x9e\x75\x39\x04\x94\xaa\x65\x1c\x5d\x05\xf2\xa3\x31\x3a\x65\x13\x88\xee\x0e\x0f\x64\x81\xbc\x88\xe6\x1f\x5c\x32\x82\xcf\x05\x4c\xb7\xbc\x78\xcb\x74\x9a\x28\xac\x1d\xd9\x00\x70\x2a\x61\xcc\x9e\x52\x9c\x67\x61\xd3\xc4\xfd\xd6\x4a\x6b\x1b\x15\xc7\x31\x99\x24\xe5\x11\x7a\xaf\x76\x70\xf5\x2e\x33\x33\x17\x3e\x3c\x66\xf6\xaa\xb3\xfc\xec\x88\x20\x2b\x07\x90\xee\x60\xd2\x7b\x6c\x21\xc6\x7f\xcb\x0d\x57\xbc\x5c\xba\xf8\x6e\x09\x77\x72\x19\x5c\xb4\x5a\xda\xa9\x6b\xbd\xd5\xce\xf1\xdd\xfb\x3b\x79\x19\x4c\x52\xba\x12\x42\x8f\xb1\x92\x39\xad\x5f\x89\xe4\x78\xbd\x26\x56\x8f\xab\xb8\x0a\x66\x38\xae\xda\xa4\x07\x5e\x3a\x6c\x43\x68\xb7\xe1\x3f\xe9\x21\x4a\x5c\xc0\x4a\x58\x7f\x75\x50\xa8\x38\x62\x00\x46\xd2\x3f\x4b\xf4\xe8\xf7\x7f\x2d\x80\x1b\x76\xd8\x50\x6c\x5d\x8b\xb4\xcf\xff\x50\x54\x93\x48\x1e\x19\x60\x36\x91\x5c\xa6\x65\xd2\x6c\x1d\xb8\x8f\x84\x29\x8b\x14\x7b\x8b\x58\xdb\x0d\x9d\xaf\x91\xf1\x56\xb9\x24\xdf\x01\x79\x9c\xed\x88\xd2\x4f\x28\xa3\x53\x0f\xb8\x20\x42\x6a\xb7\x03\xe3\x9b\x72\x4b\x1c\x07\xad\xa2\x92\x09\xb3\x0e\x1e\x29\x9e\xbf\xd4\x31\xcc\x7d\xb9\xa4\x2c\xae\x3f\xe9\xe1\xf7\x14\x85\x55\xe4\x1d\xcf\xae\x04\xe3\x54\x92\x4c\xb6\xdf\x3d\xaf\xbd\x49\xfc\xfd\x68\x79\x7b\x8d\xc7\x28\x0e\x8e\x0f\x14\xa2\x65\x64\xd1\xeb\x07\xda\xef\x83\x79\x2e\xc3\x69\x8e\x25\xe2\xd5\xba\x8f\xb1\xc1\x39\x6d\xb4\x41\xb1\xec\x33\x60\x8a\x88\xa5\xe7\xaf\x9c\x72\xf9\x58\x3f\xfc\xc2\xf4\x3f\xae\x9f\x96\xf1\xf1\xef\xf5\xb8\x1b\x97\xfc\x8a\x13\xef\xf1\x0a\xed\x86\xfa\xee\xaa\x7c\x6e\x9b\xf9\x2b\xa7\xea\xa4\x2f\x78\xa9\x52\xd3\x4d\xa1\xa9\x8d\x46\xa8\x52\x66\x82\xa9\x97\x7d\xee\x3c\x58\x40\x05\x00\x82\xd4\x0a\x82\x7c\x87\x9e\x80\x10\x4b\xaf\x01\xdf\x39\xeb\xec\x41\xe9\x86\x69\x83\x50\x24\xac\xbb\x5a\x84\xe2\x81\xfa\x52\xbc\x0a\xcc\x92\x95\x7a\x57\xbc\xed\x01\x2a\x5f\x8b\x8b\x7a\x8f\xf9\xa8\x4d\xca\x4f\x00\x2d\x38\x5c\x80\x03\x58\x00\x0a\xc7\x68\x8d\xb5\xb9\xf8\xac\x93\x04\x57\x72\xfc\xc2\xb3\x17\x4c\x5d\xd1\xbd\xd9\x85\xaa\x6a\xf0\x4b\xc9\x1f\xfe\xa9\x7e\xe5\x60\xba\x39\x11\xe0\xee\xc2\x0b\x58\x84\xc3\x74\x12\x81\x1c\xfa\xb0\xdc\xea\x07\xbc\xe0\xd1\x26\x55\xa6\x61\x4f\xf3\x4e\xa6\xc4\xc3\x24\xd0\x48\x96\xd9\xc7\x9b\xf0\xbb\x97\x48\xd5\x51\x6f\x09\xea\x94\x0d\xa1\x3c\x7f\x43\x6a\x91\xad\xfd\xa8\xcd\x20\x15\xe9\x1c\x92\x5b\xee\xc9\x27\xcf\x4c\x91\xc9\x7f\x32\x18\x56\x12\xdd\xa4\x8c\xb4\x32\x97\xd3\xf3\x74\xd6\x8f\xa1\xe5\x33\x5e\xb0\x2a\xef\xb6\x05\xb3\xb3\x79\x7c\xec\x4e\x59\xf2\xc6\x5c\x00\x07\x7d\x54\xf5\x1c\x76\xdf\xe1\x62\xa9\xbb\x8f\xa1\xfc\xee\xb1\x1f\x58\xf7\x89\x1b\x5f\x30\xfb\xb8\x01\xe3\xc9\x23\xfe\xd8\x23\xc5\xfb\xdd\x4b\x2e\x44\x7c\x30\x92\x28\xa2\x99\x4a\xd6\x5e\x13\xca\xdd\x89\xd7\x7b\x5b\x2c\xda\x9f\xc9\x85\xee\x5d\x66\x01\x97\x98\xcb\xa5\xc5\x7c\x2a\x7b\x8b\x84\xd6\x73\xbc\xd3\x53\x94\x9d\x60\x2e\xeb\xc5\x82\xc7\x6c\xa4\x9b\x1e\x56\x2e\x5c\xaa\x21\x4f\xe5\x30\x29\xaf\xee\x16\x1c\x36\xc9\x29\x16\x5d\x6f\xbe\x6a\x88\xbe\x14\x0b\xa5\x26\x35\xcd\x40\x69\x0f\xf2\x9e\x9b\xcf\x0c\x98\x45\x4c\xe2\x72\xab\x3d\x1d\xc3\x64\x82\x48\x01\x11\xbd\xba\x7a\xb7\x40\x8c\xa3\xeb\xb3\xc5\xeb\x87\xf1\xf1\x6f\xb0\xcb\x0b\x50\x68\xa8\xf7\xe9\x48\xb5\xab\xf3\xe0\x88\x02\xce\xb8\xea\x75\x10\xcf\x8d\x65\x37\x33\x6c\xf8\xac\x97\x35\xc7\x0e\xc4\x20\x16\xa5\xa2\xad\x42\x40\x00\xd3\x74\xf0\x45\xec\x8d\xa1\x34\x15\x4c\x6d\x25\xe5\x5b\xa7\x4b\xf3\xfc\x49\xce\x40\xbd\x87\x43\x65\xad\x4e\xa5\xd6\x48\x28\x72\xfb\x51\x31\x4a\x35\xa5\x46\x3a\x92\xd2\x8e\x08\x45\xe7\xa4\x9f\xc5\x0d\x56\xf5\x09\x09\xeb\x2a\x5d\xf1\xa3\x94\x51\xe5\x8a\x73\x4d\x8a\x3c\x42\x02\x0b\x77\x44\x3c\x85\x53\x9a\x34\x21\x99\x70\x98\x32\xc2\x8f\x1b\x07\x21\xbf\x86\x03\x51\xb8\x8d\xf9\x4d\xca\xa8\xf4\xc0\x39\x21\x7b\x4c\xab\x98\x06\x49\xe2\x59\xbf\x29\x14\xe3\x2e\x29\x58\xea\x9a\x61\xc6\xb9\xaa\x61\x39\x84\xfd\x3c\xb4\x6a\x61\xd9\x09\x3a\x45\xbe\x14\x9d\xee\xb9\xd7\xe1\x00\x47\x45\xaa\x68\x12\x5a\x26\x89\x74\xd2\x9e\x85\xa5\x72\xc1\xa4\x44\xd7\xa2\x41\xe3\x37\x72\x88\x99\xd5\x89\x83\x2d\xe6\x76\x73\xee\xaf\x5b\x1f\x20\xf1\xba\x2f\xa0\xfd\x38\x83\x61\xaa\xbe\x53\xc9\xd3\x7c\x07\x92\x06\x76\xee\x3b\xd0\x56\x68\x6f\x99\x48\xe7\xf3\xde\x2b\xc5\xb2\xee\xdd\xf6\x80\xfc\xc1\x44\xd9\xbb\x72\x85\x5d\xbe\xf3\xcd\x22\x65\xb6\xaf\xbe\x5a\x94\x5a\xf8\xa5\x23\x6a\xec\x74\xf5\x9e\x5c\x9f\x2e\x8d\x8d\x6d\x90\xf4\x47\xc3\x54\x8c\x0d\x62\x2f\x0b\x5a\xa7\x02\x30\x25\x43\xdc\xa7\x51\x72\xc8\x30\x5d\x4f\x07\xbb\xb6\x31\xd3\xf6\xe5\xe4\x14\x8c\x5e\xec\xfb\xc3\x39\xe3\x0f\xe5\x07\xfe\xde\x4c\x5e\x6c\x92\x65\x72\xd9\x2f\x5c\x3e\x33\x99\x07\xcd\x50\x28\xf6\x32\x7b\x18\xb3\x44\x1a\x5b\x3e\x82\x73\x3b\x16\x62\xcb\xbf\x6f\x14\x14\x62\x2b\x81\x4e\xb3\x38\xd5\x0e\x49\x57\xc6\x02\x6f\x63\x81\xb7\x83\xf0\x84\x3a\xf3\x68\xa2\x55\xda\x42\x6a\xf8\x44\x57\x23\x69\x30\x4a\x5f\xda\x4c\x56\x85\xa8\x3c\x54\xf9\x37\xbd\xe1\xa0\xca\x01\x4b\x82\xdd\xb8\xcc\x5c\x71\x87\xea\x72\x73\x68\xcd\x99\xab\x5f\xae\x12\x3e\x5b\xbd\x05\xe3\x67\xb6\x32\x34\x48\x37\x8d\xa9\x2c\x6c\x5f\x15\x5b\x79\x5c\x28\xa7\xb7\x31\x1c\x20\x67\xe3\x7e\x2a\x6c\xff\x61\x90\xd0\x01\x0b\x34\x8d\x61\xec\x47\x10\xf6\x3d\x7a\x83\xd8\xcc\xc5\x1d\x82\x9a\xce\x4c\xf7\x62\x47\x13\xef\xa1\x6e\x2f\xf3\xc5\x58\x18\xa2\x99\xbb\x4a\xaf\x55\xfe\x0a\x7d\xfa\xda\xd6\xf8\xeb\x58\x45\x22\x4c\x67\xd6\x77\x09\x89\x81\xc9\x85\x9a\xb2\x7e\x91\x6c\x88\xe0\xd3\x2c\x61\x55\x33\xf9\x20\x31\xdf\x80\x5c\x8e\x75\xfc\x0e\x5a\xc7\xaf\x04\xc0\x3d\x45\xa8\xe8\x19\x06\xea\x95\x1b\xd0\x78\xde\x1f\xea\xc0\x59\x5f\x8f\x6d\x90\xa8\xca\x02\xdb\x97\x3f\xe4\xd7\xf6\x31\x69\x06\xfe\xff\xf9\x73\xaf\x82\x2a\x4f\xa0\x42\x40\x40\x7e\xd9\xe7\xe8\x5a\xf5\x5b\x3b\xdd\x0c\xd6\xcf\x26\x8c\x90\xac\xbd\xac\xc5\xcd\xb2\xc6\x0c\xf3\x4b\x67\x6b\x46\x75\x40\xe8\x81\x94\x93\xc6\xf6\x8c\x5c\xdc\xe8\x60\xc8\x30\x50\x43\xaa\x90\x15\x6a\x44\xd3\x29\xe2\x3e\xa5\x84\x6e\x22\xed\xa7\xad\xb0\x97\xe0\x7a\x4e\xb2\x76\x48\x81\x02\x40\xb3\x5a\x6d\xbd\xa8\x5f\x71\x4c\xad\xf6\x91\x32\xbf\x9b\xee\xfb\x56\x1a\x3f\x32\x87\x2f\x2f\x37\x7e\xf8\x08\xeb\xc7\x7a\xe5\x89\xe8\xab\xf5\xab\x67\x38\x40\xf1\xe3\x3d\x6f\xef\x40\x73\x7d\xf1\x69\xc0\x7c\xa0\xba\x48\x45\xa2\x95\x2e\x5e\x10\x72\x86\x96\x35\x2a\xca\x8e\xf1\x50\xcf\xda\xd0\x90\xa5\x23\xde\x45\x03\x7c\x99\x86\x29\x94\xc5\x79\x86\x8d\xe7\x4c\xa7\x64\x1e\x7d\xa5\xec\x29\x75\x75\xe4\xd8\xed\x6b\xce\x8b\xd3\xf3\x01\xa3\xbc\xd0\x7d\xaa\x04\xeb\x29\xa7\x29\x6d\xf2\x4e\xd8\x2d\x62\x02\xe8\x31\xdf\x5c\xf7\xd3\x5a\x84\xcd\x82\x3b\xe8\xc0\xf3\xd3\xed\x81\xe0\xab\xbc\xe4\x2a\xc3\x6f\xd8\x34\x75\xae\x6f\xb7\xc4\xda\x1a\x17\x38\x0b\x53\xa5\xb0\x75\x71\x6f\xec\x31\x9d\x49\x0f\xe9\x60\x5c\xbc\x79\xb0\xdc\x87\x73\x33\x79\x71\x6e\x95\x2c\x9f\xae\xba\x59\xa4\x79\xfb\x31\xb5\xb2\xa7\xab\xd5\xd4\xde\x47\x93\xed\x92\xdc\x72\xff\xfb\x4e\xe9\x8d\x34\x9c\xa2\xb7\x2b\x69\xb8\xc4\x86\x77\xd2\x68\x3d\x03\x46\x5b\xc8\xe1\x1a\xe3\x2f\xd5\x21\x89\xc8\xa8\x70\x28\x8a\x11\x31\x7c\x5c\xe6\x6f\x21\x8f\x17\xa7\xb9\x72\x7d\x39\x24\x9a\x16\xc8\x7c\x5e\x81\x08\xf2\xce\x09\x14\x96\xfa\xeb\x50\x98\xcb\x01\x09\x4b\x9d\xb3\xad\xb7\x94\x6f\xe6\x19\xf6\x8a\x94\xb1\xd7\xe0\x8d\xb8\xbd\x38\x0d\xde\x79\x8f\x2a\x13\x7a\x4c\xdd\x94\x49\xda\x54\xe3\x1b\x7f\xa8\x92\xbd\xa2\x0c\x80\x01\xb1\xf5\x75\xa4\xff\x4a\x05\x0a\xd5\x1c\xe4\xbf\x82\xb5\x1c\xad\x36\xe5\xe5\xbe\xc9\xaf\x2f\x5b\xe5\xbc\x0e\x92\x50\x1e\x34\xd9\x75\x97\x2c\xa5\xe1\x85\xa0\x4c\xa3\xec\x39\x2c\xb3\x68\x26\x99\x7e\x10\x68\x10\x17\xa2\x9f\x10\x8e\x97\xc6\x75\x20\xb1\xa4\x26\xc2\x2b\x57\xa4\x2a\x43\xa6\x89\x56\xe9\x77\x06\x53\xba\x3a\x4e\x0d\x81\xf4\xe5\xa1\xb9\xcb\x7f\x49\x64\x69\x38\x5c\x2e\xef\x93\x1e\x7e\x0f\x6f\xa7\xf4\x40\x6d\xd1\x57\x90\xf6\x26\x77\x61\x6a\x7f\x0d\x73\x7b\x29\x9f\xdb\x59\xd3\x88\x23\x28\xb5\x81\xb8\x26\xee\xb2\xa0\x75\x2e\xf0\x32\x11\x65\xa6\x4f\x70\x4b\xea\x70\x92\xc7\xff\xcb\x00\xe3\x34\x8f\x87\x9d\x8b\x7c\x61\xfe\x62\xc4\x84\x0d\xf3\x38\xd1\xb9\x2b\xf5\x1d\x2d\x1c\xab\x75\x66\xfe\x74\xf9\xfe\x2f\x0d\x44\x4f\x6e\xf5\x87\xcd\x32\x92\x5e\xd0\xe0\xb0\x17\x68\xea\xf5\x5a\x42\x51\xf3\x64\xf8\x66\x41\x41\x74\xcd\x98\xf5\x95\xbb\xab\xd2\x3f\x6a\xc4\x6d\x34\xe2\x9e\x15\xcd\xe2\x28\xee\x76\x71\xd7\x4b\x47\x97\xbb\x38\x4c\x7d\x8c\x87\x0a\xa3\x3e\xca\xa9\x15\xc0\x3f\x72\xe6\x7b\x55\x47\x35\x6c\xa3\x38\xed\x46\xff\xc1\xd6\xc8\x24\xea\x3f\x70\x7c\xfc\x58\xc1\xa4\x43\x05\x93\x03\xa5\xe4\x5e\xf5\x5c\x80\x41\x93\xd2\xb9\x9e\xa7\xe1\x9d\xf7\x3c\x58\x54\xd5\x8a\x3b\x25\xf7\x6a\x92\xbb\x44\x4d\x52\x91\xdd\xeb\x18\x95\x3c\xcc\x1c\xea\xe0\x75\xb9\xb8\xe9\x59\x06\xe9\x5b\x1c\x51\x47\xb9\x18\x18\x88\x33\x5b\x96\x8e\x1b\x30\x4f\xd3\x32\x66\xa1\x86\x6e\x07\xc2\x40\x7b\x60\x51\xcf\x42\x6a\x9e\xe7\xd8\x65\x0d\xfe\x53\x0f\xee\xfa\x29\x36\x49\x0d\xa2\x4b\x49\xd8\x47\xc1\xb8\x07\x87\xc1\x46\xb8\x2b\xc1\x5a\x07\x74\x6d\xb2\x73\xf7\x83\xad\x86\x78\x0a\x20\x31\x4c\x34\xfd\x01\xd8\x06\x7e\x96\x0e\xec\xaf\xf0\x86\xde\xea\xf6\x3a\x7d\x93\xbe\x36\xfe\x77\xaa\x46\x99\xc6\x25\xfc\x30\xb5\xc3\x8f\xc6\x9a\x1e\x74\x11\xe8\x15\x50\x8b\xd9\x60\x2b\xcd\x70\x85\x05\xfc\xfb\x5f\xaf\xdb\x9a\x0e\x48\x32\x8b\xec\x24\x6d\x1c\x8a\xb9\xf7\xc3\x14\xfc\xab\x8d\x74\xda\x32\xa1\x2e\xda\xd3\xb0\x8a\x0e\xca\xd8\x2c\x07\x50\x84\xaf\xc8\x7f\xe8\xd1\x5b\xc0\xea\xeb\x4a\x14\x56\x95\x30\x05\x0e\x90\xc5\x5c\x0f\x4b\x0d\xab\x5d\xd7\xe2\x12\xbf\xfb\xc4\xb1\x7b\xf5\x97\x29\x24\x8a\x73\xfc\x37\xe3\x3d\x14\xa5\x38\x27\xb4\x97\x71\x17\x58\x96\x3b\x95\x77\x19\x97\x43\x6f\xc9\xa3\x2e\x72\x91\x2f\x29\x6c\x3c\x40\x31\x8e\xe0\x01\xb4\x82\xea\x0d\x9b\x9e\x4e\xc3\x7a\x73\xfb\x16\x5e\x3d\x0a\x5b\x0d\x4c\x4f\x6e\xba\x7a\x43\x15\x2a\xca\xab\x3f\x14\x9b\x39\xb0\x87\x2d\x22\xef\x7b\xf3\xa9\x0d\xc7\x7f\x94\x85\x0d\xca\xc2\xec\xc7\xcc\x38\x8f\xb8\x78\x69\x71\xd9\x83\x23\x64\x4e\xa8\x2a\xa0\x10\xbc\x3f\x74\x2c\x8e\x63\x06\x29\x27\x8c\xb1\xd2\xea\x20\x54\x32\x3d\x9e\x61\x3e\xf5\xaf\x5a\x05\xad\x11\x51\x2a\x5a\xc1\x73\x96\x61\x93\x74\x13\xc8\xb4\x58\x58\x95\x24\xe3\xd1\x83\x26\x7d\x85\xd2\x4b\x6c\x9a\xc7\xbc\xa0\x75\xe9\x12\xc3\xec\xe5\x0d\x97\x58\x90\xd5\xb8\x70\xa1\x82\x7c\xdf\x63\xa1\x89\xd6\x85\x4e\x44\x1a\x96\x5b\x4c\x37\xf1\x42\x91\x7e\x89\x6b\xad\x15\xb4\x3a\x79\x09\x39\xeb\x13\xfb\x41\x1c\x67\xca\x5e\xd6\x0c\x24\x6b\x4b\x05\x27\x9b\xd5\xd4\x09\x7e\x15\x97\x00\x7e\x9d\x23\x8e\x96\xb9\x30\x9f\x7c\xd1\xdd\x3e\xe2\x1a\xc6\xc2\xb2\xc5\xb9\x30\x1e\xa4\xa8\x6b\xe8\xfc\xd1\x6b\x45\xdb\xd0\x05\xa4\xa2\xa6\x6d\xa8\x9f\xf4\xb9\x8c\x40\x7d\xc9\xae\x62\x20\x2e\x80\x81\x97\x6c\x4d\x6c\x40\xa2\x55\xda\x05\x30\xf4\x2e\xef\x96\x69\xe0\x59\xb8\x0d\x77\x56\xb0\xb0\xed\x12\x6a\xa4\xbc\x59\xe4\x19\xac\xb1\xef\xc8\xdf\x99\x7d\x5f\x81\xbb\xf7\x5c\xbf\xa5\x09\x89\xa9\x8d\xb9\x8d\x04\x70\x82\x1d\xf2\x5d\x07\x72\x9c\x2e\xe6\x48\xa7\x39\xbe\xa1\xe7\x20\x44\xe0\xe2\x69\x31\xaa\x9a\x4b\xf3\x13\x72\xcd\x2f\xbf\xdd\xd0\x7f\xa0\x9b\x09\xa1\x3b\xec\x10\x93\x21\x5e\x01\xed\x66\x62\xbe\x7f\xf3\x99\xc4\x08\xee\x2c\xed\xaa\x1c\x7e\xd5\x6d\xcd\x25\xd7\xcc\x33\xb9\xa1\xb3\xd9\x0c\xa4\x35\x9b\xcd\x6e\xe8\xfc\x4c\xcd\xe7\x53\xf2\xcd\x87\x60\x36\x62\x03\x95\x64\x4d\x2c\xd3\xcb\x62\x36\xdc\xd0\x20\x23\x8b\x6a\xcc\x3c\x13\xf9\xa8\x55\x1c\xb8\xcb\x2c\x52\xa0\xaf\x84\xda\xd8\x4c\xbe\x26\xe0\xd8\xe8\x65\xc8\x08\x5e\x22\xd7\x17\x12\xad\x40\xd7\xe6\xff\x0e\x9c\xa1\x1d\x76\xfc\x68\x07\x94\x49\x04\x94\xf9\x9b\x2d\x92\x64\xb3\x95\x02\x49\x86\xd6\x00\x36\xda\x30\x6f\x0b\x3c\x6c\x17\xd5\x90\x78\xf9\x91\xd9\x2f\x91\xcd\x40\xbc\x94\x08\xee\x88\x90\xaa\xc9\x07\x35\x6b\x7a\xa9\x02\xf4\x83\xc8\x57\xb8\x9f\xea\x19\x91\x87\x49\x97\x37\xac\x9a\x74\x39\x3d\x87\xe9\x55\xa6\xc7\xd1\x30\xaf\xbb\x87\xc6\x20\x4a\x9a\x67\xa3\xa0\xd6\x04\xa0\xcc\x70\x0a\x78\x1a\xcd\xc5\x77\xd4\x23\x05\xb4\xd7\xd7\x6e\xb0\xb6\xc0\xf9\x7d\xcb\xb7\x84\xf7\x9c\x9b\xe3\x90\x80\x49\xea\x64\xcc\xcf\x9a\x6f\xdf\xab\xd8\xb4\x3a\x21\xb9\xaf\x85\xeb\xf9\x93\x50\x3b\xb3\x22\xd5\xb9\x18\x1d\x0d\x62\x8e\xeb\xac\x6c\x5d\x60\x9d\x43\x9a\xf6\x33\x6c\xb8\xd1\xcb\x1c\xe4\xc3\x54\x41\xf3\xb3\x06\xbb\x2d\x50\xc1\x3a\x6f\xa5\x4f\x5a\x0c\x98\x65\x6b\x8d\x32\x90\x15\x55\x96\xa8\xe5\x56\x4a\xaf\x37\x95\x4a\xcd\xf0\xc7\xd5\xd5\xa2\x56\x4c\x6f\x80\x4e\x1d\x66\x64\x89\x3e\x2b\x1e\x70\x9c\x09\x5d\x6c\x01\x50\x35\x9b\x9e\x69\x12\xdf\xaf\xa3\xcd\x4e\x3e\x10\x47\x66\x9c\x1d\xca\xdc\x72\x32\xa4\x78\x1a\x8a\x8a\xb5\x19\x02\x09\x50\x13\x48\xb0\xd1\x2d\x91\x5b\x64\x31\xd7\xc5\x33\x74\x09\x80\x6e\x26\x36\x66\x37\x13\x14\x2c\x5c\xab\x63\x2e\xe3\x80\x08\x35\x90\x24\x8c\x26\xea\xb6\x84\xb0\x33\x03\xc7\xdf\xb5\x49\x72\xf2\xcd\x87\x68\x6b\xd9\xca\xfb\xf3\xb3\x26\x1b\x29\xb9\xd5\x97\x96\xe9\x4f\x72\xb3\x68\x71\x54\x2c\x93\x7d\xcd\xda\xd2\x25\xc5\x13\x8f\xc8\x92\xfb\x90\xa9\x55\xbe\x81\x74\xad\xf6\x89\x4b\x28\x71\xf5\xa7\xb7\x39\x69\x9b\xa5\xc2\x68\x19\x5e\xca\x67\xba\x14\x42\x0b\xe0\x8b\x48\xb3\xef\x3c\x27\xf0\x65\xc3\x79\x2f\xfd\x95\x90\x44\xfa\x12\xec\x6b\xa1\x3d\xd4\x0e\x87\x9f\xab\x6d\xc8\xf2\xf4\xdb\xb1\xf9\x47\x3b\x72\x30\x8a\x56\xb0\xc5\xce\x5a\x7d\x35\x19\x12\xd4\x6f\x5a\x2b\x25\x42\x72\x2c\x99\x49\x97\xe0\x01\x57\x93\x99\x5f\xad\xf4\xb3\x53\xb8\x5b\xfd\x5e\x9d\x47\x75\xc1\x66\xcd\x0e\x8d\x6b\x43\x97\x7d\x86\x33\x07\xfe\x0c\x49\xea\x33\x73\x1b\xd7\x86\x66\x84\xa6\x86\xb8\x60\x0e\xe4\x97\x95\x2f\x79\xa6\xb6\x9a\x72\x68\x50\x00\xcb\xfb\x33\x94\xed\x27\xb3\x6e\x35\x6b\xab\x95\x47\x0c\x2b\x2c\xd5\x92\xe0\x57\x46\xba\x95\x5f\x21\x0c\xe7\xcb\x08\xc1\x89\xb0\xb6\xe0\xe2\x26\x89\xbb\x2b\x6f\x2e\x2f\x92\xfc\x5b\x2f\xf3\xc5\xcf\xff\x0b\x00\x00\xff\xff\x07\x45\x3c\xfb\x76\x39\x01\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
