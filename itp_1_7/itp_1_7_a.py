while True:
    m, f, r = map(int, input.split())
    sum = m + f

    if m == f == r == -1:
        break

    if m == -1 or f == -1:
        grade = 'F'
    elif sum >= 80:
        grade = 'A'
    elif sum >= 65:
        grade = 'B'
    elif sum >= 50:
        grade = 'C'
    elif sum >= 30:
        grade = 'C' if r >= 50 else 'D'
    else:
        grade = 'F'

    print(grade)
