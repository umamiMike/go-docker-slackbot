{ pkgs ? import <nixpkgs> {} }:

with pkgs;

let

in
mkShell {

  buildInputs = [
  bash
  go
  goimports
  fswatch
  tmux
  nodejs
];

buildPhase = ''

'';
 shellHook = ''
  glibcLocales=$(nix-build --no-out-link "<nixpkgs>" -A glibcLocales)
  export LOCALE_ARCHIVE_2_27="${glibcLocales}/lib/locale/locale-archive"

  export GOPATH="$(pwd)/.go"
  exportGOCACHE=""

 '';
}
