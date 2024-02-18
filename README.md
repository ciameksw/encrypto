# encrypto

`encrypto` is a command-line tool written in Go that encrypts and decrypts data using AES encryption.

## Installation

You can install `encrypto` by running the following command:

```bash
go install github.com/ciameksw/encrypto@latest
```

## Usage

### Encryption

To encrypt a file, use the `encrypt` command followed by the `-f` flag for the file name and the `-k` flag for the key. 
The key flag is optional, when key is not provided a new one will be generated.

```bash
encrypto encrypt -f <FILE_NAME> -k <KEY>
```

### Decryption

To decrypt a file, use the `decrypt` command followed by the `-f` flag for the file name and the `-k` flag for the key.

```bash
encrypto decrypt -f <FILE_NAME> -k <KEY>
```

