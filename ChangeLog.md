2017-04-13 Brian Inglis
* [2afd9f7](http://github.com/transcode-open/apt-cyg/commit/2afd9f7)
    apt-cyg: add function upgradable as an awk script wrapper to compare
    <cache>/<mirror>/setup.ini to /etc/setup/installed.db and display the
    number of, and details about, packages where the versions differ, and the
    size of the archives to be download, call at end of apt-update, make output
    more like Debian apt, change wrapped awk indentation, use zcat, zgrep
    instead of unpacking /etc/setup/<pkg>.lst.gz, change wget calls to remove
    wget-only output file option;
    apt-cyg(apt-list): change to list nothing unless (a) package(s) specified,
    or --upgradable or --installed requested, like Debian apt, and with similar
    output, support glob patterns and convert to regexps, scan with one pass
    over setup.ini instead of one per package, call apt-listall if no packages
    found;
    apt-cyg(apt-listall): support glob patterns and convert to regexps, scan
    with one pass over setup.ini instead of one per package, output similar to
    Debian apt,
    apt-cyg(wget): if two args passed, use first as output file adapted to
    command issued.

2017-04-13 Brian Inglis
----------
* [5779d50](http://github.com/transcode-open/apt-cyg/commit/5779d50)
  ChangeLog.md, README.md: add markup,
  apt-cyg: speed up by using sed -i instead of awk inplace, sed instead of awk,
  parameter substitution instead of sed, ensure standard collating sequence in
  commands, change some messages to warnings, change uses of script name to $0,
  add blank lines for readability, reorder command checking to put more general
  queries before actions and developer and maintenance commands after,
  apt-cyg(usage): reorder usage commands to put more general queries before
  actions and developer and maintenance commands after, add command arguments,
  OPTIONS, DIAGNOSTICS, and contributors to AUTHORS,
  apt-cyg(wget): fallback from wget to curl to lynx and complain if none
  available,
  apt-cyg(get-setup): support getting setup.xz, .bz2, or .ini if available,
  apt-cyg(apt-category): add show-categories and call if no category given,
  apt-cyg(download): don't replace /etc/setup/<pkg>.lst.gz,
  apt-cyg(apt-searchall): combine wget arg building into statement url, combine
  awk skips into one conditional skip,
  apt-cyg(apt-install): create/replace /etc/setup/<pkg>.lst.gz after successful
  download, create temp installed db in /etc/setup/ instead of /tmp/, support
  all specified permanent postinstall script names,
  apt-cyg(apt-remove): check for more useful commands whose packages should not
  be removed

2017-04-02 Brian Inglis
----------
* [9ec42b5](http://github.com/transcode-open/apt-cyg/commit/9ec42b5)
  add apt-cyg man page

2017-04-01 Brian Inglis
----------
* [9fe833f](http://github.com/transcode-open/apt-cyg/commit/9fe833f)
  rename ChangeLog and add history

2017-04-01 Brian Inglis
----------
* [db90916](http://github.com/transcode-open/apt-cyg/commit/db90916)
  rename README and update to match man page

2017-04-01 Brian Inglis
----------
* [69d8c26](http://github.com/transcode-open/apt-cyg/commit/69d8c26)
  add apt-cyg man page

2016-10-26 Brian Inglis
----------

* [1693e11](http://github.com/transcode-open/apt-cyg/commit/1693e11)
  add wget -nv non-verbose option
* [5cd1d58](http://github.com/transcode-open/apt-cyg/commit/5cd1d58)
  add apt-get like source command with --download-only and --compile
  options
  -- compile looks for cygport, configure, or makefiles and uses them
  -- with cygports, defined build dependencies (DEPEND=...) including
     cygport itself are installed
* [bdd7ea8](http://github.com/transcode-open/apt-cyg/commit/bdd7ea8)
  run all postinstall scripts after all packages installed
* [97da860](http://github.com/transcode-open/apt-cyg/commit/97da860)
  indent install package loop code, support dash scripts, don't rename
  permanent first and last to run postinstall scripts

2016-08-15 Stephen Jungels
----------

* [2cac824](http://github.com/transcode-open/apt-cyg/commit/2cac824)
  Update readme.md
  Removed bountysource link

2016-03-07 Stephen Jungels
----------

* [3a40db9](http://github.com/transcode-open/apt-cyg/commit/3a40db9)
  Delete go-cyg.go
* [5ef3cf4](http://github.com/transcode-open/apt-cyg/commit/5ef3cf4)
  Delete README.md
* [22a6874](http://github.com/transcode-open/apt-cyg/commit/22a6874)
  Delete apt-msys2
* [1d5df19](http://github.com/transcode-open/apt-cyg/commit/1d5df19)
  Merge pull request #75 from lcorbasson/patch-1
  Create a license file (MIT)

2016-03-07 Loïc Corbasson
----------

* [807a91f](http://github.com/transcode-open/apt-cyg/commit/807a91f)
  Create a license file (MIT)
  MIT license like in the source files

2016-03-01 Stephen Jungels
----------

* [0638b38](http://github.com/transcode-open/apt-cyg/commit/0638b38)
  Removed old status update

2016-01-24 Stephen Jungels
----------

* [cf0e883](http://github.com/transcode-open/apt-cyg/commit/cf0e883)
  Update readme.md
  Status update

2016-01-22 Steven Penny
----------

* [b6076c2](http://github.com/transcode-open/apt-cyg/commit/b6076c2)
  you guys love wasting my time or: fix #72

2016-01-03 Steven Penny
----------

* [b3668b1](http://github.com/transcode-open/apt-cyg/commit/b3668b1)
  apt-cyg version 1

2016-01-02 Steven Penny
----------

* [2f59c98](http://github.com/transcode-open/apt-cyg/commit/2f59c98)
  usage

2015-12-07 Steven Penny
----------

* [fd985e7](http://github.com/transcode-open/apt-cyg/commit/fd985e7)
  Fix #69

2015-11-28 Steven Penny
----------

* [0e1df15](http://github.com/transcode-open/apt-cyg/commit/0e1df15)
  modified:   readme.md
* [4a792a1](http://github.com/transcode-open/apt-cyg/commit/4a792a1)
  modified:   readme.md

2015-11-25 Steven Penny
----------

* [734e6eb](http://github.com/transcode-open/apt-cyg/commit/734e6eb)
  renamed:    CHANGELOG.md -> changelog.md
  renamed:    README.md -> readme.md

2015-09-06 Steven Penny
----------

* [395caa5](http://github.com/transcode-open/apt-cyg/commit/395caa5)
  --nodeps
  Specify this option to skip all dependency checks

2015-07-29 Steven Penny
----------

* [a517b08](http://github.com/transcode-open/apt-cyg/commit/a517b08)
  Fix #58 fix #59
  If a user sets cache as "/", this gets converted to "C:\cygwin64". Instead of
  coverting back to "/", causing an error with "mkdir", we can just use
  "C:\cygwin64" as is. "DosFileWarning" now off by default, so it should not be
  necessary to convert cache path to Unix form.
  http://cygwin.com/ml/cygwin/2015-07/msg00432.html

2015-07-07 Steven Penny
----------

* [4e05a45](http://github.com/transcode-open/apt-cyg/commit/4e05a45)
  searchall: exclude cygwin32 packages

2015-05-30 Steven Penny
----------

* [b1d9475](http://github.com/transcode-open/apt-cyg/commit/b1d9475)
  Fix #50

2015-04-18 Steven Penny
----------

* [f2fa0ab](http://github.com/transcode-open/apt-cyg/commit/f2fa0ab)
  mirror and cache operations
  mirror and cache are now operations rather than options. This is in line with
  both apt-get and pacman. Note this does mean some situations may require extra
  commands. For example if you want to temporarily change the mirror, you will
  need to do
```
    apt-cyg mirror http://alpha.com
    apt-cyg install bravo
    apt-cyg mirror http://charlie.com
```

2015-04-17 Steven Penny
----------

* [7a2044e](http://github.com/transcode-open/apt-cyg/commit/7a2044e)
  Kill info function
* [3ca57e8](http://github.com/transcode-open/apt-cyg/commit/3ca57e8)
  Update documentation
  - Update readme
  - Split “list” into “list” and “listall”
  - Kill short option names
  - Kill --file option
  - Kill --help option

2015-04-16 Steven Penny
----------

* [27570ba](http://github.com/transcode-open/apt-cyg/commit/27570ba)
  new file: go/README.md
  modified: go/go-cyg.go
* [83fb7b8](http://github.com/transcode-open/apt-cyg/commit/83fb7b8)
  modified: README.md
* [56f3961](http://github.com/transcode-open/apt-cyg/commit/56f3961)
  Fix #43

2015-04-12 Steven Penny
----------

* [210e5a3](http://github.com/transcode-open/apt-cyg/commit/210e5a3)
  renamed: go-cyg.go -> go/go-cyg.go
* [83b2425](http://github.com/transcode-open/apt-cyg/commit/83b2425)
  Begin Go rewrite
  This may end up being nothing, but would like to rewrite apt-cyg in Go. This
  will be nice because you will end up with a single EXE file, similar to
  setup-x86_64.exe. The difference is that it will be totally command line like
  apt-cyg, but better because you can build all the requirements
  (wget, bunzip2, awk) right into the EXE. My concern right now is that I will
  probably have to write an INI parser, because Cygwin setup.ini is a weird
  format. Since I am just starting with Go I may not have the skill, we shall see.

2015-01-02 Steven Penny
----------

* [aea773d](http://github.com/transcode-open/apt-cyg/commit/aea773d)
  Fix #29 Fix #30

2014-12-29 Steven Penny
----------

* [cc9b2bd](http://github.com/transcode-open/apt-cyg/commit/cc9b2bd)
  functions: download set-cache set-mirror apt-install apt-remove begin charlie
  Here are the changes, by line number:
  - remove "-c", keep long version
  - remove "-f" and "--file". I have never used this, and apt-get does does have
    it
  - remove "-m", keep long version
  - remove "--help", you can get help by just "apt-cyg"
  - lowercase ARCH variable, typically uppercase is for exported variables
  - use readonly variables where applicable
  - remove setup.ini-save, I see no point in this, and setup.exe does not use it
  - go back to using "[" over "[[", more portable
  - declare all variables local in "begin" function, this will prevent creation of
    global variables. Child functions still have access to local variables, but
    the only fix to that is running every function in a subshell, which is also
    not ideal
  - use "let" over "(("
  - stop using awk ENVIRON array. This became a problem because the array is only
    populated with exported variables. We do not want variables being passed to
    child "apt-cyg" processes just to play nice with awk
  - use "return" in functions instead of "exit", this will allow proper running of
    "charlie" function
  - add "check-packages" to some functions
  - stop creating "desc" files, you can parse this information out of setup.ini
    easy enough, and setup.exe does not create them either
  - generate md5 file from setup.ini, this way you can use it with md5sum instead
    of stdin
  - utilize "wget -rnH" to create directories
  - stop running "type -P" when searching for a file with "apt-cyg search".
    apt-get does not do this and it could be confusing
  - break "--cache" and "--mirror" out of the case statement and into their own
    functions. this will make the case statement less unwieldy
  - write directly to "installed.db" using "awk -i inplace"
  - function "charlie": this check for any leaked global variables and prints them
    to stdout just before the script finishes. this should help with debugging

2014-12-27 Steven Penny
----------

* [c325aee](http://github.com/transcode-open/apt-cyg/commit/c325aee)
  modified: apt-cyg
* [db7bada](http://github.com/transcode-open/apt-cyg/commit/db7bada)
  refactor remove function
  - use one awk command instead of two, to determine if package is not removable
  - save comparison files to /etc/setup instead of /tmp
  - save as "essential.lst" and "$pkg.lst" respectively
  - use warnings where appropriate
  - new file list logic
    - extract file list
    - if package is essential, remove file list and exit
    - if package is not essential, remove package, remove file list and continue
  - new delete logic
    - remove all files
    - remove empty directories
  - utilize awk inplace where appropriate

2014-12-26 Steven Penny
----------

* [dcd9b3d](http://github.com/transcode-open/apt-cyg/commit/dcd9b3d)
  Fix #12

2014-12-25 Steven Penny
----------

* [b1162d9](http://github.com/transcode-open/apt-cyg/commit/b1162d9)
  modified: apt-cyg

2014-12-20 Steven Penny
----------

* [44ed47e](http://github.com/transcode-open/apt-cyg/commit/44ed47e)
  modified: README.md
* [61ae91a](http://github.com/transcode-open/apt-cyg/commit/61ae91a)
  Remove wget dependency
  apt-cyg has had a long standing circular dependency with wget. This is because
  Base Cygwin does not include wget or curl. I have mentioned before workarounds
  to this
  41e1d911728653eaf314769ea29ea94df69a3702
  However none were ideal. I have since discovered lynx, which can be used in a
  pinch to download files
  http://unix.stackexchange.com/a/83987
  lynx has these features
  - pure command line, unlike setup-x86_64
  - works with "https://" URLs, unlike /dev/tcp
  - comes with Base Cygwin, unlike wget/curl
  If wget is not installed, lynx will be used as a fallback with a warning
  printed.

2014-12-16 Steven Penny
----------

* [aa4a288](http://github.com/transcode-open/apt-cyg/commit/aa4a288)
  modified: README.md
* [91c7746](http://github.com/transcode-open/apt-cyg/commit/91c7746)
  modified: README.md
* [b6cbda6](http://github.com/transcode-open/apt-cyg/commit/b6cbda6)
  rdepends: find-like output
  e9bea37b959b9cb9f933d0379f2e8db3cc73b589

2014-11-17 Steven Penny
----------

* [3f2b9d9](http://github.com/transcode-open/apt-cyg/commit/3f2b9d9)
  modified:   README.md

2014-11-14 Steven Penny
----------

* [b6f9d53](http://github.com/transcode-open/apt-cyg/commit/b6f9d53)
  minor edit

2014-11-12 Steven Penny
----------

* [23c2db7](http://github.com/transcode-open/apt-cyg/commit/23c2db7)
  searchall: fix split lines
  http://cygwin.com/ml/cygwin/2014-11/msg00287.html
* [9b5de3f](http://github.com/transcode-open/apt-cyg/commit/9b5de3f)
  searchall: account for broken lines
  http://cygwin.com/ml/cygwin/2014-11/msg00278.html
* [e9bea37](http://github.com/transcode-open/apt-cyg/commit/e9bea37)
  depends: find-like output
  when dealing with package dependency woes, such as
  http://cygwin.com/ml/cygwin/2014-10/msg00563.html
  The current "depends" tree method fails, because it only shows the packages, not
  the dependency paths. These paths are necessary to solve the "shortest path
  problem"
  http://wikipedia.org/wiki/Shortest_path_problem
  Now, every possible dependency path will print from the chosen package. Combined
  with category search you can run searches such as
      apt-cyg category Base |
      apt-cyg depends |
      awk '/perl$/ {print length,$0}' |
      sort -n
  Solving the path problem in seconds.

2014-11-01 Steven Penny
----------

* [2138f54](http://github.com/transcode-open/apt-cyg/commit/2138f54)
    check-packages

2014-10-31 Steven Penny
----------

* [8eb3151](http://github.com/transcode-open/apt-cyg/commit/8eb3151)
  Support standard input
  Can now pipe package names to apt-cyg, example
      echo bash | apt-cyg show
  If --file is provided as well, it will override stdin
      echo bash | apt-cyg show --file foo.txt
  Note the Linux analog "apt-cache show" does not provide this functionality.
* [83cac78](http://github.com/transcode-open/apt-cyg/commit/83cac78)
  category command
  This is similar to the command
      aptitude search '?section(shells)'
  http://askubuntu.com/a/473511

2014-10-26 Steven Penny
----------

* [7bee922](http://github.com/transcode-open/apt-cyg/commit/7bee922)
  Never print "No packages found" with apt-cyg list
* [c7f61ef](http://github.com/transcode-open/apt-cyg/commit/c7f61ef)
  Update README.md

2014-08-04 Steven Penny
----------

* [7bec32e](http://github.com/transcode-open/apt-cyg/commit/7bec32e)
  Consistent function declarations

2014-07-26 Steven Penny
----------

* [e83c360](http://github.com/transcode-open/apt-cyg/commit/e83c360)
  Merge pull request #20 from kuc/improve-readme
  Improve README.md
* [3ad0cc5](http://github.com/transcode-open/apt-cyg/commit/3ad0cc5)
  Improve README

2014-06-28 Steven Penny
----------

* [73702ab](http://github.com/transcode-open/apt-cyg/commit/73702ab)
  Fix #18

2014-06-18 Steven Penny
----------

* [8c91a53](http://github.com/transcode-open/apt-cyg/commit/8c91a53)
  Write to and read from setup.rc
  "last-cache" and "last-mirror" will now be written to and read from setup.rc,
  instead of the non-standard /etc/setup/last-{cache,mirror}.

2014-06-16 Steven Penny
----------

* [12f3947](http://github.com/transcode-open/apt-cyg/commit/12f3947)
  Minor fix
* [b1ad202](http://github.com/transcode-open/apt-cyg/commit/b1ad202)
  Refactor download fucntion
  Some packages were being downloaded to wrong location. For example "gcc-core" is
  supposed to be found at
      release/gcc/gcc-core
  However it was being downloaded at
      release/gcc-core
  Packages will now be downloaded as directed by "setup.ini"

2014-06-15 Steven Penny
----------

* [72fe905](http://github.com/transcode-open/apt-cyg/commit/72fe905)
  download command
  This is similar to the "apt-get --download-only install" command.
* [8882af7](http://github.com/transcode-open/apt-cyg/commit/8882af7)
  Revise usage
* [4af30cc](http://github.com/transcode-open/apt-cyg/commit/4af30cc)
  Stop using google search
  I discovered today that not all packages have been indexed by Google. For
  example the 64-bit "gcc-core" which provides "gcc.exe"
  http://cygwin.com/packages/x86_64/gcc-core
  This search
  http://google.com/search?q=%22installed+binaries%22+%22usr%2Fbin%2Fgcc.exe%22+x86_64
  currently only yields the "gcc-debuginfo" package.

2014-06-14 Steven Penny
----------

* [59c5236](http://github.com/transcode-open/apt-cyg/commit/59c5236)
  Simplify proxy check
* [67658a5](http://github.com/transcode-open/apt-cyg/commit/67658a5)
  Improve JSON parser
* [a58f4b6](http://github.com/transcode-open/apt-cyg/commit/a58f4b6)
  JSON function
* [ed6092a](http://github.com/transcode-open/apt-cyg/commit/ed6092a)
  New proxy logic
  1. create proxy file if not exist
  2. create array from proxy file if not exist
  3. if array is empty download proxy file
  4. shift continue unless country is US
  5. shift continue if proxy is slow
  6. shift continue if proxy is blocked by Google
  7. print successful proxy
  8. save remaining proxies to file

2014-06-12 Steven Penny
----------

* [062883c](http://github.com/transcode-open/apt-cyg/commit/062883c)
  --version: fix license

2014-06-11 Steven Penny
----------

* [143e7a8](http://github.com/transcode-open/apt-cyg/commit/143e7a8)
  Change proxy list
* [71a3ce0](http://github.com/transcode-open/apt-cyg/commit/71a3ce0)
  Prep for new proxy
* [62cc432](http://github.com/transcode-open/apt-cyg/commit/62cc432)
  Update README

2014-06-09 Steven Penny
----------

* [b11ab2d](http://github.com/transcode-open/apt-cyg/commit/b11ab2d)
  searchall: filter out "Index of"
  "Index of" pages do not contain list of files, only list of packages.
* [29cd844](http://github.com/transcode-open/apt-cyg/commit/29cd844)
  Stop using cygwin.com for searchall
  After reading this message
  http://cygwin.com/ml/cygwin/2014-05/msg00482.html
  I set out to find a better option. I found a solution here
  http://stackoverflow.com/a/10554535
  Granted this is a deprecated API, if it actually dies it will be easy enough to
  replace with a scraper.

2014-06-05 Steven Penny
----------

* [81c64c9](http://github.com/transcode-open/apt-cyg/commit/81c64c9)
  Golf the ARCH code

2014-06-04 Steven Penny
----------

* [e5abc1c](http://github.com/transcode-open/apt-cyg/commit/e5abc1c)
  Use proper array for packages
  I think package names are not allowed to contain spaces anyway, but if they are
  it would cause apt-cyg to fail spectacularly. Some unquoted
      $pkg
  are still lingering but that is an easy fix.

2014-06-03 Steven Penny
----------

* [e989ec9](http://github.com/transcode-open/apt-cyg/commit/e989ec9)
  Minor edit
* [6fc1dac](http://github.com/transcode-open/apt-cyg/commit/6fc1dac)
  Support /dev/clipboard

2014-06-01 Steven Penny
----------

* [a6d4d8f](http://github.com/transcode-open/apt-cyg/commit/a6d4d8f)
  Minor edit "usage"
* [1272ac4](http://github.com/transcode-open/apt-cyg/commit/1272ac4)
  Use sed to print blocks of text
  The array method works fine, but sed method uses less characters and is easier
  to read. Also it does support indenting, it would just need to be modified to
    `sed '1d;$d;s/  //'`
  or similar.

2014-05-31 Steven Penny
----------

* [200ce90](http://github.com/transcode-open/apt-cyg/commit/200ce90)
  Search "devel" packages unless looking for ".exe"
  apt-cyg searchall will now only exclude "devel" packages if you are searching
  for ".exe" or similar.

2014-05-28 Steven Penny
----------

* [c8fb260](http://github.com/transcode-open/apt-cyg/commit/c8fb260)
  Improve "depends" function
  - function will now print the packages at proper depths
  - added newline between each result if called with multiple packages

2014-05-27 Steven Penny
----------

* [3900245](http://github.com/transcode-open/apt-cyg/commit/3900245)
  listfiles download if necessary
  listfiles was only working on packages that were already installed. You should
  be able to list a packages files even if it is not installed, especially if it
  has already been downloaded.
  The function will now download the package if necessary, without installing, in
  order to view the package files.
  Other small changes
  - conditional echo instead of piping to "head" in certain loops
  - download function was exiting upon failure, when really it should just skip to
    next package
  - simplified a grep statement

2014-05-26 Steven Penny
----------

* [13a9f58](http://github.com/transcode-open/apt-cyg/commit/13a9f58)
  Download function
  Package download functionality has been moved into its own private function, to
  support future updates.

2014-05-25 Steven Penny
----------

* [6ec55ff](http://github.com/transcode-open/apt-cyg/commit/6ec55ff)
  msys2 create wrapper scripts
  The MSYS2 programs will not play nice with Cygwin paths. This change will create
  a wrapper script in "/usr/local/bin" for each ".exe" that runs the exe with a
  path it will recognize.
* [530e70c](http://github.com/transcode-open/apt-cyg/commit/530e70c)
  msys2 correct awk statement
  Files are not always at field 15, they are always last field though.
* [1ff704e](http://github.com/transcode-open/apt-cyg/commit/1ff704e)
  msys2 listfiles command
  This is similar to the "dpkg --listfiles" command.
* [feffce1](http://github.com/transcode-open/apt-cyg/commit/feffce1)
  change mode apt-msys2
* [d5a8c42](http://github.com/transcode-open/apt-cyg/commit/d5a8c42)
  apt-msys2
  This is a file similar to apt-cyg. It will install packages from the MSYS2
  project. It will only be practical for packages from the "MINGW" repo, as
  packages from the "MSYS2" repo would require the MSYS2 dll.
  I have a need for this because Cygwin Tcl/Tk requires X11. MSYS2 Tcl/Tk does not
  require X11, so it can run much faster.
* [461d595](http://github.com/transcode-open/apt-cyg/commit/461d595)
  Remove redundant license
  The license is already in the header of the file.

2014-05-24 Steven Penny
----------

* [fbf531e](http://github.com/transcode-open/apt-cyg/commit/fbf531e)
  Remove case statement
  by using functions instead of case statement
  - remove level of indentation
  - allow use of local variables via "local" keyword
* [ec36e46](http://github.com/transcode-open/apt-cyg/commit/ec36e46)
  Change default cache location
  apt-get uses /var/cache to store downloaded files
  http://askubuntu.com/q/178806/where-do-packages-installed-with-apt-get-stored
  This makes sense with Cygwin as well because Cygwin already has a /var/cache
  which is almost empty to start, compared to /setup which does not exist outside
  of apt-cyg.
* [6a62641](http://github.com/transcode-open/apt-cyg/commit/6a62641)
  Revise some awk statements

2014-05-21 Steven Penny
----------

* [73512e4](http://github.com/transcode-open/apt-cyg/commit/73512e4)
  Restore noscripts option
  This is needed as per @kou1okada comment on 44e89c6
* [e0f31ae](http://github.com/transcode-open/apt-cyg/commit/e0f31ae)
  Improve "depends" awk statement, part deux

2014-05-20 Steven Penny
----------

* [d124c46](http://github.com/transcode-open/apt-cyg/commit/d124c46)
  Improve "depends" awk statement
  We can save a set of braces by using "continue". Also removed a spurious
  variable assignment.
* [a1cf729](http://github.com/transcode-open/apt-cyg/commit/a1cf729)
  depends command
  This is similar to the "apt-cache depends" command.

2014-05-19 Steven Penny
----------

* [4818a59](http://github.com/transcode-open/apt-cyg/commit/4818a59)
  Normalize mirror directory
  User input was not being checked in any way. "mirror" must not have a trailing
  slash, while "mirrordir" must have a trailing slash before it is encoded.
  This is to match up with "setup-x86_64.exe" installer.

2014-05-15 Steven Penny
----------

* [b09f55c](http://github.com/transcode-open/apt-cyg/commit/b09f55c)
  Improve remove logic
  Thanks @kou1okada
  http://github.com/transcode-open/apt-cyg/commit/1b3c49d#commitcomment-6329645

2014-05-14 Steven Penny
----------

* [a9f9fea](http://github.com/transcode-open/apt-cyg/commit/a9f9fea)
  Minor edit
  This will allow some relief from that long line, especially if more commands are
  added.
* [c38e030](http://github.com/transcode-open/apt-cyg/commit/c38e030)
  Minor cleanup
  - test commands directly instead of using temporary variable
  - test command directly instead of exit variable
* [f4d57c9](http://github.com/transcode-open/apt-cyg/commit/f4d57c9)
  Minor fix GitHub highlighting
  GitHub is freaking about '\' even though it is valid syntax. Using double
  backslash still works with "awk", and doesnt break the highlighting.
* [1b3c49d](http://github.com/transcode-open/apt-cyg/commit/1b3c49d)
  Revise remove command
  1. print a list of all files and dependencies of those files used by apt-cyg
  2. if any are found in the files list of the package to be removed, then fail
  Unrelated, added an echo each time a postinstall script is run, the name of the
  script
* [3ecc4ee](http://github.com/transcode-open/apt-cyg/commit/3ecc4ee)
  Update setup.ini when necessary
  If you tried to use this command
      apt-cyg install --mirror <URL> <package>
  with a new mirror, it would fail because "apt-cyg install" no longer downloads
  setup.ini every time. Fix so that setup.ini will be downloaded if it does not
  already exist.
* [44e89c6](http://github.com/transcode-open/apt-cyg/commit/44e89c6)
  Remove noscripts option
  - "--noscripts" was not advertised in the help, and is not really needed now
    that updates are done manually
  - moved the function inline as it was only being called once

2014-05-13 Steven Penny
----------

* [c585c8e](http://github.com/transcode-open/apt-cyg/commit/c585c8e)
  Allow remove wget
  wget is not included with base Cygwin, so I see no reason why it cannot be
  removed. Worst case it can be reinstalled using setup-x86_64.exe
  Also
  - removed some quotes where they were not needed
  - replaced a statement piping to Bash with more appropriate xargs statement

2014-05-11 Steven Penny
----------

* [2e6305c](http://github.com/transcode-open/apt-cyg/commit/2e6305c)
  Add more filters for "searchall"

2014-05-05 Steven Penny
----------

* [a2ca937](http://github.com/transcode-open/apt-cyg/commit/a2ca937)
  Update checkpackages function
  This function doesnt really need to exit, because package loops with no
  packages will simply exit the loop immediately. Also by removing the "exit" it
  allows to do cool stuff like
      if checkpackages
* [1cfa858](http://github.com/transcode-open/apt-cyg/commit/1cfa858)
  listfiles command
  This is similar to the "dpkg --listfiles" command.

2014-05-04 Steven Penny
----------

* [97ae776](http://github.com/transcode-open/apt-cyg/commit/97ae776)
  Stop regexing "apt-cyg show"
  Packages supplied to "apt-cyg show" must now match exactly. This is in line
  with "apt-cache show".
  Also if a package is not found it will now give a message.
* [9f1fb56](http://github.com/transcode-open/apt-cyg/commit/9f1fb56)
  Stop updating setup.ini
  Applies to
      apt-cyg install
      apt-cyg list
      apt-cyg show
  This is in line with these analogues
      apt-get install
      dpkg --list
      apt-cache show
  which do not run an update each time the command is called
* [74da8e7](http://github.com/transcode-open/apt-cyg/commit/74da8e7)
  searchall, skip first line
  The "searchall" command was including the first line of the HTTP response when
  looking for package names, example
      Found  7 matches for /diff.exe
  If the search contained a "/" then the search itself would show up as a result
  package. New awk statement filters out first line.

2014-04-18 Steven Penny
----------

* [3d0f4a0](http://github.com/transcode-open/apt-cyg/commit/3d0f4a0)
  Fix install directory
  packages were installing to "$cache/$mirrordir" which I believe is a relic of
  pre 64-bit Cygwin. Packages now install to "$cache/$mirrordir/$ARCH" which is
  inline with the official Cygwin installer.
* [7fedcba](http://github.com/transcode-open/apt-cyg/commit/7fedcba)
  rdepends command
  This is similar to the "apt-cache rdepends" command.

2014-03-16 Steven Penny
----------

* [223654f](http://github.com/transcode-open/apt-cyg/commit/223654f)
  searchall command
  This is similar to the "apt-file search" command, and almost identical to the
  "cygcheck --package-query" command. Instead of only search for a file from
  installed packages, it searches all available packages.
  I have included it as a separate command because it "wget"s cygwin.com and can
  be slow at times.

2014-03-14 Steven Penny
----------

* [bbc30a0](http://github.com/transcode-open/apt-cyg/commit/bbc30a0)
  Update README.md
* [841924f](http://github.com/transcode-open/apt-cyg/commit/841924f)
  Change "apt-cyg packageof" to "apt-cyg search"
  This puts apt-cyg more in line with "dpkg --search"
* [29014aa](http://github.com/transcode-open/apt-cyg/commit/29014aa)
  Merge "find" and "list" commands
  "list" will now list all installed packages, or packages matching pattern if one
  is given. This puts it more in line with "dpkg --list"
* [ecef939](http://github.com/transcode-open/apt-cyg/commit/ecef939)
  Change "apt-cyg describe" to "apt-cyg show"
  This puts apt-cyg more in line with "apt-cache show".
* [8229393](http://github.com/transcode-open/apt-cyg/commit/8229393)
  Change "apt-cyg show" to "apt-cyg list"
  This puts apt-cyg more in line with "dpkg --list".
  The plan is to merge "apt-cyg show" and "apt-cyg find" into "apt-cyg list".

2014-03-09 Steven Penny
----------

* [ca292c5](http://github.com/transcode-open/apt-cyg/commit/ca292c5)
  Minor cleanup
  - remove quote where possible
  - remove temporary variables where it makes sense
  - use single quote over double quote where possible

2014-03-08 Steven Penny
----------

* [663a6d5](http://github.com/transcode-open/apt-cyg/commit/663a6d5)
  Code clean up part 4
  - use "[" instead of "test"
  - remove double quotes where word splitting does not occur
  - revise some "sed" statements
  - use "((" construct where possible
  - use "return" in place of long winded "if" statements
  - use "+=" where possible
  - use "type" instead of "which"
  - use "[[" construct where appropriate
* [b42a6cb](http://github.com/transcode-open/apt-cyg/commit/b42a6cb)
  Update dependency and arch checking
  Use more concise methods for both.

2014-03-07 Steven Penny
----------

* [041bc88](http://github.com/transcode-open/apt-cyg/commit/041bc88)
  Code clean up part 3
  - avoid double quotes where possible
  - replace 'echo ""' with 'echo'
  - remove "useless use of cat"
  - "awk" statement refactoring

2014-03-06 Steven Penny
----------

* [54d23fc](http://github.com/transcode-open/apt-cyg/commit/54d23fc)
  Code clean up part 2
  - avoid using escaped double quotes where possible
  - use "+=" where possible
  - remove "useless use of cat"
  - "awk" statement refactoring
  - replace 'echo ""' with 'echo'
* [135b3ab](http://github.com/transcode-open/apt-cyg/commit/135b3ab)
  Code clean up
  - utilize "||" as continuation character, remove unneeded "\"
  - remove excess empty lines
  - replace "shift; shift" with "shift 2"
  - move redirects to end of line
  - removed several "useless use of cat"
  - refactored several "awk" statements
  - replace 'echo ""' with 'echo'

2014-03-05 Steven Penny
----------

* [01244a6](http://github.com/transcode-open/apt-cyg/commit/01244a6)
  change mode apt-cyg

2014-02-28 Steven Penny
----------

* [41e1d91](http://github.com/transcode-open/apt-cyg/commit/41e1d91)
  bootstrap wget
  Currently apt-cyg has a circular dependency on wget. This can be resolved the
  following ways, but each has problems
  - setup-x86_64
    you can run a command like "setup-x86_64 -qP wget" but this is not desirable
    because it just spawns the gui, which was probably what apt-cyg was designed
    to prevent in the first place
  - using /dev/tcp
    this is a good solution but does not work for "https://" URLs
  - using telnet
    this is a windows component but is disabled by default
  - using curl
    like wget, this does not come with base Cygwin
  - using powershell
    this is available for XP and later, but only comes preinstalled with Windows 7
    and later
  This commit would allow a "wget" Bash function to be used if "wget.exe" is not
  found, for example
  http://github.com/svnpenn/dotfiles/blob/19b7b76/.bashrc#L54-L67
  My function uses powershell, but a function using any of the above methods would
  work.

2014-02-19 Stephen Lang
----------

* [c093abf](http://github.com/transcode-open/apt-cyg/commit/c093abf)
  Fix svn merge conflicts
* [58b3481](http://github.com/transcode-open/apt-cyg/commit/58b3481)
  Change default mirror to mirrors.kernel.org
* [53848c2](http://github.com/transcode-open/apt-cyg/commit/53848c2)
  Update README.md
* [ae383fb](http://github.com/transcode-open/apt-cyg/commit/ae383fb)
  Readme update.
* [b5a0136](http://github.com/transcode-open/apt-cyg/commit/b5a0136)
  Add xz archive support.
* [ff29e0f](http://github.com/transcode-open/apt-cyg/commit/ff29e0f)
  Handle the x86/x86_64 split when downloading setup.ini
  Cygwin recently introduced an x86_64 version and so all the mirrors
  now have x86 or x86_64 in their paths, depending on which version
  you've installed.
  The only change required is when the setup file is fetched.
  The setup file itself contains the paths prefixed with either
  x86 or x86_64

2014-02-19 Stephen Lang
----------

* [60a31c5](http://github.com/transcode-open/apt-cyg/commit/60a31c5)
  Change default mirror to mirrors.kernel.org
* [c42c442](http://github.com/transcode-open/apt-cyg/commit/c42c442)
  Added README

2013-08-15 Ryan Duryea
----------

* [4b71f04](http://github.com/transcode-open/apt-cyg/commit/4b71f04)
  fetch setup.ini using $arch in mirror path as well
  boothj5 pointed out that the script also needs the $arch
  in the URL when fetching setup.ini as a fall back to
  fetching setup.bz2

2013-08-09 Ryan Duryea
----------

* [0292520](http://github.com/transcode-open/apt-cyg/commit/0292520)
  Handle the x86/x86_64 split when downloading setup.ini
  Cygwin recently introduced an x86_64 version and so all the mirrors
  now have x86 or x86_64 in their paths, depending on which version
  you've installed.
  The only change required is when the setup file is fetched.
  The setup file itself contains the paths prefixed with either
  x86 or x86_64

2014-03-16 Steven Penny
----------

* [843a530](http://github.com/transcode-open/apt-cyg/commit/843a530)
  Merge commit 'origin~21'

2014-02-19 Stephen Lang
----------

* [37d6173](http://github.com/transcode-open/apt-cyg/commit/37d6173)
  Update README.md
* [7c1cecb](http://github.com/transcode-open/apt-cyg/commit/7c1cecb)
  Update README.md
* [2f84a3f](http://github.com/transcode-open/apt-cyg/commit/2f84a3f)
  Initial commit

2014-02-17 Stephen Lang
----------

* [45de191](http://github.com/transcode-open/apt-cyg/commit/45de191)
  Merge branch 0.58
* [0f361a1](http://github.com/transcode-open/apt-cyg/commit/0f361a1)
  Detect CPU architecture for new mirror paths
* [09bf3de](http://github.com/transcode-open/apt-cyg/commit/09bf3de)
  Creating version 0.58 branch

2013-10-23 Leszek Cimala
----------

* [46fb5d3](http://github.com/transcode-open/apt-cyg/commit/46fb5d3)
  Update README.md

2013-10-23 Ernie Rasta
----------

* [6c2cc7c](http://github.com/transcode-open/apt-cyg/commit/6c2cc7c)
  Readme update.
* [c437fa0](http://github.com/transcode-open/apt-cyg/commit/c437fa0)
  Add multiarch support.
* [9058643](http://github.com/transcode-open/apt-cyg/commit/9058643)
  Add xz archive support.

2013-07-28 Stephen Jungels
----------

* [fc54d5c](http://github.com/transcode-open/apt-cyg/commit/fc54d5c)
  Update README.md
* [8d8dd65](http://github.com/transcode-open/apt-cyg/commit/8d8dd65)
  Update README.md
* [2af0c86](http://github.com/transcode-open/apt-cyg/commit/2af0c86)
  Create apt-cyg
* [ab235c0](http://github.com/transcode-open/apt-cyg/commit/ab235c0)
  Initial commit

2010-04-27 Stephen Jungels
----------

* [ee3f7b8](http://github.com/transcode-open/apt-cyg/commit/ee3f7b8)
  added edge case for obsolete packages
* [25a51eb](http://github.com/transcode-open/apt-cyg/commit/25a51eb)
  1. tightened regex install expression; 2. used wget -N where appropriate

2009-08-10 Stephen Jungels
----------

* [e8891b5](http://github.com/transcode-open/apt-cyg/commit/e8891b5)
  Merged patch that adds --noupdate flag

2009-06-14 Stephen Jungels
----------

* [cbff7c9](http://github.com/transcode-open/apt-cyg/commit/cbff7c9)
  Accepted patch that prevents extra downloads of setup.bz2

2009-02-27 Stephen Jungels
----------

* [e81c40c](http://github.com/transcode-open/apt-cyg/commit/e81c40c)
  Removed cruft
* [c3ae41f](http://github.com/transcode-open/apt-cyg/commit/c3ae41f)
  Initial import
* [acabffd](http://github.com/transcode-open/apt-cyg/commit/acabffd)
  Initial directory structure.

