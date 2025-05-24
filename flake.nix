{
  description = "Bible command line tool written in go";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in {
      packages.x86_64-linux.default = pkgs.buildGoModule rec {  #rec allows the use of ${pname}

        pname = "vible-cli";
        version = "1.0.0";

        src = pkgs.fetchFromGitHub {
          owner = "Tukankamon";
          repo = "vible";
          rev = "main";    #Specific commit, will need to update the hash every update if it is a branch
          sha256 = "sha256-Gyvlkc+z4TIjYwjbbXY0H2mI++myOAJv+60p3y6FjuE=";
        };

        vendorHash = null;
        #proxyVendor = true;

        subPackages = [ "." ];  #This skips the archive folder

        nativeBuildInputs = [ pkgs.makeWrapper ];   #GPT recommendation
        postBuild = ''
          echo "Nix build directory: $PWD"
          mkdir -p $out/share
          cp -r ${./bible} $out/share/bible
        ''; # Downloads all the text files, could make a more minimal version
      };
      devShells.x86_64-linux.default = pkgs.mkShell {

        buildInputs = with pkgs; [ go ];
      };
    };
}