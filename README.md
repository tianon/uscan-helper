```
version=3
opts=\
dversionmangle=s/[+~](debian|dfsg|ds|deb)\d*$//,\
uversionmangle=s/(\d)[_\.\-\+]?((RC|rc|pre|dev|beta|alpha)\d*)$/$1~$2/,\
filenamemangle=s/.+\/v(\d\S*)\.tar\.gz/docker.io_$1.orig.tar.gz/ \
  https://uscan.tianon.xyz/github.com/docker/docker https://github.com/.*/v(\d\S*)\.tar\.gz debian ./debian/repack.sh

# /tags paginates too soon, so we lose historical versions too early for tianon's tastes
#  https://github.com/docker/docker/tags .*/v(\d\S*)\.tar\.gz debian ./debian/repack.sh
```

```console
$ uscan --verbose --report
uscan info: uscan (version 2.16.6) See uscan(1) for help
uscan info: Scan watch files in .
uscan info: Check debian/watch and debian/changelog in ./.git/refs/tags
uscan info: Check debian/watch and debian/changelog in ./.git/refs/remotes/origin
uscan info: Check debian/watch and debian/changelog in ./.git/logs/refs/remotes/origin
uscan info: Check debian/watch and debian/changelog in .
uscan info: package="docker.io" version="1.11.2~ds1-7" (as seen in debian/changelog)
uscan info: package="docker.io" version="1.11.2~ds1" (no epoch/revision)
uscan info: ./debian/changelog sets package="docker.io" version="1.11.2~ds1"
uscan info: Process ./debian/watch (package=docker.io version=1.11.2~ds1)
uscan info: opts: dversionmangle=s/[+~](debian|dfsg|ds|deb)\d*$//,uversionmangle=s/(\d)[_\.\-\+]?((RC|rc|pre|dev|beta|alpha)\d*)$/$1~$2/,filenamemangle=s/.+\/v(\d\S*)\.tar\.gz/docker.io_$1.orig.tar.gz/
uscan info: line: https://uscan.tianon.xyz/github.com/docker/docker https://github.com/.*/v(\d\S*)\.tar\.gz debian ./debian/repack.sh
uscan info: Parsing dversionmangle=s/[+~](debian|dfsg|ds|deb)\d*$//
uscan info: Parsing uversionmangle=s/(\d)[_\.\-\+]?((RC|rc|pre|dev|beta|alpha)\d*)$/$1~$2/
uscan info: Parsing filenamemangle=s/.+\/v(\d\S*)\.tar\.gz/docker.io_$1.orig.tar.gz/
uscan info: line: https://uscan.tianon.xyz/github.com/docker/docker https://github.com/.*/v(\d\S*)\.tar\.gz debian ./debian/repack.sh
uscan info: Last orig.tar.* tarball version (from debian/changelog): 1.11.2~ds1
uscan info: Last orig.tar.* tarball version (dversionmangled): 1.11.2
uscan info: Requesting URL:
   https://uscan.tianon.xyz/github.com/docker/docker
uscan info: Matching pattern:
   (?:(?:https://uscan.tianon.xyz)?\/github\.com\/docker\/docker)?https://github.com/.*/v(\d\S*)\.tar\.gz
uscan info: Found the following matching hrefs on the web page (newest first):
   https://github.com/docker/docker/archive/v1.12.1.tar.gz (1.12.1) index=1.12.1-1 
   https://github.com/docker/docker/archive/v1.12.1-rc2.tar.gz (1.12.1~rc2) index=1.12.1~rc2-1 
   https://github.com/docker/docker/archive/v1.12.1-rc1.tar.gz (1.12.1~rc1) index=1.12.1~rc1-1 
   https://github.com/docker/docker/archive/v1.12.0.tar.gz (1.12.0) index=1.12.0-1 
   https://github.com/docker/docker/archive/v1.12.0-rc5.tar.gz (1.12.0~rc5) index=1.12.0~rc5-1 
   https://github.com/docker/docker/archive/v1.12.0-rc4.tar.gz (1.12.0~rc4) index=1.12.0~rc4-1 
   https://github.com/docker/docker/archive/v1.12.0-rc3.tar.gz (1.12.0~rc3) index=1.12.0~rc3-1 
   https://github.com/docker/docker/archive/v1.12.0-rc2.tar.gz (1.12.0~rc2) index=1.12.0~rc2-1 
   https://github.com/docker/docker/archive/v1.12.0-rc1.tar.gz (1.12.0~rc1) index=1.12.0~rc1-1 
   https://github.com/docker/docker/archive/v1.11.2.tar.gz (1.11.2) index=1.11.2-1 
   https://github.com/docker/docker/archive/v1.11.2-rc1.tar.gz (1.11.2~rc1) index=1.11.2~rc1-1 
   https://github.com/docker/docker/archive/v1.11.1.tar.gz (1.11.1) index=1.11.1-1 
   https://github.com/docker/docker/archive/v1.11.1-rc1.tar.gz (1.11.1~rc1) index=1.11.1~rc1-1 
   https://github.com/docker/docker/archive/v1.11.0.tar.gz (1.11.0) index=1.11.0-1 
   https://github.com/docker/docker/archive/v1.11.0-rc5.tar.gz (1.11.0~rc5) index=1.11.0~rc5-1 
   https://github.com/docker/docker/archive/v1.11.0-rc4.tar.gz (1.11.0~rc4) index=1.11.0~rc4-1 
   https://github.com/docker/docker/archive/v1.11.0-rc3.tar.gz (1.11.0~rc3) index=1.11.0~rc3-1 
   https://github.com/docker/docker/archive/v1.11.0-rc2.tar.gz (1.11.0~rc2) index=1.11.0~rc2-1 
   https://github.com/docker/docker/archive/v1.11.0-rc1.tar.gz (1.11.0~rc1) index=1.11.0~rc1-1 
   https://github.com/docker/docker/archive/v1.10.3.tar.gz (1.10.3) index=1.10.3-1 
   https://github.com/docker/docker/archive/v1.10.3-rc2.tar.gz (1.10.3~rc2) index=1.10.3~rc2-1 
   https://github.com/docker/docker/archive/v1.10.3-rc1.tar.gz (1.10.3~rc1) index=1.10.3~rc1-1 
   https://github.com/docker/docker/archive/v1.10.2.tar.gz (1.10.2) index=1.10.2-1 
   https://github.com/docker/docker/archive/v1.10.2-rc1.tar.gz (1.10.2~rc1) index=1.10.2~rc1-1 
   https://github.com/docker/docker/archive/v1.10.1.tar.gz (1.10.1) index=1.10.1-1 
   https://github.com/docker/docker/archive/v1.10.1-rc1.tar.gz (1.10.1~rc1) index=1.10.1~rc1-1 
   https://github.com/docker/docker/archive/v1.10.0.tar.gz (1.10.0) index=1.10.0-1 
   https://github.com/docker/docker/archive/v1.10.0-rc4.tar.gz (1.10.0~rc4) index=1.10.0~rc4-1 
   https://github.com/docker/docker/archive/v1.10.0-rc3.tar.gz (1.10.0~rc3) index=1.10.0~rc3-1 
   https://github.com/docker/docker/archive/v1.10.0-rc2.tar.gz (1.10.0~rc2) index=1.10.0~rc2-1 
   https://github.com/docker/docker/archive/v1.10.0-rc1.tar.gz (1.10.0~rc1) index=1.10.0~rc1-1 
   https://github.com/docker/docker/archive/v1.9.1.tar.gz (1.9.1) index=1.9.1-1 
   https://github.com/docker/docker/archive/v1.9.1-rc1.tar.gz (1.9.1~rc1) index=1.9.1~rc1-1 
   https://github.com/docker/docker/archive/v1.9.0.tar.gz (1.9.0) index=1.9.0-1 
   https://github.com/docker/docker/archive/v1.9.0-rc5.tar.gz (1.9.0~rc5) index=1.9.0~rc5-1 
   https://github.com/docker/docker/archive/v1.9.0-rc4.tar.gz (1.9.0~rc4) index=1.9.0~rc4-1 
   https://github.com/docker/docker/archive/v1.9.0-rc3.tar.gz (1.9.0~rc3) index=1.9.0~rc3-1 
   https://github.com/docker/docker/archive/v1.9.0-rc2.tar.gz (1.9.0~rc2) index=1.9.0~rc2-1 
   https://github.com/docker/docker/archive/v1.9.0-rc1.tar.gz (1.9.0~rc1) index=1.9.0~rc1-1 
   https://github.com/docker/docker/archive/v1.8.3.tar.gz (1.8.3) index=1.8.3-1 
   https://github.com/docker/docker/archive/v1.8.2.tar.gz (1.8.2) index=1.8.2-1 
   https://github.com/docker/docker/archive/v1.8.2-rc1.tar.gz (1.8.2~rc1) index=1.8.2~rc1-1 
   https://github.com/docker/docker/archive/v1.8.1.tar.gz (1.8.1) index=1.8.1-1 
   https://github.com/docker/docker/archive/v1.8.0.tar.gz (1.8.0) index=1.8.0-1 
   https://github.com/docker/docker/archive/v1.8.0-rc3.tar.gz (1.8.0~rc3) index=1.8.0~rc3-1 
   https://github.com/docker/docker/archive/v1.8.0-rc2.tar.gz (1.8.0~rc2) index=1.8.0~rc2-1 
   https://github.com/docker/docker/archive/v1.8.0-rc1.tar.gz (1.8.0~rc1) index=1.8.0~rc1-1 
   https://github.com/docker/docker/archive/v1.7.1.tar.gz (1.7.1) index=1.7.1-1 
   https://github.com/docker/docker/archive/v1.7.1-rc3.tar.gz (1.7.1~rc3) index=1.7.1~rc3-1 
   https://github.com/docker/docker/archive/v1.7.1-rc2.tar.gz (1.7.1~rc2) index=1.7.1~rc2-1 
   https://github.com/docker/docker/archive/v1.7.1-rc1.tar.gz (1.7.1~rc1) index=1.7.1~rc1-1 
   https://github.com/docker/docker/archive/v1.7.0.tar.gz (1.7.0) index=1.7.0-1 
   https://github.com/docker/docker/archive/v1.7.0-rc5.tar.gz (1.7.0~rc5) index=1.7.0~rc5-1 
   https://github.com/docker/docker/archive/v1.7.0-rc4.tar.gz (1.7.0~rc4) index=1.7.0~rc4-1 
   https://github.com/docker/docker/archive/v1.7.0-rc3.tar.gz (1.7.0~rc3) index=1.7.0~rc3-1 
   https://github.com/docker/docker/archive/v1.7.0-rc2.tar.gz (1.7.0~rc2) index=1.7.0~rc2-1 
   https://github.com/docker/docker/archive/v1.7.0-rc1.tar.gz (1.7.0~rc1) index=1.7.0~rc1-1 
   https://github.com/docker/docker/archive/v1.6.2.tar.gz (1.6.2) index=1.6.2-1 
   https://github.com/docker/docker/archive/v1.6.1.tar.gz (1.6.1) index=1.6.1-1 
   https://github.com/docker/docker/archive/v1.6.0.tar.gz (1.6.0) index=1.6.0-1 
   https://github.com/docker/docker/archive/v1.6.0-rc7.tar.gz (1.6.0~rc7) index=1.6.0~rc7-1 
   https://github.com/docker/docker/archive/v1.6.0-rc6.tar.gz (1.6.0~rc6) index=1.6.0~rc6-1 
   https://github.com/docker/docker/archive/v1.6.0-rc5.tar.gz (1.6.0~rc5) index=1.6.0~rc5-1 
   https://github.com/docker/docker/archive/v1.6.0-rc4.tar.gz (1.6.0~rc4) index=1.6.0~rc4-1 
   https://github.com/docker/docker/archive/v1.6.0-rc3.tar.gz (1.6.0~rc3) index=1.6.0~rc3-1 
   https://github.com/docker/docker/archive/v1.6.0-rc2.tar.gz (1.6.0~rc2) index=1.6.0~rc2-1 
   https://github.com/docker/docker/archive/v1.6.0-rc1.tar.gz (1.6.0~rc1) index=1.6.0~rc1-1 
   https://github.com/docker/docker/archive/v1.5.0.tar.gz (1.5.0) index=1.5.0-1 
   https://github.com/docker/docker/archive/v1.5.0-rc4.tar.gz (1.5.0~rc4) index=1.5.0~rc4-1 
   https://github.com/docker/docker/archive/v1.5.0-rc3.tar.gz (1.5.0~rc3) index=1.5.0~rc3-1 
   https://github.com/docker/docker/archive/v1.5.0-rc2.tar.gz (1.5.0~rc2) index=1.5.0~rc2-1 
   https://github.com/docker/docker/archive/v1.5.0-rc1.tar.gz (1.5.0~rc1) index=1.5.0~rc1-1 
   https://github.com/docker/docker/archive/v1.4.1.tar.gz (1.4.1) index=1.4.1-1 
   https://github.com/docker/docker/archive/v1.4.0.tar.gz (1.4.0) index=1.4.0-1 
   https://github.com/docker/docker/archive/v1.3.3.tar.gz (1.3.3) index=1.3.3-1 
   https://github.com/docker/docker/archive/v1.3.2.tar.gz (1.3.2) index=1.3.2-1 
   https://github.com/docker/docker/archive/v1.3.1.tar.gz (1.3.1) index=1.3.1-1 
   https://github.com/docker/docker/archive/v1.3.0.tar.gz (1.3.0) index=1.3.0-1 
   https://github.com/docker/docker/archive/v1.2.0.tar.gz (1.2.0) index=1.2.0-1 
   https://github.com/docker/docker/archive/v1.1.2.tar.gz (1.1.2) index=1.1.2-1 
   https://github.com/docker/docker/archive/v1.1.1.tar.gz (1.1.1) index=1.1.1-1 
   https://github.com/docker/docker/archive/v1.1.0.tar.gz (1.1.0) index=1.1.0-1 
   https://github.com/docker/docker/archive/v1.0.1.tar.gz (1.0.1) index=1.0.1-1 
   https://github.com/docker/docker/archive/v1.0.0.tar.gz (1.0.0) index=1.0.0-1 
   https://github.com/docker/docker/archive/v0.12.0.tar.gz (0.12.0) index=0.12.0-1 
   https://github.com/docker/docker/archive/v0.11.1.tar.gz (0.11.1) index=0.11.1-1 
   https://github.com/docker/docker/archive/v0.11.0.tar.gz (0.11.0) index=0.11.0-1 
   https://github.com/docker/docker/archive/v0.10.0.tar.gz (0.10.0) index=0.10.0-1 
   https://github.com/docker/docker/archive/v0.9.1.tar.gz (0.9.1) index=0.9.1-1 
   https://github.com/docker/docker/archive/v0.9.0.tar.gz (0.9.0) index=0.9.0-1 
   https://github.com/docker/docker/archive/v0.8.1.tar.gz (0.8.1) index=0.8.1-1 
   https://github.com/docker/docker/archive/v0.8.0.tar.gz (0.8.0) index=0.8.0-1 
   https://github.com/docker/docker/archive/v0.7.6.tar.gz (0.7.6) index=0.7.6-1 
   https://github.com/docker/docker/archive/v0.7.5.tar.gz (0.7.5) index=0.7.5-1 
   https://github.com/docker/docker/archive/v0.7.4.tar.gz (0.7.4) index=0.7.4-1 
   https://github.com/docker/docker/archive/v0.7.3.tar.gz (0.7.3) index=0.7.3-1 
   https://github.com/docker/docker/archive/v0.7.2.tar.gz (0.7.2) index=0.7.2-1 
   https://github.com/docker/docker/archive/v0.7.1.tar.gz (0.7.1) index=0.7.1-1 
   https://github.com/docker/docker/archive/v0.7.0.tar.gz (0.7.0) index=0.7.0-1 
   https://github.com/docker/docker/archive/v0.7.0-rc7.tar.gz (0.7.0~rc7) index=0.7.0~rc7-1 
   https://github.com/docker/docker/archive/v0.7.0-rc6.tar.gz (0.7.0~rc6) index=0.7.0~rc6-1 
   https://github.com/docker/docker/archive/v0.7.0-rc5.tar.gz (0.7.0~rc5) index=0.7.0~rc5-1 
   https://github.com/docker/docker/archive/v0.7.0-rc4.tar.gz (0.7.0~rc4) index=0.7.0~rc4-1 
   https://github.com/docker/docker/archive/v0.7.0-rc3.tar.gz (0.7.0~rc3) index=0.7.0~rc3-1 
   https://github.com/docker/docker/archive/v0.7.0-rc2.tar.gz (0.7.0~rc2) index=0.7.0~rc2-1 
   https://github.com/docker/docker/archive/v0.7.0-rc1.tar.gz (0.7.0~rc1) index=0.7.0~rc1-1 
   https://github.com/docker/docker/archive/v0.6.7.tar.gz (0.6.7) index=0.6.7-1 
   https://github.com/docker/docker/archive/v0.6.6.tar.gz (0.6.6) index=0.6.6-1 
   https://github.com/docker/docker/archive/v0.6.5.tar.gz (0.6.5) index=0.6.5-1 
   https://github.com/docker/docker/archive/v0.6.4.tar.gz (0.6.4) index=0.6.4-1 
   https://github.com/docker/docker/archive/v0.6.3.tar.gz (0.6.3) index=0.6.3-1 
   https://github.com/docker/docker/archive/v0.6.2.tar.gz (0.6.2) index=0.6.2-1 
   https://github.com/docker/docker/archive/v0.6.1.tar.gz (0.6.1) index=0.6.1-1 
   https://github.com/docker/docker/archive/v0.6.0.tar.gz (0.6.0) index=0.6.0-1 
   https://github.com/docker/docker/archive/v0.5.3.tar.gz (0.5.3) index=0.5.3-1 
   https://github.com/docker/docker/archive/v0.5.2.tar.gz (0.5.2) index=0.5.2-1 
   https://github.com/docker/docker/archive/v0.5.1.tar.gz (0.5.1) index=0.5.1-1 
   https://github.com/docker/docker/archive/v0.5.0.tar.gz (0.5.0) index=0.5.0-1 
   https://github.com/docker/docker/archive/v0.4.8.tar.gz (0.4.8) index=0.4.8-1 
   https://github.com/docker/docker/archive/v0.4.7.tar.gz (0.4.7) index=0.4.7-1 
   https://github.com/docker/docker/archive/v0.4.6.tar.gz (0.4.6) index=0.4.6-1 
   https://github.com/docker/docker/archive/v0.4.5.tar.gz (0.4.5) index=0.4.5-1 
   https://github.com/docker/docker/archive/v0.4.4.tar.gz (0.4.4) index=0.4.4-1 
   https://github.com/docker/docker/archive/v0.4.3.tar.gz (0.4.3) index=0.4.3-1 
   https://github.com/docker/docker/archive/v0.4.2.tar.gz (0.4.2) index=0.4.2-1 
   https://github.com/docker/docker/archive/v0.4.1.tar.gz (0.4.1) index=0.4.1-1 
   https://github.com/docker/docker/archive/v0.4.0.tar.gz (0.4.0) index=0.4.0-1 
   https://github.com/docker/docker/archive/v0.3.4.tar.gz (0.3.4) index=0.3.4-1 
   https://github.com/docker/docker/archive/v0.3.3.tar.gz (0.3.3) index=0.3.3-1 
   https://github.com/docker/docker/archive/v0.3.2.tar.gz (0.3.2) index=0.3.2-1 
   https://github.com/docker/docker/archive/v0.3.1.tar.gz (0.3.1) index=0.3.1-1 
   https://github.com/docker/docker/archive/v0.3.0.tar.gz (0.3.0) index=0.3.0-1 
   https://github.com/docker/docker/archive/v0.2.2.tar.gz (0.2.2) index=0.2.2-1 
   https://github.com/docker/docker/archive/v0.2.1.tar.gz (0.2.1) index=0.2.1-1 
   https://github.com/docker/docker/archive/v0.2.0.tar.gz (0.2.0) index=0.2.0-1 
   https://github.com/docker/docker/archive/v0.1.8.tar.gz (0.1.8) index=0.1.8-1 
   https://github.com/docker/docker/archive/v0.1.7.tar.gz (0.1.7) index=0.1.7-1 
   https://github.com/docker/docker/archive/v0.1.6.tar.gz (0.1.6) index=0.1.6-1 
   https://github.com/docker/docker/archive/v0.1.5.tar.gz (0.1.5) index=0.1.5-1 
   https://github.com/docker/docker/archive/v0.1.4.tar.gz (0.1.4) index=0.1.4-1 
   https://github.com/docker/docker/archive/v0.1.3.tar.gz (0.1.3) index=0.1.3-1 
   https://github.com/docker/docker/archive/v0.1.2.tar.gz (0.1.2) index=0.1.2-1 
   https://github.com/docker/docker/archive/v0.1.1.tar.gz (0.1.1) index=0.1.1-1 
   https://github.com/docker/docker/archive/v0.1.0.tar.gz (0.1.0) index=0.1.0-1 
uscan info: Matching target for downloadurlmangle: https://github.com/docker/docker/archive/v1.12.1.tar.gz
uscan info: Upstream URL (downloadurlmangled):
   https://github.com/docker/docker/archive/v1.12.1.tar.gz
uscan info: Newest upstream tarball version selected for download (uversionmangled): 1.12.1
uscan info: Matching target for filenamemangle: https://github.com/docker/docker/archive/v1.12.1.tar.gz
uscan info: Download filename (filenamemangled): docker.io_1.12.1.orig.tar.gz
uscan: Newest version of docker.io on remote site is 1.12.1, local version is 1.11.2~ds1
 (mangled local version is 1.11.2)
uscan:    => Newer package available from
      https://github.com/docker/docker/archive/v1.12.1.tar.gz
uscan info: Scan finished
```
