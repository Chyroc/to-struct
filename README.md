# to-struct
anything to go struct

## install
```
go get -u github.com/Chyroc/to-struct
```

## use

### json

#### map

```bash
cat ./testdata/map_1.json | ./to-struct
```

```go
type Struct struct {
	Code  int
	Count int
	Result []struct {
		Aid      int
		ArtistID int
		Lrc      string
		Sid      int
		Song     string
	}
}
```

#### slice

```bash
cat ./testdata/slice_1.json | ./to-struct
```

```go
type Struct struct {
	Slice []struct {
		Aid      int
		ArtistID int
		Lrc      string
		Sid      int
		Song     string
	}
}
```

### toml

```bash
cat cat ./testdata/toml_1.toml | ./to-struct
```

```go
type Struct struct {
	Clients struct {
		Data  [][]string
		Hosts []string
	}
	Database struct {
		ConnectionMax int64
		Enabled       bool
		Ports         []int64
		Server        string
	}
	Owner struct {
		Dob  time.Time
		Name string
	}
	Servers struct {
		Alpha struct {
			Dc string
			IP string
		}
		Beta struct {
			Dc string
			IP string
		}
	}
	Title string
}
```