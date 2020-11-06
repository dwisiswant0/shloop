# shLoop

Want to execute command repeatedly without _workache_? Here is **shloop** born for it!

## Install

The installation is easy!

- You can download a prebuilt binary from [releases page](https://github.com/dwisiswant0/shloop/releases), unpack and run! or
- If you have Go compiler installed and configured:

```bash
> go get github.com/dwisiswant0/shloop/...
```

## Configuration

`shLoop` just reads environment values to run.

- `SHLOOP_COUNT` is the number of times the command is repeated _(default: 5)_,
- `SHLOOP_THREAD` to specify the number of executed command concurrently _(default: 1)_, and
- `SHLOOP_VERBOSE` to display information on what command is being executed _(default: false)_.

## Usage

Simply, `shloop` can be run with:

```bash
▶ shloop [command...]
```

### Variables

We also support variables, meanwhile we have variables as:

- `i` will generate digit incrementally _(depending on_ `SHLOOP_COUNT` _value)_,
- `str` will generate a random ASCII string with length _N_,
- `int` will generate a random string of length _N_ consists of ASCII digits _(note it can start with 0)_, and
- `uuid` _(version 4)_ will generate a random unique identifier based upon random nunbers.
	- Format: `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`.

## Workarounds

As an example, I want to repeat the `echo` command 10 times.

```bash
▶ export SHLOOP_COUNT=10
▶ shloop echo "Hello, world!"
Hello, world!
Hello, world!
Hello, world!
Hello, world!
Hello, world!
Hello, world!
Hello, world!
Hello, world!
Hello, world!
Hello, world!
```

In case if you want to test race conditions on an endpoint with `curl`.

```bash
▶ export SHLOOP_COUNT=10
▶ export SHLOOP_THREAD=5
▶ shloop curl -X POST https://httpbin.org/post -d "foo=bar"
```

### Using variables

Example of using incremental variable:

```bash
▶ export SHLOOP_COUNT=5
▶ shloop echo "No. {{i}}"
No. 1
No. 2
No. 3
No. 4
No. 5
```

**Have IDOR endpoints? Intruder on it!**

```bash
▶ export SHLOOP_COUNT=100
▶ export SHLOOP_VERBOSE=true
▶ shloop curl -s "https://domain/api/orders/{{int 8}}"
```

Another one:

```bash
▶ export SHLOOP_COUNT=100
▶ export SHLOOP_VERBOSE=true
▶ shloop curl -s "https://domain/profile.php?id={{uuid}}"
```
