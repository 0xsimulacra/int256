{
  description = "Go dev shell (gopls, delve, gotools) with CGO + direnv";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.systems.url = "github:nix-systems/default";
  inputs.flake-utils = {
    url = "github:numtide/flake-utils";
    inputs.systems.follows = "systems";
  };

  outputs =
    {
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };

        _pkgGo = pkgs.go_1_27;
        _pkgGoModule = pkgs.buildGo126Module;

        _pkgGopls = pkgs.gopls.override {
          buildGoLatestModule = _pkgGoModule;
        };

        _pkgGoDelve = pkgs.delve.override {
          buildGoModule = _pkgGoModule;
        };

        _pkgGoTools = pkgs.gotools.override {
          buildGoModule = _pkgGoModule;
          go = _pkgGo;
        };

        _pkgGoPerf = pkgs.goperf.override {
          buildGoModule = _pkgGoModule;
        };

      in
      {
        devShells.default = pkgs.mkShell {
          hardeningDisable = [ "all" ];

          nativeBuildInputs = [
            _pkgGo
            _pkgGopls
            _pkgGoDelve
            _pkgGoTools
            _pkgGoPerf

            pkgs.git
            pkgs.curl
            pkgs.cacert

            pkgs.pkg-config
            pkgs.gcc
          ];

          buildInputs = [
            # cgo libs
            pkgs.openssl
            pkgs.zlib
            pkgs.sqlite
          ];

          packages = [ ];

          GOOS = "linux";
          GOARCH = "amd64";
          GOEXPERIMENT = "simd";
          CGO_ENABLED = 1;

          LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath [
            pkgs.openssl
            pkgs.zlib
            pkgs.sqlite
            pkgs.stdenv.cc.cc.lib
          ];

          shellHook = ''
            echo "Go $(go version) • CGO=$CGO_ENABLED"
          '';

          # macOS note (harmless on Linux): uncomment if you ever build Security-dependent CGO code
          # nativeBuildInputs = nativeBuildInputs ++ [
          #   pkgs.darwin.apple_sdk.frameworks.Security
          # ];
        };
      }
    );
}
