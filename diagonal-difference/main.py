#!/bin/python3
import os


def diagonal_difference(arr):
    lhs = 0
    for i, r in enumerate(arr):
        for j, c in enumerate(arr[i]):
            if i == j:
                lhs += arr[i][j]

    for i, _ in enumerate(arr):
        arr[i] = list(reversed(arr[i]))

    rhs = 0
    for i, r in enumerate(arr):
        for j, c in enumerate(arr[i]):
            if i == j:
                rhs += arr[i][j]
    return abs(rhs-lhs)


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    arr = []

    for _ in range(n):
        arr.append(list(map(int, input().rstrip().split())))

    result = diagonal_difference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()