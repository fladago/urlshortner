package shortener

type Redirect struct {
	Code      string `json:"code,omitempty" bson:"code" msgpack:"code"`
	URL       string `json:"url,omitempty" bson:"url" msgpack:"url" validate:"empty=false & format=url"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at" msgpack:"created_at"`
}
