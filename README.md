apt-cyg
=======

apt-cyg is a Cygwin package manager. It includes a command-line installer for
Cygwin which cooperates with Cygwin Setup and uses the same repository. The
syntax is similar to apt-get. Usage examples:

* `apt-cyg install <package names>` to install packages
* `apt-cyg remove <package names>` to remove packages
* `apt-cyg update` to update setup.ini
* `apt-cyg list [pattern(s)]` to list packages matching given pattern. If no
  pattern is given, list all installed packages.
* `apt-cyg show <pattern(s)>` to show packages matching patterns
* `apt-cyg search <commands or files>` to locate parent packages

Quick start
-----------

apt-cyg is a simple script. To install:
```
# wget https://raw.githubusercontent.com/transcode-open/apt-cyg/master/apt-cyg
```
then place it in a `/bin` folder on your path:
```
# mv apt-cyg /bin/apt-cyg
```
and make it executable:
```
chmod +x /bin/apt-cyg
```

Example use of apt-cyg:
```
# apt-cyg install nano
```

Contributing
------------

This project has been re-published on GitHub to make contributing easier. Feel
free to fork and modify this script.

The [Google Code project](http://apt-cyg.googlecode.com) also has a list of
open issues.
