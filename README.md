# portablesyscall
An example operating system independent system call interface for Go

The syscall and golang.org/x/sys/unix packages provide interfaces to the UNIX system calls
but the contents is different for different target systems.
Functions may or not be present in the package for a particular target.
In particular, only some of the functions are available when running
under Windows,
which means that a program that tries to use one of the missing functions
doesn't even compile for some targets.
For example, the windows version of syscall has a Getuid but no Setuid.
It also has a Chown, which always returns an error.

Other material missing under Windows includes the stat_t structure, which is needed to find the owner
of a file under Linux.

Meanwhile, the Windows syscall defines a value EWINDOWS which it uses to create
error values such as the one returned by the Chown function whenever it's called under Windows.  
That value is not defined in the Linux syscall,
so any program that references EWINDOWS will not compile under Linux.

I was a Java programmer for many years and one of the features of Java is that 
its runtime library is the same for all targets.
This is part of its "write once, run anywhere" philosophy,
meaning that you can write and test Java software under one operating system and run it without change
on another.
Most of the Java teams I worked in developed software under Windows
which was then run under UNIX or Linux.
(Strictly, if the source code references the full path name of a file,
it will need changing between Linux and Windows.)

Go has built-in cross compilers,
so it could also be write once, run anywhere,
apart from the carbunkle of the syscall package.

My solution 
is to create a single package with a 
version for each system that I work with and
use build tags to arrange for the appropriate source to be compiled.
The package provides the same functions with the same signatures on all of those systems
plus any other material needed,
such as a common Stat_t structure.
When Windows is the target, functions that can't be made to work
return an error when called,
the same error that syscall.Chown returns.

Just to be clear,
my goal here is to be able to write software that compiles
regardless of whether the target is Linux or Windows.
I'm not trying to provide Linux-like features that work at run time under Windows.

The constant OSName is provided for all targets and
contains a string giving the name of the operating system on which
the package is running - 
"windows", "linux" or whatever, matching the build tag for that system.
This allows calling software to avoid calling functions under Windows
that are guaranteed to throw errors.
(There are other ways for a program to figure out which environment it is 
running on,
but this seems the simplest.)

An example usage of this package is my go-stripe-payments web server.
It's intended to run on a Linux target and provide an HTTPS service.
An HTTPS service is hard to achieve on a Windows desktop system.
It's possible
but
I can test most of the functionality
using an HTTP service, 
so I don't need to bother.

Thae main problem is that running an HTTPS server requires a certificate,
which is usually only available on a server system connected to the Internet.
It's possible to create a certificate for a local desktop system but
the certificate should only be readable by the admin user
so the software needs to run under that user to read them.
Running a web server under the admin user is not a great idea,
so the server should read the cerificate when it starts up
and then switch to  running as a less privileged user.
I'm happy to do all this on my Linux target system,
rather less willing to make it all work under Windows,
when I can do most of my system testing using HTTP.

My solution only implements material that I need for my own work,
initially the Setuid and Stst functions and the Stat_t structure.
I don't plan to implement it for all systems, 
just the ones I use (currently Windows and Linux).
If you want something more,
feel free to fork the project or use it as a guide to write your own.
Please don't create issues asking me to add functions.  The answer will be no, please create your own project.
Also,
please don't send pull requests asking me to accept new functions that you've written.
I will reject them.
