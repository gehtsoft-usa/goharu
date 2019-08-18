# What is GoHaru? 

GoHaru is another Go interop module for libharu PDF generation library (http://libharu.org/)

What Makes it Different from others? 

1) The package includes prebuild static libraries for 32 and 64 bit Windows, 32 and 64 ARM, 64 bit Intel Linux, and 64 bit MacOS. 

Yes, we aware that Go ideology wants you to compile the whole libharu from sources, however from our expirience, a correct cross 
platform build of these libraries appears to be challenging for an ordinary developer and is needlessly resource consuming for a reasonable project maintenance schedule. 

So, we included the  cross-platform static libraries, made from the same sources as the binaries used by our own Gehtsoft.PDFFlow library. 

These binaries also contains all recent fixes we made in haru, that may be yet to be incorporated into original libharu code. 

The binaries are tested on Windows 7, Windows 10, Cent OS 7, Ubuntu 18.04, OSX 10.14.

If you still want to accept the challenge of building haru, please feel free to take source code and building procedures here

https://github.com/gehtsoft-usa/libharu (for our fixed version)

or 

https://github.com/libharu/libharu (for original libharu)

2)  We also ported demos and provided function descriptions for godoc 

Use https://github.com/gehtsoft-usa/goharu to see types and functions documentation. 

3) To run demo go to demo subfolder of the folder where the package is installed and run

`go run demo.go help`

to get the list of demos, then run samples, e.g. 

`go run demo.go png fonts`

PDF files are created in out subfolder of demo

# How to Build the library?  

1) In order to succesfully use the library make sure that:

- On Linux: GCC is installed 

- On Windows: MinGW is installed (https://mingw.org, https://sourceforge.net/projects/mingw-w64/) and MinGW's bin folder is available in 
PATH

2) Get package using 

`go get github.com/gehtsoft-usa/goharu`
