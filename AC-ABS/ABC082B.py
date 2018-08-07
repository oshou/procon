s = str(input())
s_asc = ''.join(sorted(s))

t = str(input())
t_desc = ''.join(sorted(t, reverse=True))

if t_desc > s_asc:
    print("Yes")
else:
    print("No")
