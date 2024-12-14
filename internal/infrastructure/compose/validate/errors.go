package validate

const (
	REQUIRED_STREAM_FROM = "stream_from__required"
	VA_TIME              = "va_time"
	CHECK_REFERENCES     = "check_refs"
)

var MapErrorTags = map[string]string{
	"required":           "is required",
	REQUIRED_STREAM_FROM: "one of input_id, filter_output_name, stream_name is required",
	VA_TIME:              "please use validate format for time. example: 15:04:05.000",
	CHECK_REFERENCES:     "references are not valid",
}
