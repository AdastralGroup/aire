# aire

## Note: This only contains the `verify` and `apply` commands!

Aire is the Adastral patching tool. It's a ~~very cut down version~~ fork of itch.io's Butler, with the non-essential bits sliced off.

This should only be used by emley really - to generate packages, you'll need the `diff` command, so just use itch.io's standard butler for the time being.

## Size Savings
To build for minimum size, run the following:
``
GOARCH=386 CGO_ENABLED=0 go build -gcflags=all="-l -B -C" -ldflags="-s -w"
``

On linux, this yields a binary of only 9.1MB. Using ``upx --best --lzma --brute --ultra-brute`` leads to a size of 2.4MB. Result!


## Documentation

Documentation for butler is available as a Gitbook:

  * :memo: <https://itch.io/docs/butler>

Notes about applying and verifying are still applicable to aire.

Questions about aire are welcome on its [Issue tracker](https://github.com/AdastralGroup/aire/issues).

## License

aire is released under the MIT License. See the [LICENSE](LICENSE) file for details.
