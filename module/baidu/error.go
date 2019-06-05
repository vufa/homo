//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, June 2019
//

package baidu

import "fmt"

type ErrSpeechQuality struct {
	ErrNo  int
	ErrMsg string
}

// IsErrPixivImageUrlAlreadyExist checks if an error is a ErrPixivImageUrlAlreadyExist.
func IsErrSpeechQuality(err error) bool {
	_, ok := err.(ErrSpeechQuality)
	return ok
}

func (err ErrSpeechQuality) Error() string {
	return fmt.Sprintf("speech quality error. [ErrNo: %d, ErrMsg: %s]", err.ErrNo, err.ErrMsg)
}
