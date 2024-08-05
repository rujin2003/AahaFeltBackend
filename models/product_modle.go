package modle

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
