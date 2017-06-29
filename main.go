package main

import (
	"errors"
	"flag"
	"fmt"

	"io/ioutil"

	"os"

	gitlab "github.com/xanzy/go-gitlab"
)

const (
	perPage = 99999
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	host := flag.String("host", "example.com", "GitLab Host Name")
	pkey := flag.String("pkey", "xxxxxxxxxx", "Your GitLab Private Key")
	outdir := flag.String("outdir", "./", "git {clone/pull} directory")
	flag.Parse()

	files, err := ioutil.ReadDir(*outdir)
	if err != nil {
		panic(err)
	}

	git := gitlab.NewClient(nil, *pkey)
	git.SetBaseURL(fmt.Sprintf("%s/api/v3", *host))

	namespaces, res, err := git.Namespaces.ListNamespaces(&gitlab.ListNamespacesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: perPage,
		},
	})
	if err != nil {
		panic(err)
	}
	if res.Status != "200 OK" {
		panic(errors.New("not 200 OK"))
	}

	for _, ns := range namespaces {
		projects, res, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{
			OrderBy: gitlab.String("name"),
			Sort:    gitlab.String("asc"),
			ListOptions: gitlab.ListOptions{
				PerPage: perPage,
			},
		})
		if err != nil {
			panic(err)
		}
		if res.Status != "200 OK" {
			panic(errors.New("not 200 OK"))
		}

		if len(projects) == 0 {
			continue
		}

		for _, p := range projects {
			if ns.Path != p.Namespace.Path {
				continue
			}
			fmt.Println(p.Path)
			if exists(files, func(filename string) bool {
				return filename == p.Path
			}) {
				fmt.Println("Exists!")
			} else {
				fmt.Println("Not Exists!")
			}
		}

	}
}

func exists(files []os.FileInfo, fn func(filename string) bool) bool {
	for _, file := range files {
		if exists := fn(file.Name()); exists {
			return true
		}
	}
	return false
}
