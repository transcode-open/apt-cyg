# Why Go?

Consider this Python script

~~~py
print("Hello world")
~~~

You can compile it for Windows like this

    build_exe --bundle-files 0 --compress --optimize hello.py
    upx -9 hello.exe

and it comes out to 4.16 MB. Now consider this Go script

~~~go
package main
import "fmt"
func main() {
	fmt.Println("Hello world")
}
~~~

You can compile it like this

    go build -ldflags -s hello.go
    upx -9 hello.exe

and it comes out to 375 KB. Something else to consider is that Python is built
with Visual C++ 2010. That means that the program you create will not work on
Windows 7 for example unless the user has installed Visual C++ 2010. Go is built
with Visual C++ 6, this means that any programs created with Go will work on
Windows 7 as is.
