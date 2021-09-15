package loadBalancer

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/Aleksao998/LoadBalancer/api/pb/bankAccountpb"
	"github.com/Aleksao998/LoadBalancer/api/pb/expensespb"
	"google.golang.org/grpc"
)

var ErrServersNotExists = errors.New("servers dose not exist")

type WorkerService struct {
	BankService    bankAccountpb.BankAccountServiceClient
	ExpenseService expensespb.ExpensesServiceClient
	Url            string
}

type WorkersPool struct {
	Workers []*WorkerService
	next    uint32
}

func NewLoadbalancer() *WorkersPool {
	workersPool := make([]*WorkerService, 0)
	return &WorkersPool{
		Workers: workersPool,
	}
}

func (r *WorkersPool) RemoveUrl(oldUrl string) {
	var index int
	index = -1
	for i, worker := range r.Workers {
		if worker.Url == oldUrl {
			index = i
		}
	}
	if index != -1 {
		r.Workers[index] = r.Workers[len(r.Workers)-1]
		r.Workers = r.Workers[:len(r.Workers)-1]
	}
}

func (r *WorkersPool) AddUrl(newUrl string) {
	cc, err := grpc.Dial(newUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Print(err)
		fmt.Printf("Error\n", err)
	}
	ba := bankAccountpb.NewBankAccountServiceClient(cc)
	ex := expensespb.NewExpensesServiceClient(cc)
	worker := &WorkerService{
		BankService:    ba,
		ExpenseService: ex,
		Url:            newUrl,
	}

	r.Workers = append(r.Workers, worker)
}

func (r *WorkersPool) Next() *WorkerService {
	n := atomic.AddUint32(&r.next, 1)
	if len(r.Workers) != 0 {
		return r.Workers[(int(n)-1)%len(r.Workers)]
	}
	return nil
}
