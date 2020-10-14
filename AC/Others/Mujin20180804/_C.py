N, M = map(int, input().split())
S = [ list(map(str, input().split())) for i in range(M)]

cnt = 0
# for i in range(N):
#     for j in range(M):
#         print("i, j: ", i," ", j)
#         if S[i][j] == "." \
#                 or ((i+1 < N) and (j-1 >= 0) and (S[i][j-1] == S[i+1][j-1] == ".")) \
#                 or ((i+1 < N) and (j+1 < M) and (S[i+1][j] == S[i+1][j+1] == ".")) \
#                 or ((i-1 >= 0) and (j+1 < M) and (S[i][j+1] == S[i-1][j+1] == ".")) \
#                 or ((i-1 >= 0) and (j-1 >= 0) and (S[i-1][j] == S[i-1][j-1] == ".")):
#                 cnt += 1
#

for i in range(N):
    for j in range(M):
        if S[i][j] == ".":
            while S[i][j+k] == "." or S[i][j+k] == None:

print(cnt)
