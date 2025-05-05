# RNG Benchmark Report

## Generate 1 GB binary sample

```sh
GODEBUG=fips140=on go run cmd/rng-gen/rng-gen.go
```

## Dieharder

```sh
dieharder -a -g 201 -f rng.bin
```

## Results

[Raspberry-Pi-5](Raspberry-Pi-5.md)
