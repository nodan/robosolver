#!/bin/bash
echo
echo Real Game position
`dirname $0`/../robosolver 3 \
    "09 01 01 03 09 01 01 01 01 03 09 01 01 01 01 03\
     08 00 00 00 00 02 0c 00 00 00 00 00 00 02 09 02\
     08 04 00 00 00 00 01 00 00 00 02 0c 00 00 00 02\
     08 03 08 00 00 04 00 00 00 00 00 01 00 00 00 06\
     08 00 00 00 02 09 00 00 00 00 00 00 00 00 00 03\
     08 00 06 08 00 00 00 06 08 00 00 00 00 00 00 02\
     0c 00 01 00 00 00 00 05 04 00 04 00 00 06 08 02\
     09 00 00 00 00 00 02 0f 0f 08 03 08 00 01 00 02\
     08 00 00 00 00 00 02 0f 0f 08 00 00 04 00 00 02\
     08 04 00 02 0c 00 00 01 01 00 00 00 03 08 00 06\
     08 03 08 00 01 00 00 00 02 0c 00 00 00 00 00 03\
     0c 00 00 00 00 00 00 00 00 01 00 00 00 00 00 02\
     09 00 00 00 00 00 04 00 00 00 20 30 10 40 00 02\
     08 00 00 00 00 02 09 00 00 04 00 00 00 00 06 0a\
     08 00 06 08 00 00 00 00 02 09 00 00 00 00 01 02\
     0c 04 05 06 0c 04 04 04 04 04 06 0c 04 04 04 06" solve 22
