# The problem

A lot of password managers store credit card number as contiguous strings. And there are quite a lot of banks that provide secure online payment, but their forms are not autofillable. These cases you either ctr+c, ctrl+v your credit card number to a text editor, then ctrl+x. ctrl+v your number in chunks of four to these fields, or you can also use `sed -e 's/\(....\)/\1 /g' <<< <your contiguous credit card number>` and then copy them one by one.

`fourcut` provides a bit more simple way of doing that by cutting the provided string to chunks of four then promtping the user and copying the chunks one by one.

# Install

The install script only works on Unix based systems. Pull requests are welcome for windows.

```sh
$ go get github.com/Shadowbeetle/fourcut
```

