package static

import (
	"embed"
	"fmt"
	"github.com/benbjohnson/hashfs"
	"io/fs"
)

// Statics wraps the underlying fs.FS and provides hashed filenames for commonly used static files
type Statics struct {
	DTDFile string
	FS      fs.FS
	PAAFile string
}

//go:embed squad.dtd squad.paa
var staticFS embed.FS

// LoadStatics verifies that the embedded files are all present and fetches hashes for commonly used static files
func LoadStatics() (*Statics, error) {
	requiredFiles := []string{"squad.dtd", "squad.paa"}

	for _, file := range requiredFiles {
		_, err := staticFS.Open(file)
		if err != nil {
			return nil, fmt.Errorf("unable to read %s from embedded files", file)
		}
	}

	hashedStaticFS := hashfs.NewFS(staticFS)

	return &Statics{
		DTDFile: hashedStaticFS.HashName("squad.dtd"),
		PAAFile: hashedStaticFS.HashName("squad.paa"),
		FS:      hashedStaticFS,
	}, nil
}
