n = input();
j = map(int,raw_input().split(" "));
j.reverse();
for i in range(0,n-1):
    print j[i],
print j[n-1]
