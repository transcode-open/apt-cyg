apt-cyg
=======

apt-cyg is a command-line installer for Cygwin which cooperates with Cygwin Setup and uses the same repository. The syntax is similar to apt-get. Usage examples:

* "apt-cyg install <package names>" to install packages
* "apt-cyg remove <package names>" to remove packages
* "apt-cyg update" to update setup.ini
* "apt-cyg show" to show installed packages
* "apt-cyg find <pattern(s)>" to find packages matching patterns
* "apt-cyg describe <pattern(s)>" to describe packages matching patterns
* "apt-cyg packageof <commands or files>" to locate parent packages 

Quick start
-----------

apt-cyg is a simple script. Once you have a copy, make it executable:

  # chmod +x /bin/apt-cyg

Optionally place apt-cyg in a bin/ folder on your path.

Then use apt-cyg, for example:

  # apt-cyg install nano

True multi-architecture support
------------

Let think a case that you want to install the x86 package when you are working under the x86_64 environment.
In this case, You must only append a --charch option to the first parameter of apt-cyg.
For example:

  # apt-cyg --charch x86 install chere

As of 2013-10-26, chere package is provided for only the repository for x86.

Remarks:
Of course, you must install both environments of x86_64 and x86, beforehand.

Contributing
------------

This project has been re-published on GitHub to make contributing easier. Feel free to fork and modify this script.

The [Google Code project](https://code.google.com/p/apt-cyg/) has a list of open issues.

### Forks on the github

Caution:
Please do not merge forks that have incompatible licenses.

Ex.) Merging to the GPL from the MIT is possible. But merging to the MIT from the GPL  is impossible.

#### Official

* [transcodes-open / apt-cyg](https://github.com/transcode-open/apt-cyg/network) (MIT license)

#### Unofficial

* [cfg / apt-cyg](https://github.com/cfg/apt-cyg/network) (GPLv2)
* [ashumkin / apt-cyg](https://github.com/ashumkin/apt-cyg/network) (GPLv2)
