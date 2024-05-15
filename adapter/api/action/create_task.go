package action

import (
	"encoding/json"
	"net/http"

	"attendance-record/adapter/api/logging"
	"attendance-record/adapter/api/response"
	"attendance-record/adapter/logger"
	"attendance-record/usecase"
)

type CreateTaskAction struct {
	uc  usecase.CreateTaskUsecase
	log logger.Logger

	logKey, logMsg string
}

func NewCreateTaskAction(uc usecase.CreateTaskUsecase, log logger.Logger) CreateTaskAction {
	return CreateTaskAction{
		uc:     uc,
		log:    log,
		logKey: "create_task_action",
		logMsg: "create a new task",
	}
}

func (a CreateTaskAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.
			NewError(a.log, err, a.logKey, http.StatusBadRequest).
			Log(a.logMsg)
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.
			NewError(a.log, err, a.logKey, http.StatusInternalServerError).
			Log(a.logMsg)
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	logging.
		NewInfo(a.log, a.logKey, http.StatusCreated).
		Log(a.logMsg)

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
