xa, ya, xb, yb, xc, yc = map(int,input().split())

tmpxb = xb - xa
tmpyb = yb - ya
tmpxc = xc - xa
tmpyc = yc - ya

S = abs((tmpxb*tmpyc - tmpxc*tmpyb) / 2)
print(S)
