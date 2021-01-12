package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rlgino/monda-todo/src/api/context/task/application"
	"github.com/rlgino/monda-todo/src/api/context/task/domain"
	"github.com/rlgino/monda-todo/src/api/context/task/infraestructure/db"
)

// taskHandler handler
type taskHandler struct {
	useCases application.TaskUseCases
}

// Run task handlers
func Run() {
	db := db.InMemoryDB{}
	handler := taskHandler{
		useCases: application.New(&db),
	}

	http.HandleFunc("/task", handler.handleTaskRequest)
	http.HandleFunc("/task/check", handler.handleCheckedTaskRequest)
}

func (handler taskHandler) handleCheckedTaskRequest(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("task_id")

		if id != "" {
			err := handler.useCases.CheckTask(id)
			if err != nil {
				if errors.Is(err, domain.NotFoundError{}) {
					rw.WriteHeader(404)
					return
				}
				rw.WriteHeader(500)
				rw.Write([]byte(err.Error()))
				return
			}
			return
		}
	}

	rw.WriteHeader(400)
}

func (handler taskHandler) handleTaskRequest(rw http.ResponseWriter, r *http.Request) {
	enableCors(&rw)
	if r.Method == http.MethodPost {
		request := &taskRequest{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		creationDate, err := time.Parse("2006-01-02", request.CreationDate)
		if err != nil {
			rw.WriteHeader(400)
			rw.Write([]byte(err.Error()))
			return
		}
		handler.useCases.CreateTask(request.ID, request.Title, request.OwnerID, creationDate)
	} else if r.Method == http.MethodPut {
		request := &taskRequest{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		err = handler.useCases.UpdateTaskTitle(request.ID, request.Title)
		if err != nil {
			if errors.Is(err, domain.NotFoundError{}) {
				rw.WriteHeader(404)
				return
			}
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

	} else if r.Method == http.MethodGet {
		id := r.URL.Query().Get("task_id")

		if id != "" {
			task, err := handler.useCases.FindTask(id)
			if err != nil {
				if errors.Is(err, domain.NotFoundError{}) {
					rw.WriteHeader(404)
					return
				}
				rw.WriteHeader(500)
				rw.Write([]byte(err.Error()))
				return
			}
			resp := taskToResponse(task)
			response, _ := json.Marshal(resp)
			rw.Header().Add("Content-Type", "application/json")
			rw.Write(response)
			return
		}

		owner := r.URL.Query().Get("owner_id")
		tasks, err := handler.useCases.FindTasks(owner)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		respArray := make([]taskResponse, len(tasks))
		for i, el := range tasks {
			respArray[i] = taskToResponse(el)
		}

		response, _ := json.Marshal(respArray)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write(response)
	}
}

func taskToResponse(task domain.Task) taskResponse {
	return taskResponse{
		ID:        task.ID.Value(),
		Title:     task.Title.Value(),
		OwnerID:   task.TaskOwner.Value(),
		IsChecked: task.TaskChecked.Value(),
	}
}

type taskRequest struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	CreationDate string `json:"creation_date"`
	OwnerID      string `json:"owner"`
}

type taskResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	CreationDate string `json:"creation_date"`
	OwnerID      string `json:"owner"`
	IsChecked    bool   `json:"is_checked"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
