package main
import "fmt"

//定义一个文件接口
type IFile interface  {
	Read(buf []byte)(n int,err error)
	Write(buf []byte)(n int,err error)
	Close() error
}

type IReader interface  {
	Read(buf []byte)(n int,err error)
}

type IWriter interface  {
	Write(buf []byte)(n int,err error)

}

type ICloser interface  {
	Close() error
}

//读写接口
type IReaderWriter interface {
	IReader
	IWriter
}

//定义文件类
type File struct {
	filename string
	fid int16
}

//定义文件的方法
func (f *File)Read(buf []byte)(n int,err error){
	fmt.Println("File Opened")
	return 0,nil
}


func (f *File)Write(buf []byte) (n int,err error) {
	fmt.Println("File Writed")
	return 0,nil
}
func (f *File)Close() error {
	fmt.Println("File Closed")
	return nil
}

func main() {

	bytes := []byte{0,1,2,3,4,5,6,7,8,9}
	var file1 IFile = new(File)
	file1.Read(bytes)

	var file2 IReader = new(File)
	file2.Read(bytes)

	var file3 IWriter = new(File)
	file3.Write(bytes)

	var file4 ICloser = new(File)
	file4.Close()


	var file5  IReaderWriter = new(File)
	file5.Read(bytes)
	file5.Write(bytes)
}
