# Raspberry Pi 1B RNG Benchmark

BCM2835

## Dieharder

```sh
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
 file_input_raw|                         rng.bin|  8.91e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.52468723|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.04654988|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.13383855|  PASSED  
# The file file_input_raw was rewound 1 times
    diehard_rank_6x8|   0|    100000|     100|0.56118714|  PASSED  
# The file file_input_raw was rewound 1 times
   diehard_bitstream|   0|   2097152|     100|0.97506327|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_opso|   0|   2097152|     100|0.23826192|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_oqso|   0|   2097152|     100|0.80013622|  PASSED  
# The file file_input_raw was rewound 2 times
         diehard_dna|   0|   2097152|     100|0.02195558|  PASSED  
# The file file_input_raw was rewound 2 times
diehard_count_1s_str|   0|    256000|     100|0.10195336|  PASSED  
# The file file_input_raw was rewound 3 times
diehard_count_1s_byt|   0|    256000|     100|0.61023921|  PASSED  
# The file file_input_raw was rewound 3 times
 diehard_parking_lot|   0|     12000|     100|0.27155245|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_2dsphere|   2|      8000|     100|0.86305997|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_3dsphere|   3|      4000|     100|0.33771757|  PASSED  
# The file file_input_raw was rewound 4 times
     diehard_squeeze|   0|    100000|     100|0.46936258|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_sums|   0|       100|     100|0.54350176|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_runs|   0|    100000|     100|0.65122602|  PASSED  
        diehard_runs|   0|    100000|     100|0.90781752|  PASSED  
# The file file_input_raw was rewound 4 times
       diehard_craps|   0|    200000|     100|0.99913806|   WEAK   
       diehard_craps|   0|    200000|     100|0.83078240|  PASSED  
# The file file_input_raw was rewound 12 times
 marsaglia_tsang_gcd|   0|  10000000|     100|0.00167743|   WEAK   
 marsaglia_tsang_gcd|   0|  10000000|     100|0.00094421|   WEAK   
# The file file_input_raw was rewound 12 times
         sts_monobit|   1|    100000|     100|0.99980270|   WEAK   
# The file file_input_raw was rewound 12 times
            sts_runs|   2|    100000|     100|0.49986320|  PASSED  
# The file file_input_raw was rewound 12 times
          sts_serial|   1|    100000|     100|0.88698272|  PASSED  
          sts_serial|   2|    100000|     100|0.22019385|  PASSED  
          sts_serial|   3|    100000|     100|0.64902035|  PASSED  
          sts_serial|   3|    100000|     100|0.35918419|  PASSED  
          sts_serial|   4|    100000|     100|0.65698444|  PASSED  
          sts_serial|   4|    100000|     100|0.37363825|  PASSED  
          sts_serial|   5|    100000|     100|0.52823977|  PASSED  
          sts_serial|   5|    100000|     100|0.54251790|  PASSED  
          sts_serial|   6|    100000|     100|0.80977204|  PASSED  
          sts_serial|   6|    100000|     100|0.70491703|  PASSED  
          sts_serial|   7|    100000|     100|0.34336054|  PASSED  
          sts_serial|   7|    100000|     100|0.34650765|  PASSED  
          sts_serial|   8|    100000|     100|0.94924759|  PASSED  
          sts_serial|   8|    100000|     100|0.97426550|  PASSED  
          sts_serial|   9|    100000|     100|0.23166695|  PASSED  
          sts_serial|   9|    100000|     100|0.40857545|  PASSED  
          sts_serial|  10|    100000|     100|0.80148533|  PASSED  
          sts_serial|  10|    100000|     100|0.19277946|  PASSED  
          sts_serial|  11|    100000|     100|0.62874904|  PASSED  
          sts_serial|  11|    100000|     100|0.20762602|  PASSED  
          sts_serial|  12|    100000|     100|0.45539433|  PASSED  
          sts_serial|  12|    100000|     100|0.88659355|  PASSED  
          sts_serial|  13|    100000|     100|0.42376393|  PASSED  
          sts_serial|  13|    100000|     100|0.54343268|  PASSED  
          sts_serial|  14|    100000|     100|0.38106884|  PASSED  
          sts_serial|  14|    100000|     100|0.28444227|  PASSED  
          sts_serial|  15|    100000|     100|0.58730718|  PASSED  
          sts_serial|  15|    100000|     100|0.43087716|  PASSED  
          sts_serial|  16|    100000|     100|0.95598105|  PASSED  
          sts_serial|  16|    100000|     100|0.74953244|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   1|    100000|     100|0.72102988|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   2|    100000|     100|0.99541575|   WEAK   
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   3|    100000|     100|0.05926633|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   4|    100000|     100|0.13395506|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   5|    100000|     100|0.77005738|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   6|    100000|     100|0.47402277|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   7|    100000|     100|0.43977734|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   8|    100000|     100|0.97112536|  PASSED  
# The file file_input_raw was rewound 15 times
         rgb_bitdist|   9|    100000|     100|0.99042854|  PASSED  
# The file file_input_raw was rewound 16 times
         rgb_bitdist|  10|    100000|     100|0.99262206|  PASSED  
# The file file_input_raw was rewound 17 times
         rgb_bitdist|  11|    100000|     100|0.94487735|  PASSED  
# The file file_input_raw was rewound 18 times
         rgb_bitdist|  12|    100000|     100|0.13048304|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   2|     10000|    1000|0.61694712|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   3|     10000|    1000|0.41062907|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   4|     10000|    1000|0.68949547|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   5|     10000|    1000|0.83648659|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   2|    100000|     100|0.63748546|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   3|    100000|     100|0.79540683|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   4|    100000|     100|0.80767326|  PASSED  
# The file file_input_raw was rewound 19 times
    rgb_permutations|   5|    100000|     100|0.97161968|  PASSED  
# The file file_input_raw was rewound 19 times
      rgb_lagged_sum|   0|   1000000|     100|0.75562863|  PASSED  
# The file file_input_raw was rewound 20 times
      rgb_lagged_sum|   1|   1000000|     100|0.27504993|  PASSED  
# The file file_input_raw was rewound 21 times
      rgb_lagged_sum|   2|   1000000|     100|0.22229465|  PASSED  
# The file file_input_raw was rewound 22 times
      rgb_lagged_sum|   3|   1000000|     100|0.12539983|  PASSED  
# The file file_input_raw was rewound 24 times
      rgb_lagged_sum|   4|   1000000|     100|0.15659985|  PASSED  
# The file file_input_raw was rewound 26 times
      rgb_lagged_sum|   5|   1000000|     100|0.12847911|  PASSED  
# The file file_input_raw was rewound 29 times
      rgb_lagged_sum|   6|   1000000|     100|0.28410776|  PASSED  
# The file file_input_raw was rewound 32 times
      rgb_lagged_sum|   7|   1000000|     100|0.00041571|   WEAK   
# The file file_input_raw was rewound 35 times
      rgb_lagged_sum|   8|   1000000|     100|0.38496664|  PASSED  
# The file file_input_raw was rewound 39 times
      rgb_lagged_sum|   9|   1000000|     100|0.42337692|  PASSED  
# The file file_input_raw was rewound 43 times
      rgb_lagged_sum|  10|   1000000|     100|0.29056504|  PASSED  
# The file file_input_raw was rewound 48 times
      rgb_lagged_sum|  11|   1000000|     100|0.07856598|  PASSED  
# The file file_input_raw was rewound 53 times
      rgb_lagged_sum|  12|   1000000|     100|0.20357499|  PASSED  
# The file file_input_raw was rewound 58 times
      rgb_lagged_sum|  13|   1000000|     100|0.25747353|  PASSED  
# The file file_input_raw was rewound 63 times
      rgb_lagged_sum|  14|   1000000|     100|0.70029687|  PASSED  
# The file file_input_raw was rewound 69 times
      rgb_lagged_sum|  15|   1000000|     100|0.00000004|  FAILED  
# The file file_input_raw was rewound 76 times
      rgb_lagged_sum|  16|   1000000|     100|0.90937043|  PASSED  
# The file file_input_raw was rewound 82 times
      rgb_lagged_sum|  17|   1000000|     100|0.27386170|  PASSED  
# The file file_input_raw was rewound 89 times
      rgb_lagged_sum|  18|   1000000|     100|0.63915343|  PASSED  
# The file file_input_raw was rewound 97 times
      rgb_lagged_sum|  19|   1000000|     100|0.19792273|  PASSED  
# The file file_input_raw was rewound 105 times
      rgb_lagged_sum|  20|   1000000|     100|0.96876456|  PASSED  
# The file file_input_raw was rewound 113 times
      rgb_lagged_sum|  21|   1000000|     100|0.47515393|  PASSED  
# The file file_input_raw was rewound 121 times
      rgb_lagged_sum|  22|   1000000|     100|0.21592439|  PASSED  
# The file file_input_raw was rewound 130 times
      rgb_lagged_sum|  23|   1000000|     100|0.00013324|   WEAK   
# The file file_input_raw was rewound 140 times
      rgb_lagged_sum|  24|   1000000|     100|0.49603547|  PASSED  
# The file file_input_raw was rewound 149 times
      rgb_lagged_sum|  25|   1000000|     100|0.93826054|  PASSED  
# The file file_input_raw was rewound 159 times
      rgb_lagged_sum|  26|   1000000|     100|0.02341129|  PASSED  
# The file file_input_raw was rewound 170 times
      rgb_lagged_sum|  27|   1000000|     100|0.08610824|  PASSED  
# The file file_input_raw was rewound 181 times
      rgb_lagged_sum|  28|   1000000|     100|0.06503873|  PASSED  
# The file file_input_raw was rewound 192 times
      rgb_lagged_sum|  29|   1000000|     100|0.33758282|  PASSED  
# The file file_input_raw was rewound 203 times
      rgb_lagged_sum|  30|   1000000|     100|0.70342340|  PASSED  
# The file file_input_raw was rewound 215 times
      rgb_lagged_sum|  31|   1000000|     100|0.00045301|   WEAK   
# The file file_input_raw was rewound 228 times
      rgb_lagged_sum|  32|   1000000|     100|0.29533367|  PASSED  
# The file file_input_raw was rewound 228 times
     rgb_kstest_test|   0|     10000|    1000|0.04786015|  PASSED  
# The file file_input_raw was rewound 228 times
     dab_bytedistrib|   0|  51200000|       1|0.91657247|  PASSED  
# The file file_input_raw was rewound 228 times
             dab_dct| 256|     50000|       1|0.88533081|  PASSED  
Preparing to run test 207.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_filltree|  32|  15000000|       1|0.12415585|  PASSED  
        dab_filltree|  32|  15000000|       1|0.43888092|  PASSED  
Preparing to run test 208.  ntuple = 0
# The file file_input_raw was rewound 229 times
       dab_filltree2|   0|   5000000|       1|0.28787180|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.55429476|  PASSED  
Preparing to run test 209.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_monobit2|  12|  65000000|       1|0.78797885|  PASSED
```
