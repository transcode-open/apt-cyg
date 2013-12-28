apt-cyg
=======

apt-cyg is a command-line installer for [Cygwin](http://cygwin.com/) which cooperates with Cygwin Setup and uses the same repository. The syntax is similar to apt-get. Usage examples:

* "apt-cyg install <package names>" to install packages
* "apt-cyg remove <package names>" to remove packages
* "apt-cyg update" to update setup.ini
* "apt-cyg show" to show installed packages
* "apt-cyg find <pattern(s)>" to find packages matching patterns
* "apt-cyg describe <pattern(s)>" to describe packages matching patterns
* "apt-cyg packageof <commands or files>" to locate parent packages
* "apt-cyg pathof <cache|mirror|mirrordir|cache/mirrordir>" to show path"
* "apt-cyg key-add <files> ..." to add keys contained in <files>
* "apt-cyg key-del <keyids> ..." to remove keys <keyids>
* "apt-cyg key-list" to list keys
* "apt-cyg key-finger" to list fingerprints

Quick start
-----------

apt-cyg is a simple script. Once you have a copy, make it executable:

    # chmod +x /bin/apt-cyg

Optionally place apt-cyg in a bin/ folder on your path.

Then use apt-cyg, for example:

    # apt-cyg install nano

New features
------------

### True multi-architecture support

Let think a case that you want to install the x86 package when you are working under the x86_64 environment.
In this case, You must only append a --charch option to the first parameter of apt-cyg.
For example:

    # apt-cyg --charch x86 install chere

As of 2013-10-26, chere package is provided for only the repository for x86.

Remarks:
Of course, you must install both environments of x86_64 and x86, beforehand.

### Signature check and key management by GnuPG

The default action of apt-cyg has been changed to check signature for 'setup.ini'.
Of course you can also avoid signature check by using --no-verify or -X options.
Public keys of cygwin and cygwinports are already registered to trusted keys of embeded.
If you want to use some other public keys, please use key-* subcommands.

Contributing
------------

This project has been re-published on GitHub to make contributing easier. Feel free to fork and modify this script.

The [Google Code project](https://code.google.com/p/apt-cyg/) has a list of open issues.

### Forks on the github

Caution:
Please do not merge forks that have incompatible licenses.

Ex.) Merging to the GPL from the MIT is possible. But merging to the MIT from the GPL  is impossible.

#### Official (MIT license)

* [transcodes-open / apt-cyg](https://github.com/transcode-open/apt-cyg/network)

#### Unofficial (GPLv2)

* [cfg / apt-cyg](https://github.com/cfg/apt-cyg/network)
* [ashumkin / apt-cyg](https://github.com/ashumkin/apt-cyg/network)
* [svn2github / apt-cyg](https://github.com/svn2github/apt-cyg/network)
* [nosuchuser / apt-cyg](https://github.com/nosuchuser/apt-cyg/network)
* [kazuhisya / apt-cyg64](https://github.com/kazuhisya/apt-cyg64/network)
* [bnormsoftware / apt-cyg](https://github.com/bnormsoftware/apt-cyg/network)
* [rcmdnk / apt-cyg](https://github.com/rcmdnk/apt-cyg/network)
* [buzain / apt-cyg](https://github.com/buzain/apt-cyg/network)
* [wuyangnju / apt-cyg](https://github.com/wuyangnju/apt-cyg/network)
* [takuya / apt-cyg](https://github.com/takuya/apt-cyg/network)

Todo
------------

* Support multi mirrors: Cygwin setup can use multi mirrors. They are recorded at last-mirror section in '/etc/setup/setup.rc'. It's useful for using [Cygwinports](http://cygwinports.org/).
* Support completion: Some other forks already supported it. For example, [Milly / apt-cyg](https://github.com/Milly/apt-cyg) under the cfg / apt-cyg fork, [ashumkin / apt-cyg](https://github.com/ashumkin/apt-cyg) and etc supported it.
* Support upgrade: But maybe, busy resources can not be upgraded, and rebase problem will happen. Cygwin setup resolves by replacing them at next reboot.
* Support dependency check for remove subcommand.

Known Problem
------------
### FIXED: Check setup (cygcheck -c) can not detect .tar.xz packages
At cygwin 1.7.25, cygcheck is hardcoded for .tar.gz and .tar.bz2.
So check setup (cygcheck -c) can not detect .tar.xz packages.
This bug was already fixed with [src/winsup/utils/dump_setup.cc#rev1.28](http://cygwin.com/cgi-bin/cvsweb.cgi/src/winsup/utils/dump_setup.cc?cvsroot=src#rev1.28).
Please wait a release of cygwin 1.7.26.

This Problem was fixed [cygwin 1.7.26](http://cygwin.com/ml/cygwin-announce/2013-11/msg00027.html) which released at 2013-11-29.
