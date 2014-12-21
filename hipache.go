package main

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type Hipache struct {
	Redis redis.Conn
}

func NewHipache(host string) (Hipache, error) {
	conn, err := redis.Dial("tcp", host)
	return Hipache{conn}, err
}

func (h Hipache) Close() {
	h.Redis.Close()
}

func (h Hipache) Frontends() ([]string, error) {
	keys, err := redis.Strings(h.Redis.Do("KEYS", "frontend:*"))

	if err != nil {
		return []string{}, err
	}

	for i, val := range keys {
		keys[i] = strings.Replace(val, "frontend:", "", 1)
	}

	return keys, nil
}

func (h Hipache) Backends(frontend string) ([]string, error) {
	values, err := redis.Strings(h.Redis.Do("LRANGE", "frontend:"+frontend, "1", "-1"))

	if err != nil {
		return []string{}, err
	}

	return values, nil
}

func (h Hipache) AddFrontend(host string) error {
	return h.Redis.Send("RPUSH", "frontend:"+host, host)
}

func (h Hipache) RemoveFrontend(frontend string) error {
	return h.Redis.Send("DEL", "frontend:"+frontend)
}

func (h Hipache) AddBackend(frontend string, backend string) error {
	return h.Redis.Send("RPUSH", "frontend:"+frontend, backend)
}

func (h Hipache) RemoveBackend(frontend string, backend string) error {
	return h.Redis.Send("LREM", "frontend:"+frontend, "0", backend)
}

func (h Hipache) FrontendExists(frontend string) (bool, error) {
	return redis.Bool(h.Redis.Do("EXISTS", "frontend:"+frontend))
}

func (h Hipache) Flush() error {
	frontends, err := h.Frontends()

	if err != nil {
		return err
	}

	for _, fe := range frontends {
		err = h.Redis.Send("DEL", "frontend:"+fe)

		if err != nil {
			fmt.Println("ERROR:", err)
		}
	}

	return nil
}
