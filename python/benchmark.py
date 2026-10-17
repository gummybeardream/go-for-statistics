import timeit
import tracemalloc
import pandas as pd
import statsmodels.api as sm

datasets = {
    "Dataset I": {
        "x": [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        "y": [8.04, 6.95, 7.58, 8.81, 8.33, 9.96,
              7.24, 4.26, 10.84, 4.82, 5.68],
    },
    "Dataset II": {
        "x": [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        "y": [9.14, 8.14, 8.74, 8.77, 9.26, 8.10,
              6.13, 3.10, 9.13, 7.26, 4.74],
    },
    "Dataset III": {
        "x": [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        "y": [7.46, 6.77, 12.74, 7.11, 7.81, 8.84,
              6.08, 5.39, 8.15, 6.42, 5.73],
    },
    "Dataset IV": {
        "x": [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
        "y": [6.58, 5.76, 7.71, 8.84, 8.47, 7.04,
              5.25, 12.50, 5.56, 7.91, 6.89],
    },
}

def calculate_regression(x, y):
    design_matrix = sm.add_constant(x)
    model = sm.OLS(y, design_matrix)
    result = model.fit()

    return result.params


def benchmark_time(operation, repetitions=10000):
    elapsed = timeit.timeit(operation, number=repetitions)
    return elapsed / repetitions * 1e9


def benchmark_memory(operation):
    tracemalloc.start()

    result = operation()
    current, peak = tracemalloc.get_traced_memory()

    tracemalloc.stop()
    return peak

for name, data in datasets.items():
    x_values = data["x"]
    y_values = data["y"]

    # Define the regression operation to repeat
    operation = lambda: calculate_regression(x_values, y_values)

    # Benchmarks
    execution_time = benchmark_time(operation)
    peak_memory = benchmark_memory(operation)

    print(f"{name}:")
    print(f"  Execution time: {execution_time:.2f} ns/op")
    print(f"  Peak traced memory: {peak_memory} bytes")