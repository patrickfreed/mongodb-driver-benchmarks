use std::{fs::File, time::Instant};

use anyhow::Result;
use mongodb::{bson::RawDocumentBuf, Client};

#[tokio::main(flavor = "current_thread")]
async fn main() -> Result<()> {
    let client = Client::with_uri_str("mongodb://localhost:27017").await?;

    let coll = client.database("foo").collection::<RawDocumentBuf>("foo");
    coll.drop(None).await?;

    let mut file = File::open("../small_doc.json")?;
    let doc: RawDocumentBuf = serde_json::from_reader(&mut file)?;
    let docs = (0..10_000).map(|_| doc.clone());

    let start = Instant::now();
    coll.insert_many(docs, None).await?;
    let duration = start.elapsed();

    println!("insert: {}ms", duration.as_millis());

    let docs = coll.estimated_document_count(None).await?;
    println!("docs: {}", docs);

    Ok(())
}
