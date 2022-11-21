# Massa core plugin
This is a thyra plugin implementing Massa core features such as wallet.

## How to test ?

You will get the running port by looking at the messages in the terminal:
```shell
2022/11/21 22:11:43 Serving massa core at http://[::]:33049
warning: insecure HTTPS configuration.
	To fix this, use your own .crt and .key files using `--tls-certificate` and `--tls-key` flags
2022/11/21 22:11:43 Serving massa core at https://[::]:36869
```

In that case, you can accessthe GUI by going to http://[::]:33049/web/wallet/index.html

>**NOTE**: The port is to be changed.
