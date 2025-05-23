# This program converts a txt version of the Bible into a more readable JSON format.
# It requires that the book names are on their own line and all their verses listed below.

import re
import time
import json

book_names = [  # For books with more than one variant use 1, 2, 3, etc.
    "Genesis", "Exodus", "Leviticus", "Numbers", "Deuteronomy",
    "Joshua", "Judges", "Ruth", "1Samuel", "2Samuel", "1Kings", "2Kings",
    "1Chronicles", "2Chronicles", "Ezra", "Nehemiah", "Esther", "Job",
    "Psalms", "Proverbs", "Ecclesiastes", "Song of Solomon", "Isaiah",
    "Jeremiah", "Lamentations", "Ezekiel", "Daniel", "Hosea", "Joel", "Amos",
    "Obadiah", "Jonah", "Micah", "Nahum", "Habakkuk", "Zephaniah", "Haggai",
    "Zechariah", "Malachi", "Matthew", "Mark", "Luke", "John", "Acts",
    "Romans", "1Corinthians", "2Corinthians", "Galatians", "Ephesians",
    "Philippians", "Colossians", "1Thessalonians", "2Thessalonians",
    "1Timothy", "2Timothy", "Titus", "Philemon", "Hebrews", "James",
    "1Peter", "2Peter", "1John", "2John", "3John", "Jude", "Revelation"
]

books = {}
current_book = None
current_chapter = None

name = "kjv"    # Change this to the name of the txt file you want to convert

with open("raw_bible/" + name + ".txt", encoding="utf-8") as file:
    lines = file.readlines()

    verse_pattern = re.compile(r'^(\d+):(\d+)\s+(.*)')
    start = time.time()
    verse_buffer = ""
    verse_chapter = None
    verse_number = None

    for line in lines:
        line = line.strip()
        if not line:
            continue

        # Detect book name
        for book in book_names:
            if re.fullmatch(re.escape(book), line, re.IGNORECASE):
                current_book = book
                if current_book not in books:
                    books[current_book] = {}
                    print(f"Found book: {current_book}")
                current_chapter = None
                verse_buffer = ""
                verse_chapter = None
                verse_number = None
                break
        else:
            # Detect verse start
            verse_match = verse_pattern.match(line)
            if verse_match and current_book:
                # Save previous verse if any
                if verse_buffer and verse_chapter and verse_number:
                    if verse_chapter not in books[current_book]:
                        books[current_book][verse_chapter] = []
                    books[current_book][verse_chapter].append(f"{verse_number} {verse_buffer.strip()}")
                # Start new verse
                verse_chapter = verse_match.group(1)
                verse_number = verse_match.group(2)
                verse_buffer = verse_match.group(3)
            elif verse_buffer:
                # Continuation of previous verse (line wrap)
                verse_buffer += " " + line

    # Save the last verse in the file
    if verse_buffer and verse_chapter and verse_number:
        if verse_chapter not in books[current_book]:
            books[current_book][verse_chapter] = []
        books[current_book][verse_chapter].append(f"{verse_number} {verse_buffer.strip()}")

end = time.time()
print(f"{end-start} seconds to parse the bible")

# Save the books dictionary as JSON in the canonical order of book_names
ordered_books = {book: books[book] for book in book_names if book in books}
with open(name + ".json", "w", encoding="utf-8") as f:
    json.dump(ordered_books, f, ensure_ascii=False, indent=2)

print(f"Saved as {name+'.json'}")