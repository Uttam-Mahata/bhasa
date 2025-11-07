#!/bin/bash

# Script to run all Bhasa example programs

echo "======================================"
echo "ভাষা (Bhasa) - Running All Examples"
echo "======================================"
echo ""

examples=(
    # Basic Examples
    "hello"
    "variables"
    "bengali_variable_names"
    
    # Functions
    "functions"
    
    # Control Flow
    "conditionals"
    "logical_operators"
    "loops"
    "for_loops"
    "break_continue"
    
    # Data Structures
    "arrays"
    "array_methods"
    "array_advanced"
    "hash"
    
    # String Operations
    "string_methods"
    
    # Math Operations
    "math_functions"
    "bitwise_operations"
    "bitwise_comprehensive"
    
    # Data Types
    "datatypes_and_typecasting"
    "datatypes_edge_cases"
    
    # File I/O
    "file_io"
    
    # Advanced Examples
    "fibonacci"
    "comprehensive"
    "complete_bengali_test"
    "feature_test"
    "new_features_showcase"
    "json_support"
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

