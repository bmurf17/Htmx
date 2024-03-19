package main

import (
	"log"
	"net/http"
)

func handleGetTasks(w http.ResponseWriter, _ *http.Request) {
	items, err := fetchTasks()
	if err != nil {
		log.Printf("error fetching tasks %v", err)
		return
	}
	count, err := fetchCount()
	if err != nil {
		log.Printf("error fetching count %v", err)
		return
	}
	completeCount, err := fetchCompletedCount()
	if err != nil {
		log.Printf("error fetching Completed count %v", err)
		return
	}
	data := Tasks{
		Items:          items,
		Count:          count,
		CompletedCount: completeCount,
	}
	tmpl.ExecuteTemplate(w, "Base", data)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		return
	}
	_, err := insertTask(title)
	if err == nil {
		log.Printf("error inserting task %v", err)
		return
	}
	count, err := fetchCount()
	if err == nil {
		log.Printf("error inserting count %v", err)
		return
	}
	log.Printf("count %v", count)
	w.WriteHeader(http.StatusCreated)
	tmpl.ExecuteTemplate(w, "Form", nil)
}
