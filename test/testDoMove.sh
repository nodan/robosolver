#!/bin/bash
echo
echo Simple move on three by three:
`dirname $0`/../robosolver 3 "19 01 03 08 00 02 0c 04 06" move 0 2

echo Illegal move:
`dirname $0`/../robosolver 3 "19 01 03 08 00 02 0c 04 06" move 0 1
