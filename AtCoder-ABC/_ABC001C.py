deg, dis = map(int, input().split())
d = deg / 10

if 11.25 < d < 33.75:
    print("NNE")
elif 33.75 < d < 56.25:
    print("NE")
elif 56.25 < d < 78.75:
    print("ENE")
elif 78.75 < d < 101.25:
    print("E")
elif 101.25 < d < 123.75:
    print("ESE")
elif 123.75 < d < 146.25:
    print("SE")
elif 146.25 < d < 168.75:
    print("SSE")
elif 168.75 < d < 191.25:
    print("S")
elif 168.75 < d < 191.25:

