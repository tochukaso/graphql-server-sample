package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/tochukaso/graphql-server-sample/db"
	"github.com/tochukaso/graphql-server-sample/graph/generated"
	"github.com/tochukaso/graphql-server-sample/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	product := &model.Product{
		Name:   input.Name,
		Price:  input.Price,
		Code:   getString(input.Code, ""),
		Detail: getString(input.Detail, ""),
	}

	result := db.GetDB().Create(&product)

	if result.Error != nil {
		log.Print("商品の登録に失敗しました")
	} else {
		log.Print("商品を登録しました")
	}

	return product, result.Error
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.EditProduct) (*model.Product, error) {
	gDB := db.GetDB()
	var dbProduct *model.Product
	gDB.Find(&dbProduct, input.ID)

	dbProduct.Name = getString(input.Name, dbProduct.Name)
	dbProduct.Code = getString(input.Code, dbProduct.Code)
	dbProduct.Detail = getString(input.Detail, dbProduct.Detail)
	dbProduct.Price = getInt(input.Price, dbProduct.Price)

	result := gDB.Save(dbProduct)

	if result.Error != nil {
		log.Print("プロダクトの更新に失敗しました")
	} else {
		log.Print("プロダクトを更新しました")
		r.mu.Lock()
		for _, observer := range r.ProductObservers {
			observer <- dbProduct
		}
		r.mu.Unlock()
	}

	return dbProduct, result.Error
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, productID int) (int, error) {
	gDB := db.GetDB()
	var dbProduct *model.Product
	gDB.Find(&dbProduct, productID)
	result := gDB.Delete(dbProduct)

	if result.Error != nil {
		log.Print("プロダクトの削除に失敗しました")
	} else {
		log.Print("プロダクトを削除しました")
	}

	return productID, result.Error
}

func (r *mutationResolver) CreateSku(ctx context.Context, input model.NewSku) (*model.Sku, error) {
	sku := &model.Sku{
		ProductId: input.ProductID,
		Name:      input.Name,
		Stock:     input.Stock,
		Code:      getString(input.Code, ""),
	}

	result := db.GetDB().Create(&sku)

	if result.Error != nil {
		log.Print("SKUの登録に失敗しました")
	} else {
		log.Print("SKUを登録しました")
	}

	return sku, result.Error
}

func (r *productResolver) CreatedAt(ctx context.Context, obj *model.Product) (*string, error) {
	return tToS(obj.CreatedAt), nil
}

func (r *productResolver) UpdatedAt(ctx context.Context, obj *model.Product) (*string, error) {
	return tToS(obj.UpdatedAt), nil
}

func (r *productResolver) DeletedAt(ctx context.Context, obj *model.Product) (*string, error) {
	return dtToS(obj.DeletedAt), nil
}

func (r *productResolver) Skus(ctx context.Context, obj *model.Product) ([]*model.Sku, error) {
	return ctxLoaders(ctx).skus.Load(obj.ID)
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	db.GetDB().Find(&products)
	return products, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	var product *model.Product
	db.GetDB().First(&product, id)
	return product, nil
}

func (r *queryResolver) Skus(ctx context.Context) ([]*model.Sku, error) {
	var skus []*model.Sku
	db.GetDB().Find(&skus)
	return skus, nil
}

func (r *skuResolver) CreatedAt(ctx context.Context, obj *model.Sku) (*string, error) {
	return tToS(obj.CreatedAt), nil
}

func (r *skuResolver) UpdatedAt(ctx context.Context, obj *model.Sku) (*string, error) {
	return tToS(obj.UpdatedAt), nil
}

func (r *skuResolver) DeletedAt(ctx context.Context, obj *model.Sku) (*string, error) {
	return dtToS(obj.DeletedAt), nil
}

func (r *skuResolver) Product(ctx context.Context, obj *model.Sku) (*model.Product, error) {
	return ctxLoaders(ctx).productByID.Load(obj.ProductId)
}

func (r *subscriptionResolver) UpdateProduct(ctx context.Context, id string) (<-chan *model.Product, error) {

	var product *model.Product
	db.GetDB().First(&product, id)
	productChannel := make(chan *model.Product, 1)

	r.mu.Lock()
	r.ProductObservers[id] = productChannel
	r.mu.Unlock()

	r.ProductObservers[id] <- product
	return productChannel, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Sku returns generated.SkuResolver implementation.
func (r *Resolver) Sku() generated.SkuResolver { return &skuResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type skuResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
