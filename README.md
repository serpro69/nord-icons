# nord-icons

A simple app that converts [candy-icons](https://github.com/EliverLara/candy-icons) to use colors from the awesome [Nord](https://github.com/arcticicestudio/nord) palette.

## Usage

- initialize the `candy-icons` git submodule: `git submodule sync`
- copy existing `candy-icons` to `out` dir (this is needed to preserve symlinks): `mkdir out && cp candy-icons out/nord-icons -r`
- rename the icon theme: `sed -i 's/Name=.*/Name=nord-icons/' out/nord-icons/index.theme`
- run the `main.go` file: `go run .`
  - this can be executed multiple times until desirable colors are generated (each run generates random colors from Nord palettes)
- copy the `out/nord-icons` dir to a directory with icons, e.g. into a user's home dir `cp out/nord-icons ~/.local/share/icons/. -r`
  - (optionally) copy the icons from this repo as well: `cp nord-icons/* ~/.local/share/icons/nord-icons/. -r`
