const SourceFile = "../../data/instapaper-export.csv";
const DestFile = "../../data/instapaper-full-text.json";
const fs = require("fs");

const { writeFileSync } = require("fs");

const { Readability } = require("@mozilla/readability");
const { JSDOM } = require("jsdom");
const fetch = require("node-fetch");

const data = fs.readFileSync(SourceFile).toString();
const parseCsv = require("./csv");
const PocketDocs = parseCsv(data);
const PartialDestFile = require(DestFile);

function startsWith(str, word) {
  return str.lastIndexOf(word, 0) === 0;
}

console.log(
  `Found ${PocketDocs.length} docs, downloading and parsing using @mozilla/readability.`
);

(async function () {
  // Map of href to doc type
  FullTextDocs = loadExistingEntries();

  let i = 0;
  for (let { Title, URL } of PocketDocs) {
    Title = Title[0];
    if (i % 25 === 0) {
      // To make this process interruptible, we write a partial progress
      // cache every 25 items.
      const docsSoFar = Object.values(FullTextDocs);
      console.log(`Writing partial cache with ${docsSoFar.length} docs...`);
      writeFileSync(DestFile, JSON.stringify(docsSoFar), "utf8");
    }

    if (!startsWith(URL, "http")) {
      i++;
      continue;
    }
    const alreadyParsed = FullTextDocs[URL];
    if (alreadyParsed) {
      console.log(`Using ${URL} found in partial cache...`);
      i++;
      continue;
    }
    // Skip attempting to parse media files
    if (isMediaFile(URL)) {
      FullTextDocs[URL] = {
        title: Title,
        content: URL,
        href: URL,
      };
      i++;
      continue;
    }

    console.log(`Parsing (${i + 1}/${PocketDocs.length}) ${URL}...`);
    // For a number of reasons, either JSDOM or Readability may throw if it
    // fails to parse the page. In those cases, we bail and just keep the
    // title + href.
    try {
      // Download the HTML source
      const html = await fetch(URL).then((resp) => resp.text());

      // Create a mock document to work with Readability.js
      const doc = new JSDOM(html, { url: URL });

      // Parse with Readability.
      const reader = new Readability(doc.window.document, {
        // Default is (at time of writing) 500 chars. We want to consider
        // shorter documents valid, too, for purpose of Monocle search.
        charThreshold: 20,
      });

      const page = reader.parse();
      if (!page) {
        FullTextDocs[URL] = {
          title: Title,
          content: URL,
          href: URL,
        };

        i++;
        continue;
      }

      const { title: readabilityTitle, textContent, siteName } = page;
      // If the page is longer than ~10k words, don't cache or index.
      // It's not worth it.
      if (textContent.length > 5 * 10000) {
        FullTextDocs[URL] = {
          title: Title,
          content: URL,
          href: URL,
        };

        i++;
        continue;
      }
      FullTextDocs[URL] = {
        title: siteName
          ? `${readabilityTitle} | ${siteName}`
          : readabilityTitle,
        content: textContent || URL,
        href: URL,
      };
    } catch (e) {
      console.log(`Error during parse of ${URL} (${e})... continuing.`);
      FullTextDocs[URL] = {
        title: Title,
        content: URL,
        href: URL,
      };
    }

    i++;
  }

  writeFileSync(DestFile, JSON.stringify(Object.values(FullTextDocs)), "utf8");
  console.log("done!");
})();

function isMediaFile(URL) {
  return (
    URL.endsWith(".png") ||
    URL.endsWith(".jpg") ||
    URL.endsWith(".gif") ||
    URL.endsWith(".mp4") ||
    URL.endsWith(".mov") ||
    URL.endsWith(".pdf")
  );
}

function loadExistingEntries() {
  const FullTextDocs = {};
  console.log(`Partial cache contained ${PartialDestFile.length} docs.`);
  for (const doc of PartialDestFile) {
    FullTextDocs[doc.href] = doc;
  }
  return FullTextDocs;
}
