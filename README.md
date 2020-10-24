![logo](./doc/logo.png)

_It takes their jacket, and seats them at the best table._

Toadstool is a Gnome extension helper.

### Usage

```
A fast and convenient tool for managing Gnome extensions.

Usage:
  toadstool [command]

Available Commands:
  help        Help about any command
  install     Install a Gnome extension
  list        List installed Gnome extensions
  uninstall   Uninstall a Gnome extension

Flags:
      --config string           config file (default is $HOME/.toadstool)
      --extensions-dir string   Gnome extensions directory override (default is $HOME/.local/share/gnome-shell/extensions)
  -h, --help                    help for toadstool

Use "toadstool [command] --help" for more information about a command.
```

#### Install
Install extension at path `-e`. Must be a valid .zip file:
```
  -e, --extension string   Gnome extension to install
```

#### List
List installed extensions located in `$HOME/.local/share/gnome-shell/extensions` (or override)

#### Uninstall
Uninstall will display a list of installed extensions. Select extension to install from list, and it will be removed. Toadstool will then restart the Gnome shell, displaying a `Restarting...` message as it does.



