# A CLI tool for Bible verse lookup and continous reading

I am not a christian myself but wanted do this to get some practice with Go and TUI's.

# Installation:

## Nix / Nixos

```nix
  # flake.nix
    vible = {
      url = "github:Tukankamon/vible";
      inputs.nixpkgs.follows = "nixpkgs";
    };
```
```nix
# configuration.nix
  environment.systemPackages = wwith pkgs; [
    inputs.vible.packages.x86_64-linux.default
]
```

## Anythin else

Build from source (go build / go run)



TODO:
- [ ] Interface (not just text)
- [ ] VIM binds
- [ ] More translations/languages
