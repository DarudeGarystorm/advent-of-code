import os

with open(os.path.join(os.path.dirname(__file__), 'input.txt'), "r", encoding="utf-8") as f:
    lines = f.read().splitlines()
patterns, desired_designs = lines[0].split(', '), lines[2:]

memo = {"": 1}
def ways_to_complete_design(design: str) -> int:
    if design in memo:
        return memo[design]
    memo[design] = sum(ways_to_complete_design(design[len(p):]) for p in patterns if design.startswith(p))
    return memo[design]

print("Part 1:", sum(ways_to_complete_design(dd) > 0 for dd in desired_designs))
print("Part 2:", sum(ways_to_complete_design(dd) for dd in desired_designs))
