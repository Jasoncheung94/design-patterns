package main

import "fmt"

type Prototype interface {
	display()
	clone() Prototype
}

type Directory struct {
	name  string
	files []Prototype
}

func (d *Directory) display() {
	fmt.Println("Directory:", d.name, "with files:")
	for _, file := range d.files {
		fmt.Print("     ")
		file.display()
	}
}

func (d *Directory) clone() Prototype {
	var files []Prototype
	for _, file := range d.files {
		files = append(files, file.clone())
	}
	return &Directory{name: d.name, files: files}
}

type File struct {
	name string
}

func (f *File) display() {
	fmt.Println("File:", f.name)
}

func (f *File) clone() Prototype {
	return &File{name: f.name}
}

func main() {
	// Create files
	f1 := &File{name: "file1"}
	f2 := &File{name: "file2"}
	f3 := &File{name: "file3"}
	f1.display()
	f2.display()
	f3.display()

	// Create directory
	d1 := &Directory{name: "directory1", files: []Prototype{f1, f2, f3}}
	d1.display()
}
