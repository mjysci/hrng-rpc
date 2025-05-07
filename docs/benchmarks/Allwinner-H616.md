# Allwinner H616 RNG Benchmark

## Dieharder

```sh
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
 file_input_raw|                         rng.bin|  9.20e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.84389552|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.39747667|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.36119684|  PASSED  
# The file file_input_raw was rewound 1 times
    diehard_rank_6x8|   0|    100000|     100|0.95175676|  PASSED  
# The file file_input_raw was rewound 1 times
   diehard_bitstream|   0|   2097152|     100|0.96957935|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_opso|   0|   2097152|     100|0.84528219|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_oqso|   0|   2097152|     100|0.81773673|  PASSED  
# The file file_input_raw was rewound 2 times
         diehard_dna|   0|   2097152|     100|0.24487408|  PASSED  
# The file file_input_raw was rewound 2 times
diehard_count_1s_str|   0|    256000|     100|0.23698121|  PASSED  
# The file file_input_raw was rewound 3 times
diehard_count_1s_byt|   0|    256000|     100|0.65943204|  PASSED  
# The file file_input_raw was rewound 3 times
 diehard_parking_lot|   0|     12000|     100|0.90061093|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_2dsphere|   2|      8000|     100|0.55054135|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_3dsphere|   3|      4000|     100|0.45603817|  PASSED  
# The file file_input_raw was rewound 4 times
     diehard_squeeze|   0|    100000|     100|0.34812193|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_sums|   0|       100|     100|0.00386260|   WEAK   
# The file file_input_raw was rewound 4 times
        diehard_runs|   0|    100000|     100|0.96606116|  PASSED  
        diehard_runs|   0|    100000|     100|0.91901357|  PASSED  
# The file file_input_raw was rewound 4 times
       diehard_craps|   0|    200000|     100|0.57794951|  PASSED  
       diehard_craps|   0|    200000|     100|0.44465465|  PASSED  
# The file file_input_raw was rewound 12 times
 marsaglia_tsang_gcd|   0|  10000000|     100|0.07553628|  PASSED  
 marsaglia_tsang_gcd|   0|  10000000|     100|0.00204518|   WEAK   
# The file file_input_raw was rewound 12 times
         sts_monobit|   1|    100000|     100|0.71434176|  PASSED  
# The file file_input_raw was rewound 12 times
            sts_runs|   2|    100000|     100|0.62762646|  PASSED  
# The file file_input_raw was rewound 12 times
          sts_serial|   1|    100000|     100|0.69940398|  PASSED  
          sts_serial|   2|    100000|     100|0.97942889|  PASSED  
          sts_serial|   3|    100000|     100|0.20229275|  PASSED  
          sts_serial|   3|    100000|     100|0.00363792|   WEAK   
          sts_serial|   4|    100000|     100|0.18388447|  PASSED  
          sts_serial|   4|    100000|     100|0.15291775|  PASSED  
          sts_serial|   5|    100000|     100|0.18942580|  PASSED  
          sts_serial|   5|    100000|     100|0.50893137|  PASSED  
          sts_serial|   6|    100000|     100|0.28946439|  PASSED  
          sts_serial|   6|    100000|     100|0.39200866|  PASSED  
          sts_serial|   7|    100000|     100|0.23380716|  PASSED  
          sts_serial|   7|    100000|     100|0.90009471|  PASSED  
          sts_serial|   8|    100000|     100|0.02707584|  PASSED  
          sts_serial|   8|    100000|     100|0.20354853|  PASSED  
          sts_serial|   9|    100000|     100|0.56382755|  PASSED  
          sts_serial|   9|    100000|     100|0.48235175|  PASSED  
          sts_serial|  10|    100000|     100|0.78098429|  PASSED  
          sts_serial|  10|    100000|     100|0.47373680|  PASSED  
          sts_serial|  11|    100000|     100|0.04298893|  PASSED  
          sts_serial|  11|    100000|     100|0.05011077|  PASSED  
          sts_serial|  12|    100000|     100|0.21482094|  PASSED  
          sts_serial|  12|    100000|     100|0.79997306|  PASSED  
          sts_serial|  13|    100000|     100|0.59219735|  PASSED  
          sts_serial|  13|    100000|     100|0.35543372|  PASSED  
          sts_serial|  14|    100000|     100|0.87842513|  PASSED  
          sts_serial|  14|    100000|     100|0.77399010|  PASSED  
          sts_serial|  15|    100000|     100|0.42930606|  PASSED  
          sts_serial|  15|    100000|     100|0.17411180|  PASSED  
          sts_serial|  16|    100000|     100|0.33352442|  PASSED  
          sts_serial|  16|    100000|     100|0.88300752|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   1|    100000|     100|0.05929519|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   2|    100000|     100|0.91609320|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   3|    100000|     100|0.72411031|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   4|    100000|     100|0.89190405|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   5|    100000|     100|0.37458517|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   6|    100000|     100|0.88408324|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   7|    100000|     100|0.80307400|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   8|    100000|     100|0.20870907|  PASSED  
# The file file_input_raw was rewound 15 times
         rgb_bitdist|   9|    100000|     100|0.80121152|  PASSED  
# The file file_input_raw was rewound 16 times
         rgb_bitdist|  10|    100000|     100|0.24051451|  PASSED  
# The file file_input_raw was rewound 17 times
         rgb_bitdist|  11|    100000|     100|0.80948902|  PASSED  
# The file file_input_raw was rewound 18 times
         rgb_bitdist|  12|    100000|     100|0.52561945|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   2|     10000|    1000|0.30721618|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   3|     10000|    1000|0.90874837|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   4|     10000|    1000|0.78909969|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   5|     10000|    1000|0.72965687|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   2|    100000|     100|0.95525853|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   3|    100000|     100|0.66237713|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   4|    100000|     100|0.28638195|  PASSED  
# The file file_input_raw was rewound 19 times
    rgb_permutations|   5|    100000|     100|0.43057274|  PASSED  
# The file file_input_raw was rewound 19 times
      rgb_lagged_sum|   0|   1000000|     100|0.46553301|  PASSED  
# The file file_input_raw was rewound 20 times
      rgb_lagged_sum|   1|   1000000|     100|0.28992561|  PASSED  
# The file file_input_raw was rewound 21 times
      rgb_lagged_sum|   2|   1000000|     100|0.80788830|  PASSED  
# The file file_input_raw was rewound 22 times
      rgb_lagged_sum|   3|   1000000|     100|0.15327483|  PASSED  
# The file file_input_raw was rewound 24 times
      rgb_lagged_sum|   4|   1000000|     100|0.62730010|  PASSED  
# The file file_input_raw was rewound 26 times
      rgb_lagged_sum|   5|   1000000|     100|0.47753514|  PASSED  
# The file file_input_raw was rewound 29 times
      rgb_lagged_sum|   6|   1000000|     100|0.47582042|  PASSED  
# The file file_input_raw was rewound 32 times
      rgb_lagged_sum|   7|   1000000|     100|0.04250313|  PASSED  
# The file file_input_raw was rewound 35 times
      rgb_lagged_sum|   8|   1000000|     100|0.29269601|  PASSED  
# The file file_input_raw was rewound 39 times
      rgb_lagged_sum|   9|   1000000|     100|0.85111237|  PASSED  
# The file file_input_raw was rewound 43 times
      rgb_lagged_sum|  10|   1000000|     100|0.63124063|  PASSED  
# The file file_input_raw was rewound 48 times
      rgb_lagged_sum|  11|   1000000|     100|0.08342916|  PASSED  
# The file file_input_raw was rewound 53 times
      rgb_lagged_sum|  12|   1000000|     100|0.59300968|  PASSED  
# The file file_input_raw was rewound 58 times
      rgb_lagged_sum|  13|   1000000|     100|0.36864645|  PASSED  
# The file file_input_raw was rewound 63 times
      rgb_lagged_sum|  14|   1000000|     100|0.74153065|  PASSED  
# The file file_input_raw was rewound 69 times
      rgb_lagged_sum|  15|   1000000|     100|0.02475853|  PASSED  
# The file file_input_raw was rewound 76 times
      rgb_lagged_sum|  16|   1000000|     100|0.65621950|  PASSED  
# The file file_input_raw was rewound 82 times
      rgb_lagged_sum|  17|   1000000|     100|0.70576541|  PASSED  
# The file file_input_raw was rewound 89 times
      rgb_lagged_sum|  18|   1000000|     100|0.53943311|  PASSED  
# The file file_input_raw was rewound 97 times
      rgb_lagged_sum|  19|   1000000|     100|0.22073956|  PASSED  
# The file file_input_raw was rewound 105 times
      rgb_lagged_sum|  20|   1000000|     100|0.81377136|  PASSED  
# The file file_input_raw was rewound 113 times
      rgb_lagged_sum|  21|   1000000|     100|0.99603186|   WEAK   
# The file file_input_raw was rewound 121 times
      rgb_lagged_sum|  22|   1000000|     100|0.75631432|  PASSED  
# The file file_input_raw was rewound 130 times
      rgb_lagged_sum|  23|   1000000|     100|0.25023594|  PASSED  
# The file file_input_raw was rewound 140 times
      rgb_lagged_sum|  24|   1000000|     100|0.18235978|  PASSED  
# The file file_input_raw was rewound 149 times
      rgb_lagged_sum|  25|   1000000|     100|0.98170703|  PASSED  
# The file file_input_raw was rewound 159 times
      rgb_lagged_sum|  26|   1000000|     100|0.84753932|  PASSED  
# The file file_input_raw was rewound 170 times
      rgb_lagged_sum|  27|   1000000|     100|0.14395018|  PASSED  
# The file file_input_raw was rewound 181 times
      rgb_lagged_sum|  28|   1000000|     100|0.88831952|  PASSED  
# The file file_input_raw was rewound 192 times
      rgb_lagged_sum|  29|   1000000|     100|0.54833237|  PASSED  
# The file file_input_raw was rewound 203 times
      rgb_lagged_sum|  30|   1000000|     100|0.94940933|  PASSED  
# The file file_input_raw was rewound 215 times
      rgb_lagged_sum|  31|   1000000|     100|0.00284297|   WEAK   
# The file file_input_raw was rewound 228 times
      rgb_lagged_sum|  32|   1000000|     100|0.38925254|  PASSED  
# The file file_input_raw was rewound 228 times
     rgb_kstest_test|   0|     10000|    1000|0.06763093|  PASSED  
# The file file_input_raw was rewound 228 times
     dab_bytedistrib|   0|  51200000|       1|0.24293440|  PASSED  
# The file file_input_raw was rewound 228 times
             dab_dct| 256|     50000|       1|0.03671553|  PASSED  
Preparing to run test 207.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_filltree|  32|  15000000|       1|0.60803671|  PASSED  
        dab_filltree|  32|  15000000|       1|0.17364061|  PASSED  
Preparing to run test 208.  ntuple = 0
# The file file_input_raw was rewound 229 times
       dab_filltree2|   0|   5000000|       1|0.81937089|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.59227235|  PASSED  
Preparing to run test 209.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_monobit2|  12|  65000000|       1|0.91737310|  PASSED
```
