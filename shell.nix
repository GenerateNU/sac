{ pkgs ? import <nixpkgs> {} }:

let
  unstable = import <nixos-unstable> { config = { allowUnfree = true; }; };
in
pkgs.mkShell {
  hardeningDisable = [ "fortify" ];
  
  packages = with pkgs; [
    unstable.go_1_22
    unstable.delve
    unstable.gopls
    unstable.go-tools
  ];
}