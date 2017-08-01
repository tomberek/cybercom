/* {{{ Copyright (c) 2017, Paul R. Tagliamonte <paultag@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

package main

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func SetExpiry(c *cli.Context) error {
	store, err := NewStore(c)
	if err != nil {
		return err
	}

	client, closer, err := NewClient(c, store)
	if err != nil {
		return err
	}
	defer closer()

	var expiryTime *time.Time = nil

	args := c.Args()
	if len(args) > 1 {
		expiry := args[1]

		when, err := strconv.Atoi(expiry)
		if err != nil {
			return err
		}
		_expiryTime := time.Unix(int64(when), 0)
		expiryTime = &_expiryTime
	}

	id, err := hex.DecodeString(args[0])
	if err != nil {
		return err
	}

	return client.SetEntityExpiry(id, expiryTime)

}

var SetExpiryCommand = cli.Command{
	Name:     "set-expiry",
	Action:   Wrapper(SetExpiry),
	Category: "admin",
	Usage:    "Set an entity's expiry date",
}

// vim: foldmethod=marker
