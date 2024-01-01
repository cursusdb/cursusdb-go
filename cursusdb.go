// Package cursusdbgo
// Go Native Client
// ///////////////////////////////////////////////////////////////////////
// Copyright (C) 2023 CursusDB
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package cursusdbgo

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/textproto"
	"strings"
	"time"
)

// Client is the CursusDB cluster client structure
type Client struct {
	TLS                bool            // TLS enabled?
	ClusterHost        string          // Cluster host
	ClusterPort        uint            // Cluster port
	Username           string          // Database username
	Password           string          // Database password
	Text               *textproto.Conn // Writer and reader
	Conn               net.Conn        // Conn
	ClusterReadTimeout time.Time       // Cluster read timeout
}

// Connect - Connect to new setup client to a CursusDB cluster
func (client *Client) Connect() error {
	if client.ClusterHost == "" || client.ClusterPort == 0 {
		return errors.New("CursusDB cluster host and port required.")
	}

	if !client.TLS {
		tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", client.ClusterHost, client.ClusterPort))
		if err != nil {
			return err
		}

		client.Conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return err
		}

		// If nothing from cluster in 5 seconds, report error
		err = client.Conn.SetReadDeadline(client.ClusterReadTimeout)
		if err != nil {
			return err
		}

		client.Text = textproto.NewConn(client.Conn)

		// Authenticate
		err = client.Text.PrintfLine(fmt.Sprintf("Authentication: %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\\0%s", client.Username, client.Password)))))
		if err != nil {
			return err
		}

		read, err := client.Text.ReadLine()
		if err != nil {
			return err
		}

		if strings.HasPrefix(read, fmt.Sprintf("%d ", 0)) {
			return nil
		} else {
			return errors.New("could not authenticate to cluster")
		}

	} else {
		var err error
		config := tls.Config{ServerName: client.ClusterHost}

		client.Conn, err = tls.Dial("tcp", fmt.Sprintf("%s:%d", client.ClusterHost, client.ClusterPort), &config)
		if err != nil {
			return err
		}

		// If nothing from cluster in 5 seconds, report error
		err = client.Conn.SetReadDeadline(client.ClusterReadTimeout)
		if err != nil {
			return err
		}

		client.Text = textproto.NewConn(client.Conn)
		// Authenticate
		err = client.Text.PrintfLine(fmt.Sprintf("Authentication: %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\\0%s", client.Username, client.Password)))))
		if err != nil {
			return err
		}

		read, err := client.Text.ReadLine()
		if err != nil {
			return err
		}

		if strings.HasPrefix(read, fmt.Sprintf("%d ", 0)) {
			return nil
		} else {
			return errors.New("could not authenticate to cluster")
		}
	}

	return nil

}

// Close closes CursusDB cluster connection
func (client *Client) Close() {
	client.Text.Close()
	client.Conn.Close()
}

// Query sends query to cluster
func (client *Client) Query(query string) (string, error) {
	if !strings.HasSuffix(query, ";") {
		return "", errors.New("invalid query")
	}

	_, err := client.Conn.Write([]byte(fmt.Sprintf("%s\r\n", query)))
	if err != nil {
		return "", err
	}

	read, err := client.Text.ReadLine()
	if err != nil {
		return "", err
	}

	return read, nil
}
