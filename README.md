# hashgo

Hash files and display progress. Useful for large files or slow storage!

> [!WARNING]
> This project has been archived and moved to [Codeberg](https://codeberg.org/thumbscrw/hashgo)

## Usage

```
NAME:
   hashgo - Hash files (with a progress bar)

USAGE:
   hashgo [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --path value, -f value       path to file to hash [$HASHGO_PATH]
   --algorithm value, -a value  hashing algorithm to use (md5, sha1, sha224, sha256, sha384, sha512) [$HASHGO_ALGORITHM]
   --help, -h                   show helpNAME:
   hashgo - Hash files (with a progress bar)

USAGE:
   hashgo [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --path value, -f value       path to file to hash [$HASHGO_PATH]
   --algorithm value, -a value  hashing algorithm to use (md5, sha1, sha224, sha256, sha384, sha512) [$HASHGO_ALGORITHM]
   --help, -h                   show help
```

### Example

```shell
hashgo -f /path/to/file -a sha256
```
