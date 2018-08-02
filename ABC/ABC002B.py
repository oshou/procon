w = input()
vowel = ["a","i","u","e","o"]
for i in w:
    if i in vowel:
        w = w.replace(i,"")

print(w)
