## kubernetes-secret

Tool to help you create and update secrets in kubernetes

### Pipe secret values via stdin
```bash
cat .env.production | kubernetes-secret -n test -ns default | kubectl create -f -
```

You can also load file contents by specifying a special prefix `PEM=>>>file.pem.`. You can define your own prefix by passing `-f ###`

### Provide secrets as comma separated values (or specify a different delimeter)
```bash
echo -e "key-1,value 1\nkey-2,value 2\n" | kubernetes-secret -n test -ns test -d ',' | kubectl create -f -
```

### Releases
You can download binary releases for linux, macos and windows [here](https://github.com/checkr/kubernetes-secret/releases)
