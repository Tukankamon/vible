# A "TUI" tool for Bible verse lookup and continous reading

I am not a christian myself but wanted do this to get some practice with Go and TUI's.
![image](https://github.com/user-attachments/assets/24031af4-3b62-4281-b803-1c22064c3b51)

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
