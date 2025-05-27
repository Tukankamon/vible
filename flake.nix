{
  description = "Bible command line tool written in Go";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in {
      packages.x86_64-linux.default = pkgs.buildGoModule rec {  #if you need to do something like ${} add rec

        pname = "vible";
        version = "1.0.0";

        src = pkgs.fetchFromGitHub {
          owner = "Tukankamon";
          repo = "vible";
          rev = "main";    #Specific commit or branch, will need to update the hash every update if it is a branch
          sha256 = "sha256-g603uN0d4rRbGf95DYyIIBjL+1IvbbSwTi2d0HyXHdU=";
        };

        buildInputs = with pkgs; [ go ];

        vendorHash = "sha256-4rK69s1uTFBV20endymLw6JEUfrh51bznZEgbujUQls=";   #Couldnt find a way to do it without a vendor folder
        #proxyVendor = true;

        subPackages = [ "app/main" "app/backend" ];  #This skips the archive folder

        postBuild = ''
          echo "Nix build directory: $PWD"
          mkdir -p $out/share
          cp -r ${./bible} $out/share/bible
        ''; # Downloads all the text files, could make a more minimal version

      };

    };
}