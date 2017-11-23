package animalgrpc

import (
	"dino/databaselayer"
	fmt "fmt"

	context "golang.org/x/net/context"
)

type GrpcServer struct {
	dbHandler databaselayer.AnimalDBHandler
}

func NewAnimalGrpcServer(dbtype uint8, connstring string) (*GrpcServer, error) {
	handler, err := databaselayer.GetDatabaseHandler(dbtype, connstring) //databaselayer.MONGODB	"MONGODB://127.0.0.1"
	if err != nil {
		return nil, fmt.Errorf("Could not create a database handler object, error %v", err)
	}
	return &GrpcServer{
		dbHandler: handler,
	}, nil
}

func (server *GrpcServer) GetAnimal(ctx context.Context, r *Request) (*Animal, error) {
	animal, err := server.dbHandler.GetAnimalByNickname(r.GetNickname())
	return convertToGRPCAnimal(animal), err
}

func (server *GrpcServer) GetAllAnimals(req *Request, stream AnimalService_GetAllAnimalsServer) error {
	animals, err := server.dbHandler.GetAvailableAnimals()
	if err != nil {
		return err
	}
	for _, animal := range animals {
		grpcAnimal := convertToGRPCAnimal(animal)
		if err := stream.Send(grpcAnimal); err != nil {
			return err
		}
	}
	return nil
}

func convertToGRPCAnimal(animal databaselayer.Animal) *Animal {
	return &Animal{Id: int32(animal.ID),
		AnimalType: animal.AnimalType,
		Nickname:   animal.Nickname,
		Zone:       int32(animal.Zone),
		Age:        int32(animal.Age),
	}
}
