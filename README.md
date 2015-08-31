## Usage

Specify values manually.

```bash
echo -e "value 1\nvalue 2" | kubernetes-secret -n test-secret key-1 key-2 | kubectl create -f -
```

Or provide input as a CSV (or specify a different delimeter).

```bash
echo -e "key-1,value 1\nkey-2,value 2\n" | kubernetes-secret -n test-secret -e -d ',' | kubectl create -f -
```
