# Note from the developer

I am no longer working on apt-cyg. As far as I know, it works.

However it was written quite a while ago by reverse engineering the files used by Cygwin's setup.exe program. It doesn't
share any code with setup.exe, and if setup.exe is modified in such a way that assumptions made by apt-cyg no longer
apply, apt-cyg will break.

apt-cyg is an infrastructure program and probably belongs in Cygwin rather than here, but as far as I know, RedHat is
no longer paying more than one or two engineers to work on Cygwin and they are unlikely to do anything with apt-cyg or
other, similar programs that are out there.

Because I got dragged into a time-consuming and unpleasant law suit involving apt-cyg, I am very unlikely to put
more time and effort into it anytime soon, even though I won. I write open source software because I enjoy it, not so
I can get involved in meaningless conflicts.

If you would like to contribute to apt-cyg, the best way to do that is to fork it. You don't need my permission to
do that, just follow the license.

--Steve Jungels
