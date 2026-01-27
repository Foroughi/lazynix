{
  description = "lazy git repo";

  inputs.nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {

          buildInputs = with pkgs; [
              pkgs.go           
              pkgs.gopls
          ];

          shellHook = ''
          '';

          NIX_NAME = "Lazy Git";
        };
      });

}
