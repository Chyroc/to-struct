package main

import (
	"testing"
	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestJSONToStruct(t *testing.T) {
	var in = `{
    "count": 15,
    "code": 0,
    "result": [
        { "aid": 2848529, "lrc": "http://s.gecimi.com/lrc/344/34435/3443588.lrc", "sid": 3443588, "artist_id": 2, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2346662, "lrc": "http://s.gecimi.com/lrc/274/27442/2744281.lrc", "sid": 2744281, "artist_id": 2396, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1889264, "lrc": "http://s.gecimi.com/lrc/210/21070/2107014.lrc", "sid": 2107014, "artist_id": 8715, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2075717, "lrc": "http://s.gecimi.com/lrc/236/23651/2365157.lrc", "sid": 2365157, "artist_id": 8715, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1563419, "lrc": "http://s.gecimi.com/lrc/166/16685/1668536.lrc", "sid": 1668536, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1567586, "lrc": "http://s.gecimi.com/lrc/167/16739/1673997.lrc", "sid": 1673997, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1571906, "lrc": "http://s.gecimi.com/lrc/167/16796/1679605.lrc", "sid": 1679605, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1573814, "lrc": "http://s.gecimi.com/lrc/168/16819/1681961.lrc", "sid": 1681961, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1656038, "lrc": "http://s.gecimi.com/lrc/179/17907/1790768.lrc", "sid": 1790768, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 1718741, "lrc": "http://s.gecimi.com/lrc/187/18757/1875769.lrc", "sid": 1875769, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2003267, "lrc": "http://s.gecimi.com/lrc/226/22642/2264296.lrc", "sid": 2264296, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2020610, "lrc": "http://s.gecimi.com/lrc/228/22889/2288967.lrc", "sid": 2288967, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2051678, "lrc": "http://s.gecimi.com/lrc/233/23323/2332322.lrc", "sid": 2332322, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2412704, "lrc": "http://s.gecimi.com/lrc/283/28376/2837689.lrc", "sid": 2837689, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" },
        { "aid": 2607041, "lrc": "http://s.gecimi.com/lrc/311/31116/3111659.lrc", "sid": 3111659, "artist_id": 9208, "song": "\u6d77\u9614\u5929\u7a7a" }
    ]
}`
	var buf = new(bytes.Buffer)
	JSONToStruct(buf, "", []byte(in))
	assert.Equal(t, `type Struct struct {
	Code	int
	Count	int
	Result	[]struct {
		Aid	int
		ArtistID	int
		Lrc	string
		Sid	int
		Song	string
	}
}`, buf.String())
}
