package main

import (
	"encoding/json"
	"fmt"
	"time"
)


type Time time.Time

func (t *Time) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	tValue, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.UTC)
	if err != nil {
		tValue, err = time.ParseInLocation("2006-01-02", s, time.UTC)
	}
	if err != nil {
		tValue, err = time.Parse(time.RFC3339, s)
	}
	if err != nil {
		return err
	}
	*t = Time(tValue)
	fmt.Printf("parse %v %v\n", *t, tValue)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	dt := time.Time(t)
	return json.Marshal(dt)
}

func main() {
	t1 := Time{}
	now := time.Now()
	nowBt, _ := now.MarshalJSON()
	err := json.Unmarshal(nowBt, &t1)
	//err := t1.UnmarshalJSON(nowBt)
	t1b, e2 := json.Marshal(t1)
	t2 := time.Time(t1)
	fmt.Printf("now:%v, err:%v, nows:%s, t1:%s, e2:%v\n", now, err, string(nowBt), string(t1b), e2)
	fmt.Printf("t2: %v\n", t2)
}