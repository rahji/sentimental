# Sentimental

Sentimental is an executable that starts an OSC server and waits for a message that looks like `/text STRING`

When it receives that message, it does sentiment analysis on the string and sends the resulting float32 (negative number
is negative sentiment, positive number is positive sentiment) via OSC to a specified server address and port.

## Usage

```
Usage: sentimental [flags]

Flags:
  -h, --help                       Show context-sensitive help.
      --server-addr="127.0.0.1"    OSC server address
      --server-port=8884           OSC server port
      --client-addr="127.0.0.1"    OSC client address
      --client-port=8885           OSC client port
      --verbose                    Show extra output in the terminal
```

*Note that if you run this from WSL2 and expect to communicate with Windows on the same machine, you'll probably want to
get the server address using `ifconfig` in WSL and the client address using `ipconfig` in Command Prompt. It's unlikely to
work with the default server addresses of `127.0.0.1` in this specific use case.*

## Installation

Download the appropriate archive file from the [https://github.com/rahji/sentimental/releases/latest](Releases)
page, place the `sentimental` binary [https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/](somewhere in your path),
and run it from your terminal (eg: Terminal.app in MacOS or [https://apps.microsoft.com/store/detail/windows-terminal/9N0DX20HK701?hl=en-us&gl=us&rtc=1](Windows Terminal))
