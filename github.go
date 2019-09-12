package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

func githubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	base := fmt.Sprintf("https://github.com/%s/%s", vars["user"], vars["repo"])

	// TODO parse "git ls-remote --tags" output
	git := exec.Command("git", "ls-remote", "--tags", base+".git")
	tags, err := git.Output()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("git ls-remote %s failed: %v", base, err)
		return
	}

	data := HtmlTemplate{
		Title:    base,
		Versions: map[string]TemplateVersion{},
	}

	for _, line := range strings.Split(string(tags), "\n") {
		if len(line) == 0 {
			continue
		}
		lineParts := strings.SplitN(line, "\t", 2)
		refParts := strings.SplitN(lineParts[1], "/", 3)
		version := strings.TrimSuffix(refParts[2], "^{}")
		data.Versions[version] = TemplateVersion{
			Url: fmt.Sprintf("%s/archive/%s.tar.gz", base, version),
		}
	}

	out, err := renderTemplate(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error rendering: %v", err)
		return
	}

	w.Header().Set("Cache-Control", "max-age=3600")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, out)
}
