## Usage

Install go.

The following Assumes the program is available as `kubernetes-secret`.

You can also run the program directly, replacing `kubernetes-secret` with `go run ./kubernetes-secret.go`.

### Pipe secret values via stdin

```bash
echo -e "value 1\nvalue 2" | kubernetes-secret -n test-secret key-1 key-2 | kubectl create -f -
```

### Provide secrets as comma separated values (or specify a different delimeter)

```bash
echo -e "key-1,value 1\nkey-2,value 2\n" | kubernetes-secret -n test-secret -e -d ',' | kubectl create -f -
```
