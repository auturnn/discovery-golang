package dependency

import "os"

// 코드를 테스트할 경우 외부 리소스에 접근하는 것을 막고 싶은 경우가 많은데,
// 이때 interface를 이용한 추상화를 통해 유연하게 대처가능하다.
// kickshaw-coin을 통해 확인해보는 것도 좋을 듯 하다.
type FileSystem interface {
	Rename(oldpath, newpath string) error
	Remove(name string) error
}

type OSFileSystem struct{}

func (fs OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs OSFileSystem) Remove(name string) error {
	return os.Remove(name)
}
