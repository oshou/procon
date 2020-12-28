a,b,x,y = map(int,input().split())
up = min(y,2*x)
if a > b:
    print(((a-b)-1)*up+x)
else:
    print((b-a)*up+x)
