{
  description = "A very basic flake";

  outputs = { self, nixpkgs }:
    let
      pkgs = import nixpkgs {
        system = "x86_64-linux";
      };
    in
    {

      devShell.x86_64-linux = pkgs.mkShell {
        buildInputs = [ pkgs.gopls ];
      };
      packages.x86_64-linux.hello = nixpkgs.legacyPackages.x86_64-linux.hello;

      defaultPackage.x86_64-linux = pkgs.stdenv.mkDerivation {
        name = "git-replayer";
      };

    };
}
