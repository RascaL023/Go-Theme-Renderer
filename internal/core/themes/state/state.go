package state

type State struct {
	Version int 	 `json:"version"`

	Theme struct {
		Name 						string `json:"name"`
		WallpaperPath		string `json:"wallpaperPath"`
	} `json:"theme"`

	Waybar  string `json:"waybar"`
}

