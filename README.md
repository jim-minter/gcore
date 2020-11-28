# gcore

Generates a core file from a running process without terminating it.  The
process will be paused while the core file is being written, however.  The core
file is written to stdout.  Equivalent to `gdb`-based `gcore`, but is statically
linked, doesn't require `gdb` or its dependencies, and is container-aware (i.e.
you can harvest the core from a process running inside a container from outside
the container).

Currently only runs against 64-bit target processes on Linux/x86_64.

Usage: `gcore pid | gzip >core.gz`.
