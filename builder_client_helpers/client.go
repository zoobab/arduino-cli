/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 BCMI LABS SA (http://www.arduino.cc/)
 */

package builderClient

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the ArduinoBuilderV1 service client.
type Client struct {
	*goaclient.Client
	Oauth2Signer goaclient.Signer
	Encoder      *goa.HTTPEncoder
	Decoder      *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetOauth2Signer sets the request signer for the oauth2 security scheme.
func (c *Client) SetOauth2Signer(signer goaclient.Signer) {
	c.Oauth2Signer = signer
}
