import os
import networkx as nx

DIRECTIONS = [(-1, 0), (0, -1), (1, 0), (0, 1)]

script_dir = os.path.dirname(__file__)
file_path = os.path.join(script_dir, 'input.txt')

byte_positions = [tuple(map(int, line.split(','))) for line in open(file_path).read().splitlines()]

LARGEST_COORD = 70
start = (0, 0)
end = (LARGEST_COORD, LARGEST_COORD)

G = nx.grid_2d_graph(LARGEST_COORD + 1, LARGEST_COORD + 1)
try:
    for i, (x, y) in enumerate(byte_positions):
        print(f"{i}: {x},{y}")
        G.remove_node((x, y))
        print(nx.shortest_path_length(G, source=start, target=end))
except:
    print("oopsy daisy")
