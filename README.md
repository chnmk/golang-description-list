# golang-description-list

A study project. Scan the structure of a selected folder as a set of [description lists](https://www.w3.org/TR/2011/WD-html5-author-20110809/the-dl-element.html) and export it as JSON.

For example, the following folder structure:

```
Music
- Artist 1
  - Album 1
    - Song 1
    - Song 2
- Artist 2
  - Song 3 
```

Will be exported as 

```
[
  {
    "Name": "Artist 1", 
    "Contents": [
      {
        "Category": "Album 1",
        "Entries": [
          "Song 1",
          "Song 2"
        ]
      }
    ]
  }
  {
    "Name:" "Artist 2",
    "Contents": [
      {
        "Category": "Unsorted",
        "Entries": [
          "Song 3"
        ]
      }
    ]
  }
]
```

See also: [obsidian-mdb-plugin](https://github.com/chnmk/obsidian-mdb-plugin).

### TODO:
- User interface.
- "No file format" option.
- "Check mp3 metadata" option.
- Option to scan folders to check if all files are present.
- Compile this project as an executable binary.

### How to use (temporarily):
1. Install golang.
1. Copy this project to the folder you want to scan.
2. Input "go run ./src" in console.
3. Follow the instructions.

