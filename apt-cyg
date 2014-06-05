#!/bin/bash
# apt-cyg: install tool for Cygwin similar to debian apt-get
#
# The MIT License (MIT)
#
# Copyright (c) 2013 Trans-code Design
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

# this script requires some packages
if ! type awk bzip2 tar wget xz &>/dev/null
then
  echo You must install wget, tar, gawk, xz and bzip2 to use apt-cyg.
  exit 1
fi

ARCH=$(arch | sed s.i6.x.)

function usage () {
sed '1d;$d' <<< '
usage: apt-cyg [command] [options] [packages]

Commands:
   install <packages>     install packages
   remove <packages>      remove packages
   update                 update setup.ini
   list [patterns]        list packages matching given pattern. If no
                          pattern is given, list all installed packages.
   listfiles <packages>   list files owned by packages
   show <packages>        Displays the package records for the named packages
   depends <patterns>     performs recursive dependency listings
   rdepends <patterns>    Display packages which require X to be installed,
                          AKA show reverse dependencies
   search <patterns>      search for a filename from installed packages
   searchall <patterns>   search for a filename from all available packages

Options:
   -c, --cache <dir>      set cache
   -f, --file <file>      read package names from file
   -m, --mirror <url>     set mirror
   --help
   --version
'
}

function version () {
sed '1d;$d' <<< '
apt-cyg version 0.59
Written by Stephen Jungels

Copyright (c) 2005-9 Stephen Jungels.  Released under the GPL.
'
}

function findworkspace()
{
  # default working directory and mirror
  mirror=http://mirrors.kernel.org/sourceware/cygwin
  cache=/var/cache
  
  # work wherever setup worked last, if possible
  if [ -e /etc/setup/last-cache ]
  then
    cache=$(cygpath -f /etc/setup/last-cache)
  fi

  if [ -e /etc/setup/last-mirror ]
  then
    mirror=$(sed 's./$..' /etc/setup/last-mirror)
  fi
  mirrordir=$(sed '
  s / %2f g
  s : %3a g
  ' <<< "$mirror")

  mkdir -p "$cache/$mirrordir/$ARCH"
  cd "$cache/$mirrordir/$ARCH"
  [ -e setup.ini ] || getsetup
}

getsetup()
{
  touch setup.ini
  mv setup.ini setup.ini-save
  wget -N $mirror/$ARCH/setup.bz2
  if [ -e setup.bz2 ]
  then
    bunzip2 setup.bz2
    mv setup setup.ini
    echo Updated setup.ini
  else
    echo Error updating setup.ini, reverting
    mv setup.ini-save setup.ini
  fi
}

function checkpackages()
{
  if [[ $packages ]]
  then
    return 0
  else
    echo Nothing to do.
    return 1
  fi
}

apt-update () {
  findworkspace
  getsetup
}

apt-list () {
  if checkpackages
  then
    findworkspace
    for pkg in "${packages[@]}"
    do
      echo
      echo Searching for installed packages matching $pkg:
      awk 'NR>1 && $1~query && $0=$1' query="$pkg" /etc/setup/installed.db
      echo
      echo Searching for installable packages matching $pkg:
      awk '$1 ~ query && $0 = $1' RS='\n\n@ ' FS='\n' query="$pkg" setup.ini
    done
  else
    echo The following packages are installed: >&2
    awk 'NR>1 && $0=$1' /etc/setup/installed.db
  fi
}

apt-listfiles () {
  checkpackages
  findworkspace
  local pkg
  for pkg in "${packages[@]}"
  do
    (( notfirst++ )) && echo
    if [ ! -e /etc/setup/"$pkg".lst.gz ]
    then
      download "$pkg" || continue
    fi
    gzip -cd /etc/setup/"$pkg".lst.gz
  done
}

apt-show () {
  findworkspace
  checkpackages
  for pkg in "${packages[@]}"
  do
    (( notfirst++ )) && echo
    awk '
    $1 == query {
      print
      fd++
    }
    END {
      if (! fd)
        print "Unable to locate package " query
    }
    ' RS='\n\n@ ' FS='\n' query="$pkg" setup.ini
  done
}

apt-depends () {
  findworkspace
  checkpackages
  for pkg in "${packages[@]}"
  do
    (( sq++ )) && echo
    awk '
    $1 == "@" {
      pkg = $2
    }
    $1 == "requires:" {
      for (i=2; i<=NF; i++)
        reqs[pkg][i-1]=$i
    }
    END {
      setMinDepth(query)
      prtPkg(query)
    }
    function prtPkg(pkg,    req, i) {
      depth++
      if ( depth == mn[pkg] && !seen[pkg]++ ) {
        printf "%*s%s\n", indent, "", pkg
        indent += 2
        if (pkg in reqs)
          for (i=1; i in reqs[pkg]; i++) {
            req = reqs[pkg][i]
            prtPkg(req)
          }
        indent -= 2
      }
      depth--
    }
    function setMinDepth(pkg,    req, i) {
      depth++
      mn[pkg] = !(pkg in mn) || depth < mn[pkg] ? depth : mn[pkg]
      if (depth == mn[pkg]) {
        if (pkg in reqs)
          for (i=1; i in reqs[pkg]; i++) {
            req = reqs[pkg][i]
            setMinDepth(req)
          }
      }
      depth--
    }
    ' query="$pkg" setup.ini
  done
}

apt-rdepends () {
  findworkspace
  for pkg in "${packages[@]}"
  do
    awk '
    /^@ / {
      pn = $2
    }
    $0 ~ "^requires: .*"query {
      print pn
    }
    ' query="$pkg" setup.ini
  done
}

download () {
  local pkg install file digest digactual
  pkg=$1
  # look for package and save desc file

  mkdir -p release/"$pkg"
  awk '$1 == pc' RS='\n\n@ ' FS='\n' pc=$pkg setup.ini > release/$pkg/desc
  if [ ! -s release/$pkg/desc ]
  then
    echo Package $pkg not found or ambiguous name
    rm -r release/"$pkg"
    return 1
  fi
  echo Found package $pkg >&2

  # download and unpack the bz2 or xz file

  # pick the latest version, which comes first
  install=$(awk '/^install: / {print $2; exit}' release/"$pkg"/desc)
  if [[ ! $install ]]
  then
    echo 'Could not find "install" in package description: obsolete package?'
    exit 1
  fi

  file=$(basename $install)
  cd release/"$pkg"
  wget -nc $mirror/$install

  # check the md5
  digest=$(awk '/^install: / {print $4; exit}' desc)
  digactual=$(md5sum $file | awk NF=1)
  if [ $digest != $digactual ]
  then
    echo MD5 sum did not match, exiting
    exit 1
  fi

  tar tf $file | gzip > /etc/setup/"$pkg".lst.gz
  echo $file
}

apt-search () {
  checkpackages
  echo Searching downloaded packages...
  for pkg in "${packages[@]}"
  do
    key=$(type -P "$pkg" | sed s./..)
    [[ $key ]] || key=$pkg
    for manifest in /etc/setup/*.lst.gz
    do
      if gzip -cd $manifest | grep -q "$key"
      then
        package=$(sed '
        s,/etc/setup/,,
        s,.lst.gz,,
        ' <<< $manifest)
        echo $package
      fi
    done
  done
}

apt-searchall () {
  for pkg in "${packages[@]}"
  do
    printf -v qs 'text=1&arch=%s&grep=%s' $ARCH "$pkg"
    cd /tmp
    wget -O matches cygwin.com/cgi-bin2/package-grep.cgi?"$qs"
    awk '
    NR == 1                   {next}
    /-doc-/                   {next}
    /-debuginfo-/             {next}
    /-devel-/ && pkg~/\.exe$/ {next}
    /-src\t$/                 {next}
    mc[$2]++                  {next}
    $0 = $2
    ' FS=/ pkg="$pkg" matches
  done
}

apt-install () {
  findworkspace
  checkpackages
  local pkg file requires warn package script
  for pkg in "${packages[@]}"
  do

  if grep -q "^$pkg " /etc/setup/installed.db
  then
    echo Package $pkg is already installed, skipping
    continue
  fi
  echo
  echo Installing $pkg

  file=$(download $pkg) || continue
  cd release/$pkg
  echo Unpacking...

  tar -x -C / -f $file
  # update the package database

  awk '
  ins != 1 && pkg < $1 {
    print pkg, bz, 0
    ins = 1
  }
  1
  END {
    if (ins != 1) print pkg, bz, 0
  }
  ' pkg="$pkg" bz=$file /etc/setup/installed.db > /tmp/awk.$$
  mv /etc/setup/installed.db /etc/setup/installed.db-save
  mv /tmp/awk.$$ /etc/setup/installed.db

  # recursively install required packages

  requires=$(awk '$1=="requires", $0=$2' FS=': ' desc)
  cd ~-
  warn=0
  if [[ $requires ]]
  then
    echo Package $pkg requires the following packages, installing:
    echo $requires
    for package in $requires
    do
      if grep -q "^$package " /etc/setup/installed.db
      then
        echo Package $package is already installed, skipping
        continue
      fi
      apt-cyg install --noscripts $package || (( warn++ ))
    done
  fi
  if (( warn ))
  then
    echo 'Warning: some required packages did not install, continuing'
  fi

  # run all postinstall scripts

  (( noscripts )) && continue
  find /etc/postinstall -name '*.sh' | while read script
  do
    echo Running $script
    $script
    mv $script $script.done
  done
  echo Package $pkg installed

  done
}

apt-remove () {
  checkpackages
  for pkg in "${packages[@]}"
  do

  if ! grep -q "^$pkg " /etc/setup/installed.db
  then
    echo Package $pkg is not installed, skipping
    continue
  fi

  cygcheck awk bash bunzip2 grep gzip mv sed tar xargs xz | awk '
  /bin/ &&
  ! fd[$NF]++ &&
  $0 = $NF
  ' FS='\\' > /tmp/cygcheck.txt

  apt-cyg listfiles $pkg | awk '
  $0 = $NF
  ' FS=/ > /tmp/listfiles.txt

  if grep -xf /tmp/cygcheck.txt /tmp/listfiles.txt
  then
    echo apt-cyg cannot remove package $pkg, exiting
    exit 1
  fi
  if [ ! -e /etc/setup/"$pkg".lst.gz ]
  then
    echo Package manifest missing, cannot remove $pkg. Exiting
    exit 1
  fi
  echo Removing $pkg

  # run preremove scripts

  if [ -e /etc/preremove/"$pkg".sh ]
  then
    /etc/preremove/"$pkg".sh
    rm /etc/preremove/"$pkg".sh
  fi
  gzip -cd /etc/setup/"$pkg".lst.gz | sed '\./$.d;s.^./.' | xargs rm -f
  rm /etc/setup/"$pkg".lst.gz
  rm -f /etc/postinstall/$pkg.sh.done
  awk '$1 != pkg' pkg="$pkg" /etc/setup/installed.db > /tmp/awk.$$
  mv /etc/setup/installed.db /etc/setup/installed.db-save
  mv /tmp/awk.$$ /etc/setup/installed.db
  echo Package $pkg removed

  done
}

# process options
while (( $# ))
do
  case "$1" in

    --mirror | -m)
      echo "$2" > /etc/setup/last-mirror
      shift 2
    ;;

    --cache | -c)
      cygpath -aw "$2" > /etc/setup/last-cache
      shift 2
    ;;

    --noscripts)
      noscripts=1
      shift
    ;;

    --help)
      usage
      exit 0
    ;;

    --version)
      version
      exit 0
    ;;

    --file | -f)
      if [[ $2 ]]
      then
        file=$2
        # support /dev/clipboard
        if [ -c "$file" -o -f "$file" ]
        then
          fp=$(sed '' "$file")
          mapfile -t packages <<< "$fp"
        else
          echo File $file not found, skipping
        fi
        shift
      else
        echo No file name provided, ignoring $1 >&2
      fi
      shift
    ;;

    update)
      command=$1
      shift
    ;;

    install     \
    | list      \
    | listfiles \
    | depends   \
    | rdepends  \
    | search    \
    | searchall \
    | remove    \
    | show)
      if [[ $command ]]
      then
        packages+=("$1")
      else
        command=$1
      fi
      shift
    ;;

    *)
      packages+=("$1")
      shift
    ;;

  esac
done

if type -t apt-$command | grep -q function
then
  apt-$command
else
  usage
fi
