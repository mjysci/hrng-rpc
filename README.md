# HRNG RPC Service

An efficient Random Number Generation (RNG) service implementing JSON-RPC protocol. While this service provides robust randomization capabilities, we recommend established services such as [RANDOM.ORG](https://www.random.org/) or [ANU Quantum Numbers](https://quantumnumbers.anu.edu.au/) for applications requiring cryptographic-grade randomness.

## Configuration

The service configuration is managed through a YAML file, which can be located at:

- `config.yaml` (current directory)
- `~/.hrng-rpc/config.yaml` (home directory)

For custom configuration paths, specify the location using the `-config` parameter:

```bash
./hrng-rpc -config /path/to/config.yaml
```

### Configuration Parameters

The following parameters can be configured:

- `server.port`: Network port for the service (default: "8080")
- `trusted_proxies`: Array of authorized proxy IP addresses
- `apiKey`: Authentication key for API access (default: empty; unrestricted access permitted)

[Example configuration file](config.yaml.example)

## RPC Methods

### GenerateIntegers

Generates random integers within a user-defined range.

- `n`: Number of integers to generate (1-10000)
- `min`: Minimum value (-1e9 to 1e9)
- `max`: Maximum value (-1e9 to 1e9)
- `replacement`: Allow duplicate values (default: true)
- `base`: Output base (2, 8, 10, or 16; default: 10)

example:

```json
{
    "jsonrpc": "2.0",
    "method": "RandomService.GenerateIntegers",
    "params": {
        "apiKey": "00000000-0000-0000-0000-000000000000",
        "n": 512,
        "min": 0,
        "max": 255,
        "base": 16
    },
    "id": 42
}
```

### GenerateIntegerSequences

Generates uniform or multiform sequences of random integers.

- `n`: Number of sequences (1-1000)
- `length`: Length of each sequence or array of lengths (1-10000)
- `min`: Minimum value(s) (-1e9 to 1e9)
- `max`: Maximum value(s) (-1e9 to 1e9)
- `replacement`: Allow duplicates within sequences (default: true)
- `base`: Output base (2, 8, 10, or 16; default: 10)

### GenerateDecimalFractions

Generates random decimals from `[0,1)` interval.

- `n`: Number of decimals (1-10000)
- `decimalPlaces`: Decimal precision (1-14)
- `replacement`: Allow duplicates (default: true)

### GenerateGaussians

Generates random numbers from a Gaussian distribution.

- `n`: Number of values (1-10000)
- `mean`: Distribution mean (-1e6 to 1e6)
- `standardDeviation`: Standard deviation (-1e6 to 1e6)
- `significantDigits`: Result precision (2-14)

### GenerateStrings

Generates random strings.

- `n`: Number of strings (1-10000)
- `length`: String length (1-32)
- `characters`: Character set to use (1-128 chars)
- `replacement`: Allow duplicates (default: true)

### GenerateUUIDs

Generates version 4 UUIDs (RFC 4122).

- `n`: Number of UUIDs (1-1000)

### GenerateBlobs

Generates random binary data.

- `n`: Number of blobs (1-100)
- `size`: Blob size in bits (8-1048576, must be divisible by 8)
- `format`: Output format ("base64" or "hex")

## Testing

To validate the randomness quality, execute the following statistical tests:

```sh
sudo apt install dieharder
./rng-gen
dieharder -a -g 201 -f rng.bin
```

Detailed statistical analysis results are available in our [benchmark documentation](docs/benchmarks/README.md).

## License

Licensed under LGPL-2.1. See LICENSE for complete terms.

## Contributing

We welcome suggestions or bug reports! To submit a pull request:

1. Check the Issues tab; if there’s no relevant issue, feel free to create one.

2. Submit your PR and link it to a related issue.

## Acknowledgments

API design influenced by [RANDOM.ORG's JSON-RPC implementation](https://api.random.org/json-rpc/4/basic)
