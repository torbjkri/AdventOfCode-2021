module solutions

go 1.17

replace dec1 => ../1

require (
	dec1 v0.0.0-00010101000000-000000000000
	dec2 v0.0.0-00010101000000-000000000000
	dec3 v0.0.0-00010101000000-000000000000
)

require my_utils v0.0.0-00010101000000-000000000000 // indirect

replace my_utils => ../utils

replace dec2 => ../2

replace dec3 => ../3
