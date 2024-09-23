# avail-go-sdk

An sdk to use avail network using Golang. This is wrapper around centrigue's [go-substrate-rpc-client](https://github.com/centrifuge/go-substrate-rpc-client/tree/master)

## prerequisites

[Install](https://go.dev/doc/install) go globally if you haven't already

## Structure

This SDK is split into two main parts:

1. **GSRPC Wrapper**: This allows you to use all the Polkadot JS functions and types to interact with the chain. For more information and documentation, please refer to the [GSRPC Documentation](https://pkg.go.dev/github.com/centrifuge/go-substrate-rpc-client/v4#section-readme).
2. **Opinionated SDK**: A simpler, more streamlined way to interact with the chain. It offers less customization but provides an easier interface. This SDK will be continuously improved to include everything needed for seamless chain interaction.

### Folder Structure

- **[`src/config`](./src/config/)**: Contains the loader to read the config file when connecting to chain
- **[`src/extrinsic`](./src/extrinsic/)**: Wrapper around GSRPC so that extrinsic can be signed with avail specific AppID extension
- **[src/header`](./src/header/)**: Includes the custom header for avail with added extension field and its custom enum decoding 
- **[`src/rpc`](./src/rpc/)**: Wraper around GSRPC block specific rpc calls inorder to accustom the custom header field for avail. Also contains the structures for kate related calls.
- **[`src/sdk`](./src/sdk/)**: Contains all interfaces related to the SDK, representing the opinionated part of Avail-go-sdk.
- **[`src/sdk/call`](./src/sdk/call/)**: Contains the kate related RPC calls
- **[`src/sdk/tx`](./src/sdk/tx/)**: Includes the interfaces related to sdk


## Examples
the [examples](./examples/) folder and the [readme](./examples/README.md) contains examples and documentation on using the avail-go-sdk

## Error Reporting

In case you encounter a bug, don't hesitate to [open an issue](https://github.com/availproject/avail-go-sdk/issues) with the maximum amount of detail and we will deal with it as soon as possible.