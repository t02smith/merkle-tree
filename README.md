# Merkle Tree generator

This program breaks a file down into constant sized shards and generates a Merkle Tree
from it. This tree could be used to verify that the contents of the file have not been
tampered with and the file can be uniquely identified by the root hash of the tree.

> [More about Merkle Trees](https://en.wikipedia.org/wiki/Merkle_tree)

![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)

## How to run

> Ensure $GOPATH/bin is added to your PATH

```bash
go install
merkle-tree --help
```

```bash
# generate a merkle-tree of the generate.go file with shards of size 512 bytes
merkle-tree generate --shard-size 512 --filename "./merkletree/generate.go"
```
