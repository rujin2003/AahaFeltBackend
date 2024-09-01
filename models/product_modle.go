package model

type Product struct {
	ID            string   `json:"id"`
	Weight        string   `json:"weight"`
	Price         string   `json:"price"`
	MostPopular   bool     `json:"most_popular"`
	Bestseller    bool     `json:"bestseller"`
	Material      string   `json:"material"`
	Stock         int      `json:"stock"`
	NewArrival    bool     `json:"new_arrival"`
	Designer      string   `json:"designer"`
	Company       string   `json:"company"`
	HotCollection bool     `json:"hot_collection"`
	Colors        []string `json:"colors"`
	Category      string   `json:"category"`
	Description   string   `json:"description"`
	Reviews       int      `json:"reviews"`
	Stars         float64  `json:"stars"`
	Name          string   `json:"name"`
	Notes         string   `json:"notes"`
	Featured      bool     `json:"featured"`
	Sale          bool     `json:"sale"`
	Trending      bool     `json:"trending"`
	Shipping      string   `json:"shipping"`
	Origin        string   `json:"origin"`
	Image         string   `json:"image"`
	Images        []string `json:"images"`
	Exclusive     bool     `json:"exclusive"`
	NewInMarket   bool     `json:"new_in_market"`
}

func NewProduct(id string, weight string, price string, mostPopular bool, bestseller bool, material string, stock int, newArrival bool, designer string, company string, hotCollection bool, colors []string, category string, description string, reviews int, stars float64, name string, notes string, featured bool, sale bool, trending bool, shipping string, origin string, image string, images []string, exclusive bool, newInMarket bool) *Product {
	return &Product{
		ID:            id,
		Weight:        weight,
		Price:         price,
		MostPopular:   mostPopular,
		Bestseller:    bestseller,
		Material:      material,
		Stock:         stock,
		NewArrival:    newArrival,
		Designer:      designer,
		Company:       company,
		HotCollection: hotCollection,
		Colors:        colors,
		Category:      category,
		Description:   description,
		Reviews:       reviews,
		Stars:         stars,
		Name:          name,
		Notes:         notes,
		Featured:      featured,
		Sale:          sale,
		Trending:      trending,
		Shipping:      shipping,
		Origin:        origin,
		Image:         image,
		Images:        images,
		Exclusive:     exclusive,
		NewInMarket:   newInMarket,
	}
}
