#include <stdint.h>
#include <stdio.h>
#include <mongoc/mongoc.h>
#include <sys/types.h>

int main() {
    mongoc_client_t *client = mongoc_client_new("mongodb://localhost:27017");

    bson_t *reply = bson_new();

    bson_json_reader_t *reader = bson_json_reader_new_from_file("../small_doc.json", NULL);
    bson_t *doc = bson_new();
    bson_json_reader_read(reader, doc, NULL);

    const bson_t **docs = (const bson_t **) malloc(sizeof(bson_t *) * 10000);

    for (int i = 0; i < 10000; i++) {
        docs[i] = doc;
    }

    mongoc_collection_t *collection = mongoc_client_get_collection(client, "bench", "c");
    mongoc_collection_drop(collection, NULL);

    clock_t tic = clock();
    mongoc_collection_insert_many(collection, docs, 10000, NULL, NULL, NULL);
    clock_t toc = clock();

    printf("insert: %fms\n", ((double)(toc - tic) / CLOCKS_PER_SEC ) * 1000.0);

    int64_t n_docs = mongoc_collection_estimated_document_count(collection, NULL, NULL, NULL, NULL);
    printf("docs: %ld\n", n_docs);

    return 0;
}
