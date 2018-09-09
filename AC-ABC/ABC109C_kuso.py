N, X = map(int, input().split())
arr = list(map(int, input().split()))
arr_sort = sorted(arr)

ans = -1
previous = -1

for ele in arr_sort:
    if previous < 0:
        previous = ele
    else:
        elediff = abs(ele - previous)
        if ans < 0 or elediff < ans:
            ans = elediff
        previous = ele

    xdiff = abs(X - ele)
    if ans < 0 or xdiff < ans:
        ans = xdiff


print(ans)
