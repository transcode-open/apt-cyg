apt-cyg
=======

apt-cyg is a Cygwin package manager. It includes a command-line installer for
Cygwin which cooperates with Cygwin Setup and uses the same repository.

<table>
<thead>
<tr>
<th>Command</th>
<th>Description</th>
<th>Analog</th>
</tr>
</thead>
<tbody>
<tr>
<td>install</td>
<td>Install packages</td>
<td>apt-get install</td>
</tr>
<tr>
<td>remove</td>
<td>Remove packages</td>
<td>apt-get&nbsp;remove</td>
</tr>
<tr>
<td>update</td>
<td>Update setup.ini</td>
<td>apt-get&nbsp;update</td>
</tr>
<tr>
<td>show</td>
<td>Displays the package records for the named packages</td>
<td>apt-cache&nbsp;show</td>
</tr>
<tr>
<td>list</td>
<td>
List packages matching given pattern. If no pattern is given, list all installed
packages.
</td>
<td>dpkg --list</td>
</tr>
<tr>
<td>search</td>
<td>Search for a filename from installed packages</td>
<td>dpkg --search</td>
</tr>
<tr>
<td>download</td>
<td>Download only - do NOT install or unpack archives</td>
<td>apt-get&nbsp;install&nbsp;--download-only</td>
</tr>
<tr>
<td>depends</td>
<td> Performs recursive dependency listings</td>
<td>apt-cache depends</td>
</tr>
</tbody>
</table>

Quick start
-----------

apt-cyg is a simple script. To install:

    wget rawgit.com/transcode-open/apt-cyg/master/apt-cyg
    install apt-cyg /bin

Example use of apt-cyg:

    apt-cyg install nano
