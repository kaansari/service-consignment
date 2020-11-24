package shipping

import (

	// Import the generated protobuf code
	"fmt"
	"sync"

	pb "github.com/kaansari/service-consignment/proto/consignment"
)

type Shipping interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	Get(*pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type ShippingRepository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *ShippingRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

// GetAll consignments
func (repo *ShippingRepository) GetAll() []*pb.Consignment {

	return repo.consignments

}

func (repo *ShippingRepository) Get(consignment *pb.Consignment) (*pb.Consignment, error) {

	id := consignment.Id

	for i := range repo.consignments {
		if repo.consignments[i].Id == id {
			return repo.consignments[i], nil
		}
	}

	return nil, fmt.Errorf("Consignment not found with the given id %v", id)
}
