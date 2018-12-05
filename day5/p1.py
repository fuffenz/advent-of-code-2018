with open("input.txt", "r") as f:
    original_data = f.read().rstrip()

alphabet = "abcdefghijklmnopqrstuvwxyz"
pairs = [c + c.upper() for c in alphabet]
pairs += [c.upper() + c for c in alphabet]
def react(s):
    for p in pairs:
        s = s.replace(p, "")
    return s

def full_react(s):
    ps = data
    s = data
    while True:
        s = react(ps)
        if s == ps:
            break
        ps = s
    return s

data = original_data
print(len(full_react(data)))

lens = []
for c in alphabet:
    data = original_data
    # remember to store your results!
    data = data.replace(c, "")
    data = data.replace(c.upper(), "")
    lens.append(len(full_react(data)))
print(min(lens))