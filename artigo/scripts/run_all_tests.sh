#!/bin/bash

export PATH=$PATH:.

echo "merkled -duration 30"
merkled -duration 30s

echo "deviced -duration 30"
deviced -duration 30s

echo "acunitd -duration 30 -nodes 3"
acunitd -duration 30s -nodes 3

echo "acunitd -duration 30 -nodes 15"
acunitd -duration 30s -nodes 3

echo "acunitd -duration 30 -nodes 100"
acunitd -duration 30s -nodes 3