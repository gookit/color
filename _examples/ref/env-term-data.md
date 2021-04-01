# ENV grep term data

by run

```bash
env | grep -i term
```

## Windows

**git bash**

```text
TERM_PROGRAM=mintty
TERM_PROGRAM_VERSION=3.4.4
TERM=xterm
```

**git bash on windows terminal**

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

can by enable VTP for support True color

**cmd on windows terminal**

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

can by enable VTP for support True color

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

## Unix

contains `Mac OS X/Linux`

