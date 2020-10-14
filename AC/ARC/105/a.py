from itertools import product
items = list(map(int, input().split()))
for bits in product((0, 1), repeat=len(items)):
    comb = [x for x, bit in zip(items, bits) if bit == 0]
    print(comb)
    # choice = []
    # for j in range(n):
    #    if ((i >> j) & 1):
    #        choice.append(items[j])
    # print(choice)
