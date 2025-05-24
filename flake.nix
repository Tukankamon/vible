{
  description = "Bible command line tool written in go";
  #Doesnt work just yet

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in {
      packages.x86_64-linux.default = pkgs.buildGoModule {

        pname = "vible-cli";
        version = "1.0.0";

        src = pkgs.fetchFromGitHub {
          owner = "Tukankamon";
          repo = "vible";
          rev = "main";
          sha256 = "";
        };
        subPackages = [ "CLI" ];

        vendorHash = null;
        #proxyVendor = true;
      };
      devShells.x86_64-linux.default = pkgs.mkShell {

        buildInputs = with pkgs; [ go ];
      };
    };
}