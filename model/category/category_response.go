package category

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(category Category) CategoryResponse {
	return CategoryResponse(category)
}

func ToCategoryResponses(categories []Category) []CategoryResponse {
	var categoryResponses []CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
