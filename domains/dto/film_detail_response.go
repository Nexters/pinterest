package dto

type FilmDetailResponse struct {
	FilmID        uint                     `json:"film_id"`
	Title         string                   `json:"title"`
	Order         uint                     `json:"order"`
	PhotoCutCount uint                     `json:"photo_cut_count"`
	Likes         uint                     `json:"likes"`
	UserID        string                   `json:"user_id"`
	PhotoCuts     []PhotoCutDetailResponse `json:"photo_cuts"`
}
