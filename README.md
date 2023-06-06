# opside tool

## keystore
require go1.20.1

```shell
cd keystore
go build

./keystore 
    NAME:
    keystore - A new cli application

    USAGE:
    keystore [global options] command [command options] [arguments...]

    COMMANDS:
    encryptKey  Encrypts the privatekey with a password and create a keystore file
    help, h     Shows a list of commands or help for one command

    GLOBAL OPTIONS:
    --help, -h  show help
```

```shell
./keystore encryptKey --pk 'pk' --pw 'password' -o 'keystore file'
```

