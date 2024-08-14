package service

import (
	"context"
	"fmt"
	"io"

	"github.com/d4niells/goexpert-grpc/internal/database"
	"github.com/d4niells/goexpert-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	dbCategory, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	category := &pb.Category{
		Id:          dbCategory.ID,
		Name:        dbCategory.Name,
		Description: dbCategory.Description,
	}

	return category, nil
}

func (c *CategoryService) ListCategories(_ context.Context, _ *pb.Blank) (*pb.CategoryList, error) {
	dbCategories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categories []*pb.Category
	for _, v := range dbCategories {
		categories = append(categories, &pb.Category{Id: v.ID, Name: v.Name, Description: v.Description})
	}

	return &pb.CategoryList{Categories: categories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{Id: category.ID, Name: category.Name, Description: category.Description}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	var categories []*pb.Category

	for {
		category, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.CategoryList{Categories: categories})
			return nil
		}
		if err != nil {
			return err
		}

		dbCategory, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories = append(categories, &pb.Category{Id: dbCategory.ID, Name: dbCategory.Name, Description: dbCategory.Description})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		dbCategory, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.SendMsg(&pb.Category{Id: dbCategory.ID, Name: dbCategory.Name, Description: dbCategory.Description})
		if err != nil {
			return err
		}
	}
}

func (c *CategoryService) DeleteCategoryStreamBidirectional(stream pb.CategoryService_DeleteCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		err = c.CategoryDB.DeleteByID(category.Id)
		if err != nil {
			return err
		}

		err = stream.SendMsg(&pb.DeleteCategoryResponse{Message: fmt.Sprintf("Category %s was deleted", category.Id)})
		if err != nil {
			return err
		}
	}
}
