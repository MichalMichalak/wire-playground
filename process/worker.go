package process

type DataWarehouse interface {
	GetByID(id int) string
}

func NewWorker(dw DataWarehouse) Worker {
	return Worker{warehouse: dw}
}

type Worker struct {
	warehouse DataWarehouse
}

func (w Worker) Work(dataID int) string {
	println("Ugh oh working hard")
	result := w.warehouse.GetByID(dataID)
	return result
}
