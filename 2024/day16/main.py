import os
import networkx as nx

DIRECTIONS = [(-1, 0), (0, -1), (1, 0), (0, 1)]

script_dir = os.path.dirname(__file__)
file_path = os.path.join(script_dir, 'input.txt')

with open(file_path, "r") as infile:
    data = infile.read().splitlines()
    G = nx.DiGraph()
    start = (139, 1, 3)
    end = (1, 139, 3)

    for row_index, row in enumerate(data):
        for col_index, cell in enumerate(row):
            if cell != '#':
                for direction in range(4):
                    G.add_node((row_index, col_index, direction))

    for r, c, direction in list(G.nodes):
        rd, cd = DIRECTIONS[direction]
        rn, cn = r + rd, c + cd
        if (rn, cn, direction) in G.nodes:
            G.add_edge((r, c, direction), (rn, cn, direction), weight=1)
        for rd in range(4):
            G.add_edge((r, c, direction), (r, c, rd), weight=1000)

    p1 = nx.shortest_path_length(G, start, end, weight='weight')
    print(p1)
    seats = set()
    for path in nx.all_shortest_paths(G, start, end, weight='weight'):
        for cn in path:
            seats.add((cn[0], cn[1]))
    print(len(seats))