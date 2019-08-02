# What is GoHaru? 

GoHaru is an interop module for the libharu PDF generation library (http://libharu.org/)

What Makes it Different? 

1) The package includes prebuild static libraries for 32 and 64 bit Windows, 32 and 64 ARM, 64 bit Intel Linux, and 64 bit MacOS. 

Yes, we aware that Go ideology wants you to compile the whole libharu from sources, however from our expirience, a correct cross 
platform build of these libraries appears to be challenging for an ordinary developer and is needlessly resource consuming for a reasonable project maintenance schedule. 

So, we included the  cross-platform static libraries, made from the same sources as the binaries used by our own Gehtsoft.PDFFlow library. 

These binaries also contains all recent fixes we made in haru, that may be yet to be incorporated into original libharu code. 

If you still want to accept the challenge of building haru, please feel free to take source code and building procedures here

https://github.com/gehtsoft-usa/libharu (for our fixed version)

or 

https://github.com/libharu/libharu (for original libharu)

2)  We also ported demos and provided function descriptions for godoc 

# How to Build? 

In order to succesfully use the library make sure that:

On Linux: GCC is installed 

On Windows: MinGW is installed (https://mingw.org, https://sourceforge.net/projects/mingw-w64/) and ?:\mingw\mingw64\bin\ is available in 
PATH
