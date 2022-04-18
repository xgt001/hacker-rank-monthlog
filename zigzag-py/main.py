# this challenge is pretty weird and doesn't accept 3 modified lines
# despite same output
def findZigZagSequence(a, n):
    a.sort()
    mid = int((n - 1) / 2)
    # a[mid], a[n - 1] = a[n - 1], a[mid]
    st = mid
    ed = n - 1
    print(a)
    while st <= ed and a[st] < a[ed]:
        a[st], a[ed] = a[ed], a[st]
        st = st + 1
        ed = ed - 1

    for i in range(n):
        if i == n - 1:
            print(a[i])
        else:
            print(a[i], end=' ')
    return


if __name__ == "__main__":
    n = 7
    a = [1, 2, 3, 4, 5, 6, 7]
    findZigZagSequence(a, n)
