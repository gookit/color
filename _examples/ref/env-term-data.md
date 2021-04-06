# ENV grep term data

tests scripts:

```bash
env | grep -i term

# run examples
COLOR_DEBUG_MODE=on go run ./_examples/envcheck.go
go run ./_examples/envcheck.go

go run ./_examples/color_256.go
# test true color
sh ./_examples/ref/truecolor.sh
go run ./_examples/color_rgb.go
```

## Windows

**git bash**

```text
TERM_PROGRAM=mintty
TERM_PROGRAM_VERSION=3.4.4
TERM=xterm
```

**git bash on Windows terminal**

```text
TERM=xterm-256color
```

**git bash on vscode**

```text
COLORTERM=truecolor
TERM_PROGRAM=vscode
TERM_PROGRAM_VERSION=1.54.3
TERM=xterm-256color
```

**git bash on ConEmu**

```text
TERM=xterm-256color
```

```bash
$  env | grep -i ansi
ConEmuANSI=ON
ANSICON=104x32766 (104x26)
```

**git bash on JetBrains Terminal**

```text
TERMINAL_EMULATOR=JetBrains-JediTerm
TERM=xterm-256color
```

**cmd**

default not support color.

can enable VTP for support True color

**cmd on Windows terminal**

```text
TERM=xterm-256color
```

- support True color, but detect is `basic`

special ENV:

```text
WSLENV=WT_SESSION::WT_PROFILE_ID
WT_PROFILE_ID={574e775e-4f2a-5b96-ac1e-a2962a402336}
WT_SESSION=e68dfdf0-3f4e-4702-9b76-8768a6fbe784
```

**powerShell**

default not support color.

can enable VTP for support True color

**powerShell on windows terminal**

```text
TERM=xterm-256color
```

- support True color, but detect is `basic`

special ENV:

```text
WSLENV=WT_SESSION::WT_PROFILE_ID
WT_PROFILE_ID={574e775e-4f2a-5b96-ac1e-a2962a402336}
WT_SESSION=e68dfdf0-3f4e-4702-9b76-8768a6fbe784
```

### WSL

> tests on the `debian WSL`

- print `runtime.GOOS` is `Linux`

```bash
$ env | grep -i term
TERM=xterm-256color
```

special ENV:

```bash
$ env | grep -i wsl
WSL_DISTRO_NAME=Debian
WSLENV=WT_SESSION::WT_PROFILE_ID
```

## Linux

TODO

## Mac OS X

**zsh on Apple_Terminal**

```bash
% env | grep -i term 
TERM_PROGRAM=Apple_Terminal
TERM=xterm-256color
TERM_PROGRAM_VERSION=433
TERM_SESSION_ID=F17907FE-DCA5-488D-829B-7AFA8B323753
ZSH_TMUX_TERM=screen-256color
```

use screen:

```bash
% env | grep -i term                                                                                     
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERM_PROGRAM=Apple_Terminal
TERM_PROGRAM_VERSION=433
TERM_SESSION_ID=853CFB05-1F50-45A8-8F80-CF366958557E
ZSH_TMUX_TERM=screen
```

**zsh on iterm2**

```bash
% env | grep -i term   
LC_TERMINAL_VERSION=3.4.5beta1
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=3.4.5beta1
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
COLORTERM=truecolor
TERM=xterm-256color
ITERM_SESSION_ID=w0t2p0:3A53303E-BD72-4F1D-897D-EC15E3B4FDB5
ZSH_TMUX_TERM=screen-256color
```

use screen:

```bash
% env | grep -i term                                  
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERM_SESSION_ID=w0t2p0:3A53303E-BD72-4F1D-897D-EC15E3B4FDB5
LC_TERMINAL_VERSION=3.4.5beta1
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=3.4.5beta1
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
COLORTERM=truecolor
ITERM_SESSION_ID=w0t2p0:3A53303E-BD72-4F1D-897D-EC15E3B4FDB5
ZSH_TMUX_TERM=screen
```

**bin/zsh on JetBrains Terminal**

- support True color, but detect is `256`. so, need check `TERMINAL_EMULATOR`

```bash
% env | grep -i term 
TERM=xterm-256color
TERMINAL_EMULATOR=JetBrains-JediTerm
ZSH_TMUX_TERM=screen-256color
```

use screen:

```bash
% env | grep -i term                                                            
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERMINAL_EMULATOR=JetBrains-JediTerm
ZSH_TMUX_TERM=screen
```

**bin/zsh on Terminus Terminal**

- support True color, but detect is `256`. so, need check `TERM_PROGRAM`

```bash
% env | grep -i TERM                   
TERMINUS_PLUGINS=
TERM=xterm-256color
TERM_PROGRAM=Terminus
ZSH_TMUX_TERM=screen-256color
```
