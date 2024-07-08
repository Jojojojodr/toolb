import sys

from googlesearch import search

query = "".join(sys.argv[1:])
res = "http://google.com/search?q="

for i in search(query, num_results=10, lang="en", safe=None, advanced=True):
    print(i.title)
    print(i.url)
    print(i.description + "\n")
    
