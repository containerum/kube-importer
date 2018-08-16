// Code generated by fileb0x at "2018-08-16 15:47:57.79066898 +0300 MSK m=+0.023823208" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(7f0a49b87726c79bbd71842d3afab33e.d47d734c826bc110696e7160789432d0)

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
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5d\xdd\x8f\x1b\x37\x92\x7f\xf7\x5f\x41\xe8\x0e\x98\x04\x6b\x69\x7c\xc9\x62\x81\xf3\xdb\xc4\xb2\x63\xc1\x1e\x9f\x30\x1f\xc9\x01\x3b\x86\x40\x75\x97\x24\xc6\xdd\x64\x9b\x64\xcb\xa3\x2c\xfc\xbf\x2f\x48\xf6\xb7\xfa\x4b\xdd\x6a\x0d\x27\xab\xa7\xc4\x9a\x6a\xb2\xc8\x2a\xfe\x58\x2c\x56\x15\xff\xf5\x02\xa1\x91\xc3\xa8\x08\x7d\x10\xa3\xd7\xe8\x9f\x2f\x10\x42\x68\x84\x83\xc0\x23\x0e\x96\x84\xd1\xcb\x3f\x04\xa3\xa3\x17\x08\x7d\x7e\xa9\x68\x03\xce\xdc\xd0\x69\x47\x2b\xbe\xe1\xf5\x1a\xf8\xe8\x35\x1a\xfd\x34\x79\x35\xd2\xbf\x11\xba\x62\xa3\xd7\xe8\x5f\xe6\x5b\x17\x84\xc3\x49\xa0\xbe\x55\x54\x37\x20\x58\xc8\x1d\x18\xdf\x02\xdf\x12\x07\x10\x11\x48\x44\xff\xbb\x62\x1c\x11\x3f\x60\x5c\x12\xba\x46\x5f\xc2\x25\x70\x0a\x12\x04\xe2\xd1\x47\x02\x49\x86\x98\xdc\x00\x8f\xbf\x11\x13\xdd\x27\x42\x23\x49\xa4\x07\xaa\x87\x0f\xe1\x12\xc6\x33\xdd\x0c\xf0\xf8\xaf\x5b\xe0\x22\xe2\xe0\xd5\xe4\xd5\xe4\x7f\xd4\x08\xbe\x9b\xd1\x62\xb9\x11\x29\xbb\x97\xd8\xf3\x92\x7f\xa9\x3f\x33\x21\x33\xff\x56\x1d\xe1\x75\x3a\x35\xd1\x6f\xa6\xbb\x51\xf2\xd3\xe7\x97\x29\xbd\x08\x7d\x1f\xf3\x9d\xea\xda\x90\xa1\x2d\xf3\x94\x2c\x62\xd6\x35\x15\x0b\x80\xeb\x09\x9e\xb9\x29\xe5\x95\xe7\xbd\xc7\xd4\xf5\x92\x71\x68\x52\x0e\x22\x60\x54\x80\xc8\xf1\x85\xd0\xe8\xa7\x57\x3f\x15\x7e\xda\x9f\x7f\x33\xbf\x6a\x46\x43\x4f\x66\x5a\x35\xac\x3a\x1b\xf0\xf1\x5e\x1b\x08\x8d\xfe\x9b\xc3\x4a\x7d\xfe\x5f\x97\x2e\xac\x08\x25\xaa\x39\x71\x69\xb8\xbc\x89\xf8\xb9\x63\x12\x7b\xa3\xdc\x97\xdf\x33\xff\xfa\x9e\xed\x6c\xe4\xc2\x0a\x2b\x0e\x9a\xf8\x75\x36\xc0\xf9\x0e\x01\xe7\x8c\xf7\x65\x97\x14\x26\x75\x6a\x78\xf8\x85\xb9\xbb\x1a\xb6\x5f\x94\x0c\x60\xf4\x38\xf6\x41\x6e\x98\x3b\xde\x12\x41\x96\xc4\x23\x52\x0b\x38\x08\x97\x1e\x71\xe2\xc6\xcc\xa7\xd1\x67\xa3\x4b\x87\xd1\x15\x59\xfb\x38\xc8\x0a\x6e\xb4\x86\x16\xfa\xf5\xf6\xb1\x95\x7e\x19\x32\x64\x3a\x42\xaa\xa7\x3a\x1d\x33\xd4\x6f\x34\xf1\x35\x0e\xc4\x47\x22\x64\x4b\x3d\x7b\xd5\x28\x37\x78\x3c\xa2\x9e\x15\x78\x3c\xab\x58\xaa\x62\x2f\x87\x04\xaa\x54\x63\x9b\xb1\xaa\x93\x1e\x3d\x29\x5e\x9d\xf5\xa8\x12\xaa\x5c\x08\x3c\xb6\xf3\x81\xca\xa1\xb1\x8a\x62\x1f\x44\x80\x1d\x68\x01\x55\xd3\x94\x2b\x7b\xb1\xaa\xc8\xe4\x59\xc9\x4e\x04\x56\x19\x9d\x6d\x46\xab\x6e\xaa\x74\x86\x2b\x3b\xe1\x8a\xd0\x35\x07\x91\x17\xd9\x10\x60\x95\xf4\xd3\x8c\x55\xb3\x98\xd4\x5e\xa4\x8a\x4e\x60\xbf\x13\xb9\x99\x63\x8e\xfd\x33\x5c\x9d\x12\xae\x5a\xe9\x92\xa1\xed\xa2\x4b\x67\xa8\xb2\x13\xaa\x52\x7b\xc7\x22\xc3\xea\x53\x42\x6b\x2f\x5a\x15\x78\x3c\x6b\xd8\x89\x70\xaa\x9d\x1e\x19\xe2\x4e\x7a\x74\x46\x2a\x3b\x91\x2a\xf6\xb0\x0e\x8c\x53\x45\x47\x6e\x35\x4a\x45\xf6\xca\xd9\xa2\x3a\x23\x55\x29\x52\xb5\xd1\x24\x43\xda\x41\x93\xce\x28\x65\x29\x4a\x49\xc6\xf1\x7a\x78\x94\x8a\xba\x69\x81\x52\x11\xa5\xc5\x28\x65\x38\x3c\x83\xd3\x49\xc1\xa9\x85\x02\x45\xe0\x74\xb8\x02\x9d\xc1\xc9\x4e\x70\x8a\x2e\x7a\x07\xc6\xa6\x16\xd7\xc9\x86\xf2\x37\x43\x68\x2f\x32\x65\x19\x3c\x6b\xd5\x89\x90\xa9\x75\x34\xc2\xe1\xea\x73\xc6\x25\xbb\x70\x29\x09\x7a\xc9\xf0\x92\x86\xbe\x5c\x39\x0e\x08\xf1\x11\xb6\x90\x8b\x80\x91\xbb\x40\xc7\xd5\x08\xc9\x09\x5d\x27\x23\x1f\x3d\x8e\xd7\x6c\x1c\x60\xe7\x0b\x5e\xeb\xbf\xaf\x89\xdc\x84\xcb\x89\xc3\xfc\x4b\x87\x51\x89\x09\x05\x1e\xfa\x97\x5f\xc2\x25\x8c\x49\x14\x8e\x73\xb9\x05\xea\x32\x7e\x59\x47\xeb\x78\x04\xa8\xbc\x0c\xbe\xac\x2f\x7d\xe6\x42\x14\x51\x12\x63\xaa\xb9\x66\x7e\x47\x74\xac\xcf\x1e\x8f\x6c\xf9\x07\x38\xa9\x1e\x8d\x02\xae\x34\x59\x92\x82\x72\x8e\xcc\xc5\xf6\x62\x95\x6f\xa6\x6e\xb8\xe9\x90\x29\xf6\x35\xc1\x27\xf5\xdf\x52\x91\x44\x4d\x74\x6b\xf6\x4e\x11\xbe\x28\xca\xfc\xbb\x15\xd3\x7e\x8d\x83\xec\xac\x17\x16\x48\x42\x83\xc6\x63\xa4\xdb\xd0\x11\x5d\x69\x34\x4a\x2a\x97\x2a\x79\x71\xf8\x1a\x12\x0e\x6e\x0e\xd1\x46\x7a\x6a\x32\xf3\xeb\x62\x89\xe3\x29\xfa\xdc\x2c\x6b\x0e\x58\x82\xbb\xc0\xc5\x25\xbe\xb7\xc0\x15\x21\x61\x14\xb9\x58\x02\x22\x14\xdd\xbc\x7b\xf3\xf3\xcf\x3f\xff\xaf\x1a\x85\x8f\xf3\xe8\xd4\x5e\x9e\x6f\x4c\xf7\x57\xb2\x5c\x57\xf4\x58\x0a\x7c\x35\x04\xc1\x4c\x33\xc3\x2f\x34\x06\x1e\xb4\x1a\xaa\x21\x3c\xee\x40\xa7\xa6\xf3\xaa\x81\x46\x64\xc7\x5e\x6b\x89\x7b\xb2\x57\xdb\xa6\x85\xd2\x0e\xd8\x37\xaa\xe3\x1b\xbb\x35\xfe\x7f\xfa\x6b\x7b\x97\xf3\x34\xaf\x7e\x95\x4b\x5a\xd1\x55\x2d\x6b\xa4\x75\xb8\x71\x6d\x63\xd7\xd5\x9a\x8c\xbd\x79\xc5\x4a\x2d\xcc\xaa\x6d\x73\x65\x2c\x9f\x16\x93\xa5\x09\xab\x66\x4b\x20\x2f\x6b\x40\x75\xdc\xba\x0a\x51\x84\xb9\x96\x30\xe7\x78\x97\xd7\x49\x22\xc1\x17\xfb\x16\x4e\x03\xd0\x8c\x2a\x6d\xa6\x3c\xc2\x25\xe3\xb6\x52\xd1\x0d\x79\xbd\xdc\x0c\xcd\x9e\xc8\xa2\x9f\x09\xcd\x04\x97\x74\xdd\xc4\x88\xaf\x06\xfe\xb2\x7a\x57\xf3\x88\x4f\xa4\x38\x60\x5f\x63\xbe\x8f\xa9\x7b\x04\x35\x28\x5f\x77\x8d\x52\x8f\xba\x2f\x45\xcd\xc8\xc0\x1a\x58\x4d\x8d\x7c\xcc\xb9\xa4\x8f\xb2\x66\x18\x07\xba\x1d\x8a\xe1\xb7\x74\xdb\x92\xc9\x1c\x65\x96\x3b\xa3\x44\x5d\x37\xa3\x99\xfe\xba\xb4\xe1\x48\xf9\x5a\x19\x22\x71\x08\xff\x49\xf7\x79\x05\x08\xc3\xab\xd2\x3c\x77\x6e\xae\x95\xd1\x5c\x33\x54\xca\xab\x39\x51\x2f\x7c\x16\x52\x69\x99\xfa\x1b\xe2\x6b\xc3\x99\xbd\x68\xad\xe5\xd0\x02\xb1\x15\x5d\x1e\xb5\xa3\xd8\x9a\x14\xbd\x8f\x75\xe8\xd0\x2e\x95\xec\xbf\x39\x93\xcc\x61\x5e\x7b\xc0\x1e\x70\x65\x54\x35\x4b\xa8\x84\x35\xe4\xdd\x15\xa3\xc8\xca\x37\x7f\xfe\xc7\xdf\x47\xf5\x3a\x5e\xd1\x69\x3c\xfa\x56\x88\x31\x2f\xcc\x95\x85\x1a\x17\xad\xa2\x16\x3a\x67\x28\x95\xd6\x99\x75\x8e\x7e\xc8\xd9\x77\x3f\x22\xbd\xee\xc1\x3d\x8a\x16\xea\xb6\x16\x01\x96\x9b\xf6\x7a\xa6\x86\xd9\x59\xcf\xae\xd5\xc7\xa5\x22\xcf\xb0\xd2\xbd\xf1\x90\xca\x79\x66\x30\x27\xd9\x3a\x44\xb8\xec\xc7\xf6\x6d\xb8\x2c\x30\x6d\x8b\xf6\x6a\xdf\x42\xb3\xea\x66\xc8\xd0\x78\xdc\xfd\xec\xe1\xe1\x25\x78\x9d\xa7\xf1\xa3\xfe\xba\x5c\x44\xe6\xce\xab\xbb\x84\xa2\xef\xcb\xdd\x71\x98\x93\xd5\x6a\x41\xdc\xee\x3e\x39\xdd\xc2\x6c\x6a\x9f\x06\xa4\x71\xec\x35\xf2\x4f\x89\xf2\x7b\x65\x36\xcb\xa3\x23\x3e\x25\xec\x8a\xba\xa3\x0d\x07\x9d\x3f\x7a\xc0\xe1\x06\x3b\x92\x6c\x2b\xf5\x61\xc9\x98\x07\x98\xd6\x48\xec\xca\x7c\x5f\x75\x36\x89\x79\x1e\xda\x36\x6b\x7f\x28\x89\x39\x2a\xe7\xd8\x6e\x17\xa6\xc5\x5e\x47\x7d\x60\x5a\x04\xa1\xe7\x2d\x04\x38\x1c\xe4\x13\x9d\x96\xf5\xd1\x6b\x1e\x7a\xde\xad\xe6\x42\x9c\x7d\xa4\xd5\x3e\xd2\x32\xe0\x18\xc6\xbe\xbd\x29\xe0\x52\x61\x53\x62\x5e\xa8\x54\xb7\x64\xe7\x28\x28\xf6\x6d\x44\x89\x66\x53\xf4\x03\xa3\xde\x0e\x91\x55\x06\x5d\x11\x11\x28\xc0\x5c\x22\xb6\x42\x71\xa3\x3f\x76\xd4\xf7\xb8\xab\xdc\x66\x94\xdb\x49\xb1\x0c\x5b\x1e\xe4\xd3\x7d\xe1\xd6\x7c\x55\xbe\x7d\x32\x89\xbd\x85\x13\x84\x0d\x93\xa0\xe9\xd0\x9b\xf9\x3d\x0a\x05\x5e\x03\x5a\xee\x10\xf6\xbc\xd4\x04\x16\x6a\xcd\xcb\x0d\x11\x65\xce\xb4\x43\xa4\x1a\x36\x89\x55\x27\xa2\xbf\x99\xdf\xd7\x8d\xc7\x07\x9f\xe9\x6b\xf0\xe6\x21\xdd\x5c\x5d\xdb\x31\xa4\x6b\xc3\x73\xb9\xcf\x21\x29\x70\xd0\x42\xf0\xbf\x45\xc4\x16\x9b\x33\xb7\x45\x3d\xae\x36\x6a\x0c\xa9\x32\x6d\x32\x25\x23\xcc\x3a\x50\x2b\xee\x10\xd7\x6d\xa5\x39\xb2\xc5\xc4\xc3\x4b\x0f\x16\x03\x23\xd2\x55\xdc\x51\x3d\x34\x71\xc0\xee\x6e\x31\x38\x3a\x62\x77\xd7\xc4\xc7\x13\xe2\x73\x48\x4f\x26\x96\xfb\xb4\xab\x06\x9e\x02\x57\x5b\x6a\x43\xf3\x63\xba\x29\xe1\xc5\xba\x85\xfc\xdb\x1e\x30\x55\xaf\xe4\x88\xb6\xea\x94\x82\x22\x90\x43\x66\x96\xbb\x2f\xe7\x0a\xb0\x6c\xbd\x09\x3f\x03\xfc\x6c\xba\xac\x2c\x50\x56\x1e\x0c\x7b\xde\x56\x96\x17\x12\x38\xea\x61\x2b\x73\x04\x6e\x67\x8b\x67\x86\x6e\x9f\x04\xdf\x52\x37\x60\xa4\xf6\x38\x1f\x93\xe4\x65\x06\xf1\x87\x47\xf2\x77\x63\xd7\x8d\x72\xaa\xf7\x6e\x61\xda\x9f\xe2\x93\x36\x9e\xe6\xd4\x75\x95\xf4\xff\x0c\x8f\xd5\x03\x1d\x06\x07\x3b\xab\x0d\x7a\x43\xd7\xe3\x62\xee\xd8\x0b\x7b\x92\x21\x9a\x50\x90\x97\xce\xc6\xd0\xe2\x80\x34\xad\xea\x26\x54\xce\xd1\x95\xaf\xef\xbe\x88\x9c\xb4\x33\xdc\x35\x77\x04\x44\x6d\xef\xba\x63\x7e\x6c\x13\xd9\xb6\x56\x50\x5b\x7d\xe0\x80\xdd\x78\x8b\xbd\x10\x50\x80\x09\x57\xa7\x0d\xa0\x5b\xc2\x19\x35\xe6\x0a\xe6\x44\x99\x8c\x9d\xbd\xab\xba\xe9\x3d\xc7\xea\x53\x5f\x38\x1a\xae\x3a\xdb\x4e\xfa\x6b\xfb\xf6\x5d\x5e\x17\x22\xf4\x96\xeb\xe0\x20\x21\x31\x75\x31\x77\x91\x00\x4e\xb0\x47\xfe\x54\xd2\x45\x57\xf3\x99\x89\x0a\x7f\xa0\xd7\x20\xb4\xa3\x60\x3c\x46\x0e\xa3\x8a\x5c\x9a\x3f\x21\xdf\xfc\xe5\xf5\x03\xfd\x1b\x7a\x18\x11\xba\xc5\x1e\x71\x51\x28\x80\xab\x89\x79\x18\x99\xdf\xbf\x86\x4c\x62\x04\x8f\x0e\x80\x0b\x6e\xfc\xab\xa6\x35\x1b\x90\xe9\x67\xf4\x40\x27\x93\x09\x48\x67\x32\x99\x3c\xd0\xd9\x54\xf5\x17\x52\xf2\x35\x84\xa8\x37\xe2\x02\x95\x64\x15\xd5\x6c\x44\x0e\x73\xe1\x81\x4e\x41\x62\xe2\xe9\x83\x32\x0b\x4c\x24\x9e\x76\x65\xc0\x63\x81\x49\x81\xbe\x10\xea\x62\xd3\xf9\x8a\x80\xe7\xa2\x8b\xf8\xa8\x71\x81\xfc\x50\x48\xb4\x04\x44\x19\x1d\xff\x09\x9c\x21\xad\x0f\x31\xaf\x94\x49\x04\x94\x85\xeb\x0d\x92\x64\xbd\x91\xba\x64\xe3\x0a\xc0\x45\x6b\x16\x6c\x80\xc7\x74\x71\x49\x47\x74\xf1\x2b\x73\x2f\x90\xcb\x40\x5c\x48\x04\x8f\x44\x48\x45\xf2\x4e\xf5\x9a\x67\x55\x80\x76\xa0\xe5\x97\x9c\xe8\x63\x9a\xea\xe9\x78\x22\xd3\x24\x12\x46\xf9\x0a\xd3\x73\xde\xd2\x89\x67\x66\xaa\xc2\xfb\xed\xb6\x6b\xe3\x2d\xe7\x55\xfe\xc4\x48\x25\xba\x5f\xf9\x46\xdf\xd7\x38\x2b\x17\x1b\x29\x83\x81\x0e\xca\xc6\x37\xf4\xfe\xee\x6e\x7e\x62\xc8\x31\xd9\x22\x7b\x18\x33\x9b\xd6\xa3\x8c\x59\xcb\x1c\x02\x0e\x42\x1f\xbf\x72\xcb\x7a\x36\xed\xae\xee\x6a\x49\xb7\xd6\x86\x0f\x8a\xb8\x5c\x64\x07\xe8\xd4\xed\xc9\x6f\x4c\x2b\x66\xfd\x43\x7e\xec\x25\xf3\xae\x28\x0a\x33\xaf\x26\x4c\xef\xe8\xb9\x84\x9f\x4a\xc5\xac\xf6\xe6\x9e\x70\xa0\xb7\x4d\xfa\x75\xbb\xa7\x60\x71\xf1\xdd\xd9\xb4\x66\xb0\x27\xc8\xfd\x29\x1b\xd1\xbb\x22\x0e\x16\x47\x94\x6e\x13\x99\x11\xa5\x81\xe6\x99\xcd\xc2\x40\x6a\xcd\x10\x9f\x45\xb8\x7a\xd9\x24\xcd\x0a\x31\xa1\x07\x23\xc3\xa7\x66\x33\xb1\x14\x0a\xee\xf0\xba\xf5\x67\x96\x58\x79\x85\x0c\xc1\x6a\xc5\xca\x13\x1a\x05\x8b\xfe\x1f\xaf\x24\xf0\x4c\x49\x6a\xc3\x5d\x77\x60\x5e\x61\xe2\x81\x3b\xd4\x69\x2c\x19\x48\xe8\xb5\x3d\x91\xbd\x33\x0c\x55\xdc\xa7\x6b\x51\xd8\xc4\xee\x2c\x66\xc9\x72\x6d\x33\xe5\xb1\xdb\xaa\x9c\xa6\x56\x7a\x97\x66\xc8\x1a\xf5\x5b\x31\xae\x2f\x21\x93\xf2\x3c\x47\x40\xb2\x43\xd2\x69\xad\x99\xd4\x7c\xce\x6e\xe5\x6c\x86\x9e\xcc\x4f\xa3\xfa\x41\x4d\x22\xa3\x90\xac\xe2\xee\xab\x77\x50\xfb\xf8\x59\xc5\x83\xd8\xa2\x1f\xa6\x58\x62\x9d\x6a\x18\x8a\xbc\x7b\x2d\x2a\xc7\x78\x2c\xef\x39\x0f\x3d\x10\x7f\x99\x1c\xd5\x73\x5a\xe9\x33\x0a\x99\xd2\x9a\x37\xd0\xf6\x7c\x13\x7a\x6d\x33\x4d\x6e\xb2\x2b\xc0\x3a\x78\x80\x26\x4f\x7c\x8e\xae\x14\x2a\xa0\xaf\x27\xbe\xac\x6a\xf1\x71\x8d\xa9\x08\xd3\x5a\xda\x51\x45\x6a\x5b\x44\xf6\xa9\x64\x2d\x16\xc5\x95\xd0\x28\x51\x25\x8b\x37\x3d\x15\x6a\xd0\xec\x0a\xee\x89\xad\x7f\x48\xfc\xb2\x93\xdf\x84\x6a\xc4\x94\xad\xbb\xf1\x0c\xef\x29\x7b\xc4\xd5\x57\x39\x1d\x4b\xf3\x0c\x0a\x63\x0d\x05\x70\x53\xf3\xc4\x03\xa4\x3f\xd0\x8b\x53\x6e\x20\x95\x7f\xc7\x01\xd7\x24\x2a\xf8\xf8\x71\x01\x8f\x72\x11\x59\xdf\x1d\xbd\x96\x8d\xe1\x7e\xd7\xf8\xf1\xed\x63\x5c\xa6\xb0\x9a\x13\x42\x4f\xc1\xc9\x8c\x36\x73\x22\x39\x5e\xad\x88\x33\x20\x17\x77\x51\x0f\xa7\xdd\x4d\x75\xc3\x0b\x8f\xad\x09\xed\xd7\xfc\x47\xdd\x44\x45\x10\x5d\x0c\x2f\x07\xa5\xdf\x8a\x13\x66\xbb\x64\x23\xdc\xc4\x80\x89\x13\xf7\x02\xb8\x81\xc3\x96\x9b\xd6\xbd\xc8\x27\x4d\xd8\xb6\x65\x35\x99\x19\x79\xc2\xbc\x9d\x91\x56\x5e\xee\x69\x68\x94\xd6\x1c\x3f\xaa\xdc\x4a\x8c\xdd\x5a\xb1\xa5\xc3\xb6\x4f\x76\xf3\x9b\xe6\x90\xc5\x84\x26\xe7\x77\x46\x73\x0e\x37\xe0\x01\x16\x80\xe2\x36\x3a\x0b\x6d\x26\x3e\x85\x7e\x53\x02\x54\xe9\x0a\x8d\xba\xae\xf9\xbc\x1d\x1c\xd7\x35\x7e\x2b\xf9\x93\xfb\x7e\x97\x1e\xa6\xeb\x4b\x01\xfe\x36\x46\xef\x44\x84\xf9\x6c\xcf\x3d\xe9\x61\xb9\xd1\x3e\xa1\xc8\x11\xa0\x93\x43\x3b\xda\x86\xb9\x6f\x11\x1a\x45\xdb\xf1\xa2\xe8\x10\x88\x7f\xcf\x56\xe6\x6b\x36\x23\x7b\xa5\xad\x56\x27\xda\xe6\x98\xec\x9c\x73\x69\x1a\xa9\xc9\xbb\xcd\x0e\x79\xa0\xeb\x5d\xd3\x45\x21\x4f\xdd\x16\x20\x51\x4b\x45\x48\xa0\x51\xe5\x44\xb3\xb1\x5d\xe7\xb3\xb3\x2d\x2b\xaf\x37\x67\x75\x17\xa5\x73\xe6\x16\xab\x2d\xb8\xbd\x2a\x19\x9d\xf3\x30\x8f\xea\xa6\x0b\x3c\xb6\xeb\xbc\x9e\x4d\x58\xf4\x39\x8d\xf2\xf4\x91\xb3\x87\xe4\xec\xcd\x99\x7b\xda\x64\xbd\xec\x12\x7f\xee\x59\x7a\xc3\x8e\x65\x2f\x3d\xcf\x96\x7d\x28\x51\x99\x5a\x64\x6f\x48\xa3\xeb\x05\xf5\xc1\x06\x8b\xee\x4b\x72\xae\xbf\xae\x3a\x3b\x4b\xcc\xe5\xc2\x61\x21\x95\x83\x65\xa1\xe9\x3e\xde\xe8\x2e\xaa\x16\x30\x97\xcd\xbb\x42\xc0\x5c\xa4\x49\x8f\xbb\x2d\xdc\xaa\x26\xaf\xa4\x95\x8a\xd7\x74\xfe\x8d\x49\xf6\xcc\x8a\xbe\x67\x5e\xd5\xc4\x70\xc9\x08\x6e\xeb\x5c\x04\xd7\xc6\x13\x6e\x7d\x75\xad\xfd\xa2\x5a\x71\xce\x00\x6a\x17\xe9\x72\x2e\xa8\xf5\x17\x28\xa8\xd5\x25\x67\x62\xbe\x3f\x88\xa2\x72\x65\xfc\x25\x3a\x20\x23\x1e\x37\x52\x33\x8c\x7e\xb8\x7b\x33\x47\x8c\xa3\xfb\xe9\xfc\xc7\x27\x89\x42\x6c\xb1\x7a\x92\xf2\x87\xd5\x83\x8c\x49\x0a\x81\x89\xe9\xbd\x94\x32\xbc\x30\x75\x95\xd1\xd2\xb9\xf2\x4e\x10\x8e\x72\x71\xda\x59\xe3\xa3\x45\xc4\x41\xa3\x89\xa8\x78\x24\x14\xf9\xc3\x58\x4c\x95\x86\x5f\x2b\x93\x4f\x19\x7b\x84\xa2\x6b\x32\x0c\x73\xb6\x5a\x72\xb1\x5e\x99\xf4\xef\x42\xdc\x66\x95\x12\xa6\xc4\x51\x5a\x8a\xea\x51\x20\x0a\xdf\x92\xf0\x27\x94\xc3\xe6\x27\x2e\x0b\x66\xe5\x5c\x67\x2e\x26\xda\xce\x78\xfa\x49\x6e\xde\x43\x01\x1c\x99\xeb\xe1\x96\x85\xb3\x4f\x72\xc7\x1c\x67\x22\x75\x16\xe3\x7d\xdc\x80\xb5\x92\x14\x6d\x82\xa4\x25\x91\x1e\x64\x85\x29\x2a\x41\xdc\xd4\xcc\xd5\x38\xae\x85\x9a\xdc\xe3\x4d\x9a\xb0\x7b\x83\xb9\xdb\x1e\xaa\x35\xf5\x11\x2a\xf3\x86\x02\xba\xb7\x63\x8b\x2c\x43\xaf\x16\xf3\x42\x0f\xb2\x9e\x7d\xae\xc8\x3b\x6e\xb0\x1b\x96\x7f\xe7\xe6\xb0\x22\x9c\x9b\xe2\x4b\x3d\x87\xac\xa5\xf7\x4c\x54\x99\x89\x35\x17\x02\xbd\x8f\x35\xb9\xfb\x82\xfa\x73\x4d\xe5\xd5\x82\xf4\x44\x83\x8b\xb0\xf9\xfe\xf9\xe3\xad\x71\xef\xd9\xa8\x7e\x34\x2e\xc0\xd5\x22\x89\xa0\x84\x3a\x97\x49\x20\x19\xe2\x21\x4d\x4a\x82\x21\xa5\x87\xd0\xeb\x6d\x00\xe3\x20\x1e\xe8\xf0\x12\xb9\x9f\x2b\xea\xa6\x73\xce\xf8\x53\x25\x53\xbe\x35\x9d\x97\x3b\x83\x99\x5c\x0c\x3b\x2f\x9f\x98\xdc\x9f\x1a\x4b\x14\xf6\xb6\xb8\x14\x8b\x3a\x9a\x1e\xf2\xa3\x55\x7b\x7e\xa3\x67\xff\x62\xa5\xe4\x8d\x9e\x8a\xd9\x69\x97\x70\xd1\x23\x6d\xf8\xfc\xf6\xcf\xf9\xed\x9f\x63\x40\x42\x93\x63\x36\x43\x95\xf7\x05\x1a\x98\xe8\xeb\x9e\x8d\x5a\x19\xca\x94\x29\xda\x0f\xb5\x6b\x6a\xff\x2e\xd1\x1a\x49\x79\xe0\x48\x70\x5b\xbf\x40\x54\xfe\x41\xfd\x4b\x44\x68\xc5\x99\xaf\x2f\xcc\x32\x61\x62\x83\xe5\x94\x15\x86\x62\xd9\x44\xb7\x4d\x0e\x28\xa5\xaf\x4b\x12\x38\xed\x24\xe7\x87\x61\xcd\x1c\x17\xe3\x94\x6b\x5c\xc3\x71\x50\xf3\x11\xaf\x1a\xce\xd9\x58\x27\xd8\xe8\x07\x0c\x41\x71\x99\x8f\x7b\x04\x61\x4f\xcd\xe7\xe5\xd1\x2d\xc1\x53\x1d\x5c\x66\xf3\x73\x25\x70\x0b\xaa\xcb\x95\x86\x11\x76\x7d\xfd\xa9\x67\xd9\xf0\xb8\x1e\xc7\xd0\x35\xc3\xed\xda\x16\x1a\xae\xa4\x93\xad\x21\x99\x9e\x63\xde\x44\x4b\xcc\xd7\x20\x17\xe7\x0b\xe9\xa3\x5e\x48\x57\x4c\xf0\x30\xe3\xb8\xd3\x3d\xd8\x19\x07\x1c\xa9\xf8\x5d\xfe\x55\xe8\x66\xeb\x47\x5f\x89\xdb\x7a\x0b\x1e\x0d\xea\x77\x22\x37\x73\xcc\xb1\x5f\x6b\x2d\xe7\x49\x8b\xe7\x49\x33\xd8\x6f\x44\x6e\x90\xd9\x3e\xce\x56\xdf\xd9\xea\x1b\xde\xea\xdb\x10\xd7\x05\xda\x30\x9b\x1b\xe2\x42\xa2\xa2\xfa\x10\x67\xf2\xdf\x5e\x76\x7b\x96\xe8\xbd\xe9\xf3\x6c\x85\x9e\xad\xd0\xb3\x15\x6a\x4b\xd4\x5a\x71\x7f\x6a\x74\xfe\xec\x93\x97\x6e\x69\xfd\x7d\xa4\x51\x85\xa7\x61\xf5\x2d\xdd\xc2\xdb\xba\x4b\x23\xb6\xac\x94\xa2\x68\x27\x3d\xf1\xdc\xa5\x66\x8f\xb0\xba\x1a\x90\x11\x1c\xd4\xc9\x2a\x46\xa6\xf1\x18\xf1\x90\x52\x42\xd7\x09\x0a\x75\x35\x12\x25\xf8\x81\x97\x7d\x3c\xa4\xc4\x70\xa4\xc5\xdd\xa5\xd9\x78\x5c\x72\x4c\x9d\xee\xa9\x9a\xbf\x98\xcf\x0f\x7d\x91\xfc\xc4\x97\x8f\xd5\xcf\x92\x1f\xbf\x04\xc8\x73\x35\x3d\x12\xfd\xea\x1c\xfd\x12\x37\x50\x1e\xc3\xc5\xbb\x07\x5a\xde\xdf\x7c\xb4\x17\x06\xea\xcb\xed\x67\xa8\x74\x1d\xf6\x18\x18\x3a\x56\xdb\xaf\x7e\x1b\xc1\xce\xa5\x66\x99\xac\x3e\xea\xc0\xc7\x16\xe2\x32\x84\x39\x89\xa5\x75\x53\x4d\xf8\x64\x8f\x60\xa7\xfd\x08\xf6\xc1\x83\xd3\xf3\x8e\x8e\xa3\xf6\x79\x73\x75\x6d\xaf\xc4\x4b\x43\x68\x2b\x84\x9e\x0b\x9c\xd5\x97\x9f\x99\xc3\x43\x2a\xff\xee\x82\xaf\xac\x97\x73\xbc\xb5\x5a\x26\xcc\x12\x7f\x64\xcb\x4a\x78\x25\x05\x7b\x2c\x93\xef\xdd\xfe\xae\x55\x25\xde\x98\x34\xb7\xaa\xbf\x6d\x88\xb3\x31\x51\xd0\x0e\xa6\xca\x58\xeb\x13\xd6\x3e\xdc\xd3\xcf\x03\x18\x2a\x3a\x35\x5e\x3c\x65\x3e\x7c\x85\x07\xc0\x2b\xa2\x74\xdd\xa1\x22\x8f\xec\xa7\xb4\xc8\xfe\xb2\x16\x4d\xe3\x49\x34\x4b\x97\xc5\xca\xc3\x8f\x3a\x95\x67\xd1\xb8\x8b\xc1\x0e\xa3\x31\x8b\x2d\x4f\xa3\x09\x3f\xf6\x4a\x2d\xc6\xb7\xd6\xe2\xcb\x7d\x90\x95\x63\xf2\x64\x28\x4a\xe5\x60\xbf\x28\xf7\xcf\x1f\xcf\x55\xa4\x26\x0d\x6b\xef\x79\x84\x82\x1c\x33\x54\x79\x7f\x90\x79\x64\x34\x2d\x23\x81\x34\xd2\xb7\x0f\xd2\xaf\x48\xdb\x8c\x9b\xdb\xab\x10\xa5\x9b\x3f\xe0\xba\x2a\xdf\x50\x57\x08\x2d\x29\x91\xb3\xb7\xbb\x75\xdf\x33\xfd\x7c\x69\x73\xab\x14\x23\x39\xc8\x37\xe4\x51\x96\x50\xef\x25\x52\x66\x92\xc2\x74\x36\x63\x47\xe5\xf0\xb2\xa9\x90\x9f\xff\x33\xf3\x2e\xcd\x74\xdf\xec\xbf\x9f\x5c\x2e\x97\x98\x70\x5f\x24\xa1\xbf\x04\xfd\x24\x5f\xf2\x16\x73\xe7\xf2\xbf\xf9\x57\x96\x3f\xb7\x38\xa1\x9c\xfa\x59\x6c\x5b\x84\x57\x9a\x26\xdb\x2d\xb1\x75\xe1\xe9\x6c\xd5\x73\x7a\x6b\x8b\x39\xff\x95\xb3\x30\xa8\x5b\x2b\x31\x8d\x5a\x26\x6b\xfd\x3f\x6c\x55\xb8\xca\x7e\xda\xfc\xe3\x73\x8d\xeb\xa3\x95\xa8\x06\x05\x7c\x83\x96\xe8\xd5\xaa\x74\xad\xfb\x69\x69\x26\x5e\x47\x4c\xd5\x71\xdc\xab\x88\x53\x9b\x9a\x0e\xaa\x93\x9a\x2a\x4e\xa7\xa8\xf5\x6c\xfa\x50\x0b\xaf\x4f\x85\x66\xdd\x8b\x8d\xd1\x9c\x89\x72\x54\x6f\x02\x76\x84\xaf\x15\xb5\xb8\x05\x74\x1a\xca\x14\x40\x8d\xd6\x5a\x02\x9f\x03\x00\xd4\x7f\xc2\x96\x79\xbd\x87\x95\x0d\xe2\xcf\x39\xb5\xb5\xbf\x33\xab\x0c\xa2\xcf\x3b\x52\xcf\x02\xb5\x6d\x13\x60\x2b\xd1\x55\x08\xad\x87\xb4\xd6\xc5\xbe\x87\x11\x56\x4b\x31\x45\x33\x61\xa5\x94\xde\x03\x76\x81\x4f\xf3\x19\xd4\x35\xb1\xcf\x1b\x4d\xaf\x2b\xe3\x68\x2f\xcc\xff\x8f\x55\x2b\xe3\xf4\x7d\x17\x4c\xdd\xf8\x47\x53\x56\x39\xfa\x44\xa0\x1f\x80\x3a\xcc\x05\x57\x19\x85\x4b\x2c\xe0\x1f\x7f\xff\xb1\xeb\x91\x8f\x64\xeb\x84\x8e\xf2\x67\xfa\x14\xba\x9f\xe6\x35\x98\xc6\x80\xbe\x0d\x13\x92\xd0\xf5\x58\x59\x4c\x9c\x62\x0f\x15\x3c\x4d\x16\xbc\xd0\x52\x76\x43\xf7\xdc\x1d\x17\xcd\xaf\x06\x94\xbe\x19\x60\xea\xd7\x23\x87\xf9\x01\x96\x7a\xaa\xb6\x7d\x9f\x0e\xf8\x25\x24\x9e\x3b\xe8\x8d\x54\xa9\x4e\x5c\xe3\x3f\x18\x1f\xe0\xc9\x81\x6b\x42\x07\x69\x77\x8e\x65\x75\xc8\x56\x9f\x76\x39\x0c\x56\xa2\x27\x79\x9b\xa2\x5c\x1a\xa7\x7f\x6a\xc1\x60\x70\x9d\xd2\x1b\x90\x1e\x8f\xd1\xd6\xfc\xdf\xa1\x6f\x72\x9d\x04\x54\x23\x97\x93\x9f\x2f\xce\x5f\x27\x89\xea\xe2\xfe\xe5\xfe\x0d\x1c\x60\x87\xc8\xdd\x60\x21\x2b\x71\xfb\xcf\xb2\x70\x7d\x55\x6e\xcb\xb9\x00\xc9\x33\x7e\xd7\xaa\xbc\xac\xfd\x09\xd2\x95\xea\x0a\xe4\x0b\xc9\x38\x5e\xf7\x7d\xfa\xc4\x34\x52\xad\x18\xe7\x47\xb8\x6c\x30\xc8\x34\x30\x37\xdd\xe3\x67\xa8\xf2\xf7\xbf\x66\xbb\xea\x9b\x0e\x10\xb5\x32\x94\xac\xa2\xed\xb7\x9d\x9c\xa2\xa1\xda\x27\x29\xd3\xd2\x95\xe7\xbd\xc7\xd4\xf5\x80\x4f\x61\x85\x43\x4f\xfe\xc2\xdc\x5d\x8d\xe4\xde\x72\xed\x8d\x13\x12\x53\x17\x73\x17\x09\xe0\x04\x7b\xe4\x4f\x1d\x72\x71\x35\x9f\x21\x5d\x93\xee\x81\x46\x4f\x6c\x47\x37\x84\x8a\x5c\x9a\x3f\xa1\xe8\xf1\xee\xd7\x0f\xf4\x6f\xe8\x61\x44\xe8\x16\x7b\xc4\x14\xf3\x54\x53\xf6\x30\x32\xbf\x7f\x0d\x99\xc4\x08\x1e\x1d\x00\x17\xdc\xf8\x57\x4d\x6b\xf6\x4a\xd3\xcf\xe8\x81\x4e\x26\x13\x90\xce\x64\x32\x79\xa0\xb3\xa9\xea\x2f\xa4\xe4\x6b\x08\x51\x6f\xc4\x05\x2a\xc9\x8a\x38\xe6\x2b\x75\x56\x7d\xa0\x53\x90\x98\x78\xda\x43\xc1\x02\x13\xa0\xa8\xef\x31\xe1\xb1\xc0\xa4\x40\x5f\x08\x75\xb1\xe9\x7c\x45\xc0\x73\xd1\x45\x7c\x11\x78\x81\xfc\x50\x48\xb4\x04\x44\x19\x1d\xff\x09\x9c\xa1\x2d\xf6\xc2\x64\x04\x94\x49\x04\x94\x85\xeb\x0d\x92\x64\xbd\x91\x02\x49\x86\x56\x00\x2e\x5a\xb3\x60\x03\x3c\xa6\x4b\xea\x01\x5f\xfc\xca\xdc\x0b\xe4\x32\x10\x17\x12\xc1\x23\x11\x52\x91\xbc\x53\xbd\xe6\x59\x15\xa0\x9d\x2a\x5f\x60\x37\xd6\x3d\xa2\x00\x93\x3e\x6e\x30\xd7\x4c\xc7\x13\x85\xd3\x45\xc2\x28\xc7\x5c\x3d\xe7\x4d\xdb\x59\x3a\x45\xd9\x43\x5e\x12\x7b\x9a\x99\x28\xd3\x9c\x9a\x3c\x2d\xe6\xf2\xad\xee\x44\x51\xe7\xcd\x0b\xde\xd9\x00\xe7\xbb\x8e\x0e\x89\xb7\x9c\x9b\xe5\x90\x99\x93\xdc\xca\x98\x4d\xdb\x0f\x3f\xa8\x19\xb4\x5a\x21\x7b\xbf\x96\xf2\xf3\x81\x50\xb7\xc0\x91\xfa\xb8\x5c\x1c\x2d\x42\x83\x9b\x8c\xf5\x3e\x73\xbd\x27\x34\x65\xc6\x90\x96\x03\xbd\xdd\x9b\xf9\x38\xb5\x73\x36\x6d\x31\xda\x12\x3b\xa5\xf7\x50\x86\xd4\xc5\x08\x2c\x3b\x9b\x5d\xd1\x5e\x51\x67\xd0\x2e\x36\x52\x06\x43\xbd\x67\xa7\x7b\x78\x7f\x77\x37\x6f\xdc\xa4\xd7\x40\xc7\x1e\x33\x7b\x89\x5e\x2b\x01\x70\x5c\x08\x32\xec\x30\xa1\xaa\x37\xdd\xd3\x28\xc0\x1c\xfb\x20\xb3\x86\xe6\xe8\x0d\xa3\x14\x1c\xd5\x8b\xf1\xeb\xb6\xb9\xe0\x4b\x47\x5f\xfc\x4b\x26\x06\x2d\x6a\x35\xfd\x1b\x31\xae\x4c\xd3\x4b\x99\xb7\x56\xf2\x10\x8a\x75\x0d\x7f\x87\xe5\x2d\x73\xbe\x80\xfc\x00\xbb\x63\x32\x78\x0b\xce\x38\x69\x7b\xfc\x01\x76\xfd\xf9\x14\xba\xad\xc8\x9b\x53\xcd\x6b\x51\x9f\xf2\xba\x54\xc9\xa9\x69\x7d\xbc\xf7\x08\xe9\xe1\xdc\xde\x07\x6b\x8e\x5d\x38\xe6\x6c\x46\x4d\xf6\xe2\x4a\x00\x9f\x4d\xbb\x30\x15\x86\x19\xdf\x7e\xc2\x52\x74\xa1\x90\xd9\x87\x3a\x32\x95\x9c\xe7\xdb\xf3\x56\xe0\xe1\xd3\xde\xe3\xe5\x1d\x59\xb9\x61\x5e\x09\x17\x40\xf5\xb3\xb0\xff\xcc\x9f\x25\x73\x97\x1b\xae\x9f\x9e\xe8\x3f\xbf\x3c\x90\x7d\xd5\x6b\x27\xce\x13\xdc\x89\xcb\x63\x67\x60\xc7\x6c\x52\xd5\x27\x01\x03\x60\x85\xbd\x6c\x24\x9c\x0d\xf8\xb8\x4d\xc1\xc4\xba\xf3\xc7\x8b\x2c\x0a\x6b\x2e\x5f\x7c\xff\x77\x00\x00\x00\xff\xff\xb4\x2d\xf9\x50\x53\xe7\x00\x00")

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
