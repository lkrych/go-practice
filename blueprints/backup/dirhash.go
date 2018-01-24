package backup

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//DirHash generates an MD5 hash that represents whether or not a files/folder has changed
func DirHash(path string) (string, error) {
	hash := md5.New()
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		io.WriteString(hash, path)
		fmt.Fprintf(hash, "%v", info.IsDir())   //is item file or folder?
		fmt.Fprintf(hash, "%v", info.ModTime()) //the last modified time
		fmt.Fprintf(hash, "%v", info.Mode())    //file mode in bits
		fmt.Fprintf(hash, "%v", info.Name())    //has filename changed?
		fmt.Fprintf(hash, "%v", info.Size())    //has size of file changed?
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil //hex formatting directive
}
