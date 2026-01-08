# portablesyscall
An example operating system independent system call interface for Go

The syscall and golang.org/x/sys/unix packages provide interfaces to the UNIX system calls
but the contents is different on different systems.
Functions may or not be present in the package for a particular operating system.
In particular, only some of the functions are available under Windows,
which means that a program that tries to use one of the missing functions
doesn't even compile in that environment.
For example, the windows version of syscall has a Getuid but no Setuid.
It also has a Chown, which always returns an error.

Other missing material includes the stat_t structure, which is used to find the owner
of a file under Linux.

I was a Java programmer for many years and one of the features of Java is that compiles to 
a low-level form which is then interpreted.
The result is "write once, run anywhere",
meaning you can write and test it under one operating system and run it without change
on another.
Most of the Java teams I worked in developed software under Windows
which was then run under UNIX or Linux.

Go has built-in cross compilers,
so it could also be write once, run anywhere,
apart from the carbunkle of the system call interface.

My solution 
is to have a single package with a 
version for each system that I work with.
Each file of source code starts with a build tag
so the appropriate version is compiled on its target system.
The compiler sorts out which one to use.
The package provides the same functions with the same signatures on all of those systems.
In the windows version, functions that can't be made to work, 
all return an error when called,
the same error that syscall.Chown returns.

The constant OSName is provided in all environments and
contains a string giving the name of the operating system on which
the package is running - 
"windows", "linux" or whatever, matching the build tag for that system.
This allows calling software to avoid calling functions under Windows
that are guaranteed to throw errors.

An example of this package in use is my go-stripe-payments website.  
Running an HTTPS server requires a certificate,
which is impractical for the Windows PC on my desk.
For example, the certificate should only be readable by the admin user
so the software needs to run under that user to read them.
Running a web server under the admin user is not a great idea,
so the server should read the cerificate when it starts up
and then switch to  running as a less privileged user.
I'm happy to do all this on my Linux target system,
rather less willing to make it all work under Windows,
when I can do most of my system testing using HTTP.

Using this package, my software still compiles under Windows.
I just need to avoid using some of the features when I run it
in that environment.

My solution only implements material that I need for my own work,
initially the Setuid function and the Stat_t structure.
I don't plan to implement it for all systems, 
just the ones I use (initially Windows and Linux).
If you want something more,
feel free to fork the project or use it as a guide to write your own.
Please don't create issues asking me to add functions.  The answer will be no, please create your own project.
Please don't send pull requests asking me to accept new functions that you've written.
