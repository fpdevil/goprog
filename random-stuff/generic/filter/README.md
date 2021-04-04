# Filter
---

`Filter` program accepts some command line arguments like `minimum` and `maximum` file sizes and
`acceptable` file suffixes etc., and a *list of files*, and outputs those files from the list that
match the command-line provided criteria.

```shell
⇒  go run filter.go -min 1000 -max 100000 -suffixes .log -directories /var/log
/var/log/fsck_apfs.log
/var/log/fsck_hfs.log
/var/log/system.log
/var/log/wifi.log

⇒  go run filter.go -min 1000 -max 100000 -suffixes .log,.gz -directories /var/log
/var/log/fsck_apfs.log
/var/log/fsck_hfs.log
/var/log/system.log
/var/log/system.log.0.gz
/var/log/system.log.1.gz
/var/log/system.log.2.gz
/var/log/system.log.3.gz
/var/log/system.log.4.gz
/var/log/system.log.5.gz
/var/log/system.log.6.gz
/var/log/system.log.7.gz
/var/log/wifi.log
```
