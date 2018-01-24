package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

//Archiver implementations will be responsible for archiving the source folder and storing it
//at the destination path
type Archiver interface {
	Archive(src string, dest string) error
}

type zipper struct{}

//Zip is an Archiver that zips and unzips files
// Provide an instance of the type to save user from worrying about
//managing their own types
var ZIP Archiver = (*zipper)(nil)

//This bit of arcana defines a variable for use outside of the package,
// We assign it with nil cast to the type zipper. There is no reason anybody
//outside of the package needs to know about our zipper type, which frees us to
//change the internals without touching the externals, this is the power of interfaces

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil //skip
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path) //opens source file for reading
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path) //creates a new compressed file
		if err != nil {
			return err
		}
		_, err = io.Copy(f, in) //reads all bytes from the source file and writes them through ZipWriter object
		if err != nil {
			return err
		}
		return nil
	})
}
