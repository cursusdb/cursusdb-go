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

type CursusDB struct {
	TLS         bool
	ClusterHost string
	ClusterPort uint
	Username    string
	Password    string
	Text        *textproto.Conn
	Conn        net.Conn
}

// NewClient - Create new client connection to a CursusDB cluster
func (cursusdb *CursusDB) NewClient() error {
	if cursusdb.ClusterHost == "" || cursusdb.ClusterPort == 0 {
		return errors.New("CursusDB cluster host and port required.")
	}

	if !cursusdb.TLS {
		tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", cursusdb.ClusterHost, cursusdb.ClusterPort))
		if err != nil {
			return err
		}

		cursusdb.Conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return err
		}

		// If nothing from cluster in 5 seconds, report error
		err = cursusdb.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			return err
		}

		cursusdb.Text = textproto.NewConn(cursusdb.Conn)

		// Authenticate
		err = cursusdb.Text.PrintfLine(fmt.Sprintf("Authentication: %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\\0%s", cursusdb.Username, cursusdb.Password)))))
		if err != nil {
			return err
		}

		read, err := cursusdb.Text.ReadLine()
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
		config := tls.Config{ServerName: cursusdb.ClusterHost}

		cursusdb.Conn, err = tls.Dial("tcp", fmt.Sprintf("%s:%d", cursusdb.ClusterHost, cursusdb.ClusterPort), &config)
		if err != nil {
			return err
		}

		// If nothing from cluster in 5 seconds, report error
		err = cursusdb.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			return err
		}

		cursusdb.Text = textproto.NewConn(cursusdb.Conn)
		// Authenticate
		err = cursusdb.Text.PrintfLine(fmt.Sprintf("Authentication: %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s\\0%s", cursusdb.Username, cursusdb.Password)))))
		if err != nil {
			return err
		}

		read, err := cursusdb.Text.ReadLine()
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
func (cursusdb *CursusDB) Close() {
	cursusdb.Text.Close()
	cursusdb.Conn.Close()
}

func (cursusdb *CursusDB) Query(query string) (string, error) {
	if !strings.HasSuffix(query, ";") {
		return "", errors.New("invalid query")
	}

	_, err := cursusdb.Conn.Write([]byte(fmt.Sprintf("%s\r\n", query)))
	if err != nil {
		return "", err
	}

	read, err := cursusdb.Text.ReadLine()
	if err != nil {
		return "", err
	}

	return read, nil
}
