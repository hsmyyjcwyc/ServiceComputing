package server

import (
    "github.com/o-martini/martini" 
)

func NewServer(port string) {   
    m := martini.Classic()
    m.Get("/", func(params martini.Params) string {
        return "hey, this is Tomy, 15331151."
    })

    m.RunOnAddr(":"+port)   
}