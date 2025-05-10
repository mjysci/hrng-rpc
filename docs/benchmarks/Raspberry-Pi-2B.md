# Raspberry Pi 2B RNG Benchmark

BCM2836
3f104000.rng
Linux rpi2b 6.12.25+rpt-rpi-v7 #1 SMP Raspbian 1:6.12.25-1+rpt1 (2025-04-30) armv7l GNU/Linux

## Dieharder

```sh
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
 file_input_raw|                         rng.bin|  9.13e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.99802076|   WEAK   
      diehard_operm5|   0|   1000000|     100|0.21954417|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.47954566|  PASSED  
# The file file_input_raw was rewound 1 times
    diehard_rank_6x8|   0|    100000|     100|0.61239120|  PASSED  
# The file file_input_raw was rewound 1 times
   diehard_bitstream|   0|   2097152|     100|0.76819580|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_opso|   0|   2097152|     100|0.44912284|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_oqso|   0|   2097152|     100|0.98248650|  PASSED  
# The file file_input_raw was rewound 2 times
         diehard_dna|   0|   2097152|     100|0.23026389|  PASSED  
# The file file_input_raw was rewound 2 times
diehard_count_1s_str|   0|    256000|     100|0.86508035|  PASSED  
# The file file_input_raw was rewound 3 times
diehard_count_1s_byt|   0|    256000|     100|0.54822581|  PASSED  
# The file file_input_raw was rewound 3 times
 diehard_parking_lot|   0|     12000|     100|0.77995133|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_2dsphere|   2|      8000|     100|0.33425068|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_3dsphere|   3|      4000|     100|0.57726386|  PASSED  
# The file file_input_raw was rewound 4 times
     diehard_squeeze|   0|    100000|     100|0.95247239|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_sums|   0|       100|     100|0.02916577|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_runs|   0|    100000|     100|0.24133422|  PASSED  
        diehard_runs|   0|    100000|     100|0.21514614|  PASSED  
# The file file_input_raw was rewound 4 times
       diehard_craps|   0|    200000|     100|0.65852868|  PASSED  
       diehard_craps|   0|    200000|     100|0.81360461|  PASSED  
# The file file_input_raw was rewound 12 times
 marsaglia_tsang_gcd|   0|  10000000|     100|0.24831420|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.00007985|   WEAK   
# The file file_input_raw was rewound 12 times
         sts_monobit|   1|    100000|     100|0.91082010|  PASSED  
# The file file_input_raw was rewound 12 times
            sts_runs|   2|    100000|     100|0.05005668|  PASSED  
# The file file_input_raw was rewound 12 times
          sts_serial|   1|    100000|     100|0.69179535|  PASSED  
          sts_serial|   2|    100000|     100|0.83890394|  PASSED  
          sts_serial|   3|    100000|     100|0.97824768|  PASSED  
          sts_serial|   3|    100000|     100|0.72651837|  PASSED  
          sts_serial|   4|    100000|     100|0.91390996|  PASSED  
          sts_serial|   4|    100000|     100|0.99613344|   WEAK   
          sts_serial|   5|    100000|     100|0.51313468|  PASSED  
          sts_serial|   5|    100000|     100|0.95401085|  PASSED  
          sts_serial|   6|    100000|     100|0.70333163|  PASSED  
          sts_serial|   6|    100000|     100|0.66140438|  PASSED  
          sts_serial|   7|    100000|     100|0.27275735|  PASSED  
          sts_serial|   7|    100000|     100|0.73131183|  PASSED  
          sts_serial|   8|    100000|     100|0.79097395|  PASSED  
          sts_serial|   8|    100000|     100|0.11788110|  PASSED  
          sts_serial|   9|    100000|     100|0.11241350|  PASSED  
          sts_serial|   9|    100000|     100|0.46669305|  PASSED  
          sts_serial|  10|    100000|     100|0.36360727|  PASSED  
          sts_serial|  10|    100000|     100|0.21133860|  PASSED  
          sts_serial|  11|    100000|     100|0.87789137|  PASSED  
          sts_serial|  11|    100000|     100|0.62468125|  PASSED  
          sts_serial|  12|    100000|     100|0.86249055|  PASSED  
          sts_serial|  12|    100000|     100|0.57787321|  PASSED  
          sts_serial|  13|    100000|     100|0.67919228|  PASSED  
          sts_serial|  13|    100000|     100|0.66373553|  PASSED  
          sts_serial|  14|    100000|     100|0.20855130|  PASSED  
          sts_serial|  14|    100000|     100|0.18849764|  PASSED  
          sts_serial|  15|    100000|     100|0.89692043|  PASSED  
          sts_serial|  15|    100000|     100|0.53232607|  PASSED  
          sts_serial|  16|    100000|     100|0.77015070|  PASSED  
          sts_serial|  16|    100000|     100|0.60154555|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   1|    100000|     100|0.37952267|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   2|    100000|     100|0.36348895|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   3|    100000|     100|0.57968485|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   4|    100000|     100|0.17873936|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   5|    100000|     100|0.48558667|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   6|    100000|     100|0.74551801|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   7|    100000|     100|0.27090446|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   8|    100000|     100|0.53947545|  PASSED  
# The file file_input_raw was rewound 15 times
         rgb_bitdist|   9|    100000|     100|0.89998374|  PASSED  
# The file file_input_raw was rewound 16 times
         rgb_bitdist|  10|    100000|     100|0.00157431|   WEAK   
# The file file_input_raw was rewound 17 times
         rgb_bitdist|  11|    100000|     100|0.62019017|  PASSED  
# The file file_input_raw was rewound 18 times
         rgb_bitdist|  12|    100000|     100|0.92008350|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   2|     10000|    1000|0.87075989|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   3|     10000|    1000|0.13685966|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   4|     10000|    1000|0.89843339|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   5|     10000|    1000|0.16271961|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   2|    100000|     100|0.42954009|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   3|    100000|     100|0.90494498|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   4|    100000|     100|0.42137113|  PASSED  
# The file file_input_raw was rewound 19 times
    rgb_permutations|   5|    100000|     100|0.30273470|  PASSED  
# The file file_input_raw was rewound 19 times
      rgb_lagged_sum|   0|   1000000|     100|0.96194887|  PASSED  
# The file file_input_raw was rewound 20 times
      rgb_lagged_sum|   1|   1000000|     100|0.70980290|  PASSED  
# The file file_input_raw was rewound 21 times
      rgb_lagged_sum|   2|   1000000|     100|0.12280301|  PASSED  
# The file file_input_raw was rewound 22 times
      rgb_lagged_sum|   3|   1000000|     100|0.50215950|  PASSED  
# The file file_input_raw was rewound 24 times
      rgb_lagged_sum|   4|   1000000|     100|0.89934342|  PASSED  
# The file file_input_raw was rewound 26 times
      rgb_lagged_sum|   5|   1000000|     100|0.42379574|  PASSED  
# The file file_input_raw was rewound 29 times
      rgb_lagged_sum|   6|   1000000|     100|0.91204704|  PASSED  
# The file file_input_raw was rewound 32 times
      rgb_lagged_sum|   7|   1000000|     100|0.33600428|  PASSED  
# The file file_input_raw was rewound 35 times
      rgb_lagged_sum|   8|   1000000|     100|0.85004361|  PASSED  
# The file file_input_raw was rewound 39 times
      rgb_lagged_sum|   9|   1000000|     100|0.28553362|  PASSED  
# The file file_input_raw was rewound 43 times
      rgb_lagged_sum|  10|   1000000|     100|0.35547674|  PASSED  
# The file file_input_raw was rewound 48 times
      rgb_lagged_sum|  11|   1000000|     100|0.27686019|  PASSED  
# The file file_input_raw was rewound 53 times
      rgb_lagged_sum|  12|   1000000|     100|0.12776627|  PASSED  
# The file file_input_raw was rewound 58 times
      rgb_lagged_sum|  13|   1000000|     100|0.41511412|  PASSED  
# The file file_input_raw was rewound 63 times
      rgb_lagged_sum|  14|   1000000|     100|0.39162527|  PASSED  
# The file file_input_raw was rewound 69 times
      rgb_lagged_sum|  15|   1000000|     100|0.05912287|  PASSED  
# The file file_input_raw was rewound 76 times
      rgb_lagged_sum|  16|   1000000|     100|0.61383046|  PASSED  
# The file file_input_raw was rewound 82 times
      rgb_lagged_sum|  17|   1000000|     100|0.81000299|  PASSED  
# The file file_input_raw was rewound 89 times
      rgb_lagged_sum|  18|   1000000|     100|0.24556115|  PASSED  
# The file file_input_raw was rewound 97 times
      rgb_lagged_sum|  19|   1000000|     100|0.62274910|  PASSED  
# The file file_input_raw was rewound 105 times
      rgb_lagged_sum|  20|   1000000|     100|0.79903036|  PASSED  
# The file file_input_raw was rewound 113 times
      rgb_lagged_sum|  21|   1000000|     100|0.98924430|  PASSED  
# The file file_input_raw was rewound 121 times
      rgb_lagged_sum|  22|   1000000|     100|0.60235926|  PASSED  
# The file file_input_raw was rewound 130 times
      rgb_lagged_sum|  23|   1000000|     100|0.25016973|  PASSED  
# The file file_input_raw was rewound 140 times
      rgb_lagged_sum|  24|   1000000|     100|0.52005063|  PASSED  
# The file file_input_raw was rewound 149 times
      rgb_lagged_sum|  25|   1000000|     100|0.42229633|  PASSED  
# The file file_input_raw was rewound 159 times
      rgb_lagged_sum|  26|   1000000|     100|0.15089289|  PASSED  
# The file file_input_raw was rewound 170 times
      rgb_lagged_sum|  27|   1000000|     100|0.99748249|   WEAK   
# The file file_input_raw was rewound 181 times
      rgb_lagged_sum|  28|   1000000|     100|0.96602997|  PASSED  
# The file file_input_raw was rewound 192 times
      rgb_lagged_sum|  29|   1000000|     100|0.49913792|  PASSED  
# The file file_input_raw was rewound 203 times
      rgb_lagged_sum|  30|   1000000|     100|0.58200782|  PASSED  
# The file file_input_raw was rewound 215 times
      rgb_lagged_sum|  31|   1000000|     100|0.00001019|   WEAK   
# The file file_input_raw was rewound 228 times
      rgb_lagged_sum|  32|   1000000|     100|0.92018006|  PASSED  
# The file file_input_raw was rewound 228 times
     rgb_kstest_test|   0|     10000|    1000|0.79860094|  PASSED  
# The file file_input_raw was rewound 228 times
     dab_bytedistrib|   0|  51200000|       1|0.54423790|  PASSED  
# The file file_input_raw was rewound 228 times
             dab_dct| 256|     50000|       1|0.98923770|  PASSED  
Preparing to run test 207.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_filltree|  32|  15000000|       1|0.55389018|  PASSED  
        dab_filltree|  32|  15000000|       1|0.41088844|  PASSED  
Preparing to run test 208.  ntuple = 0
# The file file_input_raw was rewound 229 times
       dab_filltree2|   0|   5000000|       1|0.09622533|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.77657050|  PASSED  
Preparing to run test 209.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_monobit2|  12|  65000000|       1|0.52876497|  PASSED
```
