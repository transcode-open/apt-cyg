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

ARCH=${HOSTTYPE/i6/x}

function usage {
  hr '
  usage: apt-cyg [command] [options] [packages]

  Commands:
     install     Install packages
     remove      Remove packages
     update      Update setup.ini
     download    Download only - do NOT install or unpack archives
     show        Displays the package records for the named packages
     depends     Performs recursive dependency listings
     rdepends    Display packages which require X to be installed,
                 AKA show reverse dependencies
     list        List packages matching given pattern. If no pattern is given,
                 list all installed packages.
     listfiles   List files owned by packages
     search      Search for a filename from installed packages
     searchall   Search for a filename from all available packages

  Options:
     -c, --cache <dir>      set cache
     -f, --file <file>      read package names from file
     -m, --mirror <url>     set mirror
     --help
     --version
  '
}

function version {
  hr '
  apt-cyg version 0.59

  The MIT License (MIT)

  Copyright (c) 2005-9 Stephen Jungels
  '
}

function find-workspace {
  # default working directory and mirror
  
  # work wherever setup worked last, if possible
  cache=$(awk '
  /last-cache/ {
    getline
    print $1
  }
  ' /etc/setup/setup.rc | cygpath -f-)

  mirror=$(awk '
  /last-mirror/ {
    getline
    print $1
  }
  ' /etc/setup/setup.rc)
  mirrordir=$(sed '
  s / %2f g
  s : %3a g
  ' <<< "$mirror")

  mkdir -p "$cache/$mirrordir/$ARCH"
  cd "$cache/$mirrordir/$ARCH"
  if [ -e setup.ini ]
  then
    return 0
  else
    get-setup
    return 1
  fi
}

function get-setup {
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

function check-packages {
  if [[ $pks ]]
  then
    return 0
  else
    echo No packages found
    return 1
  fi
}

function warn {
  printf '\e[36m%s\e[m\n' "$*" >&2
}

function apt-update {
  if find-workspace
  then
    get-setup
  fi
}

function apt-list {
  find-workspace
  local sbq
  for pkg in "${pks[@]}"
  do
    (( sbq++ )) && echo
    warn Searching for installed packages matching $pkg:
    awk 'NR>1 && $1~query && $0=$1' query="$pkg" /etc/setup/installed.db
    echo
    warn Searching for installable packages matching $pkg:
    awk '$1 ~ query && $0 = $1' RS='\n\n@ ' FS='\n' query="$pkg" setup.ini
  done
  (( sbq )) && return
  warn The following packages are installed:
  awk 'NR>1 && $0=$1' /etc/setup/installed.db
}

function apt-listfiles {
  check-packages
  find-workspace
  local pkg sbq
  for pkg in "${pks[@]}"
  do
    (( sbq++ )) && echo
    if [ ! -e /etc/setup/"$pkg".lst.gz ]
    then
      download "$pkg"
    fi
    gzip -cd /etc/setup/"$pkg".lst.gz
  done
}

function apt-show {
  find-workspace
  check-packages
  for pkg in "${pks[@]}"
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

function apt-depends {
  find-workspace
  check-packages
  for pkg in "${pks[@]}"
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

function apt-rdepends {
  find-workspace
  for pkg in "${pks[@]}"
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

function apt-download {
  check-packages
  find-workspace
  local pkg sbq
  for pkg in "${pks[@]}"
  do
    (( sbq++ )) && echo
    download "$pkg"
  done
}

function download {
  local pkg file digest digactual
  pkg=$1
  # look for package and save desc file

  awk '$1 == pc' RS='\n\n@ ' FS='\n' pc=$pkg setup.ini > desc
  if [ ! -s desc ]
  then
    echo Unable to locate package $pkg
    exit 1
  fi

  # download and unpack the bz2 or xz file

  # pick the latest version, which comes first
  set -- $(awk '$1 == "install:"' desc)
  if (( ! $# ))
  then
    echo 'Could not find "install" in package description: obsolete package?'
    exit 1
  fi

  dn=$(dirname $2)
  bn=$(basename $2)

  # check the md5
  digest=$4
  mkdir -p $cache/$mirrordir/$dn
  cd $cache/$mirrordir/$dn
  wget -nc $mirror/$dn/$bn
  digactual=$(md5sum $bn | awk NF=1)
  if [ $digest != $digactual ]
  then
    echo MD5 sum did not match, exiting
    exit 1
  fi

  tar tf $bn | gzip > /etc/setup/"$pkg".lst.gz
  cd ~-
  mv desc $cache/$mirrordir/$dn
  echo $dn $bn > /tmp/dwn
}

function apt-search {
  check-packages
  echo Searching downloaded packages...
  for pkg in "${pks[@]}"
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

function apt-searchall {
  cd /tmp
  for pkg in "${pks[@]}"
  do
    printf -v qs 'text=1&arch=%s&grep=%s' $ARCH "$pkg"
    wget -O matches cygwin.com/cgi-bin2/package-grep.cgi?"$qs"
    awk '
    NR == 1                   {next}
    /-doc-/                   {next}
    /-debuginfo-/             {next}
    /-devel-/ && pkg~/\.exe$/ {next}
    /-src\t$/                 {next}
    mc[$1]++                  {next}
    $0 = $1
    ' FS=-[[:digit:]] pkg="$pkg" matches
  done
}

function apt-install {
  check-packages
  find-workspace
  local pkg dn bn requires wr package script
  for pkg in "${pks[@]}"
  do

  if grep -q "^$pkg " /etc/setup/installed.db
  then
    echo Package $pkg is already installed, skipping
    continue
  fi
  echo
  echo Installing $pkg

  download $pkg
  read dn bn </tmp/dwn
  echo Unpacking...

  cd $cache/$mirrordir/$dn
  tar -x -C / -f $bn
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
  ' pkg="$pkg" bz=$bn /etc/setup/installed.db > /tmp/awk.$$
  mv /etc/setup/installed.db /etc/setup/installed.db-save
  mv /tmp/awk.$$ /etc/setup/installed.db

  # recursively install required packages

  requires=$(awk '$1=="requires", $0=$2' FS=': ' desc)
  cd ~-
  wr=0
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
      apt-cyg install --noscripts $package || (( wr++ ))
    done
  fi
  if (( wr ))
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

function apt-remove {
  check-packages
  for pkg in "${pks[@]}"
  do

  if ! grep -q "^$pkg " /etc/setup/installed.db
  then
    echo Package $pkg is not installed, skipping
    continue
  fi

  cygcheck awk bash bunzip2 grep gzip mv sed tar xargs xz | awk '
  /bin/      &&
  !fd[$NF]++ &&
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

function hr {
  sed '
  1d
  $d
  s/  //
  ' <<< "$1"
}

# process options
while (( $# ))
do
  case "$1" in

    --mirror | -m)
      awk -i inplace '
      1
      /last-mirror/ {
        getline
        print "\t" rpc
      }
      ' rpc="$2" /etc/setup/setup.rc
      shift 2
    ;;

    --cache | -c)
      rpc=$(cygpath -aw "$2" | sed 's \\ \\\\ g')
      awk -i inplace '
      1
      /last-cache/ {
        getline
        print "\t" rpc
      }
      ' rpc="$rpc" /etc/setup/setup.rc
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
          mapfile -t pks <<< "$fp"
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
    | remove    \
    | download  \
    | show      \
    | depends   \
    | rdepends  \
    | list      \
    | listfiles \
    | search    \
    | searchall)
      if [[ $command ]]
      then
        pks+=("$1")
      else
        command=$1
      fi
      shift
    ;;

    *)
      pks+=("$1")
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
