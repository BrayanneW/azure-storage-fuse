/*
    _____           _____   _____   ____          ______  _____  ------
   |     |  |      |     | |     | |     |     | |       |            |
   |     |  |      |     | |     | |     |     | |       |            |
   | --- |  |      |     | |-----| |---- |     | |-----| |-----  ------
   |     |  |      |     | |     | |     |     |       | |       |
   | ____|  |_____ | ____| | ____| |     |_____|  _____| |_____  |_____


   Licensed under the MIT License <http://opensource.org/licenses/MIT>.

   Copyright © 2020-2024 Microsoft Corporation. All rights reserved.
   Author : <blobfusedev@microsoft.com>

   Permission is hereby granted, free of charge, to any person obtaining a copy
   of this software and associated documentation files (the "Software"), to deal
   in the Software without restriction, including without limitation the rights
   to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
   copies of the Software, and to permit persons to whom the Software is
   furnished to do so, subject to the following conditions:

   The above copyright notice and this permission notice shall be included in all
   copies or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
   OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
   SOFTWARE
*/

package block_pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type blockpoolTestSuite struct {
	suite.Suite
	assert *assert.Assertions
}

func (suite *blockpoolTestSuite) SetupTest() {
}

func (suite *blockpoolTestSuite) cleanupTest() {
}

func (suite *blockpoolTestSuite) TestAllocate() {
	suite.assert = assert.New(suite.T())

	bp := Allocate[interface{}](0, 0)
	suite.assert.Nil(bp)

	bp = Allocate[interface{}](1, 0)
	suite.assert.Nil(bp)

	bp = Allocate[interface{}](1, 1)
	suite.assert.NotNil(bp)
	suite.assert.NotNil(bp.blocksCh)

	bp.Terminate()
	suite.assert.Equal(len(bp.blocksCh), 0)
}

func (suite *blockpoolTestSuite) TestGetReleaseBlock() {
	suite.assert = assert.New(suite.T())

	bp := Allocate[interface{}](1, 5)
	suite.assert.NotNil(bp)
	suite.assert.NotNil(bp.blocksCh)
	suite.assert.Equal(len(bp.blocksCh), 5)

	b := bp.GetBlock()
	suite.assert.NotNil(b)
	suite.assert.Equal(len(bp.blocksCh), 4)

	bp.ReleaseBlock(b)
	suite.assert.Equal(len(bp.blocksCh), 5)

	b = bp.GetBlock()
	suite.assert.NotNil(b)
	suite.assert.Equal(len(bp.blocksCh), 4)

	bp.ReleaseBlock(b)
	suite.assert.Equal(len(bp.blocksCh), 5)

	bp.Terminate()
	suite.assert.Equal(len(bp.blocksCh), 0)
}

func (suite *blockpoolTestSuite) TestUsage() {
	suite.assert = assert.New(suite.T())

	bp := Allocate[interface{}](1, 5)
	suite.assert.NotNil(bp)
	suite.assert.NotNil(bp.blocksCh)
	suite.assert.Equal(len(bp.blocksCh), 5)

	var blocks []*Block[interface{}]
	b := bp.GetBlock()
	suite.assert.NotNil(b)
	blocks = append(blocks, b)

	usage := bp.Usage()
	suite.assert.Equal(usage, uint32(20))

	b = bp.GetBlock()
	suite.assert.NotNil(b)
	blocks = append(blocks, b)

	usage = bp.Usage()
	suite.assert.Equal(usage, uint32(40))

	for _, blk := range blocks {
		bp.ReleaseBlock(blk)
	}

	bp.Terminate()
	suite.assert.Equal(len(bp.blocksCh), 0)
}

func (suite *blockpoolTestSuite) TestBufferExhaution() {
	suite.assert = assert.New(suite.T())

	bp := Allocate[interface{}](1, 5)
	suite.assert.NotNil(bp)
	suite.assert.NotNil(bp.blocksCh)
	suite.assert.Equal(len(bp.blocksCh), 5)

	var blocks []*Block[interface{}]
	for i := 0; i < 5; i++ {
		b := bp.GetBlock()
		suite.assert.NotNil(b)
		blocks = append(blocks, b)
	}

	usage := bp.Usage()
	suite.assert.Equal(usage, uint32(100))

	b := bp.GetBlockNoWait()
	suite.assert.Nil(b)

	for _, blk := range blocks {
		bp.ReleaseBlock(blk)
	}

	bp.Terminate()
	suite.assert.Equal(len(bp.blocksCh), 0)
}

func TestBlockPoolSuite(t *testing.T) {
	suite.Run(t, new(blockpoolTestSuite))
}
