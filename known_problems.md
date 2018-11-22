
### Known Problems

---

### 2016-02-12: Problems about the ACL

In recent version, the cygwin changed the ACL mechanism .
So, in the cygwin current version, if it will be installed by `setup.exe` with `-B` or `--no-admin` option,
the cygwin root (/) does not have correct ACL.

A new subcommand `repair-acl` tries to repair it.
But some package, that are failed to install by the ACL problem, need to be reinstalled.



### 2015-04-09: gpgv seems not work correctly on 32 bit environment

To solve this problem, the backend of `verify_signatures` function was changed from `gpgv` to `gpg`.
thienhv reported this problem [#14](https://github.com/kou1okada/apt-cyg/issues/14). Thanks.

### 2014-01-17: ca-certificates package is not setup correct at x86_64 with Windows 8.

After clean installing with `setup-x86_64.exe`, there are something wrong about ca-certificate package as below:

    $ ls -lnG \
         /cygdrive/c/cygwin/etc/pki/ca-trust/extracted/openssl/ca-bundle.trust.crt \
         /cygdrive/c/cygwin64/etc/pki/ca-trust/extracted/openssl/ca-bundle.trust.crt
    -r--r--r-- 1 1001 314336 Jan 17 18:17 /cygdrive/c/cygwin/etc/pki/ca-trust/extracted/openssl/ca-bundle.trust.crt
    -rw-r--r-- 1 1001      0 Oct 16 12:35 /cygdrive/c/cygwin64/etc/pki/ca-trust/extracted/openssl/ca-bundle.trust.crt
    $ ls -lnG \
         /cygdrive/c/cygwin/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem \
         /cygdrive/c/cygwin64/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem
    -r--r--r-- 1 1001 232342 Jan 17 18:17 /cygdrive/c/cygwin/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem
    -rw-r--r-- 1 1001      0 Oct 16 12:35 /cygdrive/c/cygwin64/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem

The x86 environment seems no problem except that need an additional setting for wget and ca-certificate. Hmm,,,    

It seems that `/usr/bin/update-ca-trust` is failed at x86_64.

It's ad hoc, but effective way to fix it, is to copy files from cygwin32 to cygwin64, as below:

    cp -a /cygdrive/c/cygwin/etc/pki/ca-trust/extracted/openssl/ca-bundle.trust.crt \
          /cygdrive/c/cygwin64/etc/pki/ca-trust/extracted/openssl/ca-bundle.trust.crt
    cp -a /cygdrive/c/cygwin/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem \
          /cygdrive/c/cygwin64/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem

This problem is fixed at 2014-01-24.
version 0.18.7-1 of p11-kit, p11-kit-trust and libp11-kit0 did not work at the cygwin64 under the Windows 8.
If you face the above problem, please upgrade these three packages from version 0.18.7-1 to version 0.18.7-2.

For more details, see a thread of below:

* Cygwin mailing list / cygwin / [Re: Is there someone who have a same problem ?](http://cygwin.com/ml/cygwin/2014-01/msg00368.html)

### 2014-01-15: Signature check failed at cygwinports x86_64.

Oops, setup.bz2 is newer than setup.bz2.sig.
And it seems to be corrupted.

    $ date
    Wed Jan 15 01:30:19 JST 2014
    $ w3m -dump ftp://ftp.cygwinports.org/pub/cygwinports/x86_64/
    Index of ftp://ftp.cygwinports.org/pub/cygwinports/x86_64/
    
    [Upper Directory]
    md5.sum. . . . . . . Jan 10 14:24    184
    release/ . . . . . . Jan 10 18:00
    setup.bz2. . . . . . Jan 10 13:59   579K
    setup.bz2.sig. . . . Dec  6 11:28     72
    setup.ini. . . . . . Dec  6 11:28  3.20M
    setup.ini.sig. . . . Dec  6 11:28     72
    
    $ wget -q -O - ftp://ftp.cygwinports.org/pub/cygwinports/x86_64/setup.bz2 | bzip2 -tvv
      (stdin):
        [1: huff+mtf rt+rld]
        [2: huff+mtf data integrity (CRC) error in data
    
    You can use the `bzip2recover' program to attempt to recover
    data from undamaged sections of corrupted files.

As of 2014-01-24, above problem is recovered at least.

### FIXED: Check setup (`cygcheck -c`) can not detect .tar.xz packages

At cygwin 1.7.25, cygcheck is hardcoded for .tar.gz and .tar.bz2.
So check setup (`cygcheck -c`) can not detect .tar.xz packages.
This bug was already fixed with [src/winsup/utils/dump_setup.cc#rev1.28](http://cygwin.com/cgi-bin/cvsweb.cgi/src/winsup/utils/dump_setup.cc?cvsroot=src#rev1.28).
Please wait a release of cygwin 1.7.26.

This Problem was fixed [cygwin 1.7.26](http://cygwin.com/ml/cygwin-announce/2013-11/msg00027.html) which released at 2013-11-29.
