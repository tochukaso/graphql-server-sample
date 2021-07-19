package graph

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tochukaso/graphql-server-sample/db"
	"github.com/tochukaso/graphql-server-sample/graph/generated"
	"github.com/tochukaso/graphql-server-sample/graph/model"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type loaders struct {
	skus        *generated.SkuSliceLoader
	productByID *generated.ProductLoader
}

func LoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		ldrs.skus = generateSkuSliceLoader()
		ldrs.productByID = generateProductLoader()

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func generateSkuSliceLoader() *generated.SkuSliceLoader {

	config := &generated.SkuSliceLoaderConfig{
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(ids []int) ([][]*model.Sku, []error) {
			var skus []*model.Sku
			var errs []error
			if dbResult := db.GetDB().Where(" product_id in ? ", ids).Order("product_id").Find(&skus); dbResult.Error != nil {
				errs = append(errs, dbResult.Error)
				return nil, errs
			}

			var result [][]*model.Sku
			if len(skus) == 0 {
				return result, errs
			}

			skuMap := make(map[int][]*model.Sku, len(skus))
			for _, sku := range skus {
				ss := skuMap[sku.ProductId]
				if ss == nil {
					ss = make([]*model.Sku, 0)
				}
				ss = append(ss, sku)
				skuMap[sku.ProductId] = ss
			}

			for _, product_id := range ids {
				result = append(result, skuMap[product_id])
			}

			for _, r := range result {
				log.Println(&r)
			}
			return result, errs
		},
	}
	return generated.NewSkuSliceLoader(*config)
}

func generateProductLoader() *generated.ProductLoader {

	config := &generated.ProductLoaderConfig{
		Wait:     1 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(ids []int) ([]*model.Product, []error) {
			var products []*model.Product
			gdb := db.GetDB()
			var errs []error
			if result := gdb.Where(" id in ? ", ids).Order("id").Find(&products); result.Error != nil {
				errs = append(errs, result.Error)
			}

			return products, errs
		},
	}
	return generated.NewProductLoader(*config)
}

func ctxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
