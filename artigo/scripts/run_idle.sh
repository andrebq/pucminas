#!/bin/bash

echo "idle - $(date)"
sleep 30

echo "primes - $(date)"
sysbench --test=cpu --cpu-max-prime=20000 run &
sys_pid=$!

sleep 30
kill $sys_pid
echo "done - $(date)"