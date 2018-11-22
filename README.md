apt-cyg
=======

apt-cyg is a command-line installer for [Cygwin](http://cygwin.com/) which cooperates with Cygwin Setup and uses the same repository. The syntax is similar to apt-get.

Usage
-----

### Command line

    apt-cyg [OPTIONS ...] [SUBCOMMAND [PARAMETERS ...]]

### Sub commands

| Sub command | Description |
|:------- |:----------- |
| `install PACKAGE_NAMES` |  to install packages |
| `remove PACKAGE_NAMES` |  to remove packages |
| `update` |  to update *setup.ini* |
| `show` |  to show installed packages |
| `find PATTERNS` |  to find packages matching patterns |
| `describe PATTERNS` |  to describe packages matching patterns |
| `packageof COMMAND_OR_FILE_NAMES ...` |  to locate parent packages |
| `pathof "cache"\|"mirror"\|"mirrordir"\|"cache/mirrordir"\|"setup.ini"` |  to show path |
| `key-add FILES ...` |  to add keys contained in FILES |
| `key-del KEYIDS ...` |  to remove keys `KEYIDS` |
| `key-list` |  to list keys |
| `key-finger` |  to list fingerprints |
| `upgrade-self` |  to upgrade apt-cyg |
| `depends PACKAGE_NAMES ...` |  to show forward dependency information for packages with depth. |
| `rdepends PACKAGE_NAMES ...` |  to show reverse dependency information for packages with depth. |
| `completion-install` |  to install completion. |
| `completion-uninstall` |  to uninstall completion. |
| `mirrors-list` |  to show list of mirrors. |
| `benchmark-mirrors URLs ...` |  to benchmark mirrors. |
| `benchmark-parallel-mirrors URLs ...` |  to benchmark mirrors in parallel. |
| `benchmark-parallel-mirrors-list` |  to benchmark mirrors-list in parallel. |
| `scriptinfo` |  to show script infomations. |
| `show-packages-busyness PACKAGE_NAMES ...` |  to show if packages are busy or not. |
| `dist-upgrade` |  to upgrade all packages that is installed. This subcommand uses **`setup-*.exe`** |
| `update-setup` |  to update setup.exe |
| `setup [PARAMS ...]` |  to call setup.exe |
| `packages-total-count` |  count number of total packages from setup.ini |
| `packages-total-size [PATTERN_OF_SECTION]` |  count size of total packages from setup.ini |
| `packages-cached-count` |  count number of cached packages in cache/mirrordir. |
| `packages-cached-size` |  count size of cached packages in cache/mirrordir. |
| `repair-acl` |  repair the windows ACL (Access Control List). |
| `source PACKAGE_NAMES ...` |  download source archive. |
| `download PACKAGE_NAMES ...` |  download the binary package into the current directory. |
| `mirror PACKAGE_NAMES ...` |  download the binary package into the current cache/mirrordir as mirror. |
| `listfiles PACKAGE_NAMES ...` |  List files 'owned' by package(s). |
| `get-proxy` |  Get proxies for eval. |
| `ls-categories` |  List categories. |

### Options

| Option | Description |
|:------ |:----------- |
| `--ag` | use the silver searcher (currently work only at packageof subcommand) |
| `--charch ARCH` | change archetecture |
| `--use-setuprc` | set cache and mirror with /etc/setup/setup.rc |
| `--use-own-conf` | use own cache and mirror settings when apt-cyg calls setup.exe |
| `--ignore-case`, `-i` | ignore case distinctions for `<patterns>` |
| `--force-remove` | force remove
| `--force-fetch-trustedkeys` | force fetch trustedkeys |
| `--force-update-packageof-cache` | force update packageof cache |
| `--no-verify`, `-X` | Don't verify setup.ini signatures |
| `--no-check-certificate` | Don't validate the server's certificate |
| `--no-update-setup` | Don't update setup.exe
| `--no-header` | Don't print header |
| `--proxy`, `-p "auto"\|"inherit"\|"none"\|URL` | set proxy (default: ${APT_CYG_PROXY:-auto}) |
| `--completion-get-subcommand` | get subcommand (for completion internal use) |
| `--completion-disable-autoupdate` | disable completion autoupdate |
| `--max-jobs`, `-j N` | Run N jobs in parallel |
| `--mirror`, `-m URL` | set mirror
| `--cache`, `-c DIR` | set cache |
| `--file`, `-f FILE` | read package names from file |
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

`wget`, `ca-certificates`, `gnupg`

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

### True multi-architecture support

Let think a case that you want to install the x86 package when you are working under the x86_64 environment.
For example:

    apt-cyg --charch x86 install chere

As of 2013-10-26, chere package is provided for only the repository for x86.

Remarks:
Of course, you must install both environments of x86_64 and x86, beforehand.

### Signature check and key management by GnuPG

The default action of apt-cyg has been changed to check signature for 'setup.ini'.
Of course you can also avoid signature check by using `--no-verify` or `-X` options.
Public keys of cygwin and cygwinports are already registered to trusted keys of embeded.
If you want to use some other public keys, please use `key-*` subcommands.

### Upgrade apt-cyg

If apt-cyg is under the git version control, this fork can upgrade itself by `upgrade-self` subcommand.
Therefore, the most recommended way to deploy this fork is below:

    git clone HTTPS_clone_URL
    ln -s "$(realpath apt-cyg/apt-cyg)" /usr/local/bin/

`HTTPS_clone_URL` is like a `https://github.com/USERNAME/apt-cyg.git`.
It can be found from the right side menu in each fork pages on github.
    
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


---

### Contributing

This project has been re-published on GitHub to make contributing easier.  
Feel free to fork and modify this script.

The [Google Code project](https://code.google.com/p/apt-cyg/) has a list of open issues.

#### Forks on the github

See [known_forks.md](known_forks.md)

#### Todo

- [ ] Support multi mirrors: Cygwin setup can use multi mirrors. They are recorded at last-mirror section in '/etc/setup/setup.rc'. It's useful for using [Cygwinports](http://cygwinports.org/).
- [ ] Support upgrade: But maybe, busy resources can not be upgraded, and rebase problem will happen. Cygwin setup resolves by replacing them at next reboot.
- [ ] Support dependency check for remove subcommand.

#### Known Problems

For older known problems see: [**`known problems`**](known_problems.md)
