#n = int(input())
#s = input()
#cnt = 0
# for i in range(0, n-2):
#    if s[i:i+3] == "ABC":
#        cnt += 1
# print(cnt)

n = int(input())
s = input()
cnt = 0
for i in range(0, len(s)-2):
    if s[i] == "A" and s[i+1] == "B"and s[i+2] == "C":
        cnt += 1
print(cnt)
