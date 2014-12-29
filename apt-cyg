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

mapfile usage <<+
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
   category    List packages matching given category
   listfiles   List files owned by packages
   search      Search for a filename from installed packages
   searchall   Search for a filename from all available packages

Options:
   --cache <dir>    set cache
   --mirror <url>   set mirror
   --version
+

mapfile version <<+
apt-cyg version 0.59

The MIT License (MIT)

Copyright (c) 2005-9 Stephen Jungels
+

readonly usage version

function wget {
  if command wget -h &>/dev/null
  then
    command wget "$@"
  else
    warn wget is not installed, using lynx as fallback
    set "${*: -1}"
    lynx -source "$1" > "${1##*/}"
  fi
}

function find-workspace {
  # default working directory and mirror
  
  # work wherever setup worked last, if possible
  readonly cache=$(awk '
  /last-cache/ {
    getline
    print $1
  }
  ' /etc/setup/setup.rc | cygpath -f-)

  readonly mirror=$(awk '
  /last-mirror/ {
    getline
    print $1
  }
  ' /etc/setup/setup.rc)
  readonly mirrordir=$(sed '
  s / %2f g
  s : %3a g
  ' <<< "$mirror")

  mkdir -p "$cache/$mirrordir/$arch"
  cd "$cache/$mirrordir/$arch"
  if [ -e setup.ini ]
  then
    return 0
  else
    get-setup
    return 1
  fi
}

function get-setup {
  if wget -N $mirror/$arch/setup.bz2
  then
    bunzip2 setup.bz2
    mv setup setup.ini
    echo Updated setup.ini
  fi
}

function check-packages {
  if [ "$pks" ]
  then
    return 0
  else
    echo No packages found
    return 1
  fi
}

function info {
  printf '\e[36m%s\e[m\n' "$*" >&2
}

function warn {
  printf '\e[1;31m%s\e[m\n' "$*" >&2
}

function apt-update {
  if find-workspace
  then
    get-setup
  fi
}

function apt-category {
  check-packages
  find-workspace
  for vas in "${pks[@]}"
  do
    awk '
    $1 == "@" {
      pck = $2
    }
    $1 == "category:" && $0 ~ vas {
      print pck
    }
    ' vas="$vas" setup.ini
  done
}

function apt-list {
  find-workspace
  for vas in "${pks[@]}"
  do
    let sbq++ && echo
    info Searching for installed packages matching "$vas":
    awk 'NR>1 && $1~vas && $0=$1' vas="$vas" /etc/setup/installed.db
    echo
    info Searching for installable packages matching "$vas":
    awk '$1~vas && $0=$1' RS='\n\n@ ' FS='\n' vas="$vas" setup.ini
  done
  let sbq && return
  info The following packages are installed:
  awk 'NR>1 && $0=$1' /etc/setup/installed.db
}

function apt-listfiles {
  check-packages
  find-workspace
  for vas in "${pks[@]}"
  do
    let sbq++ && echo
    if [ ! -e /etc/setup/"$vas".lst.gz ]
    then
      download || return
    fi
    gzip -cd /etc/setup/"$vas".lst.gz
  done
}

function apt-show {
  find-workspace
  check-packages
  for vas in "${pks[@]}"
  do
    let sbq++ && echo
    awk '
    $1 == vas {
      print
      fd++
    }
    END {
      if (! fd)
        print "Unable to locate package", vas
    }
    ' RS='\n\n@ ' FS='\n' vas="$vas" setup.ini
  done
}

function apt-depends {
  find-workspace
  check-packages
  for vas in "${pks[@]}"
  do
    awk '
    @include "join"
    $1 == "@" {
      apg = $2
    }
    $1 == "requires:" {
      for (z=2; z<=NF; z++)
        reqs[apg][z-1] = $z
    }
    END {
      prpg(vas)
    }
    function smartmatch(small, large,    values) {
      for (each in large)
        values[large[each]]
      return small in values
    }
    function prpg(fpg) {
      if (smartmatch(fpg, spath)) return
      spath[length(spath)+1] = fpg
      print join(spath, 1, length(spath), " > ")
      if (isarray(reqs[fpg]))
        for (each in reqs[fpg])
          prpg(reqs[fpg][each])
      delete spath[length(spath)]
    }
    ' vas="$vas" setup.ini
  done
}

function apt-rdepends {
  check-packages
  find-workspace
  for vas in "${pks[@]}"
  do
    awk '
    @include "join"
    $1 == "@" {
      apg = $2
    }
    $1 == "requires:" {
      for (z=2; z<=NF; z++)
        reqs[$z][length(reqs[$z])+1] = apg
    }
    END {
      prpg(vas)
    }
    function smartmatch(small, large,    values) {
      for (each in large)
        values[large[each]]
      return small in values
    }
    function prpg(fpg) {
      if (smartmatch(fpg, spath)) return
      spath[length(spath)+1] = fpg
      print join(spath, 1, length(spath), " < ")
      if (isarray(reqs[fpg]))
        for (each in reqs[fpg])
          prpg(reqs[fpg][each])
      delete spath[length(spath)]
    }
    ' vas="$vas" setup.ini
  done
}

function apt-download {
  check-packages
  find-workspace
  for vas in "${pks[@]}"
  do
    let sbq++ && echo
    download
  done
}

function download {
  # look for package and save md5 file

  # download and unpack the bz2 or xz file

  # pick the latest version, which comes first
  awk '
  $1 == "@" {
    c = $2
  }
  $1 == "install:" && c == vas {
    print $4, $2
    exit
  }
  ' vas="$vas" setup.ini > md5.sum
  if [ ! -s md5.sum ]
  then
    warn Unable to locate package $vas
    return 1
  fi

  read _ acv < md5.sum

  # check the md5
  mv md5.sum ..
  cd ..
  if ! test -e $acv || ! md5sum -c md5.sum
  then
    wget -rnH $mirror/$acv
    md5sum -c md5.sum || return
  fi

  tar tf $acv | gzip > /etc/setup/"$vas".lst.gz
}

function apt-search {
  check-packages || return
  echo Searching downloaded packages...
  for vas in "${pks[@]}"
  do
    for mft in /etc/setup/*.lst.gz
    do
      if gzip -cd $mft | grep -q "$vas"
      then
        awk '$0=$4' FS='[./]' <<< $mft
      fi
    done
  done
}

function apt-searchall {
  check-packages
  cd /etc/setup
  for vas in "${pks[@]}"
  do
    printf -v qry 'text=1&arch=%s&grep=%s' $arch "$vas"
    wget -O matches cygwin.com/cgi-bin2/package-grep.cgi?"$qry"
    awk '
    {
      if (NR == 1)
        next
      if (mc[$1]++)
        next
      if (/-debuginfo-/)
        next
      print $1
    }
    ' FS=-[[:digit:]] matches
  done
}

function set-cache {
  if [ "$1" ]
  then
    vas=$(cygpath -aw "$1")
    awk -i inplace '
    1
    /last-cache/ {
      getline
      print "\t" vas
    }
    ' vas="${vas//\\/\\\\}" /etc/setup/setup.rc
    echo Cache set to "$vas"
  else
    warn No path provided, exiting
  fi
}

function set-mirror {
  if [ "$1" ]
  then
    awk -i inplace '
    1
    /last-mirror/ {
      getline
      print "\t" vas
    }
    ' vas="$1" /etc/setup/setup.rc
    echo Mirror set to "$1"
  else
    warn No URL provided, exiting
  fi
}

function apt-install {
  check-packages
  find-workspace
  for vas in "${pks[@]}"
  do

  if grep -q "^$vas " /etc/setup/installed.db
  then
    echo Package $vas is already installed, skipping
    continue
  fi
  let sbq++ && echo
  echo Installing $vas

  download || return
  echo Unpacking...

  tar -x -C / -f $acv
  cd ~-
  # update the package database

  awk -e '
  {
    foo[$0]
  }
  ENDFILE {
    $1 = vas
    $2 = acv
    $3 = 0
    foo[$0]
    asorti(foo)
    for (bar in foo)
      print foo[bar]
  }
  ' -i inplace vas="$vas" acv=$acv /etc/setup/installed.db

  # recursively install required packages

  awk '
  $1 == "@" {
    c = $2
  }
  $1 == "requires:" && c == vas {
    print
    exit
  }
  ' vas="$vas" setup.ini > deps.ini
  if [ -s deps.ini ]
  then
    read _ rqs < deps.ini
    echo Package $vas requires the following packages, installing:
    echo $rqs
    apt-cyg install --deps $rqs ||
    info some required packages did not install, continuing
  fi

  # run all postinstall scripts

  let dps && continue
  find /etc/postinstall -name '*.sh' | while read script
  do
    echo Running $script
    $script
    mv $script $script.done
  done
  echo Package $vas installed

  done
}

function apt-remove {
  check-packages
  cd /etc
  cygcheck awk bash bunzip2 grep gzip mv sed tar xz > setup/essential.lst
  for vas in "${pks[@]}"
  do

  if ! grep -q "^$vas " setup/installed.db
  then
    echo Package $vas is not installed, skipping
    continue
  fi

  if [ ! -e setup/"$vas".lst.gz ]
  then
    warn Package manifest missing, cannot remove $vas. Exiting
    return 1
  fi
  gzip -dk setup/"$vas".lst.gz
  awk '
  NR == FNR {
    if ($NF) ess[$NF]
    next
  }
  $NF in ess {
    exit 1
  }
  ' FS='[/\\\\]' setup/{essential,$vas}.lst
  esn=$?
  if [ $esn = 0 ]
  then
    echo Removing $vas
    if [ -e preremove/"$vas".sh ]
    then
      preremove/"$vas".sh
      rm preremove/"$vas".sh
    fi
    mapfile mft < setup/"$vas".lst
    for emt in ${mft[*]}
    do
      [ -f /$emt ] && rm /$emt
    done
    for emt in ${mft[*]}
    do
      [ -d /$emt ] && rmdir --i /$emt
    done
    rm -f setup/"$vas".lst.gz postinstall/"$vas".sh.done
    awk -i inplace '$1 != vas' vas="$vas" setup/installed.db
    echo Package $vas removed
  fi
  rm setup/"$vas".lst
  if [ $esn = 1 ]
  then
    warn apt-cyg cannot remove package $vas, exiting
    return 1
  fi

  done
}

function begin {
  local acv dps emt esn mft pks qry rqs sbq tsk vas arch cache mirror mirrordir
  if [ -p /dev/stdin ]
  then
    mapfile -t pks
  fi

  # process options
  while let $#
  do
  case "$1" in

    --mirror)
      set-mirror "$2"
      return
    ;;

    --cache)
      set-cache "$2"
      return
    ;;

    --deps)
      dps=1
      shift
    ;;

    --version)
      printf %s "${version[@]}"
      return
    ;;

    list | remove | update  | install  | download | listfiles |\
    show | search | depends | category | rdepends | searchall )
      tsk=$1
      shift
    ;;

    *)
      pks+=("$1")
      shift
    ;;

  esac
  done
  if type -t apt-$tsk | grep -q function
  then
    readonly arch=${HOSTTYPE/i6/x}
    apt-$tsk
  else
    printf %s "${usage[@]}"
  fi
}

function charlie {
  cd /etc/setup
  compgen -v > $1.lst
  if [ $1 = fin ]
  then
    comm -3 {start,fin}.lst
  fi
}

charlie start
begin "$@"
charlie fin
