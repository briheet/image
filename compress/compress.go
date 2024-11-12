package compress

import (
	"io"
	"os"
)

func CompressTofile(got string, name string) (os.File, error) {
	given, err := os.Open(got)
	if err != nil {
		return os.File{}, err
	}
	defer given.Close()

	compressedFile, err := os.Create(name)
	if err != nil {
		return os.File{}, nil
	}
	defer compressedFile.Close()

	buf := make([]byte, 1024)

	for {
		n, err := given.Read(buf)
		if err != nil && err != io.EOF {
			return os.File{}, nil
		}
		if n == 0 {
			break
		}

		_, err = compressedFile.Write(buf)
		if err != nil {
			return os.File{}, nil
		}
	}

	return *compressedFile, nil
}
