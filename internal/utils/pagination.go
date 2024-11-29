package utils

import "strconv"

// PreparePaginationParams calculates the offset and limit for pagination.
func PreparePaginationParams(after *int64, first int64) (int64, int64) {
	offset := int64(0)

	// If `after` is provided, use it as the offset.
	if after != nil {
		offset = *after
	}

	return offset, first
}

// PrepareSorting determines the sort field and order based on input or defaults.
func PrepareSorting(sortField, sortOrder, field, order string) (string, string) {
	// If the provided field is empty, use the default.
	if field == "" {
		field = sortField
	}

	// If the provided order is empty, use the default.
	if order == "" {
		order = sortOrder
	}

	return field, order
}

// PrepareSorting determines the sort field and order based on input or defaults.
// It will use the provided values if they exist, otherwise default to the provided defaults.
func PrepareSortingTwo(sortField, sortOrder string, sortFieldParam *string, sortOrderParam *string) (string, string) {
	// If the sort parameters are nil, use the default values
	if sortFieldParam == nil || *sortFieldParam == "" {
		sortFieldParam = &sortField
	}

	if sortOrderParam == nil || *sortOrderParam == "" {
		sortOrderParam = &sortOrder
	}

	// Return the final sort field and order
	return *sortFieldParam, *sortOrderParam
}

// GenerateCursor generates a cursor string for pagination.
func GenerateCursor(offset, index int64) string {
	return strconv.FormatInt(offset+index+1, 10)
}

// CalculateHasNextPage determines if there are more pages based on offset, current length, and total count.
func CalculateHasNextPage(offset, currentLength, totalCount int64) bool {
	return offset+currentLength < totalCount
}
