package payload

type AssignCategoryProductPayload struct {
	ProductID   string   `json:"product_id"`
	CategoryIDs []string `json:"category_ids"`
}
