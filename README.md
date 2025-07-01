# A "TUI" tool for Bible verse lookup and continous reading

I am not a christian myself but wanted do this to get some practice with Go and TUI's.
![image](https://github.com/user-attachments/assets/a699edbd-6ad1-4c40-8b0b-fa4d8fa04f5c)

![image](https://github.com/user-attachments/assets/fbb8caa9-4d05-44fc-a2b7-3016785178f9)



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

## Anything else

Build from source (go build / go run)



# TODO:
- [x] Interface (not just text)
- [X] VIM binds
- [ ] More translations/languages
