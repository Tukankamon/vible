{
  description = "Bible command line tool written in go";

  inputs.nixpkgs.url = "github:nixos/nixpkgs";

  outputs = { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in {
      packages.x86_64-linux.default = pkgs.buildGoModule {
        pname = "vible-cli";
        version = "1.0.0";
        src = ./.;
        vendorHash = null;
        proxyVendor = true;
      };
      devShells.x86_64-linux.default = pkgs.mkShell {
        buildInputs = with pkgs; [ go ];
      };
    };
}