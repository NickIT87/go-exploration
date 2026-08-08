package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

// Serching recursively for keyword rose in folder Folder2
// Searching for keyword rose in file File2
// Searching for keyword rose in file File3
// Serching recursively for keyword rose in folder Folder1
// Searching for keyword rose in file File1
