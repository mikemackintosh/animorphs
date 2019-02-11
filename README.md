# animorphs
Animorphs is a config conversion utility

## Installation

Using `go get`, download the package:

    go get -u  github.com/mikemackintosh/animorphs/...

The binaries will be added to `$GOPATH/bin/` automatically. If your `$PATH` executes from `$GOPATH/bin` you're all set.

## Usage

The utilities read from `stdin`.

    $ echo '{"json": "value", "subkey":{"array":["a", "b", "c"]}}' | json2yaml
    json: value
    subkey:
      array:
      - a
      - b
      - c 
    
