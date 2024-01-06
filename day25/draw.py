import graphviz
import sys

lines = open(sys.argv[1]).read().strip().split("\n")
g = graphviz.Graph("G", filename="process.gv", engine="sfdp")
for line in lines:
    tokens = line.split(": ")
    for t in tokens[1].split():
        g.edge(tokens[0], t)
g.view()
