#!/bin/bash

export PATH=$PATH:.

echo "merkled -duration 30"
merkled.arm -duration 30s

echo "deviced -duration 30"
deviced.arm -duration 30s

echo "acunitd -duration 30 -nodes 3"
acunitd.arm -duration 30s -nodes 3

echo "acunitd -duration 30 -nodes 15"
acunitd.arm -duration 30s -nodes 3

echo "acunitd -duration 30 -nodes 100"
acunitd.arm -duration 30s -nodes 3