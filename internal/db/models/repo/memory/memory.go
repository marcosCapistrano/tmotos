package memory

import (
	"fmt"
	"snippetbox/internal/db/models"
)

type MemoryRepo struct {
	productTree ProductTree
}

func NewMemoryRepo() *MemoryRepo {
	repo := &MemoryRepo{
		productTree: NewProductTree(),
	}

	return repo
}

func (repo MemoryRepo) SearchProducts(url string) []models.Product {
	root := repo.productTree.root
	subtree := root.FindSubtree(url)

	leaves := subtree.CollectLeaves()
	for _, leaf := range leaves {
		fmt.Println(leaf.Name)
	}

	return []models.Product{}
}
