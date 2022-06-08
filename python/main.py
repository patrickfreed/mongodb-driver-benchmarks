import pymongo
import time
import json
import copy

def ms_between(start, end):
    return str(int((end - start) * 1000)) + "ms"

with open('../small_doc.json', 'r') as f:
    doc = json.load(f)

client = pymongo.MongoClient()
coll = client.bench.python

coll.drop()

docs = []
for i in range(10000):
    docs.append(copy.copy(doc))

start = time.time()
coll.insert_many(docs)
end = time.time()

print("insert: " + ms_between(start, end))

print("docs: " + str(coll.estimated_document_count()))
