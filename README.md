[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner2-direct.svg)](https://stand-with-ukraine.pp.ua)

# nord-icons

A simple app that converts [candy-icons](https://github.com/EliverLara/candy-icons) to use colors from the awesome [Nord](https://github.com/arcticicestudio/nord) palette.

## Usage

You can install the icons manually or use a small script which will do (almost) everything for you.

Whichever you choose, you'll need to clone this repo first somewhere and then proceed with one of the below installation options.

```shell
tempdir=$(mktemp -d -u) && git clone git@github.com:serpro69/nord-icons.git "$tempdir" && cd "$tempdir"
```

### Install script

Run the [`install.sh`](install.sh) script which will handle everything for you: `./install.sh --copy -p <path_to_icons_dir>`

Use the `--help` option to get all usage details.

### Manual

You can also do the same manually with the following steps:

- initialize the `candy-icons` git submodule: `git submodule sync`
- copy existing `candy-icons` to `out` dir (this is needed to preserve symlinks): `mkdir out && cp candy-icons out/nord-icons -r`
- rename the icon theme: `sed -i 's/Name=.*/Name=nord-icons/' out/nord-icons/index.theme`
- run the `main.go` file: `go run .`
  - this can be executed multiple times until desirable colors are generated (each run generates random colors from Nord palettes)
- copy the `out/nord-icons` dir to a directory with icons, e.g. into a user's home dir `cp out/nord-icons ~/.local/share/icons/. -r`
  - (optionally) copy the icons from this repo as well: `cp nord-icons/* ~/.local/share/icons/nord-icons/. -r`
