#!/bin/bash

# Script to run all Bhasa example programs

echo "======================================"
echo "ভাষা (Bhasa) - Running All Examples"
echo "======================================"
echo ""

examples=(
    "hello"
    "variables"
    "functions"
    "conditionals"
    "loops"
    "arrays"
    "hash"
    "fibonacci"
    "comprehensive"
    "for_loops"
    "break_continue"
    "logical_operators"
    "string_methods"
    "math_functions"
)

for example in "${examples[@]}"; do
    echo "--------------------------------------"
    echo "Running: examples/${example}.bhasa"
    echo "--------------------------------------"
    ./bhasa "examples/${example}.bhasa"
    exit_code=$?
    if [ $exit_code -ne 0 ]; then
        echo "ERROR: Failed to run ${example}.bhasa (exit code: $exit_code)"
        exit 1
    fi
    echo ""
done

echo "======================================"
echo "All examples completed successfully!"
echo "======================================"

