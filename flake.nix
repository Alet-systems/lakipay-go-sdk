{
  description = "A reproducible Go development environment for lakipay go sdk";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable"; # Or a specific branch/commit
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go # The Golang compiler and tools
            gopls # Language server
            delve # Debugger (if needed)
            # Other tools like git, make, etc.
          ];

          # Optional: set environment variables or run commands on shell entry
          shellHook = ''
            echo "Welcome to the Go development environment!"
          '';
        };
      });
}
