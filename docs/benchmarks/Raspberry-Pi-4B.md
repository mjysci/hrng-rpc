# Raspberry Pi 4B RNG Benchmark

BCM2711
fe104000.rng
Linux rpi4b 6.12.20+rpt-rpi-v8 #1 SMP PREEMPT Debian 1:6.12.20-1+rpt1~bpo12+1 (2025-03-19) aarch64 GNU/Linux

## Dieharder

```sh
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
   rng_name    |           filename             |rands/second|
 file_input_raw|                         rng.bin|  9.18e+07  |
#=============================================================================#
        test_name   |ntup| tsamples |psamples|  p-value |Assessment
#=============================================================================#
   diehard_birthdays|   0|       100|     100|0.64491002|  PASSED  
      diehard_operm5|   0|   1000000|     100|0.46977181|  PASSED  
  diehard_rank_32x32|   0|     40000|     100|0.96437220|  PASSED  
# The file file_input_raw was rewound 1 times
    diehard_rank_6x8|   0|    100000|     100|0.41620653|  PASSED  
# The file file_input_raw was rewound 1 times
   diehard_bitstream|   0|   2097152|     100|0.55585544|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_opso|   0|   2097152|     100|0.56395299|  PASSED  
# The file file_input_raw was rewound 2 times
        diehard_oqso|   0|   2097152|     100|0.81829061|  PASSED  
# The file file_input_raw was rewound 2 times
         diehard_dna|   0|   2097152|     100|0.38860045|  PASSED  
# The file file_input_raw was rewound 2 times
diehard_count_1s_str|   0|    256000|     100|0.78648049|  PASSED  
# The file file_input_raw was rewound 3 times
diehard_count_1s_byt|   0|    256000|     100|0.61440805|  PASSED  
# The file file_input_raw was rewound 3 times
 diehard_parking_lot|   0|     12000|     100|0.78081639|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_2dsphere|   2|      8000|     100|0.96831361|  PASSED  
# The file file_input_raw was rewound 3 times
    diehard_3dsphere|   3|      4000|     100|0.52830336|  PASSED  
# The file file_input_raw was rewound 4 times
     diehard_squeeze|   0|    100000|     100|0.79085217|  PASSED  
# The file file_input_raw was rewound 4 times
        diehard_sums|   0|       100|     100|0.00461225|   WEAK   
# The file file_input_raw was rewound 4 times
        diehard_runs|   0|    100000|     100|0.03064652|  PASSED  
        diehard_runs|   0|    100000|     100|0.39092543|  PASSED  
# The file file_input_raw was rewound 4 times
       diehard_craps|   0|    200000|     100|0.88976259|  PASSED  
       diehard_craps|   0|    200000|     100|0.36022435|  PASSED  
# The file file_input_raw was rewound 12 times
 marsaglia_tsang_gcd|   0|  10000000|     100|0.00036706|   WEAK   
 marsaglia_tsang_gcd|   0|  10000000|     100|0.34515683|  PASSED  
# The file file_input_raw was rewound 12 times
         sts_monobit|   1|    100000|     100|0.03883630|  PASSED  
# The file file_input_raw was rewound 12 times
            sts_runs|   2|    100000|     100|0.08436889|  PASSED  
# The file file_input_raw was rewound 12 times
          sts_serial|   1|    100000|     100|0.63133175|  PASSED  
          sts_serial|   2|    100000|     100|0.95614812|  PASSED  
          sts_serial|   3|    100000|     100|0.28453260|  PASSED  
          sts_serial|   3|    100000|     100|0.67161160|  PASSED  
          sts_serial|   4|    100000|     100|0.28518525|  PASSED  
          sts_serial|   4|    100000|     100|0.57274493|  PASSED  
          sts_serial|   5|    100000|     100|0.76925271|  PASSED  
          sts_serial|   5|    100000|     100|0.01585859|  PASSED  
          sts_serial|   6|    100000|     100|0.83011405|  PASSED  
          sts_serial|   6|    100000|     100|0.65168595|  PASSED  
          sts_serial|   7|    100000|     100|0.94922036|  PASSED  
          sts_serial|   7|    100000|     100|0.90645599|  PASSED  
          sts_serial|   8|    100000|     100|0.71807996|  PASSED  
          sts_serial|   8|    100000|     100|0.79855607|  PASSED  
          sts_serial|   9|    100000|     100|0.59686321|  PASSED  
          sts_serial|   9|    100000|     100|0.87398483|  PASSED  
          sts_serial|  10|    100000|     100|0.93530814|  PASSED  
          sts_serial|  10|    100000|     100|0.95890065|  PASSED  
          sts_serial|  11|    100000|     100|0.80162050|  PASSED  
          sts_serial|  11|    100000|     100|0.30550570|  PASSED  
          sts_serial|  12|    100000|     100|0.92491030|  PASSED  
          sts_serial|  12|    100000|     100|0.90239925|  PASSED  
          sts_serial|  13|    100000|     100|0.90943607|  PASSED  
          sts_serial|  13|    100000|     100|0.95010329|  PASSED  
          sts_serial|  14|    100000|     100|0.18361203|  PASSED  
          sts_serial|  14|    100000|     100|0.37502241|  PASSED  
          sts_serial|  15|    100000|     100|0.11018377|  PASSED  
          sts_serial|  15|    100000|     100|0.40903652|  PASSED  
          sts_serial|  16|    100000|     100|0.85477525|  PASSED  
          sts_serial|  16|    100000|     100|0.31678463|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   1|    100000|     100|0.31769367|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   2|    100000|     100|0.36299274|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   3|    100000|     100|0.07191572|  PASSED  
# The file file_input_raw was rewound 12 times
         rgb_bitdist|   4|    100000|     100|0.60442249|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   5|    100000|     100|0.99352573|  PASSED  
# The file file_input_raw was rewound 13 times
         rgb_bitdist|   6|    100000|     100|0.94492375|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   7|    100000|     100|0.67349058|  PASSED  
# The file file_input_raw was rewound 14 times
         rgb_bitdist|   8|    100000|     100|0.18321946|  PASSED  
# The file file_input_raw was rewound 15 times
         rgb_bitdist|   9|    100000|     100|0.99285919|  PASSED  
# The file file_input_raw was rewound 16 times
         rgb_bitdist|  10|    100000|     100|0.94862793|  PASSED  
# The file file_input_raw was rewound 17 times
         rgb_bitdist|  11|    100000|     100|0.28720161|  PASSED  
# The file file_input_raw was rewound 18 times
         rgb_bitdist|  12|    100000|     100|0.90925733|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   2|     10000|    1000|0.13839729|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   3|     10000|    1000|0.85134028|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   4|     10000|    1000|0.87866185|  PASSED  
# The file file_input_raw was rewound 18 times
rgb_minimum_distance|   5|     10000|    1000|0.13642347|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   2|    100000|     100|0.08603907|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   3|    100000|     100|0.99344756|  PASSED  
# The file file_input_raw was rewound 18 times
    rgb_permutations|   4|    100000|     100|0.64069312|  PASSED  
# The file file_input_raw was rewound 19 times
    rgb_permutations|   5|    100000|     100|0.63853523|  PASSED  
# The file file_input_raw was rewound 19 times
      rgb_lagged_sum|   0|   1000000|     100|0.84161336|  PASSED  
# The file file_input_raw was rewound 20 times
      rgb_lagged_sum|   1|   1000000|     100|0.47363273|  PASSED  
# The file file_input_raw was rewound 21 times
      rgb_lagged_sum|   2|   1000000|     100|0.42452975|  PASSED  
# The file file_input_raw was rewound 22 times
      rgb_lagged_sum|   3|   1000000|     100|0.10592482|  PASSED  
# The file file_input_raw was rewound 24 times
      rgb_lagged_sum|   4|   1000000|     100|0.92473703|  PASSED  
# The file file_input_raw was rewound 26 times
      rgb_lagged_sum|   5|   1000000|     100|0.03227402|  PASSED  
# The file file_input_raw was rewound 29 times
      rgb_lagged_sum|   6|   1000000|     100|0.99635070|   WEAK   
# The file file_input_raw was rewound 32 times
      rgb_lagged_sum|   7|   1000000|     100|0.01957778|  PASSED  
# The file file_input_raw was rewound 35 times
      rgb_lagged_sum|   8|   1000000|     100|0.99685092|   WEAK   
# The file file_input_raw was rewound 39 times
      rgb_lagged_sum|   9|   1000000|     100|0.63253130|  PASSED  
# The file file_input_raw was rewound 43 times
      rgb_lagged_sum|  10|   1000000|     100|0.84950774|  PASSED  
# The file file_input_raw was rewound 48 times
      rgb_lagged_sum|  11|   1000000|     100|0.28653388|  PASSED  
# The file file_input_raw was rewound 53 times
      rgb_lagged_sum|  12|   1000000|     100|0.70650196|  PASSED  
# The file file_input_raw was rewound 58 times
      rgb_lagged_sum|  13|   1000000|     100|0.88854931|  PASSED  
# The file file_input_raw was rewound 63 times
      rgb_lagged_sum|  14|   1000000|     100|0.08523196|  PASSED  
# The file file_input_raw was rewound 69 times
      rgb_lagged_sum|  15|   1000000|     100|0.97657304|  PASSED  
# The file file_input_raw was rewound 76 times
      rgb_lagged_sum|  16|   1000000|     100|0.91485943|  PASSED  
# The file file_input_raw was rewound 82 times
      rgb_lagged_sum|  17|   1000000|     100|0.30636262|  PASSED  
# The file file_input_raw was rewound 89 times
      rgb_lagged_sum|  18|   1000000|     100|0.51889774|  PASSED  
# The file file_input_raw was rewound 97 times
      rgb_lagged_sum|  19|   1000000|     100|0.47643844|  PASSED  
# The file file_input_raw was rewound 105 times
      rgb_lagged_sum|  20|   1000000|     100|0.30312093|  PASSED  
# The file file_input_raw was rewound 113 times
      rgb_lagged_sum|  21|   1000000|     100|0.13112494|  PASSED  
# The file file_input_raw was rewound 121 times
      rgb_lagged_sum|  22|   1000000|     100|0.90612385|  PASSED  
# The file file_input_raw was rewound 130 times
      rgb_lagged_sum|  23|   1000000|     100|0.18835080|  PASSED  
# The file file_input_raw was rewound 140 times
      rgb_lagged_sum|  24|   1000000|     100|0.19506512|  PASSED  
# The file file_input_raw was rewound 149 times
      rgb_lagged_sum|  25|   1000000|     100|0.26749981|  PASSED  
# The file file_input_raw was rewound 159 times
      rgb_lagged_sum|  26|   1000000|     100|0.52302788|  PASSED  
# The file file_input_raw was rewound 170 times
      rgb_lagged_sum|  27|   1000000|     100|0.54127962|  PASSED  
# The file file_input_raw was rewound 181 times
      rgb_lagged_sum|  28|   1000000|     100|0.19573357|  PASSED  
# The file file_input_raw was rewound 192 times
      rgb_lagged_sum|  29|   1000000|     100|0.98754224|  PASSED  
# The file file_input_raw was rewound 203 times
      rgb_lagged_sum|  30|   1000000|     100|0.86472441|  PASSED  
# The file file_input_raw was rewound 215 times
      rgb_lagged_sum|  31|   1000000|     100|0.04638593|  PASSED  
# The file file_input_raw was rewound 228 times
      rgb_lagged_sum|  32|   1000000|     100|0.39903131|  PASSED  
# The file file_input_raw was rewound 228 times
     rgb_kstest_test|   0|     10000|    1000|0.01040127|  PASSED  
# The file file_input_raw was rewound 228 times
     dab_bytedistrib|   0|  51200000|       1|0.72187562|  PASSED  
# The file file_input_raw was rewound 228 times
             dab_dct| 256|     50000|       1|0.64584835|  PASSED  
Preparing to run test 207.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_filltree|  32|  15000000|       1|0.30855313|  PASSED  
        dab_filltree|  32|  15000000|       1|0.11910032|  PASSED  
Preparing to run test 208.  ntuple = 0
# The file file_input_raw was rewound 229 times
       dab_filltree2|   0|   5000000|       1|0.63176795|  PASSED  
       dab_filltree2|   1|   5000000|       1|0.86522071|  PASSED  
Preparing to run test 209.  ntuple = 0
# The file file_input_raw was rewound 229 times
        dab_monobit2|  12|  65000000|       1|0.41551716|  PASSED
```
