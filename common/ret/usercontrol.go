package ret

type MainPage struct {
	DiskTotal     	int64 	 	`json:"disk_total"`
	DiskFree 		int64       `json:"disk_free"`
	MusicTotal      int32       `json:"music_total"`
	ForgetTotal 	int32  		`json:"forget_total"`
	ContactTotal 	int32 		`json:"contact_total"`
	VideoTotal 		int32 		`json:"video_total"`
	AccessLog 		string 		`json:"access_log"`
}
