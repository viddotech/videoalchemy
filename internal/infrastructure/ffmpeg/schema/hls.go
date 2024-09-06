package schema

type HLS struct {
	Time               float64 `yaml:"time"`                                                           // Duration of each segment in seconds
	ListSize           int     `yaml:"list_size"`                                                      // Maximum number of entries in the playlist
	SegmentFilename    string  `yaml:"segment_filename"`                                               // Filename pattern for the segments
	PlaylistType       string  `validate:"omitempty,oneof=event vod vod-live" yaml:"playlist_type"`    // Playlist type (event, vod, vod-live)
	SegmentType        string  `validate:"omitempty,oneof=mpegts fmp4" yaml:"segment_type"`            // Segment type (mpegts, fmp4)
	Flags              string  `validate:"omitempty,oneof=single_file program_date_time" yaml:"flags"` // Additional HLS flags
	MasterPlaylistName string  `yaml:"master_playlist_name"`                                           // Name of the master playlist file
	SegmentList        string  `yaml:"segment_list"`                                                   // Filename for the segment list
	SegmentListSize    int     `yaml:"segment_list_size"`                                              // Number of segments in the segment list
	MaxEntries         int     `yaml:"max_entries"`                                                    // Limits the number of entries in the playlist
	AllowCache         bool    `yaml:"allow_cache"`                                                    // Allows or disallows caching of segments
	KeyInfoFile        string  `yaml:"key_info_file"`                                                  // File containing encryption key information
	KeyURL             string  `yaml:"key_url"`                                                        // URL for the encryption key
}
