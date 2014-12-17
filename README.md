# check-ssh-chat

Used to check if a [ssh-chat](https://github.com/shazow/ssh-chat)
server is up, and responding.

## Usage

```bash
Usage of ./check-ssh-chat:
  -h="localhost": Hostname
  -n="check-ssh-chat": Username
  -p=22: Port
  -t=500ms: Timeout for the check
  -v=false: Verbose output
```

## Example

```bash
$ ./check-ssh-chat -h pi.c7.se -v
Checking: pi.c7.se:22
The ssh-chat server seems to be working
```

## MIT License

*Copyright (C) 2014 Peter Hellberg*

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
