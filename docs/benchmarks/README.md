# RNG Benchmark Report

## Generate 1 GB binary sample

```sh
go install github.com/mjysci/hrng-rpc/cmd/rng-gen@latest
GODEBUG=fips140=on rng-gen
```

## Dieharder

```sh
dieharder -a -g 201 -f rng.bin
```

## Results

[Raspberry Pi 5](Raspberry-Pi-5.md)  
[Raspberry Pi 4B](Raspberry-Pi-4B.md)  
[NVIDIA Jetson Nano 2GB Developer Kit](NVIDIA-Jetson-Nano-2GB-Developer-Kit.md)  
[x86-64](x86-64.md)  
[i.MX 6ULL](i.MX-6ULL.md)  
[Raspberry Pi 1B](Raspberry-Pi-1B.md)  
[STM32MP157](STM32MP157.md)  
