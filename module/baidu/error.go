//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, June 2019
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
