apt-cyg
=======

**apt-cyg** is a Cygwin package manager.
It includes a command-line installer for **Cygwin** which 
cooperates with Cygwin **Setup** and uses the same repository.

[github.com/transcode-open/apt-cyg][1]

[1]:https://github.com/transcode-open/apt-cyg

Operations
----------

**list**[**all**] [_pkg_...]

  Search for package names that match regexp.  If no package names are
  provided in the command line, all installed packages will be listed.  If
  listall is used, searches the master package list (setup.ini).

**show** _pkg_...

  Display information on given package name(s).

**listfiles** _pkg_...

  List all files owned by a given package. Multiple packages can be specified
  on the command line.

**install** _pkg_...

  Install package(s) and any dependencies.

**source** _pkg_...

  Retrieve package source(s) from the server into package directory created
  under current directory and unpack under the package directory.

**download** _pkg_...

  Retrieve package(s) from the server, but do not install/upgrade anything.

**remove** _pkg_...

  Remove package(s) from the system.

**depends** _pkg_...

  Produce a tree of all dependencies for a package.

**rdepends** _pkg_...

  Produce a tree of packages that depend on the named package.

**search**[**all**] _file_...

  Search for packages that own the specified file(s).  The path can be
  relative or absolute, and one or more files can be specified.  If searchall
  is used searches cygwin.com for packages that own the specified file(s).

**category** [_cat_...]

  Display all packages that are members of a named category.  If no category
  is provided in the command line, all categories used will be listed.

**update**

  Download a fresh copy of the master package list (setup.ini) from the
  mirror.

**mirror** [_URL_]

  Set the mirror: a full URL to a location where the database, packages, and
  signatures for this repository can be found.  If no URL is provided,
  display the current mirror.

**cache** [_directory_]

  Set the package cache directory.  Unix and Windows forms are accepted, as
  well as absolute or regular paths.  If no directory is provided, display
  current cache.  If a package to install is not found in the cache
  directory, it will be downloaded.

  **--build**|**--compile**

  With source: install any build dependencies; if cygport is part of the
  package, include cygport and any of its build dependencies, and build
  package using cygport; otherwise try to configure, then try to make; if
  cygport is not part of the package, build dependencies may be missing, and
  need to be installed manually.

  **--download**|**--download-only**

  With source: just download and do not unpack source package.

  **--nodeps**

  Specify this option to skip all dependency checks, and not download or
  install packages on which specified packages are dependent.

  **--nopick**

  Internal option for install of a package dependency.

  **--noscripts**

  Specify this option to skip running any preremove or postinstall scripts.
  Used internally during install to defer running postinstall scripts until
  all requested packages and dependencies have been installed.

**--help**

  Display usage and exit.

**--version**

  Display version and exit.

Quick start
-----------

**apt-cyg** is a simple script. To install:

    lynx -source rawgit.com/transcode-open/apt-cyg/master/apt-cyg > apt-cyg
    install apt-cyg /bin

Example use of **apt-cyg**:

    apt-cyg install nano

