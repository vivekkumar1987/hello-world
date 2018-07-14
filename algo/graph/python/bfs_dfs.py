from graph_utils import Queue, Graph


def DFS_util(g, v, visited):
    print('visited=', v)
    visited[v] = True
    for u in g.adj[v]:
        if not visited[u]:
            DFS_util(g, u, visited)


def DFS(g, v):
    visited = [False] * g.V
    DFS_util(g, v, visited)


def BFS(g, v):
    visited = [False] * g.V
    q = Queue()
    q.enqueue(v)
    while len(q) > 0:
        t = q.dequeue()
        if not visited[t]:
            print('visited=', t)
            visited[t] = True
        for u in g.adj[t]:
            if not visited[u]:
                q.enqueue(u)

if __name__ == '__main__':

    g = Graph(8)
    g.addUEdge(0, 1)
    g.addUEdge(1, 7)
    g.addUEdge(1, 2)
    g.addUEdge(2, 4)
    g.addUEdge(2, 3)
    g.addUEdge(4, 5)
    g.addUEdge(4, 6)
    g.addUEdge(4, 7)
    print(g.adj)

    print('----- DFS -----')
    DFS(g, 0)

    print('----- BFS -----')
    BFS(g, 0)
