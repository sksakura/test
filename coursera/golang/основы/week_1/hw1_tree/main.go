package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Elem struct {
	Prev  *Elem
	Path  string
	Index int
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	var root_node Elem = Elem{}
	root_node.Index = -1
	root_node.Path = "./testdata"
	cur_node := &root_node
	queue := make([]string, 0)
	for {
		files, err := ioutil.ReadDir(cur_node.Path)
		if err != nil {
			return err
		}

		/*если файлы печатать не надо - грохнем их перед началом обработки папки*/
		if !printFiles {
			i := 0
			for {
				if i >= len(files) {
					break
				}
				if files[i].IsDir() {
					i++
					continue
				}
				files = append(files[:i], files[i+1:]...)

			}
		}

		//пройдя все файлы
		if cur_node.Index == len(files)-1 {
			//выходим если вернулись в корень,
			if cur_node.Prev == nil {
				break
			}

			//уходим на уровень вверх
			cur_node = cur_node.Prev
			//queue = queue[1:]
			queue = queue[:len(queue)-1]
			continue
		}

		for i, f := range files {
			//пробегаем без остановки по уже обработанным
			if cur_node.Index >= i {
				continue
			}

			cur_node.Index = i

			for _, q := range queue {
				fmt.Fprint(out, q)
			}

			fdata := strings.TrimSpace(f.Name())
			if !f.IsDir() {
				fdata = strings.Replace(fmt.Sprintf("%s (%db)", fdata, f.Size()), "(0b)", "(empty)", 1)
			}

			if i == len(files)-1 {
				fmt.Fprint(out, "└───", fdata, "\n")
			} else {
				fmt.Fprint(out, "├───", fdata, "\n")
			}

			if f.IsDir() {
				//выделяем новый объект с ссылкой на текущий
				new_node := Elem{
					Path:  cur_node.Path + string(os.PathSeparator) + f.Name(),
					Prev:  cur_node,
					Index: -1,
				}
				if i == len(files)-1 {
					queue = append(queue, "\t")
				} else {
					queue = append(queue, "│\t")
				}
				cur_node = &new_node
				break
			}

			//если дошли до конца списка - прыгнем на уровень вверх
			if i == len(files)-1 {
				if cur_node.Prev != nil {
					cur_node = cur_node.Prev
					//queue = queue[1:]
					queue = queue[:len(queue)-1]
				}
			}
		}

	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
