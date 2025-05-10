# Raspberry Pi 5 RNG Benchmark

BCM2712
107d208000.rng
Linux rpi5 6.6.74+rpt-rpi-2712 #1 SMP PREEMPT Debian 1:6.6.74-1+rpt1 (2025-01-27) aarch64 GNU/Linux

## Dieharder

```sh
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
 file_input_raw|                rpi5_go_fips.bin|  8.04e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.77112663|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.86666208|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.86521141|  PASSED  
# The file file_input_raw was rewound 1 times
    diehard_rank_6x8|   0|    100000|     100|0.14149581|  PASSED  
# The file file_input_raw was rewound 1 times
   diehard_bitstream|   0|   2097152|     100|0.81996065|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_opso|   0|   2097152|     100|0.30049032|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_oqso|   0|   2097152|     100|0.93711782|  PASSED  
# The file file_input_raw was rewound 2 times
         diehard_dna|   0|   2097152|     100|0.58603315|  PASSED  
# The file file_input_raw was rewound 2 times
diehard_count_1s_str|   0|    256000|     100|0.35053481|  PASSED  
# The file file_input_raw was rewound 3 times
diehard_count_1s_byt|   0|    256000|     100|0.30509808|  PASSED  
# The file file_input_raw was rewound 3 times
 diehard_parking_lot|   0|     12000|     100|0.56445279|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_2dsphere|   2|      8000|     100|0.56525735|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_3dsphere|   3|      4000|     100|0.00012586|   WEAK   
# The file file_input_raw was rewound 4 times
     diehard_squeeze|   0|    100000|     100|0.92786477|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_sums|   0|       100|     100|0.22332556|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_runs|   0|    100000|     100|0.63151161|  PASSED  
        diehard_runs|   0|    100000|     100|0.33529452|  PASSED  
# The file file_input_raw was rewound 4 times
       diehard_craps|   0|    200000|     100|0.35172165|  PASSED  
       diehard_craps|   0|    200000|     100|0.93521333|  PASSED  
# The file file_input_raw was rewound 12 times
 marsaglia_tsang_gcd|   0|  10000000|     100|0.01128387|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.25502881|  PASSED  
# The file file_input_raw was rewound 12 times
         sts_monobit|   1|    100000|     100|0.30774783|  PASSED  
# The file file_input_raw was rewound 12 times
            sts_runs|   2|    100000|     100|0.79306728|  PASSED  
# The file file_input_raw was rewound 12 times
          sts_serial|   1|    100000|     100|0.92377764|  PASSED  
          sts_serial|   2|    100000|     100|0.30975399|  PASSED  
          sts_serial|   3|    100000|     100|0.86139211|  PASSED  
          sts_serial|   3|    100000|     100|0.23866549|  PASSED  
          sts_serial|   4|    100000|     100|0.63229825|  PASSED  
          sts_serial|   4|    100000|     100|0.94136446|  PASSED  
          sts_serial|   5|    100000|     100|0.31590341|  PASSED  
          sts_serial|   5|    100000|     100|0.42545680|  PASSED  
          sts_serial|   6|    100000|     100|0.14425251|  PASSED  
          sts_serial|   6|    100000|     100|0.76652869|  PASSED  
          sts_serial|   7|    100000|     100|0.24348031|  PASSED  
          sts_serial|   7|    100000|     100|0.40314201|  PASSED  
          sts_serial|   8|    100000|     100|0.63751978|  PASSED  
          sts_serial|   8|    100000|     100|0.33309866|  PASSED  
          sts_serial|   9|    100000|     100|0.61139314|  PASSED  
          sts_serial|   9|    100000|     100|0.58413581|  PASSED  
          sts_serial|  10|    100000|     100|0.38094335|  PASSED  
          sts_serial|  10|    100000|     100|0.30132032|  PASSED  
          sts_serial|  11|    100000|     100|0.75718819|  PASSED  
          sts_serial|  11|    100000|     100|0.49036478|  PASSED  
          sts_serial|  12|    100000|     100|0.79625174|  PASSED  
          sts_serial|  12|    100000|     100|0.46325943|  PASSED  
          sts_serial|  13|    100000|     100|0.72489260|  PASSED  
          sts_serial|  13|    100000|     100|0.22795342|  PASSED  
          sts_serial|  14|    100000|     100|0.35242955|  PASSED  
          sts_serial|  14|    100000|     100|0.99717687|   WEAK   
          sts_serial|  15|    100000|     100|0.91808419|  PASSED  
          sts_serial|  15|    100000|     100|0.19709209|  PASSED  
          sts_serial|  16|    100000|     100|0.10030547|  PASSED  
          sts_serial|  16|    100000|     100|0.02656306|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   1|    100000|     100|0.53795212|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   2|    100000|     100|0.95226333|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   3|    100000|     100|0.88723433|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   4|    100000|     100|0.64045122|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   5|    100000|     100|0.03704820|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   6|    100000|     100|0.43165258|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   7|    100000|     100|0.98739707|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   8|    100000|     100|0.41534248|  PASSED  
# The file file_input_raw was rewound 15 times
         rgb_bitdist|   9|    100000|     100|0.91798418|  PASSED  
# The file file_input_raw was rewound 16 times
         rgb_bitdist|  10|    100000|     100|0.82321674|  PASSED  
# The file file_input_raw was rewound 17 times
         rgb_bitdist|  11|    100000|     100|0.99552373|   WEAK   
# The file file_input_raw was rewound 18 times
         rgb_bitdist|  12|    100000|     100|0.89429314|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   2|     10000|    1000|0.52999174|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   3|     10000|    1000|0.14454258|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   4|     10000|    1000|0.31676942|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   5|     10000|    1000|0.43566173|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   2|    100000|     100|0.15308100|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   3|    100000|     100|0.40267995|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   4|    100000|     100|0.76810515|  PASSED  
# The file file_input_raw was rewound 19 times
    rgb_permutations|   5|    100000|     100|0.99247483|  PASSED  
# The file file_input_raw was rewound 19 times
      rgb_lagged_sum|   0|   1000000|     100|0.88337698|  PASSED  
# The file file_input_raw was rewound 20 times
      rgb_lagged_sum|   1|   1000000|     100|0.82956449|  PASSED  
# The file file_input_raw was rewound 21 times
      rgb_lagged_sum|   2|   1000000|     100|0.86444730|  PASSED  
# The file file_input_raw was rewound 22 times
      rgb_lagged_sum|   3|   1000000|     100|0.56093865|  PASSED  
# The file file_input_raw was rewound 24 times
      rgb_lagged_sum|   4|   1000000|     100|0.38118599|  PASSED  
# The file file_input_raw was rewound 26 times
      rgb_lagged_sum|   5|   1000000|     100|0.60871403|  PASSED  
# The file file_input_raw was rewound 29 times
      rgb_lagged_sum|   6|   1000000|     100|0.55317871|  PASSED  
# The file file_input_raw was rewound 32 times
      rgb_lagged_sum|   7|   1000000|     100|0.28970886|  PASSED  
# The file file_input_raw was rewound 35 times
      rgb_lagged_sum|   8|   1000000|     100|0.16557280|  PASSED  
# The file file_input_raw was rewound 39 times
      rgb_lagged_sum|   9|   1000000|     100|0.36716281|  PASSED  
# The file file_input_raw was rewound 43 times
      rgb_lagged_sum|  10|   1000000|     100|0.98376144|  PASSED  
# The file file_input_raw was rewound 48 times
      rgb_lagged_sum|  11|   1000000|     100|0.44647920|  PASSED  
# The file file_input_raw was rewound 53 times
      rgb_lagged_sum|  12|   1000000|     100|0.98842860|  PASSED  
# The file file_input_raw was rewound 58 times
      rgb_lagged_sum|  13|   1000000|     100|0.69126968|  PASSED  
# The file file_input_raw was rewound 63 times
      rgb_lagged_sum|  14|   1000000|     100|0.56989113|  PASSED  
# The file file_input_raw was rewound 69 times
      rgb_lagged_sum|  15|   1000000|     100|0.65279942|  PASSED  
# The file file_input_raw was rewound 76 times
      rgb_lagged_sum|  16|   1000000|     100|0.26951164|  PASSED  
# The file file_input_raw was rewound 82 times
      rgb_lagged_sum|  17|   1000000|     100|0.34672812|  PASSED  
# The file file_input_raw was rewound 89 times
      rgb_lagged_sum|  18|   1000000|     100|0.09963983|  PASSED  
# The file file_input_raw was rewound 97 times
      rgb_lagged_sum|  19|   1000000|     100|0.51843003|  PASSED  
# The file file_input_raw was rewound 105 times
      rgb_lagged_sum|  20|   1000000|     100|0.91272268|  PASSED  
# The file file_input_raw was rewound 113 times
      rgb_lagged_sum|  21|   1000000|     100|0.86148492|  PASSED  
# The file file_input_raw was rewound 121 times
      rgb_lagged_sum|  22|   1000000|     100|0.87596553|  PASSED  
# The file file_input_raw was rewound 130 times
      rgb_lagged_sum|  23|   1000000|     100|0.29500762|  PASSED  
# The file file_input_raw was rewound 140 times
      rgb_lagged_sum|  24|   1000000|     100|0.49776731|  PASSED  
# The file file_input_raw was rewound 149 times
      rgb_lagged_sum|  25|   1000000|     100|0.99921597|   WEAK   
# The file file_input_raw was rewound 159 times
      rgb_lagged_sum|  26|   1000000|     100|0.68152451|  PASSED  
# The file file_input_raw was rewound 170 times
      rgb_lagged_sum|  27|   1000000|     100|0.94650744|  PASSED  
# The file file_input_raw was rewound 181 times
      rgb_lagged_sum|  28|   1000000|     100|0.62132320|  PASSED  
# The file file_input_raw was rewound 192 times
      rgb_lagged_sum|  29|   1000000|     100|0.99847803|   WEAK   
# The file file_input_raw was rewound 203 times
      rgb_lagged_sum|  30|   1000000|     100|0.69497119|  PASSED  
# The file file_input_raw was rewound 215 times
      rgb_lagged_sum|  31|   1000000|     100|0.00051177|   WEAK   
# The file file_input_raw was rewound 228 times
      rgb_lagged_sum|  32|   1000000|     100|0.38000035|  PASSED  
# The file file_input_raw was rewound 228 times
     rgb_kstest_test|   0|     10000|    1000|0.13814620|  PASSED  
# The file file_input_raw was rewound 228 times
     dab_bytedistrib|   0|  51200000|       1|0.27857782|  PASSED  
# The file file_input_raw was rewound 228 times
             dab_dct| 256|     50000|       1|0.28495887|  PASSED  
Preparing to run test 207.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_filltree|  32|  15000000|       1|0.81034777|  PASSED  
        dab_filltree|  32|  15000000|       1|0.00574365|  PASSED  
Preparing to run test 208.  ntuple = 0
# The file file_input_raw was rewound 229 times
       dab_filltree2|   0|   5000000|       1|0.84498101|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.91137004|  PASSED  
Preparing to run test 209.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_monobit2|  12|  65000000|       1|0.58840568|  PASSED
```
