// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MediaContentRatingGermany undocumented
type MediaContentRatingGermany struct {
	// Object is the base model of MediaContentRatingGermany
	Object
	// MovieRating Movies rating selected for Germany
	MovieRating *RatingGermanyMoviesType `json:"movieRating,omitempty"`
	// TvRating TV rating selected for Germany
	TvRating *RatingGermanyTelevisionType `json:"tvRating,omitempty"`
}
