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
          rev = "main";    #Specific commit, will need to update the hash every update if it is a branch
          sha256 = "sha256-1Y/sha256-pGkt/D+kQ9ebVrbjoy42oa2YwKDgCn5E+ztBly/ELKY=";
        };
        #subPackages = [ "CLI" ];

        vendorHash = null;
        #proxyVendor = true;

        subPackages = [ "." ];  #This skips the archive folder
      };
      devShells.x86_64-linux.default = pkgs.mkShell {

        buildInputs = with pkgs; [ go ];
      };
    };
}