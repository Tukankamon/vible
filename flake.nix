{
  description = "Bible command line tool written in Go";

  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = {
    self, #Lsp might say its not used bu it is
    nixpkgs,
  }: let
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in {
    packages.x86_64-linux.default = pkgs.buildGoModule {
      #if you need to do something like ${} add rec

      pname = "vible";
      version = "1.0.0";

      src = pkgs.fetchFromGitHub {
        owner = "Tukankamon";
        repo = "vible";
        rev = "main"; #Specific commit or branch, will need to update the hash every update if it is a branch
        sha256 = "sha256-E436U1lp0sZK9tE6n/oHUAjdJ9sWjenknUV2ObMTDlM=";
      };

      buildInputs = with pkgs; [go];

      vendorHash = null; #Couldnt find a way to do it without a vendor folder
      #proxyVendor = true;

      subPackages = ["app/main" "app/backend"]; #This skips the archive folder

      postBuild = ''
        echo "Nix build directory: $PWD"
        mkdir -p $out/share
        cp -r ${./bible} $out/share/bible
      ''; # Downloads all the text files, could make a more minimal version

      postInstall = ''
        mv $out/bin/main $out/bin/vible
      '';
    };
  };
}
