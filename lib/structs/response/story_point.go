package response

type StoryPoints struct {
	StoryPoints []StoryPoint `json:"storyPoints"`
}

type StoryPoint struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Team   *Team   `json:"team"`
	Member *Member `json:"member"`
}
