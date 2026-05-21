{
  description = "Go dev shell (gopls, delve, gotools) with CGO + direnv";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";
  inputs.nixpkgs-unstable.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.systems.url = "github:nix-systems/default";
  inputs.flake-utils = {
    url = "github:numtide/flake-utils";
    inputs.systems.follows = "systems";
  };
  inputs.foundry.url = "github:shazow/foundry.nix/stable";

  outputs =
    { nixpkgs, nixpkgs-unstable, flake-utils, foundry, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ foundry.overlay ];
        };

        pkgsUnstable = import nixpkgs-unstable {
          inherit system;
        };

        _pkgGo = pkgsUnstable.go_1_26;
        _pkgGoModule = pkgsUnstable.buildGo126Module;

        _pkgGopls = pkgsUnstable.gopls.override {
          buildGoLatestModule = _pkgGoModule;
        };

        _pkgGoDelve = pkgsUnstable.delve.override {
          buildGoModule = _pkgGoModule;
        };

        _pkgGoTools = pkgsUnstable.gotools.override {
          buildGoModule = _pkgGoModule;
          go = _pkgGo;
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

          packages = [
            pkgs.foundry-bin
          ];

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
            go telemetry off

            echo "Go $(go version) • Telemetry=$(go telemetry 2>/dev/null || echo n/a) • CGO=$CGO_ENABLED"
          '';

          # macOS note (harmless on Linux): uncomment if you ever build Security-dependent CGO code
          # nativeBuildInputs = nativeBuildInputs ++ [
          #   pkgs.darwin.apple_sdk.frameworks.Security
          # ];
        };
      }
    );
}
