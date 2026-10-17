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


# 2. Regression function to benchmark
def calculate_regression(x, y):
    # TODO: Use sm.add_constant() to include the intercept
    # TODO: Construct and fit the sm.OLS() regression model
    # TODO: Return the fitted slope and intercept
    pass


# 3. Execution-time benchmark
def benchmark_time(operation, repetitions=10000):
    elapsed = timeit.timeit(operation, number=repetitions)
    return elapsed / repetitions * 1e9  # nanoseconds per operation


# 4. Python-tracked memory benchmark
def benchmark_memory(operation):
    tracemalloc.start()

    # TODO: Execute operation and measure traced peak memory
    # Use tracemalloc.get_traced_memory()

    tracemalloc.stop()

    # TODO: Return peak memory in bytes


# 5. Table-driven benchmark
for name, dataset in datasets.items():
    x = pd.Series(dataset["x"])
    y = pd.Series(dataset["y"])

    # TODO: Define an operation that runs calculate_regression()
    # TODO: Pass operation to benchmark_time()
    # TODO: Pass operation to benchmark_memory()
    # TODO: Print results for this dataset