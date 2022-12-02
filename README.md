apt-cyg
=======

apt-cyg is a command-line installer for [Cygwin](http://cygwin.com/) which cooperates with Cygwin Setup and uses the same repository. The syntax is similar to apt-get.

Usage
-----

### Command line

    apt-cyg [<options>] [<subcommand> [<parameters> ...]]

### Subcommands

| Subcommand | Description |
|:------- |:----------- |
| `install <package names>` |  to install packages |
| `remove <package names>` |  to remove packages |
| `update` |  to update *setup.ini* |
| `show` |  to show installed packages |
| `find <patterns>` |  to find packages matching patterns |
| `describe <patterns>` |  to describe packages matching patterns |
| `packageof <command or file names> ...` |  to locate parent packages |
| `pathof {cache\|mirror\|mirrordir\|cache/mirrordir\|setup.ini}` |  to show path |
| `key-add <files> ...` |  to add keys contained in \<files\> |
| `key-del <keyids> ...` |  to remove keys \<keyids\> |
| `key-list` |  to list keys |
| `key-finger` |  to list fingerprints |
| `upgrade-self` |  to upgrade apt-cyg |
| `depends <package names> ...` |  to show forward dependency information for packages with depth. |
| `rdepends <package names> ...` |  to show reverse dependency information for packages with depth. |
| `completion-install` |  to install completion. |
| `completion-uninstall` |  to uninstall completion. |
| `mirrors-list` |  to show list of mirrors. |
| `mirrors-list-long` |  to show list of mirros with full details. |
| `mirrors-list-online` |  to show list of mirrors from online. |
| `benchmark-mirrors <urls> ...` |  to benchmark mirrors. |
| `benchmark-parallel-mirrors <urls> ...` |  to benchmark mirrors in parallel. |
| `benchmark-parallel-mirrors-list` |  to benchmark mirrors-list in parallel. |
| `scriptinfo` |  to show script information. |
| `show-packages-busyness <package names> ...` |  to show if packages are busy or not. |
| `dist-upgrade` |  to upgrade all packages that is installed. This subcommand uses **`setup-*.exe`** |
| `update-setup` |  to update setup.exe |
| `setup [<params> ...]` |  to call setup.exe |
| `packages-total-count` |  count number of total packages from setup.ini |
| `packages-total-size [<pattern of section>]` |  count size of total packages from setup.ini |
| `packages-cached-count` |  count number of cached packages in cache/mirrordir. |
| `packages-cached-size` |  count size of cached packages in cache/mirrordir. |
| `repair-acl` |  to repair the windows ACL (Access Control List). |
| `repair-postinstall` | Repair postinstall scripts. |
| `source <package names> ...` |  download source archive. |
| `mirror-source <package names> ...` | download the source package into the current cache/mirrordir as mirror. |
| `download <package names> ...` |  download the binary package into the current directory. |
| `mirror <package names> ...` |  download the binary package into the current cache/mirrordir as mirror. |
| `browse-homepage-with-mirror-source [<package names> ...]` | Browse homepages of packages with mirror-source. |
| `browse-homepage [<package names> ...]` | Browse homepages of packages. |
| `browse-summary [<package names> ...]` | Browse summaries of packages. |
| `listfiles <package names> ...` |  List files 'owned' by package(s). |
| `get-proxy` |  Get proxies for eval. |
| `ls-categories` |  List categories. |
| `ls-pkg-with-category` |  List packages with category. |
| `category <category>` | List all packages in given \<category\>.|
| `setuprc-get <section>` | Get section from **`setup.rc`**. |
| `set-cache [<cache>]` | Set cache. |
| `set-mirror [<mirrors> ...]` | Set mirrors. Note: `setup-x86{,_64}.exe` uses all of them but currently `apt-cyg` uses the first one only. |
| `mark-auto [<packages> ...]` | Mark the given packages as automatically installed. |
| `mark-manual [<packages> ...]` | Mark the given packages as manually installed. |
| `mark-showauto` |Print the list of automatically installed packages. |
| `mark-showmanual` | Print the list of manually installed packages. |
| `call [<internal_function> [<args> ...]]` | Call internal function in apt-cyg. |
| `time [<internal_function> [<args> ...]]` | Report time consumed to call internal function in apt-cyg. |
| `filelist [<package>]` | File list like apt-file list. |
| `filesearch [<pattern>]` | File search like apt-file search. |

### Options

| Option | Description |
|:------ |:----------- |
| `--ag` | use the silver searcher (currently work only at packageof subcommand) |
| `--ignore-case`, `-i` | ignore case distinctions for `<patterns>` |
| `--force-remove` | force remove
| `--force-fetch-trustedkeys` | force fetch trustedkeys |
| `--force-update-packageof-cache` | force update packageof cache |
| `--no-verify`, `-X` | Don't verify setup.ini signatures |
| `--no-check-certificate` | Don't validate the server's certificate |
| `--no-update-setup` | Don't update setup.exe
| `--no-header` | Don't print header |
| `--proxy`, `-p {auto\|inherit\|none\|<url>}` | set proxy (default: ${APT_CYG_PROXY:-auto}) |
| `--completion-get-subcommand` | get subcommand (for completion internal use) |
| `--completion-disable-autoupdate` | disable completion autoupdate |
| `--max-jobs`, `-j <n>` | Run \<n\> jobs in parallel |
| `--mirror`, `-m <url>` | set mirror
| `--cache`, `-c <dir>` | set cache |
| `--file`, `-f <file>` | read package names from \<file\> |
| `--noupdate`, `-u` | don't update setup.ini from mirror |
| `--ipv4`, `-4` | wget prefer ipv4 |
| `--no-progress` | hide the progress bar in any verbosity mode |
| `--quiet`, `-q` | quiet (no output) |
| `--verbose`, `-v` | verbose |
| `--help` | Display help and exit |
| `--version` | Display version and exit |

Requirements
------------

`apt-cyg` requires the Cygwin default environment and the additional *Cygwin* packages:

`wget`, `ca-certificates`, `gnupg`, `libiconv`

In **32bit** version of cygwin, `wget` requires an additional setting for the `ca-certificates` package.
Choose one of below settings.

    # 1. Create symbolic link for the default ca-directory of wget. 
    ln -s /usr/ssl /etc/
    
    # or
    # 2. Set ca-directory paramete in '/etc/wgetrc'. 
    echo "ca-directory = /usr/ssl/certs" >> /etc/wgetrc
    
    # or
    # 3. Set ca-directory paramete in '~/.wgetrc'. 
    echo "ca-directory = /usr/ssl/certs" >> ~/.wgetrc

Remarks:
Above additional settings for wget is not required for 64bit version of cygwin.
But, as of 2014-01-17, perhaps `ca-certificates` package makes fail of certification in 64bit version of cygwin with Windows 8. See below:

* Known Problem / [2014-01-17: ca-certificates package is not setup correct at x86_64 with Windows 8.](#2014-01-17-ca-certificates-package-is-not-setup-correct-at-x86_64-with-windows-8)

Quick start
-----------

The most recommended way to deploy this fork can be seen from a link below:
* New features / [Upgrade apt-cyg](#upgrade-apt-cyg)

apt-cyg is a simple script. Once you have a copy, make it executable:

    chmod +x /bin/apt-cyg

Optionally place apt-cyg in a bin/ folder on your path.

Then use apt-cyg, for example:

    apt-cyg install nano

New features
------------

### dist-upgrade support

This fork has achieved `dist-upgrade` command by using `setup-x86.exe` and `setup-x86_64.exe` as a backend.
Note that all of running tasks on the cygwin will be killed before starting dist-upgrade.

### Multiple hash algorithms support

After the middle of 2015-03, the cygwin project changed the hash algorithm for checking tarball from md5 to sha512.
But, as of 2015-04-09, the cygwinports project seems still using md5.
This fork is available for both of cygwin and cygwinports by supporting algorithm of md5, sha1, sha224, sha256 and sha512.

### Signature check and key management by GnuPG

The default action of apt-cyg has been changed to check signature for 'setup.ini'.
Of course you can also avoid signature check by using `--no-verify` or `-X` options.
Public keys of cygwin and cygwinports are already registered to trusted keys of embeded.
If you want to use some other public keys, please use `key-*` subcommands.

### Upgrade apt-cyg

If apt-cyg is under Git version control, this fork can upgrade itself by `upgrade-self` subcommand.
Therefore, the most recommended way to deploy this fork is `copy and paste` below commands to cygwin console:

    # cd $DIR # Change working directory where you want to install apt-cyg
    git clone https://github.com/kou1okada/apt-cyg.git
    ln -s "$(realpath apt-cyg/apt-cyg)" /usr/local/bin/

If you want to use another fork, which forked from https://github.com/kou1okada/apt-cyg, rewrite the URL for apropriate one.

### Proxy support

Use `--proxy`, `-p` option.
This option must take a parameter from one of "auto", "inherit", "none" and URL.

* "auto" will determine a proxy using a part of the [Web Proxy Auto-Discovery Protocol (WPAD)](http://en.wikipedia.org/wiki/Web_Proxy_Autodiscovery_Protocol).
The current implementation will look for a string of "PROXY URL" from "http://wpad/wpad.dat".
If "wpad.dat" could not be downloaded, the proxy settings are inherited from the parent environment.
* "inherit" will inherit the proxy settings from the parent environment.
* "none" will not use the proxy.
* URL can take a string like "protocol://hostname:port".

For example:

    apt-cyg --proxy http://proxy.home:8080 update

The default parameter is "${APT\_CYG\_PROXY:-auto}".
At the environment where is not provided the WPAD server, it makes the lag for a few seconds at startup.
So, if you don't want to use WPAD, please define APT\_CYG\_PROXY environment variable as below:

    export APT_CYG_PROXY=inherit

### Bash completion support

Bash completion script can be installed  to "/etc/bash_completion.d/apt-cyg" by `completion-install` subcommand.
It will be automatically updated when apt-cyg is upgraded to newer version.
If you don't want to update it automatically, execute `completion-install` subcommand in conjunction with `--completion-disable-autoupdate` option. And `completion-uninstall` subcommand removes "/etc/bash_completion.d/apt-cyg".

Some other forks, [Milly / apt-cyg](https://github.com/Milly/apt-cyg) under the cfg / apt-cyg fork, [ashumkin / apt-cyg](https://github.com/ashumkin/apt-cyg) and etc, are also supported it.

Contributing
------------

This project has been re-published on GitHub to make contributing easier. Feel free to fork and modify this script.

The [Google Code project](https://code.google.com/p/apt-cyg/) has a list of open issues.

### Forks on the github

Caution:
Please do not merge forks that have incompatible licenses.

Ex.) Merging to the GPL from the MIT is possible. But merging to the MIT from the GPL  is impossible.

See [other_forks.md](other_forks.md)

Todo
------------

- [ ] Support multi mirrors: Cygwin setup can use multi mirrors. They are recorded at last-mirror section in '/etc/setup/setup.rc'. It's useful for using [Cygwinports](http://cygwinports.org/).
- [ ] Support upgrade: But maybe, busy resources can not be upgraded, and rebase problem will happen. Cygwin setup resolves by replacing them at next reboot.
- [ ] Support dependency check for remove subcommand.

Known Problems
------------

For older known problems see [known_problems.md](known_problems.md).

License
-------

[![The MIT license](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE) 
