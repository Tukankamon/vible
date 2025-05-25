{
  description = "Bible command line tool written in go";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in {
      packages.x86_64-linux.default = pkgs.buildGoModule rec {  #rec allows the use of ${pname}

        pname = "vible";
        version = "1.0.0";

        src = pkgs.fetchFromGitHub {
          owner = "Tukankamon";
          repo = "vible";
          rev = "main";    #Specific commit, will need to update the hash every update if it is a branch
          sha256 = "sha256-PssycwV47kxoW5dW2L6E3MTOQmViioTy4mr6ftjuGx0=";
        };

        vendorHash = null;
        #proxyVendor = true;

        subPackages = [ "app" ];  #This skips the archive folder

        nativeBuildInputs = [ pkgs.makeWrapper ];   #GPT recommendation
        postBuild = ''
          echo "Nix build directory: $PWD"
          mkdir -p $out/share
          cp -r ${./app/bible} $out/share/bible
        ''; # Downloads all the text files, could make a more minimal version

      };
      devShells.x86_64-linux.default = pkgs.mkShell {

        buildInputs = with pkgs; [ go ];
      };
    };
}