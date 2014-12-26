apt-cyg
=======

apt-cyg is a Cygwin package manager. It includes a command-line installer for
Cygwin which cooperates with Cygwin Setup and uses the same repository.

| Command | Description | Analog |
| --------| ----------- | ------ |
| install | Install packages | apt-get install |
| remove  | Remove packages  | apt-get remove |
| update  | Update setup.ini | apt-get update |
| show    | Displays the package records for the named packages | apt-cache show |
| list    | List packages matching given pattern. If no pattern is given, list all installed packages. | dpkg --list |
| search  | Search for a filename from installed packages | dpkg --search |
| download | Download only - do NOT install or unpack archives | apt-get&nbsp;install&nbsp;--download-only |
| depends | Performs recursive dependency listings | apt-cache depends|
| listfiles | List files owned by packages | dpkg --listfiles |

Quick start
-----------

apt-cyg is a simple script. To install:

```bash
lynx -source https://github.com/transcode-open/apt-cyg/raw/master/apt-cyg > apt-cyg
install apt-cyg /bin
```

Example use of apt-cyg:

```bash
apt-cyg install nano
```
