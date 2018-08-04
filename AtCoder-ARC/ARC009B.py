rule = list(map(str, input().split()))
N = int(input())
arr = [ str(input()) for _ in range(N)]
cvn = [ str("") for _ in range(N)]

for i in range(N):
    print("arr[i]: ", arr[i])
    print("len(arr[i]): ", len(arr[i]))
    for j in range(len(arr[i])):
        for k in range(10):
            cvn[i][j] = arr[i][j].replace(rule[k], str(k))

print(cvn)
# sorted = cvn.sort()
# for i in range(N):
#     print(arr[cvn[i]])




# max_len = 0
#
# for i in range(N):
#     if len(str(arr[i])) > max_len:
#         max_len = len(str(arr[0]))
#         max_index = 1
#         max_value = arr[0]
#     elif len(str(i)) == max_len:
#         for i in range(max_len):
#             if rule.find(i) > rule.find(max_value[i])
#             i.find()
#
#     max_len =
#
#
#     sorted = arr.pop
