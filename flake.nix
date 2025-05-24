{
  description = "Bible command line tool written in go";

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
          sha256 = "sha256-srp7WEs+KLpdQgGfLlsWLJJsthr9xPSm+hwbe5VkmFw=";
        };
        #subPackages = [ "CLI" ];

        vendorHash = "sha256-srp7WEs+KLpdQgGfLlsWLJJsthr9xPSm+hwbe5VkmFw=";
        #proxyVendor = true;

        subPackages = [ "." ];  #This skips the archive folder
      };
      devShells.x86_64-linux.default = pkgs.mkShell {

        buildInputs = with pkgs; [ go ];
      };
    };
}