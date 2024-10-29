# tdl-whoami

`tdl-whoami` is a tdl extension that tells you who you are.

Further, it provides the details of the authenticated user, app remote config, and general config.

## Installation

1. **Install the `tdl` CLI**

    Refer to the [tdl documentation](https://docs.iyear.me/tdl/getting-started/installation/) for instructions.
    
2. **Install the `tdl-whoami` Extension**

    ```sh
    tdl extension install iyear/tdl-whoami
    ```

    Use proxy if you are behind a firewall:
    ```sh
    tdl extension install --proxy <PROXY> iyear/tdl-whoami
    ```

## Usage

### Basic Usage

To display name and ID of the authenticated user:

```sh
tdl whoami
```

### Verbose Mode

To display detailed information of authenticated user, app remote config, and general config.

```sh
tdl whoami -v
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

AGPL-3.0 License

