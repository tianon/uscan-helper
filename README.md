```
version=3
opts=\
dversionmangle=s/[+~](debian|dfsg|ds|deb)\d*$//,\
uversionmangle=s/(\d)[_\.\-\+]?((RC|rc|pre|dev|beta|alpha)\d*)$/$1~$2/,\
filenamemangle=s/.+\/v(\d\S*)\.tar\.gz/docker.io_$1.orig.tar.gz/ \
  https://path.to.uscan-helper.service/github.com/docker/docker .*/v(\d\S*)\.tar\.gz debian ./debian/repack.sh
```
