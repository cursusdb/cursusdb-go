// Package cursusdbgo
// Go Native Client Tests
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
	"net"
	"net/textproto"
	"testing"
	"time"
)

func TestClient_Close(t *testing.T) {
	type fields struct {
		TLS                bool
		ClusterHost        string
		ClusterPort        uint
		Username           string
		Password           string
		Text               *textproto.Conn
		Conn               net.Conn
		ClusterReadTimeout time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// more test cases if you want
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				TLS:                tt.fields.TLS,
				ClusterHost:        tt.fields.ClusterHost,
				ClusterPort:        tt.fields.ClusterPort,
				Username:           tt.fields.Username,
				Password:           tt.fields.Password,
				Text:               tt.fields.Text,
				Conn:               tt.fields.Conn,
				ClusterReadTimeout: tt.fields.ClusterReadTimeout,
			}
			client.Close()
		})
	}
}

func TestClient_Connect(t *testing.T) {
	type fields struct {
		TLS                bool
		ClusterHost        string
		ClusterPort        uint
		Username           string
		Password           string
		Text               *textproto.Conn
		Conn               net.Conn
		ClusterReadTimeout time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// more test cases if you want
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				TLS:                tt.fields.TLS,
				ClusterHost:        tt.fields.ClusterHost,
				ClusterPort:        tt.fields.ClusterPort,
				Username:           tt.fields.Username,
				Password:           tt.fields.Password,
				Text:               tt.fields.Text,
				Conn:               tt.fields.Conn,
				ClusterReadTimeout: tt.fields.ClusterReadTimeout,
			}
			if err := client.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Query(t *testing.T) {
	type fields struct {
		TLS                bool
		ClusterHost        string
		ClusterPort        uint
		Username           string
		Password           string
		Text               *textproto.Conn
		Conn               net.Conn
		ClusterReadTimeout time.Time
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// more test cases if you want
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				TLS:                tt.fields.TLS,
				ClusterHost:        tt.fields.ClusterHost,
				ClusterPort:        tt.fields.ClusterPort,
				Username:           tt.fields.Username,
				Password:           tt.fields.Password,
				Text:               tt.fields.Text,
				Conn:               tt.fields.Conn,
				ClusterReadTimeout: tt.fields.ClusterReadTimeout,
			}
			got, err := client.Query(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
		})
	}
}
